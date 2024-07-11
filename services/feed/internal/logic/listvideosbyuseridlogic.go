package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"min-tiktok/services/feed/feed"
	"min-tiktok/services/feed/internal/svc"
)

type ListVideosByUserIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListVideosByUserIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListVideosByUserIDLogic {
	return &ListVideosByUserIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ListVideosByUserID query by user_id 获取某个用户的视频列表
func (l *ListVideosByUserIDLogic) ListVideosByUserID(in *feed.ListVideosByUserIDRequest) (*feed.ListVideosByUserIDResponse, error) {
	videoList, err := l.svcCtx.VideoModel.ListVideoByUserId(l.ctx, int64(in.UserId))
	if err != nil {
		logx.Errorw("query video list by user_id error", logx.Field("err", err))
		return nil, err
	}

	videos, err := FetchVideoDetails(l.ctx, videoList, in.ActorId, l.svcCtx.UserRpc, l.svcCtx.Rdb)
	if err != nil {
		l.Errorw("fetch video details failed", logx.Field("err", err))
		return nil, err
	}
	resp := &feed.ListVideosByUserIDResponse{
		VideoList: videos,
	}

	return resp, nil
}
