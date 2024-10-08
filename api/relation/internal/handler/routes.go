// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"min-tiktok/api/relation/internal/svc"

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
					Handler: RelationActionHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/follow/list",
					Handler: RelationFollowListHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/follower/list",
					Handler: RelationFollowerListHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/friend/list",
					Handler: RelationFriendListHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/douyin/relation"),
	)
}
