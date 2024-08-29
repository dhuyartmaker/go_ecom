package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-ecommerce/internal/wire"
)

type UserRouter struct{}

func (ur *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	UserController, _ := wire.InitUserRouterHandler()

	userPublicRouter := Router.Group("/user")
	{
		userPublicRouter.GET("/", UserController.GetUser)
	}
	userPrivateRouter := Router.Group("/user")
	{
		userPrivateRouter.GET("/info")
	}
}
