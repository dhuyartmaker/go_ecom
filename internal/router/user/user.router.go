package user

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (ur *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userPublicRouter := Router.Group("/user")
	{
		userPublicRouter.GET("/")
	}
	userPrivateRouter := Router.Group("/user")
	{
		userPrivateRouter.GET("/info")
	}
}