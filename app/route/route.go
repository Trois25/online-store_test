package route

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func New(e *echo.Echo, db *gorm.DB) {
	loginPage := e.Group("/")
	category := e.Group("/category")
	product := e.Group("/product")
	cart := e.Group("/cart")
	order := e.Group("/order")
	payment := e.Group("/pay")

	RouteCustomer(db, loginPage)
	RouteCategory(db, category)
	RouteProduct(db, product)
	RouteCart(db, cart)
	RouteOrder(db, order)
	RoutePayment(db, payment)
}
