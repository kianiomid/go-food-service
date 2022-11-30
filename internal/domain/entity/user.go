package entity

import (
	"food-service/internal/domain/dto"
	"food-service/pkg/security"
	"github.com/badoux/checkmail"
)

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UserDTO struct {
	UserViewModel *dto.UserViewModel
}
type RegisterDTO struct {
	RegisterViewModel *dto.RegisterViewModel
}
type LoginDTO struct {
	LoginViewModel *dto.LoginViewModel
}

func (u *User) EncryptPassword(password string) (string, error) {
	hashPassword, err := security.Hash(password)
	if err != nil {
		return "", err
	}

	return string(hashPassword), nil
}

func (u *User) Validate() map[string]string {
	var errorMessages = make(map[string]string)
	var err error
	if u.Email == "" {
		errorMessages["email_required"] = "email required"
	}
	if u.Email != "" {
		if err = checkmail.ValidateFormat(u.Email); err != nil {
			errorMessages["invalid_email"] = "email email"
		}
	}

	return errorMessages
}

func (lvm *LoginDTO) Validate() map[string]string {
	var errorMessages = make(map[string]string)
	var err error
	if lvm.LoginViewModel.Password == "" {
		errorMessages["password_required"] = "password is required"
	}
	if lvm.LoginViewModel.Email == "" {
		errorMessages["email_required"] = "email is required"
	}
	if lvm.LoginViewModel.Email != "" {
		if err = checkmail.ValidateFormat(lvm.LoginViewModel.Email); err != nil {
			errorMessages["invalid_email"] = "please provide a valid email"
		}
	}

	return errorMessages
}

func (r *RegisterDTO) Validate() map[string]string {
	var errorMessages = make(map[string]string)
	var err error

	if r.RegisterViewModel.FirstName == "" {
		errorMessages["firstname_required"] = "first name is required"
	}
	if r.RegisterViewModel.LastName == "" {
		errorMessages["lastname_required"] = "last name is required"
	}
	if r.RegisterViewModel.Password == "" {
		errorMessages["password_required"] = "password is required"
	}
	if r.RegisterViewModel.Password != "" && len(r.RegisterViewModel.Password) < 6 {
		errorMessages["invalid_password"] = "password should be at least 6 characters"
	}
	if r.RegisterViewModel.Email == "" {
		errorMessages["email_required"] = "email is required"
	}
	if r.RegisterViewModel.Email != "" {
		if err = checkmail.ValidateFormat(r.RegisterViewModel.Email); err != nil {
			errorMessages["invalid_email"] = "please provide a valid email"
		}
	}

	return errorMessages
}
