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

func (r *UserRepo) GetUserByCondition(condition string, args ...interface{}) (*biz.Users, int64, error) {
	var user biz.Users
	var count int64 = 0

	if err := r.db.Model(&biz.Users{}).Where(condition, args...).First(&user).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	return &user, count, nil
}

// 根据id查询用户
func (r *UserRepo) GetUserByID(id int) (*biz.Users, error) {

	if id == 0 {
		return nil, errors.New("id is required")
	}

	user, _, err := r.GetUserByCondition("id = ?", id)
	return user, err
}

// 根据账号查询用户
func (r *UserRepo) GetUserByAccount(account string) (*biz.Users, int64, error) {
	if account == "" {
		return nil, 0, errors.New("账号是必填项")
	}

	return r.GetUserByCondition("email = ? or phone = ? or username = ?", account, account, account)
}

// 根据用户名查询用户
func (r *UserRepo) GetUserByUsername(username string) (*biz.Users, error) {
	if username == "" {
		return nil, errors.New("username is required")
	}

	user, _, err := r.GetUserByCondition("username = ?", username)
	return user, err
}

func (r *UserRepo) Create(user *biz.Users) error {
	if user == nil {
		return errors.New("user is required")
	}

	return r.db.Create(user).Error
}
