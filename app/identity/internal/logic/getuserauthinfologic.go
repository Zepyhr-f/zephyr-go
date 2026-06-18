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

type GetUserAuthInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserAuthInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserAuthInfoLogic {
	return &GetUserAuthInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 【内部校验】提供给 Auth 认证中心专用的内部接口，用于拉取密码 Hash 及基本身份
func (l *GetUserAuthInfoLogic) GetUserAuthInfo(in *pb.GetUserAuthInfoReq) (*pb.GetUserAuthInfoResp, error) {
	var user model.SysUser

	// 根据 username 查询用户（支持用 code 或 nick_name 登录）
	query := l.svcCtx.DB.Where("code = ? OR nick_name = ?", in.Username, in.Username)
	if in.TenantCode != "" {
		query = query.Where("tenant_code = ?", in.TenantCode)
	}

	if err := query.First(&user).Error; err != nil {
		l.Logger.Errorf("Failed to query user %s: %v", in.Username, err)
		return nil, status.Error(codes.NotFound, "用户不存在")
	}

	return &pb.GetUserAuthInfoResp{
		UserCode:     user.Code,
		PasswordHash: user.Password,
		TenantCode:   user.TenantCode,
		Status:       int32(user.Status),
	}, nil
}
