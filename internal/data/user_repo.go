package data

import (
	"errors"

	"gorm.io/gorm"
	"northal.com/internal/biz"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) GetUserByID(id int) (*biz.Users, error) {

	if id == 0 {
		return nil, errors.New("id is required")
	}

	var user biz.Users
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) Create(user *biz.Users) error {
	if user == nil {
		return errors.New("user is required")
	}

	return r.db.Create(user).Error
}
