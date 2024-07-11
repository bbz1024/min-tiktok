// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"min-tiktok/api/favorite/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.WithMiddleware, serverCtx.AuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/action",
					Handler: FavoriteActionHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/list",
					Handler: FavoriteListHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/douyin/favorite"),
	)
}