package logic

import (
	"context"

	"zephyr-go/app/identity/internal/repository/model"
	"zephyr-go/app/identity/internal/svc"
	"zephyr-go/app/identity/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleSubmitLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleSubmitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleSubmitLogic {
	return &RoleSubmitLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleSubmitLogic) RoleSubmit(in *pb.RoleSubmitReq) (*pb.SuccessResp, error) {
	if in.Id != "" {
		// update
		var role model.SysRole
		if err := l.svcCtx.DB.Where("code = ?", in.Code).First(&role).Error; err != nil {
			return nil, err
		}
		role.RoleName = in.RoleName
		role.OrderNum = int(in.OrderNum)
		role.Status = int16(in.Status)
		if err := l.svcCtx.DB.Save(&role).Error; err != nil {
			return nil, err
		}
	} else {
		// create
		role := model.SysRole{
			RoleName: in.RoleName,
			OrderNum: int(in.OrderNum),
			Status:   int16(in.Status),
		}
		role.Code = in.Code
		if err := l.svcCtx.DB.Create(&role).Error; err != nil {
			return nil, err
		}
	}

	return &pb.SuccessResp{Success: true}, nil
}
