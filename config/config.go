package config

import "github.com/kelseyhightower/envconfig"

// FreelanceConfig ...
type FreelanceConfig struct {
	DatabaseURL string `envconfig:"DATABASE_URL"`
	DbType      string `envconfig:"DB_TYPE"`
	DoMigration string `envconfig:"DO_MIGRATION"`
}

// GetConfig ...
func GetConfig() (*FreelanceConfig, error) {
	var c FreelanceConfig
	err := envconfig.Process("freelance", &c)
	return &c, err
}
