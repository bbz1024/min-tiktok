package svc

import (
	"context"
	"fmt"
	"github.com/willf/bloom"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"min-tiktok/api/auths/internal/config"
	"min-tiktok/models/user"
	"min-tiktok/services/auths/authsclient"
)

type ServiceContext struct {
	Config     config.Config
	AuthsRpc   authsclient.Auths
	UserModel  user.UsersModel
	UserFilter *bloom.BloomFilter
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlConn := sqlx.NewMysql(c.MySQL.DataSource)
	userModel := user.NewUsersModel(mysqlConn)
	userFilter := bloom.NewWithEstimates(100000, 0.01)
	//  push user id to bloom filter
	names, err := userModel.GetNamesCtx(context.TODO())
	if err != nil {
		panic(err)
	}
	for _, name := range names {
		fmt.Println(name)
		userFilter.AddString(name)
	}

	return &ServiceContext{
		Config:     c,
		AuthsRpc:   authsclient.NewAuths(zrpc.MustNewClient(c.AuthsRpc)),
		UserModel:  userModel,
		UserFilter: userFilter,
	}
}
