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
