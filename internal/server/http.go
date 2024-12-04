package server

import (
	"context"

	"github.com/gin-gonic/gin"
	"northal.com/internal/model"
)

func InitHttp() *gin.Engine {
	engine := gin.Default()
	ctx := context.Background()

	db := InitDatabase()
	redis := InitRedis()
	err := redis.Set(ctx, "key", "123", 0).Err()
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.Users{})

	var user model.Users
	db.Find(&user)
	val, err := redis.Get(ctx, "key").Result()

	api := engine.Group("/api")
	api.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":     "123",
			"data":        user,
			"redis_value": val,
			"hello":       "123",
		})
	})

	engine.Run(":8080")
	return engine
}
