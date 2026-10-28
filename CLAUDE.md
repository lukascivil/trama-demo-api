# trama-demo-api

Demo backend for trama. Go, standard library only.

- Run: `go run ./cmd/api` (listens on http://localhost:8080)
- Test: `go test ./...`
- `GET /hello` returns `{"message": "Hello, World"}`.
- The allowed browser origin comes from `WEB_ORIGIN` (default `http://localhost:5173`).
