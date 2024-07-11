package logic

import (
	"context"
	"min-tiktok/common/consts/code"
	"min-tiktok/services/user/user"

	"min-tiktok/api/user/internal/svc"
	"min-tiktok/api/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.GetUserInfoRequest) (resp *types.GetUserInfoResponse, err error) {
	// call rpc
	res, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &user.UserRequest{
		ActorId: req.ActorID, //
		UserId:  req.UserID,  //
	})
	resp = new(types.GetUserInfoResponse)
	if err != nil {
		resp.StatusCode = code.ServerError
		resp.StatusMsg = code.ServerErrorMsg
		l.Errorw("call rpc UserRpc.GetUserInfo error ", logx.Field("err", err))
		return
	}
	if res.StatusCode != code.OK {
		resp.StatusCode = uint32(res.StatusCode)
		resp.StatusMsg = res.StatusMsg
		return
	}
	resp.User = types.User{
		ID:              res.User.Id,
		Name:            res.User.Name,
		FollowCount:     res.User.FollowCount,
		FollowerCount:   res.User.FollowerCount,
		IsFollow:        res.User.IsFollow,
		Avatar:          res.User.Avatar,
		BackgroundImage: res.User.BackgroundImage,
		Signature:       res.User.Signature,
		TotalFavorited:  res.User.TotalFavorited,
		WorkCount:       res.User.WorkCount,
		FavoriteCount:   res.User.FavoriteCount,
	}
	return
}
