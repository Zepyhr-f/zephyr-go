package logic

import (
	"context"

	"zephyr-go/app/identity/internal/repository/model"
	"zephyr-go/app/identity/internal/svc"
	"zephyr-go/app/identity/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuRemoveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuRemoveLogic {
	return &MenuRemoveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MenuRemoveLogic) MenuRemove(in *pb.MenuRemoveReq) (*pb.SuccessResp, error) {
	if len(in.Codes) == 0 {
		return &pb.SuccessResp{Success: true}, nil
	}
	if err := l.svcCtx.DB.Where("code IN ?", in.Codes).Delete(&model.SysMenu{}).Error; err != nil {
		l.Errorf("remove menus failed: %v", err)
		return nil, err
	}
	if err := l.svcCtx.DB.Where("menu_code IN ?", in.Codes).Delete(&model.SysRoleMenu{}).Error; err != nil {
		l.Errorf("remove menu relations failed: %v", err)
		return nil, err
	}
	return &pb.SuccessResp{Success: true}, nil
}
