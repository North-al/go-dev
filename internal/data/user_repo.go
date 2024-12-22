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

// 根据id查询用户
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

// 根据账号查询用户
func (r *UserRepo) GetUserByAccount(account string) (int64, error) {
	if account == "" {
		return 0, errors.New("账号是必填项")
	}

	var count int64 = 0
	if err := r.db.Model(&biz.Users{}).Where("email = ? or phone = ? or username = ?", account, account, account).Count(&count).Error; err != nil {
		return count, err
	}

	return count, nil
}

// 根据用户名查询用户
func (r *UserRepo) GetUserByUsername(username string) (*biz.Users, error) {
	if username == "" {
		return nil, errors.New("username is required")
	}

	var user biz.Users
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, errors.New("当前用户不存在")
	}

	return &user, nil
}

func (r *UserRepo) Create(user *biz.Users) error {
	if user == nil {
		return errors.New("user is required")
	}

	return r.db.Create(user).Error
}
