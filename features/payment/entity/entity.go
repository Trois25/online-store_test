package entity

type PaymentCore struct {
	ID           string
	UserId       string
	OrderId      string
	Payment      string
	Number       string
	PaymentTotal int
}
