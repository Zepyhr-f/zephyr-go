package logic

import (
	"context"

	"zephyr-go/app/identity/internal/repository/model"
	"zephyr-go/app/identity/internal/svc"
	"zephyr-go/app/identity/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeptSaveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeptSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptSaveLogic {
	return &DeptSaveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeptSaveLogic) DeptSave(in *pb.DeptSaveReq) (*pb.SuccessResp, error) {
	dept := model.SysDept{
		ParentCode: in.ParentCode,
		DeptName:   in.DeptName,
		OrderNum:   int(in.OrderNum),
		Status:     int16(in.Status),
	}
	dept.Code = in.Code
	if err := l.svcCtx.DB.Create(&dept).Error; err != nil {
		l.Errorf("save dept failed: %v", err)
		return nil, err
	}

	return &pb.SuccessResp{Success: true}, nil
}
