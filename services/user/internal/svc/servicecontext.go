package svc

import (
	"context"
	"github.com/willf/bloom"
	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"min-tiktok/models/user"
	"min-tiktok/services/user/internal/config"
	"time"
)

type ServiceContext struct {
	Config     config.Config
	UserModel  user.UsersModel
	Rdb        *redis.Redis
	UserBloom  *bloom.BloomFilter
	LocalCache *collection.Cache
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlConn := sqlx.NewMysql(c.MySQL.DataSource)
	userModel := user.NewUsersModel(mysqlConn, c.CacheConf)
	cache, err := collection.NewCache(time.Hour, collection.WithLimit(50))
	if err != nil {
		panic(err)
	}
	rdb, err := redis.NewRedis(c.CacheConf[0].RedisConf)
	if err != nil {
		panic(err)
	}
	// init bloom c.Bloom.EstimateN, c.Bloom.EstimateFP
	estimates := bloom.NewWithEstimates(10000, 0.01)
	// put user_id into bloom filter
	ids, err := userModel.QueryAllUserID(context.TODO())
	for _, id := range ids {
		estimates.AddString(id)
	}
	logx.Info("bloom filter init finished")
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:     c,
		UserModel:  userModel,
		Rdb:        rdb,
		UserBloom:  estimates,
		LocalCache: cache,
	}
}
