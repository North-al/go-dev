package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"northal.com/internal/pkg/jwt"
	"northal.com/internal/pkg/response"
)

func AuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == "" {
			response.Error(c, http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		jwtInstance := jwt.NewJwt()

		claims, err := jwtInstance.ParseToken(token)
		fmt.Println(claims)
		if err != nil {
			response.Error(c, http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		c.Set("user_id", claims)
		c.Next()
	}
}
