package db

import (
	"go-apiserver-template/internal/db/entities"

	"gorm.io/gorm"
)

// Migrate creates or updates the database tables.
// Append new database entities here, database migragtions will be handled
// on the application startup.
func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&entities.User{},
		&entities.Order{},
	)
}
