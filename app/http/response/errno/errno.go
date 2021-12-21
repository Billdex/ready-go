package errno

// 基础状态码
const (
	StatusOK          = 0
	StatusBadRequest  = 400
	StatusServerError = 500
)

// 校验与检查相关的错误码
const (
	ErrorCheckToken   = 1001
	ErrorUnauthorized = 1002
	ErrorTokenExpired = 1003

	ErrorValidateParamsFail = 2001
)

// 业务相关的错误码
const ()

var statusMsg = map[int]string{
	// 基础状态码信息
	StatusOK:          "ok",
	StatusBadRequest:  "fail",
	StatusServerError: "服务器内部错误",

	// 请求相关的错误信息
	ErrorCheckToken:         "权限信息验证失败, 请提供有效的 Token",
	ErrorUnauthorized:       "权限不足",
	ErrorTokenExpired:       "身份验证信息过期",
	ErrorValidateParamsFail: "请求参数有误",

	// 业务相关的错误信息
}

func GetMsg(code int) string {
	if msg, ok := statusMsg[code]; ok {
		return msg
	}
	return statusMsg[StatusBadRequest]
}
