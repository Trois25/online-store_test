package route

import (
	"store/features/orders/handler"
	"store/features/orders/repository"
	"store/features/orders/service"
	middlewares "store/utils/jwt"

	cartRepository "store/features/carts/repository"
	productRepository "store/features/products/repository"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RouteOrder(db *gorm.DB, e *echo.Group) {

	productRepo := productRepository.New(db)
	cartRepo := cartRepository.NewCartRepository(db, productRepo)
	orderRepository := repository.NewOrderRepository(db, cartRepo, productRepo)
	orderUsecase := service.NewOrderService(orderRepository)
	orderController := handler.NewOrderHandler(orderUsecase)

	e.POST("/:id", orderController.AddOrder, middlewares.JWTMiddleware())
}
