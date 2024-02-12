package entity

import "time"

type OrdersCore struct {
	ID            string
	UserId        string
	CartId        string
	TotalPrice    int
	PaymentStatus string
	UpdatedAt     time.Time
}
