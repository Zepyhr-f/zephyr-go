package adminextra

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zephyr-go/app/gateway/internal/logic/adminextra"
	"zephyr-go/app/gateway/internal/svc"
	"zephyr-go/app/gateway/internal/types"
	"zephyr-go/pkg/core/response"
)

func MenuListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response.Success(w, adminextra.MenuList(r.Context()))
	}
}

func MenuDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CodeReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Error(w, err)
			return
		}
		resp, ok := adminextra.MenuDetail(r.Context(), req.Code)
		if !ok {
			response.Success(w, &types.OperationResp{Success: false, Message: "菜单不存在"})
			return
		}
		response.Success(w, resp)
	}
}

func MenuMutationHandler(message string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response.Success(w, &types.OperationResp{Success: true, Message: message})
	}
}

func RoleDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CodeReq
		_ = httpx.Parse(r, &req)
		response.Success(w, adminextra.RoleDetail(r.Context(), req.Code))
	}
}

func RoleMenuTreeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CodeReq
		_ = httpx.Parse(r, &req)
		response.Success(w, adminextra.RoleMenuTree(r.Context(), req.Code))
	}
}

func RoleAssignMenusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RoleAssignMenusReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Error(w, err)
			return
		}
		response.Success(w, adminextra.AssignRoleMenus(r.Context(), &req))
	}
}

func RoleDataScopeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RoleDataScopeReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Error(w, err)
			return
		}
		response.Success(w, adminextra.RoleDataScope(r.Context(), &req))
	}
}

func PlaceholderListHandler(module string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response.Success(w, adminextra.PlaceholderList(module))
	}
}

func ServerMetricsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response.Success(w, adminextra.ServerMetrics())
	}
}

func CacheMetricsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response.Success(w, adminextra.CacheMetrics())
	}
}

func DatasourceMetricsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response.Success(w, adminextra.DatasourceMetrics())
	}
}

func DevtoolStatusHandler(tool string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response.Success(w, adminextra.DevtoolStatus(tool))
	}
}
