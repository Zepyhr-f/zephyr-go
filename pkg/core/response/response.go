package response

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zephyr-go/pkg/core/xerr"
)

type Body struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// Success 统一成功返回
func Success(w http.ResponseWriter, data interface{}) {
	httpx.OkJson(w, Body{
		Code: xerr.OK,
		Msg:  "success",
		Data: data,
	})
}

// Error 统一错误返回（也可以结合 httpx.SetErrorHandler 使用）
func Error(w http.ResponseWriter, err error) {
	// 判断如果是系统自定义的错误
	if codeErr, ok := err.(*xerr.CodeError); ok {
		httpx.WriteJson(w, http.StatusOK, Body{
			Code: codeErr.GetErrCode(),
			Msg:  codeErr.GetErrMsg(),
		})
		return
	}
	// 兜底返回服务器异常
	httpx.WriteJson(w, http.StatusOK, Body{
		Code: xerr.SERVER_COMMON_ERROR,
		Msg:  err.Error(),
	})
}
