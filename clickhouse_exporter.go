package main

import (
	"fmt"

	"github.com/hot-wifi/clickhouse_exporter/pkg/collector"
	"github.com/hot-wifi/clickhouse_exporter/pkg/config"
	"github.com/hot-wifi/clickhouse_exporter/pkg/handler"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/prometheus/client_golang/prometheus"

	_ "github.com/kshvakov/clickhouse"
)

func main() {
	e := echo.New()

	cfg, err := config.NewFromEnv()
	if err != nil {
		e.Logger.Fatalf("failed to init config: %v", err)
	}

	db, err := sqlx.Connect("clickhouse", cfg.ClickHouseDSN)
	if err != nil {
		e.Logger.Fatalf("failed to connect db: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			e.Logger.Fatalf("failed to close db: %v", err)
		}
	}()

	e.Debug = cfg.Debug
	e.HideBanner = !e.Debug

	e.Use(middleware.Recover())

	if cfg.Debug {
		e.Use(middleware.Logger())
	}

	c := collector.NewClickHouseCollector(db)

	if err := prometheus.Register(c); err != nil {
		e.Logger.Fatalf("failed to register collector: %v", err)
	}

	h := handler.NewHandler(db)

	e.GET(cfg.TelemetryEndpoint, h.MetricsHandler())

	e.GET("/_ah/health", h.HealthCheck)

	addr := fmt.Sprintf("[::]:%d", cfg.TelemetryPort)

	e.Logger.Fatal(e.Start(addr))
}
