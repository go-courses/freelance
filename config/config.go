package config

import "github.com/kelseyhightower/envconfig"

// FreelanceConfig ...
type FreelanceConfig struct {
	DatabaseURL string `envconfig:"DATABASE_URL"`
}

// GetConfig ...
func GetConfig() (*FreelanceConfig, error) {
	var c FreelanceConfig
	err := envconfig.Process("freelance", &c)
	return &c, err
}
