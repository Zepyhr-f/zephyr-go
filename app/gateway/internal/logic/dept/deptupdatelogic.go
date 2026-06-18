// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package dept

import (
	"context"

	"zephyr-go/app/gateway/internal/svc"
	"zephyr-go/app/gateway/internal/types"
	"zephyr-go/app/identity/identityservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeptUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeptUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptUpdateLogic {
	return &DeptUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeptUpdateLogic) DeptUpdate(req *types.DeptUpdateReq) error {
	_, err := l.svcCtx.IdentityRpc.DeptUpdate(l.ctx, &identityservice.DeptUpdateReq{
		Id:         req.Id,
		Code:       req.Code,
		ParentCode: req.ParentCode,
		DeptName:   req.DeptName,
		OrderNum:   req.OrderNum,
		Status:     req.Status,
	})
	return err
}
