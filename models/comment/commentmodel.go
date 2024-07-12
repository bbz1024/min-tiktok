package comment

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CommentModel = (*customCommentModel)(nil)

type (
	// CommentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCommentModel.
	CommentModel interface {
		commentModel
		GetCommentList(ctx context.Context, videoId uint32) ([]*Comment, error)
		withSession(session sqlx.Session) CommentModel
	}

	customCommentModel struct {
		*defaultCommentModel
	}
)

func (m *customCommentModel) GetCommentList(ctx context.Context, videoId uint32) ([]*Comment, error) {
	query := fmt.Sprintf("select %s from %s where videoid = ? order by ? ", commentRows, m.table)
	var resp []*Comment
	err := m.conn.QueryRowsCtx(ctx, &resp, query, videoId, "createdat")
	switch {
	case err == nil:
		return resp, nil
	case errors.Is(err, sqlx.ErrNotFound):
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// NewCommentModel returns a model for the database table.
func NewCommentModel(conn sqlx.SqlConn) CommentModel {
	return &customCommentModel{
		defaultCommentModel: newCommentModel(conn),
	}
}

func (m *customCommentModel) withSession(session sqlx.Session) CommentModel {
	return NewCommentModel(sqlx.NewSqlConnFromSession(session))
}
