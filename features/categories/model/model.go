package model

import "time"

type Categories struct {
	ID        int       `gorm:"int;primaryKey;not null" json:"id"`
	Category  string    `gorm:"varchar(50);not null" json:"category"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
}
