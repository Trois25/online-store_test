package handler

type ProductRequest struct {
	Product    string `json:"product"`
	CategoryId int    `json:"category_id"`
	Price      int    `json:"price"`
}
