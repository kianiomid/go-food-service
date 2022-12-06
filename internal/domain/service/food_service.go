package service

import (
	"food-service/internal/domain/dto"
	"food-service/internal/domain/repository/repositoryInterfaces"
	"food-service/internal/domain/service/serviceInterfaces"
	"food-service/internal/domain/transformer"
)

type FoodService struct {
	foodRepository repositoryInterfaces.IFoodRepository
	userService    serviceInterfaces.IUserService
}

func NewFoodService(foodRepository repositoryInterfaces.IFoodRepository, userService serviceInterfaces.IUserService) *FoodService {
	return &FoodService{foodRepository: foodRepository, userService: userService}
	//var foodService = FoodService{}
	//foodService.foodRepository = foodRepository
	//foodService.userRepository = userRepository
	//return &foodService
}

func (foodService *FoodService) SaveFood(foodViewModel *dto.FoodViewModel) (*dto.FoodViewModel, error) {
	var food = transformer.FoodViewModelDTOToEntity(foodViewModel)
	result, err := foodService.foodRepository.SaveFood(food)
	if err != nil {
		return nil, err
	}
	if result != nil {
		foodViewModel = transformer.FoodEntityToViewModelDTO(result)
	}

	return foodViewModel, nil
}

func (foodService *FoodService) FindFoodById(id int) (*dto.FoodDetailViewModel, error) {
	result, err := foodService.foodRepository.FindFoodById(id)
	if err != nil {
		return nil, err
	}

	var foodVM dto.FoodDetailViewModel
	if result != nil {
		var username = foodService.userService.GetUserNameById(result.UserID)
		foodVM = transformer.FoodEntityToDetailViewModelDTO(*result, username)
	}

	return &foodVM, nil
}

func (foodService *FoodService) GetAllFood() ([]dto.FoodDetailViewModel, error) {
	result, err := foodService.foodRepository.GetAllFood()
	if err != nil {
		return nil, err
	}

	var foodListVM []dto.FoodDetailViewModel
	if result != nil {
		for _, item := range result {
			var username = foodService.userService.GetUserNameById(item.UserID)
			foodVM := transformer.FoodEntityToDetailViewModelDTO(item, username)
			foodListVM = append(foodListVM, foodVM)
		}
	}

	return foodListVM, nil
}

func (foodService *FoodService) UpdateFood(foodVM *dto.FoodViewModel) (*dto.FoodViewModel, error) {
	var food = transformer.FoodUpdateViewModelDTOToEntity(foodVM)
	_, err := foodService.foodRepository.UpdateFood(food)
	if err != nil {
		return nil, err
	}
	return foodVM, nil
}

func (foodService *FoodService) DeleteFoodById(id int) error {
	err := foodService.foodRepository.DeleteFoodById(id)
	if err != nil {
		return err
	}
	return nil
}
