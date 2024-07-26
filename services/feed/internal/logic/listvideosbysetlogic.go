package logic

import (
	"context"
	"min-tiktok/services/feed/feed"
	"min-tiktok/services/feed/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListVideosBySetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListVideosBySetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListVideosBySetLogic {
	return &ListVideosBySetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ListVideosBySet query by set of video_id
func (l *ListVideosBySetLogic) ListVideosBySet(in *feed.ListVideosBySetRequest) (*feed.ListVideosBySetResponse, error) {
	videoList, err := l.svcCtx.VideoModel.ListVideoByVideoSet(l.ctx, in.VideoIdSet)
	if err != nil {
		l.Errorw("query video by video set failed", logx.Field("err", err))
		return nil, err
	}
	videos, err := FetchVideoDetails(l.ctx, videoList, in.ActorId, l.svcCtx.UserRpc, l.svcCtx.Rdb)
	if err != nil {
		l.Errorw("fetch video details failed", logx.Field("err", err))
		return nil, err
	}
	resp := &feed.ListVideosBySetResponse{
		VideoList: videos,
	}

	return resp, nil
}
