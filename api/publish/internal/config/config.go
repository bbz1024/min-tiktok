package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

type Config struct {
	rest.RestConf
	Consul consul.Conf
	// -------------------- rpc --------------------
	AuthsRpc   zrpc.RpcClientConf
	PublishRpc zrpc.RpcClientConf
	UserRpc    zrpc.RpcClientConf
}
