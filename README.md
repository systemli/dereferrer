# Dereferrer Service

This small service aims to prevent links to be tracked by the website they are linking to. It takes the URL as a parameter and returns a redirect to the URL with the referrer header removed. Especially useful for privacy-aware sites.

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
