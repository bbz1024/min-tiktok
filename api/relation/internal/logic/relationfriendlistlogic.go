package logic

import (
	"context"
	"min-tiktok/common/consts/code"
	"min-tiktok/services/relation/relationclient"
	"min-tiktok/services/user/userclient"

	"min-tiktok/api/relation/internal/svc"
	"min-tiktok/api/relation/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RelationFriendListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRelationFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationFriendListLogic {
	return &RelationFriendListLogic{
		Logger: logx.WithContext(ctx).WithFields(logx.Field("type", "api")),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RelationFriendListLogic) RelationFriendList(req *types.RelationFriendListRequest) (resp *types.RelationFriendListResponse, err error) {
	// check user_id exist
	exist, err := l.svcCtx.UserRpc.CheckUserExist(l.ctx, &userclient.UserExistRequest{
		UserId: req.UserID,
	})
	resp = new(types.RelationFriendListResponse)
	if err != nil {
		resp.StatusCode = code.ServerError
		resp.StatusMsg = code.ServerErrorMsg
		l.Errorw("call rpc UserRpc.CheckUserExist", logx.Field("err", err))
		return
	}
	if !exist.Exist {
		l.Infow("user not found", logx.Field("user_id", req.UserID))
		return &types.RelationFriendListResponse{
			StatusCode: code.UserNotFoundCode,
			StatusMsg:  code.UserNotFoundMsg,
		}, nil
	}
	res, err := l.svcCtx.RelationRpc.GetFriendList(l.ctx, &relationclient.FriendListRequest{
		UserId:  req.UserID,
		ActorId: req.ActorID,
	})
	resp = new(types.RelationFriendListResponse)
	if err != nil {
		resp.StatusCode = code.ServerError
		resp.StatusMsg = code.ServerErrorMsg
		l.Errorw("call rpc RelationRpc.GetFriendList", logx.Field("err", err))
		return
	}
	if res.StatusCode != code.OK {
		resp.StatusCode = res.StatusCode
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
