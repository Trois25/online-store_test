package model

import (
	categories "store/features/categories/model"

	"time"

	"github.com/google/uuid"
)

type Products struct {
	ID         uuid.UUID `gorm:"varchar(50);primaryKey;not null" json:"id"`
	Product    string    `gorm:"varchar(50);not null" json:"product"`
	CategoryId string    `gorm:"varchar(50);not null" json:"category_id"`
	Price      int       `gorm:"varchar(50);not null" json:"price"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"update_at"`
	Category   categories.Categories
}
