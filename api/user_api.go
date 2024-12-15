package api

import (
	"github.com/gin-gonic/gin"
	"northal.com/internal/pkg/response"
	"northal.com/internal/services"
)

type UserApi struct {
	service *services.UserService
}

const (
	ApiPrefix = "/user"
)

func NewUserApi(service *services.UserService) *UserApi {
	return &UserApi{service: service}
}

func (u *UserApi) SetupPublicRoutes(r *gin.RouterGroup) *UserApi {
	prefix := r.Group(ApiPrefix)
	prefix.GET("/", u.GetUser)
	return u
}

func (u *UserApi) SetupAuthRoutes(r *gin.RouterGroup) *UserApi {
	// prefix := r.Group(ApiPrefix)
	// prefix.GET("/", u.GetUser)
	return u
}

func (u *UserApi) GetUser(c *gin.Context) {

	response.Success(c, gin.H{
		"message": "ok",
	})
}
