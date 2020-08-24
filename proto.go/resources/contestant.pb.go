// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: isuxportal/resources/contestant.proto

package resources

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

type Contestant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               int64                        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TeamId           int64                        `protobuf:"varint,2,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	Name             string                       `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	ContestantDetail *Contestant_ContestantDetail `protobuf:"bytes,7,opt,name=contestant_detail,json=contestantDetail,proto3" json:"contestant_detail,omitempty"` // FIXME: rename to "detail"
}

func (x *Contestant) Reset() {
	*x = Contestant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isuxportal_resources_contestant_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Contestant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contestant) ProtoMessage() {}

func (x *Contestant) ProtoReflect() protoreflect.Message {
	mi := &file_isuxportal_resources_contestant_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Contestant.ProtoReflect.Descriptor instead.
func (*Contestant) Descriptor() ([]byte, []int) {
	return file_isuxportal_resources_contestant_proto_rawDescGZIP(), []int{0}
}

func (x *Contestant) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Contestant) GetTeamId() int64 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

func (x *Contestant) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Contestant) GetContestantDetail() *Contestant_ContestantDetail {
	if x != nil {
		return x.ContestantDetail
	}
	return nil
}

type Contestant_ContestantDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GithubLogin string `protobuf:"bytes,1,opt,name=github_login,json=githubLogin,proto3" json:"github_login,omitempty"`
	DiscordTag  string `protobuf:"bytes,2,opt,name=discord_tag,json=discordTag,proto3" json:"discord_tag,omitempty"`
	IsStudent   bool   `protobuf:"varint,3,opt,name=is_student,json=isStudent,proto3" json:"is_student,omitempty"`
	AvatarUrl   string `protobuf:"bytes,4,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	GithubId    string `protobuf:"bytes,16,opt,name=github_id,json=githubId,proto3" json:"github_id,omitempty"`
	DiscordId   string `protobuf:"bytes,17,opt,name=discord_id,json=discordId,proto3" json:"discord_id,omitempty"`
}

func (x *Contestant_ContestantDetail) Reset() {
	*x = Contestant_ContestantDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isuxportal_resources_contestant_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Contestant_ContestantDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contestant_ContestantDetail) ProtoMessage() {}

func (x *Contestant_ContestantDetail) ProtoReflect() protoreflect.Message {
	mi := &file_isuxportal_resources_contestant_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Contestant_ContestantDetail.ProtoReflect.Descriptor instead.
func (*Contestant_ContestantDetail) Descriptor() ([]byte, []int) {
	return file_isuxportal_resources_contestant_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Contestant_ContestantDetail) GetGithubLogin() string {
	if x != nil {
		return x.GithubLogin
	}
	return ""
}

func (x *Contestant_ContestantDetail) GetDiscordTag() string {
	if x != nil {
		return x.DiscordTag
	}
	return ""
}

func (x *Contestant_ContestantDetail) GetIsStudent() bool {
	if x != nil {
		return x.IsStudent
	}
	return false
}

func (x *Contestant_ContestantDetail) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

func (x *Contestant_ContestantDetail) GetGithubId() string {
	if x != nil {
		return x.GithubId
	}
	return ""
}

func (x *Contestant_ContestantDetail) GetDiscordId() string {
	if x != nil {
		return x.DiscordId
	}
	return ""
}

var File_isuxportal_resources_contestant_proto protoreflect.FileDescriptor

var file_isuxportal_resources_contestant_proto_rawDesc = []byte{
	0x0a, 0x25, 0x69, 0x73, 0x75, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x61, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x69, 0x73, 0x75, 0x78, 0x70, 0x6f, 0x72,
	0x74, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x22, 0x82, 0x03, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x61,
	0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x64, 0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x5f, 0x64, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x69, 0x73, 0x75,
	0x78, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x61,
	0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x1a, 0xd0, 0x01, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x73,
	0x74, 0x61, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1f, 0x0a,
	0x0b, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x74, 0x61, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x61, 0x67, 0x12, 0x1d,
	0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x64, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x73, 0x75, 0x63, 0x6f, 0x6e, 0x2f, 0x69, 0x73,
	0x75, 0x63, 0x6f, 0x6e, 0x31, 0x30, 0x2d, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_isuxportal_resources_contestant_proto_rawDescOnce sync.Once
	file_isuxportal_resources_contestant_proto_rawDescData = file_isuxportal_resources_contestant_proto_rawDesc
)

func file_isuxportal_resources_contestant_proto_rawDescGZIP() []byte {
	file_isuxportal_resources_contestant_proto_rawDescOnce.Do(func() {
		file_isuxportal_resources_contestant_proto_rawDescData = protoimpl.X.CompressGZIP(file_isuxportal_resources_contestant_proto_rawDescData)
	})
	return file_isuxportal_resources_contestant_proto_rawDescData
}

var file_isuxportal_resources_contestant_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_isuxportal_resources_contestant_proto_goTypes = []interface{}{
	(*Contestant)(nil),                  // 0: isuxportal.proto.resources.Contestant
	(*Contestant_ContestantDetail)(nil), // 1: isuxportal.proto.resources.Contestant.ContestantDetail
}
var file_isuxportal_resources_contestant_proto_depIdxs = []int32{
	1, // 0: isuxportal.proto.resources.Contestant.contestant_detail:type_name -> isuxportal.proto.resources.Contestant.ContestantDetail
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_isuxportal_resources_contestant_proto_init() }
func file_isuxportal_resources_contestant_proto_init() {
	if File_isuxportal_resources_contestant_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_isuxportal_resources_contestant_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Contestant); i {
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
		file_isuxportal_resources_contestant_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Contestant_ContestantDetail); i {
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
			RawDescriptor: file_isuxportal_resources_contestant_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_isuxportal_resources_contestant_proto_goTypes,
		DependencyIndexes: file_isuxportal_resources_contestant_proto_depIdxs,
		MessageInfos:      file_isuxportal_resources_contestant_proto_msgTypes,
	}.Build()
	File_isuxportal_resources_contestant_proto = out.File
	file_isuxportal_resources_contestant_proto_rawDesc = nil
	file_isuxportal_resources_contestant_proto_goTypes = nil
	file_isuxportal_resources_contestant_proto_depIdxs = nil
}
