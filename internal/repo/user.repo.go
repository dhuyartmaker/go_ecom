package repo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserRepository struct {}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (ur *UserRepository) GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
