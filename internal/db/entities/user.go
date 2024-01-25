package entities

type User struct {
	IDBase
	Name     string `json:"name" gorm:"not null"`
	Password string `json:"password" gorm:"size:255;not null"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	AuditBase
}

func (User) TableName() string {
	return "users"
}
