package service

import (
	"fmt"

	"github.com/go-ecommerce/internal/repo"
)

type IUserService interface {
	GetUser() string
	Register(email string)
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
	ur.UserRepository.GetUserByEmail("duchuy2411itd@gmail.com")
	return "ok"
}

func (ur *userService) Register(email string) {
	otp := 123456

	fmt.Printf("Otp is %d", otp)
	// Save in Redis
}
