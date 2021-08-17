// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.8
// source: isuxportal/services/admin/clarifications.proto

package admin

import (
	resources "github.com/isucon/isucon11-portal/proto.go/isuxportal/resources"
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

type ListClarificationsQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// optional to filter
	TeamId         int64 `protobuf:"varint,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	UnansweredOnly bool  `protobuf:"varint,2,opt,name=unanswered_only,json=unansweredOnly,proto3" json:"unanswered_only,omitempty"`
}

func (x *ListClarificationsQuery) Reset() {
	*x = ListClarificationsQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isuxportal_services_admin_clarifications_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListClarificationsQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListClarificationsQuery) ProtoMessage() {}

func (x *ListClarificationsQuery) ProtoReflect() protoreflect.Message {
	mi := &file_isuxportal_services_admin_clarifications_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListClarificationsQuery.ProtoReflect.Descriptor instead.
func (*ListClarificationsQuery) Descriptor() ([]byte, []int) {
	return file_isuxportal_services_admin_clarifications_proto_rawDescGZIP(), []int{0}
}

func (x *ListClarificationsQuery) GetTeamId() int64 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

func (x *ListClarificationsQuery) GetUnansweredOnly() bool {
	if x != nil {
		return x.UnansweredOnly
	}
	return false
}

type ListClarificationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clarifications []*resources.Clarification `protobuf:"bytes,1,rep,name=clarifications,proto3" json:"clarifications,omitempty"`
}

func (x *ListClarificationsResponse) Reset() {
	*x = ListClarificationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isuxportal_services_admin_clarifications_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListClarificationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListClarificationsResponse) ProtoMessage() {}

func (x *ListClarificationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_isuxportal_services_admin_clarifications_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListClarificationsResponse.ProtoReflect.Descriptor instead.
func (*ListClarificationsResponse) Descriptor() ([]byte, []int) {
	return file_isuxportal_services_admin_clarifications_proto_rawDescGZIP(), []int{1}
}

func (x *ListClarificationsResponse) GetClarifications() []*resources.Clarification {
	if x != nil {
		return x.Clarifications
	}
	return nil
}

type GetClarificationQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetClarificationQuery) Reset() {
	*x = GetClarificationQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isuxportal_services_admin_clarifications_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClarificationQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClarificationQuery) ProtoMessage() {}

func (x *GetClarificationQuery) ProtoReflect() protoreflect.Message {
	mi := &file_isuxportal_services_admin_clarifications_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClarificationQuery.ProtoReflect.Descriptor instead.
func (*GetClarificationQuery) Descriptor() ([]byte, []int) {
	return file_isuxportal_services_admin_clarifications_proto_rawDescGZIP(), []int{2}
}

func (x *GetClarificationQuery) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetClarificationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clarification *resources.Clarification `protobuf:"bytes,1,opt,name=clarification,proto3" json:"clarification,omitempty"`
}

func (x *GetClarificationResponse) Reset() {
	*x = GetClarificationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isuxportal_services_admin_clarifications_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClarificationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClarificationResponse) ProtoMessage() {}

func (x *GetClarificationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_isuxportal_services_admin_clarifications_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClarificationResponse.ProtoReflect.Descriptor instead.
func (*GetClarificationResponse) Descriptor() ([]byte, []int) {
	return file_isuxportal_services_admin_clarifications_proto_rawDescGZIP(), []int{3}
}

func (x *GetClarificationResponse) GetClarification() *resources.Clarification {
	if x != nil {
		return x.Clarification
	}
	return nil
}

type RespondClarificationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Disclose bool   `protobuf:"varint,2,opt,name=disclose,proto3" json:"disclose,omitempty"`
	Answer   string `protobuf:"bytes,3,opt,name=answer,proto3" json:"answer,omitempty"`
	// optional to override original question
	Question string `protobuf:"bytes,4,opt,name=question,proto3" json:"question,omitempty"`
}

func (x *RespondClarificationRequest) Reset() {
	*x = RespondClarificationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isuxportal_services_admin_clarifications_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespondClarificationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespondClarificationRequest) ProtoMessage() {}

func (x *RespondClarificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_isuxportal_services_admin_clarifications_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespondClarificationRequest.ProtoReflect.Descriptor instead.
func (*RespondClarificationRequest) Descriptor() ([]byte, []int) {
	return file_isuxportal_services_admin_clarifications_proto_rawDescGZIP(), []int{4}
}

func (x *RespondClarificationRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RespondClarificationRequest) GetDisclose() bool {
	if x != nil {
		return x.Disclose
	}
	return false
}

func (x *RespondClarificationRequest) GetAnswer() string {
	if x != nil {
		return x.Answer
	}
	return ""
}

func (x *RespondClarificationRequest) GetQuestion() string {
	if x != nil {
		return x.Question
	}
	return ""
}

type RespondClarificationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clarification *resources.Clarification `protobuf:"bytes,1,opt,name=clarification,proto3" json:"clarification,omitempty"`
}

func (x *RespondClarificationResponse) Reset() {
	*x = RespondClarificationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isuxportal_services_admin_clarifications_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespondClarificationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespondClarificationResponse) ProtoMessage() {}

func (x *RespondClarificationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_isuxportal_services_admin_clarifications_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespondClarificationResponse.ProtoReflect.Descriptor instead.
func (*RespondClarificationResponse) Descriptor() ([]byte, []int) {
	return file_isuxportal_services_admin_clarifications_proto_rawDescGZIP(), []int{5}
}

func (x *RespondClarificationResponse) GetClarification() *resources.Clarification {
	if x != nil {
		return x.Clarification
	}
	return nil
}

type CreateClarificationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Answer   string `protobuf:"bytes,2,opt,name=answer,proto3" json:"answer,omitempty"`
	Question string `protobuf:"bytes,3,opt,name=question,proto3" json:"question,omitempty"`
	TeamId   int64  `protobuf:"varint,4,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
}

func (x *CreateClarificationRequest) Reset() {
	*x = CreateClarificationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isuxportal_services_admin_clarifications_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateClarificationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateClarificationRequest) ProtoMessage() {}

func (x *CreateClarificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_isuxportal_services_admin_clarifications_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateClarificationRequest.ProtoReflect.Descriptor instead.
func (*CreateClarificationRequest) Descriptor() ([]byte, []int) {
	return file_isuxportal_services_admin_clarifications_proto_rawDescGZIP(), []int{6}
}

func (x *CreateClarificationRequest) GetAnswer() string {
	if x != nil {
		return x.Answer
	}
	return ""
}

func (x *CreateClarificationRequest) GetQuestion() string {
	if x != nil {
		return x.Question
	}
	return ""
}

func (x *CreateClarificationRequest) GetTeamId() int64 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

type CreateClarificationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clarification *resources.Clarification `protobuf:"bytes,1,opt,name=clarification,proto3" json:"clarification,omitempty"`
}

func (x *CreateClarificationResponse) Reset() {
	*x = CreateClarificationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isuxportal_services_admin_clarifications_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateClarificationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateClarificationResponse) ProtoMessage() {}

func (x *CreateClarificationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_isuxportal_services_admin_clarifications_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateClarificationResponse.ProtoReflect.Descriptor instead.
func (*CreateClarificationResponse) Descriptor() ([]byte, []int) {
	return file_isuxportal_services_admin_clarifications_proto_rawDescGZIP(), []int{7}
}

func (x *CreateClarificationResponse) GetClarification() *resources.Clarification {
	if x != nil {
		return x.Clarification
	}
	return nil
}

var File_isuxportal_services_admin_clarifications_proto protoreflect.FileDescriptor

var file_isuxportal_services_admin_clarifications_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x69, 0x73, 0x75, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x63, 0x6c, 0x61, 0x72,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1f, 0x69, 0x73, 0x75, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x1a, 0x28, 0x69, 0x73, 0x75, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x6c, 0x61, 0x72, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5b, 0x0a, 0x17, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x6c, 0x61, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12,
	0x27, 0x0a, 0x0f, 0x75, 0x6e, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x65, 0x64, 0x5f, 0x6f, 0x6e,
	0x6c, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x75, 0x6e, 0x61, 0x6e, 0x73, 0x77,
	0x65, 0x72, 0x65, 0x64, 0x4f, 0x6e, 0x6c, 0x79, 0x22, 0x6f, 0x0a, 0x1a, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x6c, 0x61, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x0e, 0x63, 0x6c, 0x61, 0x72, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29,
	0x2e, 0x69, 0x73, 0x75, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6c, 0x61, 0x72,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x63, 0x6c, 0x61, 0x72, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x27, 0x0a, 0x15, 0x47, 0x65, 0x74,
	0x43, 0x6c, 0x61, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x6b, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x61, 0x72, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f,
	0x0a, 0x0d, 0x63, 0x6c, 0x61, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x69, 0x73, 0x75, 0x78, 0x70, 0x6f, 0x72, 0x74,
	0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2e, 0x43, 0x6c, 0x61, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0d, 0x63, 0x6c, 0x61, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x7d, 0x0a, 0x1b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x43, 0x6c, 0x61, 0x72, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6e,
	0x73, 0x77, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6e, 0x73, 0x77,
	0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x6f,
	0x0a, 0x1c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x43, 0x6c, 0x61, 0x72, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f,
	0x0a, 0x0d, 0x63, 0x6c, 0x61, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x69, 0x73, 0x75, 0x78, 0x70, 0x6f, 0x72, 0x74,
	0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2e, 0x43, 0x6c, 0x61, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0d, 0x63, 0x6c, 0x61, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x69, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x61, 0x72, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x22, 0x6e, 0x0a, 0x1b, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x61, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0d, 0x63, 0x6c, 0x61,
	0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x29, 0x2e, 0x69, 0x73, 0x75, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6c,
	0x61, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x63, 0x6c, 0x61,
	0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x46, 0x5a, 0x44, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x73, 0x75, 0x63, 0x6f, 0x6e, 0x2f,
	0x69, 0x73, 0x75, 0x63, 0x6f, 0x6e, 0x31, 0x30, 0x2d, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x6f, 0x2f, 0x69, 0x73, 0x75, 0x78, 0x70, 0x6f, 0x72,
	0x74, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_isuxportal_services_admin_clarifications_proto_rawDescOnce sync.Once
	file_isuxportal_services_admin_clarifications_proto_rawDescData = file_isuxportal_services_admin_clarifications_proto_rawDesc
)

func file_isuxportal_services_admin_clarifications_proto_rawDescGZIP() []byte {
	file_isuxportal_services_admin_clarifications_proto_rawDescOnce.Do(func() {
		file_isuxportal_services_admin_clarifications_proto_rawDescData = protoimpl.X.CompressGZIP(file_isuxportal_services_admin_clarifications_proto_rawDescData)
	})
	return file_isuxportal_services_admin_clarifications_proto_rawDescData
}

var file_isuxportal_services_admin_clarifications_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_isuxportal_services_admin_clarifications_proto_goTypes = []interface{}{
	(*ListClarificationsQuery)(nil),      // 0: isuxportal.proto.services.admin.ListClarificationsQuery
	(*ListClarificationsResponse)(nil),   // 1: isuxportal.proto.services.admin.ListClarificationsResponse
	(*GetClarificationQuery)(nil),        // 2: isuxportal.proto.services.admin.GetClarificationQuery
	(*GetClarificationResponse)(nil),     // 3: isuxportal.proto.services.admin.GetClarificationResponse
	(*RespondClarificationRequest)(nil),  // 4: isuxportal.proto.services.admin.RespondClarificationRequest
	(*RespondClarificationResponse)(nil), // 5: isuxportal.proto.services.admin.RespondClarificationResponse
	(*CreateClarificationRequest)(nil),   // 6: isuxportal.proto.services.admin.CreateClarificationRequest
	(*CreateClarificationResponse)(nil),  // 7: isuxportal.proto.services.admin.CreateClarificationResponse
	(*resources.Clarification)(nil),      // 8: isuxportal.proto.resources.Clarification
}
var file_isuxportal_services_admin_clarifications_proto_depIdxs = []int32{
	8, // 0: isuxportal.proto.services.admin.ListClarificationsResponse.clarifications:type_name -> isuxportal.proto.resources.Clarification
	8, // 1: isuxportal.proto.services.admin.GetClarificationResponse.clarification:type_name -> isuxportal.proto.resources.Clarification
	8, // 2: isuxportal.proto.services.admin.RespondClarificationResponse.clarification:type_name -> isuxportal.proto.resources.Clarification
	8, // 3: isuxportal.proto.services.admin.CreateClarificationResponse.clarification:type_name -> isuxportal.proto.resources.Clarification
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_isuxportal_services_admin_clarifications_proto_init() }
func file_isuxportal_services_admin_clarifications_proto_init() {
	if File_isuxportal_services_admin_clarifications_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_isuxportal_services_admin_clarifications_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListClarificationsQuery); i {
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
		file_isuxportal_services_admin_clarifications_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListClarificationsResponse); i {
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
		file_isuxportal_services_admin_clarifications_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClarificationQuery); i {
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
		file_isuxportal_services_admin_clarifications_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClarificationResponse); i {
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
		file_isuxportal_services_admin_clarifications_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespondClarificationRequest); i {
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
		file_isuxportal_services_admin_clarifications_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespondClarificationResponse); i {
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
		file_isuxportal_services_admin_clarifications_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateClarificationRequest); i {
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
		file_isuxportal_services_admin_clarifications_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateClarificationResponse); i {
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
			RawDescriptor: file_isuxportal_services_admin_clarifications_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_isuxportal_services_admin_clarifications_proto_goTypes,
		DependencyIndexes: file_isuxportal_services_admin_clarifications_proto_depIdxs,
		MessageInfos:      file_isuxportal_services_admin_clarifications_proto_msgTypes,
	}.Build()
	File_isuxportal_services_admin_clarifications_proto = out.File
	file_isuxportal_services_admin_clarifications_proto_rawDesc = nil
	file_isuxportal_services_admin_clarifications_proto_goTypes = nil
	file_isuxportal_services_admin_clarifications_proto_depIdxs = nil
}
