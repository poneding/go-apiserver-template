package db

import (
	"go-apiserver-template/internal/config"
	"go-apiserver-template/internal/global"
	"go-apiserver-template/pkg/log"
	"os"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// newDB is a helper function to create a new *gorm.DB instance with
// default configurations
func newDB(dialector gorm.Dialector) (*gorm.DB, error) {
	loglevel := logger.Error
	if os.Getenv(global.RunModeEnvKey) == global.RunModeDev {
		loglevel = logger.Info
	}

	return gorm.Open(dialector, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		},
		Logger: logger.Default.LogMode(loglevel),
	})
}

// instance returns a singleton instance of the database
func instance() *gorm.DB {
	var result *gorm.DB
	var err error
	switch config.Instance().DB.Type {
	case global.DBTypeMySQL:
		result, err = NewMySQL(config.Instance().DB.DNS)
	case global.DBTypePostgres:
		result, err = NewPostgres(config.Instance().DB.DNS)
	case global.DBTypeSQLite:
		result, err = NewSQLite(config.Instance().DB.DNS)
	case global.DBTypeSQLServer:
		result, err = NewSQLServer(config.Instance().DB.DNS)
	default:
		log.Fatalf("unsupported database type, %v", config.Instance().DB.Type)
	}

	if err != nil {
		log.Fatalf("error connecting to database, %v", err)
	}

	// Migrate the database
	Migrate(result)

	return result
}
