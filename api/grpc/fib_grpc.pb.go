// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package apiGRPC

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

// FibClient is the client API for Fib service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FibClient interface {
	GetFib(ctx context.Context, in *FibRequest, opts ...grpc.CallOption) (*FibResponse, error)
}

type fibClient struct {
	cc grpc.ClientConnInterface
}

func NewFibClient(cc grpc.ClientConnInterface) FibClient {
	return &fibClient{cc}
}

func (c *fibClient) GetFib(ctx context.Context, in *FibRequest, opts ...grpc.CallOption) (*FibResponse, error) {
	out := new(FibResponse)
	err := c.cc.Invoke(ctx, "/apiGRPC.Fib/GetFib", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FibServer is the server API for Fib service.
// All implementations must embed UnimplementedFibServer
// for forward compatibility
type FibServer interface {
	GetFib(context.Context, *FibRequest) (*FibResponse, error)
	mustEmbedUnimplementedFibServer()
}

// UnimplementedFibServer must be embedded to have forward compatible implementations.
type UnimplementedFibServer struct {
}

func (UnimplementedFibServer) GetFib(context.Context, *FibRequest) (*FibResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFib not implemented")
}
func (UnimplementedFibServer) mustEmbedUnimplementedFibServer() {}

// UnsafeFibServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FibServer will
// result in compilation errors.
type UnsafeFibServer interface {
	mustEmbedUnimplementedFibServer()
}

func RegisterFibServer(s grpc.ServiceRegistrar, srv FibServer) {
	s.RegisterService(&Fib_ServiceDesc, srv)
}

func _Fib_GetFib_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FibRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FibServer).GetFib(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apiGRPC.Fib/GetFib",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FibServer).GetFib(ctx, req.(*FibRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Fib_ServiceDesc is the grpc.ServiceDesc for Fib service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Fib_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "apiGRPC.Fib",
	HandlerType: (*FibServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFib",
			Handler:    _Fib_GetFib_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fib.proto",
}
