package handler

type PaymentRequest struct {
	Payment      string    `json:"payment"`
	Number       string    `json:"number"`
	PaymentTotal int       `json:"payment_total"`
}
