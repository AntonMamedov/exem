// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: protos/arithmetical.proto

package arithmetical

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

// ArithmeticalClient is the client API for Arithmetical service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArithmeticalClient interface {
	Mul(ctx context.Context, in *MulRequest, opts ...grpc.CallOption) (*MulResponse, error)
	MulOn2(ctx context.Context, in *MulOn2Request, opts ...grpc.CallOption) (*CodeResponse, error)
	Code(ctx context.Context, in *CodeRequest, opts ...grpc.CallOption) (*CodeResponse, error)
}

type arithmeticalClient struct {
	cc grpc.ClientConnInterface
}

func NewArithmeticalClient(cc grpc.ClientConnInterface) ArithmeticalClient {
	return &arithmeticalClient{cc}
}

func (c *arithmeticalClient) Mul(ctx context.Context, in *MulRequest, opts ...grpc.CallOption) (*MulResponse, error) {
	out := new(MulResponse)
	err := c.cc.Invoke(ctx, "/Arithmetical/Mul", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *arithmeticalClient) MulOn2(ctx context.Context, in *MulOn2Request, opts ...grpc.CallOption) (*CodeResponse, error) {
	out := new(CodeResponse)
	err := c.cc.Invoke(ctx, "/Arithmetical/MulOn2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *arithmeticalClient) Code(ctx context.Context, in *CodeRequest, opts ...grpc.CallOption) (*CodeResponse, error) {
	out := new(CodeResponse)
	err := c.cc.Invoke(ctx, "/Arithmetical/Code", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArithmeticalServer is the server API for Arithmetical service.
// All implementations must embed UnimplementedArithmeticalServer
// for forward compatibility
type ArithmeticalServer interface {
	Mul(context.Context, *MulRequest) (*MulResponse, error)
	MulOn2(context.Context, *MulOn2Request) (*CodeResponse, error)
	Code(context.Context, *CodeRequest) (*CodeResponse, error)
	mustEmbedUnimplementedArithmeticalServer()
}

// UnimplementedArithmeticalServer must be embedded to have forward compatible implementations.
type UnimplementedArithmeticalServer struct {
}

func (UnimplementedArithmeticalServer) Mul(context.Context, *MulRequest) (*MulResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Mul not implemented")
}
func (UnimplementedArithmeticalServer) MulOn2(context.Context, *MulOn2Request) (*CodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MulOn2 not implemented")
}
func (UnimplementedArithmeticalServer) Code(context.Context, *CodeRequest) (*CodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Code not implemented")
}
func (UnimplementedArithmeticalServer) mustEmbedUnimplementedArithmeticalServer() {}

// UnsafeArithmeticalServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArithmeticalServer will
// result in compilation errors.
type UnsafeArithmeticalServer interface {
	mustEmbedUnimplementedArithmeticalServer()
}

func RegisterArithmeticalServer(s grpc.ServiceRegistrar, srv ArithmeticalServer) {
	s.RegisterService(&Arithmetical_ServiceDesc, srv)
}

func _Arithmetical_Mul_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MulRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArithmeticalServer).Mul(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Arithmetical/Mul",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArithmeticalServer).Mul(ctx, req.(*MulRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Arithmetical_MulOn2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MulOn2Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArithmeticalServer).MulOn2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Arithmetical/MulOn2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArithmeticalServer).MulOn2(ctx, req.(*MulOn2Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Arithmetical_Code_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArithmeticalServer).Code(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Arithmetical/Code",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArithmeticalServer).Code(ctx, req.(*CodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Arithmetical_ServiceDesc is the grpc.ServiceDesc for Arithmetical service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Arithmetical_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Arithmetical",
	HandlerType: (*ArithmeticalServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Mul",
			Handler:    _Arithmetical_Mul_Handler,
		},
		{
			MethodName: "MulOn2",
			Handler:    _Arithmetical_MulOn2_Handler,
		},
		{
			MethodName: "Code",
			Handler:    _Arithmetical_Code_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/arithmetical.proto",
}
