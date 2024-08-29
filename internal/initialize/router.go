package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/go-ecommerce/global"
	"github.com/go-ecommerce/internal/router"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	userRouter := router.RouterApp.UserRouter
	MainGroupRouter := r.Group("v1")
	{
		MainGroupRouter.GET("ping")
	}
	{
		userRouter.InitProductRouter(MainGroupRouter)
		userRouter.InitUserRouter(MainGroupRouter)
	}

	return r
}
