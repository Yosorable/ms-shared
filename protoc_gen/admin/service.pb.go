// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.10
// source: protos/admin/service.proto

package admin

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_protos_admin_service_proto protoreflect.FileDescriptor

var file_protos_admin_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2f, 0x72, 0x65, 0x71, 0x5f, 0x72, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0x8f, 0x05, 0x0a, 0x05, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x31, 0x0a, 0x05, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x12, 0x13, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0a,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x18, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3a,
	0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x12, 0x19, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12,
	0x7f, 0x0a, 0x1f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x66, 0x4e, 0x6f, 0x74, 0x45, 0x78, 0x69,
	0x73, 0x74, 0x12, 0x2d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x49, 0x66, 0x4e, 0x6f, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2b, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x49,
	0x66, 0x4e, 0x6f, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x12, 0x4f, 0x0a, 0x0f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x12, 0x1d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x12, 0x6a, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x26, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x52, 0x0a,
	0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x12, 0x1e, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x59, 0x6f, 0x73, 0x6f, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x2f, 0x6d, 0x73, 0x2d, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_protos_admin_service_proto_goTypes = []interface{}{
	(*LoginRequest)(nil),                           // 0: admin.LoginRequest
	(*CheckTokenRequest)(nil),                      // 1: admin.CheckTokenRequest
	(*RegisterRequest)(nil),                        // 2: admin.RegisterRequest
	(*GetUserByIDRequest)(nil),                     // 3: admin.GetUserByIDRequest
	(*CreateUserRecordTableIfNotExistRequest)(nil), // 4: admin.CreateUserRecordTableIfNotExistRequest
	(*QueryUserRecordRequest)(nil),                 // 5: admin.QueryUserRecordRequest
	(*CreateOrUpdateUserRecordRequest)(nil),        // 6: admin.CreateOrUpdateUserRecordRequest
	(*DeleteUserRecordRequest)(nil),                // 7: admin.DeleteUserRecordRequest
	(*LoginReply)(nil),                             // 8: admin.LoginReply
	(*CheckTokenReply)(nil),                        // 9: admin.CheckTokenReply
	(*RegisterReply)(nil),                          // 10: admin.RegisterReply
	(*GetUserByIDReply)(nil),                       // 11: admin.GetUserByIDReply
	(*CreateUserRecordTableIfNotExistReply)(nil),   // 12: admin.CreateUserRecordTableIfNotExistReply
	(*QueryUserRecordReply)(nil),                   // 13: admin.QueryUserRecordReply
	(*CreateOrUpdateUserRecordReply)(nil),          // 14: admin.CreateOrUpdateUserRecordReply
	(*DeleteUserRecordReply)(nil),                  // 15: admin.DeleteUserRecordReply
}
var file_protos_admin_service_proto_depIdxs = []int32{
	0,  // 0: admin.Admin.Login:input_type -> admin.LoginRequest
	1,  // 1: admin.Admin.CheckToken:input_type -> admin.CheckTokenRequest
	2,  // 2: admin.Admin.Register:input_type -> admin.RegisterRequest
	3,  // 3: admin.Admin.GetUserByID:input_type -> admin.GetUserByIDRequest
	4,  // 4: admin.Admin.CreateUserRecordTableIfNotExist:input_type -> admin.CreateUserRecordTableIfNotExistRequest
	5,  // 5: admin.Admin.QueryUserRecord:input_type -> admin.QueryUserRecordRequest
	6,  // 6: admin.Admin.CreateOrUpdateUserRecord:input_type -> admin.CreateOrUpdateUserRecordRequest
	7,  // 7: admin.Admin.DeleteUserRecord:input_type -> admin.DeleteUserRecordRequest
	8,  // 8: admin.Admin.Login:output_type -> admin.LoginReply
	9,  // 9: admin.Admin.CheckToken:output_type -> admin.CheckTokenReply
	10, // 10: admin.Admin.Register:output_type -> admin.RegisterReply
	11, // 11: admin.Admin.GetUserByID:output_type -> admin.GetUserByIDReply
	12, // 12: admin.Admin.CreateUserRecordTableIfNotExist:output_type -> admin.CreateUserRecordTableIfNotExistReply
	13, // 13: admin.Admin.QueryUserRecord:output_type -> admin.QueryUserRecordReply
	14, // 14: admin.Admin.CreateOrUpdateUserRecord:output_type -> admin.CreateOrUpdateUserRecordReply
	15, // 15: admin.Admin.DeleteUserRecord:output_type -> admin.DeleteUserRecordReply
	8,  // [8:16] is the sub-list for method output_type
	0,  // [0:8] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_protos_admin_service_proto_init() }
func file_protos_admin_service_proto_init() {
	if File_protos_admin_service_proto != nil {
		return
	}
	file_protos_admin_req_res_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_admin_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_admin_service_proto_goTypes,
		DependencyIndexes: file_protos_admin_service_proto_depIdxs,
	}.Build()
	File_protos_admin_service_proto = out.File
	file_protos_admin_service_proto_rawDesc = nil
	file_protos_admin_service_proto_goTypes = nil
	file_protos_admin_service_proto_depIdxs = nil
}