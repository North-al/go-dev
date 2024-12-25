package api

import (
	"fmt"
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
		prefix.POST("/register", u.Register)
	}

	return u
}

func (u *UserApi) SetupAuthRoutes(r *gin.RouterGroup) *UserApi {
	prefix := r.Group(ApiPrefix)
	{
		prefix.GET("/info", u.GetUserInfo)
		prefix.GET("/list", u.GetUserList)
	}

	return u
}

// @Summary 用户登录
// @Description 用户登录，返回token
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param user body services.LoginParams true "用户信息"
// @Success 200 {object} response.Response{data=string}  成功后返回值
// @Failure 500 {object} response.Response  失败后返回值
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
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param user body services.RegisterParams true "用户信息"
// @Success 200 {object} response.Response{data=int}  成功后返回值
// @Failure 500 {object} response.Response  失败后返回值
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
// @Tags 用户模块
// @Accept json
// @Produce json
// @security ApiKeyAuth
// @Success 200 {object} response.Response{data=biz.Users}  成功后返回值
// @Failure 500 {object} response.Response  失败后返回值
// @Router /user/info [get]
func (u *UserApi) GetUserInfo(c *gin.Context) {
	userID := c.GetInt("user_id")
	user, err := u.service.GetUserInfo(userID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, user)
}

// @Summary 分页获取用户列表
// @Description 分页获取用户列表
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param params query biz.PaginationRequest true "分页参数"
// @Success 200 {object} response.Response{data=biz.PaginationResponse{list=[]biz.Users,total=int64}}  成功后返回值
// @Failure 500 {object} response.Response  失败后返回值
// @Router /user/list [get]
func (u *UserApi) GetUserList(c *gin.Context) {
	var params biz.PaginationRequest
	if err := c.ShouldBindQuery(&params); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Println(params, "params")

	users, err := u.service.GetUserList(params)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, users)
}
