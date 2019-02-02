package handler

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Handler
type Handler struct {
	db *sqlx.DB
}

// NewHandler
func NewHandler(db *sqlx.DB) *Handler {
	return &Handler{db}
}

// MetricsHandler
func (h *Handler) MetricsHandler() echo.HandlerFunc {
	return echo.WrapHandler(promhttp.Handler())
}

// HealthCheck
func (h *Handler) HealthCheck(c echo.Context) error {
	if err := h.db.Ping(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.NoContent(http.StatusNoContent)
}
