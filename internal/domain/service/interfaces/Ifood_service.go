package interfaces

import "food-service/internal/domain/dto"

type IFoodService interface {
	SaveFood(*dto.FoodViewModel) (*dto.FoodViewModel, error)
	FindFoodById(int) (*dto.FoodDetailViewModel, error)
	GetAllFood() ([]dto.FoodDetailViewModel, error)
	UpdateFood(*dto.FoodViewModel) (*dto.FoodViewModel, error)
	DeleteFoodById(int) error
}
