package handler

import (
	"fmt"
	"food-service/internal/domain/dto"
	"food-service/internal/domain/rule"
	"food-service/internal/domain/service/interfaces"
	"food-service/pkg/jwttoken"
	"food-service/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"path/filepath"
	"strconv"
)

type FoodHandler struct {
	foodService interfaces.IFoodService
}

func NewFoodHandler(foodService interfaces.IFoodService) *FoodHandler {
	return &FoodHandler{foodService: foodService}
	//var foodHandler = FoodHandler{}
	//foodHandler.foodService = foodService
	//return &foodHandler
}
func (foodHandler *FoodHandler) SaveFood(c *gin.Context) {
	title := c.DefaultPostForm("title", "title")
	description := c.DefaultPostForm("description", "description")

	//Get token
	tokenMetaDate, err := jwttoken.ExtractTokenMetadata(c.Request)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}
	//Get userId
	userId := tokenMetaDate.UserID
	//Get file
	file, err := c.FormFile("file")
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}
	//Upload file
	path := viper.GetString("Files.Path")
	filename := filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, fmt.Sprintf("%s/%s", path, filename)); err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
	}

	var foodViewModelDTO = dto.FoodViewModel{
		UserID:      int(userId),
		Title:       title,
		Description: description,
		FoodImage:   filename,
	}

	foodRule := rule.FoodDTO{FoodViewModelDTO: &foodViewModelDTO}
	saveFoodError := foodRule.Validate("")
	if len(saveFoodError) > 0 {
		response.ResponseCustomError(c, saveFoodError, http.StatusBadRequest)
		return
	}

	result, err := foodHandler.foodService.SaveFood(&foodViewModelDTO)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
	}
	response.ResponseCreated(c, result)
}

func (foodHandler *FoodHandler) GetAllFood(c *gin.Context) {
	result, err := foodHandler.foodService.GetAllFood()
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
	}
	if result == nil {
		result = []dto.FoodDetailViewModel{}
	}
	response.ResponseOkWithData(c, result)
}

func (foodHandler *FoodHandler) FindFoodById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := foodHandler.foodService.FindFoodById(id)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
	}
	response.ResponseOkWithData(c, result)
}

func (foodHandler *FoodHandler) UpdateFood(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	title := c.DefaultPostForm("title", "title")
	description := c.DefaultPostForm("description", "description")
	userId := 0

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	filename := filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, filename); err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	var foodViewModelDTO = dto.FoodViewModel{
		ID:          id,
		UserID:      userId,
		Title:       title,
		Description: description,
		FoodImage:   filename,
	}

	foodRule := rule.FoodDTO{FoodViewModelDTO: &foodViewModelDTO}
	saveFoodError := foodRule.Validate("")
	if len(saveFoodError) > 0 {
		response.ResponseCustomError(c, saveFoodError, http.StatusBadRequest)
		return
	}

	result, err := foodHandler.foodService.UpdateFood(&foodViewModelDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (foodHandler *FoodHandler) DeleteById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	err = foodHandler.foodService.DeleteFoodById(id)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}
	response.ResponseOK(c, "Successfully Food Deleted.")
}
