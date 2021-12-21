package middleware

import (
	"github.com/gin-gonic/gin"
	"ready-go/app/http/response"
	"ready-go/pkg/logger"
)

// Recover 捕获 panic 并记录错误
func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Errorf("[%s]%s triggered panic! %+v", c.Request.Method, c.Request.RequestURI, err)
				response.Error(c)
				c.Abort()
			}
		}()
		c.Next()
	}
}
