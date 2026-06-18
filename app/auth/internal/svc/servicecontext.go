package svc

import (
	"zephyr-go/app/auth/internal/config"
	"zephyr-go/app/identity/identityservice"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	IdentityRpc identityservice.IdentityService
	BizRedis    *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		IdentityRpc: identityservice.NewIdentityService(zrpc.MustNewClient(c.IdentityRpc)),
		BizRedis: redis.New(c.BizRedis.Host, func(r *redis.Redis) {
			r.Type = redis.NodeType
			r.Pass = c.BizRedis.Pass
		}),
	}
}
