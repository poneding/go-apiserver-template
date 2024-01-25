package models

type User struct {
	ID     uint64  `json:"id" uri:"id"`
	Name   string  `json:"name" form:"name" binding:"required"`
	Orders []Order `json:"orders" form:"orders"`
}
