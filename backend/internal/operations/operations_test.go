package operations

import (
	"errors"
	"math"
	"testing"

	"github.com/Rubenrv95/calculator-app/backend/internal/models"
)

func f(v float64) *float64 { return &v }

func TestAdd(t *testing.T) {
	if got := Add(2, 3); got != 5 {
		t.Errorf("Add(2, 3) = %v, want 5", got)
	}
	if got := Add(-2, -3); got != -5 {
		t.Errorf("Add(-2, -3) = %v, want -5", got)
	}
}

func TestSubtract(t *testing.T) {
	if got := Subtract(5, 3); got != 2 {
		t.Errorf("Subtract(5, 3) = %v, want 2", got)
	}
}

func TestMultiply(t *testing.T) {
	if got := Multiply(4, 3); got != 12 {
		t.Errorf("Multiply(4, 3) = %v, want 12", got)
	}
	if got := Multiply(4, 0); got != 0 {
		t.Errorf("Multiply(4, 0) = %v, want 0", got)
	}
}

func TestDivide(t *testing.T) {
	got, err := Divide(10, 2)
	if err != nil {
		t.Fatalf("Divide(10, 2) unexpected error: %v", err)
	}
	if got != 5 {
		t.Errorf("Divide(10, 2) = %v, want 5", got)
	}

	_, err = Divide(10, 0)
	if !errors.Is(err, ErrDivisionByZero) {
		t.Errorf("Divide(10, 0) error = %v, want ErrDivisionByZero", err)
	}
}

func TestPower(t *testing.T) {
	if got := Power(2, 10); got != 1024 {
		t.Errorf("Power(2, 10) = %v, want 1024", got)
	}
	if got := Power(2, 0); got != 1 {
		t.Errorf("Power(2, 0) = %v, want 1", got)
	}
}

func TestSqrt(t *testing.T) {
	got, err := Sqrt(16)
	if err != nil {
		t.Fatalf("Sqrt(16) unexpected error: %v", err)
	}
	if got != 4 {
		t.Errorf("Sqrt(16) = %v, want 4", got)
	}

	_, err = Sqrt(-4)
	if !errors.Is(err, ErrNegativeSqrt) {
		t.Errorf("Sqrt(-4) error = %v, want ErrNegativeSqrt", err)
	}
}

func TestPercent(t *testing.T) {
	if got := Percent(50, 200); got != 100 {
		t.Errorf("Percent(50, 200) = %v, want 100", got)
	}
}

func TestCalculate(t *testing.T) {
	cases := []struct {
		name    string
		op      models.Operation
		a, b    *float64
		want    float64
		wantErr error
	}{
		{"add", models.OpAdd, f(2), f(3), 5, nil},
		{"subtract", models.OpSubtract, f(5), f(3), 2, nil},
		{"multiply", models.OpMultiply, f(4), f(3), 12, nil},
		{"divide", models.OpDivide, f(10), f(2), 5, nil},
		{"divide by zero", models.OpDivide, f(10), f(0), 0, ErrDivisionByZero},
		{"power", models.OpPower, f(2), f(3), 8, nil},
		{"sqrt", models.OpSqrt, f(9), nil, 3, nil},
		{"sqrt negative", models.OpSqrt, f(-9), nil, 0, ErrNegativeSqrt},
		{"percent", models.OpPercent, f(10), f(50), 5, nil},
		{"missing a", models.OpAdd, nil, f(3), 0, ErrMissingOperand},
		{"missing b for add", models.OpAdd, f(2), nil, 0, ErrMissingOperand},
		{"missing b for subtract", models.OpSubtract, f(2), nil, 0, ErrMissingOperand},
		{"missing b for multiply", models.OpMultiply, f(2), nil, 0, ErrMissingOperand},
		{"missing b for power", models.OpPower, f(2), nil, 0, ErrMissingOperand},
		{"missing b for percent", models.OpPercent, f(2), nil, 0, ErrMissingOperand},
		{"unsupported operation", models.Operation("modulo"), f(2), f(3), 0, ErrUnsupportedOperation},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Calculate(tc.op, tc.a, tc.b)
			if tc.wantErr != nil {
				if !errors.Is(err, tc.wantErr) {
					t.Fatalf("Calculate(%s) error = %v, want %v", tc.name, err, tc.wantErr)
				}
				return
			}
			if err != nil {
				t.Fatalf("Calculate(%s) unexpected error: %v", tc.name, err)
			}
			if math.Abs(got-tc.want) > 1e-9 {
				t.Errorf("Calculate(%s) = %v, want %v", tc.name, got, tc.want)
			}
		})
	}
}
