// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: isuxportal/services/audience/team_list.proto

package audience

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ListTeamsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListTeamsRequest) Reset() {
	*x = ListTeamsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isuxportal_services_audience_team_list_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTeamsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTeamsRequest) ProtoMessage() {}

func (x *ListTeamsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_isuxportal_services_audience_team_list_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTeamsRequest.ProtoReflect.Descriptor instead.
func (*ListTeamsRequest) Descriptor() ([]byte, []int) {
	return file_isuxportal_services_audience_team_list_proto_rawDescGZIP(), []int{0}
}

type ListTeamsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Teams []*ListTeamsResponse_TeamListItem `protobuf:"bytes,1,rep,name=teams,proto3" json:"teams,omitempty"`
}

func (x *ListTeamsResponse) Reset() {
	*x = ListTeamsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isuxportal_services_audience_team_list_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTeamsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTeamsResponse) ProtoMessage() {}

func (x *ListTeamsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_isuxportal_services_audience_team_list_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTeamsResponse.ProtoReflect.Descriptor instead.
func (*ListTeamsResponse) Descriptor() ([]byte, []int) {
	return file_isuxportal_services_audience_team_list_proto_rawDescGZIP(), []int{1}
}

func (x *ListTeamsResponse) GetTeams() []*ListTeamsResponse_TeamListItem {
	if x != nil {
		return x.Teams
	}
	return nil
}

type ListTeamsResponse_TeamListItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamId             int64    `protobuf:"varint,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	Name               string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	MemberNames        []string `protobuf:"bytes,3,rep,name=member_names,json=memberNames,proto3" json:"member_names,omitempty"`
	FinalParticipation bool     `protobuf:"varint,4,opt,name=final_participation,json=finalParticipation,proto3" json:"final_participation,omitempty"`
	IsStudent          bool     `protobuf:"varint,5,opt,name=is_student,json=isStudent,proto3" json:"is_student,omitempty"`
}

func (x *ListTeamsResponse_TeamListItem) Reset() {
	*x = ListTeamsResponse_TeamListItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isuxportal_services_audience_team_list_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTeamsResponse_TeamListItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTeamsResponse_TeamListItem) ProtoMessage() {}

func (x *ListTeamsResponse_TeamListItem) ProtoReflect() protoreflect.Message {
	mi := &file_isuxportal_services_audience_team_list_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTeamsResponse_TeamListItem.ProtoReflect.Descriptor instead.
func (*ListTeamsResponse_TeamListItem) Descriptor() ([]byte, []int) {
	return file_isuxportal_services_audience_team_list_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ListTeamsResponse_TeamListItem) GetTeamId() int64 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

func (x *ListTeamsResponse_TeamListItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ListTeamsResponse_TeamListItem) GetMemberNames() []string {
	if x != nil {
		return x.MemberNames
	}
	return nil
}

func (x *ListTeamsResponse_TeamListItem) GetFinalParticipation() bool {
	if x != nil {
		return x.FinalParticipation
	}
	return false
}

func (x *ListTeamsResponse_TeamListItem) GetIsStudent() bool {
	if x != nil {
		return x.IsStudent
	}
	return false
}

var File_isuxportal_services_audience_team_list_proto protoreflect.FileDescriptor

var file_isuxportal_services_audience_team_list_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x69, 0x73, 0x75, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x2f, 0x74,
	0x65, 0x61, 0x6d, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22,
	0x69, 0x73, 0x75, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e,
	0x63, 0x65, 0x22, 0x12, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x9e, 0x02, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x54,
	0x65, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x58, 0x0a, 0x05,
	0x74, 0x65, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x42, 0x2e, 0x69, 0x73,
	0x75, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x05, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x1a, 0xae, 0x01, 0x0a, 0x0c, 0x54, 0x65, 0x61, 0x6d, 0x4c,
	0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x2f, 0x0a, 0x13, 0x66, 0x69, 0x6e, 0x61, 0x6c,
	0x5f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x50, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x69, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x73,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73,
	0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x73, 0x75, 0x63, 0x6f, 0x6e, 0x2f, 0x69, 0x73, 0x75,
	0x63, 0x6f, 0x6e, 0x31, 0x30, 0x2d, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x61,
	0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_isuxportal_services_audience_team_list_proto_rawDescOnce sync.Once
	file_isuxportal_services_audience_team_list_proto_rawDescData = file_isuxportal_services_audience_team_list_proto_rawDesc
)

func file_isuxportal_services_audience_team_list_proto_rawDescGZIP() []byte {
	file_isuxportal_services_audience_team_list_proto_rawDescOnce.Do(func() {
		file_isuxportal_services_audience_team_list_proto_rawDescData = protoimpl.X.CompressGZIP(file_isuxportal_services_audience_team_list_proto_rawDescData)
	})
	return file_isuxportal_services_audience_team_list_proto_rawDescData
}

var file_isuxportal_services_audience_team_list_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_isuxportal_services_audience_team_list_proto_goTypes = []interface{}{
	(*ListTeamsRequest)(nil),               // 0: isuxportal.proto.services.audience.ListTeamsRequest
	(*ListTeamsResponse)(nil),              // 1: isuxportal.proto.services.audience.ListTeamsResponse
	(*ListTeamsResponse_TeamListItem)(nil), // 2: isuxportal.proto.services.audience.ListTeamsResponse.TeamListItem
}
var file_isuxportal_services_audience_team_list_proto_depIdxs = []int32{
	2, // 0: isuxportal.proto.services.audience.ListTeamsResponse.teams:type_name -> isuxportal.proto.services.audience.ListTeamsResponse.TeamListItem
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_isuxportal_services_audience_team_list_proto_init() }
func file_isuxportal_services_audience_team_list_proto_init() {
	if File_isuxportal_services_audience_team_list_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_isuxportal_services_audience_team_list_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTeamsRequest); i {
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
		file_isuxportal_services_audience_team_list_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTeamsResponse); i {
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
		file_isuxportal_services_audience_team_list_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTeamsResponse_TeamListItem); i {
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
			RawDescriptor: file_isuxportal_services_audience_team_list_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_isuxportal_services_audience_team_list_proto_goTypes,
		DependencyIndexes: file_isuxportal_services_audience_team_list_proto_depIdxs,
		MessageInfos:      file_isuxportal_services_audience_team_list_proto_msgTypes,
	}.Build()
	File_isuxportal_services_audience_team_list_proto = out.File
	file_isuxportal_services_audience_team_list_proto_rawDesc = nil
	file_isuxportal_services_audience_team_list_proto_goTypes = nil
	file_isuxportal_services_audience_team_list_proto_depIdxs = nil
}
