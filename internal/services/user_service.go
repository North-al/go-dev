package services

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
	"northal.com/internal/biz"
	"northal.com/internal/data"
	"northal.com/internal/pkg/jwt"
	"northal.com/internal/pkg/random"
	"northal.com/internal/pkg/verify"
)

type UserService struct {
	repo *data.UserRepo
}

type LoginParams struct {
	Account  string `json:"account" binding:"required" default:"zhangsan"` // 账号 邮箱/手机号/用户名
	Password string `json:"password" binding:"required" default:"123456"`  // 密码 6-16位
}

type RegisterParams struct {
	LoginParams
}

type LoginResponse struct {
	Token string `json:"token"` // 令牌
}

func NewUserService(repo *data.UserRepo) *UserService {
	return &UserService{repo: repo}
}

func getEmailIfValid(account string) string {
	if verify.IsEmail(account) {
		return account
	}
	return ""
}

func getPhoneIfValid(account string) string {
	if verify.IsPhone(account) {
		return account
	}
	return ""
}

func (u *UserService) Login(params LoginParams) (*LoginResponse, error) {
	// 1. 查询用户是否存在
	user, _, err := u.repo.GetUserByAccount(params.Account)
	if err != nil {
		return nil, err
	}

	// 2. 验证密码是否正确
	if user.Password != params.Password {
		return nil, errors.New("password is incorrect")
	}

	// 3. 生成token
	jwtInstance := jwt.NewJwt()
	token, err := jwtInstance.GenerateToken(int(user.ID))
	if err != nil {
		return nil, err
	}

	// 4. redis缓存token
	u.repo.SetToken(int(user.ID), token)

	return &LoginResponse{Token: token}, nil
}

func (u *UserService) Register(params RegisterParams) (int64, error) {

	// 1. 查询用户是否存在
	_, count, err := u.repo.GetUserByAccount(params.Account)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, err
	}

	if count > 0 {
		return count, errors.New("用户已存在")
	}

	// 2. 不存在则新建用户
	user := biz.Users{
		Username: random.GenerateRandomUsername(),
		Password: params.Password,
		Email:    getEmailIfValid(params.Account),
		Phone:    getPhoneIfValid(params.Account),
	}

	// 判断是邮箱还是手机号、用户名
	if getEmailIfValid(params.Account) != "" {
		user.Email = getEmailIfValid(params.Account)
	} else if getPhoneIfValid(params.Account) != "" {
		user.Phone = getPhoneIfValid(params.Account)
	} else {
		user.Username = params.Account
	}

	fmt.Println(user)

	err = u.repo.Create(&user)

	if err != nil {
		return 0, err
	}

	// 3. 返回结果
	return int64(user.ID), nil
}

// 获取用户信息
func (u *UserService) GetUserInfo(id int) (*biz.Users, error) {
	return u.repo.GetUserByID(id)
}

// 分页获取用户列表
func (u *UserService) GetUserList(params biz.PaginationRequest) (*biz.PaginationResponse, error) {
	users, total, err := u.repo.GetUserList(params)
	if err != nil {
		return nil, err
	}

	return &biz.PaginationResponse{
		Total: total,
		List:  users,
	}, nil
}
