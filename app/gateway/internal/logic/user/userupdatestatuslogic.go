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

type UserUpdateStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserUpdateStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUpdateStatusLogic {
	return &UserUpdateStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserUpdateStatusLogic) UserUpdateStatus(req *types.UserUpdateStatusReq) error {
	_, err := l.svcCtx.IdentityRpc.UserUpdateStatus(l.ctx, &identityservice.UserUpdateStatusReq{
		Id:     req.Id,
		Status: req.Status,
	})
	return err
}
