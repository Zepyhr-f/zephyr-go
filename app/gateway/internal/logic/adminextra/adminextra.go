package adminextra

import (
	"context"
	"runtime"
	"time"

	"zephyr-go/app/gateway/internal/types"
)

func MenuList(ctx context.Context) *types.MenuListResp {
	items := flattenMenus(seedMenus())
	return &types.MenuListResp{Records: items, Total: int64(len(items)), Current: 1, Size: len(items)}
}

func MenuDetail(ctx context.Context, code string) (*types.MenuDetailResp, bool) {
	for _, item := range flattenMenus(seedMenus()) {
		if item.MenuCode == code {
			return &types.MenuDetailResp{Menu: item}, true
		}
	}
	return nil, false
}

func RoleDetail(ctx context.Context, code string) *types.RoleDetailExtraResp {
	roleCode := code
	if roleCode == "" {
		roleCode = "admin"
	}
	checked := []string{"dashboard", "system", "user", "user_query"}
	if roleCode == "admin" {
		checked = allMenuCodes()
	}
	return &types.RoleDetailExtraResp{
		Role: types.RoleDetail{Code: roleCode, RoleName: roleCode, Status: 1},
		MenuCodes: checked,
	}
}

func RoleMenuTree(ctx context.Context, code string) *types.RoleMenuTreeResp {
	return &types.RoleMenuTreeResp{
		Menus: seedMenus(),
		CheckedKeys: RoleDetail(ctx, code).MenuCodes,
	}
}

func AssignRoleMenus(ctx context.Context, req *types.RoleAssignMenusReq) *types.OperationResp {
	return &types.OperationResp{Success: true, Message: "角色菜单授权接口已接收请求，当前最小实现不执行危险写入"}
}

func RoleDataScope(ctx context.Context, req *types.RoleDataScopeReq) *types.OperationResp {
	return &types.OperationResp{Success: true, Message: "数据权限能力尚未启用，当前为安全占位实现"}
}

func PlaceholderList(module string) *types.AdminListResp {
	now := time.Now().Format(time.RFC3339)
	return &types.AdminListResp{
		Total: 1,
		Current: 1,
		Size: 10,
		Records: []map[string]any{{
			"module": module,
			"status": "placeholder",
			"message": "接口已按菜单文档预留，后续接入真实数据源",
			"updatedAt": now,
		}},
	}
}

func ServerMetrics() map[string]any {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return map[string]any{
		"runtime": "go",
		"goVersion": runtime.Version(),
		"goroutines": runtime.NumGoroutine(),
		"memory": map[string]any{
			"allocBytes": m.Alloc,
			"sysBytes": m.Sys,
		},
		"status": "ok",
	}
}

func CacheMetrics() map[string]any {
	return map[string]any{
		"status": "placeholder",
		"message": "缓存监控仅返回安全摘要，不暴露 Redis 连接信息",
	}
}

func DatasourceMetrics() map[string]any {
	return map[string]any{
		"status": "placeholder",
		"message": "数据源监控仅返回安全摘要，不暴露数据库连接串",
	}
}

func DevtoolStatus(tool string) map[string]any {
	enabled := false
	message := "该开发工具当前为安全占位，默认禁用危险操作"
	if tool == "api-doc" {
		message = "接口文档入口已预留，待接入 OpenAPI 聚合"
	}
	return map[string]any{"tool": tool, "enabled": enabled, "message": message}
}

func seedMenus() []types.MenuDetail {
	return []types.MenuDetail{
		{MenuCode: "dashboard", ParentCode: "-1", MenuName: "概览", MenuType: "C", Perms: "sys:dashboard", Path: "/", Component: "dashboard/Overview", OrderNum: 0},
		{MenuCode: "system", ParentCode: "-1", MenuName: "系统管理", MenuType: "M", Path: "/system", OrderNum: 1, Children: []types.MenuDetail{
			{MenuCode: "user", ParentCode: "system", MenuName: "用户管理", MenuType: "C", Perms: "sys:user:list", Path: "/system/users", Component: "system/UserManagement", OrderNum: 1, Children: []types.MenuDetail{
				{MenuCode: "user_query", ParentCode: "user", MenuName: "用户查询", MenuType: "F", Perms: "sys:user:query", OrderNum: 1},
				{MenuCode: "user_add", ParentCode: "user", MenuName: "用户新增", MenuType: "F", Perms: "sys:user:add", OrderNum: 2},
				{MenuCode: "user_edit", ParentCode: "user", MenuName: "用户修改", MenuType: "F", Perms: "sys:user:edit", OrderNum: 3},
				{MenuCode: "user_remove", ParentCode: "user", MenuName: "用户删除", MenuType: "F", Perms: "sys:user:remove", OrderNum: 4},
			}},
			{MenuCode: "dept", ParentCode: "system", MenuName: "部门管理", MenuType: "C", Perms: "sys:dept:list", Path: "/system/depts", Component: "system/DepartmentManagement", OrderNum: 2},
			{MenuCode: "post", ParentCode: "system", MenuName: "岗位管理", MenuType: "C", Perms: "sys:post:list", Path: "/system/posts", Component: "system/PostManagement", OrderNum: 3},
			{MenuCode: "menu", ParentCode: "system", MenuName: "菜单管理", MenuType: "C", Perms: "sys:menu:list", Path: "/system/menus", Component: "system/MenuManagement", OrderNum: 4},
			{MenuCode: "role", ParentCode: "system", MenuName: "角色管理", MenuType: "C", Perms: "sys:role:list", Path: "/system/roles", Component: "system/RoleManagement", OrderNum: 5},
		}},
		{MenuCode: "security", ParentCode: "-1", MenuName: "安全审计", MenuType: "M", Path: "/security", OrderNum: 2, Children: []types.MenuDetail{
			{MenuCode: "login_log", ParentCode: "security", MenuName: "登录日志", MenuType: "C", Perms: "sys:loginlog:list", Path: "/security/login-log", Component: "security/LoginLog", OrderNum: 1},
			{MenuCode: "oper_log", ParentCode: "security", MenuName: "操作日志", MenuType: "C", Perms: "sys:operlog:list", Path: "/security/op-log", Component: "security/OperationLog", OrderNum: 2},
			{MenuCode: "online", ParentCode: "security", MenuName: "在线用户", MenuType: "C", Perms: "sys:online:list", Path: "/security/online", Component: "security/OnlineUsers", OrderNum: 3},
		}},
		{MenuCode: "monitor", ParentCode: "-1", MenuName: "系统监控", MenuType: "M", Path: "/monitor", OrderNum: 3, Children: []types.MenuDetail{
			{MenuCode: "server", ParentCode: "monitor", MenuName: "服务监控", MenuType: "C", Perms: "sys:server:list", Path: "/monitor/server", Component: "monitor/ServiceMonitoring", OrderNum: 1},
			{MenuCode: "cache", ParentCode: "monitor", MenuName: "缓存监控", MenuType: "C", Perms: "sys:cache:list", Path: "/monitor/cache", Component: "monitor/CacheMonitoring", OrderNum: 2},
			{MenuCode: "datasource", ParentCode: "monitor", MenuName: "数据源监控", MenuType: "C", Perms: "sys:db:list", Path: "/monitor/datasource", Component: "monitor/DataSourceMonitoring", OrderNum: 3},
			{MenuCode: "cron", ParentCode: "monitor", MenuName: "任务调度", MenuType: "C", Perms: "sys:job:list", Path: "/monitor/cron", Component: "monitor/CronJobs", OrderNum: 4},
		}},
		{MenuCode: "infra", ParentCode: "-1", MenuName: "基础设施", MenuType: "M", Path: "/infrastructure", OrderNum: 4, Children: []types.MenuDetail{
			{MenuCode: "dict", ParentCode: "infra", MenuName: "字典管理", MenuType: "C", Perms: "sys:dict:list", Path: "/infrastructure/dict", Component: "infrastructure/Dictionary", OrderNum: 1},
			{MenuCode: "param", ParentCode: "infra", MenuName: "参数配置", MenuType: "C", Perms: "sys:config:list", Path: "/infrastructure/params", Component: "infrastructure/Params", OrderNum: 2},
			{MenuCode: "file", ParentCode: "infra", MenuName: "文件管理", MenuType: "C", Perms: "sys:file:list", Path: "/infrastructure/files", Component: "infrastructure/FileCenter", OrderNum: 3},
			{MenuCode: "notice", ParentCode: "infra", MenuName: "通知公告", MenuType: "C", Perms: "sys:notice:list", Path: "/infrastructure/notices", Component: "infrastructure/Notices", OrderNum: 4},
		}},
		{MenuCode: "devtools", ParentCode: "-1", MenuName: "开发工具", MenuType: "M", Path: "/devtools", OrderNum: 5, Children: []types.MenuDetail{
			{MenuCode: "codegen", ParentCode: "devtools", MenuName: "代码生成", MenuType: "C", Perms: "dev:codegen:list", Path: "/devtools/codegen", Component: "devtools/Codegen", OrderNum: 1},
			{MenuCode: "api_doc", ParentCode: "devtools", MenuName: "接口文档", MenuType: "C", Perms: "dev:api:list", Path: "/devtools/api-doc", Component: "devtools/ApiDoc", OrderNum: 2},
			{MenuCode: "sql_terminal", ParentCode: "devtools", MenuName: "SQL 终端", MenuType: "C", Perms: "dev:sql:list", Path: "/devtools/sql", Component: "devtools/SqlTerminal", OrderNum: 3},
		}},
	}
}

func flattenMenus(items []types.MenuDetail) []types.MenuDetail {
	out := make([]types.MenuDetail, 0)
	var walk func([]types.MenuDetail)
	walk = func(nodes []types.MenuDetail) {
		for _, n := range nodes {
			copyNode := n
			copyNode.Children = nil
			out = append(out, copyNode)
			if len(n.Children) > 0 {
				walk(n.Children)
			}
		}
	}
	walk(items)
	return out
}

func allMenuCodes() []string {
	items := flattenMenus(seedMenus())
	codes := make([]string, 0, len(items))
	for _, item := range items {
		codes = append(codes, item.MenuCode)
	}
	return codes
}
