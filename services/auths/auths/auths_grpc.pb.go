// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.19.4
// source: auths.proto

package auths

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Auths_Register_FullMethodName       = "/auths.Auths/Register"
	Auths_Login_FullMethodName          = "/auths.Auths/Login"
	Auths_Authentication_FullMethodName = "/auths.Auths/Authentication"
)

// AuthsClient is the client API for Auths service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthsClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Authentication(ctx context.Context, in *AuthsRequest, opts ...grpc.CallOption) (*AuthsResponse, error)
}

type authsClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthsClient(cc grpc.ClientConnInterface) AuthsClient {
	return &authsClient{cc}
}

func (c *authsClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, Auths_Register_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authsClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, Auths_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authsClient) Authentication(ctx context.Context, in *AuthsRequest, opts ...grpc.CallOption) (*AuthsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AuthsResponse)
	err := c.cc.Invoke(ctx, Auths_Authentication_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthsServer is the server API for Auths service.
// All implementations must embed UnimplementedAuthsServer
// for forward compatibility
type AuthsServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Authentication(context.Context, *AuthsRequest) (*AuthsResponse, error)
	mustEmbedUnimplementedAuthsServer()
}

// UnimplementedAuthsServer must be embedded to have forward compatible implementations.
type UnimplementedAuthsServer struct {
}

func (UnimplementedAuthsServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedAuthsServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthsServer) Authentication(context.Context, *AuthsRequest) (*AuthsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authentication not implemented")
}
func (UnimplementedAuthsServer) mustEmbedUnimplementedAuthsServer() {}

// UnsafeAuthsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthsServer will
// result in compilation errors.
type UnsafeAuthsServer interface {
	mustEmbedUnimplementedAuthsServer()
}

func RegisterAuthsServer(s grpc.ServiceRegistrar, srv AuthsServer) {
	s.RegisterService(&Auths_ServiceDesc, srv)
}

func _Auths_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthsServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auths_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthsServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auths_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthsServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auths_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthsServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auths_Authentication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthsServer).Authentication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auths_Authentication_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthsServer).Authentication(ctx, req.(*AuthsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Auths_ServiceDesc is the grpc.ServiceDesc for Auths service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Auths_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auths.Auths",
	HandlerType: (*AuthsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Auths_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Auths_Login_Handler,
		},
		{
			MethodName: "Authentication",
			Handler:    _Auths_Authentication_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auths.proto",
}
