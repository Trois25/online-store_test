package route

import (
	"store/features/payment/handler"
	"store/features/payment/repository"
	"store/features/payment/service"
	middlewares "store/utils/jwt"

	orderRepository "store/features/orders/repository"
	cartRepository "store/features/carts/repository"
	productRepository "store/features/products/repository"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RoutePayment(db *gorm.DB, e *echo.Group) {

	productRepo := productRepository.New(db)
	cartRepo := cartRepository.NewCartRepository(db, productRepo)
	orderRepo := orderRepository.NewOrderRepository(db, cartRepo, productRepo)
	paymentRepository := repository.NewPaymentRepository(db, orderRepo)
	paymentUsecase := service.NewPaymentService(paymentRepository)
	paymentController := handler.NewPaymentHandler(paymentUsecase)

	e.PUT("/:id", paymentController.PayOrder, middlewares.JWTMiddleware())
}
