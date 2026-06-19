package adminextra

import (
	"context"
	"runtime"

	"zephyr-go/app/gateway/internal/types"
	"zephyr-go/app/identity/identityservice"
)

func MenuList(ctx context.Context, rpc identityservice.IdentityService) (*types.MenuListResp, error) {
	rpcResp, err := rpc.MenuList(ctx, &identityservice.EmptyReq{})
	if err != nil {
		return nil, err
	}
	items := make([]types.MenuDetail, 0, len(rpcResp.Records))
	for _, item := range rpcResp.Records {
		items = append(items, mapMenuToTypes(item))
	}
	return &types.MenuListResp{Records: items, Total: rpcResp.Total, Current: int(rpcResp.Current), Size: int(rpcResp.Size)}, nil
}

func MenuDetail(ctx context.Context, rpc identityservice.IdentityService, code string) (*types.MenuDetailResp, error) {
	rpcResp, err := rpc.MenuDetail(ctx, &identityservice.MenuDetailReq{Code: code})
	if err != nil {
		return nil, err
	}
	return &types.MenuDetailResp{Menu: mapMenuToTypes(rpcResp.Menu)}, nil
}

func MenuSave(ctx context.Context, rpc identityservice.IdentityService, req *types.MenuSaveReq) error {
	_, err := rpc.MenuSave(ctx, &identityservice.MenuSaveReq{
		Code: req.Code, ParentCode: req.ParentCode, MenuName: req.MenuName, Icon: req.Icon,
		MenuType: req.MenuType, Perms: req.Perms, Path: req.Path, Component: req.Component,
		OrderNum: req.OrderNum, Status: req.Status,
	})
	return err
}

func MenuUpdate(ctx context.Context, rpc identityservice.IdentityService, req *types.MenuUpdateReq) error {
	_, err := rpc.MenuUpdate(ctx, &identityservice.MenuUpdateReq{
		Id: req.Id, Code: req.Code, ParentCode: req.ParentCode, MenuName: req.MenuName, Icon: req.Icon,
		MenuType: req.MenuType, Perms: req.Perms, Path: req.Path, Component: req.Component,
		OrderNum: req.OrderNum, Status: req.Status,
	})
	return err
}

func MenuRemove(ctx context.Context, rpc identityservice.IdentityService, req *types.MenuRemoveReq) error {
	_, err := rpc.MenuRemove(ctx, &identityservice.MenuRemoveReq{Codes: req.Codes})
	return err
}

func MenuStatus(ctx context.Context, rpc identityservice.IdentityService, req *types.MenuStatusReq) error {
	_, err := rpc.MenuStatus(ctx, &identityservice.MenuStatusReq{Code: req.Code, Status: req.Status})
	return err
}

func RoleDetail(ctx context.Context, rpc identityservice.IdentityService, code string) (*types.RoleDetailExtraResp, error) {
	rpcResp, err := rpc.RoleDetail(ctx, &identityservice.RoleDetailReq{Code: code})
	if err != nil {
		return nil, err
	}
	return &types.RoleDetailExtraResp{
		Role:      mapRoleToTypes(rpcResp.Role),
		MenuCodes: rpcResp.MenuCodes,
	}, nil
}

func RoleMenuTree(ctx context.Context, rpc identityservice.IdentityService, code string) (*types.RoleMenuTreeResp, error) {
	rpcResp, err := rpc.RoleMenuTree(ctx, &identityservice.RoleMenuTreeReq{RoleCode: code})
	if err != nil {
		return nil, err
	}
	menus := make([]types.MenuDetail, 0, len(rpcResp.Menus))
	for _, item := range rpcResp.Menus {
		menus = append(menus, mapMenuToTypes(item))
	}
	return &types.RoleMenuTreeResp{Menus: menus, CheckedKeys: rpcResp.CheckedKeys}, nil
}

func AssignRoleMenus(ctx context.Context, rpc identityservice.IdentityService, req *types.RoleAssignMenusReq) error {
	_, err := rpc.RoleAssignMenus(ctx, &identityservice.RoleAssignMenusReq{RoleCode: req.RoleCode, MenuCodes: req.MenuCodes})
	return err
}

func RoleDataScope(ctx context.Context, rpc identityservice.IdentityService, req *types.RoleDataScopeReq) error {
	_, err := rpc.RoleDataScope(ctx, &identityservice.RoleDataScopeReq{RoleCode: req.RoleCode, DataScope: req.DataScope, DeptCodes: req.DeptCodes})
	return err
}

func ServerMetrics() map[string]any {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return map[string]any{
		"runtime":    "go",
		"goVersion":  runtime.Version(),
		"goroutines": runtime.NumGoroutine(),
		"memory": map[string]any{
			"allocBytes": m.Alloc,
			"sysBytes":   m.Sys,
		},
		"status": "ok",
	}
}

func CacheMetrics() map[string]any {
	return map[string]any{
		"status":  "safe-summary",
		"message": "缓存监控仅返回安全摘要，不暴露 Redis 连接信息",
	}
}

func DatasourceMetrics() map[string]any {
	return map[string]any{
		"status":  "safe-summary",
		"message": "数据源监控仅返回安全摘要，不暴露数据库连接串",
	}
}

func DevtoolStatus(tool string) map[string]any {
	enabled := false
	message := "该开发工具默认禁用危险操作"
	if tool == "api-doc" {
		message = "接口文档入口已预留，待接入 OpenAPI 聚合"
	}
	return map[string]any{"tool": tool, "enabled": enabled, "message": message}
}

func PlaceholderList(module string) *types.AdminListResp {
	return &types.AdminListResp{
		Total:   0,
		Current: 1,
		Size:    10,
		Records: []map[string]any{},
	}
}

func mapMenuToTypes(v *identityservice.MenuDetail) types.MenuDetail {
	if v == nil {
		return types.MenuDetail{}
	}
	m := types.MenuDetail{
		Id:         v.Id,
		MenuCode:   v.MenuCode,
		ParentCode: v.ParentCode,
		MenuName:   v.MenuName,
		Icon:       v.Icon,
		MenuType:   v.MenuType,
		Perms:      v.Perms,
		Path:       v.Path,
		Component:  v.Component,
		OrderNum:   v.OrderNum,
		Status:     v.Status,
		CreateTime: v.CreateTime,
	}
	for _, child := range v.Children {
		m.Children = append(m.Children, mapMenuToTypes(child))
	}
	return m
}

func mapRoleToTypes(v *identityservice.RoleDetail) types.RoleDetail {
	if v == nil {
		return types.RoleDetail{}
	}
	return types.RoleDetail{
		Id:         v.Id,
		Code:       v.Code,
		RoleName:   v.RoleName,
		RoleSort:   v.RoleSort,
		Status:     v.Status,
		CreateTime: v.CreateTime,
	}
}
