package server

import (
	"context"

	"github.com/gin-gonic/gin"
	"northal.com/config"
	"northal.com/internal/biz"
)

func InitHttp() *gin.Engine {
	config.InitViper()
	engine := gin.Default()
	ctx := context.Background()

	db := InitDatabase()
	redis := InitRedis()
	err := redis.Set(ctx, "key", "123", 0).Err()
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&biz.Users{}, &biz.Role{}, &biz.Permission{})

	var user biz.Users
	db.Find(&user)
	val, _ := redis.Get(ctx, "key").Result()

	api := engine.Group("/api")
	api.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":     "123",
			"data":        user,
			"redis_value": val,
			"hello":       "123",
			"config":      config.GetAppConfig(),
			"database":    config.GetDatabaseConfig(),
		})
	})

	engine.Run(":8080")
	return engine
}
