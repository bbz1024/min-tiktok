package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
	"min-tiktok/common/config"
)

type Config struct {
	zrpc.RpcServerConf
	Consul    consul.Conf
	CacheConf cache.CacheConf
	FeedRpc   zrpc.RpcClientConf

	MySQL      config.MysqlStructure
	QiNiu      config.QiNiuStructure
	RabbitMQ   config.RabbitMQStructure
	AlibabaNsl config.AlibabaNslStructure
	Gpt        config.GptStructure
	Gorse      config.GorseStructure
}
