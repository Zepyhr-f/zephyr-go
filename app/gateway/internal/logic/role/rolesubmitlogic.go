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

type RoleSubmitLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleSubmitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleSubmitLogic {
	return &RoleSubmitLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleSubmitLogic) RoleSubmit(req *types.RoleSubmitReq) error {
	_, err := l.svcCtx.IdentityRpc.RoleSubmit(l.ctx, &identityservice.RoleSubmitReq{
		Id:       req.Id,
		Code:     req.Code,
		RoleName: req.RoleName,
		OrderNum: req.OrderNum,
		Status:   req.Status,
		Remark:   req.Remark,
	})
	return err
}
