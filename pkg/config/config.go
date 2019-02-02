package config

import (
	"github.com/kelseyhightower/envconfig"
)

type configKey struct{}

// Config
type Config struct {
	Debug             bool
	TelemetryPort     int    `envconfig:"TELEMETRY_PORT" default:"9116"`
	TelemetryEndpoint string `envconfig:"TELEMETRY_ENDPOINT" default:"/metrics"`
	ClickHouseDSN     string `envconfig:"CLICKHOUSE_DSN" required:"true"`
}

// NewContext creates a config from env.
func NewFromEnv() (*Config, error) {
	cfg := &Config{}
	err := envconfig.Process("", cfg)
	return cfg, err
}