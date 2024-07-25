package main

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	client "min-tiktok/common/util/gorse"
	"min-tiktok/services/auths/internal/config"
	"min-tiktok/services/auths/internal/svc"
	"testing"
)

var c config.Config
var ctx *svc.ServiceContext

func init() {
	conf.MustLoad(*configFile, &c)
	ctx = svc.NewServiceContext(c)
}

// 加载用户数据到Gorse
func TestInsert2Gorse(t *testing.T) {
	gorseClient := client.NewGorseClient(ctx.Config.Gorse.GorseAddr, ctx.Config.Gorse.GorseApikey)
	// -------------------- insert user --------------------
	allUser, err := ctx.UserModel.GetAllUser(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	var users []client.User
	for _, user := range allUser {
		users = append(users, client.User{
			UserId:  fmt.Sprintf("%d", user.Id),
			Comment: user.Signature.String,
		})
	}
	rowAffected, err := gorseClient.InsertUsers(context.Background(), users)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("insert %d items", rowAffected)
}
