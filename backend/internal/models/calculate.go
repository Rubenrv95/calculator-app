package models

// Operation identifies which arithmetic operation to perform.
type Operation string

const (
	OpAdd      Operation = "add"
	OpSubtract Operation = "subtract"
	OpMultiply Operation = "multiply"
	OpDivide   Operation = "divide"
	OpPower    Operation = "power"
	OpSqrt     Operation = "sqrt"
	OpPercent  Operation = "percent"
)

// CalculateRequest is the JSON body accepted by POST /calculate.
// B is optional for unary operations (sqrt only needs A).
type CalculateRequest struct {
	Operation Operation `json:"operation"`
	A         *float64  `json:"a"`
	B         *float64  `json:"b,omitempty"`
}

// CalculateResponse is the JSON body returned on success.
type CalculateResponse struct {
	Result float64 `json:"result"`
}

// ErrorResponse is the JSON body returned on any error.
type ErrorResponse struct {
	Error string `json:"error"`
}
