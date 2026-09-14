package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouter_CalculateEndToEnd(t *testing.T) {
	router := NewRouter()

	req := httptest.NewRequest(http.MethodPost, "/calculate", bytes.NewBufferString(`{"operation":"multiply","a":6,"b":7}`))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Origin", "http://localhost:5173")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d, body = %s", rec.Code, http.StatusOK, rec.Body.String())
	}
	if got := rec.Header().Get("Access-Control-Allow-Origin"); got != "http://localhost:5173" {
		t.Errorf("Access-Control-Allow-Origin = %q, want the request origin", got)
	}
}

func TestRouter_HealthEndpoint(t *testing.T) {
	router := NewRouter()

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}
}

func TestRouter_CORSPreflight(t *testing.T) {
	router := NewRouter()

	req := httptest.NewRequest(http.MethodOptions, "/calculate", nil)
	req.Header.Set("Origin", "http://localhost:5173")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusNoContent)
	}
}

func TestRouter_CORSUntrustedOrigin(t *testing.T) {
	router := NewRouter()

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	req.Header.Set("Origin", "http://evil.example.com")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if got := rec.Header().Get("Access-Control-Allow-Origin"); got != "" {
		t.Errorf("Access-Control-Allow-Origin = %q, want empty for untrusted origin", got)
	}
}
