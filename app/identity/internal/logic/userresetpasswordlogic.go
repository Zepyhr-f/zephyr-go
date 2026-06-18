package logic

import (
	"context"

	"zephyr-go/app/identity/internal/repository/model"
	"zephyr-go/app/identity/internal/svc"
	"zephyr-go/app/identity/pb"
	
	"golang.org/x/crypto/bcrypt"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserResetPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserResetPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserResetPasswordLogic {
	return &UserResetPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserResetPasswordLogic) UserResetPassword(in *pb.UserResetPasswordReq) (*pb.SuccessResp, error) {
	hash, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	if err := l.svcCtx.DB.Model(&model.SysUser{}).Where("id = ?", in.Id).Update("password", string(hash)).Error; err != nil {
		l.Errorf("reset user password failed: %v", err)
		return nil, err
	}

	return &pb.SuccessResp{Success: true}, nil
}
