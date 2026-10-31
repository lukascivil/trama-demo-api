# trama-demo-api

A hello world backend used to demonstrate [trama](https://github.com/lukascivil/trama). Its frontend is [trama-demo-web](https://github.com/lukascivil/trama-demo-web).

```bash
go run ./cmd/api      # http://localhost:8080
go test ./...
curl localhost:8080/hello
# {"greeting":"Hello, World","message":"Hello, World"}
```

`message` is deprecated and repeats `greeting`; it will be removed once every client reads `greeting`.

`WEB_ORIGIN` sets the browser origin allowed by CORS (default `http://localhost:5173`, where the frontend runs).

License: [MIT](LICENSE)
