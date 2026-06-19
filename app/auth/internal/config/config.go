package config

import (
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	IdentityRpc zrpc.RpcClientConf
	BizRedis    struct {
		Host string
		Pass string
		DB   int
	}

	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
}
