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

type UserRemoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRemoveLogic {
	return &UserRemoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRemoveLogic) UserRemove(req *types.UserRemoveReq) error {
	_, err := l.svcCtx.IdentityRpc.UserRemove(l.ctx, &identityservice.UserRemoveReq{
		Ids: req.Ids,
	})
	return err
}
