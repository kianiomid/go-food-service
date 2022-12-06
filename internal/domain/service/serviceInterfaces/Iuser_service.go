package serviceInterfaces

import (
	"food-service/internal/domain/dto"
	"food-service/internal/domain/entity"
)

type IUserService interface {
	SaveUser(model *dto.RegisterViewModel) (*dto.UserResponseDTO, error)
	GetAllUser() (*[]dto.UserResponseDTO, error)
	FindUserById(id int) (*dto.UserResponseDTO, error)
	UpdateUser(userViewModel *dto.UserViewModel) (*dto.UserResponseDTO, error)
	DeleteUserById(id int) error
	GetUserNameById(id int) string
	GetUserByEmailPassword(loginViewModel dto.LoginViewModel) (*entity.User, error)
}
