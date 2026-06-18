package logic

import (
	"context"

	"zephyr-go/app/identity/internal/repository/model"
	"zephyr-go/app/identity/internal/svc"
	"zephyr-go/app/identity/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserUpdateStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserUpdateStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUpdateStatusLogic {
	return &UserUpdateStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserUpdateStatusLogic) UserUpdateStatus(in *pb.UserUpdateStatusReq) (*pb.SuccessResp, error) {
	if err := l.svcCtx.DB.Model(&model.SysUser{}).Where("id = ?", in.Id).Update("status", in.Status).Error; err != nil {
		l.Errorf("update user status failed: %v", err)
		return nil, err
	}

	return &pb.SuccessResp{Success: true}, nil
}
