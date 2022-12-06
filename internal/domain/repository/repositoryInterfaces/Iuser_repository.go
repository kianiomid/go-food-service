package repositoryInterfaces

import (
	"food-service/internal/domain/dto"
	"food-service/internal/domain/entity"
)

type IUserRepository interface {
	SaveUser(*entity.User) (*entity.User, error)
	FindUserById(int) (*entity.User, error)
	GetAllUser() ([]entity.User, error)
	UpdateUser(*entity.User) (*entity.User, error)
	DeleteUserById(int) error
	GetUserNameById(id int) string
	GetUserByEmailPassword(loginViewModel dto.LoginViewModel) (*entity.User, error)
}
