package logic

import (
	"context"

	"zephyr-go/app/identity/internal/repository/model"
	"zephyr-go/app/identity/internal/svc"
	"zephyr-go/app/identity/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeptRemoveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeptRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptRemoveLogic {
	return &DeptRemoveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeptRemoveLogic) DeptRemove(in *pb.DeptRemoveReq) (*pb.SuccessResp, error) {
	if len(in.Ids) == 0 {
		return &pb.SuccessResp{Success: true}, nil
	}
	if err := l.svcCtx.DB.Where("code IN ?", in.Ids).Delete(&model.SysDept{}).Error; err != nil {
		l.Errorf("remove dept failed: %v", err)
		return nil, err
	}

	return &pb.SuccessResp{Success: true}, nil
}
