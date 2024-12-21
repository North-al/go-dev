package main

import (
	"github.com/gin-gonic/gin"
	"northal.com/api"
	"northal.com/config"
	_ "northal.com/docs"
	"northal.com/internal/server"
)

//	@title			权限管理系统
//	@version		1.0
//	@description	权限管理系统swagger api介绍
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	North
//	@contact.email	1227379879@qq.com

//	@host		localhost:8080
//	@BasePath	/

//	@securityDefinitions.basic	BasicAuth

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	config.InitViper()
	router := server.InitHttp()
	db := server.InitDatabase()
	redis := server.InitRedis()

	// 加载所有api
	api.NewSetupApi(router, db, redis).SetupRoutes()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 加载swagger
	server.InitDocs(router)

	router.Run(":" + config.GetAppConfig().Port)
}
