package logic

import (
	"strconv"

	"zephyr-go/app/identity/internal/repository/model"
	"zephyr-go/app/identity/pb"
)

func menuToPB(menu model.SysMenu) *pb.MenuDetail {
	idStr := ""
	if menu.Id != 0 {
		idStr = strconv.FormatInt(menu.Id, 10)
	}
	createTime := ""
	if !menu.CreateTime.IsZero() {
		createTime = menu.CreateTime.Format("2006-01-02 15:04:05")
	}
	return &pb.MenuDetail{
		MenuCode:   menu.Code,
		ParentCode: menu.ParentCode,
		MenuName:   menu.MenuName,
		Icon:       menu.Icon,
		MenuType:   menu.MenuType,
		Perms:      menu.Perms,
		Path:       menu.Path,
		Component:  menu.Component,
		OrderNum:   int32(menu.OrderNum),
		Id:         idStr,
		Status:     int32(menu.Status),
		CreateTime: createTime,
	}
}

func buildMenuTree(menus []model.SysMenu, parentCode string) []*pb.MenuDetail {
	var list []*pb.MenuDetail
	for _, menu := range menus {
		if menu.ParentCode == parentCode || (parentCode == "-1" && menu.ParentCode == "") {
			detail := menuToPB(menu)
			detail.Children = buildMenuTree(menus, menu.Code)
			list = append(list, detail)
		}
	}
	return list
}

func roleToPB(role model.SysRole) *pb.RoleDetail {
	idStr := ""
	if role.Id != 0 {
		idStr = strconv.FormatInt(role.Id, 10)
	}
	createTime := ""
	if !role.CreateTime.IsZero() {
		createTime = role.CreateTime.Format("2006-01-02 15:04:05")
	}
	return &pb.RoleDetail{
		Id:         idStr,
		Code:       role.Code,
		RoleName:   role.RoleName,
		RoleSort:   int32(role.OrderNum),
		Status:     int32(role.Status),
		CreateTime: createTime,
	}
}
