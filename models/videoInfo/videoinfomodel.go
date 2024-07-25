package videoInfo

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ VideoinfoModel = (*customVideoinfoModel)(nil)

type (
	// VideoinfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customVideoinfoModel.
	VideoinfoModel interface {
		videoinfoModel
		withSession(session sqlx.Session) VideoinfoModel
		GetAll(ctx context.Context) ([]*Videoinfo, error)
	}

	customVideoinfoModel struct {
		*defaultVideoinfoModel
	}
)

func (m *defaultVideoinfoModel) GetAll(ctx context.Context) ([]*Videoinfo, error) {
	var videoinfo []*Videoinfo
	query := fmt.Sprintf("select %s from %s", videoinfoRows, m.table)
	err := m.conn.QueryRowsCtx(ctx, &videoinfo, query)
	switch {
	case err == nil:
		return videoinfo, nil
	case errors.Is(err, sqlx.ErrNotFound):
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// NewVideoinfoModel returns a model for the database table.
func NewVideoinfoModel(conn sqlx.SqlConn) VideoinfoModel {
	return &customVideoinfoModel{
		defaultVideoinfoModel: newVideoinfoModel(conn),
	}
}

func (m *customVideoinfoModel) withSession(session sqlx.Session) VideoinfoModel {
	return NewVideoinfoModel(sqlx.NewSqlConnFromSession(session))
}
