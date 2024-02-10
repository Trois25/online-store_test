package entity

type OrderRepositoryInterface interface {
	AddOrder(userId string, cartId string) error
}

type OrderServiceInterface interface {
	AddOrder(userId string, cartId string) error
}
