package main

import (
	"log"
	"net/http"
	"os"

	"github.com/lukascivil/trama-demo-api/internal/api"
)

const addr = ":8080"

func main() {
	origin := os.Getenv("WEB_ORIGIN")
	if origin == "" {
		origin = api.DefaultWebOrigin
	}
	log.Printf("listening on http://localhost%s", addr)
	log.Fatal(http.ListenAndServe(addr, api.NewServer(origin)))
}
