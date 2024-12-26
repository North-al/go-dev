package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"northal.com/internal/biz"
	"northal.com/internal/pkg/response"
	"northal.com/internal/services"
)

type RoleApi struct {
	service   *services.RoleService
	apiPrefix string
}

func NewRoleApi(service *services.RoleService) *RoleApi {
	return &RoleApi{service: service, apiPrefix: "/role"}
}

func (u *RoleApi) SetupPublicRoutes(r *gin.RouterGroup) *RoleApi {
	prefix := r.Group(u.apiPrefix)
	{
		prefix.POST("/create", u.CreateRole)
	}

	return u
}

func (u *RoleApi) SetupAuthRoutes(r *gin.RouterGroup) *RoleApi {
	// prefix := r.Group(u.apiPrefix)
	// {
	// prefix.GET("/info", u.GetUserInfo)
	// prefix.GET("/list", u.GetUserList)
	// }

	return u
}

// @Summary 创建角色
// @Description 创建角色
// @Tags 角色模块
// @Accept json
// @Produce json
// @Param role body biz.Role true "角色信息"
// @Success 200 {object} response.Response{data=int64}  成功后返回值
// @Failure 500 {object} response.Response  失败后返回值
// @Router /role/create [post]
func (r *RoleApi) CreateRole(c *gin.Context) {
	var req biz.Role
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := r.service.CreateRole(&req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessWithMessage(c, id, "创建角色成功")
}
