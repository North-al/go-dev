package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"northal.com/internal/biz"
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
	{
		prefix.POST("/login", u.Login)
	}
	prefix.GET("/", u.GetUser)
	return u
}

func (u *UserApi) SetupAuthRoutes(r *gin.RouterGroup) *UserApi {
	// prefix := r.Group(ApiPrefix)
	// prefix.GET("/", u.GetUser)
	return u
}

func (u *UserApi) GetUser(c *gin.Context) {

}

// @Summary 登录
// @Description 登录
// @Tags 用户
// @Accept json
// @Produce json
// @Param user body biz.Users true "用户信息"
// @Success 200 {object} response.Response
// @Router /user/login [post]
func (u *UserApi) Login(c *gin.Context) {
	var params biz.Users
	if err := c.ShouldBindJSON(&params); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	success, err := u.service.Login(params)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, success)
}
