// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"min-tiktok/api/user/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.WithMiddleware, serverCtx.AuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/",
					Handler: GetUserInfoHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/douyin/user"),
	)
}
