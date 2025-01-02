package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"northal.com/internal/biz"
	"northal.com/internal/pkg/response"
	"northal.com/internal/services"
)

type MenuApi struct {
	service   *services.MenuService
	apiPrefix string
}

func NewMenuApi(service *services.MenuService) *MenuApi {
	return &MenuApi{service: service, apiPrefix: "/menu"}
}

func (m *MenuApi) SetupAuthRoutes(r *gin.RouterGroup) *MenuApi {
	prefix := r.Group(m.apiPrefix)
	{
		prefix.GET("/list", m.GetMenuList)
		prefix.POST("/create", m.CreateMenu)
		prefix.PUT("/update", m.UpdateMenu)
	}

	return m
}

//	@Summary		获取菜单列表
//	@Description	获取菜单列表
//	@Tags			菜单模块
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	response.Response{data=[]biz.Menu}	成功后返回值
//	@Failure		500	{object}	response.Response					失败后返回值
//	@Router			/menu/list [get]
func (m *MenuApi) GetMenuList(c *gin.Context) {
	menus, err := m.service.GetAllMenus()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, menus)
}

//	@Summary		创建菜单
//	@Description	创建菜单
//	@Tags			菜单模块
//	@Accept			json
//	@Produce		json
//	@Param			menu	body		biz.Menu						true	"菜单信息"
//	@Success		200		{object}	response.Response{data=string}	成功后返回值
//	@Failure		500		{object}	response.Response				失败后返回值
//	@Router			/menu/create [post]
func (m *MenuApi) CreateMenu(c *gin.Context) {
	var menu biz.Menu
	if err := c.ShouldBindJSON(&menu); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	err := m.service.CreateMenu(&menu)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessWithMessage(c, true, "创建菜单成功")
}

//	@Summary		更新菜单
//	@Description	更新菜单
//	@Tags			菜单模块
//	@Accept			json
//	@Produce		json
//	@Param			menu	body		biz.Menu						true	"菜单信息"
//	@Success		200		{object}	response.Response{data=bool}	成功后返回值
//	@Failure		500		{object}	response.Response				失败后返回值
//	@Router			/menu/update [put]
func (m *MenuApi) UpdateMenu(c *gin.Context) {
	var menu biz.Menu
	if err := c.ShouldBindJSON(&menu); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	err := m.service.UpdateMenu(&menu)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessWithMessage(c, true, "更新菜单成功")
}
