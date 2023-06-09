// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: server.proto

package pb

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

// FileCheckClient is the client API for FileCheck service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileCheckClient interface {
	Execute(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Resp, error)
}

type fileCheckClient struct {
	cc grpc.ClientConnInterface
}

func NewFileCheckClient(cc grpc.ClientConnInterface) FileCheckClient {
	return &fileCheckClient{cc}
}

func (c *fileCheckClient) Execute(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Resp, error) {
	out := new(Resp)
	err := c.cc.Invoke(ctx, "/demo.File_check/Execute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileCheckServer is the server API for FileCheck service.
// All implementations must embed UnimplementedFileCheckServer
// for forward compatibility
type FileCheckServer interface {
	Execute(context.Context, *Req) (*Resp, error)
	mustEmbedUnimplementedFileCheckServer()
}

// UnimplementedFileCheckServer must be embedded to have forward compatible implementations.
type UnimplementedFileCheckServer struct {
}

func (UnimplementedFileCheckServer) Execute(context.Context, *Req) (*Resp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Execute not implemented")
}
func (UnimplementedFileCheckServer) mustEmbedUnimplementedFileCheckServer() {}

// UnsafeFileCheckServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileCheckServer will
// result in compilation errors.
type UnsafeFileCheckServer interface {
	mustEmbedUnimplementedFileCheckServer()
}

func RegisterFileCheckServer(s grpc.ServiceRegistrar, srv FileCheckServer) {
	s.RegisterService(&FileCheck_ServiceDesc, srv)
}

func _FileCheck_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileCheckServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo.File_check/Execute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileCheckServer).Execute(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

// FileCheck_ServiceDesc is the grpc.ServiceDesc for FileCheck service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileCheck_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "demo.File_check",
	HandlerType: (*FileCheckServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Execute",
			Handler:    _FileCheck_Execute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server.proto",
}
