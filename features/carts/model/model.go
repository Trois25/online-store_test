package model

import (
	"store/features/products/model"

	"github.com/google/uuid"
)

type Carts struct {
	ID         uuid.UUID        `gorm:"type:varchar(50);primaryKey;not null" json:"id"`
	CustomerId string           `gorm:"type:varchar(50);not null" json:"customer_id"`
	ProductId  string           `gorm:"type:varchar(50);not null" json:"product_id"`
	Products   []model.Products `gorm:"many2many:CartProduct"`
	Quantity   int              `json:"quantity"`
}

type CartProduct struct {
	ProductID string
	CartID    string
}
