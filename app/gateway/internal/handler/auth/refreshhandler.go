// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package auth

import (
	"net/http"

	"zephyr-go/app/gateway/internal/logic/auth"
	"zephyr-go/app/gateway/internal/svc"
	"zephyr-go/pkg/core/response"
)

// 刷新Token
func RefreshHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := auth.NewRefreshLogic(r.Context(), svcCtx)
		resp, err := l.Refresh()
		if err != nil {
			response.Error(w, err)
		} else {
			response.Success(w, resp)
		}
	}
}
