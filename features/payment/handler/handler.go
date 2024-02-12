package handler

import (
	"net/http"
	"store/features/payment/entity"
	middlewares "store/utils/jwt"

	"github.com/labstack/echo/v4"
)

type paymentHandler struct {
	paymentService entity.PaymentServiceInterface
}

func NewPaymentHandler(payment entity.PaymentServiceInterface) *paymentHandler {
	return &paymentHandler{
		paymentService: payment,
	}
}

func (handler *paymentHandler) PayOrder(e echo.Context) error {
	Id := middlewares.ExtractTokenUserId(e)
	if Id == "" {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to get id",
		})
	}

	orderId := e.Param("id")
	dataPayment := new(PaymentRequest)
	errBind := e.Bind(&dataPayment)
	if errBind != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data",
		})
	}

	paymentInput := entity.PaymentCore{
		OrderId: orderId,
		Payment: dataPayment.Payment,
		Number: dataPayment.Number,
		PaymentTotal: dataPayment.PaymentTotal,
	}

	errPay := handler.paymentService.PayOrder(Id, orderId, paymentInput)
	if errPay != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed pay order",
			"error":   errPay.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "payment success",
	})

}