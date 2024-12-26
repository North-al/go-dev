package data

import (
	"errors"
	"reflect"

	"gorm.io/gorm"
	"northal.com/internal/biz"
)

type RoleRepo struct {
	db *gorm.DB
}

func NewRoleRepo(db *gorm.DB) *RoleRepo {
	return &RoleRepo{db: db}
}

func (r *RoleRepo) CreateRole(role *biz.Role) error {
	if role == nil {
		return errors.New("role is required")
	}

	// 或者role是一个空结构体
	if reflect.DeepEqual(role, biz.Role{}) {
		return errors.New("role is empty")
	}

	return r.db.Create(role).Error
}

// 根据条件查询角色
func (r *RoleRepo) GetRoleByCondition(condition string, args ...interface{}) (*biz.Role, int64, error) {
	var role biz.Role
	var count int64 = 0

	if err := r.db.Model(&biz.Role{}).
		Where(condition, args...).
		First(&role).
		Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if count == 0 {
		return nil, count, gorm.ErrRecordNotFound
	}

	return &role, count, nil
}

// 根据id查询角色
func (r *RoleRepo) GetRoleByID(id int) (*biz.Role, int64, error) {
	return r.GetRoleByCondition("id = ?", id)
}

// 根据角色名查询角色
func (r *RoleRepo) GetRoleByName(name string) (*biz.Role, int64, error) {
	return r.GetRoleByCondition("name = ?", name)
}

// 分页获取角色列表
func (r *RoleRepo) GetRoleList(params biz.PaginationRequest) ([]biz.Role, int64, error) {
	var roles []biz.Role
	var count int64 = 0

	if err := r.db.Model(&biz.Role{}).
		Count(&count).
		Offset((params.Page - 1) * params.PageSize).
		Find(&roles).Error; err != nil {
		return nil, 0, err
	}

	return roles, count, nil
}
