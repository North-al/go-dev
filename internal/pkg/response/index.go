package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 通用响应结构
// @Description 接口统一返回格式
type Response struct {
	Code    int         `json:"code" example:"200"`        // 状态码 (200-成功, 500-失败)
	Message string      `json:"message" example:"success"` // 提示信息
	Data    interface{} `json:"data,omitempty"`            // 数据内容
}

// Success 返回成功响应
func Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	})
}

func SuccessWithMessage(ctx *gin.Context, data interface{}, message string) {
	ctx.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: message,
		Data:    data,
	})
}

func SuccessWithCodeAndMessage(ctx *gin.Context, code int, data interface{}, message string) {
	ctx.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

// Error 返回错误响应
func Error(ctx *gin.Context, code int, message string) {
	ctx.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    nil,
	})

}

func ErrorWithMessage(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusOK, Response{
		Code:    http.StatusInternalServerError,
		Message: message,
		Data:    nil,
	})
}
