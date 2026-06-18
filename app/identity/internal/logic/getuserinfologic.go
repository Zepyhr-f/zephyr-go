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

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 【用户信息】提供给网关或其他服务拉取用户的全量聚合信息（带上缓存设计）
func (l *GetUserInfoLogic) GetUserInfo(in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	var user model.SysUser

	// 预加载 Roles 及其关联的 Menus
	query := l.svcCtx.DB.Preload("Roles").Preload("Roles.Menus").Where("code = ?", in.UserCode)
	if in.TenantCode != "" {
		query = query.Where("tenant_code = ?", in.TenantCode)
	}

	if err := query.First(&user).Error; err != nil {
		l.Logger.Errorf("Failed to fetch user %s: %v", in.UserCode, err)
		return nil, status.Error(codes.NotFound, "用户不存在")
	}

	resp := &pb.GetUserInfoResp{
		User: &pb.UserEntity{
			UserCode: user.Code,
			Username: user.NickName, // 或者用 nickname，根据实际情况
			Avatar:   user.Avatar,
		},
		Roles: make([]string, 0),
		Perms: make([]string, 0),
	}

	permSet := make(map[string]struct{})

	for _, role := range user.Roles {
		resp.Roles = append(resp.Roles, role.Code)
		for _, menu := range role.Menus {
			if menu.Perms != "" {
				if _, ok := permSet[menu.Perms]; !ok {
					permSet[menu.Perms] = struct{}{}
					resp.Perms = append(resp.Perms, menu.Perms)
				}
			}
		}
	}

	return resp, nil
}
