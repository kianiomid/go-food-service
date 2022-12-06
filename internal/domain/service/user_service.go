package service

import (
	"fmt"
	"food-service/internal/domain/dto"
	"food-service/internal/domain/entity"
	"food-service/internal/domain/repository/repositoryInterfaces"
	"food-service/internal/domain/transformer"
	"food-service/pkg/security"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository repositoryInterfaces.IUserRepository
}

func NewUserService(userRepository repositoryInterfaces.IUserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (userService *UserService) GetAllUser() (*[]dto.UserResponseDTO, error) {
	result, err := userService.userRepository.GetAllUser()
	if err != nil {
		return nil, err
	}

	var users []dto.UserResponseDTO
	for _, item := range result {
		var userResponseDTO = transformer.UserEntityToResponseDTO(item)
		users = append(users, userResponseDTO)
	}

	return &users, nil
}

func (userService *UserService) FindUserById(id int) (*dto.UserResponseDTO, error) {
	var userResponseDTO dto.UserResponseDTO
	result, err := userService.userRepository.FindUserById(id)
	if err != nil {
		return nil, err
	}

	if result != nil {
		userResponseDTO = transformer.UserEntityToResponseDTO(*result)
	}
	return &userResponseDTO, nil
}

func (userService *UserService) SaveUser(registerViewModel *dto.RegisterViewModel) (*dto.UserResponseDTO, error) {
	var user = transformer.UserRegisterViewModelDTOToEntity(registerViewModel)

	password, err := user.EncryptPassword(registerViewModel.Password)
	if err != nil {
		return nil, err
	}
	user.Password = password
	result, err := userService.userRepository.SaveUser(user)
	if err != nil {
		return nil, err
	}

	var afterRegisterViewModel dto.UserResponseDTO
	if result != nil {
		afterRegisterViewModel = dto.UserResponseDTO{
			ID:       result.ID,
			FullName: fmt.Sprintf("%s %s", result.FirstName, result.LastName),
			Email:    result.Email,
		}
	}
	return &afterRegisterViewModel, nil
}

func (userService *UserService) UpdateUser(userViewModel *dto.UserViewModel) (*dto.UserResponseDTO, error) {
	var user = &entity.User{}
	password, err := user.EncryptPassword(user.Password)
	if err != nil {
		return nil, err
	}

	userViewModel.Password = password
	var userEntity = transformer.UserUpdateViewModelDTOToEntity(user, userViewModel)

	result, err := userService.userRepository.UpdateUser(userEntity)
	if err != nil {
		return nil, err
	}

	var userResponseDTO = transformer.UserEntityToResponseDTO(*result)

	return &userResponseDTO, err
}

func (userService *UserService) DeleteUserById(id int) error {
	err := userService.userRepository.DeleteUserById(id)
	if err != nil {
		return err
	}

	return nil
}

func (userService *UserService) GetUserNameById(id int) string {
	var fullName = userService.userRepository.GetUserNameById(id)
	return fullName
}

func (userService *UserService) GetUserByEmailPassword(loginViewModel dto.LoginViewModel) (*entity.User, error) {
	result, err := userService.userRepository.GetUserByEmailPassword(loginViewModel)
	if err != nil {
		return nil, err
	}
	// Verify Password
	err = security.VerifyPassword(result.Password, loginViewModel.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, fmt.Errorf("incorrect Password. Error %s", err.Error())
	}
	return result, nil
}
