//  Copyright (c) 2022 Avesha, Inc. All rights reserved.
//
//  SPDX-License-Identifier: Apache-2.0
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.3
// source: edgeservice/edgeservice.proto

package edgeservice

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// GwEdgeResponse represents the Sidecar response format.
type GwEdgeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusMsg string `protobuf:"bytes,1,opt,name=statusMsg,proto3" json:"statusMsg,omitempty"`
}

func (x *GwEdgeResponse) Reset() {
	*x = GwEdgeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_edgeservice_edgeservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GwEdgeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GwEdgeResponse) ProtoMessage() {}

func (x *GwEdgeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_edgeservice_edgeservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GwEdgeResponse.ProtoReflect.Descriptor instead.
func (*GwEdgeResponse) Descriptor() ([]byte, []int) {
	return file_edgeservice_edgeservice_proto_rawDescGZIP(), []int{0}
}

func (x *GwEdgeResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

type SliceGwServiceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GwName          string `protobuf:"bytes,1,opt,name=GwName,proto3" json:"GwName,omitempty"`
	GwSvcName       string `protobuf:"bytes,2,opt,name=GwSvcName,proto3" json:"GwSvcName,omitempty"`
	GwSvcClusterIP  string `protobuf:"bytes,3,opt,name=GwSvcClusterIP,proto3" json:"GwSvcClusterIP,omitempty"`
	GwSvcTargetPort uint32 `protobuf:"varint,4,opt,name=GwSvcTargetPort,proto3" json:"GwSvcTargetPort,omitempty"`
	GwSvcNodePort   uint32 `protobuf:"varint,5,opt,name=GwSvcNodePort,proto3" json:"GwSvcNodePort,omitempty"`
}

func (x *SliceGwServiceInfo) Reset() {
	*x = SliceGwServiceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_edgeservice_edgeservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SliceGwServiceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SliceGwServiceInfo) ProtoMessage() {}

func (x *SliceGwServiceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_edgeservice_edgeservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SliceGwServiceInfo.ProtoReflect.Descriptor instead.
func (*SliceGwServiceInfo) Descriptor() ([]byte, []int) {
	return file_edgeservice_edgeservice_proto_rawDescGZIP(), []int{1}
}

func (x *SliceGwServiceInfo) GetGwName() string {
	if x != nil {
		return x.GwName
	}
	return ""
}

func (x *SliceGwServiceInfo) GetGwSvcName() string {
	if x != nil {
		return x.GwSvcName
	}
	return ""
}

func (x *SliceGwServiceInfo) GetGwSvcClusterIP() string {
	if x != nil {
		return x.GwSvcClusterIP
	}
	return ""
}

func (x *SliceGwServiceInfo) GetGwSvcTargetPort() uint32 {
	if x != nil {
		return x.GwSvcTargetPort
	}
	return 0
}

func (x *SliceGwServiceInfo) GetGwSvcNodePort() uint32 {
	if x != nil {
		return x.GwSvcNodePort
	}
	return 0
}

type SliceGwServiceMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SliceName          string                `protobuf:"bytes,1,opt,name=sliceName,proto3" json:"sliceName,omitempty"`
	SliceGwServiceList []*SliceGwServiceInfo `protobuf:"bytes,2,rep,name=sliceGwServiceList,proto3" json:"sliceGwServiceList,omitempty"`
}

func (x *SliceGwServiceMap) Reset() {
	*x = SliceGwServiceMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_edgeservice_edgeservice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SliceGwServiceMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SliceGwServiceMap) ProtoMessage() {}

func (x *SliceGwServiceMap) ProtoReflect() protoreflect.Message {
	mi := &file_edgeservice_edgeservice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SliceGwServiceMap.ProtoReflect.Descriptor instead.
func (*SliceGwServiceMap) Descriptor() ([]byte, []int) {
	return file_edgeservice_edgeservice_proto_rawDescGZIP(), []int{2}
}

func (x *SliceGwServiceMap) GetSliceName() string {
	if x != nil {
		return x.SliceName
	}
	return ""
}

func (x *SliceGwServiceMap) GetSliceGwServiceList() []*SliceGwServiceInfo {
	if x != nil {
		return x.SliceGwServiceList
	}
	return nil
}

var File_edgeservice_edgeservice_proto protoreflect.FileDescriptor

var file_edgeservice_edgeservice_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x65, 0x64, 0x67, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x65, 0x64,
	0x67, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x65, 0x64, 0x67, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x2e, 0x0a, 0x0e,
	0x47, 0x77, 0x45, 0x64, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x22, 0xc2, 0x01, 0x0a,
	0x12, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x47, 0x77, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x47, 0x77, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x47, 0x77, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x47,
	0x77, 0x53, 0x76, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x47, 0x77, 0x53, 0x76, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x47, 0x77, 0x53,
	0x76, 0x63, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x50, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x47, 0x77, 0x53, 0x76, 0x63, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49,
	0x50, 0x12, 0x28, 0x0a, 0x0f, 0x47, 0x77, 0x53, 0x76, 0x63, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x50, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x47, 0x77, 0x53, 0x76,
	0x63, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x47,
	0x77, 0x53, 0x76, 0x63, 0x4e, 0x6f, 0x64, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0d, 0x47, 0x77, 0x53, 0x76, 0x63, 0x4e, 0x6f, 0x64, 0x65, 0x50, 0x6f, 0x72,
	0x74, 0x22, 0x82, 0x01, 0x0a, 0x11, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x47, 0x77, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x4d, 0x61, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x6c, 0x69, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x6c, 0x69, 0x63,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x4f, 0x0a, 0x12, 0x73, 0x6c, 0x69, 0x63, 0x65, 0x47, 0x77,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x53, 0x6c, 0x69, 0x63, 0x65, 0x47, 0x77, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x12, 0x73, 0x6c, 0x69, 0x63, 0x65, 0x47, 0x77, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x32, 0x69, 0x0a, 0x0d, 0x47, 0x77, 0x45, 0x64, 0x67, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x58, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x47, 0x77, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d,
	0x61, 0x70, 0x12, 0x1e, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x47, 0x77, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d,
	0x61, 0x70, 0x1a, 0x1b, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x47, 0x77, 0x45, 0x64, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x3b, 0x65, 0x64, 0x67, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_edgeservice_edgeservice_proto_rawDescOnce sync.Once
	file_edgeservice_edgeservice_proto_rawDescData = file_edgeservice_edgeservice_proto_rawDesc
)

func file_edgeservice_edgeservice_proto_rawDescGZIP() []byte {
	file_edgeservice_edgeservice_proto_rawDescOnce.Do(func() {
		file_edgeservice_edgeservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_edgeservice_edgeservice_proto_rawDescData)
	})
	return file_edgeservice_edgeservice_proto_rawDescData
}

var file_edgeservice_edgeservice_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_edgeservice_edgeservice_proto_goTypes = []interface{}{
	(*GwEdgeResponse)(nil),     // 0: edgeservice.GwEdgeResponse
	(*SliceGwServiceInfo)(nil), // 1: edgeservice.SliceGwServiceInfo
	(*SliceGwServiceMap)(nil),  // 2: edgeservice.SliceGwServiceMap
}
var file_edgeservice_edgeservice_proto_depIdxs = []int32{
	1, // 0: edgeservice.SliceGwServiceMap.sliceGwServiceList:type_name -> edgeservice.SliceGwServiceInfo
	2, // 1: edgeservice.GwEdgeService.UpdateSliceGwServiceMap:input_type -> edgeservice.SliceGwServiceMap
	0, // 2: edgeservice.GwEdgeService.UpdateSliceGwServiceMap:output_type -> edgeservice.GwEdgeResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_edgeservice_edgeservice_proto_init() }
func file_edgeservice_edgeservice_proto_init() {
	if File_edgeservice_edgeservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_edgeservice_edgeservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GwEdgeResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_edgeservice_edgeservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SliceGwServiceInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_edgeservice_edgeservice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SliceGwServiceMap); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_edgeservice_edgeservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_edgeservice_edgeservice_proto_goTypes,
		DependencyIndexes: file_edgeservice_edgeservice_proto_depIdxs,
		MessageInfos:      file_edgeservice_edgeservice_proto_msgTypes,
	}.Build()
	File_edgeservice_edgeservice_proto = out.File
	file_edgeservice_edgeservice_proto_rawDesc = nil
	file_edgeservice_edgeservice_proto_goTypes = nil
	file_edgeservice_edgeservice_proto_depIdxs = nil
}