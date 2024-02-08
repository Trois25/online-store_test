package entity

type ProductCore struct {
	ID         string    `json:"id"`
	Product    string    `json:"product"`
	CategoryId int       `json:"category_id"`
	Price      int       `json:"price"`
}
