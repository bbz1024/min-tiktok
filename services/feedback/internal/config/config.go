package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
	"min-tiktok/common/config"
)

type Config struct {
	zrpc.RpcServerConf
	Consul    consul.Conf
	Gorse     config.GorseStructure
	RabbitMQ  config.RabbitMQStructure
	RedisConf redis.RedisConf
}