package logic

import (
	"context"
	"min-tiktok/common/consts/code"
	"min-tiktok/models/video"
	"min-tiktok/services/feed/feed"
	"min-tiktok/services/feed/internal/svc"
	"min-tiktok/services/feedback/feedbackclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListRecommendVideosLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListRecommendVideosLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListRecommendVideosLogic {
	return &ListRecommendVideosLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListRecommendVideosLogic) ListRecommendVideos(in *feed.ListRecommendRequest) (*feed.ListFeedResponse, error) {
	res, err := l.svcCtx.FeedBackRpc.ListRecommendWithFeedback(l.ctx, &feedbackclient.ListRecommendRequest{
		ActorId: in.ActorId,
		Count:   video.Count,
	})
	if err != nil {
		logx.Errorf("get recommend error: %v", err)
		return nil, err
	}
	if res.StatusCode != code.OK {
		return nil, err
	}
	videoSet, err := l.svcCtx.VideoModel.ListVideoByVideoSet(l.ctx, res.GetVideoIds())
	if err != nil {
		logx.Errorf("list video by video set error: %v", err)
		return nil, err
	}

	videoList, err := FetchVideoDetails(l.ctx, videoSet, in.ActorId, l.svcCtx.UserRpc, l.svcCtx.Rdb)
	if err != nil {
		logx.Errorf("fetch video details error: %v", err)
		return nil, err
	}
	return &feed.ListFeedResponse{
		VideoList: videoList,
	}, nil
}
