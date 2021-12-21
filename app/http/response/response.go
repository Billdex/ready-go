package response

import (
	"net/http"
	"ready-go/app/http/response/errno"

	"github.com/gin-gonic/gin"
)

// 封装常见请求返回情况

// JsonResponse json 格式返回数据
func JsonResponse(c *gin.Context, code int, data interface{}) {
	response := gin.H{
		"code":    code,
		"message": errno.GetMsg(code),
	}
	if data != nil {
		response["data"] = data
	}
	c.JSON(http.StatusOK, response)
}

// Success 成功返回正常数据
func Success(c *gin.Context, data interface{}) {
	JsonResponse(c, errno.StatusOK, data)
}

// Fail 业务逻辑失败
func Fail(c *gin.Context, code int, data interface{}) {
	JsonResponse(c, code, data)
}

// Error 服务内部错误
func Error(c *gin.Context) {
	JsonResponse(c, errno.StatusServerError, nil)
}

// ErrorCheckToken 鉴权失败
func ErrorCheckToken(c *gin.Context) {
	JsonResponse(c, errno.ErrorCheckToken, nil)
}

// ErrorUnauthorized 用户没有权限
func ErrorUnauthorized(c *gin.Context) {
	JsonResponse(c, errno.ErrorUnauthorized, nil)
}

// ErrorTokenExpired 鉴权信息过期
func ErrorTokenExpired(c *gin.Context) {
	JsonResponse(c, errno.ErrorTokenExpired, nil)
}

// ErrorValidateParamsFail 请求参数有误
func ErrorValidateParamsFail(c *gin.Context) {
	JsonResponse(c, errno.ErrorValidateParamsFail, nil)
}
