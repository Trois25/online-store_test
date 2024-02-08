package service

import (
	"errors"
	categories "store/features/categories/entity"
)

type categoryUseCase struct {
	categoryRepository categories.CategoryUseCaseInterface
}

func New(Categoryuc categories.CategoryDataInterface) categories.CategoryUseCaseInterface {
	return categoryUseCase{
		categoryRepository: Categoryuc,
	}
}

// CreateCategory implements entity.CategoryUseCaseInterface.
func (CategoryUC categoryUseCase) CreateCategory(data categories.CategoryCore) (err error) {
	if data.Category == "" {
		return errors.New("category can't be empty")
	}

	check,_ := CategoryUC.categoryRepository.ReadAllCategory()
	
	for _, existingCategory := range check {
		if existingCategory.Category == data.Category {
			return errors.New("category name already exists")
		}
	}

	errCategory := CategoryUC.categoryRepository.CreateCategory(data)
	if errCategory != nil {
		return errors.New("can't create category")
	}

	return nil
}

// DeleteCategory implements entity.CategoryUseCaseInterface.
func (CategoryUC categoryUseCase) DeleteCategory(id uint64) (err error) {
	if id == 0 {
		return errors.New("category not found")
	}

	errCategory := CategoryUC.categoryRepository.DeleteCategory(id)
	if errCategory != nil {
		return errors.New("can't delete category")
	}

	return nil
}

// ReadAllCategory implements entity.CategoryUseCaseInterface.
func (CategoryUC categoryUseCase) ReadAllCategory() ([]categories.CategoryCore, error) {
	categories, err := CategoryUC.categoryRepository.ReadAllCategory()
	if err != nil {
		return nil, errors.New("error get data")
	}

	return categories, nil
}