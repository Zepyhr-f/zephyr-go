package logic

import (
	"context"

	"zephyr-go/app/identity/internal/repository/model"
	"zephyr-go/app/identity/internal/svc"
	"zephyr-go/app/identity/pb"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeptTreeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeptTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptTreeLogic {
	return &DeptTreeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Dept
func (l *DeptTreeLogic) DeptTree(in *pb.EmptyReq) (*pb.DeptTreeResp, error) {
	var depts []model.SysDept
	if err := l.svcCtx.DB.Find(&depts).Error; err != nil {
		l.Errorf("fetch depts failed: %v", err)
		return nil, err
	}

	// build map
	deptMap := make(map[string]*pb.DeptDetail)
	var rootNodes []*pb.DeptDetail

	for _, d := range depts {
		deptMap[d.Code] = &pb.DeptDetail{
			Id:         fmt.Sprintf("%d", d.Id),
			Code:       d.Code,
			ParentCode: d.ParentCode,
			Leaf:       int32(d.Leaf),
			DeptName:   d.DeptName,
			OrderNum:   int32(d.OrderNum),
			Status:     int32(d.Status),
			CreateTime: d.CreateTime.Format("2006-01-02 15:04:05"),
			Children:   []*pb.DeptDetail{},
		}
	}

	for _, d := range depts {
		node := deptMap[d.Code]
		if d.ParentCode == "" || d.ParentCode == "0" {
			rootNodes = append(rootNodes, node)
		} else {
			if parent, ok := deptMap[d.ParentCode]; ok {
				parent.Children = append(parent.Children, node)
			} else {
				// parent not found, treat as root
				rootNodes = append(rootNodes, node)
			}
		}
	}

	return &pb.DeptTreeResp{
		List: rootNodes,
	}, nil
}
