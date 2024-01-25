package entities

import (
	"go-apiserver-template/pkg/util"
	"time"

	"gorm.io/gorm"
)

// IDBase is the base entity for all database entities with uint64 ID as primary key.
type IDBase struct {
	ID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
}

// UUIDBase is the base entity for all database entities with UUID as primary key.
type UUIDBase struct {
	ID string `gorm:"primary_key;column:id;size:36" json:"id"`
}

// BeforeCreate is a gorm hook to generate a new UUID before creating a new UUIDBase entity.
func (b *UUIDBase) BeforeCreate(tx *gorm.DB) error {
	b.ID = util.NewUUID(util.UUIDOption{
		RemoveHyphen: true,
	})
	return nil
}

// AuditBase is the base entity for all database entities with audit fields.
type AuditBase struct {
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	CreatedBy uint64         `gorm:"column:created_by" json:"created_by"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy uint64         `gorm:"column:updated_by" json:"updated_by"`
	DeletedAt gorm.DeletedAt `gorm:"index;column:deleted_at" json:"deleted_at"`
	DeletedBy uint64         `gorm:"column:deleted_by" json:"deleted_by"`
}
