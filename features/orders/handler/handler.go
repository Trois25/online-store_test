package handler

import (
	"net/http"
	"store/features/orders/entity"
	middlewares "store/utils/jwt"

	"github.com/labstack/echo/v4"
)

type orderHandler struct {
	orderService entity.OrderServiceInterface
}

func NewOrderHandler(order entity.OrderServiceInterface) *orderHandler {
	return &orderHandler{
		orderService: order,
	}
}

func (handler *orderHandler) AddOrder(e echo.Context) error {
	Id := middlewares.ExtractTokenUserId(e)
	if Id == "" {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to get id",
		})
	}

	cartId := e.Param("id")

	orderInput := entity.OrdersCore{
		CartId: cartId,
	}

	errAdd := handler.orderService.AddOrder(Id, orderInput.CartId)
	if errAdd != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed make order",
			"error":   errAdd.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "order success",
	})
}
