package logic

import (
	"context"
	"fmt"
	"min-tiktok/common/consts/keys"
	"min-tiktok/services/relation/internal/svc"
	"min-tiktok/services/relation/relation"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFriendListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendListLogic {
	return &GetFriendListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetFriendList follow each other
func (l *GetFriendListLogic) GetFriendList(in *relation.FriendListRequest) (*relation.FriendListResponse, error) {

	// user follow
	userFollowKey := fmt.Sprintf(keys.UserFollow, in.UserId)
	// user follower
	userFollowerKey := fmt.Sprintf(keys.UserFollower, in.UserId)
	friendKey := fmt.Sprintf(keys.UserFriendKey, in.UserId)
	//intersection
	if _, err := l.svcCtx.Rdb.SinterstoreCtx(l.ctx, friendKey, userFollowKey, userFollowerKey); err != nil {
		return nil, err
	}
	userList, err := fetchUserList(l.ctx, friendKey, in.UserId, l.svcCtx.Rdb, l.svcCtx.UserRpc)
	if err != nil {
		return nil, err
	}
	return &relation.FriendListResponse{
		UserList: userList,
	}, nil
}
