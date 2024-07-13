package video

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
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
		ListVideoByVideoSet(ctx context.Context, videoSet []string) ([]*Video, error)
		GetVideoIds(ctx context.Context) ([]uint32, error)
	}

	customVideoModel struct {
		*defaultVideoModel
	}
)

func (m *customVideoModel) GetVideoIds(ctx context.Context) ([]uint32, error) {
	query := fmt.Sprintf("SELECT `id` FROM %s", m.table)
	var videoIds []uint32
	err := m.conn.QueryRowsCtx(ctx, &videoIds, query)
	switch {
	case err == nil:
		return videoIds, nil
	case errors.Is(err, sqlx.ErrNotFound):
		return videoIds, nil
	default:
		return nil, err
	}
}

func (m *customVideoModel) ListVideoByVideoSet(ctx context.Context, videoSet []string) ([]*Video, error) {
	// 构建视频ID IN 查询的条件字符串，例如 "video_id IN ('id1', 'id2', ...)"
	var videoIdsPlaceholder strings.Builder
	if len(videoSet) == 0 {
		return []*Video{}, nil
	}
	for i, id := range videoSet {
		if i > 0 {
			videoIdsPlaceholder.WriteString(", ")
		}
		videoIdsPlaceholder.WriteString(id)
	}

	query := fmt.Sprintf("SELECT %s FROM %s WHERE `id` IN (%s)", videoRows, m.table, videoIdsPlaceholder.String())

	var videos []*Video
	err := m.conn.QueryRowsCtx(ctx, &videos, query)
	switch {
	case err == nil:
		return videos, nil
	case errors.Is(err, sqlx.ErrNotFound):
		return []*Video{}, nil // 返回空切片表示没有找到匹配项，而不是nil
	default:
		return nil, err
	}
}
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
