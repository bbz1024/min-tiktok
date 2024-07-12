// Code generated by goctl. DO NOT EDIT.

package video

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	videoFieldNames          = builder.RawFieldNames(&Video{})
	videoRows                = strings.Join(videoFieldNames, ",")
	videoRowsExpectAutoSet   = strings.Join(stringx.Remove(videoFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	videoRowsWithPlaceHolder = strings.Join(stringx.Remove(videoFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	videoModel interface {
		Insert(ctx context.Context, data *Video) (sql.Result, error)
		FindOne(ctx context.Context, id uint64) (*Video, error)
		Update(ctx context.Context, data *Video) error
		Delete(ctx context.Context, id uint64) error
	}

	defaultVideoModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Video struct {
		Id        uint64         `db:"id"`
		Userid    uint64         `db:"userid"`     // 用户id
		Title     string         `db:"title"`      // 标题
		Playurl   string         `db:"playurl"`    // 文件名
		Coverurl  string         `db:"coverurl"`   // 封面名
		Content   sql.NullString `db:"content"`    // 摘要
		CreatedAt time.Time      `db:"created_at"` // 创建时间
		UpdatedAt time.Time      `db:"updated_at"` // 更新时间
	}
)

func newVideoModel(conn sqlx.SqlConn) *defaultVideoModel {
	return &defaultVideoModel{
		conn:  conn,
		table: "`video`",
	}
}

func (m *defaultVideoModel) Delete(ctx context.Context, id uint64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultVideoModel) FindOne(ctx context.Context, id uint64) (*Video, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", videoRows, m.table)
	var resp Video
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultVideoModel) Insert(ctx context.Context, data *Video) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, videoRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Userid, data.Title, data.Playurl, data.Coverurl, data.Content)
	return ret, err
}

func (m *defaultVideoModel) Update(ctx context.Context, data *Video) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, videoRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Userid, data.Title, data.Playurl, data.Coverurl, data.Content, data.Id)
	return err
}

func (m *defaultVideoModel) tableName() string {
	return m.table
}
