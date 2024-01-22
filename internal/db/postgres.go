package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewPostgres creates a new PostgreSQL database connection.
func NewPostgres(dns string) (*gorm.DB, error) {
	return newDB(postgres.New(postgres.Config{
		DSN: dns,
	}))
}
