package server

import (
	"github.com/gin-gonic/gin"
	"northal.com/internal/middleware"
)

func InitHttp() *gin.Engine {
	engine := gin.Default()
	engine.Use(middleware.ErrorHandler())
	return engine
}
