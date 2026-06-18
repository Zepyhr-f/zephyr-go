package logic

import (
	"context"
	"strconv"

	"zephyr-go/app/identity/internal/repository/model"
	"zephyr-go/app/identity/internal/svc"
	"zephyr-go/app/identity/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RolePageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRolePageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RolePageLogic {
	return &RolePageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Role
func (l *RolePageLogic) RolePage(in *pb.RolePageReq) (*pb.RolePageResp, error) {
	var roles []model.SysRole
	var total int64
	
	query := l.svcCtx.DB.Model(&model.SysRole{})
	if in.RoleName != "" {
		query = query.Where("role_name LIKE ?", "%"+in.RoleName+"%")
	}
	if in.Status != 0 {
		query = query.Where("status = ?", in.Status)
	}

	query.Count(&total)

	offset := (in.Current - 1) * in.Size
	if err := query.Order("order_num asc, id desc").Offset(int(offset)).Limit(int(in.Size)).Find(&roles).Error; err != nil {
		l.Errorf("find roles failed: %v", err)
		return nil, err
	}

	var records []*pb.RoleDetail
	for _, r := range roles {
		// Use integer format for Id but store as string since it's int64 in DB and string in protobuf
		idStr := ""
		if r.Id != 0 {
			idStr = strconv.FormatInt(r.Id, 10)
		}
		
		records = append(records, &pb.RoleDetail{
			Id:         idStr,
			Code:       r.Code,
			RoleName:   r.RoleName,
			RoleSort:   int32(r.OrderNum),
			Status:     int32(r.Status),
			CreateTime: r.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	pages := total / int64(in.Size)
	if total%int64(in.Size) != 0 {
		pages++
	}

	return &pb.RolePageResp{
		Total:   total,
		Records: records,
		Size:    in.Size,
		Current: in.Current,
		Pages:   int32(pages),
	}, nil
}
