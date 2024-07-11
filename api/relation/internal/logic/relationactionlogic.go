package logic

import (
	"context"
	"min-tiktok/api/relation/internal/svc"
	"min-tiktok/api/relation/internal/types"
	"min-tiktok/common/consts/code"
	"min-tiktok/services/relation/relation"

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

	resp = new(types.RealtionActionResponse)
	if err != nil {
		resp.StatusCode = code.ServerError
		resp.StatusMsg = code.ServerErrorMsg
		return
	}
	if res.StatusCode != code.OK {
		resp.StatusCode = uint32(res.StatusCode)
		resp.StatusMsg = res.StatusMsg
		return
	}
	return &types.RealtionActionResponse{
		StatusCode: uint32(res.StatusCode),
		StatusMsg:  res.StatusMsg,
	}, nil
}
