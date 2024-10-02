package logic

import (
	"context"
	"min-tiktok/api/relation/internal/svc"
	"min-tiktok/api/relation/internal/types"
	"min-tiktok/common/consts/code"
	"min-tiktok/services/relation/relation"
	"min-tiktok/services/user/userclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type RelationActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRelationActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationActionLogic {
	return &RelationActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RelationActionLogic) RelationAction(req *types.RealtionActionReuqest) (resp *types.RealtionActionResponse, err error) {
	// check user_id exist
	exist, err := l.svcCtx.UserRpc.CheckUserExist(l.ctx, &userclient.UserExistRequest{
		UserId: req.UserID,
	})
	resp = new(types.RealtionActionResponse)
	if err != nil {
		resp.StatusCode = code.ServerError
		resp.StatusMsg = code.ServerErrorMsg
		l.Errorw("call rpc UserRpc.CheckUserExist", logx.Field("err", err))
		return
	}
	if !exist.Exist {
		l.Infow("user not found", logx.Field("user_id", req.UserID))
		return &types.RealtionActionResponse{
			StatusCode: code.UserNotFoundCode,
			StatusMsg:  code.UserNotFoundMsg,
		}, nil
	}
	// check follow yourself
	if req.ActorID == req.UserID {
		resp = new(types.RealtionActionResponse)
		resp.StatusCode = code.ForbidFollowSelf
		resp.StatusMsg = code.ForbidFollowSelfMsg
		return
	}
	var actionReq = &relation.RelationActionRequest{
		ActorId: req.ActorID,
		UserId:  req.UserID,
	}

	var res *relation.RelationActionResponse
	switch relation.ActionType(req.ActionType) {
	case relation.ActionType_Follow:
		res, err = l.svcCtx.RelationRpc.Follow(l.ctx, actionReq)
	case relation.ActionType_UnFollow:
		res, err = l.svcCtx.RelationRpc.Unfollow(l.ctx, actionReq)
	}

	if err != nil {
		resp.StatusCode = code.ServerError
		resp.StatusMsg = code.ServerErrorMsg
		return
	}
	if res.StatusCode != code.OK {
		resp.StatusCode = res.StatusCode
		resp.StatusMsg = res.StatusMsg
		return
	}
	return &types.RealtionActionResponse{
		StatusCode: res.StatusCode,
		StatusMsg:  res.StatusMsg,
	}, nil
}
