package service

import (
	"errors"
	products "store/features/products/entity"
)

type productUseCase struct {
	productRepository products.ProductDataInterface
}

func New(ProductUseCase products.ProductDataInterface) products.ProductUseCaseInterface {
	return &productUseCase{
		productRepository: ProductUseCase,
	}
}

// CreateProduct implements entity.ProductUseCaseInterface.
func (productUC *productUseCase) PostProduct(data products.ProductCore) (dataInput products.ProductCore,err error) {
	if data.Product == "" || data.Price == 0 {
		return products.ProductCore{},errors.New("error, Product name or price can't be empty")
	}

	if data.Price < 0 {
		return products.ProductCore{},errors.New("error, Price must be a positive integer")
	}
	
	data,errProduct := productUC.productRepository.PostProduct(data)
	if errProduct != nil{
		return products.ProductCore{},errProduct
	}

	return dataInput,nil
}

// DeleteUser implements entity.ProductUseCaseInterface.
func (productUC *productUseCase) DeleteProduct(id string) (err error) {
	if id == "" {
		return errors.New("product id can't empty")
	}

	errPurchase := productUC.productRepository.DeleteProduct(id)
	if errPurchase != nil {
		return errors.New("can't delete product")
	}

	return nil
}

// ReadAllProduct implements entity.ProductUseCaseInterface.
func (productUC *productUseCase) ReadAllProductByCategory(categoryId int) ([]products.ProductCore, error) {
	products, err := productUC.productRepository.ReadAllProductByCategory(categoryId)
	if err != nil {
		return nil, errors.New("error get data")
	}
	return products, nil
}
