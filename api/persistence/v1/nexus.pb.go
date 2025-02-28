// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Copyright (c) 2019 Temporal Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by protoc-gen-go. DO NOT EDIT.
// plugins:
// 	protoc-gen-go
// 	protoc
// source: temporal/server/api/persistence/v1/nexus.proto

package persistence

import (
	reflect "reflect"
	sync "sync"

	v1 "go.temporal.io/server/api/clock/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NexusIncomingService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The last recorded cluster-local Hybrid Logical Clock timestamp for _this_ service.
	// Updated whenever the service is directly updated due to a user action but not when applying replication events.
	// The clock is referenced when new timestamps are generated to ensure it produces monotonically increasing
	// timestamps.
	Clock *v1.HybridLogicalClock `protobuf:"bytes,1,opt,name=clock,proto3" json:"clock,omitempty"`
	// Name of this service.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// UUID of the namespace to dispatch service requests to.
	NamespaceId string `protobuf:"bytes,3,opt,name=namespace_id,json=namespaceId,proto3" json:"namespace_id,omitempty"`
	// Task queue to dispatch service requests to.
	TaskQueue string `protobuf:"bytes,4,opt,name=task_queue,json=taskQueue,proto3" json:"task_queue,omitempty"`
	// Arbitrary user provided data. For use in the authorizer.
	Metadata map[string]*anypb.Any `protobuf:"bytes,5,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *NexusIncomingService) Reset() {
	*x = NexusIncomingService{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_server_api_persistence_v1_nexus_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NexusIncomingService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NexusIncomingService) ProtoMessage() {}

func (x *NexusIncomingService) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_server_api_persistence_v1_nexus_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NexusIncomingService.ProtoReflect.Descriptor instead.
func (*NexusIncomingService) Descriptor() ([]byte, []int) {
	return file_temporal_server_api_persistence_v1_nexus_proto_rawDescGZIP(), []int{0}
}

func (x *NexusIncomingService) GetClock() *v1.HybridLogicalClock {
	if x != nil {
		return x.Clock
	}
	return nil
}

func (x *NexusIncomingService) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NexusIncomingService) GetNamespaceId() string {
	if x != nil {
		return x.NamespaceId
	}
	return ""
}

func (x *NexusIncomingService) GetTaskQueue() string {
	if x != nil {
		return x.TaskQueue
	}
	return ""
}

func (x *NexusIncomingService) GetMetadata() map[string]*anypb.Any {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// Container for a version, a UUID, and a NexusIncomingService.
type NexusIncomingServiceEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version int64                 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Id      string                `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Service *NexusIncomingService `protobuf:"bytes,3,opt,name=service,proto3" json:"service,omitempty"`
}

func (x *NexusIncomingServiceEntry) Reset() {
	*x = NexusIncomingServiceEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_server_api_persistence_v1_nexus_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NexusIncomingServiceEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NexusIncomingServiceEntry) ProtoMessage() {}

func (x *NexusIncomingServiceEntry) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_server_api_persistence_v1_nexus_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NexusIncomingServiceEntry.ProtoReflect.Descriptor instead.
func (*NexusIncomingServiceEntry) Descriptor() ([]byte, []int) {
	return file_temporal_server_api_persistence_v1_nexus_proto_rawDescGZIP(), []int{1}
}

func (x *NexusIncomingServiceEntry) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *NexusIncomingServiceEntry) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NexusIncomingServiceEntry) GetService() *NexusIncomingService {
	if x != nil {
		return x.Service
	}
	return nil
}

var File_temporal_server_api_persistence_v1_nexus_proto protoreflect.FileDescriptor

var file_temporal_server_api_persistence_v1_nexus_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x6e, 0x65, 0x78, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x74,
	0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a,
	0x2a, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x87, 0x03, 0x0a, 0x14, 0x4e, 0x65, 0x78, 0x75, 0x73, 0x49, 0x6e, 0x63, 0x6f,
	0x6d, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x05, 0x63, 0x6c,
	0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f,
	0x72, 0x61, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6c,
	0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x79, 0x62, 0x72, 0x69, 0x64, 0x4c, 0x6f, 0x67, 0x69,
	0x63, 0x61, 0x6c, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x05, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x42,
	0x02, 0x68, 0x00, 0x12, 0x16, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x02, 0x68, 0x00, 0x12, 0x25, 0x0a, 0x0c, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x42, 0x02, 0x68, 0x00, 0x12, 0x21,
	0x0a, 0x0a, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x74, 0x61, 0x73, 0x6b, 0x51, 0x75, 0x65, 0x75, 0x65, 0x42, 0x02, 0x68, 0x00, 0x12,
	0x66, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x46, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x78, 0x75, 0x73, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x69, 0x6e,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x02,
	0x68, 0x00, 0x1a, 0x59, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x14, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x42, 0x02, 0x68, 0x00, 0x12, 0x2e, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x42, 0x02, 0x68, 0x00, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xa5, 0x01, 0x0a, 0x19, 0x4e, 0x65, 0x78,
	0x75, 0x73, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x02, 0x68,
	0x00, 0x12, 0x12, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x42, 0x02, 0x68, 0x00, 0x12, 0x56, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73,
	0x74, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x78, 0x75, 0x73, 0x49, 0x6e, 0x63,
	0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x07, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x42, 0x02, 0x68, 0x00, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x6f, 0x2e, 0x74, 0x65,
	0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x69, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x2f,
	0x76, 0x31, 0x3b, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_temporal_server_api_persistence_v1_nexus_proto_rawDescOnce sync.Once
	file_temporal_server_api_persistence_v1_nexus_proto_rawDescData = file_temporal_server_api_persistence_v1_nexus_proto_rawDesc
)

func file_temporal_server_api_persistence_v1_nexus_proto_rawDescGZIP() []byte {
	file_temporal_server_api_persistence_v1_nexus_proto_rawDescOnce.Do(func() {
		file_temporal_server_api_persistence_v1_nexus_proto_rawDescData = protoimpl.X.CompressGZIP(file_temporal_server_api_persistence_v1_nexus_proto_rawDescData)
	})
	return file_temporal_server_api_persistence_v1_nexus_proto_rawDescData
}

var file_temporal_server_api_persistence_v1_nexus_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_temporal_server_api_persistence_v1_nexus_proto_goTypes = []interface{}{
	(*NexusIncomingService)(nil),      // 0: temporal.server.api.persistence.v1.NexusIncomingService
	(*NexusIncomingServiceEntry)(nil), // 1: temporal.server.api.persistence.v1.NexusIncomingServiceEntry
	nil,                               // 2: temporal.server.api.persistence.v1.NexusIncomingService.MetadataEntry
	(*v1.HybridLogicalClock)(nil),     // 3: temporal.server.api.clock.v1.HybridLogicalClock
	(*anypb.Any)(nil),                 // 4: google.protobuf.Any
}
var file_temporal_server_api_persistence_v1_nexus_proto_depIdxs = []int32{
	3, // 0: temporal.server.api.persistence.v1.NexusIncomingService.clock:type_name -> temporal.server.api.clock.v1.HybridLogicalClock
	2, // 1: temporal.server.api.persistence.v1.NexusIncomingService.metadata:type_name -> temporal.server.api.persistence.v1.NexusIncomingService.MetadataEntry
	0, // 2: temporal.server.api.persistence.v1.NexusIncomingServiceEntry.service:type_name -> temporal.server.api.persistence.v1.NexusIncomingService
	4, // 3: temporal.server.api.persistence.v1.NexusIncomingService.MetadataEntry.value:type_name -> google.protobuf.Any
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_temporal_server_api_persistence_v1_nexus_proto_init() }
func file_temporal_server_api_persistence_v1_nexus_proto_init() {
	if File_temporal_server_api_persistence_v1_nexus_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_temporal_server_api_persistence_v1_nexus_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NexusIncomingService); i {
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
		file_temporal_server_api_persistence_v1_nexus_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NexusIncomingServiceEntry); i {
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
			RawDescriptor: file_temporal_server_api_persistence_v1_nexus_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_temporal_server_api_persistence_v1_nexus_proto_goTypes,
		DependencyIndexes: file_temporal_server_api_persistence_v1_nexus_proto_depIdxs,
		MessageInfos:      file_temporal_server_api_persistence_v1_nexus_proto_msgTypes,
	}.Build()
	File_temporal_server_api_persistence_v1_nexus_proto = out.File
	file_temporal_server_api_persistence_v1_nexus_proto_rawDesc = nil
	file_temporal_server_api_persistence_v1_nexus_proto_goTypes = nil
	file_temporal_server_api_persistence_v1_nexus_proto_depIdxs = nil
}
