//go:build wireinject

package wire

import (
	"github.com/go-ecommerce/internal/controller"
	"github.com/go-ecommerce/internal/repo"
	"github.com/go-ecommerce/internal/service"
	"github.com/google/wire"
)

func InitUserRouterHandler() (*controller.UserController, error) {
	wire.Build(
		repo.NewUserRepository,
		service.NewUserService,
		controller.NewUserController,
	)

	return new(controller.UserController), nil
}