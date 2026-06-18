package logic

import (
	"context"

	"zephyr-go/app/identity/internal/repository/model"
	"zephyr-go/app/identity/internal/svc"
	"zephyr-go/app/identity/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleUpdateStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleUpdateStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleUpdateStatusLogic {
	return &RoleUpdateStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleUpdateStatusLogic) RoleUpdateStatus(in *pb.RoleUpdateStatusReq) (*pb.SuccessResp, error) {
	if err := l.svcCtx.DB.Model(&model.SysRole{}).Where("id = ?", in.Id).Update("status", in.Status).Error; err != nil {
		l.Errorf("update role status failed: %v", err)
		return nil, err
	}

	return &pb.SuccessResp{Success: true}, nil
}
