package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

type Config struct {
	zrpc.RpcServerConf
	IdentityRpc zrpc.RpcClientConf
	BizRedis    struct {
		Host string
		Pass string
		DB   int
	}
	Consul consul.Conf
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
}
