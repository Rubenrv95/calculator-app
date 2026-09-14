package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Rubenrv95/calculator-app/backend/internal/models"
)

func doCalculate(t *testing.T, body string) *httptest.ResponseRecorder {
	t.Helper()
	req := httptest.NewRequest(http.MethodPost, "/calculate", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	Calculate(rec, req)
	return rec
}

func TestCalculateHandler_Success(t *testing.T) {
	rec := doCalculate(t, `{"operation":"add","a":2,"b":3}`)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d, body = %s", rec.Code, http.StatusOK, rec.Body.String())
	}

	var resp models.CalculateResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if resp.Result != 5 {
		t.Errorf("result = %v, want 5", resp.Result)
	}
}

func TestCalculateHandler_DivisionByZero(t *testing.T) {
	rec := doCalculate(t, `{"operation":"divide","a":10,"b":0}`)

	if rec.Code != http.StatusUnprocessableEntity {
		t.Fatalf("status = %d, want %d, body = %s", rec.Code, http.StatusUnprocessableEntity, rec.Body.String())
	}

	var resp models.ErrorResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if resp.Error == "" {
		t.Error("expected non-empty error message")
	}
}

func TestCalculateHandler_NegativeSqrt(t *testing.T) {
	rec := doCalculate(t, `{"operation":"sqrt","a":-9}`)

	if rec.Code != http.StatusUnprocessableEntity {
		t.Fatalf("status = %d, want %d, body = %s", rec.Code, http.StatusUnprocessableEntity, rec.Body.String())
	}
}

func TestCalculateHandler_MalformedJSON(t *testing.T) {
	rec := doCalculate(t, `{"operation":"add",`)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d, body = %s", rec.Code, http.StatusBadRequest, rec.Body.String())
	}
}

func TestCalculateHandler_MissingOperation(t *testing.T) {
	rec := doCalculate(t, `{"a":1,"b":2}`)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d, body = %s", rec.Code, http.StatusBadRequest, rec.Body.String())
	}
}

func TestCalculateHandler_MissingOperand(t *testing.T) {
	rec := doCalculate(t, `{"operation":"add","a":1}`)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d, body = %s", rec.Code, http.StatusBadRequest, rec.Body.String())
	}
}

func TestCalculateHandler_NonNumericInput(t *testing.T) {
	rec := doCalculate(t, `{"operation":"add","a":"two","b":3}`)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d, body = %s", rec.Code, http.StatusBadRequest, rec.Body.String())
	}
}

func TestCalculateHandler_UnsupportedOperation(t *testing.T) {
	rec := doCalculate(t, `{"operation":"modulo","a":5,"b":2}`)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d, body = %s", rec.Code, http.StatusBadRequest, rec.Body.String())
	}
}

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()
	Health(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}
}
