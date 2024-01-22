package models

type Order struct {
	ID    uint64  `json:"id" uri:"id"`
	Price float64 `json:"price" form:"price"`
}
