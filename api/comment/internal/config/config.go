package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
	"min-tiktok/common/config"
)

type Config struct {
	rest.RestConf
	AuthsRpc   zrpc.RpcClientConf
	Consul     consul.Conf
	MySQL      config.MysqlStructure
	CommentRpc zrpc.RpcClientConf
}
