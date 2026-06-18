package logic

import (
	"context"

	"zephyr-go/app/identity/internal/repository/model"
	"zephyr-go/app/identity/internal/svc"
	"zephyr-go/app/identity/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPostStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostStatusLogic {
	return &PostStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PostStatusLogic) PostStatus(in *pb.PostStatusReq) (*pb.SuccessResp, error) {
	if err := l.svcCtx.DB.Model(&model.SysPost{}).Where("code = ?", in.Code).Update("status", in.Status).Error; err != nil {
		l.Errorf("update post status failed: %v", err)
		return nil, err
	}

	return &pb.SuccessResp{Success: true}, nil
}
