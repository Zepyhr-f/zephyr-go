package logic

import (
	"context"

	"zephyr-go/app/identity/internal/repository/model"
	"zephyr-go/app/identity/internal/svc"
	"zephyr-go/app/identity/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MenuStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuStatusLogic {
	return &MenuStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MenuStatusLogic) MenuStatus(in *pb.MenuStatusReq) (*pb.SuccessResp, error) {
	if in.Code == "" {
		return nil, status.Error(codes.InvalidArgument, "菜单编码不能为空")
	}
	if err := l.svcCtx.DB.Model(&model.SysMenu{}).Where("code = ? AND if_deleted = ?", in.Code, 0).Update("status", in.Status).Error; err != nil {
		l.Errorf("update menu status failed: %v", err)
		return nil, err
	}
	return &pb.SuccessResp{Success: true}, nil
}
