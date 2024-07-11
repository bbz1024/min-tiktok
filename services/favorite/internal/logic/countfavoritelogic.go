package logic

import (
	"context"

	"min-tiktok/services/favorite/favorite"
	"min-tiktok/services/favorite/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CountFavoriteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCountFavoriteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CountFavoriteLogic {
	return &CountFavoriteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CountFavoriteLogic) CountFavorite(in *favorite.CountFavoriteRequest) (*favorite.CountFavoriteResponse, error) {
	// todo: add your logic here and delete this line

	return &favorite.CountFavoriteResponse{}, nil
}
