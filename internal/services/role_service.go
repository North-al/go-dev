package services

import (
	"errors"

	"gorm.io/gorm"
	"northal.com/internal/biz"
	"northal.com/internal/data"
)

type RoleService struct {
	repo *data.RoleRepo
}

func NewRoleService(repo *data.RoleRepo) *RoleService {
	return &RoleService{repo: repo}
}

func (r *RoleService) CreateRole(role *biz.Role) (int64, error) {

	// 1. 查询角色是否存在
	_, count, err := r.repo.GetRoleByName(role.Name)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, err
	}

	if count > 0 {
		return 0, errors.New("角色已存在")
	}

	_role := biz.Role{
		Name:        role.Name,
		Description: role.Description,
	}

	// 2. 不存在则新建角色
	return int64(_role.ID), r.repo.CreateRole(&_role)
}
