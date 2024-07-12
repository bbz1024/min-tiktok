package logic

import (
	"context"
	"min-tiktok/common/consts/variable"
	recommend "min-tiktok/services/feedback/internal/mq"

	"min-tiktok/services/feedback/feedback"
	"min-tiktok/services/feedback/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeedbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFeedbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedbackLogic {
	return &FeedbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FeedbackLogic) Feedback(in *feedback.FeedbackRequest) (*feedback.FeedbackResponse, error) {
	if err := recommend.GetInstance().Product(recommend.GorseRecommendReq{
		Type:     variable.FeedType(in.Type),
		UserID:   in.UserId,
		VideoIds: in.VideoIds,
	}); err != nil {
		l.Errorw("feedback error", logx.Field("err", err))
		return nil, err
	}
	return &feedback.FeedbackResponse{}, nil
}
