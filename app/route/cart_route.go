package route

import (
	"store/features/carts/handler"
	"store/features/carts/repository"
	"store/features/carts/service"
	middlewares "store/utils/jwt"

	productRepository "store/features/products/repository"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RouteCart(db *gorm.DB, e *echo.Group) {

	productRepo := productRepository.New(db)
	cartRepository := repository.NewCartRepository(db, productRepo)
	cartUsecase := service.NewCartService(cartRepository)
	cartController := handler.NewCartHandler(cartUsecase)

	e.POST("", cartController.AddCartProduct, middlewares.JWTMiddleware())
	e.GET("", cartController.GetAllCartProduct, middlewares.JWTMiddleware())
	e.DELETE("/:id", cartController.DeleteCartProduct, middlewares.JWTMiddleware())
}
