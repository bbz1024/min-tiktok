package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
	"min-tiktok/common/config"
)

type Config struct {
	zrpc.RpcServerConf
	Consul               consul.Conf
	MySQL                config.MysqlStructure
	UserRpc, FeedBackRpc zrpc.RpcClientConf
	CacheConf            cache.CacheConf
}
