package service

import (
	"github.com/gin-gonic/gin"
)

type UserService struct {}

func NewUserService() *UserService {
	return &UserService{}
}

func (ur *UserService) GetUser(c *gin.Context) string {
	return "Huy ngu"
}
