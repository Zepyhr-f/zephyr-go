// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package dept

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"zephyr-go/app/gateway/internal/svc"
	"zephyr-go/app/gateway/internal/types"
	"zephyr-go/app/identity/identityservice"
)

type DeptRemoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeptRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptRemoveLogic {
	return &DeptRemoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeptRemoveLogic) DeptRemove(req *types.DeptRemoveReq) error {
	_, err := l.svcCtx.IdentityRpc.DeptRemove(l.ctx, &identityservice.DeptRemoveReq{
		Ids: req.Ids,
	})
	return err
}
