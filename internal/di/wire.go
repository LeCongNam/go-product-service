// go:build wireinject
//go:build wireinject
// +build wireinject

package diContainer

import (
	user_controllers "product_srv/internal/domain/users/controllers"
	"product_srv/internal/domain/users/services"
	"product_srv/internal/repositories"

	"github.com/google/wire"
)

// User Domain Injection
func InitUserDomain() (*user_controllers.UserController, error) {
	panic(wire.Build(
		repositories.NewUserRepository,
		services.NewUserService,
		user_controllers.NewUserController,
	))
}
