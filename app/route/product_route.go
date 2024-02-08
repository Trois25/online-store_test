package route

import (
	"store/features/products/handler"
	"store/features/products/repository"
	"store/features/products/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RouteProduct(db *gorm.DB, e *echo.Group) {
	productRepository := repository.New(db)
	productUsecase := service.New(productRepository)
	productController := handler.New(productUsecase)

	e.POST("", productController.PostProduct)
	e.GET("/:category_id", productController.ReadAllProductByCategory)
	e.DELETE("/:id", productController.DeleteProduct)
}
