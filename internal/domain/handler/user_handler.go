package handler

import (
	"errors"
	"food-service/internal/domain/dto"
	"food-service/internal/domain/entity"
	"food-service/internal/domain/rule"
	"food-service/internal/domain/service"
	"food-service/pkg/jwttoken"
	"food-service/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserHandler struct {
	userService service.IUserService
}

func NewUserHandler(userService service.IUserService) *UserHandler {
	//return &UserHandler{userService: userService}
	var userHandler = UserHandler{}
	userHandler.userService = userService
	return &userHandler
}

func (userHandler *UserHandler) RegisterUser(c *gin.Context) {
	var registerViewModelDTO dto.RegisterViewModel
	err := c.ShouldBindJSON(&registerViewModelDTO)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	userRule := rule.RegisterDTO{RegisterViewModel: &registerViewModelDTO}
	registerUserError := userRule.Validate()
	if len(registerUserError) > 0 {
		response.ResponseCustomError(c, registerUserError, http.StatusBadRequest)
		return
	}

	result, err := userHandler.userService.SaveUser(&registerViewModelDTO)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	response.ResponseCreated(c, result)
}

func (userHandler *UserHandler) GetAllUser(c *gin.Context) {
	result, err := userHandler.userService.GetAllUser()
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}
	if result == nil {
		result = &[]dto.UserResponseDTO{}
	}

	response.ResponseOkWithData(c, result)
}

func (userHandler *UserHandler) FindUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := userHandler.userService.FindUserById(id)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}
	if result == nil {
		result = &dto.UserResponseDTO{}
	}

	response.ResponseOkWithData(c, result)
}

func (userHandler *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.ResponseError(c, errors.New("invalid User ID").Error(), http.StatusBadRequest)
		return
	}

	//var updateUser entity.User
	var userViewModelDTO dto.UserViewModel
	err = c.ShouldBindJSON(&userViewModelDTO)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	//updateUser.ID = id
	userViewModelDTO.ID = id

	userRule := rule.UserDTO{UserViewModel: &userViewModelDTO}
	updateUserError := userRule.Validate()
	if len(updateUserError) > 0 {
		response.ResponseCustomError(c, updateUserError, http.StatusBadRequest)
		return
	}

	result, err := userHandler.userService.UpdateUser(&userViewModelDTO)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	if result == nil {
		result = &dto.UserResponseDTO{}
	}

	response.ResponseOkWithData(c, result)
}

func (userHandler *UserHandler) DeleteUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.ResponseError(c, errors.New("invalid User ID").Error(), http.StatusBadRequest)
		return
	}

	err = userHandler.userService.DeleteUserById(id)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	response.ResponseOK(c, "Successfully User Deleted")
}

func (userHandler *UserHandler) Login(c *gin.Context) {
	var loginVM dto.LoginViewModel

	err := c.ShouldBindJSON(&loginVM)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	validateUser, err := userHandler.userService.GetUserByEmailPassword(loginVM)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	if validateUser == nil {
		validateUser = &entity.User{}
	}

	//Generate JWT
	token, err := jwttoken.CreateToken(int64(validateUser.ID))
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	userData := map[string]interface{}{
		"access_token": token.AccessToken,
		"expired":      token.ExpiredToken,
		"user_id":      validateUser.ID,
	}

	response.ResponseOkWithData(c, userData)
}
