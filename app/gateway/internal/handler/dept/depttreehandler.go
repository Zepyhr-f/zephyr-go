// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package dept

import (
	"net/http"

	"zephyr-go/app/gateway/internal/logic/dept"
	"zephyr-go/app/gateway/internal/svc"
	"zephyr-go/pkg/core/response"
)

func DeptTreeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := dept.NewDeptTreeLogic(r.Context(), svcCtx)
		resp, err := l.DeptTree()
		if err != nil {
			response.Error(w, err)
		} else {
			response.Success(w, resp)
		}
	}
}
