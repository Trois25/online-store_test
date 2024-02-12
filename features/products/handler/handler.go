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

func (handler *ProductController) PostProduct(e echo.Context) error {
	input := new(ProductRequest)
	errBind := e.Bind(&input)
	if errBind != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
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
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error post product",
			"error":   errProduct.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]any{
		"message": "success post product",
		"data":    data,
	})
}

func (handler *ProductController) ReadAllProductByCategory(e echo.Context) error {
	categoryId := e.Param("category_id")

	id, _ := strconv.Atoi(categoryId)
	data, err := handler.productUseCase.ReadAllProductByCategory(id)
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

func (handler *ProductController) ReadAllProduct(e echo.Context) error {
	data, err := handler.productUseCase.ReadAllProduct()
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

func (handler *ProductController) DeleteProduct(e echo.Context) error {
	idParams := e.Param("id")
	err := handler.productUseCase.DeleteProduct(idParams)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error deleting product",
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "product deleted successfully",
	})
}

func (handler *ProductController) GetProductByID(e echo.Context) error {
	Id := e.Param("id")

	data, err := handler.productUseCase.GetProductByID(Id)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get product",
		})
	}

	return e.JSON(http.StatusOK, map[string]any{
		"message": "get product",
		"data":    data,
	})
}