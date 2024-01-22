package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewMySQL creates a new MySQL database connection.
func NewMySQL(dns string) (*gorm.DB, error) {
	return newDB(mysql.New(mysql.Config{
		DSN:                      dns,
		DefaultStringSize:        256,
		DisableDatetimePrecision: true,
	}))
}
