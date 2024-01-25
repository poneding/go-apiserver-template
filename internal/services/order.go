package services

import (
	"go-apiserver-template/internal/db"
	"go-apiserver-template/internal/db/entities"
)

type OrderServiceInterface interface {
	GetOrder(oid string) (*entities.Order, error)
}

func NewOrderService() OrderServiceInterface {
	return &orderService{}
}

type orderService struct {
}

func (s *orderService) GetOrder(oid string) (*entities.Order, error) {
	var order entities.Order
	if err := db.Repository().Where("id = ?", oid).First(&order).Error; err != nil {
		return nil, err
	}

	return &order, nil
}
