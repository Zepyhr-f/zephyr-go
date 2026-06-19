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

type MenuSaveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuSaveLogic {
	return &MenuSaveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MenuSaveLogic) MenuSave(in *pb.MenuSaveReq) (*pb.SuccessResp, error) {
	if in.Code == "" || in.MenuName == "" {
		return nil, status.Error(codes.InvalidArgument, "菜单编码和名称不能为空")
	}
	menu := model.SysMenu{
		MenuName:   in.MenuName,
		ParentCode: in.ParentCode,
		OrderNum:   int(in.OrderNum),
		Path:       in.Path,
		Component:  in.Component,
		MenuType:   in.MenuType,
		Status:     int16(in.Status),
		Perms:      in.Perms,
		Icon:       in.Icon,
	}
	menu.Code = in.Code
	if menu.ParentCode == "" {
		menu.ParentCode = "-1"
	}
	if menu.Status == 0 {
		menu.Status = 1
	}
	if err := l.svcCtx.DB.Create(&menu).Error; err != nil {
		l.Errorf("create menu failed: %v", err)
		return nil, err
	}
	return &pb.SuccessResp{Success: true}, nil
}
