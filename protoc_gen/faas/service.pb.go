// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.10
// source: protos/faas/service.proto

package faas

import (
	common "github.com/Yosorable/ms-shared/protoc_gen/common"
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

var File_protos_faas_service_proto protoreflect.FileDescriptor

var file_protos_faas_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x66, 0x61, 0x61, 0x73, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x66, 0x61, 0x61,
	0x73, 0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x66, 0x61, 0x61, 0x73, 0x2f, 0x72, 0x65, 0x71, 0x5f, 0x72,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xf4, 0x02, 0x0a, 0x04, 0x46, 0x61, 0x61,
	0x73, 0x12, 0x45, 0x0a, 0x0b, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x43, 0x61, 0x6c, 0x6c,
	0x12, 0x1a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69,
	0x63, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x43, 0x61, 0x6c,
	0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x46, 0x61, 0x61, 0x73, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x1b, 0x2e, 0x66, 0x61, 0x61,
	0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x61, 0x61, 0x73, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x66, 0x61, 0x61, 0x73, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x61, 0x61, 0x73, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x46, 0x61, 0x61, 0x73, 0x54,
	0x61, 0x73, 0x6b, 0x42, 0x79, 0x50, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x2e, 0x66, 0x61, 0x61, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x46, 0x61, 0x61, 0x73, 0x54, 0x61, 0x73, 0x6b, 0x42, 0x79, 0x50, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x66, 0x61, 0x61, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x46, 0x61, 0x61, 0x73, 0x54, 0x61, 0x73, 0x6b, 0x42, 0x79, 0x50, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x46, 0x61, 0x61, 0x73, 0x54, 0x61, 0x73, 0x6b, 0x42, 0x79, 0x49, 0x44, 0x12, 0x1c, 0x2e, 0x66,
	0x61, 0x61, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x61, 0x61, 0x73, 0x54, 0x61, 0x73, 0x6b, 0x42,
	0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x66, 0x61, 0x61,
	0x73, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x61, 0x61, 0x73, 0x54, 0x61, 0x73, 0x6b, 0x42, 0x79, 0x49,
	0x44, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x07, 0x52, 0x75, 0x6e, 0x54,
	0x61, 0x73, 0x6b, 0x12, 0x14, 0x2e, 0x66, 0x61, 0x61, 0x73, 0x2e, 0x52, 0x75, 0x6e, 0x54, 0x61,
	0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x66, 0x61, 0x61, 0x73,
	0x2e, 0x52, 0x75, 0x6e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42,
	0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x59, 0x6f,
	0x73, 0x6f, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x2f, 0x6d, 0x73, 0x2d, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x66, 0x61, 0x61,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_protos_faas_service_proto_goTypes = []interface{}{
	(*common.DynamicCallRequest)(nil), // 0: common.DynamicCallRequest
	(*CreateFaasTaskRequest)(nil),     // 1: faas.CreateFaasTaskRequest
	(*GetFaasTaskByPageRequest)(nil),  // 2: faas.GetFaasTaskByPageRequest
	(*GetFaasTaskByIDRequest)(nil),    // 3: faas.GetFaasTaskByIDRequest
	(*RunTaskRequest)(nil),            // 4: faas.RunTaskRequest
	(*common.DynamicCallReply)(nil),   // 5: common.DynamicCallReply
	(*CreateFaasTaskReply)(nil),       // 6: faas.CreateFaasTaskReply
	(*GetFaasTaskByPageReply)(nil),    // 7: faas.GetFaasTaskByPageReply
	(*GetFaasTaskByIDReply)(nil),      // 8: faas.GetFaasTaskByIDReply
	(*RunTaskReply)(nil),              // 9: faas.RunTaskReply
}
var file_protos_faas_service_proto_depIdxs = []int32{
	0, // 0: faas.Faas.DynamicCall:input_type -> common.DynamicCallRequest
	1, // 1: faas.Faas.CreateFaasTask:input_type -> faas.CreateFaasTaskRequest
	2, // 2: faas.Faas.GetFaasTaskByPage:input_type -> faas.GetFaasTaskByPageRequest
	3, // 3: faas.Faas.GetFaasTaskByID:input_type -> faas.GetFaasTaskByIDRequest
	4, // 4: faas.Faas.RunTask:input_type -> faas.RunTaskRequest
	5, // 5: faas.Faas.DynamicCall:output_type -> common.DynamicCallReply
	6, // 6: faas.Faas.CreateFaasTask:output_type -> faas.CreateFaasTaskReply
	7, // 7: faas.Faas.GetFaasTaskByPage:output_type -> faas.GetFaasTaskByPageReply
	8, // 8: faas.Faas.GetFaasTaskByID:output_type -> faas.GetFaasTaskByIDReply
	9, // 9: faas.Faas.RunTask:output_type -> faas.RunTaskReply
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_faas_service_proto_init() }
func file_protos_faas_service_proto_init() {
	if File_protos_faas_service_proto != nil {
		return
	}
	file_protos_faas_req_res_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_faas_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_faas_service_proto_goTypes,
		DependencyIndexes: file_protos_faas_service_proto_depIdxs,
	}.Build()
	File_protos_faas_service_proto = out.File
	file_protos_faas_service_proto_rawDesc = nil
	file_protos_faas_service_proto_goTypes = nil
	file_protos_faas_service_proto_depIdxs = nil
}