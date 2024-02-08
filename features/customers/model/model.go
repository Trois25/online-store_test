package model

import (
	"time"

	"github.com/google/uuid"
)

type Customers struct {
	ID        uuid.UUID `gorm:"type:varchar(50);primaryKey;not null" json:"id"`
	Name      string    `gorm:"varchar(50);not null" json:"name"`
	Email     string    `gorm:"varchar(50);not null" json:"email"`
	Password  string    `gorm:"varchar(50);not null" json:"password"`
	Address   string    `gorm:"varchar(50);not null" json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
}
