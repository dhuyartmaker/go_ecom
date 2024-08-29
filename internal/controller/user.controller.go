package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-ecommerce/internal/service"
	"github.com/go-ecommerce/response"
)

type UserController struct {
	UserService service.IUserService
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

func (uc *UserController) GetUser(c *gin.Context) {
	result := uc.UserService.GetUser()
	response.SuccessResponse(c, 20001, []string {result})
}
