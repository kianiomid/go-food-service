package service

import (
	"food-service/internal/domain/dto"
	"food-service/internal/domain/entity"
	"food-service/internal/domain/repository"
)

type FoodService struct {
	foodRepository repository.FoodRepository
	userRepository repository.UserRepository
}

type IFoodService interface {
	SaveFood(*dto.FoodViewModel) (*dto.FoodViewModel, error)
	FindFoodById(int) (*dto.FoodDetailViewModel, error)
	GetAllFood() ([]dto.FoodDetailViewModel, error)
	UpdateFood(*dto.FoodViewModel) (*dto.FoodViewModel, error)
	DeleteFoodByIdById(int) error
}

func NewFoodService(foodRepository repository.FoodRepository, userRepository repository.UserRepository) *FoodService {
	//return &FoodService{foodRepository: foodRepository, userRepository: userRepository}
	var foodService = FoodService{}
	foodService.foodRepository = foodRepository
	foodService.userRepository = userRepository
	return &foodService
}

func (foodService *FoodService) SaveFood(foodViewModel *dto.FoodViewModel) (*dto.FoodViewModel, error) {
	var food = entity.Food{
		UserID:      foodViewModel.UserID,
		Title:       foodViewModel.Title,
		Description: foodViewModel.Description,
		FoodImage:   foodViewModel.FoodImage,
	}

	result, err := foodService.foodRepository.SaveFood(&food)
	if err != nil {
		return nil, err
	}
	if result != nil {
		foodViewModel = &dto.FoodViewModel{
			ID:          result.ID,
			UserID:      result.UserID,
			Title:       result.Title,
			Description: result.Description,
			FoodImage:   result.FoodImage,
		}
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
		foodVM = dto.FoodDetailViewModel{
			UserName:    foodService.userRepository.GetUserNameById(result.UserID),
			Title:       result.Title,
			FoodImage:   result.FoodImage,
			Description: result.Description,
		}
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
			foodVM := dto.FoodDetailViewModel{
				UserName:    foodService.userRepository.GetUserNameById(item.UserID),
				Title:       item.Title,
				FoodImage:   item.FoodImage,
				Description: item.Description,
			}

			foodListVM = append(foodListVM, foodVM)
		}
	}
	return foodListVM, nil
}

func (foodService *FoodService) UpdateFood(foodVM *dto.FoodViewModel) (*dto.FoodViewModel, error) {
	var food = entity.Food{
		ID:          foodVM.ID,
		UserID:      foodVM.UserID,
		Title:       foodVM.Title,
		Description: foodVM.Description,
		FoodImage:   foodVM.FoodImage,
	}
	_, err := foodService.foodRepository.UpdateFood(&food)
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