package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

type Config struct {
	rest.RestConf
	AuthsRpc zrpc.RpcClientConf // 认证rpc
	FeedRpc  zrpc.RpcClientConf // 认证rpc
	Consul   consul.Conf
}
