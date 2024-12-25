package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"northal.com/internal/pkg/jwt"
	"northal.com/internal/pkg/response"
)

func AuthHandler(getToken func(userID int) (string, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == "" {
			response.Error(c, http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		jwtInstance := jwt.NewJwt()

		userId, err := jwtInstance.ParseToken(token)
		if err != nil {
			response.Error(c, http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		userToken, err := getToken(userId)
		if err != nil || userToken != token {
			response.Error(c, http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		c.Set("user_id", userId)
		c.Next()
	}
}
