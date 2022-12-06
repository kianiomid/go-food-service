package transformer

import (
	"food-service/internal/domain/dto"
	"food-service/internal/domain/entity"
)

func FoodViewModelDTOToEntity(foodViewModel *dto.FoodViewModel) *entity.Food {
	var food = entity.Food{
		UserID:      foodViewModel.UserID,
		Title:       foodViewModel.Title,
		Description: foodViewModel.Description,
		FoodImage:   foodViewModel.FoodImage,
	}
	return &food
}

func FoodUpdateViewModelDTOToEntity(foodViewModel *dto.FoodViewModel) *entity.Food {
	var food = entity.Food{
		ID:          foodViewModel.ID,
		UserID:      foodViewModel.UserID,
		Title:       foodViewModel.Title,
		Description: foodViewModel.Description,
		FoodImage:   foodViewModel.FoodImage,
	}
	return &food
}

func FoodEntityToViewModelDTO(food *entity.Food) *dto.FoodViewModel {
	var foodViewModel = &dto.FoodViewModel{
		ID:          food.ID,
		UserID:      food.UserID,
		Title:       food.Title,
		Description: food.Description,
		FoodImage:   food.FoodImage,
	}
	return foodViewModel
}

func FoodEntityToDetailViewModelDTO(food entity.Food, username string) dto.FoodDetailViewModel {
	var foodVM = dto.FoodDetailViewModel{
		ID:          food.ID,
		UserName:    username,
		Title:       food.Title,
		FoodImage:   food.FoodImage,
		Description: food.Description,
	}
	return foodVM
}
