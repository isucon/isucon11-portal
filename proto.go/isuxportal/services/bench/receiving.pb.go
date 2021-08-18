// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.8
// source: isuxportal/services/bench/receiving.proto

package bench

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

type ReceiveBenchmarkJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token        string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	InstanceName string `protobuf:"bytes,2,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	TeamId       int64  `protobuf:"varint,3,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
}

func (x *ReceiveBenchmarkJobRequest) Reset() {
	*x = ReceiveBenchmarkJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isuxportal_services_bench_receiving_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiveBenchmarkJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveBenchmarkJobRequest) ProtoMessage() {}

func (x *ReceiveBenchmarkJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_isuxportal_services_bench_receiving_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveBenchmarkJobRequest.ProtoReflect.Descriptor instead.
func (*ReceiveBenchmarkJobRequest) Descriptor() ([]byte, []int) {
	return file_isuxportal_services_bench_receiving_proto_rawDescGZIP(), []int{0}
}

func (x *ReceiveBenchmarkJobRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *ReceiveBenchmarkJobRequest) GetInstanceName() string {
	if x != nil {
		return x.InstanceName
	}
	return ""
}

func (x *ReceiveBenchmarkJobRequest) GetTeamId() int64 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

type ReceiveBenchmarkJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// optional
	JobHandle *ReceiveBenchmarkJobResponse_JobHandle `protobuf:"bytes,1,opt,name=job_handle,json=jobHandle,proto3" json:"job_handle,omitempty"`
}

func (x *ReceiveBenchmarkJobResponse) Reset() {
	*x = ReceiveBenchmarkJobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isuxportal_services_bench_receiving_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiveBenchmarkJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveBenchmarkJobResponse) ProtoMessage() {}

func (x *ReceiveBenchmarkJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_isuxportal_services_bench_receiving_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveBenchmarkJobResponse.ProtoReflect.Descriptor instead.
func (*ReceiveBenchmarkJobResponse) Descriptor() ([]byte, []int) {
	return file_isuxportal_services_bench_receiving_proto_rawDescGZIP(), []int{1}
}

func (x *ReceiveBenchmarkJobResponse) GetJobHandle() *ReceiveBenchmarkJobResponse_JobHandle {
	if x != nil {
		return x.JobHandle
	}
	return nil
}

type CancelOwnedBenchmarkJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token        string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	InstanceName string `protobuf:"bytes,2,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
}

func (x *CancelOwnedBenchmarkJobRequest) Reset() {
	*x = CancelOwnedBenchmarkJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isuxportal_services_bench_receiving_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelOwnedBenchmarkJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelOwnedBenchmarkJobRequest) ProtoMessage() {}

func (x *CancelOwnedBenchmarkJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_isuxportal_services_bench_receiving_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelOwnedBenchmarkJobRequest.ProtoReflect.Descriptor instead.
func (*CancelOwnedBenchmarkJobRequest) Descriptor() ([]byte, []int) {
	return file_isuxportal_services_bench_receiving_proto_rawDescGZIP(), []int{2}
}

func (x *CancelOwnedBenchmarkJobRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *CancelOwnedBenchmarkJobRequest) GetInstanceName() string {
	if x != nil {
		return x.InstanceName
	}
	return ""
}

type CancelOwnedBenchmarkJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CancelOwnedBenchmarkJobResponse) Reset() {
	*x = CancelOwnedBenchmarkJobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isuxportal_services_bench_receiving_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelOwnedBenchmarkJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelOwnedBenchmarkJobResponse) ProtoMessage() {}

func (x *CancelOwnedBenchmarkJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_isuxportal_services_bench_receiving_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelOwnedBenchmarkJobResponse.ProtoReflect.Descriptor instead.
func (*CancelOwnedBenchmarkJobResponse) Descriptor() ([]byte, []int) {
	return file_isuxportal_services_bench_receiving_proto_rawDescGZIP(), []int{3}
}

type ReceiveBenchmarkJobResponse_JobHandle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobId             int64    `protobuf:"varint,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	Handle            string   `protobuf:"bytes,2,opt,name=handle,proto3" json:"handle,omitempty"`
	TargetIpv4Address string   `protobuf:"bytes,3,opt,name=target_ipv4_address,json=targetIpv4Address,proto3" json:"target_ipv4_address,omitempty"`
	DescriptionHuman  string   `protobuf:"bytes,4,opt,name=description_human,json=descriptionHuman,proto3" json:"description_human,omitempty"`
	AllIpv4Addresses  []string `protobuf:"bytes,5,rep,name=all_ipv4_addresses,json=allIpv4Addresses,proto3" json:"all_ipv4_addresses,omitempty"`
}

func (x *ReceiveBenchmarkJobResponse_JobHandle) Reset() {
	*x = ReceiveBenchmarkJobResponse_JobHandle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isuxportal_services_bench_receiving_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiveBenchmarkJobResponse_JobHandle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveBenchmarkJobResponse_JobHandle) ProtoMessage() {}

func (x *ReceiveBenchmarkJobResponse_JobHandle) ProtoReflect() protoreflect.Message {
	mi := &file_isuxportal_services_bench_receiving_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveBenchmarkJobResponse_JobHandle.ProtoReflect.Descriptor instead.
func (*ReceiveBenchmarkJobResponse_JobHandle) Descriptor() ([]byte, []int) {
	return file_isuxportal_services_bench_receiving_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ReceiveBenchmarkJobResponse_JobHandle) GetJobId() int64 {
	if x != nil {
		return x.JobId
	}
	return 0
}

func (x *ReceiveBenchmarkJobResponse_JobHandle) GetHandle() string {
	if x != nil {
		return x.Handle
	}
	return ""
}

func (x *ReceiveBenchmarkJobResponse_JobHandle) GetTargetIpv4Address() string {
	if x != nil {
		return x.TargetIpv4Address
	}
	return ""
}

func (x *ReceiveBenchmarkJobResponse_JobHandle) GetDescriptionHuman() string {
	if x != nil {
		return x.DescriptionHuman
	}
	return ""
}

func (x *ReceiveBenchmarkJobResponse_JobHandle) GetAllIpv4Addresses() []string {
	if x != nil {
		return x.AllIpv4Addresses
	}
	return nil
}

var File_isuxportal_services_bench_receiving_proto protoreflect.FileDescriptor

var file_isuxportal_services_bench_receiving_proto_rawDesc = []byte{
	0x0a, 0x29, 0x69, 0x73, 0x75, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x2f, 0x72, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x69, 0x73, 0x75,
	0x78, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x22, 0x70, 0x0a, 0x1a,
	0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x42, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b,
	0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x22, 0xcc,
	0x02, 0x0a, 0x1b, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x42, 0x65, 0x6e, 0x63, 0x68, 0x6d,
	0x61, 0x72, 0x6b, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x65,
	0x0a, 0x0a, 0x6a, 0x6f, 0x62, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x46, 0x2e, 0x69, 0x73, 0x75, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62,
	0x65, 0x6e, 0x63, 0x68, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x42, 0x65, 0x6e, 0x63,
	0x68, 0x6d, 0x61, 0x72, 0x6b, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x4a, 0x6f, 0x62, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x09, 0x6a, 0x6f, 0x62, 0x48,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x1a, 0xc5, 0x01, 0x0a, 0x09, 0x4a, 0x6f, 0x62, 0x48, 0x61, 0x6e,
	0x64, 0x6c, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x70, 0x76,
	0x34, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x11, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x70, 0x76, 0x34, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x68, 0x75, 0x6d, 0x61, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x75, 0x6d, 0x61, 0x6e, 0x12,
	0x2c, 0x0a, 0x12, 0x61, 0x6c, 0x6c, 0x5f, 0x69, 0x70, 0x76, 0x34, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x61, 0x6c, 0x6c,
	0x49, 0x70, 0x76, 0x34, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x22, 0x5b, 0x0a,
	0x1e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4f, 0x77, 0x6e, 0x65, 0x64, 0x42, 0x65, 0x6e, 0x63,
	0x68, 0x6d, 0x61, 0x72, 0x6b, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x21, 0x0a, 0x1f, 0x43, 0x61,
	0x6e, 0x63, 0x65, 0x6c, 0x4f, 0x77, 0x6e, 0x65, 0x64, 0x42, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61,
	0x72, 0x6b, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xc2, 0x02,
	0x0a, 0x0e, 0x42, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x51, 0x75, 0x65, 0x75, 0x65,
	0x12, 0x90, 0x01, 0x0a, 0x13, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x42, 0x65, 0x6e, 0x63,
	0x68, 0x6d, 0x61, 0x72, 0x6b, 0x4a, 0x6f, 0x62, 0x12, 0x3b, 0x2e, 0x69, 0x73, 0x75, 0x78, 0x70,
	0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x42, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x4a, 0x6f, 0x62, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x69, 0x73, 0x75, 0x78, 0x70, 0x6f, 0x72, 0x74,
	0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x42,
	0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x9c, 0x01, 0x0a, 0x17, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4f, 0x77,
	0x6e, 0x65, 0x64, 0x42, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x4a, 0x6f, 0x62, 0x12,
	0x3f, 0x2e, 0x69, 0x73, 0x75, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x65, 0x6e, 0x63,
	0x68, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4f, 0x77, 0x6e, 0x65, 0x64, 0x42, 0x65, 0x6e,
	0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x40, 0x2e, 0x69, 0x73, 0x75, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x65, 0x6e,
	0x63, 0x68, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4f, 0x77, 0x6e, 0x65, 0x64, 0x42, 0x65,
	0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x69, 0x73, 0x75, 0x63, 0x6f, 0x6e, 0x2f, 0x69, 0x73, 0x75, 0x63, 0x6f, 0x6e, 0x31, 0x30,
	0x2d, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x6f,
	0x2f, 0x69, 0x73, 0x75, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2f, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_isuxportal_services_bench_receiving_proto_rawDescOnce sync.Once
	file_isuxportal_services_bench_receiving_proto_rawDescData = file_isuxportal_services_bench_receiving_proto_rawDesc
)

func file_isuxportal_services_bench_receiving_proto_rawDescGZIP() []byte {
	file_isuxportal_services_bench_receiving_proto_rawDescOnce.Do(func() {
		file_isuxportal_services_bench_receiving_proto_rawDescData = protoimpl.X.CompressGZIP(file_isuxportal_services_bench_receiving_proto_rawDescData)
	})
	return file_isuxportal_services_bench_receiving_proto_rawDescData
}

var file_isuxportal_services_bench_receiving_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_isuxportal_services_bench_receiving_proto_goTypes = []interface{}{
	(*ReceiveBenchmarkJobRequest)(nil),            // 0: isuxportal.proto.services.bench.ReceiveBenchmarkJobRequest
	(*ReceiveBenchmarkJobResponse)(nil),           // 1: isuxportal.proto.services.bench.ReceiveBenchmarkJobResponse
	(*CancelOwnedBenchmarkJobRequest)(nil),        // 2: isuxportal.proto.services.bench.CancelOwnedBenchmarkJobRequest
	(*CancelOwnedBenchmarkJobResponse)(nil),       // 3: isuxportal.proto.services.bench.CancelOwnedBenchmarkJobResponse
	(*ReceiveBenchmarkJobResponse_JobHandle)(nil), // 4: isuxportal.proto.services.bench.ReceiveBenchmarkJobResponse.JobHandle
}
var file_isuxportal_services_bench_receiving_proto_depIdxs = []int32{
	4, // 0: isuxportal.proto.services.bench.ReceiveBenchmarkJobResponse.job_handle:type_name -> isuxportal.proto.services.bench.ReceiveBenchmarkJobResponse.JobHandle
	0, // 1: isuxportal.proto.services.bench.BenchmarkQueue.ReceiveBenchmarkJob:input_type -> isuxportal.proto.services.bench.ReceiveBenchmarkJobRequest
	2, // 2: isuxportal.proto.services.bench.BenchmarkQueue.CancelOwnedBenchmarkJob:input_type -> isuxportal.proto.services.bench.CancelOwnedBenchmarkJobRequest
	1, // 3: isuxportal.proto.services.bench.BenchmarkQueue.ReceiveBenchmarkJob:output_type -> isuxportal.proto.services.bench.ReceiveBenchmarkJobResponse
	3, // 4: isuxportal.proto.services.bench.BenchmarkQueue.CancelOwnedBenchmarkJob:output_type -> isuxportal.proto.services.bench.CancelOwnedBenchmarkJobResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_isuxportal_services_bench_receiving_proto_init() }
func file_isuxportal_services_bench_receiving_proto_init() {
	if File_isuxportal_services_bench_receiving_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_isuxportal_services_bench_receiving_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiveBenchmarkJobRequest); i {
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
		file_isuxportal_services_bench_receiving_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiveBenchmarkJobResponse); i {
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
		file_isuxportal_services_bench_receiving_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelOwnedBenchmarkJobRequest); i {
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
		file_isuxportal_services_bench_receiving_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelOwnedBenchmarkJobResponse); i {
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
		file_isuxportal_services_bench_receiving_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiveBenchmarkJobResponse_JobHandle); i {
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
			RawDescriptor: file_isuxportal_services_bench_receiving_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_isuxportal_services_bench_receiving_proto_goTypes,
		DependencyIndexes: file_isuxportal_services_bench_receiving_proto_depIdxs,
		MessageInfos:      file_isuxportal_services_bench_receiving_proto_msgTypes,
	}.Build()
	File_isuxportal_services_bench_receiving_proto = out.File
	file_isuxportal_services_bench_receiving_proto_rawDesc = nil
	file_isuxportal_services_bench_receiving_proto_goTypes = nil
	file_isuxportal_services_bench_receiving_proto_depIdxs = nil
}
