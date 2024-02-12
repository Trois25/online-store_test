package model

import (
	"time"

	"github.com/google/uuid"
)

type Orders struct {
	ID            uuid.UUID `gorm:"type:varchar(50);primaryKey;not null" json:"id"`
	UserId        string    `gorm:"type:varchar(50);not null" json:"user_id"`
	TotalPrice    int       `gorm:"type:varchar(50);not null" json:"total_price"`
	CartId        string    `gorm:"type:varchar(50);not null" json:"cart_id"`
	PaymentStatus string    `gorm:"type:enum('Paid','Unpaid');default:'Unpaid'"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"update_at"`
}
