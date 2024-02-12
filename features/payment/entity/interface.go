package entity

type PaymentRepositoryInterface interface {
	PayOrder(userId string, orderId string, data PaymentCore) error
}

type PaymentServiceInterface interface {
	PayOrder(userId string, orderId string, data PaymentCore) error
}