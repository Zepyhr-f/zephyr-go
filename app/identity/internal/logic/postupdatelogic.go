package logic

import (
	"context"

	"zephyr-go/app/identity/internal/repository/model"
	"zephyr-go/app/identity/internal/svc"
	"zephyr-go/app/identity/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPostUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostUpdateLogic {
	return &PostUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PostUpdateLogic) PostUpdate(in *pb.PostUpdateReq) (*pb.SuccessResp, error) {
	var post model.SysPost
	if err := l.svcCtx.DB.Where("code = ?", in.Code).First(&post).Error; err != nil {
		l.Errorf("find post failed: %v", err)
		return nil, err
	}
	
	post.PostName = in.PostName
	post.OrderNum = int(in.OrderNum)
	post.Status = int16(in.Status)

	if err := l.svcCtx.DB.Save(&post).Error; err != nil {
		l.Errorf("update post failed: %v", err)
		return nil, err
	}

	return &pb.SuccessResp{Success: true}, nil
}
