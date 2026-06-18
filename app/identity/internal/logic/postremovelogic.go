package logic

import (
	"context"

	"zephyr-go/app/identity/internal/repository/model"
	"zephyr-go/app/identity/internal/svc"
	"zephyr-go/app/identity/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostRemoveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPostRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostRemoveLogic {
	return &PostRemoveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PostRemoveLogic) PostRemove(in *pb.PostRemoveReq) (*pb.SuccessResp, error) {
	if len(in.Codes) == 0 {
		return &pb.SuccessResp{Success: true}, nil
	}
	if err := l.svcCtx.DB.Where("code IN ?", in.Codes).Delete(&model.SysPost{}).Error; err != nil {
		l.Errorf("remove post failed: %v", err)
		return nil, err
	}

	return &pb.SuccessResp{Success: true}, nil
}
