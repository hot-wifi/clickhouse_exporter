package collector

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/log"
)

// Prometheus metrics namespace
const metricsNamespace = "clickhouse"

// ClickHouseCollector collects clickhouse metrics.
type ClickHouseCollector struct {
	db *sqlx.DB
}

// NewClickHouseCollector returns an initialized ClickHouseCollector.
func NewClickHouseCollector(db *sqlx.DB) *ClickHouseCollector {
	return &ClickHouseCollector{db}
}

// Describe describes all the metrics.
func (c *ClickHouseCollector) Describe(ch chan<- *prometheus.Desc) {
	metricCh := make(chan prometheus.Metric)
	doneCh := make(chan struct{})

	go func() {
		for m := range metricCh {
			ch <- m.Desc()
		}
		close(doneCh)
	}()

	c.Collect(metricCh)
	close(metricCh)
	<-doneCh
}

// Collect fetches the metrics from clickhouse.
func (c *ClickHouseCollector) Collect(ch chan<- prometheus.Metric) {
	if err := c.collect(ch); err != nil {
		log.Printf("failed to collect clickhouse metrics: %s", err)
	}
}

func (c *ClickHouseCollector) collect(ch chan<- prometheus.Metric) error {
	var metrics []Metric
	err := c.db.Select(&metrics, "SELECT metric, value FROM system.metrics")
	if err != nil {
		return err
	}
	for _, m := range metrics {
		vec := prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: metricsNamespace,
			Name:      fixMetricName(m.Metric),
			Help:      fmt.Sprintf("Number of %s currently processed", m.Metric),
		}, []string{}).WithLabelValues()
		vec.Set(m.Value)
		vec.Collect(ch)
	}

	var asyncMetrics []Metric
	err = c.db.Select(&asyncMetrics, "SELECT metric, value FROM system.asynchronous_metrics")
	if err != nil {
		return err
	}
	for _, m := range asyncMetrics {
		vec := prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: metricsNamespace,
			Name:      fixMetricName(m.Metric),
			Help:      fmt.Sprintf("Number of %s async processed", m.Metric),
		}, []string{}).WithLabelValues()
		vec.Set(m.Value)
		vec.Collect(ch)
	}

	var events []Event
	err = c.db.Select(&events, "SELECT event, value FROM system.events")
	if err != nil {
		return err
	}
	for _, e := range events {
		m, err := prometheus.NewConstMetric(
			prometheus.NewDesc(
				fmt.Sprintf("%s_%s_total", metricsNamespace, fixMetricName(e.Event)),
				fmt.Sprintf("Number of %s total processed", e.Event),
				[]string{}, nil),
			prometheus.CounterValue, e.Value)
		if err != nil {
			log.Printf("failed to init const metric: %v", err)
			continue
		}
		ch <- m
	}

	return nil
}

var _ prometheus.Collector = (*ClickHouseCollector)(nil)
