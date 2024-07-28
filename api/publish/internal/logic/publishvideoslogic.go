package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"min-tiktok/api/publish/internal/svc"
	"min-tiktok/api/publish/internal/types"
	"min-tiktok/common/consts/code"
	"min-tiktok/services/publish/publish"
)

type PublishVideosLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishVideosLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishVideosLogic {
	return &PublishVideosLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishVideosLogic) PublishVideos(req *types.PublishActionReq) (resp *types.PublishActionResp, err error) {

	res, err := l.svcCtx.PublishRpc.ActionVideo(l.ctx, &publish.ActionVideoReq{
		ActorId: req.ActorId,
		Data:    req.Data,
		Title:   req.Title,
	})
	if err != nil {
		resp.StatusMsg = code.ServerErrorMsg
		resp.StatusCode = code.ServerError
		l.Errorw("call rpc PublishRpc.PublishVideos error ", logx.Field("err", err))
		return
	}
	if res.StatusCode != code.OK {
		resp.StatusMsg = res.StatusMsg
		resp.StatusCode = res.StatusCode
		return
	}
	resp = new(types.PublishActionResp)
	resp.StatusMsg = res.StatusMsg
	resp.StatusCode = res.StatusCode
	return
}
