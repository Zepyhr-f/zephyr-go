package logic

import (
	"context"

	"zephyr-go/app/identity/internal/repository/model"
	"zephyr-go/app/identity/internal/svc"
	"zephyr-go/app/identity/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeptUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeptUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptUpdateLogic {
	return &DeptUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeptUpdateLogic) DeptUpdate(in *pb.DeptUpdateReq) (*pb.SuccessResp, error) {
	var dept model.SysDept
	if err := l.svcCtx.DB.Where("code = ?", in.Code).First(&dept).Error; err != nil {
		l.Errorf("find dept failed: %v", err)
		return nil, err
	}
	
	dept.DeptName = in.DeptName
	dept.ParentCode = in.ParentCode
	dept.OrderNum = int(in.OrderNum)
	dept.Status = int16(in.Status)

	if err := l.svcCtx.DB.Save(&dept).Error; err != nil {
		l.Errorf("update dept failed: %v", err)
		return nil, err
	}

	return &pb.SuccessResp{Success: true}, nil
}
