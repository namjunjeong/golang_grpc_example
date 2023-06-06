// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.1
// source: unary.proto

package unary

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

// UnaryClient is the client API for Unary service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UnaryClient interface {
	Multiply(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Res, error)
}

type unaryClient struct {
	cc grpc.ClientConnInterface
}

func NewUnaryClient(cc grpc.ClientConnInterface) UnaryClient {
	return &unaryClient{cc}
}

func (c *unaryClient) Multiply(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Res, error) {
	out := new(Res)
	err := c.cc.Invoke(ctx, "/Unary/Multiply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UnaryServer is the server API for Unary service.
// All implementations must embed UnimplementedUnaryServer
// for forward compatibility
type UnaryServer interface {
	Multiply(context.Context, *Req) (*Res, error)
	mustEmbedUnimplementedUnaryServer()
}

// UnimplementedUnaryServer must be embedded to have forward compatible implementations.
type UnimplementedUnaryServer struct {
}

func (UnimplementedUnaryServer) Multiply(context.Context, *Req) (*Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Multiply not implemented")
}
func (UnimplementedUnaryServer) mustEmbedUnimplementedUnaryServer() {}

// UnsafeUnaryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UnaryServer will
// result in compilation errors.
type UnsafeUnaryServer interface {
	mustEmbedUnimplementedUnaryServer()
}

func RegisterUnaryServer(s grpc.ServiceRegistrar, srv UnaryServer) {
	s.RegisterService(&Unary_ServiceDesc, srv)
}

func _Unary_Multiply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UnaryServer).Multiply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Unary/Multiply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UnaryServer).Multiply(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

// Unary_ServiceDesc is the grpc.ServiceDesc for Unary service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Unary_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Unary",
	HandlerType: (*UnaryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Multiply",
			Handler:    _Unary_Multiply_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "unary.proto",
}