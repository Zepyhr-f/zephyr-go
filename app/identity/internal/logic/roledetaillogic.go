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

type RoleDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleDetailLogic {
	return &RoleDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleDetailLogic) RoleDetail(in *pb.RoleDetailReq) (*pb.RoleDetailResp, error) {
	if in.Code == "" {
		return nil, status.Error(codes.InvalidArgument, "角色编码不能为空")
	}
	var role model.SysRole
	if err := l.svcCtx.DB.Where("code = ? AND if_deleted = ?", in.Code, 0).First(&role).Error; err != nil {
		return nil, status.Error(codes.NotFound, "角色不存在")
	}
	var relations []model.SysRoleMenu
	if err := l.svcCtx.DB.Where("role_code = ?", role.Code).Find(&relations).Error; err != nil {
		l.Errorf("query role menus failed: %v", err)
		return nil, err
	}
	menuCodes := make([]string, 0, len(relations))
	for _, relation := range relations {
		menuCodes = append(menuCodes, relation.MenuCode)
	}
	return &pb.RoleDetailResp{Role: roleToPB(role), MenuCodes: menuCodes}, nil
}
