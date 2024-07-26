package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-ecommerce/internal/controller"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	
	v1 := r.Group("/v1")
	{
		v1.GET("/user", controller.NewUserController().GetUser)
	}
	return r
}
