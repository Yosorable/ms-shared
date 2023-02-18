// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.10
// source: protos/faas/service.proto

package faas

import (
	context "context"
	common "github.com/Yosorable/ms-shared/protoc_gen/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FaasClient is the client API for Faas service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FaasClient interface {
	DynamicCall(ctx context.Context, in *common.DynamicCallRequest, opts ...grpc.CallOption) (*common.DynamicCallReply, error)
	CreateFaasTask(ctx context.Context, in *CreateFaasTaskRequest, opts ...grpc.CallOption) (*CreateFaasTaskReply, error)
	GetFaasTaskByPage(ctx context.Context, in *GetFaasTaskByPageRequest, opts ...grpc.CallOption) (*GetFaasTaskByPageReply, error)
	GetFaasTaskByID(ctx context.Context, in *GetFaasTaskByIDRequest, opts ...grpc.CallOption) (*GetFaasTaskByIDReply, error)
	UpdateFaasTaskByID(ctx context.Context, in *UpdateFaasTaskByIDRequest, opts ...grpc.CallOption) (*UpdateFaasTaskByIDReply, error)
	RunTask(ctx context.Context, in *RunTaskRequest, opts ...grpc.CallOption) (*RunTaskReply, error)
}

type faasClient struct {
	cc grpc.ClientConnInterface
}

func NewFaasClient(cc grpc.ClientConnInterface) FaasClient {
	return &faasClient{cc}
}

func (c *faasClient) DynamicCall(ctx context.Context, in *common.DynamicCallRequest, opts ...grpc.CallOption) (*common.DynamicCallReply, error) {
	out := new(common.DynamicCallReply)
	err := c.cc.Invoke(ctx, "/faas.Faas/DynamicCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *faasClient) CreateFaasTask(ctx context.Context, in *CreateFaasTaskRequest, opts ...grpc.CallOption) (*CreateFaasTaskReply, error) {
	out := new(CreateFaasTaskReply)
	err := c.cc.Invoke(ctx, "/faas.Faas/CreateFaasTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *faasClient) GetFaasTaskByPage(ctx context.Context, in *GetFaasTaskByPageRequest, opts ...grpc.CallOption) (*GetFaasTaskByPageReply, error) {
	out := new(GetFaasTaskByPageReply)
	err := c.cc.Invoke(ctx, "/faas.Faas/GetFaasTaskByPage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *faasClient) GetFaasTaskByID(ctx context.Context, in *GetFaasTaskByIDRequest, opts ...grpc.CallOption) (*GetFaasTaskByIDReply, error) {
	out := new(GetFaasTaskByIDReply)
	err := c.cc.Invoke(ctx, "/faas.Faas/GetFaasTaskByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *faasClient) UpdateFaasTaskByID(ctx context.Context, in *UpdateFaasTaskByIDRequest, opts ...grpc.CallOption) (*UpdateFaasTaskByIDReply, error) {
	out := new(UpdateFaasTaskByIDReply)
	err := c.cc.Invoke(ctx, "/faas.Faas/UpdateFaasTaskByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *faasClient) RunTask(ctx context.Context, in *RunTaskRequest, opts ...grpc.CallOption) (*RunTaskReply, error) {
	out := new(RunTaskReply)
	err := c.cc.Invoke(ctx, "/faas.Faas/RunTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FaasServer is the server API for Faas service.
// All implementations must embed UnimplementedFaasServer
// for forward compatibility
type FaasServer interface {
	DynamicCall(context.Context, *common.DynamicCallRequest) (*common.DynamicCallReply, error)
	CreateFaasTask(context.Context, *CreateFaasTaskRequest) (*CreateFaasTaskReply, error)
	GetFaasTaskByPage(context.Context, *GetFaasTaskByPageRequest) (*GetFaasTaskByPageReply, error)
	GetFaasTaskByID(context.Context, *GetFaasTaskByIDRequest) (*GetFaasTaskByIDReply, error)
	UpdateFaasTaskByID(context.Context, *UpdateFaasTaskByIDRequest) (*UpdateFaasTaskByIDReply, error)
	RunTask(context.Context, *RunTaskRequest) (*RunTaskReply, error)
	mustEmbedUnimplementedFaasServer()
}

// UnimplementedFaasServer must be embedded to have forward compatible implementations.
type UnimplementedFaasServer struct {
}

func (UnimplementedFaasServer) DynamicCall(context.Context, *common.DynamicCallRequest) (*common.DynamicCallReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DynamicCall not implemented")
}
func (UnimplementedFaasServer) CreateFaasTask(context.Context, *CreateFaasTaskRequest) (*CreateFaasTaskReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFaasTask not implemented")
}
func (UnimplementedFaasServer) GetFaasTaskByPage(context.Context, *GetFaasTaskByPageRequest) (*GetFaasTaskByPageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFaasTaskByPage not implemented")
}
func (UnimplementedFaasServer) GetFaasTaskByID(context.Context, *GetFaasTaskByIDRequest) (*GetFaasTaskByIDReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFaasTaskByID not implemented")
}
func (UnimplementedFaasServer) UpdateFaasTaskByID(context.Context, *UpdateFaasTaskByIDRequest) (*UpdateFaasTaskByIDReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFaasTaskByID not implemented")
}
func (UnimplementedFaasServer) RunTask(context.Context, *RunTaskRequest) (*RunTaskReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunTask not implemented")
}
func (UnimplementedFaasServer) mustEmbedUnimplementedFaasServer() {}

// UnsafeFaasServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FaasServer will
// result in compilation errors.
type UnsafeFaasServer interface {
	mustEmbedUnimplementedFaasServer()
}

func RegisterFaasServer(s grpc.ServiceRegistrar, srv FaasServer) {
	s.RegisterService(&Faas_ServiceDesc, srv)
}

func _Faas_DynamicCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.DynamicCallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FaasServer).DynamicCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/faas.Faas/DynamicCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FaasServer).DynamicCall(ctx, req.(*common.DynamicCallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Faas_CreateFaasTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFaasTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FaasServer).CreateFaasTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/faas.Faas/CreateFaasTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FaasServer).CreateFaasTask(ctx, req.(*CreateFaasTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Faas_GetFaasTaskByPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFaasTaskByPageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FaasServer).GetFaasTaskByPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/faas.Faas/GetFaasTaskByPage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FaasServer).GetFaasTaskByPage(ctx, req.(*GetFaasTaskByPageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Faas_GetFaasTaskByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFaasTaskByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FaasServer).GetFaasTaskByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/faas.Faas/GetFaasTaskByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FaasServer).GetFaasTaskByID(ctx, req.(*GetFaasTaskByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Faas_UpdateFaasTaskByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFaasTaskByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FaasServer).UpdateFaasTaskByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/faas.Faas/UpdateFaasTaskByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FaasServer).UpdateFaasTaskByID(ctx, req.(*UpdateFaasTaskByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Faas_RunTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FaasServer).RunTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/faas.Faas/RunTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FaasServer).RunTask(ctx, req.(*RunTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Faas_ServiceDesc is the grpc.ServiceDesc for Faas service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Faas_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "faas.Faas",
	HandlerType: (*FaasServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DynamicCall",
			Handler:    _Faas_DynamicCall_Handler,
		},
		{
			MethodName: "CreateFaasTask",
			Handler:    _Faas_CreateFaasTask_Handler,
		},
		{
			MethodName: "GetFaasTaskByPage",
			Handler:    _Faas_GetFaasTaskByPage_Handler,
		},
		{
			MethodName: "GetFaasTaskByID",
			Handler:    _Faas_GetFaasTaskByID_Handler,
		},
		{
			MethodName: "UpdateFaasTaskByID",
			Handler:    _Faas_UpdateFaasTaskByID_Handler,
		},
		{
			MethodName: "RunTask",
			Handler:    _Faas_RunTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/faas/service.proto",
}
