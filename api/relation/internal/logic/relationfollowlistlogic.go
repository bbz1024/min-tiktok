package logic

import (
	"context"
	"min-tiktok/common/consts/code"
	"min-tiktok/services/relation/relationclient"

	"min-tiktok/api/relation/internal/svc"
	"min-tiktok/api/relation/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RelationFollowListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRelationFollowListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationFollowListLogic {
	return &RelationFollowListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

/*
	res, err := l.svcCtx.RelationRpc.GetFollowerList(l.ctx, &relationclient.FollowerListRequest{
		ActorId: req.ActorID,
		UserId:  req.UserID,
	})

	if err != nil {
		resp.StatusMsg = code.ServerErrorMsg
		resp.StatusCode = code.ServerError
		return
	}

resp = new(types.RelationFollowerListResponse)

	if res.StatusCode != 0 {
		resp.StatusMsg = res.StatusMsg
		resp.StatusCode = uint32(res.StatusCode)
		return resp, nil
	}

var followerList = make([]*types.User, len(res.UserList))

	for i, user := range res.UserList {
		followerList[i] = &types.User{
			ID:              user.Id,
			Name:            user.Name,
			FollowCount:     user.FollowCount,
			FollowerCount:   user.FollowerCount,
			IsFollow:        user.IsFollow,
			Avatar:          user.Avatar,
			BackgroundImage: user.BackgroundImage,
			Signature:       user.Signature,
			TotalFavorited:  user.TotalFavorited,
			WorkCount:       user.WorkCount,
			FavoriteCount:   user.FavoriteCount,
		}
	}

resp.UserList = followerList
return resp, nil
*/
func (l *RelationFollowListLogic) RelationFollowList(req *types.RelationFollowListRequest) (resp *types.RelationFollowListResponse, err error) {
	res, err := l.svcCtx.RelationRpc.GetFollowList(
		l.ctx,
		&relationclient.FollowListRequest{
			ActorId: req.ActorID,
			UserId:  req.UserID,
		})
	resp = new(types.RelationFollowListResponse)
	if err != nil {
		resp.StatusCode = code.ServerError
		resp.StatusMsg = code.ServerErrorMsg
		return
	}
	resp = new(types.RelationFollowListResponse)
	if res.StatusCode != code.OK {
		resp.StatusCode = uint32(res.StatusCode)
		resp.StatusMsg = res.StatusMsg
		return
	}
	resp.UserList = make([]*types.User, len(res.UserList))
	for i, user := range res.UserList {
		resp.UserList[i] = &types.User{
			ID:              user.Id,
			Name:            user.Name,
			FollowCount:     user.FollowCount,
			FollowerCount:   user.FollowerCount,
			IsFollow:        user.IsFollow,
			Avatar:          user.Avatar,
			BackgroundImage: user.BackgroundImage,
			Signature:       user.Signature,
			TotalFavorited:  user.TotalFavorited,
			WorkCount:       user.WorkCount,
			FavoriteCount:   user.FavoriteCount,
		}
	}
	return
}
