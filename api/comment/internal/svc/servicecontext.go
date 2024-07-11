package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"min-tiktok/api/comment/internal/config"
	"min-tiktok/common/middleware"
	"min-tiktok/services/comment/commentclient"
)

type ServiceContext struct {
	Config         config.Config
	WithMiddleware rest.Middleware
	AuthMiddleware rest.Middleware
	CommentRpc     commentclient.Comment
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		AuthMiddleware: middleware.WrapperAuthMiddleware(c.AuthsRpc),
		WithMiddleware: middleware.WithMiddleware,
		//		FavoriteRpc:    favoriteclient.NewFavorite(zrpc.MustNewClient(c.FavoriteRpc)),
		CommentRpc: commentclient.NewComment(zrpc.MustNewClient(c.CommentRpc)),
	}
}
