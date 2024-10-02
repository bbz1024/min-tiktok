package logic

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"min-tiktok/common/consts/keys"
	"min-tiktok/services/relation/internal/svc"
	"min-tiktok/services/relation/relation"
)

type GetFollowListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFollowListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowListLogic {
	return &GetFollowListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx).WithFields(logx.Field("type", "service")),
	}
}

func (l *GetFollowListLogic) GetFollowList(in *relation.FollowListRequest) (*relation.FollowListResponse, error) {

	userFollowKey := fmt.Sprintf(keys.UserFollow, in.UserId)
	resp := new(relation.FollowListResponse)
	userList, err := fetchUserList(l.ctx, userFollowKey, in.ActorId, l.svcCtx.Rdb, l.svcCtx.UserRpc)
	if err != nil {
		l.Errorw("fetch user list failed", logx.Field("err", err))
		return nil, err
	}
	resp.UserList = userList
	return resp, nil
}
