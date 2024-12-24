package api

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"northal.com/internal/data"
	"northal.com/internal/middleware"
	"northal.com/internal/services"
)

type SetupApi struct {
	router *gin.Engine
	DB     *gorm.DB
	Redis  *redis.Client
}

var prefix = "/api"

func NewSetupApi(router *gin.Engine, db *gorm.DB, redis *redis.Client) *SetupApi {
	return &SetupApi{router: router, DB: db, Redis: redis}
}

func (s *SetupApi) SetupRoutes() {
	public := s.router.Group(prefix)

	auth := s.router.Group(prefix)
	auth.Use(middleware.AuthHandler())

	NewUserApi(services.NewUserService(data.NewUserRepo(s.DB, s.Redis))).SetupPublicRoutes(public).SetupAuthRoutes(auth)
}
