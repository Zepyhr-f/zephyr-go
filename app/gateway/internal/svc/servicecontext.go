// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package svc

import (
	"zephyr-go/app/gateway/internal/config"
	"zephyr-go/app/auth/authservice"
	"zephyr-go/app/identity/identityservice"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	AuthRpc     authservice.AuthService
	IdentityRpc identityservice.IdentityService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		AuthRpc:     authservice.NewAuthService(zrpc.MustNewClient(c.AuthRpc)),
		IdentityRpc: identityservice.NewIdentityService(zrpc.MustNewClient(c.IdentityRpc)),
	}
}
