package services

import (
	"errors"

	"northal.com/internal/biz"
	"northal.com/internal/data"
	"northal.com/internal/pkg/jwt"
)

type UserService struct {
	repo *data.UserRepo
}

func NewUserService(repo *data.UserRepo) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) Login(params biz.Users) (*struct{ Token string }, error) {
	// 1. 查询用户是否存在
	user, err := u.repo.GetUserByUsername(params.Username)
	if err != nil {
		return nil, err
	}

	// 2. 验证密码是否正确
	if user.Password != params.Password {
		return nil, errors.New("password is incorrect")
	}

	// 3. 生成token
	jwtInstance := jwt.NewJwt("Northal")
	token, err := jwtInstance.GenerateToken(int(user.ID))
	if err != nil {
		return nil, err
	}

	return &struct{ Token string }{Token: token}, nil
}
