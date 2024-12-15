package services

import (
	"northal.com/internal/data"
)

type UserService struct {
	repo *data.UserRepo
}

func NewUserService(repo *data.UserRepo) *UserService {
	return &UserService{repo: repo}
}
