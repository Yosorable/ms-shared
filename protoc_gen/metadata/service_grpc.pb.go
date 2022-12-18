// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.10
// source: protos/metadata/service.proto

package metadata

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MetadataClient is the client API for Metadata service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MetadataClient interface {
	SayHello(ctx context.Context, in *SayHelloRequest, opts ...grpc.CallOption) (*SayHelloReply, error)
}

type metadataClient struct {
	cc grpc.ClientConnInterface
}

func NewMetadataClient(cc grpc.ClientConnInterface) MetadataClient {
	return &metadataClient{cc}
}

func (c *metadataClient) SayHello(ctx context.Context, in *SayHelloRequest, opts ...grpc.CallOption) (*SayHelloReply, error) {
	out := new(SayHelloReply)
	err := c.cc.Invoke(ctx, "/metadata.Metadata/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetadataServer is the server API for Metadata service.
// All implementations must embed UnimplementedMetadataServer
// for forward compatibility
type MetadataServer interface {
	SayHello(context.Context, *SayHelloRequest) (*SayHelloReply, error)
	mustEmbedUnimplementedMetadataServer()
}

// UnimplementedMetadataServer must be embedded to have forward compatible implementations.
type UnimplementedMetadataServer struct {
}

func (UnimplementedMetadataServer) SayHello(context.Context, *SayHelloRequest) (*SayHelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedMetadataServer) mustEmbedUnimplementedMetadataServer() {}

// UnsafeMetadataServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MetadataServer will
// result in compilation errors.
type UnsafeMetadataServer interface {
	mustEmbedUnimplementedMetadataServer()
}

func RegisterMetadataServer(s grpc.ServiceRegistrar, srv MetadataServer) {
	s.RegisterService(&Metadata_ServiceDesc, srv)
}

func _Metadata_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SayHelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metadata.Metadata/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServer).SayHello(ctx, req.(*SayHelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Metadata_ServiceDesc is the grpc.ServiceDesc for Metadata service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Metadata_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "metadata.Metadata",
	HandlerType: (*MetadataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Metadata_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/metadata/service.proto",
}