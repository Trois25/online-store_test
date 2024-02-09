package repository

import (
	"errors"
	products "store/features/products/entity"
	"store/features/products/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) products.ProductDataInterface {
	return &productRepository{
		db: db,
	}
}

// CreateProduct implements entity.ProductDataInterface.
func (productRep *productRepository) PostProduct(data products.ProductCore) (dataInput products.ProductCore, err error) {
	newUUID, UUIDerr := uuid.NewRandom()
	if UUIDerr != nil {
		return products.ProductCore{}, UUIDerr
	}

	var input = model.Products{
		ID:         newUUID,
		Product:    data.Product,
		Price:      data.Price,
		CategoryId: data.CategoryId,
	}

	errData := productRep.db.Save(&input)
	if errData != nil {
		return products.ProductCore{}, errData.Error
	}

	return dataInput, nil
}

// DeleteUser implements entity.ProductDataInterface.
func (productRep *productRepository) DeleteProduct(id string) (err error) {
	var checkId model.Products

	errData := productRep.db.Where("id = ?", id).Delete(&checkId)
	if errData != nil {
		return errData.Error
	}

	if errData.RowsAffected == 0 {
		return errors.New("data not found")
	}

	return nil
}

// ReadAllProduct implements entity.ProductDataInterface.
func (productRep *productRepository) ReadAllProductByCategory(categoryId int) ([]products.ProductCore, error) {
	var dataProduct []model.Products

	errData := productRep.db.Where("category_id = ?", categoryId).Find(&dataProduct).Error
	if errData != nil {
		return nil, errData
	}

	mapData := make([]products.ProductCore, len(dataProduct))
	for i, value := range dataProduct {
		mapData[i] = products.ProductCore{
			ID:         value.ID.String(),
			Product:    value.Product,
			Price:      value.Price,
			CategoryId: value.CategoryId,
		}
	}
	return mapData, nil
}

// ReadAllProduct implements entity.ProductDataInterface.
func (productRep *productRepository) ReadAllProduct() ([]products.ProductCore, error) {
	var dataProduct []model.Products

	errData := productRep.db.Find(&dataProduct).Error
	if errData != nil {
		return nil, errData
	}

	mapData := make([]products.ProductCore, len(dataProduct))
	for i, value := range dataProduct {
		mapData[i] = products.ProductCore{
			ID:         value.ID.String(),
			Product:    value.Product,
			Price:      value.Price,
			CategoryId: value.CategoryId,
		}
	}
	return mapData, nil
}

// GetProductByID implements entity.ProductDataInterface.
func (productRepo *productRepository) GetProductByID(id string) (products.ProductCore, error) {
	var product model.Products

	err := productRepo.db.Where("id = ?", id).First(&product).Error
	if err != nil {
		return products.ProductCore{}, err
	}

	ProductCore := products.ProductCore{
		Product: product.Product,
		Price:   product.Price,
	}

	return ProductCore, nil
}
