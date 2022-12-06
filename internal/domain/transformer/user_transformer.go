package transformer

import (
	"fmt"
	"food-service/internal/domain/dto"
	"food-service/internal/domain/entity"
)

func UserRegisterViewModelDTOToEntity(registerViewModel *dto.RegisterViewModel) *entity.User {
	var user = entity.User{
		FirstName: registerViewModel.FirstName,
		LastName:  registerViewModel.LastName,
		Email:     registerViewModel.Email,
	}
	return &user
}

func UserEntityToResponseDTO(userEntity entity.User) dto.UserResponseDTO {
	var userResponseDTO dto.UserResponseDTO
	userResponseDTO.ID = userEntity.ID
	userResponseDTO.FullName = fmt.Sprintf("%s %s", userEntity.FirstName, userEntity.LastName)
	userResponseDTO.Email = userEntity.Email

	return userResponseDTO
}

func UserUpdateViewModelDTOToEntity(user *entity.User, userViewModel *dto.UserViewModel) *entity.User {
	user.ID = userViewModel.ID
	user.FirstName = userViewModel.FirstName
	user.LastName = userViewModel.LastName
	user.Email = userViewModel.Email
	user.Password = userViewModel.Password

	return user
}
