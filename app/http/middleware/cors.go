package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		// 允许的请求来源
		AllowOrigins: []string{"localhost", "http://127.0.0.1"},
		// 允许的请求方法
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		// 请求头应当携带的字段
		AllowHeaders: []string{"Origin", "Content-Length", "Content-Type"},
		// 是否允许携带 Cookie 等身份信息
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	})
}
