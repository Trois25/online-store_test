package route

import (
	"store/features/customers/handler"
	"store/features/customers/repository"
	"store/features/customers/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RouteCustomer(db *gorm.DB, e *echo.Group) {
	customerRepository := repository.New(db)
	customerUsecase := service.New(customerRepository)
	customerController := handler.New(customerUsecase)

	e.POST("register", customerController.Register)
	e.POST("login", customerController.Login)
}
