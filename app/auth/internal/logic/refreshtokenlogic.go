package logic

import (
	"context"

	"zephyr-go/app/auth/internal/svc"
	"zephyr-go/app/auth/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshTokenLogic {
	return &RefreshTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 【Token刷新】
func (l *RefreshTokenLogic) RefreshToken(in *pb.RefreshTokenReq) (*pb.LoginVerifyResp, error) {
	// todo: add your logic here and delete this line

	return &pb.LoginVerifyResp{}, nil
}
