# ClickHouse Exporter for Prometheus

[![Go Report Card](https://goreportcard.com/badge/github.com/hot-wifi/clickhouse_exporter)](https://goreportcard.com/report/github.com/hot-wifi/clickhouse_exporter)

## Building and Running

### Configuration

Configuration using environment variables:

| Variable             | Description                                                                        | Default    |
|----------------------|------------------------------------------------------------------------------------|------------|
| `DEBUG`              | Debug flag                                                                         | `false`    |
| `TELEMETRY_PORT`     | Telemetry port                                                                     | `9116`     |
| `TELEMETRY_ENDPOINT` | Telemetry endpoint                                                                 | `/metrics` |
| `CLICKHOUSE_DSN`     | ClickHouse data source name ([Format](https://github.com/kshvakov/clickhouse#dsn)) |            |

### Using Docker

```bash
$ cat env-example
CLICKHOUSE_DSN="tcp://host1:9000?username=user&password=qwerty&database=clicks&read_timeout=10&write_timeout=20&alt_hosts=host2:9000,host3:9000"
$ docker run -d --env-file ./env-example hotwifi/clickhouse-exporter
```