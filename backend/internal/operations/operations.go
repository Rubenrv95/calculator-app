// Package operations implements the calculator's arithmetic logic, kept
// free of any HTTP concerns so it can be tested and reused in isolation.
package operations

import (
	"errors"
	"math"

	"github.com/Rubenrv95/calculator-app/backend/internal/models"
)

// ErrDivisionByZero is returned when a division's divisor is zero.
var ErrDivisionByZero = errors.New("Division by zero")

// ErrNegativeSqrt is returned when the square root of a negative number is requested.
var ErrNegativeSqrt = errors.New("square root of a negative number is not a real number")

// ErrMissingOperand is returned when a required operand is absent.
var ErrMissingOperand = errors.New("missing required operand")

// ErrUnsupportedOperation is returned for an operation the calculator does not know.
var ErrUnsupportedOperation = errors.New("unsupported operation")

// Calculate dispatches to the requested operation and returns its result.
func Calculate(op models.Operation, a, b *float64) (float64, error) {
	if a == nil {
		return 0, ErrMissingOperand
	}

	switch op {
	case models.OpAdd:
		if b == nil {
			return 0, ErrMissingOperand
		}
		return Add(*a, *b), nil
	case models.OpSubtract:
		if b == nil {
			return 0, ErrMissingOperand
		}
		return Subtract(*a, *b), nil
	case models.OpMultiply:
		if b == nil {
			return 0, ErrMissingOperand
		}
		return Multiply(*a, *b), nil
	case models.OpDivide:
		if b == nil {
			return 0, ErrMissingOperand
		}
		return Divide(*a, *b)
	case models.OpPower:
		if b == nil {
			return 0, ErrMissingOperand
		}
		return Power(*a, *b), nil
	case models.OpSqrt:
		return Sqrt(*a)
	case models.OpPercent:
		if b == nil {
			return 0, ErrMissingOperand
		}
		return Percent(*a, *b), nil
	default:
		return 0, ErrUnsupportedOperation
	}
}

func Add(a, b float64) float64 {
	return a + b
}

func Subtract(a, b float64) float64 {
	return a - b
}

func Multiply(a, b float64) float64 {
	return a * b
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	}
	return a / b, nil
}

func Power(base, exponent float64) float64 {
	return math.Pow(base, exponent)
}

func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, ErrNegativeSqrt
	}
	return math.Sqrt(a), nil
}

// Percent returns a as a percentage of b (i.e. a% of b), e.g. Percent(50, 200) == 100.
func Percent(a, b float64) float64 {
	return (a / 100) * b
}
