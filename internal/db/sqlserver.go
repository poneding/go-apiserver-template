package db

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// NewSQLServer creates a new SQL Server database connection.
func NewSQLServer(dns string) (*gorm.DB, error) {
	return newDB(sqlserver.Open(dns))
}
