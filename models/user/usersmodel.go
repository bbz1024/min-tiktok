package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UsersModel = (*customUsersModel)(nil)

type (
	// UsersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUsersModel.
	UsersModel interface {
		usersModel
		withSession(session sqlx.Session) UsersModel
		GetNamesCtx(crx context.Context) ([]string, error)
		GetAllUserId(ctx context.Context) ([]string, error)
	}

	customUsersModel struct {
		*defaultUsersModel
	}
)

func (m *customUsersModel) GetAllUserId(ctx context.Context) ([]string, error) {
	var resp []string
	query := fmt.Sprintf("select id from %s ", m.table)
	err := m.conn.QueryRowsCtx(ctx, &resp, query)
	switch {
	case err == nil:
		return resp, nil
	case errors.Is(err, sqlx.ErrNotFound):
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customUsersModel) GetNamesCtx(ctx context.Context) ([]string, error) {
	var resp []*Users
	query := fmt.Sprintf("select %s from %s ", usersRows, m.table)
	err := m.conn.QueryRowsCtx(ctx, &resp, query)

	switch {
	case err == nil:
		names := make([]string, 0, len(resp))
		for _, user := range resp {
			names = append(names, user.Username)
		}
		return names, nil
	case errors.Is(err, sqlx.ErrNotFound):
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// NewUsersModel returns a model for the database table.
func NewUsersModel(conn sqlx.SqlConn) UsersModel {
	return &customUsersModel{
		defaultUsersModel: newUsersModel(conn),
	}
}

func (m *customUsersModel) withSession(session sqlx.Session) UsersModel {
	return NewUsersModel(sqlx.NewSqlConnFromSession(session))
}
