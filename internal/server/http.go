package server

import (
	"github.com/gin-gonic/gin"
	"northal.com/internal/model"
)

func InitHttp() *gin.Engine {
	engine := gin.Default()

	db := InitDatabase()
	db.AutoMigrate(&model.Users{})

	var user model.Users
	db.Find(&user)

	api := engine.Group("/api")
	api.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "123",
			"data":    user,
		})
	})

	engine.Run(":8080")
	return engine
}
