// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package user

import (
	"context"

	"zephyr-go/app/gateway/internal/svc"
	"zephyr-go/app/gateway/internal/types"
	"zephyr-go/app/identity/identityservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserResetPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserResetPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserResetPasswordLogic {
	return &UserResetPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserResetPasswordLogic) UserResetPassword(req *types.UserResetPasswordReq) error {
	_, err := l.svcCtx.IdentityRpc.UserResetPassword(l.ctx, &identityservice.UserResetPasswordReq{
		Id: req.Id,
	})
	return err
}
