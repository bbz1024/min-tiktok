package logic

import (
	"context"
	"min-tiktok/common/consts/variable"
	recommend "min-tiktok/services/feedback/internal/mq"
	"strconv"

	"min-tiktok/services/feedback/feedback"
	"min-tiktok/services/feedback/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListRecommendWithFeedbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListRecommendWithFeedbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListRecommendWithFeedbackLogic {
	return &ListRecommendWithFeedbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListRecommendWithFeedbackLogic) ListRecommendWithFeedback(in *feedback.ListRecommendRequest) (*feedback.ListRecommendResponse, error) {
	recommendSet, err := l.svcCtx.Recommend.GetRecommend(l.ctx, strconv.Itoa(int(in.ActorId)), "", int(in.Count))
	if err != nil {
		l.Errorw("feedback error", logx.Field("err", err))
		return nil, err
	}
	if err := recommend.GetInstance().Product(recommend.GorseRecommendReq{
		Type:     variable.ReadFeedBack,
		UserID:   in.ActorId,
		VideoIds: nil,
	}); err != nil {
		l.Errorw("feedback error", logx.Field("err", err))
		return nil, err
	}

	return &feedback.ListRecommendResponse{
		VideoIds: recommendSet,
	}, nil
}
