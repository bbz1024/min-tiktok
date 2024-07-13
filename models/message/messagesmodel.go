package message

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ MessagesModel = (*customMessagesModel)(nil)

type (
	// MessagesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMessagesModel.
	MessagesModel interface {
		messagesModel
		withSession(session sqlx.Session) MessagesModel
		QueryMessageListByTime(ctx context.Context, userId, actorId uint32, preMsgTime uint64) ([]*Messages, error)
	}

	customMessagesModel struct {
		*defaultMessagesModel
	}
)

func (m *customMessagesModel) QueryMessageListByTime(ctx context.Context, userId, actorId uint32, preMsgTime uint64) ([]*Messages, error) {

	var messageList []*Messages
	conversationid := fmt.Sprintf("%d-%d", userId, actorId)
	if userId > actorId {
		conversationid = fmt.Sprintf("%d-%d", actorId, userId)
	}
	if preMsgTime == 0 {
		query := fmt.Sprintf("select %s from %s where conversationid = ?  order by createdat", messagesRows, m.table)
		if err := m.conn.QueryRowsCtx(ctx, &messageList, query, conversationid); err != nil {
			return nil, err
		}
	} else {
		query := fmt.Sprintf(
			"select %s from %s where conversationid = ? and createdat > ? order by createdat", messagesRows, m.table,
		)
		if err := m.conn.QueryRowsCtx(ctx, &messageList, query, conversationid, preMsgTime); err != nil {
			return nil, err
		}
	}
	return messageList, nil

}

// NewMessagesModel returns a model for the database table.
func NewMessagesModel(conn sqlx.SqlConn) MessagesModel {
	return &customMessagesModel{
		defaultMessagesModel: newMessagesModel(conn),
	}
}

func (m *customMessagesModel) withSession(session sqlx.Session) MessagesModel {
	return NewMessagesModel(sqlx.NewSqlConnFromSession(session))
}
