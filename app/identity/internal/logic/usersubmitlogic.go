package logic

import (
	"context"

	"zephyr-go/app/identity/internal/repository/model"
	"zephyr-go/app/identity/internal/svc"
	"zephyr-go/app/identity/pb"
	
	"golang.org/x/crypto/bcrypt"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserSubmitLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserSubmitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserSubmitLogic {
	return &UserSubmitLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserSubmitLogic) UserSubmit(in *pb.UserSubmitReq) (*pb.SuccessResp, error) {
	if in.Id != "" {
		// update
		var user model.SysUser
		if err := l.svcCtx.DB.Where("code = ?", in.Code).First(&user).Error; err != nil {
			return nil, err
		}
		user.NickName = in.NickName
		user.RealName = in.RealName
		user.Phone = in.Phone
		user.Email = in.Email
		user.DeptCode = in.DeptCode
		user.Sex = int16(in.Sex)
		user.Status = int16(in.Status)
		if err := l.svcCtx.DB.Save(&user).Error; err != nil {
			return nil, err
		}
	} else {
		hash, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
		user := model.SysUser{
			NickName: in.NickName,
			RealName: in.RealName,
			Phone:    in.Phone,
			Email:    in.Email,
			DeptCode: in.DeptCode,
			Sex:      int16(in.Sex),
			Status:   int16(in.Status),
			Password: string(hash),
		}
		user.Code = in.Code
		if err := l.svcCtx.DB.Create(&user).Error; err != nil {
			return nil, err
		}
	}

	return &pb.SuccessResp{Success: true}, nil
}
