// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: hello/hello.proto

package hello

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

const (
	D8Grpc_Hello_FullMethodName = "/hello.D8grpc/Hello"
)

// D8GrpcClient is the client API for D8Grpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type D8GrpcClient interface {
	// Sends a greeting
	Hello(ctx context.Context, in *MsgReq, opts ...grpc.CallOption) (*MsgRes, error)
}

type d8GrpcClient struct {
	cc grpc.ClientConnInterface
}

func NewD8GrpcClient(cc grpc.ClientConnInterface) D8GrpcClient {
	return &d8GrpcClient{cc}
}

func (c *d8GrpcClient) Hello(ctx context.Context, in *MsgReq, opts ...grpc.CallOption) (*MsgRes, error) {
	out := new(MsgRes)
	err := c.cc.Invoke(ctx, D8Grpc_Hello_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// D8GrpcServer is the server API for D8Grpc service.
// All implementations must embed UnimplementedD8GrpcServer
// for forward compatibility
type D8GrpcServer interface {
	// Sends a greeting
	Hello(context.Context, *MsgReq) (*MsgRes, error)
	mustEmbedUnimplementedD8GrpcServer()
}

// UnimplementedD8GrpcServer must be embedded to have forward compatible implementations.
type UnimplementedD8GrpcServer struct {
}

func (UnimplementedD8GrpcServer) Hello(context.Context, *MsgReq) (*MsgRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (UnimplementedD8GrpcServer) mustEmbedUnimplementedD8GrpcServer() {}

// UnsafeD8GrpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to D8GrpcServer will
// result in compilation errors.
type UnsafeD8GrpcServer interface {
	mustEmbedUnimplementedD8GrpcServer()
}

func RegisterD8GrpcServer(s grpc.ServiceRegistrar, srv D8GrpcServer) {
	s.RegisterService(&D8Grpc_ServiceDesc, srv)
}

func _D8Grpc_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(D8GrpcServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: D8Grpc_Hello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(D8GrpcServer).Hello(ctx, req.(*MsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

// D8Grpc_ServiceDesc is the grpc.ServiceDesc for D8Grpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var D8Grpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hello.D8grpc",
	HandlerType: (*D8GrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _D8Grpc_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello/hello.proto",
}
