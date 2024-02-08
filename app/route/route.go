package route

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func New(e *echo.Echo, db *gorm.DB) {
	loginPage := e.Group("/")
	category := e.Group("/category")
	product := e.Group("/product")

	RouteCustomer(db, loginPage)
	RouteCategory(db, category)
	RouteProduct(db, product)
}
