package entity

import "time"

type CategoryCore struct {
	ID            uint64    `json:"id"`
	Category string    `json:"category_name"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
