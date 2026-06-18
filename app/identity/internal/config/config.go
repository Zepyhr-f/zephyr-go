package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

type Config struct {
	zrpc.RpcServerConf
	Postgres struct {
		DSN string
	}
	Consul consul.Conf
}
