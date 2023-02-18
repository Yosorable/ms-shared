// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.10
// source: protos/faas/req_res.proto

package faas

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

type CreateFaasTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name            string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	HttpTriggerPath *string `protobuf:"bytes,2,opt,name=http_trigger_path,json=httpTriggerPath,proto3,oneof" json:"http_trigger_path,omitempty"`
	CronTriggerExp  *string `protobuf:"bytes,3,opt,name=cron_trigger_exp,json=cronTriggerExp,proto3,oneof" json:"cron_trigger_exp,omitempty"`
	FunctionCode    string  `protobuf:"bytes,4,opt,name=function_code,json=functionCode,proto3" json:"function_code,omitempty"`
}

func (x *CreateFaasTaskRequest) Reset() {
	*x = CreateFaasTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_faas_req_res_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFaasTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFaasTaskRequest) ProtoMessage() {}

func (x *CreateFaasTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_faas_req_res_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFaasTaskRequest.ProtoReflect.Descriptor instead.
func (*CreateFaasTaskRequest) Descriptor() ([]byte, []int) {
	return file_protos_faas_req_res_proto_rawDescGZIP(), []int{0}
}

func (x *CreateFaasTaskRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateFaasTaskRequest) GetHttpTriggerPath() string {
	if x != nil && x.HttpTriggerPath != nil {
		return *x.HttpTriggerPath
	}
	return ""
}

func (x *CreateFaasTaskRequest) GetCronTriggerExp() string {
	if x != nil && x.CronTriggerExp != nil {
		return *x.CronTriggerExp
	}
	return ""
}

func (x *CreateFaasTaskRequest) GetFunctionCode() string {
	if x != nil {
		return x.FunctionCode
	}
	return ""
}

type CreateFaasTaskReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateFaasTaskReply) Reset() {
	*x = CreateFaasTaskReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_faas_req_res_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFaasTaskReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFaasTaskReply) ProtoMessage() {}

func (x *CreateFaasTaskReply) ProtoReflect() protoreflect.Message {
	mi := &file_protos_faas_req_res_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFaasTaskReply.ProtoReflect.Descriptor instead.
func (*CreateFaasTaskReply) Descriptor() ([]byte, []int) {
	return file_protos_faas_req_res_proto_rawDescGZIP(), []int{1}
}

type RunTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Params *string `protobuf:"bytes,1,opt,name=params,proto3,oneof" json:"params,omitempty"`
}

func (x *RunTaskRequest) Reset() {
	*x = RunTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_faas_req_res_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunTaskRequest) ProtoMessage() {}

func (x *RunTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_faas_req_res_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunTaskRequest.ProtoReflect.Descriptor instead.
func (*RunTaskRequest) Descriptor() ([]byte, []int) {
	return file_protos_faas_req_res_proto_rawDescGZIP(), []int{2}
}

func (x *RunTaskRequest) GetParams() string {
	if x != nil && x.Params != nil {
		return *x.Params
	}
	return ""
}

type RunTaskReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *string `protobuf:"bytes,1,opt,name=result,proto3,oneof" json:"result,omitempty"`
}

func (x *RunTaskReply) Reset() {
	*x = RunTaskReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_faas_req_res_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunTaskReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunTaskReply) ProtoMessage() {}

func (x *RunTaskReply) ProtoReflect() protoreflect.Message {
	mi := &file_protos_faas_req_res_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunTaskReply.ProtoReflect.Descriptor instead.
func (*RunTaskReply) Descriptor() ([]byte, []int) {
	return file_protos_faas_req_res_proto_rawDescGZIP(), []int{3}
}

func (x *RunTaskReply) GetResult() string {
	if x != nil && x.Result != nil {
		return *x.Result
	}
	return ""
}

type GetFaasTaskByPageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int32   `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int32   `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Name     *string `protobuf:"bytes,3,opt,name=name,proto3,oneof" json:"name,omitempty"`
}

func (x *GetFaasTaskByPageRequest) Reset() {
	*x = GetFaasTaskByPageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_faas_req_res_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFaasTaskByPageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFaasTaskByPageRequest) ProtoMessage() {}

func (x *GetFaasTaskByPageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_faas_req_res_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFaasTaskByPageRequest.ProtoReflect.Descriptor instead.
func (*GetFaasTaskByPageRequest) Descriptor() ([]byte, []int) {
	return file_protos_faas_req_res_proto_rawDescGZIP(), []int{4}
}

func (x *GetFaasTaskByPageRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetFaasTaskByPageRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *GetFaasTaskByPageRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

type GetFaasTaskByPageReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FaasTasks []*FaasTask `protobuf:"bytes,1,rep,name=faas_tasks,json=faasTasks,proto3" json:"faas_tasks,omitempty"`
	Total     int32       `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *GetFaasTaskByPageReply) Reset() {
	*x = GetFaasTaskByPageReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_faas_req_res_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFaasTaskByPageReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFaasTaskByPageReply) ProtoMessage() {}

func (x *GetFaasTaskByPageReply) ProtoReflect() protoreflect.Message {
	mi := &file_protos_faas_req_res_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFaasTaskByPageReply.ProtoReflect.Descriptor instead.
func (*GetFaasTaskByPageReply) Descriptor() ([]byte, []int) {
	return file_protos_faas_req_res_proto_rawDescGZIP(), []int{5}
}

func (x *GetFaasTaskByPageReply) GetFaasTasks() []*FaasTask {
	if x != nil {
		return x.FaasTasks
	}
	return nil
}

func (x *GetFaasTaskByPageReply) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetFaasTaskByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetFaasTaskByIDRequest) Reset() {
	*x = GetFaasTaskByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_faas_req_res_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFaasTaskByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFaasTaskByIDRequest) ProtoMessage() {}

func (x *GetFaasTaskByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_faas_req_res_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFaasTaskByIDRequest.ProtoReflect.Descriptor instead.
func (*GetFaasTaskByIDRequest) Descriptor() ([]byte, []int) {
	return file_protos_faas_req_res_proto_rawDescGZIP(), []int{6}
}

func (x *GetFaasTaskByIDRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetFaasTaskByIDReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FaasTask *FaasTask `protobuf:"bytes,1,opt,name=faas_task,json=faasTask,proto3" json:"faas_task,omitempty"`
}

func (x *GetFaasTaskByIDReply) Reset() {
	*x = GetFaasTaskByIDReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_faas_req_res_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFaasTaskByIDReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFaasTaskByIDReply) ProtoMessage() {}

func (x *GetFaasTaskByIDReply) ProtoReflect() protoreflect.Message {
	mi := &file_protos_faas_req_res_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFaasTaskByIDReply.ProtoReflect.Descriptor instead.
func (*GetFaasTaskByIDReply) Descriptor() ([]byte, []int) {
	return file_protos_faas_req_res_proto_rawDescGZIP(), []int{7}
}

func (x *GetFaasTaskByIDReply) GetFaasTask() *FaasTask {
	if x != nil {
		return x.FaasTask
	}
	return nil
}

var File_protos_faas_req_res_proto protoreflect.FileDescriptor

var file_protos_faas_req_res_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x66, 0x61, 0x61, 0x73, 0x2f, 0x72, 0x65,
	0x71, 0x5f, 0x72, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x66, 0x61, 0x61,
	0x73, 0x1a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x66, 0x61, 0x61, 0x73, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdb, 0x01, 0x0a, 0x15, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x61, 0x61, 0x73, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x11, 0x68, 0x74, 0x74, 0x70,
	0x5f, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0f, 0x68, 0x74, 0x74, 0x70, 0x54, 0x72, 0x69, 0x67, 0x67,
	0x65, 0x72, 0x50, 0x61, 0x74, 0x68, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x10, 0x63, 0x72, 0x6f,
	0x6e, 0x5f, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x65, 0x78, 0x70, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0e, 0x63, 0x72, 0x6f, 0x6e, 0x54, 0x72, 0x69, 0x67, 0x67,
	0x65, 0x72, 0x45, 0x78, 0x70, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x75, 0x6e, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x14, 0x0a,
	0x12, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x70,
	0x61, 0x74, 0x68, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x63, 0x72, 0x6f, 0x6e, 0x5f, 0x74, 0x72, 0x69,
	0x67, 0x67, 0x65, 0x72, 0x5f, 0x65, 0x78, 0x70, 0x22, 0x15, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x46, 0x61, 0x61, 0x73, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x38, 0x0a, 0x0e, 0x52, 0x75, 0x6e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1b, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x88, 0x01, 0x01, 0x42, 0x09,
	0x0a, 0x07, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x36, 0x0a, 0x0c, 0x52, 0x75, 0x6e,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1b, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x22, 0x6d, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x46, 0x61, 0x61, 0x73, 0x54, 0x61, 0x73, 0x6b,
	0x42, 0x79, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x17,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x5d, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x46, 0x61, 0x61, 0x73, 0x54, 0x61, 0x73, 0x6b, 0x42,
	0x79, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2d, 0x0a, 0x0a, 0x66, 0x61,
	0x61, 0x73, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x66, 0x61, 0x61, 0x73, 0x2e, 0x46, 0x61, 0x61, 0x73, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x09,
	0x66, 0x61, 0x61, 0x73, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22,
	0x28, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x46, 0x61, 0x61, 0x73, 0x54, 0x61, 0x73, 0x6b, 0x42, 0x79,
	0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x43, 0x0a, 0x14, 0x47, 0x65, 0x74,
	0x46, 0x61, 0x61, 0x73, 0x54, 0x61, 0x73, 0x6b, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x2b, 0x0a, 0x09, 0x66, 0x61, 0x61, 0x73, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x66, 0x61, 0x61, 0x73, 0x2e, 0x46, 0x61, 0x61, 0x73,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x08, 0x66, 0x61, 0x61, 0x73, 0x54, 0x61, 0x73, 0x6b, 0x42, 0x30,
	0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x59, 0x6f, 0x73,
	0x6f, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x2f, 0x6d, 0x73, 0x2d, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x66, 0x61, 0x61, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_faas_req_res_proto_rawDescOnce sync.Once
	file_protos_faas_req_res_proto_rawDescData = file_protos_faas_req_res_proto_rawDesc
)

func file_protos_faas_req_res_proto_rawDescGZIP() []byte {
	file_protos_faas_req_res_proto_rawDescOnce.Do(func() {
		file_protos_faas_req_res_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_faas_req_res_proto_rawDescData)
	})
	return file_protos_faas_req_res_proto_rawDescData
}

var file_protos_faas_req_res_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_protos_faas_req_res_proto_goTypes = []interface{}{
	(*CreateFaasTaskRequest)(nil),    // 0: faas.CreateFaasTaskRequest
	(*CreateFaasTaskReply)(nil),      // 1: faas.CreateFaasTaskReply
	(*RunTaskRequest)(nil),           // 2: faas.RunTaskRequest
	(*RunTaskReply)(nil),             // 3: faas.RunTaskReply
	(*GetFaasTaskByPageRequest)(nil), // 4: faas.GetFaasTaskByPageRequest
	(*GetFaasTaskByPageReply)(nil),   // 5: faas.GetFaasTaskByPageReply
	(*GetFaasTaskByIDRequest)(nil),   // 6: faas.GetFaasTaskByIDRequest
	(*GetFaasTaskByIDReply)(nil),     // 7: faas.GetFaasTaskByIDReply
	(*FaasTask)(nil),                 // 8: faas.FaasTask
}
var file_protos_faas_req_res_proto_depIdxs = []int32{
	8, // 0: faas.GetFaasTaskByPageReply.faas_tasks:type_name -> faas.FaasTask
	8, // 1: faas.GetFaasTaskByIDReply.faas_task:type_name -> faas.FaasTask
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protos_faas_req_res_proto_init() }
func file_protos_faas_req_res_proto_init() {
	if File_protos_faas_req_res_proto != nil {
		return
	}
	file_protos_faas_model_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protos_faas_req_res_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFaasTaskRequest); i {
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
		file_protos_faas_req_res_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFaasTaskReply); i {
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
		file_protos_faas_req_res_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunTaskRequest); i {
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
		file_protos_faas_req_res_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunTaskReply); i {
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
		file_protos_faas_req_res_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFaasTaskByPageRequest); i {
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
		file_protos_faas_req_res_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFaasTaskByPageReply); i {
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
		file_protos_faas_req_res_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFaasTaskByIDRequest); i {
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
		file_protos_faas_req_res_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFaasTaskByIDReply); i {
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
	file_protos_faas_req_res_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_protos_faas_req_res_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_protos_faas_req_res_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_protos_faas_req_res_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_faas_req_res_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_faas_req_res_proto_goTypes,
		DependencyIndexes: file_protos_faas_req_res_proto_depIdxs,
		MessageInfos:      file_protos_faas_req_res_proto_msgTypes,
	}.Build()
	File_protos_faas_req_res_proto = out.File
	file_protos_faas_req_res_proto_rawDesc = nil
	file_protos_faas_req_res_proto_goTypes = nil
	file_protos_faas_req_res_proto_depIdxs = nil
}
