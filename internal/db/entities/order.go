package entities

type Order struct {
	IDBase
	AuditBase
	UserID uint64  `json:"user_id" gorm:"not null"`
	Price  float64 `json:"price" gorm:"not null"`
}

func (Order) TableName() string {
	return "orders"
}
