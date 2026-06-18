// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package menu

import (
	"net/http"

	"zephyr-go/app/gateway/internal/logic/menu"
	"zephyr-go/app/gateway/internal/svc"
	"zephyr-go/pkg/core/response"
)

func MenuTreeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := menu.NewMenuTreeLogic(r.Context(), svcCtx)
		resp, err := l.MenuTree()
		if err != nil {
			response.Error(w, err)
		} else {
			response.Success(w, resp)
		}
	}
}
