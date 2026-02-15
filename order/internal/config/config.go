package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	App AppConfig
	DB  struct {
		ConnectionString string `envconfig:"DSN" required:"true"`
		MigrationPath    string `envconfig:"MIGRATION_PATH" default:"/app/migrations"`
	}
}

type AppConfig struct {
	Port        string `envconfig:"SERVER_PORT" default:"8080"`
	IsDebugMode bool   `envconfig:"DEBUG" default:"true"`
	Environment string `envconfig:"ENVIRONMENT" default:"dev"`
}

func NewConfig() (*Config, error) {
	const op = "config.NewConfig"

	var config Config
	var err error

	// App config
	if err = envconfig.Process("", &config.App); err != nil {
		return nil, fmt.Errorf("%s: failed to parse App config: %w", op, err)
	}
	if err = envconfig.Process("", &config.DB); err != nil {
		return nil, fmt.Errorf("%s: failed to parse DB config: %w", op, err)
	}

	return &config, nil
}
