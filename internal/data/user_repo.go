package data

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"northal.com/config"
	"northal.com/internal/biz"
)

type UserRepo struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewUserRepo(db *gorm.DB, redis *redis.Client) *UserRepo {
	return &UserRepo{db: db, redis: redis}
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

// redis缓存token
func (r *UserRepo) SetToken(userID int, token string) error {
	ctxBackground := context.Background()
	fmt.Println(config.GetJwtConfig().TokenExpire, "setToken config.GetJwtConfig().TokenExpire")
	expireTime := time.Duration(config.GetJwtConfig().TokenExpire) * time.Hour
	return r.redis.Set(ctxBackground, fmt.Sprintf("user:%d:token", userID), token, expireTime).Err()
}
