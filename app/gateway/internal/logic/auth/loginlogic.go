// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package auth

import (
	"context"

	authpb "zephyr-go/app/auth/pb/pb"
	"zephyr-go/app/gateway/internal/svc"
	"zephyr-go/app/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户登录
func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// 调用 Auth 微服务进行密码比对和 Token 签发
	rpcResp, err := l.svcCtx.AuthRpc.LoginVerify(l.ctx, &authpb.LoginVerifyReq{
		Username:   req.Username,
		Password:   req.Password,
		TenantCode: req.TenantCode,
	})
	if err != nil {
		return nil, err
	}

	return &types.LoginResp{
		Token:        rpcResp.AccessToken,
		RefreshToken: rpcResp.RefreshToken,
		AccessExpire: rpcResp.Expire,
	}, nil
}
