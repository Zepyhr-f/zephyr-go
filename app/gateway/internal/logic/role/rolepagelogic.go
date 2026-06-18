// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package role

import (
	"context"

	"zephyr-go/app/gateway/internal/svc"
	"zephyr-go/app/gateway/internal/types"
	"zephyr-go/app/identity/identityservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type RolePageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRolePageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RolePageLogic {
	return &RolePageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RolePageLogic) RolePage(req *types.RolePageReq) (resp *types.RolePageResp, err error) {
	rpcResp, err := l.svcCtx.IdentityRpc.RolePage(l.ctx, &identityservice.RolePageReq{
		Current:  int32(req.Current),
		Size:     int32(req.Size),
		RoleName: req.RoleName,
		Status:   req.Status,
	})
	if err != nil {
		return nil, err
	}

	var records []types.RoleDetail
	for _, v := range rpcResp.Records {
		records = append(records, types.RoleDetail{
			Id:         v.Id,
			Code:       v.Code,
			RoleName:   v.RoleName,
			RoleSort:   v.RoleSort,
			Status:     v.Status,
			CreateTime: v.CreateTime,
		})
	}

	return &types.RolePageResp{
		Total:   rpcResp.Total,
		Records: records,
		Size:    int(rpcResp.Size),
		Current: int(rpcResp.Current),
		Pages:   int(rpcResp.Pages),
	}, nil
}
