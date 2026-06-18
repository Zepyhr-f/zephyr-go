package logic

import (
	"context"

	"zephyr-go/app/identity/internal/repository/model"
	"zephyr-go/app/identity/internal/svc"
	"zephyr-go/app/identity/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPostPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostPageLogic {
	return &PostPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Post
func (l *PostPageLogic) PostPage(in *pb.PostPageReq) (*pb.PostPageResp, error) {
	var posts []model.SysPost
	var total int64
	
	query := l.svcCtx.DB.Model(&model.SysPost{})
	if in.PostName != "" {
		query = query.Where("post_name LIKE ?", "%"+in.PostName+"%")
	}
	if in.Status != 0 {
		query = query.Where("status = ?", in.Status)
	}

	query.Count(&total)

	offset := (in.Current - 1) * in.Size
	if err := query.Order("order_num asc, id desc").Offset(int(offset)).Limit(int(in.Size)).Find(&posts).Error; err != nil {
		l.Errorf("find posts failed: %v", err)
		return nil, err
	}

	var records []*pb.PostDetail
	for _, p := range posts {
		records = append(records, &pb.PostDetail{
			Id:         p.Id,
			Code:       p.Code,
			PostName:   p.PostName,
			OrderNum:   int32(p.OrderNum),
			Status:     int32(p.Status),
			CreateTime: p.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	pages := total / int64(in.Size)
	if total%int64(in.Size) != 0 {
		pages++
	}

	return &pb.PostPageResp{
		Total:   total,
		Records: records,
		Size:    in.Size,
		Current: in.Current,
		Pages:   int32(pages),
	}, nil
}
