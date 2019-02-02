package config

import (
	"context"

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

// NewContext creates a context with config.
func NewContext(ctx context.Context, cfg *Config) context.Context {
	return context.WithValue(ctx, configKey{}, cfg)
}

// FromContext returns the config from context.
func FromContext(ctx context.Context) (*Config, bool) {
	c, ok := ctx.Value(configKey{}).(*Config)
	return c, ok
}
