package middleware

import (
	"net/http"

	"northal.com/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 捕获 panic 错误
				response.Error(ctx, http.StatusInternalServerError, err.(string))
			}
		}()
		ctx.Next()

		// 检查是否有错误
		if len(ctx.Errors) > 0 {
			response.Error(ctx, http.StatusInternalServerError, ctx.Errors.Last().Error())
		}
	}
}
