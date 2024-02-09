package entity

type ProductDataInterface interface {
	PostProduct(data ProductCore) (dataInput ProductCore, err error)
	ReadAllProduct() ([]ProductCore, error)
	GetProductByID(id string) (ProductCore, error)
	ReadAllProductByCategory(categoryId int) ([]ProductCore, error)
	DeleteProduct(id string) (err error)
}

type ProductUseCaseInterface interface {
	PostProduct(data ProductCore) (dataInput ProductCore, err error)
	ReadAllProduct() ([]ProductCore, error)
	GetProductByID(id string) (ProductCore, error)
	ReadAllProductByCategory(categoryId int) ([]ProductCore, error)
	DeleteProduct(id string) (err error)
}
