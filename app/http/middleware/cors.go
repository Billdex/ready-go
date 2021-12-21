package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	conf := cors.DefaultConfig()
	conf.AllowOrigins = []string{}
	return cors.New(conf)
}
