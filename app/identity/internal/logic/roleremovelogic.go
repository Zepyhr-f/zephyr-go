package logic

import (
	"context"

	"zephyr-go/app/identity/internal/repository/model"
	"zephyr-go/app/identity/internal/svc"
	"zephyr-go/app/identity/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleRemoveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleRemoveLogic {
	return &RoleRemoveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleRemoveLogic) RoleRemove(in *pb.RoleRemoveReq) (*pb.SuccessResp, error) {
	if len(in.Ids) == 0 {
		return &pb.SuccessResp{Success: true}, nil
	}
	if err := l.svcCtx.DB.Where("id IN ?", in.Ids).Delete(&model.SysRole{}).Error; err != nil {
		l.Errorf("remove roles failed: %v", err)
		return nil, err
	}

	return &pb.SuccessResp{Success: true}, nil
}
