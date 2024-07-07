package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

type Config struct {
	zrpc.RpcServerConf
	Consul consul.Conf
	MySQL  struct {
		DataSource string
	}
	RedisConf redis.RedisConf
	QiNiu     struct {
		AccessKey   string
		SecretKey   string
		VideoDomain string
		Bucket      string
	}
	FeedRpc zrpc.RpcClientConf
}
