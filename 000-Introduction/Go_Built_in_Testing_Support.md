# Built-in Testing Support in Go

## 1. What is Built-in Testing Support?

In Go, **Built-in Testing Support** refers mainly to Go's standard `testing` package and the `go test` command. Go provides testing directly as part of its standard toolchain, so you don't need a third-party testing framework for basic unit and integration testing.

Go provides built-in support for:

- `testing` package — provides testing functionality.
- `go test` — discovers and executes tests.
- `*_test.go` — files containing test code.
- `TestXxx` functions — ordinary unit tests.
- `BenchmarkXxx` — performance benchmarks.
- `ExampleXxx` — executable examples.
- `FuzzXxx` — fuzz testing.

### Purpose

Its main purpose is to help developers:

- Verify that functions work correctly.
- Catch bugs early.
- Prevent existing functionality from breaking.
- Measure performance.
- Test code with different inputs.

### When is it commonly used?

It is commonly used during:

1. **Development** — test individual functions while writing code.
2. **Refactoring** — ensure existing behavior still works.
3. **CI/CD** — automatically run tests before deploying applications.
4. **API development** — verify HTTP handlers and responses.
5. **Production maintenance** — prevent previously fixed bugs from returning.

---

## 2. Simple Code Example

Suppose we have a function that adds two numbers.

### `calculator.go`

```go
package calculator

func Add(a int, b int) int {
    return a + b
}
```

Now create a test file.

### `calculator_test.go`

```go
package calculator

import "testing"

func TestAdd(t *testing.T) {
    result := Add(10, 20)

    expected := 30

    if result != expected {
        t.Errorf("expected %d, got %d", expected, result)
    }
}
```

The important flow is:

```text
calculator.go
calculator_test.go
       ↓
go test
       ↓
TestAdd()
       ↓
PASS / FAIL
```

Run the test:

```bash
go test
```

For more detailed output:

```bash
go test -v
```

Example output:

```text
=== RUN   TestAdd
--- PASS: TestAdd (0.00s)
PASS
ok      calculator
```

### Why is this "built-in"?

You don't need to install a separate testing framework for basic testing.

The testing functionality comes with Go itself:

```go
import "testing"
```

and the Go toolchain runs tests with:

```bash
go test
```

---

## 3. Three Common Mistakes and Misconceptions

### Mistake 1: Not following Go's test naming convention

Beginners sometimes write:

```go
func checkAdd(t *testing.T) {
    // ...
}
```

and expect Go to automatically execute it.

Go expects test functions to follow the convention:

```go
func TestAdd(t *testing.T) {
    // ...
}
```

The file should normally end with:

```text
_test.go
```

**Avoid it:** Use the standard naming convention:

```text
*_test.go
TestXxx
```

---

### Mistake 2: Thinking `go test` tests the entire application automatically

`go test` doesn't magically understand what behavior your application should have.

You need to write tests describing the expected behavior.

For example:

```go
expected := 30

if result != expected {
    t.Errorf("wrong result")
}
```

**Avoid it:** Think of tests as executable specifications of expected behavior.

---

### Mistake 3: Only testing the "happy path"

A beginner might test:

```text
10 + 20 → 30
```

but never consider:

```text
0 + 0
-10 + 20
large numbers
```

For real applications, invalid input and edge cases are often where bugs occur.

**Avoid it:** Test normal cases, boundary cases, and invalid cases where applicable.

---

## 4. Real-World Applications

### Scenario 1: REST API Testing

Suppose you're building a Go backend:

```text
React Frontend
      ↓
REST API
      ↓
Go Handler
      ↓
Service
      ↓
Database
```

Built-in testing can verify:

- HTTP status codes.
- Response JSON.
- Request validation.
- Error responses.
- Authentication behavior.

For example:

```text
POST /users
       ↓
Valid request
       ↓
Expected: HTTP 201
```

A test can automatically verify that behavior.

This is particularly useful when changing backend code and making sure existing APIs still work.

---

### Scenario 2: CI/CD Pipeline

Testing is extremely useful before deploying a Go application.

A typical pipeline might look like:

```text
Developer
   ↓
git push
   ↓
CI/CD
   ↓
go test ./...
   ↓
Build
   ↓
Docker Image
   ↓
Deployment
```

If a test fails:

```text
go test ./...
      ↓
    FAIL
      ↓
Deployment stopped
```

This prevents known bugs from being deployed automatically.

---

## 5. Progressive Exercises

### Exercise 1 — Beginner: Temperature Converter

Create a Go function that converts Celsius to Fahrenheit.

Requirements:

- Create a function for the conversion.
- Create a `_test.go` file.
- Write tests for:
  - `0°C`
  - `100°C`
  - `25°C`
- Verify the expected results using the `testing` package.

**Goal:** Practice creating and running basic unit tests.

---

### Exercise 2 — Intermediate: User Validation

Create a function that validates a user:

```text
Name
Age
Email
```

Define reasonable validation rules.

For example, you might decide that:

- Name cannot be empty.
- Age must be within a reasonable range.
- Email must follow a basic valid format.

Write tests for:

- Valid user.
- Empty name.
- Invalid age.
- Invalid email.
- Multiple invalid fields.

**Goal:** Practice testing multiple input conditions and edge cases.

---

### Exercise 3 — Advanced: REST API Handler Testing

Create a small Go HTTP API with:

```text
GET /users/{id}
```

The endpoint should return user information as JSON.

Write tests using Go's built-in testing support to verify:

1. A valid user ID returns the expected user.
2. An unknown user returns the appropriate HTTP status.
3. An invalid ID is handled correctly.
4. The response contains the expected JSON.
5. The appropriate HTTP method is required.

Use Go's HTTP testing capabilities rather than manually starting a real server.

**Goal:** Move from testing simple functions to testing real backend/API behavior.

---

## Thought-Provoking Question

If Go already provides a built-in testing framework, **why do you think large Go projects might still use additional testing libraries or tools?**

Consider whether the built-in `testing` package is enough for:

- Unit tests
- API tests
- Integration tests
- Mocking
- Test data generation
- Large-scale CI/CD testing
