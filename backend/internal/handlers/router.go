package handlers

import "net/http"

// NewRouter builds the application's HTTP handler, including middleware.
func NewRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /calculate", Calculate)
	mux.HandleFunc("GET /health", Health)

	return WithCORS(mux)
}

// Health is a simple liveness endpoint.
func Health(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}
