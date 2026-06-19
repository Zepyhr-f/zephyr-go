package logic

import (
	"context"

	"zephyr-go/app/identity/internal/repository/model"
	"zephyr-go/app/identity/internal/svc"
	"zephyr-go/app/identity/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuTreeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuTreeLogic {
	return &MenuTreeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Menu
func (l *MenuTreeLogic) MenuTree(in *pb.EmptyReq) (*pb.MenuTreeResp, error) {
	var menus []model.SysMenu
	err := l.svcCtx.DB.Where("if_deleted = ?", 0).Order("order_num asc, id asc").Find(&menus).Error
	if err != nil {
		l.Logger.Errorf("Failed to query menus: %v", err)
		return nil, err
	}

	return &pb.MenuTreeResp{
		List: buildMenuTree(menus, "-1"),
	}, nil
}
