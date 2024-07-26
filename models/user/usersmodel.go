package user

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UsersModel = (*customUsersModel)(nil)

type (
	// UsersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUsersModel.
	UsersModel interface {
		usersModel
		QueryAllUsername(ctx context.Context) ([]string, error)
		QueryAllUserID(ctx context.Context) ([]string, error)
	}

	customUsersModel struct {
		*defaultUsersModel
	}
)

func (c customUsersModel) QueryAllUserID(ctx context.Context) ([]string, error) {
	var res []string
	query := fmt.Sprintf("select id from %s", c.table)
	if err := c.CachedConn.QueryRowsNoCacheCtx(ctx, &res, query); err != nil {
		return nil, err
	}
	return res, nil
}

func (c customUsersModel) QueryAllUsername(ctx context.Context) ([]string, error) {
	var res []string
	query := fmt.Sprintf("select username from %s", c.table)
	if err := c.CachedConn.QueryRowsNoCacheCtx(ctx, &res, query); err != nil {
		return nil, err
	}
	return res, nil
}

// NewUsersModel returns a model for the database table.
func NewUsersModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UsersModel {
	return &customUsersModel{
		defaultUsersModel: newUsersModel(conn, c, opts...),
	}
}
