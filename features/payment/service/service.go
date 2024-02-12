package service

import (
	"errors"
	"store/features/payment/entity"
)

type paymentService struct {
	paymentRepository entity.PaymentRepositoryInterface
}

func NewPaymentService(payment entity.PaymentRepositoryInterface) entity.PaymentServiceInterface {
	return &paymentService{
		paymentRepository: payment,
	}
}

// PayOrder implements entity.PaymentServiceInterface.
func (payment *paymentService) PayOrder(userId string, orderId string, data entity.PaymentCore) error {
	if userId == "" {
		return errors.New("you need to login first")
	}

	if orderId == "" {
		return errors.New("item you order is not available")
	}

	errAdd := payment.paymentRepository.PayOrder(userId, orderId, data)
	if errAdd != nil {
		return errAdd
	}

	return nil
}
