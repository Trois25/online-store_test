package database

import (
	carts "store/features/carts/model"
	categories "store/features/categories/model"
	customers "store/features/customers/model"
	orders "store/features/orders/model"
	products "store/features/products/model"

	"fmt"
	"store/app/configs"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDBMysql(cfg *configs.AppConfig) *gorm.DB {

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.DBUSER, cfg.DBPASS, cfg.DBHOST, cfg.DBPORT, cfg.DBNAME)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}

func InitMigration(db *gorm.DB) {
	db.AutoMigrate(&carts.Carts{})
	db.AutoMigrate(categories.Categories{})
	db.AutoMigrate(customers.Customers{})
	db.AutoMigrate(orders.Orders{})
	db.AutoMigrate(products.Products{})
}
