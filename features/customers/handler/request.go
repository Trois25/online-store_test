package handler

type CustomerRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
	Name string `json:"name"`
	Address string `json:"address"`
}