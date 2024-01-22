package global

const (
	CurrentVersion = "v1.0.0"

	ConfigFileEnvKey = "CONFIG_FILE"
	RunModeEnvKey    = "RUN_MODE"

	RunModeDev  = "dev"
	RunModeProd = "prod"

	HealthzRelativePath = "/healthz"

	DefaultTimeFormat = "2006-01-02 15:04:05"
)

type DBType string

const (
	DBTypeMySQL     = "mysql"
	DBTypePostgres  = "postgres"
	DBTypeSQLite    = "sqlite"
	DBTypeSQLServer = "sqlserver"
)
