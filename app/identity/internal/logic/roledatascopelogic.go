package logic

import (
	"context"

	"zephyr-go/app/identity/internal/repository/model"
	"zephyr-go/app/identity/internal/svc"
	"zephyr-go/app/identity/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type RoleDataScopeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleDataScopeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleDataScopeLogic {
	return &RoleDataScopeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleDataScopeLogic) RoleDataScope(in *pb.RoleDataScopeReq) (*pb.SuccessResp, error) {
	if in.RoleCode == "" {
		return nil, status.Error(codes.InvalidArgument, "角色编码不能为空")
	}
	var role model.SysRole
	if err := l.svcCtx.DB.Where("code = ? AND if_deleted = ?", in.RoleCode, 0).First(&role).Error; err != nil {
		return nil, status.Error(codes.NotFound, "角色不存在")
	}

	err := l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("role_code = ?", in.RoleCode).Delete(&model.SysRoleDept{}).Error; err != nil {
			return err
		}
		for _, deptCode := range in.DeptCodes {
			if deptCode == "" {
				continue
			}
			if err := tx.Create(&model.SysRoleDept{RoleCode: in.RoleCode, DeptCode: deptCode}).Error; err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		l.Errorf("save role data scope failed: %v", err)
		return nil, err
	}
	return &pb.SuccessResp{Success: true}, nil
}
