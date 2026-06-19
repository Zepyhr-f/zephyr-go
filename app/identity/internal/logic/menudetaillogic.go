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

type MenuDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuDetailLogic {
	return &MenuDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MenuDetailLogic) MenuDetail(in *pb.MenuDetailReq) (*pb.MenuDetailResp, error) {
	if in.Code == "" {
		return nil, status.Error(codes.InvalidArgument, "菜单编码不能为空")
	}

	var menu model.SysMenu
	if err := l.svcCtx.DB.Where("code = ? AND if_deleted = ?", in.Code, 0).First(&menu).Error; err != nil {
		return nil, status.Error(codes.NotFound, "菜单不存在")
	}

	return &pb.MenuDetailResp{Menu: menuToPB(menu)}, nil
}
