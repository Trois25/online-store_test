package handler

import (
	"net/http"
	categories "store/features/categories/entity"
	"strconv"

	"github.com/labstack/echo/v4"
)

type categoryController struct {
	categoryUseCase categories.CategoryUseCaseInterface
}

func New(categoryUC categories.CategoryUseCaseInterface) *categoryController {
	return &categoryController{
		categoryUseCase: categoryUC,
	}
}

func (handler *categoryController) CreateCategory(e echo.Context) error {

	input := new(CategoryRequest)
	errBind := e.Bind(&input)
	if errBind != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data",
		})
	}

	data := categories.CategoryCore{
		Category: input.Category,
	}

	errcategory := handler.categoryUseCase.CreateCategory(data)
	if errcategory != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error create category",
			"error":   errcategory.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]any{
		"message": "success create category",
	})
}

func (handler *categoryController) ReadAllCategory(e echo.Context) error {

	data, err := handler.categoryUseCase.ReadAllCategory()
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get all category",
		})
	}

	return e.JSON(http.StatusOK, map[string]any{
		"message": "get all category",
		"data":    data,
	})
}

func (handler *categoryController) DeleteCategory(e echo.Context) error {

	idParams := e.Param("id")
	if idParams == "" {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "id can't empty",
		})
	}

	inputId, errParse := strconv.ParseUint(idParams, 10, 64)
	if errParse != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed delete role",
		})
	}

	err := handler.categoryUseCase.DeleteCategory(inputId)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed delete category",
		})
	}

	return e.JSON(http.StatusOK, map[string]any{
		"message": "success delete data",
	})
}
