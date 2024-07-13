package logic

import (
	"context"
	"fmt"
	"min-tiktok/common/consts/code"
	"min-tiktok/services/relation/relation"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	message2 "min-tiktok/models/message"
	"min-tiktok/services/message/internal/svc"
	"min-tiktok/services/message/message"
)

type MessageActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMessageActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessageActionLogic {
	return &MessageActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MessageActionLogic) MessageAction(in *message.MessageActionRequest) (*message.MessageActionResponse, error) {
	// check is friend
	res, err := l.svcCtx.RelationRpc.IsFriend(l.ctx, &relation.IsFriendRequest{
		ActorId: in.ActorId,
		UserId:  in.UserId,
	})
	if err != nil {
		l.Errorw("call rpc RelationRpc.IsFriend", logx.Field("err", err))
		return nil, err
	}
	if !res.Result {
		return &message.MessageActionResponse{
			StatusCode: code.IsNotFriendCode,
			StatusMsg:  code.IsNotFriendMsg,
		}, nil
	}
	conversationid := fmt.Sprintf("%d-%d", in.ActorId, in.UserId)
	if in.ActorId > in.UserId {
		conversationid = fmt.Sprintf("%d-%d", in.UserId, in.ActorId)
	}
	msg := &message2.Messages{
		Content:        in.Content,
		Fromuserid:     uint64(in.ActorId),
		Touserid:       uint64(in.UserId),
		Conversationid: conversationid,
		Createdat:      time.Now(),
	}
	if _, err := l.svcCtx.MessageModel.Insert(l.ctx, msg); err != nil {
		logx.Errorw("message insert error", logx.Field("err", err))
		return nil, err
	}
	return &message.MessageActionResponse{}, nil
}
