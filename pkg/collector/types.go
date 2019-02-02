package collector

// Metric from ClickHouse.
type Metric struct {
	Metric string  `db:"metric"`
	Value  float64 `db:"value"`
}

// Event from ClickHouse.
type Event struct {
	Event string  `db:"event"`
	Value float64 `db:"value"`
}
