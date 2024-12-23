package api

import (
	"net/http"

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
	{
		prefix.POST("/login", u.Login)
		prefix.POST("/register", u.Register)
	}

	return u
}

func (u *UserApi) SetupAuthRoutes(r *gin.RouterGroup) *UserApi {
	prefix := r.Group(ApiPrefix)
	{
		prefix.GET("/info", u.GetUserInfo)
	}

	return u
}

// @Summary 登录
// @Description 登录
// @Tags 用户
// @Accept json
// @Produce json
// @Param user body services.LoginParams true "用户信息"
// @Success 200 {object} response.Response
// @Router /user/login [post]
func (u *UserApi) Login(c *gin.Context) {
	var params services.LoginParams
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

// @Summary 注册
// @Description 注册
// @Tags 用户
// @Accept json
// @Produce json
// @Param user body services.RegisterParams true "用户信息"
// @Success 200 {object} response.Response
// @Router /user/register [post]
func (u *UserApi) Register(c *gin.Context) {
	var body services.RegisterParams

	if err := c.ShouldBindJSON(&body); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	result, err := u.service.Register(body)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, result)
}

// @Summary 获取用户信息
// @Description 获取用户信息
// @Tags 用户
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /user/info [get]
func (u *UserApi) GetUserInfo(c *gin.Context) {
	userID := c.GetInt("user_id")
	response.Success(c, userID)
}
