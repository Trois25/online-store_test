package entity

type CartRepositoryInterface interface {
	AddCartProduct(cartInput CartsCore, userId string) error
	GetAllCartProduct(userId string) ([]CartsCore, error)
	DeleteCartProduct(id string, userId string) error
}

type CartServiceInterface interface {
	AddCartProduct(cartInput CartsCore, userId string) error
	GetAllCartProduct(userId string) ([]CartsCore, error)
	DeleteCartProduct(id string, userId string) error
}
