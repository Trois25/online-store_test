package route

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func New(e *echo.Echo, db *gorm.DB) {
	loginPage := e.Group("/")
	category := e.Group("/category")

	RouteCustomer(db, loginPage)
	RouteCategory(db, category)
}
