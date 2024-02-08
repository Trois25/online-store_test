package handler

import (
	"net/http"
	products "store/features/products/entity"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductController struct {
	productUseCase products.ProductUseCaseInterface
}

func New(productUC products.ProductUseCaseInterface) *ProductController {
	return &ProductController{
		productUseCase: productUC,
	}
}

func (handler *ProductController) PostProduct(c echo.Context) error {
	input := new(ProductRequest)
	errBind := c.Bind(&input)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data",
		})
	}

	data := products.ProductCore{
		Product:    input.Product,
		CategoryId: input.CategoryId,
		Price:      input.Price,
	}

	_, errProduct := handler.productUseCase.PostProduct(data)
	if errProduct != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error post product",
			"error":   errProduct.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success post product",
		"data":    data,
	})
}

func (handler *ProductController) ReadAllProductByCategory(c echo.Context) error {
	categoryId := c.Param("category_id")

	id, _ := strconv.Atoi(categoryId)
	data, err := handler.productUseCase.ReadAllProductByCategory(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get all product",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "get all product",
		"data":    data,
	})
}

func (handler *ProductController) DeleteProduct(c echo.Context) error {
	idParams := c.Param("id")
	err := handler.productUseCase.DeleteProduct(idParams)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error deleting product",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "product deleted successfully",
	})
}
