package video

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
	"time"
)

var _ VideoModel = (*customVideoModel)(nil)

const Count = 5

type (
	// VideoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customVideoModel.
	VideoModel interface {
		videoModel
		ListVideoByVideoSet(ctx context.Context, ids []string) ([]*Video, error)
		ListVideoByUserId(ctx context.Context, i int64) ([]*Video, error)
		ListVideoByCreateTime(ctx context.Context, time time.Time) ([]*Video, error)
		GetVideoIds(ctx context.Context) ([]int, error)
	}

	customVideoModel struct {
		*defaultVideoModel
	}
)

func (c customVideoModel) ListVideoByCreateTime(ctx context.Context, createTime time.Time) ([]*Video, error) {
	query := fmt.Sprintf("select %s from %s where `created_at` < ? order by `created_at` desc limit ? ", videoRows, c.table)
	var resp []*Video

	err := c.CachedConn.QueryRowsNoCacheCtx(ctx, &resp, query, createTime, Count)
	switch {
	case err == nil:
		return resp, nil
	case errors.Is(err, sqlx.ErrNotFound):
		return resp, nil
	default:
		return nil, err
	}
}

func (c customVideoModel) ListVideoByUserId(ctx context.Context, userId int64) ([]*Video, error) {
	query := fmt.Sprintf("select %s from %s where `userid` = ? ", videoRows, c.table)
	var resp []*Video
	err := c.CachedConn.QueryRowsNoCacheCtx(ctx, &resp, query, userId)
	switch {
	case err == nil:
		return resp, nil
	case errors.Is(err, sqlx.ErrNotFound):
		return resp, nil
	default:
		return nil, err
	}
}

func (c customVideoModel) ListVideoByVideoSet(ctx context.Context, videoSet []string) ([]*Video, error) {
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

	query := fmt.Sprintf("SELECT %s FROM %s WHERE `id` IN (%s)", videoRows, c.table, videoIdsPlaceholder.String())

	var videos []*Video
	err := c.CachedConn.QueryRowsNoCacheCtx(ctx, &videos, query)
	switch {
	case err == nil:
		return videos, nil
	case errors.Is(err, sqlx.ErrNotFound):
		return []*Video{}, nil // 返回空切片表示没有找到匹配项，而不是nil
	default:
		return nil, err
	}
}
func (c customVideoModel) GetVideoIds(ctx context.Context) ([]int, error) {
	query := fmt.Sprintf("select %s from %s ", "id", c.table)
	var resp []int
	err := c.CachedConn.QueryRowsNoCacheCtx(ctx, &resp, query)
	switch {
	case err == nil:
		return resp, nil
	case errors.Is(err, sqlx.ErrNotFound):
		return resp, nil
	}
	return nil, err
}

// NewVideoModel returns a model for the database table.
func NewVideoModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) VideoModel {
	return &customVideoModel{
		defaultVideoModel: newVideoModel(conn, c, opts...),
	}
}
