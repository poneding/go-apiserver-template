package config

import (
	"go-apiserver-template/internal/global"
	"log"
	"os"

	"github.com/spf13/viper"
)

// Config is the struct for config file
type Config struct {
	DB DBConfig `yaml:"db"`
}

// DBConfig is the struct for database config
type DBConfig struct {
	Type global.DBType `yaml:"type"`
	DNS  string        `yaml:"dns"`
}

// configInstance is the singleton instance of the config
var configInstance *Config

// Instance returns a singleton instance of the config
func Instance() *Config {
	if configInstance == nil {
		configPath := os.Getenv(global.ConfigFileEnvKey)
		if configPath == "" {
			configPath = "config.yaml"
		}

		viper.SetConfigFile(configPath)

		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("error reading config file, %v", err)
		}

		if err := viper.Unmarshal(&configInstance); err != nil {
			log.Fatalf("unable to decode config into struct, %v", err)
		}
	}

	return configInstance
}
