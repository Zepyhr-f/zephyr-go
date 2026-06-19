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

type RoleMenuTreeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleMenuTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleMenuTreeLogic {
	return &RoleMenuTreeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleMenuTreeLogic) RoleMenuTree(in *pb.RoleMenuTreeReq) (*pb.RoleMenuTreeResp, error) {
	if in.RoleCode == "" {
		return nil, status.Error(codes.InvalidArgument, "角色编码不能为空")
	}
	var menus []model.SysMenu
	if err := l.svcCtx.DB.Where("if_deleted = ?", 0).Order("order_num asc, id asc").Find(&menus).Error; err != nil {
		l.Errorf("query menus failed: %v", err)
		return nil, err
	}
	var relations []model.SysRoleMenu
	if err := l.svcCtx.DB.Where("role_code = ?", in.RoleCode).Find(&relations).Error; err != nil {
		l.Errorf("query role menus failed: %v", err)
		return nil, err
	}
	checked := make([]string, 0, len(relations))
	for _, relation := range relations {
		checked = append(checked, relation.MenuCode)
	}
	return &pb.RoleMenuTreeResp{Menus: buildMenuTree(menus, "-1"), CheckedKeys: checked}, nil
}
