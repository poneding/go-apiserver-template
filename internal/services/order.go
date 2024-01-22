package services

import (
	"go-apiserver-template/internal/db"
	"go-apiserver-template/internal/db/entities"
)

type OrderServiceInterface interface {
	GetUserOrder(uid, oid uint64) (*entities.Order, error)
}

func NewOrderService() OrderServiceInterface {
	return &orderService{}
}

type orderService struct {
}

func (s *orderService) GetUserOrder(uid, oid uint64) (*entities.Order, error) {
	var order entities.Order
	if err := db.Repository().Where("user_id = ? AND id = ?", uid, oid).First(&order).Error; err != nil {
		return nil, err
	}

	return &order, nil
}
