package services

import (
	"product_srv/internal/domain/users/models"
	"product_srv/internal/repositories"
)

type UserServiceInterface interface {
	GetListUser()
	CreateUser(user models.User) error
	GetUserByID(id int) (models.User, error)
}

type UserService struct {
	userRep *repositories.UserRepository
}

// NewUserService là constructor cho UserService
func NewUserService(userRep *repositories.UserRepository) *UserService {
	return &UserService{
		userRep: userRep,
	}
}

func (us *UserService) GetListUser(filter repositories.GetList) ([]models.User, int64, error) {
	return us.userRep.GetListUser(filter)
}

func (us *UserService) CreateUser(user models.User) (models.User, error) {
	return us.userRep.CreateUser(user)
}
