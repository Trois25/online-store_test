package repository

import (
	"errors"
	categories "store/features/categories/entity"
	"store/features/categories/model"

	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) categories.CategoryDataInterface {
	return &categoryRepository{
		db: db,
	}
}

// CreateCategory implements entity.CategoryDataInterface.
func (categoryRepo *categoryRepository) CreateCategory(data categories.CategoryCore) (err error) {
	var input = model.Categories{
		Category: data.Category,
	}

	errData := categoryRepo.db.Save(&input)
	if errData != nil {
		return errData.Error
	}

	return nil
}

// DeleteCategory implements entity.CategoryDataInterface.
func (categoryRepo *categoryRepository) DeleteCategory(id uint64) (err error) {
	var checkId model.Categories

	errData := categoryRepo.db.Where("id = ?", id).Delete(&checkId)
	if errData != nil {
		return errData.Error
	}

	if errData.RowsAffected == 0 {
		return errors.New("data not found")
	}

	return nil
}

// ReadAllCategory implements entity.CategoryDataInterface.
func (categoryRepo *categoryRepository) ReadAllCategory() ([]categories.CategoryCore, error) {
	var dataCategories []model.Categories

	errData := categoryRepo.db.Find(&dataCategories).Error
	if errData != nil {
		return nil, errData
	}

	mapData := make([]categories.CategoryCore, len(dataCategories))
	for i, value := range dataCategories {
		mapData[i] = categories.CategoryCore{
			ID:        uint64(value.ID),
			Category:  value.Category,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
		}
	}

	return mapData, nil
}