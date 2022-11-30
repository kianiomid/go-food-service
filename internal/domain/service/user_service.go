package service

import (
	"fmt"
	"food-service/internal/domain/dto"
	"food-service/internal/domain/entity"
	"food-service/internal/domain/repository"
	"food-service/pkg/security"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository repository.IUserRepository
}

type IUserService interface {
	SaveUser(model *dto.RegisterViewModel) (*dto.UserViewModel, error)
	GetAllUser() (*[]dto.UserViewModel, error)
	FindUserById(id int) (*dto.UserViewModel, error)
	UpdateUser(userViewModel *entity.User) (*dto.UserViewModel, error)
	DeleteUserById(id int) error
	GetUserByEmailPassword(loginViewModel dto.LoginViewModel) (*entity.User, error)
}

func NewUserService(userRepository repository.IUserRepository) *UserService {
	//return &UserService{userRepository: userRepository}
	var userService = UserService{}
	userService.userRepository = userRepository
	return &userService
}

func (userService *UserService) GetAllUser() (*[]dto.UserViewModel, error) {
	result, err := userService.userRepository.GetAllUser()
	if err != nil {
		return nil, err
	}

	var users []dto.UserViewModel
	for _, item := range result {
		var user dto.UserViewModel
		user.Email = item.Email
		user.FullName = fmt.Sprintf("%s %s", item.FirstName, item.LastName)
		user.Email = item.Email
		users = append(users, user)
	}
	return &users, nil
}

func (userService *UserService) FindUserById(id int) (*dto.UserViewModel, error) {
	var viewModel dto.UserViewModel
	result, err := userService.userRepository.FindUserById(id)
	if err != nil {
		return nil, err
	}

	if result != nil {
		viewModel = dto.UserViewModel{
			ID:       result.ID,
			FullName: fmt.Sprintf("%s %s", result.FirstName, result.LastName),
			Email:    result.Email,
		}
	}
	return &viewModel, nil
}

func (userService *UserService) SaveUser(registerViewModel *dto.RegisterViewModel) (*dto.UserViewModel, error) {
	var user = entity.User{
		FirstName: registerViewModel.FirstName,
		LastName:  registerViewModel.LastName,
		Email:     registerViewModel.Email,
	}

	password, err := user.EncryptPassword(registerViewModel.Password)
	if err != nil {
		return nil, err
	}
	user.Password = password
	result, err := userService.userRepository.SaveUser(&user)
	if err != nil {
		return nil, err
	}

	var afterRegisterViewModel dto.UserViewModel
	if result != nil {
		afterRegisterViewModel = dto.UserViewModel{
			ID:       result.ID,
			FullName: fmt.Sprintf("%s %s", result.FirstName, result.LastName),
			Email:    result.Email,
		}
	}
	return &afterRegisterViewModel, nil
}

func (userService *UserService) UpdateUser(user *entity.User) (*dto.UserViewModel, error) {
	password, err := user.EncryptPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = password
	result, err := userService.userRepository.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	var userAfterUpdate dto.UserViewModel
	userAfterUpdate = dto.UserViewModel{
		ID:       result.ID,
		FullName: fmt.Sprintf("%s %s", result.FirstName, result.LastName),
		Email:    result.Email,
	}
	return &userAfterUpdate, err
}

func (userService *UserService) DeleteUserById(id int) error {
	err := userService.userRepository.DeleteUserById(id)
	if err != nil {
		return err
	}

	return nil
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
