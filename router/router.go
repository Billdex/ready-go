package router

import (
	"github.com/gin-gonic/gin"
	"ready-go/app/http/middleware"
	"ready-go/config"
)

func InitRouter() *gin.Engine {
	gin.SetMode(config.GetConfig().App.RunMode)
	g := gin.New()
	g.Use(middleware.Recover(), gin.Logger(), middleware.Cors())

	return g
}
