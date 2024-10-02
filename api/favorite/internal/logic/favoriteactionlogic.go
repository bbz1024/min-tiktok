package logic

import (
	"context"
	"min-tiktok/common/consts/code"
	"min-tiktok/services/favorite/favorite"

	"min-tiktok/api/favorite/internal/svc"
	"min-tiktok/api/favorite/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteActionLogic {
	return &FavoriteActionLogic{
		Logger: logx.WithContext(ctx).WithFields(logx.Field("type", "api")),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteActionLogic) FavoriteAction(req *types.ActionRequest) (resp *types.ActionResponse, err error) {
	res, err := l.svcCtx.FavoriteRpc.FavoriteAction(l.ctx, &favorite.FavoriteRequest{
		ActionType: favorite.ActionType(req.ActionType),
		ActorId:    req.ActorID, // 点赞的人
		VideoId:    uint32(req.VideoID),
	})
	resp = new(types.ActionResponse)
	if err != nil {
		resp.StatusCode = code.ServerError
		resp.StatusMsg = code.ServerErrorMsg
		l.Errorw("call rpc FavoriteRpc.FavoriteAction error ", logx.Field("err", err))
		return
	}
	if res.StatusCode != code.OK {
		resp.StatusCode = uint32(res.StatusCode)
		resp.StatusMsg = res.StatusMsg
		return
	}
	resp.StatusCode = uint32(res.StatusCode)
	resp.StatusMsg = res.StatusMsg
	return
}
