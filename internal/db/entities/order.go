package entities

type Order struct {
	UUIDBase
	UserID uint64  `json:"user_id" gorm:"not null"`
	Price  float64 `json:"price" gorm:"not null"`
	AuditBase
}

func (Order) TableName() string {
	return "orders"
}
