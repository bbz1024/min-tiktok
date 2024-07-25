package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"min-tiktok/common/consts/keys"
	client "min-tiktok/common/util/gorse"
	"min-tiktok/services/publish/internal/config"
	"min-tiktok/services/publish/internal/svc"
	"strings"
	"testing"
)

var c config.Config
var ctx *svc.ServiceContext

func init() {
	conf.MustLoad(*configFile, &c)
	ctx = svc.NewServiceContext(c)
}

// 加载视频信息到redis
func TestPreloadVideo2Redis(t *testing.T) {
	//get all video
	ids, err := ctx.VideoModel.GetVideoIds(context.Background())
	if err != nil {
		t.Error(err)
		return
	}
	userID := 1
	for _, id := range ids {
		key := fmt.Sprintf(keys.UserInfoKey, userID)
		if _, err := ctx.Rdb.HincrbyCtx(context.Background(), key, keys.WorkCount, 1); err != nil && !errors.Is(err, redis.Nil) {
			logx.Errorw("incr user work count ", logx.Field("err", err))
			return
		}

		// 4. add video id to user video list
		key = fmt.Sprintf(keys.UserWorkKey, userID)
		if _, err := ctx.Rdb.SaddCtx(context.Background(), key, id); err != nil {
			logx.Errorw("add video id to user video list ", logx.Field("err", err))
			return
		}
		videoInfoKey := fmt.Sprintf(keys.VideoInfoKey, id)
		if err := ctx.Rdb.HsetCtx(context.Background(), videoInfoKey, keys.VideoAuthorID, fmt.Sprintf("%d", userID)); err != nil {
			logx.Errorw("set video author id ", logx.Field("err", err))
			return
		}
	}
}

// 插入视频到gorse
func TestInsert2Gorse(t *testing.T) {
	gorseClient := client.NewGorseClient(ctx.Config.Gorse.GorseAddr, ctx.Config.Gorse.GorseApikey)
	// -------------------- insert video --------------------
	videinfo, err := ctx.VideoInfoModel.GetAll(context.Background())
	if err != nil {
		t.Error(err)
		return
	}
	var items []client.Item
	for _, v := range videinfo {
		items = append(items, client.Item{
			ItemId:     fmt.Sprintf("%d", v.Id),
			Timestamp:  v.CreatedAt.String(),
			Categories: strings.Split(v.Category.String, "|"),
			Labels:     strings.Split(v.Keyword.String, "|"),
		})
	}
	rowAffected, err := gorseClient.InsertItems(context.Background(), items)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("insert %d items", rowAffected)
}
