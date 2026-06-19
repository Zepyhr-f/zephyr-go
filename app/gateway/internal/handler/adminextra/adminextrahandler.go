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
		resp, err := adminextra.MenuList(r.Context(), svcCtx.IdentityRpc)
		if err != nil {
			response.Error(w, err)
			return
		}
		response.Success(w, resp)
	}
}

func MenuDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CodeReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Error(w, err)
			return
		}
		resp, err := adminextra.MenuDetail(r.Context(), svcCtx.IdentityRpc, req.Code)
		if err != nil {
			response.Error(w, err)
			return
		}
		response.Success(w, resp)
	}
}

func MenuSaveHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MenuSaveReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Error(w, err)
			return
		}
		if err := adminextra.MenuSave(r.Context(), svcCtx.IdentityRpc, &req); err != nil {
			response.Error(w, err)
			return
		}
		response.Success(w, nil)
	}
}

func MenuUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MenuUpdateReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Error(w, err)
			return
		}
		if err := adminextra.MenuUpdate(r.Context(), svcCtx.IdentityRpc, &req); err != nil {
			response.Error(w, err)
			return
		}
		response.Success(w, nil)
	}
}

func MenuRemoveHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MenuRemoveReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Error(w, err)
			return
		}
		if err := adminextra.MenuRemove(r.Context(), svcCtx.IdentityRpc, &req); err != nil {
			response.Error(w, err)
			return
		}
		response.Success(w, nil)
	}
}

func MenuStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MenuStatusReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Error(w, err)
			return
		}
		if err := adminextra.MenuStatus(r.Context(), svcCtx.IdentityRpc, &req); err != nil {
			response.Error(w, err)
			return
		}
		response.Success(w, nil)
	}
}

func RoleDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CodeReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Error(w, err)
			return
		}
		resp, err := adminextra.RoleDetail(r.Context(), svcCtx.IdentityRpc, req.Code)
		if err != nil {
			response.Error(w, err)
			return
		}
		response.Success(w, resp)
	}
}

func RoleMenuTreeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CodeReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Error(w, err)
			return
		}
		resp, err := adminextra.RoleMenuTree(r.Context(), svcCtx.IdentityRpc, req.Code)
		if err != nil {
			response.Error(w, err)
			return
		}
		response.Success(w, resp)
	}
}

func RoleAssignMenusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RoleAssignMenusReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Error(w, err)
			return
		}
		if err := adminextra.AssignRoleMenus(r.Context(), svcCtx.IdentityRpc, &req); err != nil {
			response.Error(w, err)
			return
		}
		response.Success(w, nil)
	}
}

func RoleDataScopeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RoleDataScopeReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Error(w, err)
			return
		}
		if err := adminextra.RoleDataScope(r.Context(), svcCtx.IdentityRpc, &req); err != nil {
			response.Error(w, err)
			return
		}
		response.Success(w, nil)
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
