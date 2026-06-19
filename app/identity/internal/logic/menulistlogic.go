package logic

import (
	"context"

	"zephyr-go/app/identity/internal/repository/model"
	"zephyr-go/app/identity/internal/svc"
	"zephyr-go/app/identity/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuListLogic {
	return &MenuListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MenuListLogic) MenuList(in *pb.EmptyReq) (*pb.MenuListResp, error) {
	var menus []model.SysMenu
	if err := l.svcCtx.DB.Where("if_deleted = ?", 0).Order("order_num asc, id asc").Find(&menus).Error; err != nil {
		l.Errorf("query menu list failed: %v", err)
		return nil, err
	}

	records := make([]*pb.MenuDetail, 0, len(menus))
	for _, menu := range menus {
		records = append(records, menuToPB(menu))
	}

	return &pb.MenuListResp{
		Total:   int64(len(records)),
		Records: records,
		Size:    int32(len(records)),
		Current: 1,
	}, nil
}
