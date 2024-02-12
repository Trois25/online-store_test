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

func (handler *orderHandler) GetSpecificCart(e echo.Context) error {

	Id := middlewares.ExtractTokenUserId(e)
	if Id == "" {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to get id",
		})
	}
	orderId := e.Param("id")

	data, err := handler.orderService.GetSpecificOrder(Id, orderId)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get order",
		})
	}

	return e.JSON(http.StatusOK, map[string]any{
		"message": "get order",
		"data":    data,
	})
}

func (handler *orderHandler) EditOrder(e echo.Context) error {
	userId := middlewares.ExtractTokenUserId(e)

	if userId == "" {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get userId",
		})
	}

	orderId := e.Param("id")

	err := handler.orderService.EditOrder(userId, orderId)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error updating event",
			"error":   err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Payment Status updated",
	})
}