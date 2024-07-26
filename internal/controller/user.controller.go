package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-ecommerce/internal/service"
	"github.com/go-ecommerce/response"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (uc *UserController) GetUser(c *gin.Context) {
	response.SuccessResponse(c, 20001, []string {"huy","khung","met","moi"})
}
