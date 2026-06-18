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

type RoleUpdateStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleUpdateStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleUpdateStatusLogic {
	return &RoleUpdateStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleUpdateStatusLogic) RoleUpdateStatus(req *types.RoleUpdateStatusReq) error {
	_, err := l.svcCtx.IdentityRpc.RoleUpdateStatus(l.ctx, &identityservice.RoleUpdateStatusReq{
		Id:     req.Id,
		Status: req.Status,
	})
	return err
}
