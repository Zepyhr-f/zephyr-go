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

type DeptSaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeptSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptSaveLogic {
	return &DeptSaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeptSaveLogic) DeptSave(req *types.DeptSaveReq) error {
	_, err := l.svcCtx.IdentityRpc.DeptSave(l.ctx, &identityservice.DeptSaveReq{
		Code:       req.Code,
		ParentCode: req.ParentCode,
		DeptName:   req.DeptName,
		OrderNum:   req.OrderNum,
		Status:     req.Status,
	})
	return err
}
