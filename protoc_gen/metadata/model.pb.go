// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.10
// source: protos/metadata/model.proto

package metadata

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

type ObjIndex_IndexType int32

const (
	ObjIndex_UNSET_NOTALLOW ObjIndex_IndexType = 0 // 占位, 不允许使用
	ObjIndex_COMMON         ObjIndex_IndexType = 1
	ObjIndex_UNIQUE         ObjIndex_IndexType = 2
)

// Enum value maps for ObjIndex_IndexType.
var (
	ObjIndex_IndexType_name = map[int32]string{
		0: "UNSET_NOTALLOW",
		1: "COMMON",
		2: "UNIQUE",
	}
	ObjIndex_IndexType_value = map[string]int32{
		"UNSET_NOTALLOW": 0,
		"COMMON":         1,
		"UNIQUE":         2,
	}
)

func (x ObjIndex_IndexType) Enum() *ObjIndex_IndexType {
	p := new(ObjIndex_IndexType)
	*p = x
	return p
}

func (x ObjIndex_IndexType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ObjIndex_IndexType) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_metadata_model_proto_enumTypes[0].Descriptor()
}

func (ObjIndex_IndexType) Type() protoreflect.EnumType {
	return &file_protos_metadata_model_proto_enumTypes[0]
}

func (x ObjIndex_IndexType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ObjIndex_IndexType.Descriptor instead.
func (ObjIndex_IndexType) EnumDescriptor() ([]byte, []int) {
	return file_protos_metadata_model_proto_rawDescGZIP(), []int{0, 0}
}

type ObjField_FieldType int32

const (
	ObjField_UNSET_NOTALLOW ObjField_FieldType = 0 // 占位, 不允许使用
	ObjField_ID             ObjField_FieldType = 1 // 本对象的id, 无符号自增int
	ObjField_LOOKUPID       ObjField_FieldType = 2 // 其他对象的id, 无符号int
	ObjField_INTEGER        ObjField_FieldType = 3 // int
	ObjField_DOUBLE         ObjField_FieldType = 4 // double
	ObjField_STRING         ObjField_FieldType = 5 // varchar(255)
	ObjField_TEXT           ObjField_FieldType = 6 // text
	ObjField_DATETIME       ObjField_FieldType = 7 // datetime
)

// Enum value maps for ObjField_FieldType.
var (
	ObjField_FieldType_name = map[int32]string{
		0: "UNSET_NOTALLOW",
		1: "ID",
		2: "LOOKUPID",
		3: "INTEGER",
		4: "DOUBLE",
		5: "STRING",
		6: "TEXT",
		7: "DATETIME",
	}
	ObjField_FieldType_value = map[string]int32{
		"UNSET_NOTALLOW": 0,
		"ID":             1,
		"LOOKUPID":       2,
		"INTEGER":        3,
		"DOUBLE":         4,
		"STRING":         5,
		"TEXT":           6,
		"DATETIME":       7,
	}
)

func (x ObjField_FieldType) Enum() *ObjField_FieldType {
	p := new(ObjField_FieldType)
	*p = x
	return p
}

func (x ObjField_FieldType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ObjField_FieldType) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_metadata_model_proto_enumTypes[1].Descriptor()
}

func (ObjField_FieldType) Type() protoreflect.EnumType {
	return &file_protos_metadata_model_proto_enumTypes[1]
}

func (x ObjField_FieldType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ObjField_FieldType.Descriptor instead.
func (ObjField_FieldType) EnumDescriptor() ([]byte, []int) {
	return file_protos_metadata_model_proto_rawDescGZIP(), []int{1, 0}
}

type ObjIndex struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjFieldNames []string           `protobuf:"bytes,1,rep,name=obj_field_names,json=objFieldNames,proto3" json:"obj_field_names,omitempty"`
	IndexType     ObjIndex_IndexType `protobuf:"varint,2,opt,name=index_type,json=indexType,proto3,enum=metadata.ObjIndex_IndexType" json:"index_type,omitempty"`
}

func (x *ObjIndex) Reset() {
	*x = ObjIndex{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_metadata_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjIndex) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjIndex) ProtoMessage() {}

func (x *ObjIndex) ProtoReflect() protoreflect.Message {
	mi := &file_protos_metadata_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjIndex.ProtoReflect.Descriptor instead.
func (*ObjIndex) Descriptor() ([]byte, []int) {
	return file_protos_metadata_model_proto_rawDescGZIP(), []int{0}
}

func (x *ObjIndex) GetObjFieldNames() []string {
	if x != nil {
		return x.ObjFieldNames
	}
	return nil
}

func (x *ObjIndex) GetIndexType() ObjIndex_IndexType {
	if x != nil {
		return x.IndexType
	}
	return ObjIndex_UNSET_NOTALLOW
}

type ObjField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	FieldType   ObjField_FieldType `protobuf:"varint,2,opt,name=field_type,json=fieldType,proto3,enum=metadata.ObjField_FieldType" json:"field_type,omitempty"`
	LookupObjId *int32             `protobuf:"varint,3,opt,name=lookup_obj_id,json=lookupObjId,proto3,oneof" json:"lookup_obj_id,omitempty"`
	Description *string            `protobuf:"bytes,4,opt,name=description,proto3,oneof" json:"description,omitempty"`
	IsAllowNull bool               `protobuf:"varint,5,opt,name=is_allow_null,json=isAllowNull,proto3" json:"is_allow_null,omitempty"`
}

func (x *ObjField) Reset() {
	*x = ObjField{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_metadata_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjField) ProtoMessage() {}

func (x *ObjField) ProtoReflect() protoreflect.Message {
	mi := &file_protos_metadata_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjField.ProtoReflect.Descriptor instead.
func (*ObjField) Descriptor() ([]byte, []int) {
	return file_protos_metadata_model_proto_rawDescGZIP(), []int{1}
}

func (x *ObjField) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ObjField) GetFieldType() ObjField_FieldType {
	if x != nil {
		return x.FieldType
	}
	return ObjField_UNSET_NOTALLOW
}

func (x *ObjField) GetLookupObjId() int32 {
	if x != nil && x.LookupObjId != nil {
		return *x.LookupObjId
	}
	return 0
}

func (x *ObjField) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *ObjField) GetIsAllowNull() bool {
	if x != nil {
		return x.IsAllowNull
	}
	return false
}

type ObjDao struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32            `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string           `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description *string          `protobuf:"bytes,3,opt,name=description,proto3,oneof" json:"description,omitempty"`
	TimeInfo    *common.TimeInfo `protobuf:"bytes,4,opt,name=time_info,json=timeInfo,proto3" json:"time_info,omitempty"`
}

func (x *ObjDao) Reset() {
	*x = ObjDao{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_metadata_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjDao) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjDao) ProtoMessage() {}

func (x *ObjDao) ProtoReflect() protoreflect.Message {
	mi := &file_protos_metadata_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjDao.ProtoReflect.Descriptor instead.
func (*ObjDao) Descriptor() ([]byte, []int) {
	return file_protos_metadata_model_proto_rawDescGZIP(), []int{2}
}

func (x *ObjDao) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ObjDao) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ObjDao) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *ObjDao) GetTimeInfo() *common.TimeInfo {
	if x != nil {
		return x.TimeInfo
	}
	return nil
}

type ObjFieldDao struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32              `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ObjId       int32              `protobuf:"varint,2,opt,name=obj_id,json=objId,proto3" json:"obj_id,omitempty"`
	Name        int32              `protobuf:"varint,3,opt,name=name,proto3" json:"name,omitempty"`
	FieldType   ObjField_FieldType `protobuf:"varint,4,opt,name=field_type,json=fieldType,proto3,enum=metadata.ObjField_FieldType" json:"field_type,omitempty"`
	LookupObjId *int32             `protobuf:"varint,5,opt,name=lookup_obj_id,json=lookupObjId,proto3,oneof" json:"lookup_obj_id,omitempty"`
	Description *string            `protobuf:"bytes,6,opt,name=description,proto3,oneof" json:"description,omitempty"`
	IsAllowNull bool               `protobuf:"varint,7,opt,name=is_allow_null,json=isAllowNull,proto3" json:"is_allow_null,omitempty"`
	TimeInfo    *common.TimeInfo   `protobuf:"bytes,8,opt,name=time_info,json=timeInfo,proto3" json:"time_info,omitempty"`
}

func (x *ObjFieldDao) Reset() {
	*x = ObjFieldDao{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_metadata_model_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjFieldDao) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjFieldDao) ProtoMessage() {}

func (x *ObjFieldDao) ProtoReflect() protoreflect.Message {
	mi := &file_protos_metadata_model_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjFieldDao.ProtoReflect.Descriptor instead.
func (*ObjFieldDao) Descriptor() ([]byte, []int) {
	return file_protos_metadata_model_proto_rawDescGZIP(), []int{3}
}

func (x *ObjFieldDao) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ObjFieldDao) GetObjId() int32 {
	if x != nil {
		return x.ObjId
	}
	return 0
}

func (x *ObjFieldDao) GetName() int32 {
	if x != nil {
		return x.Name
	}
	return 0
}

func (x *ObjFieldDao) GetFieldType() ObjField_FieldType {
	if x != nil {
		return x.FieldType
	}
	return ObjField_UNSET_NOTALLOW
}

func (x *ObjFieldDao) GetLookupObjId() int32 {
	if x != nil && x.LookupObjId != nil {
		return *x.LookupObjId
	}
	return 0
}

func (x *ObjFieldDao) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *ObjFieldDao) GetIsAllowNull() bool {
	if x != nil {
		return x.IsAllowNull
	}
	return false
}

func (x *ObjFieldDao) GetTimeInfo() *common.TimeInfo {
	if x != nil {
		return x.TimeInfo
	}
	return nil
}

type ObjDaoWithFieldsDao struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Obj    *ObjDao        `protobuf:"bytes,1,opt,name=obj,proto3" json:"obj,omitempty"`
	Fields []*ObjFieldDao `protobuf:"bytes,2,rep,name=fields,proto3" json:"fields,omitempty"`
}

func (x *ObjDaoWithFieldsDao) Reset() {
	*x = ObjDaoWithFieldsDao{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_metadata_model_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjDaoWithFieldsDao) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjDaoWithFieldsDao) ProtoMessage() {}

func (x *ObjDaoWithFieldsDao) ProtoReflect() protoreflect.Message {
	mi := &file_protos_metadata_model_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjDaoWithFieldsDao.ProtoReflect.Descriptor instead.
func (*ObjDaoWithFieldsDao) Descriptor() ([]byte, []int) {
	return file_protos_metadata_model_proto_rawDescGZIP(), []int{4}
}

func (x *ObjDaoWithFieldsDao) GetObj() *ObjDao {
	if x != nil {
		return x.Obj
	}
	return nil
}

func (x *ObjDaoWithFieldsDao) GetFields() []*ObjFieldDao {
	if x != nil {
		return x.Fields
	}
	return nil
}

var File_protos_metadata_model_proto protoreflect.FileDescriptor

var file_protos_metadata_model_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xa8, 0x01, 0x0a, 0x08, 0x4f, 0x62, 0x6a, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x26, 0x0a, 0x0f, 0x6f, 0x62, 0x6a, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x6f, 0x62, 0x6a, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x3b, 0x0a, 0x0a, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4f, 0x62, 0x6a, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x54, 0x79, 0x70, 0x65, 0x22, 0x37, 0x0a, 0x09, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x4e, 0x53, 0x45, 0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x41,
	0x4c, 0x4c, 0x4f, 0x57, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x4f, 0x4d, 0x4d, 0x4f, 0x4e,
	0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x4e, 0x49, 0x51, 0x55, 0x45, 0x10, 0x02, 0x22, 0xe5,
	0x02, 0x0a, 0x08, 0x4f, 0x62, 0x6a, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x3b, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4f,
	0x62, 0x6a, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x27, 0x0a, 0x0d,
	0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x5f, 0x6f, 0x62, 0x6a, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x0b, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x4f, 0x62, 0x6a,
	0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0d,
	0x69, 0x73, 0x5f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x6e, 0x75, 0x6c, 0x6c, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x4e, 0x75, 0x6c, 0x6c,
	0x22, 0x72, 0x0a, 0x09, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a,
	0x0e, 0x55, 0x4e, 0x53, 0x45, 0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x10,
	0x00, 0x12, 0x06, 0x0a, 0x02, 0x49, 0x44, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x4c, 0x4f, 0x4f,
	0x4b, 0x55, 0x50, 0x49, 0x44, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x54, 0x45, 0x47,
	0x45, 0x52, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x4f, 0x55, 0x42, 0x4c, 0x45, 0x10, 0x04,
	0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x05, 0x12, 0x08, 0x0a, 0x04,
	0x54, 0x45, 0x58, 0x54, 0x10, 0x06, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x41, 0x54, 0x45, 0x54, 0x49,
	0x4d, 0x45, 0x10, 0x07, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x5f,
	0x6f, 0x62, 0x6a, 0x5f, 0x69, 0x64, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x92, 0x01, 0x0a, 0x06, 0x4f, 0x62, 0x6a, 0x44, 0x61,
	0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x0e, 0x0a, 0x0c, 0x5f,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xca, 0x02, 0x0a, 0x0b,
	0x4f, 0x62, 0x6a, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x44, 0x61, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x6f,
	0x62, 0x6a, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6f, 0x62, 0x6a,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4f, 0x62, 0x6a, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x27, 0x0a, 0x0d, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x5f, 0x6f, 0x62,
	0x6a, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x0b, 0x6c, 0x6f,
	0x6f, 0x6b, 0x75, 0x70, 0x4f, 0x62, 0x6a, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x01, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f,
	0x6e, 0x75, 0x6c, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x41, 0x6c,
	0x6c, 0x6f, 0x77, 0x4e, 0x75, 0x6c, 0x6c, 0x12, 0x2d, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x74, 0x69,
	0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x6c, 0x6f, 0x6f, 0x6b, 0x75,
	0x70, 0x5f, 0x6f, 0x62, 0x6a, 0x5f, 0x69, 0x64, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x68, 0x0a, 0x13, 0x4f, 0x62, 0x6a, 0x44,
	0x61, 0x6f, 0x57, 0x69, 0x74, 0x68, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x44, 0x61, 0x6f, 0x12,
	0x22, 0x0a, 0x03, 0x6f, 0x62, 0x6a, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4f, 0x62, 0x6a, 0x44, 0x61, 0x6f, 0x52, 0x03,
	0x6f, 0x62, 0x6a, 0x12, 0x2d, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4f,
	0x62, 0x6a, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x44, 0x61, 0x6f, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x59, 0x6f, 0x73, 0x6f, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x2f, 0x6d, 0x73, 0x2d, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x5f, 0x67, 0x65, 0x6e, 0x2f,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_metadata_model_proto_rawDescOnce sync.Once
	file_protos_metadata_model_proto_rawDescData = file_protos_metadata_model_proto_rawDesc
)

func file_protos_metadata_model_proto_rawDescGZIP() []byte {
	file_protos_metadata_model_proto_rawDescOnce.Do(func() {
		file_protos_metadata_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_metadata_model_proto_rawDescData)
	})
	return file_protos_metadata_model_proto_rawDescData
}

var file_protos_metadata_model_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_protos_metadata_model_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protos_metadata_model_proto_goTypes = []interface{}{
	(ObjIndex_IndexType)(0),     // 0: metadata.ObjIndex.IndexType
	(ObjField_FieldType)(0),     // 1: metadata.ObjField.FieldType
	(*ObjIndex)(nil),            // 2: metadata.ObjIndex
	(*ObjField)(nil),            // 3: metadata.ObjField
	(*ObjDao)(nil),              // 4: metadata.ObjDao
	(*ObjFieldDao)(nil),         // 5: metadata.ObjFieldDao
	(*ObjDaoWithFieldsDao)(nil), // 6: metadata.ObjDaoWithFieldsDao
	(*common.TimeInfo)(nil),     // 7: common.TimeInfo
}
var file_protos_metadata_model_proto_depIdxs = []int32{
	0, // 0: metadata.ObjIndex.index_type:type_name -> metadata.ObjIndex.IndexType
	1, // 1: metadata.ObjField.field_type:type_name -> metadata.ObjField.FieldType
	7, // 2: metadata.ObjDao.time_info:type_name -> common.TimeInfo
	1, // 3: metadata.ObjFieldDao.field_type:type_name -> metadata.ObjField.FieldType
	7, // 4: metadata.ObjFieldDao.time_info:type_name -> common.TimeInfo
	4, // 5: metadata.ObjDaoWithFieldsDao.obj:type_name -> metadata.ObjDao
	5, // 6: metadata.ObjDaoWithFieldsDao.fields:type_name -> metadata.ObjFieldDao
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_protos_metadata_model_proto_init() }
func file_protos_metadata_model_proto_init() {
	if File_protos_metadata_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_metadata_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjIndex); i {
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
		file_protos_metadata_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjField); i {
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
		file_protos_metadata_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjDao); i {
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
		file_protos_metadata_model_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjFieldDao); i {
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
		file_protos_metadata_model_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjDaoWithFieldsDao); i {
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
	file_protos_metadata_model_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_protos_metadata_model_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_protos_metadata_model_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_metadata_model_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_metadata_model_proto_goTypes,
		DependencyIndexes: file_protos_metadata_model_proto_depIdxs,
		EnumInfos:         file_protos_metadata_model_proto_enumTypes,
		MessageInfos:      file_protos_metadata_model_proto_msgTypes,
	}.Build()
	File_protos_metadata_model_proto = out.File
	file_protos_metadata_model_proto_rawDesc = nil
	file_protos_metadata_model_proto_goTypes = nil
	file_protos_metadata_model_proto_depIdxs = nil
}
