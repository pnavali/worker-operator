/*
 *  Copyright (c) 2022 Avesha, Inc. All rights reserved.
 *
 *  SPDX-License-Identifier: Apache-2.0
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

package slice

import (
	"context"
	"fmt"
	"os"
	"time"

	ossEvents "github.com/kubeslice/worker-operator/events"

	kubeslicev1beta1 "github.com/kubeslice/worker-operator/api/v1beta1"
	nsmv1 "github.com/networkservicemesh/sdk-k8s/pkg/tools/k8s/apis/networkservicemesh.io/v1"

	"github.com/kubeslice/worker-operator/controllers"
	"github.com/kubeslice/worker-operator/pkg/logger"
	"github.com/kubeslice/worker-operator/pkg/utils"
	webhook "github.com/kubeslice/worker-operator/pkg/webhook/pod"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

const (
	nsmDataplaneVpp                 string = "vpp"
	nsmDataplaneKernel              string = "kernel"
	nsmVppDataplaneCfgStr           string = "nsm-vpp-plane"
	nsmKernelDataplaneCfgStr        string = "nsm-kernel-plane"
	sliceRouterDeploymentNamePrefix string = "vl3-slice-router-"
)

func labelsForSliceRouterDeployment(name string) map[string]string {
	return map[string]string{
		"networkservicemesh.io/app":                      "vl3-nse-" + name,
		"networkservicemesh.io/impl":                     "vl3-service-" + name,
		"spiffe.io/spiffe-id":                            "true",
		webhook.PodInjectLabelKey:                        "router",
		controllers.ApplicationNamespaceSelectorLabelKey: name,
	}
}

func getSliceRouterSidecarImageAndPullPolicy() (string, corev1.PullPolicy) {
	pullPolicy := corev1.PullAlways

	sliceRouterSidecarImage := os.Getenv("AVESHA_VL3_SIDECAR_IMAGE")
	sliceRouterSidecarImagePullPolicy := os.Getenv("AVESHA_VL3_SIDECAR_IMAGE_PULLPOLICY")

	if len(sliceRouterSidecarImagePullPolicy) > 0 {
		pullPolicy = corev1.PullPolicy(sliceRouterSidecarImagePullPolicy)
	}

	return sliceRouterSidecarImage, pullPolicy
}

func (r *SliceReconciler) getNsmDataplaneMode(ctx context.Context, slice *kubeslicev1beta1.Slice) (string, error) {
	log := logger.FromContext(ctx).WithValues("type", "SliceRouter")

	vppPodList := &corev1.PodList{}
	listOpts := []client.ListOption{
		client.InNamespace(slice.Namespace),
		client.MatchingLabels{"app": nsmVppDataplaneCfgStr},
	}
	if err := r.List(ctx, vppPodList, listOpts...); err != nil {
		log.Error(err, "Failed to list nsm vpp dataplane pods")
		return "", err
	}
	if len(vppPodList.Items) > 0 {
		return nsmDataplaneVpp, nil
	}

	return nsmDataplaneKernel, nil
}

func (r *SliceReconciler) getContainerSpecForSliceRouter(s *kubeslicev1beta1.Slice, dataplane string) corev1.Container {
	vl3ImagePullPolicy := corev1.PullAlways

	vl3Image := os.Getenv("AVESHA_VL3_ROUTER_IMAGE")
	vl3RouterPullPolicy := os.Getenv("AVESHA_VL3_ROUTER_PULLPOLICY")

	if len(vl3RouterPullPolicy) != 0 {
		vl3ImagePullPolicy = corev1.PullPolicy(vl3RouterPullPolicy)
	}

	privileged := true

	sliceRouterContainer := corev1.Container{
		Image:           vl3Image,
		Name:            "vl3-nse",
		ImagePullPolicy: vl3ImagePullPolicy,
		Env: []corev1.EnvVar{
			{
				Name:  "SPIFFE_ENDPOINT_SOCKET",
				Value: "unix:///run/spire/sockets/agent.sock",
			},
			{
				Name:  "NSM_CONNECT_TO",
				Value: "unix:///var/lib/networkservicemesh/nsm.io.sock",
			},
			{
				Name:  "NSM_SERVICE_NAMES",
				Value: "vl3-service-" + s.Name,
			},
			{
				Name:  "NSM_LABELS",
				Value: "app:" + "vl3-nse-" + s.Name,
			},
			{
				Name:  "NSM_NAME",
				Value: "vl3-nse-" + s.Name,
			},
		},
		SecurityContext: &corev1.SecurityContext{
			Privileged: &privileged,
		},
	}

	dnsConfig := fmt.Sprintf("[{\"dns_server_ips\": [\"%s\"], \"search_domains\": [\"slice.local\"]}]", s.Status.DNSIP)

	if dataplane == nsmDataplaneKernel {
		sliceRouterContainer.Env = append(sliceRouterContainer.Env,
			corev1.EnvVar{
				Name:  "NSM_VL3_PREFIX",
				Value: s.Status.SliceConfig.ClusterSubnetCIDR,
			},
			corev1.EnvVar{
				Name:  "NSM_DNS_CONFIGS",
				Value: dnsConfig,
			},
		)
		sliceRouterContainer.VolumeMounts = append(sliceRouterContainer.VolumeMounts,
			corev1.VolumeMount{
				Name:      "spire-agent-socket",
				MountPath: "/run/spire/sockets",
			},
			corev1.VolumeMount{
				Name:      "nsm-socket",
				MountPath: "/var/lib/networkservicemesh",
			},
		)
	}

	if dataplane == nsmDataplaneVpp {
		sliceRouterContainer.Env = append(sliceRouterContainer.Env,
			corev1.EnvVar{
				Name:  "NSE_IPAM_UNIQUE_OCTET",
				Value: s.Status.SliceConfig.ClusterSubnetCIDR,
			},
		)
		sliceRouterContainer.VolumeMounts = append(sliceRouterContainer.VolumeMounts,
			corev1.VolumeMount{
				Name:      "universal-cnf-config-volume",
				MountPath: "/etc/universal-cnf/config.yaml",
				SubPath:   "config.yaml",
			},
		)
	}

	return sliceRouterContainer
}

func (r *SliceReconciler) getContainerSpecForSliceRouterSidecar(dataplane string) corev1.Container {
	vl3SidecarImage, vl3SidecarImagePullPolicy := getSliceRouterSidecarImageAndPullPolicy()

	privileged := true

	sliceRouterSidecarContainer := corev1.Container{
		Name:            "kubeslice-vl3-sidecar",
		Image:           vl3SidecarImage,
		ImagePullPolicy: vl3SidecarImagePullPolicy,
		Env: []corev1.EnvVar{
			{
				Name:  "DATAPLANE",
				Value: dataplane,
			},
			{
				Name:  "POD_TYPE",
				Value: "SLICEROUTER_POD",
			},
		},
		SecurityContext: &corev1.SecurityContext{
			Privileged:               &privileged,
			AllowPrivilegeEscalation: &privileged,
			Capabilities: &corev1.Capabilities{
				Add: []corev1.Capability{
					"NET_ADMIN",
				},
			},
		},
		VolumeMounts: []corev1.VolumeMount{
			{
				Name:      "shared-volume",
				MountPath: "/config",
			},
		},
	}

	return sliceRouterSidecarContainer

}

func (r *SliceReconciler) getVolumeSpecForSliceRouter(s *kubeslicev1beta1.Slice, dataplane string) []corev1.Volume {
	var spireHostPathType corev1.HostPathType = "Directory"
	var nsmHostPathType corev1.HostPathType = "DirectoryOrCreate"
	sliceRouterVolumeSpec := []corev1.Volume{{
		Name: "shared-volume",
		VolumeSource: corev1.VolumeSource{
			EmptyDir: &corev1.EmptyDirVolumeSource{},
		},
	},
		{
			Name: "spire-agent-socket",
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: "/run/spire/sockets",
					Type: &spireHostPathType,
				},
			},
		},
		{
			Name: "nsm-socket",
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: "/var/lib/networkservicemesh",
					Type: &nsmHostPathType,
				},
			},
		},
	}

	if dataplane == nsmDataplaneVpp {
		sliceRouterVolumeSpec = append(sliceRouterVolumeSpec, corev1.Volume{
			Name: "universal-cnf-config-volume",
			VolumeSource: corev1.VolumeSource{
				ConfigMap: &corev1.ConfigMapVolumeSource{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: "ucnf-vl3-service-" + s.Name,
					},
				},
			},
		},
		)
	}

	return sliceRouterVolumeSpec
}

// Creates a deployment spec for the vL3 slice router
func (r *SliceReconciler) deploymentForSliceRouter(s *kubeslicev1beta1.Slice, dataplane string) *appsv1.Deployment {
	var replicas int32 = 1

	ls := labelsForSliceRouterDeployment(s.Name)

	dep := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      sliceRouterDeploymentNamePrefix + s.Name,
			Namespace: s.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: ls,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: ls,
				},
				Spec: corev1.PodSpec{
					ServiceAccountName: "slice-router",
					Affinity: &corev1.Affinity{
						NodeAffinity: &corev1.NodeAffinity{
							RequiredDuringSchedulingIgnoredDuringExecution: &corev1.NodeSelector{
								NodeSelectorTerms: []corev1.NodeSelectorTerm{{
									MatchExpressions: []corev1.NodeSelectorRequirement{{
										Key:      controllers.NodeTypeSelectorLabelKey,
										Operator: corev1.NodeSelectorOpIn,
										Values:   []string{"gateway"},
									}},
								}},
							},
						},
					},

					Containers: []corev1.Container{
						r.getContainerSpecForSliceRouter(s, dataplane),
						r.getContainerSpecForSliceRouterSidecar(dataplane),
					},
					Volumes: r.getVolumeSpecForSliceRouter(s, dataplane),
					Tolerations: []corev1.Toleration{{
						Key:      controllers.NodeTypeSelectorLabelKey,
						Operator: "Equal",
						Effect:   "NoSchedule",
						Value:    "gateway",
					}, {
						Key:      controllers.NodeTypeSelectorLabelKey,
						Operator: "Equal",
						Effect:   "NoExecute",
						Value:    "gateway",
					}},
				},
			},
		},
	}

	if len(controllers.ImagePullSecretName) != 0 {
		dep.Spec.Template.Spec.ImagePullSecrets = []corev1.LocalObjectReference{{
			Name: controllers.ImagePullSecretName,
		}}
	}

	ctrl.SetControllerReference(s, dep, r.Scheme)
	return dep
}

// Deploys the vL3 slice router.
// The configmap needed for the NSE is created first before the NSE is launched.
func (r *SliceReconciler) deploySliceRouter(ctx context.Context, slice *kubeslicev1beta1.Slice) error {
	log := logger.FromContext(ctx).WithName("slice-router")

	dataplane, err := r.getNsmDataplaneMode(ctx, slice)
	if err != nil {
		log.Error(err, "Failed to get nsm dataplane mode. Cannot deploy slice router")
		return err
	}
	if dataplane != nsmDataplaneKernel && dataplane != nsmDataplaneVpp {
		return fmt.Errorf("invalid dataplane: %v", dataplane)
	}

	dep := r.deploymentForSliceRouter(slice, dataplane)
	err = r.Create(ctx, dep)
	if err != nil {
		log.Error(err, "Failed to create deployment for slice router")
		utils.RecordEvent(ctx, r.EventRecorder, slice, nil, ossEvents.EventSliceRouterDeploymentFailed, controllerName)
		return err
	}
	log.Info("Created deployment spec for slice router: ", "Name: ", slice.Name, "cluster subnet: ", slice.Status.SliceConfig.ClusterSubnetCIDR)
	return nil
}

// Deploys the vL3 slice router service
func (r *SliceReconciler) deploySliceRouterSvc(ctx context.Context, slice *kubeslicev1beta1.Slice) error {
	log := logger.FromContext(ctx).WithName("slice-router-svc")

	ls := labelsForSliceRouterDeployment(slice.Name)

	svc := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      sliceRouterDeploymentNamePrefix + slice.Name,
			Namespace: controllers.ControlPlaneNamespace,
		},
		Spec: corev1.ServiceSpec{
			Selector: ls,
			Ports: []corev1.ServicePort{{
				Port: 5000,
				Name: "grpc",
			}},
		},
	}
	if err := ctrl.SetControllerReference(slice, svc, r.Scheme); err != nil {
		return err
	}

	err := r.Create(ctx, svc)
	if err != nil {
		log.Error(err, "Failed to create svc for slice router")
		utils.RecordEvent(ctx, r.EventRecorder, slice, nil, ossEvents.EventSliceRouterServiceFailed, controllerName)
		return err
	}
	log.Info("Created svc spec for slice router: ", "Name: ", slice.Name)
	return nil
}

func (r *SliceReconciler) ReconcileSliceRouter(ctx context.Context, slice *kubeslicev1beta1.Slice) (ctrl.Result, error, bool) {
	log := logger.FromContext(ctx).WithName("slice-router")
	// Spawn the slice router for the slice if not done already
	foundSliceRouter := &appsv1.Deployment{}
	err := r.Get(ctx, types.NamespacedName{Name: sliceRouterDeploymentNamePrefix + slice.Name, Namespace: slice.Namespace}, foundSliceRouter)
	if err != nil {
		if errors.IsNotFound(err) {
			if !sliceConfigDefined(slice) {
				log.Info("Slice subnet config not available yet, cannot deploy slice router. Waiting...")
				return ctrl.Result{
					RequeueAfter: 10 * time.Second,
				}, nil, true
			}
			// Define and create a new deployment for the slice router
			return newDeploymentSliceRouter(r, ctx, slice)
		}
		return ctrl.Result{}, err, true
	}

	foundSvc := &corev1.Service{}
	err = r.Get(ctx, types.NamespacedName{
		Name:      sliceRouterDeploymentNamePrefix + slice.Name,
		Namespace: controllers.ControlPlaneNamespace,
	}, foundSvc)

	if err != nil {
		if errors.IsNotFound(err) {
			if slice.Status.SliceConfig == nil {
				return ctrl.Result{
					RequeueAfter: 10 * time.Second,
				}, nil, true
			}
			// Define and create a new service for the slice router
			return newServiceSliceRouter(r, ctx, slice)
		}
		return ctrl.Result{}, err, true
	}

	return ctrl.Result{}, nil, false
}

func newServiceSliceRouter(r *SliceReconciler, ctx context.Context, slice *kubeslicev1beta1.Slice) (reconcile.Result, error, bool) {
	log := logger.FromContext(ctx).WithName("slice-router")
	err := r.deploySliceRouterSvc(ctx, slice)
	if err != nil {
		log.Error(err, "Failed to deploy slice router service")
		return ctrl.Result{}, err, true
	}
	log.Info("Creating slice router", "Namespace svc", slice.Namespace, "Name", "vl3-nse-"+slice.Name)
	return ctrl.Result{
		RequeueAfter: 10 * time.Second,
	}, nil, true
}

func newDeploymentSliceRouter(r *SliceReconciler, ctx context.Context, slice *kubeslicev1beta1.Slice) (reconcile.Result, error, bool) {
	log := logger.FromContext(ctx).WithName("slice-router")
	err := r.deploySliceRouter(ctx, slice)
	if err != nil {
		log.Error(err, "Failed to deploy slice router")
		return ctrl.Result{}, err, true
	}
	log.Info("Creating slice router", "Namespace", slice.Namespace, "Name", "vl3-nse-"+slice.Name)
	return ctrl.Result{
		RequeueAfter: 10 * time.Second,
	}, nil, true
}
func sliceConfigDefined(slice *kubeslicev1beta1.Slice) bool {
	return slice.Status.SliceConfig != nil && slice.Status.SliceConfig.SliceSubnet != "" && slice.Status.SliceConfig.ClusterSubnetCIDR != ""
}
func (r *SliceReconciler) cleanupVl3NSE(ctx context.Context, sliceName string) error {
	log := logger.FromContext(ctx)

	vl3Nse := &nsmv1.NetworkService{}
	err := r.Get(ctx, types.NamespacedName{Name: "vl3-service-" + sliceName, Namespace: controllers.ControlPlaneNamespace}, vl3Nse)
	if err != nil {
		if errors.IsNotFound(err) {
			return nil
		}
		log.Error(err, "Slice router cleanup: Failed to get vl3 nse")
		return err
	}

	err = r.Delete(ctx, vl3Nse)
	if err != nil {
		log.Error(err, "Slice router cleanup: Failed to delete vl3 nse")
		return err
	}
	return nil
}

func (r *SliceReconciler) deleteSliceRouterComponentsIfPresent(ctx context.Context, slice *kubeslicev1beta1.Slice) error {
	log := logger.FromContext(ctx)
	// debug := log.V(1)
	// delete router deployment if present
	dep := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      sliceRouterDeploymentNamePrefix + slice.Name,
			Namespace: slice.Namespace,
		},
	}
	if err := r.Delete(ctx, dep); err != nil {
		if !errors.IsNotFound(err) {
			log.Error(err, "Failed to delete Slice router deployment")
			return err
		}
	}

	// delete router service if present
	svc := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      sliceRouterDeploymentNamePrefix + slice.Name,
			Namespace: controllers.ControlPlaneNamespace,
		},
	}
	if err := r.Delete(ctx, svc); err != nil {
		if !errors.IsNotFound(err) {
			log.Error(err, "Failed to delete Slice router service")
			return err
		}
	}

	// cleanup slice router network service if present
	if err := r.cleanupVl3NSE(ctx, slice.Name); err != nil {
		return err
	}

	return nil
}
