package models

type Order struct {
	ID    string  `json:"id" uri:"id"`
	Price float64 `json:"price" form:"price"`
}
