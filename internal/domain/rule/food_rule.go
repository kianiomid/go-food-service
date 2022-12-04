package rule

import (
	"food-service/internal/domain/dto"
	"strings"
)

type FoodDTO struct {
	FoodViewModelDTO *dto.FoodViewModel
}

func (foodDTO *FoodDTO) Validate(action string) map[string]string {
	var errorMessages = make(map[string]string)
	switch strings.ToLower(action) {
	case "update":
		if foodDTO.FoodViewModelDTO.Title == "" || foodDTO.FoodViewModelDTO.Title == "null" {
			errorMessages["title_required"] = "title is required"
		}
		if foodDTO.FoodViewModelDTO.Title == "" || foodDTO.FoodViewModelDTO.Title == "null" {
			errorMessages["desc_required"] = "description is required"
		}
	default:
		if foodDTO.FoodViewModelDTO.Title == "" || foodDTO.FoodViewModelDTO.Title == "null" {
			errorMessages["title_required"] = "title is required"
		}
		if foodDTO.FoodViewModelDTO.Title == "" || foodDTO.FoodViewModelDTO.Title == "null" {
			errorMessages["desc_required"] = "description is required"
		}
	}
	return errorMessages
}
