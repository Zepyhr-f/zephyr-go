// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package user

import (
	"context"

	"zephyr-go/app/gateway/internal/svc"
	"zephyr-go/app/gateway/internal/types"
	"zephyr-go/app/identity/identityservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserListLogic) UserList(req *types.UserListReq) (resp *types.UserListResp, err error) {
	rpcResp, err := l.svcCtx.IdentityRpc.UserList(l.ctx, &identityservice.UserListReq{
		Username: req.Username,
		RealName: req.RealName,
		Phone:    req.Phone,
		Status:   req.Status,
		DeptCode: req.DeptCode,
	})
	if err != nil {
		return nil, err
	}

	list := make([]types.UserDetail, 0)
	for _, v := range rpcResp.List {
		list = append(list, types.UserDetail{
			Id:         v.Id,
			Code:       v.Code,
			Username:   v.Username,
			RealName:   v.RealName,
			Phone:      v.Phone,
			Email:      v.Email,
			Sex:        v.Sex,
			DeptId:     v.DeptId,
			DeptName:   v.DeptName,
			Status:     v.Status,
			Avatar:     v.Avatar,
			RoleCodes:  v.RoleCodes,
			CreateTime: v.CreateTime,
		})
	}

	return &types.UserListResp{
		List: list,
	}, nil
}
