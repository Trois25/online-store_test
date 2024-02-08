package route

import (
	"store/features/categories/handler"
	"store/features/categories/repository"
	"store/features/categories/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RouteCategory(db *gorm.DB, e *echo.Group) {
	categoryRepository := repository.New(db)
	categoryUsecase := service.New(categoryRepository)
	categoryController := handler.New(categoryUsecase)

	e.POST("", categoryController.CreateCategory)
	e.GET("", categoryController.ReadAllCategory)
	e.DELETE("/:id", categoryController.DeleteCategory)
}
