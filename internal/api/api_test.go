package api_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/lukascivil/trama-demo-api/internal/api"
)

func newServer(t *testing.T) *httptest.Server {
	t.Helper()
	srv := httptest.NewServer(api.NewServer(api.DefaultWebOrigin))
	t.Cleanup(srv.Close)
	return srv
}

func TestHelloReturnsGreeting(t *testing.T) {
	srv := newServer(t)
	res, err := http.Get(srv.URL + "/hello")
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("status = %d, want 200", res.StatusCode)
	}
	if ct := res.Header.Get("Content-Type"); ct != "application/json" {
		t.Fatalf("Content-Type = %q", ct)
	}
	// Decode into a map, not api.HelloResponse, so a renamed JSON key fails here.
	var body map[string]string
	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		t.Fatal(err)
	}
	if got := body["greeting"]; got != "Hello, World" {
		t.Fatalf("greeting = %q", got)
	}
	if got := body["message"]; got != "Hello, World" {
		t.Fatalf("message = %q, want it kept until every client reads greeting", got)
	}
	if got := res.Header.Get("Access-Control-Allow-Origin"); got != "http://localhost:5173" {
		t.Fatalf("Access-Control-Allow-Origin = %q", got)
	}
}

func TestHelloRejectsPost(t *testing.T) {
	srv := newServer(t)
	res, err := http.Post(srv.URL+"/hello", "application/json", nil)
	if err != nil {
		t.Fatal(err)
	}
	res.Body.Close()
	if res.StatusCode != http.StatusMethodNotAllowed {
		t.Fatalf("status = %d, want 405", res.StatusCode)
	}
}

func TestPreflightReturnsNoContent(t *testing.T) {
	srv := newServer(t)
	req, _ := http.NewRequest(http.MethodOptions, srv.URL+"/hello", nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	res.Body.Close()
	if res.StatusCode != http.StatusNoContent {
		t.Fatalf("status = %d, want 204", res.StatusCode)
	}
	if got := res.Header.Get("Access-Control-Allow-Origin"); got != "http://localhost:5173" {
		t.Fatalf("Access-Control-Allow-Origin = %q", got)
	}
}
