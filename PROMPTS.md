# Prompts used

This project was built with the assistance of Claude Code. Below is a
summary of the main prompts used during development, in the order they were
given.

## 1. Initial prompt (context and full requirements)

Claude was given the project context (existing repo with an empty scaffold
for a React + TypeScript frontend and a Go backend), and the full
functional and non-functional requirements of the technical test were
pasted verbatim (operations to support, API design, error handling, tests,
documentation, optional Docker). Claude was asked to work in stages — first
a complete backend with tests, then a complete frontend with tests, then
documentation — reporting back after each stage before moving to the next,
prioritizing clean, idiomatic code over extra features, and making
separate, descriptive commits per logical stage.

> "I'm building a technical test for a hiring process [...] Go step by
> step: first the complete backend with tests, then the complete frontend
> with tests, then documentation. Let me know when each stage is done [...]
> Let's start with the backend."

## 2. Backend

Claude was asked to build the Go server following the initial prompt:
separated layers (handlers, business logic, models), deciding and
justifying the API design (single endpoint vs. one per operation), edge
case handling (division by zero, square root of negatives, malformed JSON,
missing fields), CORS for the local frontend, and unit tests with coverage
using the standard `testing` package and `httptest`.

Claude implemented the server, ran `go build`, `go vet`, and
`go test ./... -cover`, and reported coverage metrics before continuing.

## 3. Go module adjustment

The user was asked for their GitHub username to replace the `TU_USUARIO`
placeholder in the Go module path (`go.mod` and all internal imports),
since a real module path was needed for the project to build.

## 4. Change in commit workflow

> "Can you undo the commit? I want to see the changes made. From now on
> leave the commits and push to me unless I tell you otherwise"

From this point on, Claude stopped running `git commit`/`git push` on its
own: it only leaves changes in the working tree for the user to review and
commit manually.

## 5. Manual backend verification

The user asked for the commands to start the backend and run the tests
manually. Along the way, a couple of Windows/PowerShell-specific issues
were resolved: PowerShell's `curl` alias (`Invoke-WebRequest`) doesn't
accept Unix `curl` syntax, and an issue generating the HTML coverage report
(the file was created without the `.out` extension due to a command-copying
issue), plus PowerShell's script execution policy blocking `npm.ps1`.

## 6. Frontend

Claude was asked to continue with the frontend following the same initial
prompt: a simple UI with an operation selector, client-side input
validation, backend error handling without exposing raw errors, responsive
design, an API client isolated from the UI component, and tests with
Vitest + React Testing Library mocking `fetch`.

Claude installed the testing dependencies (`vitest`,
`@testing-library/react`, `@testing-library/jest-dom`,
`@testing-library/user-event`, `jsdom`, `@vitest/coverage-v8`),
implemented the API client, validation, the `Calculator` component and
responsive styles, and ran `tsc -b`, `eslint`, and `vitest run --coverage`
to verify everything compiled, had no lint errors, and had good coverage
before continuing.

## 7. Documentation

> "Ok, let's do the README and PROMPTS"

Claude was asked to complete the root `README.md` (setup, how to run the
backend and frontend, JSON API examples, design decisions and assumptions,
how to run tests and generate coverage on both sides) and this
`PROMPTS.md` file, as required by the test's deliverables.

## 8. English documentation

> "I need you to add documentation in English to the code. Also leave the
> README and PROMPTS in English"

Code comments were already written in English throughout the backend and
frontend (user-facing strings in the UI remained in Spanish, since that's
application content rather than documentation). `README.md` and this file
were translated to English.
