package logic

import (
	"context"

	"strconv"
	
	"zephyr-go/app/identity/internal/repository/model"
	"zephyr-go/app/identity/internal/svc"
	"zephyr-go/app/identity/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// User
func (l *UserListLogic) UserList(in *pb.UserListReq) (*pb.UserListResp, error) {
	var users []model.SysUser
	query := l.svcCtx.DB.Model(&model.SysUser{})

	if in.Username != "" {
		query = query.Where("nick_name LIKE ?", "%"+in.Username+"%")
	}
	if in.RealName != "" {
		query = query.Where("real_name LIKE ?", "%"+in.RealName+"%")
	}
	if in.Phone != "" {
		query = query.Where("phone LIKE ?", "%"+in.Phone+"%")
	}
	if in.Status != 0 {
		query = query.Where("status = ?", in.Status)
	}
	if in.DeptCode != "" {
		query = query.Where("dept_code = ?", in.DeptCode)
	}

	if err := query.Preload("Roles").Preload("Dept").Order("id desc").Find(&users).Error; err != nil {
		l.Errorf("find users failed: %v", err)
		return nil, err
	}

	var list []*pb.UserDetail
	for _, u := range users {
		var roleCodes []string
		for _, r := range u.Roles {
			roleCodes = append(roleCodes, r.Code)
		}

		idStr := ""
		if u.Id != 0 {
			idStr = strconv.FormatInt(u.Id, 10)
		}

		deptIdStr := ""
		if u.Dept.Id != 0 {
			deptIdStr = strconv.FormatInt(u.Dept.Id, 10)
		}

		list = append(list, &pb.UserDetail{
			Id:         idStr,
			Code:       u.Code,
			Username:   u.NickName,
			RealName:   u.RealName,
			Phone:      u.Phone,
			Email:      u.Email,
			Sex:        int32(u.Sex),
			DeptId:     deptIdStr,
			DeptName:   u.Dept.DeptName,
			Status:     int32(u.Status),
			Avatar:     u.Avatar,
			RoleCodes:  roleCodes,
			CreateTime: u.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &pb.UserListResp{
		List: list,
	}, nil
}
