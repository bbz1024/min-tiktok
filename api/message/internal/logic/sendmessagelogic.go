package logic

import (
	"context"
	"min-tiktok/api/message/internal/svc"
	"min-tiktok/api/message/internal/types"
	"min-tiktok/common/consts/code"
	"min-tiktok/services/message/messageclient"
	"min-tiktok/services/user/userclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendMessageLogic {
	return &SendMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendMessageLogic) SendMessage(req *types.MessageActionReq) (resp *types.MessageActionResp, err error) {
	// check user_id exist
	exist1, err := l.svcCtx.UserRpc.CheckUserExist(l.ctx, &userclient.UserExistRequest{
		UserId: req.ActorId,
	})
	if err != nil {
		l.Errorw("call rpc UserRpc.CheckUserExist failed", logx.Field("err", err), logx.Field("user_id", req.ActorId))
		return nil, err
	}
	exist2, err := l.svcCtx.UserRpc.CheckUserExist(l.ctx, &userclient.UserExistRequest{
		UserId: req.ToUserID,
	})
	if err != nil {
		l.Errorw("call rpc UserRpc.CheckUserExist failed", logx.Field("err", err), logx.Field("user_id", req.ToUserID))
		return nil, err
	}
	if !exist1.Exist || !exist2.Exist {
		return &types.MessageActionResp{
			StatusCode: code.UserNotFoundCode,
			StatusMsg:  code.UserNotFoundMsg,
		}, nil
	}

	res, err := l.svcCtx.MessageRpc.MessageAction(l.ctx, &messageclient.MessageActionRequest{
		UserId:     req.ToUserID,
		ActorId:    req.ActorId,
		ActionType: req.ActionType,
		Content:    req.Content,
	})
	if err != nil {
		l.Errorw("call rpc MessageRpc.MessageAction failed", logx.Field("err", err))
		return nil, err
	}
	resp = new(types.MessageActionResp)
	if res.StatusCode != code.OK {
		resp.StatusCode = res.StatusCode
		resp.StatusMsg = res.StatusMsg
		return resp, nil
	}
	return
}
