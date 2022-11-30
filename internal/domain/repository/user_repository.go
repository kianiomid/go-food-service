package repository

import (
	"fmt"
	"food-service/internal/domain/dto"
	"food-service/internal/domain/entity"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

type IUserRepository interface {
	SaveUser(*entity.User) (*entity.User, error)
	FindUserById(int) (*entity.User, error)
	GetAllUser() ([]entity.User, error)
	UpdateUser(*entity.User) (*entity.User, error)
	DeleteUserById(int) error
	GetUserNameById(id int) string
	GetUserByEmailPassword(loginViewModel dto.LoginViewModel) (*entity.User, error)
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	//return &UserRepository{db}
	var userRepository = UserRepository{}
	userRepository.db = db
	return &userRepository
}

func (userRepository *UserRepository) SaveUser(user *entity.User) (*entity.User, error) {
	err := userRepository.db.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (userRepository *UserRepository) FindUserById(id int64) (*entity.User, error) {
	var user entity.User
	err := userRepository.db.Where("id = ?", id).Take(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (userRepository *UserRepository) GetUserNameById(id int64) string {
	userDetail, _ := userRepository.FindUserById(id)
	var fullname = fmt.Sprintf("%s %s", userDetail.FirstName, userDetail.LastName)
	return fullname
}

func (userRepository *UserRepository) GetAllUser() ([]entity.User, error) {
	var users []entity.User
	err := userRepository.db.Order("id desc").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (userRepository *UserRepository) UpdateUser(user *entity.User) (*entity.User, error) {
	err := userRepository.db.Save(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (userRepository *UserRepository) DeleteUserById(id int) error {
	var user entity.User
	err := userRepository.db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (userRepository *UserRepository) GetUserByEmailPassword(loginViewModel dto.LoginViewModel) (*entity.User, error) {
	var user entity.User
	err := userRepository.db.Where("email = ?", loginViewModel.Email).Take(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
