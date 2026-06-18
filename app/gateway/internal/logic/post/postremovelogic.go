// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package post

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"zephyr-go/app/gateway/internal/svc"
	"zephyr-go/app/gateway/internal/types"
	"zephyr-go/app/identity/identityservice"
)

type PostRemoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostRemoveLogic {
	return &PostRemoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostRemoveLogic) PostRemove(req *types.PostRemoveReq) error {
	_, err := l.svcCtx.IdentityRpc.PostRemove(l.ctx, &identityservice.PostRemoveReq{
		Codes: req.Codes,
	})
	return err
}
