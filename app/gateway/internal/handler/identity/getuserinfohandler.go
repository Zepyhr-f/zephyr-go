// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package identity

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zephyr-go/app/gateway/internal/logic/identity"
	"zephyr-go/app/gateway/internal/svc"
	"zephyr-go/app/gateway/internal/types"
	"zephyr-go/pkg/core/response"
)

// 获取当前登录用户信息
func GetUserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Error(w, err)
			return
		}

		l := identity.NewGetUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetUserInfo(&req)
		if err != nil {
			response.Error(w, err)
		} else {
			response.Success(w, resp)
		}
	}
}
