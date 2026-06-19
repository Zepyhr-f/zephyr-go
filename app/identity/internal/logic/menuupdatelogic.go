package logic

import (
	"context"
	"strconv"

	"zephyr-go/app/identity/internal/repository/model"
	"zephyr-go/app/identity/internal/svc"
	"zephyr-go/app/identity/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MenuUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuUpdateLogic {
	return &MenuUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MenuUpdateLogic) MenuUpdate(in *pb.MenuUpdateReq) (*pb.SuccessResp, error) {
	if in.Code == "" && in.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "菜单ID或编码不能为空")
	}

	var menu model.SysMenu
	query := l.svcCtx.DB.Where("if_deleted = ?", 0)
	if in.Id != "" {
		id, err := strconv.ParseInt(in.Id, 10, 64)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "菜单ID格式错误")
		}
		query = query.Where("id = ?", id)
	} else {
		query = query.Where("code = ?", in.Code)
	}
	if err := query.First(&menu).Error; err != nil {
		return nil, status.Error(codes.NotFound, "菜单不存在")
	}

	menu.Code = in.Code
	menu.MenuName = in.MenuName
	menu.ParentCode = in.ParentCode
	menu.OrderNum = int(in.OrderNum)
	menu.Path = in.Path
	menu.Component = in.Component
	menu.MenuType = in.MenuType
	menu.Perms = in.Perms
	menu.Icon = in.Icon
	menu.Status = int16(in.Status)
	if menu.ParentCode == "" {
		menu.ParentCode = "-1"
	}
	if err := l.svcCtx.DB.Save(&menu).Error; err != nil {
		l.Errorf("update menu failed: %v", err)
		return nil, err
	}
	return &pb.SuccessResp{Success: true}, nil
}
