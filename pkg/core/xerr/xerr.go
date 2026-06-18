package xerr

const (
	OK                  = 200
	SERVER_COMMON_ERROR = 500
	REUQEST_PARAM_ERROR = 400
	TOKEN_EXPIRE_ERROR  = 401
	DB_ERROR            = 500
)

type CodeError struct {
	errCode int
	errMsg  string
}

func (e *CodeError) GetErrCode() int {
	return e.errCode
}

func (e *CodeError) GetErrMsg() string {
	return e.errMsg
}

func (e *CodeError) Error() string {
	return e.errMsg
}

func NewErrCodeMsg(errCode int, errMsg string) *CodeError {
	return &CodeError{errCode: errCode, errMsg: errMsg}
}

func NewErrCode(errCode int) *CodeError {
	return &CodeError{errCode: errCode, errMsg: "服务器开小差啦，稍后再来试一试"}
}

func NewErrMsg(errMsg string) *CodeError {
	return &CodeError{errCode: SERVER_COMMON_ERROR, errMsg: errMsg}
}
