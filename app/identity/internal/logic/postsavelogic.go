package logic

import (
	"context"

	"zephyr-go/app/identity/internal/repository/model"
	"zephyr-go/app/identity/internal/svc"
	"zephyr-go/app/identity/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostSaveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPostSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostSaveLogic {
	return &PostSaveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PostSaveLogic) PostSave(in *pb.PostSaveReq) (*pb.SuccessResp, error) {
	post := model.SysPost{
		PostName: in.PostName,
		OrderNum: int(in.OrderNum),
		Status:   int16(in.Status),
	}
	post.Code = in.Code
	if err := l.svcCtx.DB.Create(&post).Error; err != nil {
		l.Errorf("save post failed: %v", err)
		return nil, err
	}

	return &pb.SuccessResp{Success: true}, nil
}
