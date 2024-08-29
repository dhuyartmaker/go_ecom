package service

import (
	"github.com/go-ecommerce/internal/repo"
)

type IUserService interface {
	GetUser() string
}

type userService struct {
	UserRepository repo.IUserRepository
}

func NewUserService(userRepo repo.IUserRepository) IUserService {
	return &userService{
		UserRepository: userRepo,
	}
}

func (ur *userService) GetUser() string {
	result := ur.UserRepository.GetUserByEmail("duchuy2411itd@gmail.com")
	return result
}
