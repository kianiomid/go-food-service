package interfaces

import "food-service/internal/domain/entity"

type IFoodRepository interface {
	SaveFood(*entity.Food) (*entity.Food, error)
	FindFoodById(int) (*entity.Food, error)
	GetAllFood() ([]entity.Food, error)
	UpdateFood(*entity.Food) (*entity.Food, error)
	DeleteFoodById(int) error
}
