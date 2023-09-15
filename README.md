# Dereferrer Service

[![Integration](https://github.com/systemli/dereferrer/actions/workflows/integration.yml/badge.svg)](https://github.com/systemli/dereferrer/actions/workflows/integration.yml) [![Quality](https://github.com/systemli/dereferrer/actions/workflows/quality.yml/badge.svg)](https://github.com/systemli/dereferrer/actions/workflows/quality.yml) [![codecov](https://codecov.io/gh/systemli/dereferrer/graph/badge.svg?token=TXTUP2J7MW)](https://codecov.io/gh/systemli/dereferrer)

This small service aims to prevent links to be tracked by the website they are linking to. It takes the URL as a parameter and returns a redirect to the URL with the referrer header removed. Especially useful for privacy-aware sites.

Imagine you serve this service on `https://dereferrer.example.com` and you want to link to `https://example.com`. Instead of linking to `https://example.com` directly, you link to `https://dereferrer.example.com/?https%3A%2F%2Fexample.com`. The service will then redirect you to `https://example.com` with the referrer header replaced with `https://dereferrer.example.com`.

Important: The query parameter should be URL encoded.

## Usage

### Environment Variables

| Variable | Description | Default |
| -------- | ----------- | ------- |
| `LISTEN_ADDR` | Address to listen on | `:8080` |
| `METRICS_ADDR` | Address to listen on for metrics | `:8081` |
| `LOG_LEVEL` | Log level | `info` |

### Docker

```bash
docker run -p 8080:8080 -p 8081:8081 -d --name dereferrer \
docker.io/systemli/dereferrer:latest
```

The service needs no special capabilities and can start as read-only.

```bash
docker run -p 8080:8080 -p 8081:8081 -d --name dereferrer \
--read-only --cap-drop all \
docker.io/systemli/dereferrer:latest
```

### Go

```bash
go install github.com/systemli/dereferrer@latest
dereferrer
```

## Metrics

The service exposes metrics on the `/metrics` endpoint on port `8081`. The metrics are compatible with Prometheus.

Exported metrics:

```text
# HELP requests_total Number of requests
# TYPE requests_total counter
requests_total{status="200"} 6145
requests_total{status="400"} 619
requests_total{status="404"} 651
requests_total{status="405"} 35
```

## License

This project is licensed under the GPL-3.0 License - see the [LICENSE](LICENSE) file for details.
