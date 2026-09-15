# Calculator App

Full-stack calculator application: a Go backend exposing a REST API for
arithmetic operations, and a React + TypeScript frontend that consumes it.

## Repository structure

```
calculator-app/
├── backend/    # REST API in Go
├── frontend/   # UI in React + TypeScript (Vite)
├── PROMPTS.md  # Main prompts used during development
└── README.md
```

## Prerequisites

- [Go](https://go.dev/dl/) 1.22 or later (uses the `"POST /calculate"`
  pattern from `net/http`, available since Go 1.22).
- [Node.js](https://nodejs.org/) 18 or later and npm.

## Installation and setup

```bash
# Backend: no external dependencies, standard library only
cd backend
go build ./...

# Frontend
cd frontend
npm install
```

## Running in development

**Backend** (listens on `http://localhost:8080`):

```bash
cd backend
go run ./cmd/server
```

**Frontend** (listens on `http://localhost:5173`, with hot reload):

```bash
cd frontend
npm run dev
```

With both running, open `http://localhost:5173` in the browser. The
frontend client points to `http://localhost:8080` by default; this can be
overridden with the `VITE_API_BASE_URL` environment variable (for example
in a `frontend/.env.local` file).

## API

### Design choice

A **single `POST /calculate` endpoint** was chosen, with the operation
specified in the request body, instead of one endpoint per operation
(`/add`, `/subtract`, etc.). Reasons:

- The API contract is uniform: a single request/response shape for all
  seven operations, which simplifies both the frontend client (one
  `calculate()` function) and backend validation.
- Adding a new operation doesn't require a new route or router changes:
  it's just an extra case in the `internal/operations` switch.
- It's the most common pattern for calculator-style APIs (the "operator" is
  a domain parameter, not a distinct HTTP resource).

The trade-off is a less pure REST semantics (a single POST verb does
everything), but for the size and purpose of this project, simplicity on
the client and router side was prioritized.

### `POST /calculate`

**Request body:**

```json
{
  "operation": "add",
  "a": 10,
  "b": 5
}
```

| Field       | Type   | Required | Description                                                                 |
|-------------|--------|----------|-------------------------------------------------------------------------------|
| `operation` | string | yes      | `add`, `subtract`, `multiply`, `divide`, `power`, `sqrt`, `percent`          |
| `a`         | number | yes      | First operand                                                               |
| `b`         | number | depends  | Second operand. Not used (nor required) for `sqrt`                          |

For `percent`, the result is "`a`% of `b`" (e.g. `a=50, b=200` → `100`).

**Successful response (`200 OK`):**

```json
{ "result": 15 }
```

**Examples (one per operation):**

```bash
# Addition
curl -X POST http://localhost:8080/calculate \
  -H "Content-Type: application/json" \
  -d '{"operation":"add","a":10,"b":5}'
# → {"result":15}

# Subtraction
curl -X POST http://localhost:8080/calculate \
  -H "Content-Type: application/json" \
  -d '{"operation":"subtract","a":10,"b":5}'
# → {"result":5}

# Multiplication
curl -X POST http://localhost:8080/calculate \
  -H "Content-Type: application/json" \
  -d '{"operation":"multiply","a":10,"b":5}'
# → {"result":50}

# Division
curl -X POST http://localhost:8080/calculate \
  -H "Content-Type: application/json" \
  -d '{"operation":"divide","a":10,"b":5}'
# → {"result":2}

# Exponentiation
curl -X POST http://localhost:8080/calculate \
  -H "Content-Type: application/json" \
  -d '{"operation":"power","a":2,"b":10}'
# → {"result":1024}

# Square root (b is not needed)
curl -X POST http://localhost:8080/calculate \
  -H "Content-Type: application/json" \
  -d '{"operation":"sqrt","a":16}'
# → {"result":4}

# Percentage (a% of b)
curl -X POST http://localhost:8080/calculate \
  -H "Content-Type: application/json" \
  -d '{"operation":"percent","a":50,"b":200}'
# → {"result":100}

# Error example: division by zero
curl -X POST http://localhost:8080/calculate \
  -H "Content-Type: application/json" \
  -d '{"operation":"divide","a":10,"b":0}'
# → 422 Unprocessable Entity
# {"error":"division by zero"}
```

**Error codes:**

| Status | Case                                                                 |
|--------|------------------------------------------------------------------------|
| `400`  | Malformed JSON, missing or unknown `operation`, missing or non-numeric operand |
| `422`  | Division by zero, square root of a negative number                    |

### `GET /health`

Liveness endpoint, returns `{"status":"ok"}` with `200 OK`.

## Design decisions

**Backend:**

- **Router**: `net/http` with `ServeMux` (standard library), using the
  `"POST /calculate"` pattern available since Go 1.22. For an API with a
  handful of routes, pulling in a third-party router like `chi` wasn't
  justified; the standard library already supports method matching, and
  avoiding the dependency keeps the project simpler to audit and maintain.
- **Separated layers**: `internal/models` (request/response types),
  `internal/operations` (pure arithmetic logic, no HTTP knowledge, easy to
  test in isolation), and `internal/handlers` (decodes JSON, calls
  `operations`, translates errors into HTTP status codes). This separation
  allows testing business logic without spinning up an HTTP server.
- **Error handling**: `operations` returns typed errors
  (`ErrDivisionByZero`, `ErrNegativeSqrt`, etc.) that handlers map to HTTP
  status codes via `errors.Is`, instead of comparing strings or mixing
  status-code logic into the calculation itself.
- **CORS**: enabled only for the frontend's development origins
  (`localhost:5173` / `127.0.0.1:5173`), via middleware.

**Frontend:**

- **HTTP client**: native `fetch` instead of `axios`. For a single endpoint
  with no need for interceptors, advanced cancellation, or global
  request/response transforms, `fetch` avoids an extra dependency without
  losing readability.
- **Separation of concerns**: the API call lives only in
  `src/api/calculatorClient.ts`; the `Calculator` component never calls
  `fetch` directly, which allows mocking the client in UI tests without
  touching the network.
- **Client-side validation**: `src/validation.ts` validates that fields are
  non-empty and numeric before calling the backend, giving immediate
  feedback without an unnecessary round-trip.
- **Backend errors vs. network errors**: `calculatorClient` distinguishes a
  backend error response (uses the message returned by the API) from a
  network failure (server down), always showing a readable message in the
  UI instead of a raw error.
- **Dark mode**: a toggle in the top-right corner (`useTheme` +
  `ThemeToggle`) switches between a light and a dark, red-accented theme.
  The choice is persisted in `localStorage` and falls back to the OS's
  `prefers-color-scheme` on first load; the themes are implemented as CSS
  custom properties swapped via a `data-theme` attribute on `<html>`, with
  no extra styling library.

**Assumptions:**

- `percent` is interpreted as "`a`% of `b`" (requires both operands), not
  as a unary operation on a single number.
- No authentication or persistence was implemented: the requirements don't
  ask for it and the scope is a stateless calculator.

## Tests and coverage

**Backend:**

```bash
cd backend
go test ./... -cover
```

For a browsable HTML report:

```bash
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out -o=coverage.html
```

Current coverage: ~98% in `internal/handlers`, ~97% in `internal/operations`.

**Frontend:**

```bash
cd frontend
npm run test            # run tests once
npm run test:watch      # watch mode
npm run test:coverage   # tests + coverage report (frontend/coverage/index.html)
```

Current coverage: ~97% of statements.

## Docker

_Not included in this submission (optional step of the requirements)._
