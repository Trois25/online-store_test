package entity

type OrderRepositoryInterface interface {
	AddOrder(userId string, cartId string) error
	GetSpecificOrder(userId string, orderId string) (OrdersCore, error)
	EditOrder(userId string, orderId string) error
}

type OrderServiceInterface interface {
	AddOrder(userId string, cartId string) error
	GetSpecificOrder(userId string, orderId string) (OrdersCore, error)
	EditOrder(userId string, orderId string) error
}
