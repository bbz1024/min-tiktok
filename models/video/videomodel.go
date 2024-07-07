package video

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
)

var _ VideoModel = (*customVideoModel)(nil)

type (
	// VideoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customVideoModel.
	VideoModel interface {
		videoModel
		withSession(session sqlx.Session) VideoModel
		ListVideoByCreateTime(ctx context.Context, createTime time.Time) ([]*Video, error)
		ListVideoByUserId(ctx context.Context, userId int64) ([]*Video, error)
	}

	customVideoModel struct {
		*defaultVideoModel
	}
)

func (m *customVideoModel) ListVideoByUserId(ctx context.Context, userId int64) ([]*Video, error) {
	query := fmt.Sprintf("select %s from %s where `userid` = ? ", videoRows, m.table)
	var resp []*Video
	err := m.conn.QueryRowsCtx(ctx, &resp, query, userId)
	switch {
	case err == nil:
		return resp, nil
	case errors.Is(err, sqlx.ErrNotFound):
		return resp, nil
	default:
		return nil, err
	}
}

const Count = 5

func (m *customVideoModel) ListVideoByCreateTime(ctx context.Context, createTime time.Time) ([]*Video, error) {
	query := fmt.Sprintf("select %s from %s where `created_at` < ? order by `created_at` desc limit ? ", videoRows, m.table)
	var resp []*Video

	err := m.conn.QueryRowsCtx(ctx, &resp, query, createTime, Count)
	switch {
	case err == nil:
		return resp, nil
	case errors.Is(err, sqlx.ErrNotFound):
		return resp, nil
	default:
		return nil, err
	}
}

// NewVideoModel returns a model for the database table.
func NewVideoModel(conn sqlx.SqlConn) VideoModel {
	return &customVideoModel{
		defaultVideoModel: newVideoModel(conn),
	}
}

func (m *customVideoModel) withSession(session sqlx.Session) VideoModel {
	return NewVideoModel(sqlx.NewSqlConnFromSession(session))
}
