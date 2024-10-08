// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.19.4
// source: favorite.proto

package favorite

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
	Favorite_FavoriteAction_FullMethodName = "/favorite.Favorite/FavoriteAction"
	Favorite_FavoriteList_FullMethodName   = "/favorite.Favorite/FavoriteList"
)

// FavoriteClient is the client API for Favorite service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FavoriteClient interface {
	FavoriteAction(ctx context.Context, in *FavoriteRequest, opts ...grpc.CallOption) (*FavoriteResponse, error)
	FavoriteList(ctx context.Context, in *FavoriteListRequest, opts ...grpc.CallOption) (*FavoriteListResponse, error)
}

type favoriteClient struct {
	cc grpc.ClientConnInterface
}

func NewFavoriteClient(cc grpc.ClientConnInterface) FavoriteClient {
	return &favoriteClient{cc}
}

func (c *favoriteClient) FavoriteAction(ctx context.Context, in *FavoriteRequest, opts ...grpc.CallOption) (*FavoriteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FavoriteResponse)
	err := c.cc.Invoke(ctx, Favorite_FavoriteAction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoriteClient) FavoriteList(ctx context.Context, in *FavoriteListRequest, opts ...grpc.CallOption) (*FavoriteListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FavoriteListResponse)
	err := c.cc.Invoke(ctx, Favorite_FavoriteList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FavoriteServer is the server API for Favorite service.
// All implementations must embed UnimplementedFavoriteServer
// for forward compatibility
type FavoriteServer interface {
	FavoriteAction(context.Context, *FavoriteRequest) (*FavoriteResponse, error)
	FavoriteList(context.Context, *FavoriteListRequest) (*FavoriteListResponse, error)
	mustEmbedUnimplementedFavoriteServer()
}

// UnimplementedFavoriteServer must be embedded to have forward compatible implementations.
type UnimplementedFavoriteServer struct {
}

func (UnimplementedFavoriteServer) FavoriteAction(context.Context, *FavoriteRequest) (*FavoriteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FavoriteAction not implemented")
}
func (UnimplementedFavoriteServer) FavoriteList(context.Context, *FavoriteListRequest) (*FavoriteListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FavoriteList not implemented")
}
func (UnimplementedFavoriteServer) mustEmbedUnimplementedFavoriteServer() {}

// UnsafeFavoriteServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FavoriteServer will
// result in compilation errors.
type UnsafeFavoriteServer interface {
	mustEmbedUnimplementedFavoriteServer()
}

func RegisterFavoriteServer(s grpc.ServiceRegistrar, srv FavoriteServer) {
	s.RegisterService(&Favorite_ServiceDesc, srv)
}

func _Favorite_FavoriteAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavoriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavoriteServer).FavoriteAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Favorite_FavoriteAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavoriteServer).FavoriteAction(ctx, req.(*FavoriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Favorite_FavoriteList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavoriteListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavoriteServer).FavoriteList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Favorite_FavoriteList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavoriteServer).FavoriteList(ctx, req.(*FavoriteListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Favorite_ServiceDesc is the grpc.ServiceDesc for Favorite service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Favorite_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "favorite.Favorite",
	HandlerType: (*FavoriteServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FavoriteAction",
			Handler:    _Favorite_FavoriteAction_Handler,
		},
		{
			MethodName: "FavoriteList",
			Handler:    _Favorite_FavoriteList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "favorite.proto",
}
