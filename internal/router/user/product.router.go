package user

import "github.com/gin-gonic/gin"

type ProductRouter struct{}

func (ur *ProductRouter) InitProductRouter(Router *gin.RouterGroup) {
	productPublicRouter := Router.Group("/product")
	{
		productPublicRouter.GET("/")
	}
	productPrivateRouter := Router.Group("/product")
	{
		productPrivateRouter.GET("/info")
	}
}