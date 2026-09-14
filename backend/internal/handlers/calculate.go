// Package handlers wires HTTP requests to the operations package and
// translates results and errors into JSON responses.
package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Rubenrv95/calculator-app/backend/internal/models"
	"github.com/Rubenrv95/calculator-app/backend/internal/operations"
)

// Calculate handles POST /calculate: it decodes a CalculateRequest, runs the
// requested operation, and writes a CalculateResponse or ErrorResponse.
func Calculate(w http.ResponseWriter, r *http.Request) {
	var req models.CalculateRequest

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}

	if req.Operation == "" {
		writeError(w, http.StatusBadRequest, "operation is required")
		return
	}

	result, err := operations.Calculate(req.Operation, req.A, req.B)
	if err != nil {
		status, message := statusFor(err)
		writeError(w, status, message)
		return
	}

	writeJSON(w, http.StatusOK, models.CalculateResponse{Result: result})
}

// statusFor maps a domain error to the HTTP status and message to report.
func statusFor(err error) (int, string) {
	switch {
	case errors.Is(err, operations.ErrMissingOperand):
		return http.StatusBadRequest, err.Error()
	case errors.Is(err, operations.ErrUnsupportedOperation):
		return http.StatusBadRequest, err.Error()
	case errors.Is(err, operations.ErrDivisionByZero):
		return http.StatusUnprocessableEntity, err.Error()
	case errors.Is(err, operations.ErrNegativeSqrt):
		return http.StatusUnprocessableEntity, err.Error()
	default:
		return http.StatusInternalServerError, "internal server error"
	}
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, models.ErrorResponse{Error: message})
}
