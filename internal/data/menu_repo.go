package data

import (
	"errors"
	"reflect"

	"gorm.io/gorm"
	"northal.com/internal/biz"
)

type MenuRepo struct {
	db *gorm.DB
}

func NewMenuRepo(db *gorm.DB) *MenuRepo {
	return &MenuRepo{db: db}
}

// 创建菜单
func (m *MenuRepo) CreateMenu(menu *biz.Menu) error {
	if menu == nil {
		return errors.New("menu is required")
	}

	if reflect.DeepEqual(menu, biz.Menu{}) {
		return errors.New("menu is empty")
	}

	return m.db.Create(menu).Error
}

func (m *MenuRepo) GetMenuByCondition(condition string, args ...interface{}) (*biz.Menu, int64, error) {
	var menu biz.Menu
	var count int64 = 0

	if err := m.db.Model(&biz.Menu{}).
		Where(condition, args...).
		First(&menu).
		Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if count == 0 {
		return nil, count, gorm.ErrRecordNotFound
	}

	return &menu, count, nil
}

// 根据path查询菜单
func (m *MenuRepo) GetMenuByPathOrName(pathOrName string) (*biz.Menu, int64, error) {
	return m.GetMenuByCondition("route_path = ? OR route_name = ?", pathOrName, pathOrName)
}

// 查询全部菜单
func (m *MenuRepo) GetAllMenus() ([]biz.Menu, error) {
	var menus []biz.Menu
	return menus, m.db.Find(&menus).Error
}

// 更新菜单
func (m *MenuRepo) UpdateMenu(menu *biz.Menu) error {
	if menu == nil {
		return errors.New("menu is required")
	}

	if reflect.DeepEqual(menu, biz.Menu{}) {
		return errors.New("menu is empty")
	}

	return m.db.Model(&biz.Menu{}).Where("id = ?", menu.ID).Updates(menu).Error
}

// 删除菜单
func (m *MenuRepo) DeleteMenu(id uint) error {
	return m.db.Delete(&biz.Menu{}, id).Error
}

// 设置菜单给角色
func (m *MenuRepo) SetMenuToRole(roleId uint, menuIds []uint) error {
	return m.db.Model(&biz.Role{}).Where("id = ?", roleId).Association("Menus").Replace(menuIds)
}
