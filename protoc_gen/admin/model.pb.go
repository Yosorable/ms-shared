// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.11
// source: protos/admin/model.proto

package admin

import (
	common "github.com/Yosorable/ms-shared/protoc_gen/common"
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

type OrderBy_Order int32

const (
	OrderBy_ASC  OrderBy_Order = 0
	OrderBy_DESC OrderBy_Order = 1
)

// Enum value maps for OrderBy_Order.
var (
	OrderBy_Order_name = map[int32]string{
		0: "ASC",
		1: "DESC",
	}
	OrderBy_Order_value = map[string]int32{
		"ASC":  0,
		"DESC": 1,
	}
)

func (x OrderBy_Order) Enum() *OrderBy_Order {
	p := new(OrderBy_Order)
	*p = x
	return p
}

func (x OrderBy_Order) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderBy_Order) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_admin_model_proto_enumTypes[0].Descriptor()
}

func (OrderBy_Order) Type() protoreflect.EnumType {
	return &file_protos_admin_model_proto_enumTypes[0]
}

func (x OrderBy_Order) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderBy_Order.Descriptor instead.
func (OrderBy_Order) EnumDescriptor() ([]byte, []int) {
	return file_protos_admin_model_proto_rawDescGZIP(), []int{2, 0}
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32            `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username string           `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	TimeInfo *common.TimeInfo `protobuf:"bytes,3,opt,name=time_info,json=timeInfo,proto3" json:"time_info,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_admin_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_protos_admin_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_protos_admin_model_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetTimeInfo() *common.TimeInfo {
	if x != nil {
		return x.TimeInfo
	}
	return nil
}

type UserRecordTableOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceName   string `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	TableTag      string `protobuf:"bytes,2,opt,name=table_tag,json=tableTag,proto3" json:"table_tag,omitempty"`
	ForeignIdName string `protobuf:"bytes,3,opt,name=foreign_id_name,json=foreignIdName,proto3" json:"foreign_id_name,omitempty"`
}

func (x *UserRecordTableOption) Reset() {
	*x = UserRecordTableOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_admin_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRecordTableOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRecordTableOption) ProtoMessage() {}

func (x *UserRecordTableOption) ProtoReflect() protoreflect.Message {
	mi := &file_protos_admin_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRecordTableOption.ProtoReflect.Descriptor instead.
func (*UserRecordTableOption) Descriptor() ([]byte, []int) {
	return file_protos_admin_model_proto_rawDescGZIP(), []int{1}
}

func (x *UserRecordTableOption) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *UserRecordTableOption) GetTableTag() string {
	if x != nil {
		return x.TableTag
	}
	return ""
}

func (x *UserRecordTableOption) GetForeignIdName() string {
	if x != nil {
		return x.ForeignIdName
	}
	return ""
}

type OrderBy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderByFiledName string        `protobuf:"bytes,1,opt,name=order_by_filed_name,json=orderByFiledName,proto3" json:"order_by_filed_name,omitempty"`
	Order            OrderBy_Order `protobuf:"varint,2,opt,name=order,proto3,enum=admin.OrderBy_Order" json:"order,omitempty"`
}

func (x *OrderBy) Reset() {
	*x = OrderBy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_admin_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderBy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderBy) ProtoMessage() {}

func (x *OrderBy) ProtoReflect() protoreflect.Message {
	mi := &file_protos_admin_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderBy.ProtoReflect.Descriptor instead.
func (*OrderBy) Descriptor() ([]byte, []int) {
	return file_protos_admin_model_proto_rawDescGZIP(), []int{2}
}

func (x *OrderBy) GetOrderByFiledName() string {
	if x != nil {
		return x.OrderByFiledName
	}
	return ""
}

func (x *OrderBy) GetOrder() OrderBy_Order {
	if x != nil {
		return x.Order
	}
	return OrderBy_ASC
}

type PageOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *PageOption) Reset() {
	*x = PageOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_admin_model_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageOption) ProtoMessage() {}

func (x *PageOption) ProtoReflect() protoreflect.Message {
	mi := &file_protos_admin_model_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageOption.ProtoReflect.Descriptor instead.
func (*PageOption) Descriptor() ([]byte, []int) {
	return file_protos_admin_model_proto_rawDescGZIP(), []int{3}
}

func (x *PageOption) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *PageOption) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type UserRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    int32  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ForeignId int32  `protobuf:"varint,2,opt,name=foreign_id,json=foreignId,proto3" json:"foreign_id,omitempty"`
	CreatedAt string `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *UserRecord) Reset() {
	*x = UserRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_admin_model_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRecord) ProtoMessage() {}

func (x *UserRecord) ProtoReflect() protoreflect.Message {
	mi := &file_protos_admin_model_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRecord.ProtoReflect.Descriptor instead.
func (*UserRecord) Descriptor() ([]byte, []int) {
	return file_protos_admin_model_proto_rawDescGZIP(), []int{4}
}

func (x *UserRecord) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserRecord) GetForeignId() int32 {
	if x != nil {
		return x.ForeignId
	}
	return 0
}

func (x *UserRecord) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *UserRecord) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

var File_protos_admin_model_proto protoreflect.FileDescriptor

var file_protos_admin_model_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x61, 0x0a,
	0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x2d, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x22, 0x7f, 0x0a, 0x15, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x61,
	0x62, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x74, 0x61, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x61, 0x67, 0x12, 0x26, 0x0a, 0x0f, 0x66, 0x6f, 0x72,
	0x65, 0x69, 0x67, 0x6e, 0x5f, 0x69, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x66, 0x6f, 0x72, 0x65, 0x69, 0x67, 0x6e, 0x49, 0x64, 0x4e, 0x61, 0x6d,
	0x65, 0x22, 0x80, 0x01, 0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x12, 0x2d, 0x0a,
	0x13, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x64, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x42, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x05,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x1a, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x53, 0x43, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x45,
	0x53, 0x43, 0x10, 0x01, 0x22, 0x3a, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x22, 0x82, 0x01, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x6f, 0x72, 0x65,
	0x69, 0x67, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x66, 0x6f,
	0x72, 0x65, 0x69, 0x67, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x59, 0x6f, 0x73, 0x6f, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x2f, 0x6d, 0x73,
	0x2d, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x5f, 0x67,
	0x65, 0x6e, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_admin_model_proto_rawDescOnce sync.Once
	file_protos_admin_model_proto_rawDescData = file_protos_admin_model_proto_rawDesc
)

func file_protos_admin_model_proto_rawDescGZIP() []byte {
	file_protos_admin_model_proto_rawDescOnce.Do(func() {
		file_protos_admin_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_admin_model_proto_rawDescData)
	})
	return file_protos_admin_model_proto_rawDescData
}

var file_protos_admin_model_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protos_admin_model_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protos_admin_model_proto_goTypes = []interface{}{
	(OrderBy_Order)(0),            // 0: admin.OrderBy.Order
	(*User)(nil),                  // 1: admin.User
	(*UserRecordTableOption)(nil), // 2: admin.UserRecordTableOption
	(*OrderBy)(nil),               // 3: admin.OrderBy
	(*PageOption)(nil),            // 4: admin.PageOption
	(*UserRecord)(nil),            // 5: admin.UserRecord
	(*common.TimeInfo)(nil),       // 6: common.TimeInfo
}
var file_protos_admin_model_proto_depIdxs = []int32{
	6, // 0: admin.User.time_info:type_name -> common.TimeInfo
	0, // 1: admin.OrderBy.order:type_name -> admin.OrderBy.Order
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protos_admin_model_proto_init() }
func file_protos_admin_model_proto_init() {
	if File_protos_admin_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_admin_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_protos_admin_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRecordTableOption); i {
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
		file_protos_admin_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderBy); i {
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
		file_protos_admin_model_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageOption); i {
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
		file_protos_admin_model_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRecord); i {
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
			RawDescriptor: file_protos_admin_model_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_admin_model_proto_goTypes,
		DependencyIndexes: file_protos_admin_model_proto_depIdxs,
		EnumInfos:         file_protos_admin_model_proto_enumTypes,
		MessageInfos:      file_protos_admin_model_proto_msgTypes,
	}.Build()
	File_protos_admin_model_proto = out.File
	file_protos_admin_model_proto_rawDesc = nil
	file_protos_admin_model_proto_goTypes = nil
	file_protos_admin_model_proto_depIdxs = nil
}
