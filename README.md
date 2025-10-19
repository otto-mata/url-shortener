# URL Shortener (Go + Chi)

A minimal scaffolding for a URL shortener web app using Go and the Chi router.

## Features
- Chi-based HTTP server with middleware
- Endpoints:
  - GET /healthz
  - POST /api/shorten {"url": "https://example.com"}
  - GET /{code} -> 302 redirect
- In-memory storage implementation (swap later for Redis/DB)
- Graceful shutdown
- Simple JSON response helpers

## Quick start

### Prerequisites
- Go 1.20+

### Run
```sh
# from repo root
go run ./cmd/server
```

### Example
```sh
# shorten a URL
curl -sX POST localhost:8080/api/shorten \
  -H 'content-type: application/json' \
  -d '{"url":"https://golang.org"}'
# => {"code":"abc123","short_url":"localhost:8080/abc123"}

# follow the redirect
curl -i localhost:8080/abc123
```

## Configuration
Environment variables:
- ADDR (default :8080)
- BASE_URL (default http://localhost:8080)
- SECRET (default dev-secret)

## Project layout
- cmd/server/main.go – app entrypoint
- internal/config – env config
- internal/http – router and handlers
- internal/shortener – business logic
- internal/storage – storage interface and in-memory impl
- pkg/respond – response helpers

## Notes and next steps
- No persistence yet. Consider Redis or Postgres storage.
- No authentication or rate limiting yet.
- No custom domains or analytics yet.
- Collision handling is naive; improve by checking for existing code.

PRs welcome.
