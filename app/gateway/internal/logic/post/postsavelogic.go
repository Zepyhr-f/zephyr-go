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

type PostSaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostSaveLogic {
	return &PostSaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostSaveLogic) PostSave(req *types.PostSaveReq) error {
	_, err := l.svcCtx.IdentityRpc.PostSave(l.ctx, &identityservice.PostSaveReq{
		Code:     req.Code,
		PostName: req.PostName,
		OrderNum: req.OrderNum,
		Status:   req.Status,
	})
	return err
}
