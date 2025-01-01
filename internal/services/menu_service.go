package services

import (
	"errors"

	"gorm.io/gorm"
	"northal.com/internal/biz"
	"northal.com/internal/data"
)

type MenuService struct {
	repo *data.MenuRepo
}

func NewMenuService(repo *data.MenuRepo) *MenuService {
	return &MenuService{repo: repo}
}

func (m *MenuService) CreateMenu(menu *biz.Menu) error {

	// 1. 查询菜单是否存在
	_menu, _, err := m.repo.GetMenuByPathOrName(menu.RoutePath)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if _menu.RouteName == menu.RouteName {
		return errors.New("菜单的Name不能相同")
	}

	if _menu.RoutePath == menu.RoutePath {
		return errors.New("菜单的Path不能相同")
	}

	// 2. 新建菜单
	return m.repo.CreateMenu(menu)
}

func (m *MenuService) GetAllMenus() ([]biz.Menu, error) {
	return m.repo.GetAllMenus()
}
