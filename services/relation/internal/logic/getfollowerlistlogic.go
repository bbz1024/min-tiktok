package logic

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"min-tiktok/common/consts/keys"
	"min-tiktok/services/relation/internal/svc"
	"min-tiktok/services/relation/relation"
)

type GetFollowerListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFollowerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowerListLogic {
	return &GetFollowerListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx).WithFields(logx.Field("type", "service")),
	}
}

// GetFollowerList get user follower list 粉丝
func (l *GetFollowerListLogic) GetFollowerList(in *relation.FollowerListRequest) (*relation.FollowerListResponse, error) {

	userFollowerKey := fmt.Sprintf(keys.UserFollower, in.UserId)
	resp := new(relation.FollowerListResponse)
	userList, err := fetchUserList(l.ctx, userFollowerKey, in.ActorId, l.svcCtx.Rdb, l.svcCtx.UserRpc)
	if err != nil {
		l.Errorw("fetchUserList failed", logx.Field("err", err))
		return nil, err
	}
	resp.UserList = userList
	return resp, nil
}
