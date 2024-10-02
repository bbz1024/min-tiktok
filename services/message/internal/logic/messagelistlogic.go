package logic

import (
	"context"
	"min-tiktok/services/message/internal/svc"
	"min-tiktok/services/message/message"

	"github.com/zeromicro/go-zero/core/logx"
)

type MessageListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMessageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessageListLogic {
	return &MessageListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx).WithFields(logx.Field("type", "service")),
	}
}

func (l *MessageListLogic) MessageList(in *message.MessageListRequest) (*message.MessageListResponse, error) {

	messageList, err := l.svcCtx.MessageModel.QueryMessageListByTime(l.ctx, in.ToUserId, in.ActorId, in.PreMsgTime)
	if err != nil {
		l.Errorw("query message list failed", logx.Field("err", err))
		return nil, err
	}
	resp := new(message.MessageListResponse)
	resp.MessageList = make([]*message.MessageInfo, 0, len(messageList))
	for _, v := range messageList {
		resp.MessageList = append(resp.MessageList, &message.MessageInfo{
			Id:         uint32(v.Id),
			ToUserId:   uint32(v.Touserid),
			FromUserId: uint32(v.Fromuserid),
			Content:    v.Content,
			CreateTime: uint64(v.Createdat.UTC().UnixMilli()),
		})
	}

	return resp, nil
}
