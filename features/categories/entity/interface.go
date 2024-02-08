package entity

type CategoryDataInterface interface {
	CreateCategory(data CategoryCore) (err error)
	ReadAllCategory() ([]CategoryCore, error)
	DeleteCategory(id uint64) (err error)
}

type CategoryUseCaseInterface interface {
	CreateCategory(data CategoryCore) (err error)
	ReadAllCategory() ([]CategoryCore, error)
	DeleteCategory(id uint64) (err error)
}