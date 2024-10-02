package logic

import (
	"context"
	"fmt"
	"min-tiktok/common/consts/keys"
	"min-tiktok/services/relation/internal/svc"
	"min-tiktok/services/relation/relation"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsFriendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsFriendLogic {
	return &IsFriendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx).WithFields(logx.Field("type", "service")),
	}
}

func (l *IsFriendLogic) IsFriend(in *relation.IsFriendRequest) (*relation.IsFriendResponse, error) {
	actorKey := fmt.Sprintf(keys.UserFollow, in.ActorId)
	exist, err := l.svcCtx.Rdb.SismemberCtx(l.ctx, actorKey, in.UserId)
	if err != nil {
		l.Errorf("IsFriend: SismemberCtx error: %v", err)
		return nil, err
	}
	userKey := fmt.Sprintf(keys.UserFollow, in.UserId)
	exist2, err := l.svcCtx.Rdb.SismemberCtx(l.ctx, userKey, in.ActorId)
	if err != nil {
		l.Errorf("IsFriend: SismemberCtx error: %v", err)
		return nil, err
	}
	return &relation.IsFriendResponse{
		Result: exist && exist2,
	}, nil
}
