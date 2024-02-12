package repository

import (
	"errors"
	ordere "store/features/orders/entity"
	"store/features/payment/entity"
	"store/features/payment/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type paymentRepository struct {
	db    *gorm.DB
	order ordere.OrderServiceInterface
}

func NewPaymentRepository(db *gorm.DB, order ordere.OrderServiceInterface) entity.PaymentRepositoryInterface {
	return &paymentRepository{
		db:    db,
		order: order,
	}
}

// PayOrder implements entity.PaymentRepositoryInterface.
func (paymentRepo *paymentRepository) PayOrder(userId string, orderId string, data entity.PaymentCore) error {

	// Mendapatkan data order spesifik berdasarkan orderId
	orderData, errCart := paymentRepo.order.GetSpecificOrder(userId, orderId)
	if errCart != nil {
		return errCart
	}

	newUUID, UUIDerr := uuid.NewRandom()
	if UUIDerr != nil {
		return UUIDerr
	}

	var paymentData = model.Payment{
		ID : newUUID,
		UserId: userId,
		OrderId: orderId,
		Payment: data.Payment,
		Number: data.Number,
		PaymentTotal: data.PaymentTotal,
	}

	if orderData.TotalPrice == data.PaymentTotal {

		paymentRepo.order.EditOrder(userId, orderId)

		tx := paymentRepo.db.Save(&paymentData)
		if tx.Error != nil{
			return tx.Error
		}
		return nil
	}

	return errors.New("payment total is not true")

}
