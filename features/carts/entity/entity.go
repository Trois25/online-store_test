package entity

type CartsCore struct {
	ID         string
	CustomerId string
	ProductId  string
	Products   CartProductCore
	Quantity   int
}

type CartProductCore struct {
	Product   string
	Price     int
}
