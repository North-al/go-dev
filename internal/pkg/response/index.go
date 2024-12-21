package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`           // 状态码
	Message string      `json:"message"`        // 提示信息
	Data    interface{} `json:"data,omitempty"` // 数据内容
}

// Success 返回成功响应
func Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	})
}

func SuccessWithMessage(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: message,
		Data:    nil,
	})
}

// Error 返回错误响应
func Error(ctx *gin.Context, code int, message string) {
	ctx.JSON(code, Response{
		Code:    code,
		Message: message,
		Data:    nil,
	})

}

func ErrorWithMessage(message string) *Response {
	return &Response{
		Code:    http.StatusInternalServerError,
		Message: message,
		Data:    nil,
	}
}
