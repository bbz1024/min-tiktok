package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/mr"
	"min-tiktok/common/consts/code"
	"min-tiktok/services/message/messageclient"
	"min-tiktok/services/user/userclient"

	"min-tiktok/api/message/internal/svc"
	"min-tiktok/api/message/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListMessageChatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListMessageChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListMessageChatLogic {
	return &ListMessageChatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListMessageChatLogic) ListMessageChat(req *types.MessageChatReq) (resp *types.MessageChatResp, err error) {
	var exist bool
	err = mr.Finish(func() error {
		res, err := l.svcCtx.UserRpc.CheckUserExist(l.ctx, &userclient.UserExistRequest{
			UserId: req.ActorId,
		})
		if err != nil {
			return err
		}
		exist = res.Exist
		return nil
	}, func() error {
		res, err := l.svcCtx.UserRpc.CheckUserExist(l.ctx, &userclient.UserExistRequest{
			UserId: req.ToUserID,
		})
		if err != nil {
			return err
		}
		exist = res.Exist
		return nil
	})
	if err != nil {
		l.Errorw("call rpc UserRpc.CheckUserExist failed", logx.Field("err", err), logx.Field("user_id", req.ToUserID))
		return nil, err
	}
	if !exist {
		return &types.MessageChatResp{
			StatusCode: code.UserNotFoundCode,
			StatusMsg:  code.UserNotFoundMsg,
		}, nil
	}
	res, err := l.svcCtx.MessageRpc.MessageList(l.ctx, &messageclient.MessageListRequest{
		UserId:     req.ToUserID,
		ActorId:    req.ActorId,
		PreMsgTime: req.PreMsgTime,
	})
	if err != nil {
		l.Errorw("call rpc MessageRpc.MessageList failed")
		return nil, err
	}
	resp = new(types.MessageChatResp)
	resp.MessageList = make([]*types.Message, 0, len(res.MessageList))
	if res.StatusCode != code.OK {
		resp.StatusCode = res.StatusCode
		resp.StatusMsg = res.StatusMsg
		return resp, nil
	}
	for _, v := range res.MessageList {
		resp.MessageList = append(resp.MessageList, &types.Message{
			ID:         v.Id,
			Content:    v.Content,
			CreateTime: v.CreateTime,
		})
	}
	return
}
