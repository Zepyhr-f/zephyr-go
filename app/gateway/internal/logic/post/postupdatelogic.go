// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package post

import (
	"context"

	"zephyr-go/app/gateway/internal/svc"
	"zephyr-go/app/gateway/internal/types"
	"zephyr-go/app/identity/identityservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostUpdateLogic {
	return &PostUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostUpdateLogic) PostUpdate(req *types.PostUpdateReq) error {
	_, err := l.svcCtx.IdentityRpc.PostUpdate(l.ctx, &identityservice.PostUpdateReq{
		Id:       req.Id,
		Code:     req.Code,
		PostName: req.PostName,
		OrderNum: req.OrderNum,
		Status:   req.Status,
	})
	return err
}
