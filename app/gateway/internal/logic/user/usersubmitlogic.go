// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package user

import (
	"context"

	"zephyr-go/app/gateway/internal/svc"
	"zephyr-go/app/gateway/internal/types"
	"zephyr-go/app/identity/identityservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserSubmitLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserSubmitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserSubmitLogic {
	return &UserSubmitLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserSubmitLogic) UserSubmit(req *types.UserSubmitReq) error {
	_, err := l.svcCtx.IdentityRpc.UserSubmit(l.ctx, &identityservice.UserSubmitReq{
		Id:       req.Id,
		Code:     req.Code,
		NickName: req.NickName,
		RealName: req.RealName,
		Phone:    req.Phone,
		Email:    req.Email,
		DeptCode: req.DeptCode,
		Sex:      req.Sex,
		Status:   req.Status,
	})
	return err
}
