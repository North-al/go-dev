package server

import (
	_ "northal.com/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitDocs(router *gin.Engine) {
	// 我想访问swagger.json文件
	router.GET("/swagger/swagger.json", func(c *gin.Context) {
		c.File("docs/swagger.json")
	})

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
