package service

import (
	"errors"
	"store/features/orders/entity"
)

type orderService struct {
	orderRepository entity.OrderRepositoryInterface
}

func NewOrderService(order entity.OrderRepositoryInterface) entity.OrderServiceInterface {
	return &orderService{
		orderRepository: order,
	}
}

// addOrder implements entity.OrderServiceInterface.
func (order *orderService) AddOrder(userId string, cartId string) error {
	if userId == "" {
		return errors.New("you need to login first")
	}

	if cartId == "" {
		return errors.New("item you order is not available")
	}

	errAdd := order.orderRepository.AddOrder(userId, cartId)
	if errAdd != nil {
		return errAdd
	}

	return nil
}

// GetSpecificOrder implements entity.OrderServiceInterface.
func (order *orderService) GetSpecificOrder(userId string, orderId string) (entity.OrdersCore, error) {
	if orderId == "" {
		return entity.OrdersCore{}, errors.New("id invalid")
	}

	if userId == "" {
		return entity.OrdersCore{}, errors.New("you need to login first")
	}

	orderData, err := order.orderRepository.GetSpecificOrder(userId, orderId)
	if err != nil {
		return entity.OrdersCore{}, errors.New("can't read data")
	}

	return orderData, nil
}

// EditOrder implements entity.OrderServiceInterface.
func (order *orderService) EditOrder(userId string, orderId string) error {
	if orderId == "" {
		return errors.New("error, Purchase ID is required")
	}

	errUpdate := order.orderRepository.EditOrder(userId, orderId)
	if errUpdate != nil {
		return errUpdate
	}

	return nil
}
