package logic

import (
	"context"

	"min-tiktok/services/favorite/favorite"
	"min-tiktok/services/favorite/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CountUserTotalFavoritedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCountUserTotalFavoritedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CountUserTotalFavoritedLogic {
	return &CountUserTotalFavoritedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CountUserTotalFavoritedLogic) CountUserTotalFavorited(in *favorite.CountUserTotalFavoritedRequest) (*favorite.CountUserTotalFavoritedResponse, error) {
	// todo: add your logic here and delete this line

	return &favorite.CountUserTotalFavoritedResponse{}, nil
}
