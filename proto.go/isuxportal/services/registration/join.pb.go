// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: isuxportal/services/registration/join.proto

package registration

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

type JoinTeamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TeamId      int64  `protobuf:"varint,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	InviteToken string `protobuf:"bytes,2,opt,name=invite_token,json=inviteToken,proto3" json:"invite_token,omitempty"`
	Name        string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	IsStudent   bool   `protobuf:"varint,4,opt,name=is_student,json=isStudent,proto3" json:"is_student,omitempty"`
}

func (x *JoinTeamRequest) Reset() {
	*x = JoinTeamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isuxportal_services_registration_join_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinTeamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinTeamRequest) ProtoMessage() {}

func (x *JoinTeamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_isuxportal_services_registration_join_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinTeamRequest.ProtoReflect.Descriptor instead.
func (*JoinTeamRequest) Descriptor() ([]byte, []int) {
	return file_isuxportal_services_registration_join_proto_rawDescGZIP(), []int{0}
}

func (x *JoinTeamRequest) GetTeamId() int64 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

func (x *JoinTeamRequest) GetInviteToken() string {
	if x != nil {
		return x.InviteToken
	}
	return ""
}

func (x *JoinTeamRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *JoinTeamRequest) GetIsStudent() bool {
	if x != nil {
		return x.IsStudent
	}
	return false
}

type JoinTeamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *JoinTeamResponse) Reset() {
	*x = JoinTeamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isuxportal_services_registration_join_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinTeamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinTeamResponse) ProtoMessage() {}

func (x *JoinTeamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_isuxportal_services_registration_join_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinTeamResponse.ProtoReflect.Descriptor instead.
func (*JoinTeamResponse) Descriptor() ([]byte, []int) {
	return file_isuxportal_services_registration_join_proto_rawDescGZIP(), []int{1}
}

var File_isuxportal_services_registration_join_proto protoreflect.FileDescriptor

var file_isuxportal_services_registration_join_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x69, 0x73, 0x75, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x6a, 0x6f, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x26, 0x69,
	0x73, 0x75, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x80, 0x01, 0x0a, 0x0f, 0x4a, 0x6f, 0x69, 0x6e, 0x54, 0x65,
	0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x61,
	0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d,
	0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f,
	0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69,
	0x73, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x4a, 0x6f, 0x69, 0x6e,
	0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x4d, 0x5a, 0x4b,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x73, 0x75, 0x63, 0x6f,
	0x6e, 0x2f, 0x69, 0x73, 0x75, 0x63, 0x6f, 0x6e, 0x31, 0x30, 0x2d, 0x70, 0x6f, 0x72, 0x74, 0x61,
	0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x6f, 0x2f, 0x69, 0x73, 0x75, 0x78, 0x70,
	0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_isuxportal_services_registration_join_proto_rawDescOnce sync.Once
	file_isuxportal_services_registration_join_proto_rawDescData = file_isuxportal_services_registration_join_proto_rawDesc
)

func file_isuxportal_services_registration_join_proto_rawDescGZIP() []byte {
	file_isuxportal_services_registration_join_proto_rawDescOnce.Do(func() {
		file_isuxportal_services_registration_join_proto_rawDescData = protoimpl.X.CompressGZIP(file_isuxportal_services_registration_join_proto_rawDescData)
	})
	return file_isuxportal_services_registration_join_proto_rawDescData
}

var file_isuxportal_services_registration_join_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_isuxportal_services_registration_join_proto_goTypes = []interface{}{
	(*JoinTeamRequest)(nil),  // 0: isuxportal.proto.services.registration.JoinTeamRequest
	(*JoinTeamResponse)(nil), // 1: isuxportal.proto.services.registration.JoinTeamResponse
}
var file_isuxportal_services_registration_join_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_isuxportal_services_registration_join_proto_init() }
func file_isuxportal_services_registration_join_proto_init() {
	if File_isuxportal_services_registration_join_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_isuxportal_services_registration_join_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinTeamRequest); i {
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
		file_isuxportal_services_registration_join_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinTeamResponse); i {
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
			RawDescriptor: file_isuxportal_services_registration_join_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_isuxportal_services_registration_join_proto_goTypes,
		DependencyIndexes: file_isuxportal_services_registration_join_proto_depIdxs,
		MessageInfos:      file_isuxportal_services_registration_join_proto_msgTypes,
	}.Build()
	File_isuxportal_services_registration_join_proto = out.File
	file_isuxportal_services_registration_join_proto_rawDesc = nil
	file_isuxportal_services_registration_join_proto_goTypes = nil
	file_isuxportal_services_registration_join_proto_depIdxs = nil
}
