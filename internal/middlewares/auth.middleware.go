package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-ecommerce/response"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		fmt.Printf("Token: %s", token)
		if token == "" {
			response.ErrorResponse(c, response.ErrorCodeAuthorization, "")
			c.Abort()
			return
		}

		c.Next()
	}
}
