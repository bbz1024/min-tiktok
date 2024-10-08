// Code generated by goctl. DO NOT EDIT.
// Source: auths.proto

package authsclient

import (
	"context"

	"min-tiktok/services/auths/auths"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AuthsRequest     = auths.AuthsRequest
	AuthsResponse    = auths.AuthsResponse
	LoginRequest     = auths.LoginRequest
	LoginResponse    = auths.LoginResponse
	RegisterRequest  = auths.RegisterRequest
	RegisterResponse = auths.RegisterResponse

	Auths interface {
		Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
		Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
		Authentication(ctx context.Context, in *AuthsRequest, opts ...grpc.CallOption) (*AuthsResponse, error)
	}

	defaultAuths struct {
		cli zrpc.Client
	}
)

func NewAuths(cli zrpc.Client) Auths {
	return &defaultAuths{
		cli: cli,
	}
}

func (m *defaultAuths) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	client := auths.NewAuthsClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

func (m *defaultAuths) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	client := auths.NewAuthsClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultAuths) Authentication(ctx context.Context, in *AuthsRequest, opts ...grpc.CallOption) (*AuthsResponse, error) {
	client := auths.NewAuthsClient(m.cli.Conn())
	return client.Authentication(ctx, in, opts...)
}
