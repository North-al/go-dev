package server

import (
	"context"

	"github.com/gin-gonic/gin"
	"northal.com/api"
	"northal.com/config"
	"northal.com/internal/biz"
	"northal.com/internal/data"
	"northal.com/internal/middleware"
	"northal.com/internal/pkg/response"
	"northal.com/internal/services"
)

func InitHttp() *gin.Engine {
	config.InitViper()
	engine := gin.Default()
	engine.Use(middleware.ErrorHandler())
	ctx := context.Background()

	db := InitDatabase()
	redis := InitRedis()

	publicApi := engine.Group("/api")
	authApi := engine.Group("/api")

	api.NewUserApi(services.NewUserService(data.NewUserRepo(db))).SetupPublicRoutes(publicApi).SetupAuthRoutes(authApi)

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
		response.Success(c, gin.H{
			"message":     "123",
			"data":        user,
			"redis_value": val,
			"hello":       "123",
			"config":      config.GetAppConfig(),
			"database":    config.GetDatabaseConfig(),
		})
	})

	api.GET("/error", func(c *gin.Context) {
		panic("error")
	})

	engine.Run(":8080")
	return engine
}
