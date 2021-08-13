// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: isuxportal/services/contestant/instances.proto

package contestant

import (
	resources "github.com/isucon/isucon10-portal/proto.go/isuxportal/resources"
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

type ListContestantInstancesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListContestantInstancesRequest) Reset() {
	*x = ListContestantInstancesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isuxportal_services_contestant_instances_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListContestantInstancesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListContestantInstancesRequest) ProtoMessage() {}

func (x *ListContestantInstancesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_isuxportal_services_contestant_instances_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListContestantInstancesRequest.ProtoReflect.Descriptor instead.
func (*ListContestantInstancesRequest) Descriptor() ([]byte, []int) {
	return file_isuxportal_services_contestant_instances_proto_rawDescGZIP(), []int{0}
}

type ListContestantInstancesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContestantInstances []*resources.ContestantInstance `protobuf:"bytes,1,rep,name=contestant_instances,json=contestantInstances,proto3" json:"contestant_instances,omitempty"`
}

func (x *ListContestantInstancesResponse) Reset() {
	*x = ListContestantInstancesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isuxportal_services_contestant_instances_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListContestantInstancesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListContestantInstancesResponse) ProtoMessage() {}

func (x *ListContestantInstancesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_isuxportal_services_contestant_instances_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListContestantInstancesResponse.ProtoReflect.Descriptor instead.
func (*ListContestantInstancesResponse) Descriptor() ([]byte, []int) {
	return file_isuxportal_services_contestant_instances_proto_rawDescGZIP(), []int{1}
}

func (x *ListContestantInstancesResponse) GetContestantInstances() []*resources.ContestantInstance {
	if x != nil {
		return x.ContestantInstances
	}
	return nil
}

var File_isuxportal_services_contestant_instances_proto protoreflect.FileDescriptor

var file_isuxportal_services_contestant_instances_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x69, 0x73, 0x75, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x74,
	0x2f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x24, 0x69, 0x73, 0x75, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x1a, 0x2e, 0x69, 0x73, 0x75, 0x78, 0x70, 0x6f, 0x72, 0x74,
	0x61, 0x6c, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x20, 0x0a, 0x1e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x84, 0x01, 0x0a, 0x1f, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x61, 0x0a, 0x14,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x69, 0x73, 0x75,
	0x78, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x61,
	0x6e, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x13, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x42,
	0x4b, 0x5a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x73,
	0x75, 0x63, 0x6f, 0x6e, 0x2f, 0x69, 0x73, 0x75, 0x63, 0x6f, 0x6e, 0x31, 0x30, 0x2d, 0x70, 0x6f,
	0x72, 0x74, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x6f, 0x2f, 0x69, 0x73,
	0x75, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_isuxportal_services_contestant_instances_proto_rawDescOnce sync.Once
	file_isuxportal_services_contestant_instances_proto_rawDescData = file_isuxportal_services_contestant_instances_proto_rawDesc
)

func file_isuxportal_services_contestant_instances_proto_rawDescGZIP() []byte {
	file_isuxportal_services_contestant_instances_proto_rawDescOnce.Do(func() {
		file_isuxportal_services_contestant_instances_proto_rawDescData = protoimpl.X.CompressGZIP(file_isuxportal_services_contestant_instances_proto_rawDescData)
	})
	return file_isuxportal_services_contestant_instances_proto_rawDescData
}

var file_isuxportal_services_contestant_instances_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_isuxportal_services_contestant_instances_proto_goTypes = []interface{}{
	(*ListContestantInstancesRequest)(nil),  // 0: isuxportal.proto.services.contestant.ListContestantInstancesRequest
	(*ListContestantInstancesResponse)(nil), // 1: isuxportal.proto.services.contestant.ListContestantInstancesResponse
	(*resources.ContestantInstance)(nil),    // 2: isuxportal.proto.resources.ContestantInstance
}
var file_isuxportal_services_contestant_instances_proto_depIdxs = []int32{
	2, // 0: isuxportal.proto.services.contestant.ListContestantInstancesResponse.contestant_instances:type_name -> isuxportal.proto.resources.ContestantInstance
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_isuxportal_services_contestant_instances_proto_init() }
func file_isuxportal_services_contestant_instances_proto_init() {
	if File_isuxportal_services_contestant_instances_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_isuxportal_services_contestant_instances_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListContestantInstancesRequest); i {
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
		file_isuxportal_services_contestant_instances_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListContestantInstancesResponse); i {
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
			RawDescriptor: file_isuxportal_services_contestant_instances_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_isuxportal_services_contestant_instances_proto_goTypes,
		DependencyIndexes: file_isuxportal_services_contestant_instances_proto_depIdxs,
		MessageInfos:      file_isuxportal_services_contestant_instances_proto_msgTypes,
	}.Build()
	File_isuxportal_services_contestant_instances_proto = out.File
	file_isuxportal_services_contestant_instances_proto_rawDesc = nil
	file_isuxportal_services_contestant_instances_proto_goTypes = nil
	file_isuxportal_services_contestant_instances_proto_depIdxs = nil
}
