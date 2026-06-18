// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package role

import (
	"context"

	"zephyr-go/app/gateway/internal/svc"
	"zephyr-go/app/gateway/internal/types"
	"zephyr-go/app/identity/identityservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleRemoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleRemoveLogic {
	return &RoleRemoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleRemoveLogic) RoleRemove(req *types.RoleRemoveReq) error {
	_, err := l.svcCtx.IdentityRpc.RoleRemove(l.ctx, &identityservice.RoleRemoveReq{
		Ids: req.Ids,
	})
	return err
}
