package handler

import (
	"net/http"
	"store/features/carts/entity"
	middlewares "store/utils/jwt"

	"github.com/labstack/echo/v4"
)

type cartHandler struct {
	cartService entity.CartServiceInterface
}

func NewCartHandler(cart entity.CartServiceInterface) *cartHandler {
	return &cartHandler{
		cartService: cart,
	}
}

func (handler *cartHandler) AddCartProduct(e echo.Context) error {
	Id := middlewares.ExtractTokenUserId(e)
	if Id == "" {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to get id",
		})
	}

	newCart := new(CartRequest)
	errBind := e.Bind(&newCart)
	if errBind != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data",
		})
	}

	cartInput := entity.CartsCore{
		ProductId: newCart.ProductId,
		Quantity:  newCart.Quantity,
	}

	errAdd := handler.cartService.AddCartProduct(cartInput, Id)
	if errAdd != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed add product",
			"error":   errAdd.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "product added to cart",
	})
}

func (handler *cartHandler) DeleteCartProduct(e echo.Context) error {
	userId := middlewares.ExtractTokenUserId(e)
	if userId == "" {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to get id",
		})
	}

	idParams := e.Param("id")
	err := handler.cartService.DeleteCartProduct(idParams, userId)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error deleting product",
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "product deleted successfully",
	})
}

func (handler *cartHandler) GetAllCartProduct(e echo.Context) error {
	userId := middlewares.ExtractTokenUserId(e)
	if userId == "" {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to get id",
		})
	}

	data, err := handler.cartService.GetAllCartProduct(userId)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get all product",
		})
	}

	return e.JSON(http.StatusOK, map[string]any{
		"message": "get all product",
		"data":    data,
	})
}
