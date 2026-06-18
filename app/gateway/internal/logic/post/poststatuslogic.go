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

type PostStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostStatusLogic {
	return &PostStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostStatusLogic) PostStatus(req *types.PostStatusReq) error {
	_, err := l.svcCtx.IdentityRpc.PostStatus(l.ctx, &identityservice.PostStatusReq{
		Code:   req.Code,
		Status: req.Status,
	})
	return err
}
