package logic

import (
	"context"

	"zephyr-go/app/identity/internal/repository/model"
	"zephyr-go/app/identity/internal/svc"
	"zephyr-go/app/identity/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRemoveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRemoveLogic {
	return &UserRemoveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserRemoveLogic) UserRemove(in *pb.UserRemoveReq) (*pb.SuccessResp, error) {
	if len(in.Ids) == 0 {
		return &pb.SuccessResp{Success: true}, nil
	}
	if err := l.svcCtx.DB.Where("id IN ?", in.Ids).Delete(&model.SysUser{}).Error; err != nil {
		l.Errorf("remove users failed: %v", err)
		return nil, err
	}

	return &pb.SuccessResp{Success: true}, nil
}
