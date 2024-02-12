package model

import (
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	ID           uuid.UUID `gorm:"type:varchar(50);primaryKey;not null" json:"id"`
	UserId       string    `gorm:"type:varchar(50);primaryKey;not null" json:"user_id"`
	OrderId      string    `gorm:"type:varchar(50);primaryKey;not null" json:"order_id"`
	Payment      string    `gorm:"not null" json:"payment"`
	Number       string    `gorm:"not null" json:"number"`
	PaymentTotal int       `gorm:"not null" json:"payment_total"`
	CreatedAt    time.Time `json:"created_at"`
}
