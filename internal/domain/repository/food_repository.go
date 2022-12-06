package repository

import (
	"food-service/internal/domain/entity"
	"github.com/jinzhu/gorm"
)

type FoodRepository struct {
	db *gorm.DB
}

func NewFoodRepository(db *gorm.DB) *FoodRepository {
	return &FoodRepository{db}
	//var foodRepository = FoodRepository{}
	//foodRepository.db = db
	//return &foodRepository
}

func (foodRepository *FoodRepository) SaveFood(food *entity.Food) (*entity.Food, error) {
	err := foodRepository.db.Create(&food).Error
	if err != nil {
		return nil, err
	}
	return food, nil
}

func (foodRepository *FoodRepository) FindFoodById(id int) (*entity.Food, error) {
	var food entity.Food
	err := foodRepository.db.Where("id = ?", id).Take(&food).Error
	if err != nil {
		return nil, err
	}
	return &food, nil
}

func (foodRepository *FoodRepository) GetAllFood() ([]entity.Food, error) {
	var foods []entity.Food
	err := foodRepository.db.Order("id desc").Find(&foods).Error
	if err != nil {
		return nil, err
	}
	return foods, nil
}

func (foodRepository *FoodRepository) UpdateFood(food *entity.Food) (*entity.Food, error) {
	err := foodRepository.db.Save(&food).Error
	if err != nil {
		return nil, err
	}
	return food, nil
}

func (foodRepository *FoodRepository) DeleteFoodById(id int) error {
	var food entity.Food
	err := foodRepository.db.Where("id = ?", id).Delete(&food).Error
	if err != nil {
		return err
	}
	return nil
}
