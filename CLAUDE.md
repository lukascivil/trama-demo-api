# trama-demo-api

Demo backend for trama. Go, standard library only.

- Run: `go run ./cmd/api` (listens on http://localhost:8080)
- Test: `go test ./...`
- `GET /hello` returns `{"greeting": "Hello, World", "message": "Hello, World"}`. `message` is deprecated: it repeats `greeting` until every client reads `greeting`.
- The allowed browser origin comes from `WEB_ORIGIN` (default `http://localhost:5173`).
