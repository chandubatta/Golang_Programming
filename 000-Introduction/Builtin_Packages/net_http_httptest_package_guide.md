# `net/http/httptest` Package in Go

The `net/http/httptest` package is part of Go's standard library and provides utilities for **testing HTTP servers, handlers, and clients without needing to start a real network server**.

It is especially useful when writing unit and integration tests for code built with `net/http`.

## 1. What is `net/http/httptest`?

You typically use `net/http/httptest` to:

- Test an `http.Handler` directly.
- Create fake HTTP requests.
- Capture and inspect HTTP responses.
- Run a temporary local HTTP server.
- Test HTTP clients against a controlled server.
- Test middleware.
- Verify status codes, headers, cookies, and response bodies.
- Test HTTPS/TLS behavior.

For example, instead of actually starting an API server on a port and making a network request to it, you can test the handler entirely in memory:

```text
Test
 │
 ├── Create fake request
 │
 ├── Create fake response recorder
 │
 ├── Call HTTP handler
 │
 └── Inspect response
       ├── Status code
       ├── Headers
       └── Body
```

# 2. Important Functions and Types in `net/http/httptest`

The package is small, but several of its functions and types are extremely useful.

## `httptest.NewRequest`

```go
func NewRequest(method, target string, body io.Reader) *http.Request
```

Creates an `*http.Request` suitable for testing an HTTP handler.

### Parameters

**`method`**

HTTP method such as:

```go
"GET"
"POST"
"PUT"
"DELETE"
```

**`target`**

The URL or request target.

For example:

```go
"/users"
"/users?id=10"
"https://example.com/users"
```

**`body`**

The request body.

You can use:

```go
nil
```

when there is no body.

Or:

```go
strings.NewReader(`{"name":"Alice"}`)
```

for JSON data.

### Example

```go
req := httptest.NewRequest(
    http.MethodGet,
    "/users",
    nil,
)
```

Unlike constructing a request manually with `http.NewRequest`, `httptest.NewRequest` is specifically designed for server-handler testing.

---

# `httptest.NewRecorder`

```go
func NewRecorder() *ResponseRecorder
```

Creates a `ResponseRecorder`.

A `ResponseRecorder` records what an HTTP handler writes to its response.

For example, if your handler does:

```go
w.WriteHeader(http.StatusCreated)
w.Write([]byte("created"))
```

the recorder allows your test to inspect:

- status code
- headers
- response body

### Example

```go
recorder := httptest.NewRecorder()

handler.ServeHTTP(recorder, req)

fmt.Println(recorder.Code)
fmt.Println(recorder.Body.String())
```

---

# `httptest.NewServer`

```go
func NewServer(handler http.Handler) *Server
```

Creates and starts an HTTP server using a temporary local address.

This is particularly useful for testing HTTP clients.

Example:

```go
server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Hello"))
}))

defer server.Close()

resp, err := http.Get(server.URL)
```

The important difference is:

```text
NewRecorder
    ↓
Tests a handler directly

NewServer
    ↓
Starts a real local HTTP server
    ↓
Tests an HTTP client against it
```

---

# `httptest.NewTLSServer`

```go
func NewTLSServer(handler http.Handler) *Server
```

Creates a test HTTPS server.

It is similar to `NewServer`, except the server uses TLS/HTTPS.

Example:

```go
server := httptest.NewTLSServer(http.HandlerFunc(
    func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Secure response"))
    },
))

defer server.Close()
```

You can use:

```go
server.Client()
```

to obtain an HTTP client configured to communicate with this test TLS server.

This is useful when testing code that communicates over HTTPS.

---

# `httptest.NewUnstartedServer`

```go
func NewUnstartedServer(handler http.Handler) *Server
```

Creates a test HTTP server **without starting it**.

This is useful when you need to configure the server before starting it.

For example, you might want to configure TLS or modify the listener.

```go
server := httptest.NewUnstartedServer(handler)

// Configure server here.

server.Start()
defer server.Close()
```

Compare:

```go
httptest.NewServer(handler)
```

which starts immediately, with:

```go
httptest.NewUnstartedServer(handler)
```

which gives you an opportunity to configure the server first.

---

# `httptest.NewTLSServer` vs `NewUnstartedServer`

There is also:

```go
NewUnstartedServer
```

which can subsequently be configured for TLS:

```go
server := httptest.NewUnstartedServer(handler)

server.StartTLS()

defer server.Close()
```

This is useful when you need control over the server setup before TLS is enabled.

---

# `httptest.NewRequestWithContext`

Modern Go versions also provide:

```go
func NewRequestWithContext(
    ctx context.Context,
    method string,
    target string,
    body io.Reader,
) *http.Request
```

It creates a test HTTP request with a specific `context.Context`.

This is useful when your handler depends on context values, cancellation, or deadlines.

Example:

```go
ctx := context.WithValue(
    context.Background(),
    "userID",
    123,
)

req := httptest.NewRequestWithContext(
    ctx,
    http.MethodGet,
    "/profile",
    nil,
)
```

Your handler can then access the request context:

```go
userID := r.Context().Value("userID")
```

**Note:** Depending on the Go version you are using, this API may differ from older tutorials. Always check the documentation for your installed Go version.

---

# 3. `httptest.ResponseRecorder`

`ResponseRecorder` is one of the most important types in the package.

```go
type ResponseRecorder struct
```

It implements:

```go
http.ResponseWriter
```

so you can pass it to an HTTP handler.

It records the response produced by the handler.

Important fields include the following.

## `Code`

```go
Code int
```

Contains the HTTP status code recorded by the recorder.

Example:

```go
if recorder.Code != http.StatusOK {
    t.Errorf("expected 200, got %d", recorder.Code)
}
```

---

## `HeaderMap`

```go
HeaderMap http.Header
```

Historically used to inspect response headers.

However, when testing response headers, you should generally prefer:

```go
recorder.Result().Header
```

rather than relying directly on `HeaderMap`.

---

## `Body`

```go
Body *bytes.Buffer
```

Contains the response body.

Example:

```go
body := recorder.Body.String()

if body != "Hello" {
    t.Errorf("unexpected body: %s", body)
}
```

---

## `Flushed`

```go
Flushed bool
```

Indicates whether the handler called:

```go
Flush()
```

on the recorder.

This can be useful when testing streaming-style HTTP responses.

---

# `ResponseRecorder.Header`

`ResponseRecorder` implements:

```go
Header() http.Header
```

This allows the handler to set headers:

```go
w.Header().Set("Content-Type", "application/json")
```

and lets your test inspect them.

Example:

```go
contentType := recorder.Header().Get("Content-Type")
```

---

# `ResponseRecorder.Write`

```go
func (rw *ResponseRecorder) Write(buf []byte) (int, error)
```

Records response data written by the handler.

Normally, you don't call this yourself.

Instead, your handler calls:

```go
w.Write(...)
```

and because `ResponseRecorder` implements `http.ResponseWriter`, the data is captured.

---

# `ResponseRecorder.WriteHeader`

```go
func (rw *ResponseRecorder) WriteHeader(code int)
```

Records the HTTP status code.

For example:

```go
w.WriteHeader(http.StatusNotFound)
```

causes the recorder to capture:

```text
404
```

---

# `ResponseRecorder.Flush`

```go
func (rw *ResponseRecorder) Flush()
```

Implements the HTTP flushing behavior.

It can be used to test handlers that send data incrementally rather than waiting for the entire response.

For example, streaming or Server-Sent Events (SSE) handlers may use:

```go
w.(http.Flusher).Flush()
```

A `ResponseRecorder` can help test such behavior.

---

# `ResponseRecorder.Result`

```go
func (rw *ResponseRecorder) Result() *http.Response
```

Returns the recorded response as an `*http.Response`.

This is generally the preferred way to inspect the final response.

Example:

```go
result := recorder.Result()

fmt.Println(result.StatusCode)
fmt.Println(result.Header)
```

You can also read its body:

```go
body, err := io.ReadAll(result.Body)
```

Remember to close the body when you're finished:

```go
defer result.Body.Close()
```

---

# 4. `httptest.Server`

The `Server` type represents a test HTTP server.

```go
type Server struct
```

Important fields and methods include the following.

## `Server.URL`

```go
URL string
```

Contains the server's URL.

For example:

```text
http://127.0.0.1:54321
```

You can use it when making requests:

```go
resp, err := http.Get(server.URL)
```

---

## `Server.Client`

```go
func (s *Server) Client() *http.Client
```

Returns an HTTP client configured to communicate with the test server.

This is particularly important for TLS servers.

Example:

```go
server := httptest.NewTLSServer(handler)
defer server.Close()

client := server.Client()

resp, err := client.Get(server.URL)
```

---

## `Server.Close`

```go
func (s *Server) Close()
```

Stops the test server.

Typically you should immediately defer it:

```go
server := httptest.NewServer(handler)
defer server.Close()
```

This ensures cleanup even if the test fails.

---

## `Server.CloseClientConnections`

```go
func (s *Server) CloseClientConnections()
```

Closes connections that have been made to the test server.

This can be useful when testing connection behavior or ensuring that clients don't keep connections alive.

---

## `Server.Start`

```go
func (s *Server) Start()
```

Starts a server created using:

```go
httptest.NewUnstartedServer(...)
```

Example:

```go
server := httptest.NewUnstartedServer(handler)

server.Start()

defer server.Close()
```

---

## `Server.StartTLS`

```go
func (s *Server) StartTLS()
```

Starts an unstarted test server using TLS.

Example:

```go
server := httptest.NewUnstartedServer(handler)

server.StartTLS()

defer server.Close()
```

---

## `Server.TLS`

```go
TLS *tls.Config
```

Contains TLS configuration for the test server.

This is useful when testing HTTPS behavior.

---

## `Server.Listener`

```go
Listener net.Listener
```

Represents the network listener used by the test server.

It can be useful when you need lower-level control over the server.

---

## `Server.Config`

```go
Config *http.Server
```

Provides access to the underlying `http.Server`.

This can be useful when configuring or inspecting server behavior.

---

## `Server.EnableHTTP2`

```go
EnableHTTP2 bool
```

When enabled before starting the server, the test server can support HTTP/2.

This is useful when testing applications whose behavior depends on HTTP/2.

---

# 5. Complete Simple Example

Suppose we have this handler:

```go
package main

import (
    "fmt"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/plain")
    w.WriteHeader(http.StatusOK)

    fmt.Fprint(w, "Hello, Go!")
}
```

We can test it using `httptest`.

```go
package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/plain")
    w.WriteHeader(http.StatusOK)

    w.Write([]byte("Hello, Go!"))
}

func TestHelloHandler(t *testing.T) {
    req := httptest.NewRequest(
        http.MethodGet,
        "/hello",
        nil,
    )

    recorder := httptest.NewRecorder()

    helloHandler(recorder, req)

    if recorder.Code != http.StatusOK {
        t.Errorf(
            "expected status %d, got %d",
            http.StatusOK,
            recorder.Code,
        )
    }

    expected := "Hello, Go!"

    if recorder.Body.String() != expected {
        t.Errorf(
            "expected body %q, got %q",
            expected,
            recorder.Body.String(),
        )
    }

    if recorder.Header().Get("Content-Type") != "text/plain" {
        t.Error("expected Content-Type to be text/plain")
    }
}
```

The important sequence is:

```go
req := httptest.NewRequest(...)
```

↓

```go
recorder := httptest.NewRecorder()
```

↓

```go
helloHandler(recorder, req)
```

↓

```go
recorder.Code
recorder.Body
recorder.Header()
```

This allows you to test the handler **without starting an actual HTTP server**.

---

# 6. Testing an HTTP Client with `NewServer`

`httptest.NewServer` becomes particularly useful when the code you're testing is an HTTP **client**.

```go
func TestAPIClient(t *testing.T) {
    server := httptest.NewServer(
        http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            if r.URL.Path != "/users" {
                http.NotFound(w, r)
                return
            }

            w.Header().Set("Content-Type", "application/json")
            w.Write([]byte(`{"name":"Alice"}`))
        }),
    )

    defer server.Close()

    resp, err := http.Get(server.URL + "/users")
    if err != nil {
        t.Fatal(err)
    }

    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        t.Fatalf("expected 200, got %d", resp.StatusCode)
    }
}
```

Here the architecture is:

```text
Your HTTP client
       |
       | HTTP request
       ↓
httptest.Server
       |
       ↓
Fake/test handler
```

This lets you test your client against predictable responses.

---

# 7. Three Common Beginner Mistakes

## Mistake 1: Starting a real server unnecessarily

A beginner might write:

```go
http.ListenAndServe(":8080", handler)
```

inside a test.

This introduces unnecessary complications:

- port conflicts
- cleanup problems
- slower tests
- external networking
- possible flaky tests

### Better approach

For handler tests:

```go
httptest.NewRecorder()
```

For client tests:

```go
httptest.NewServer()
```

---

## Mistake 2: Forgetting `server.Close()`

This is a common mistake:

```go
server := httptest.NewServer(handler)
```

without cleanup.

### Better

```go
server := httptest.NewServer(handler)
defer server.Close()
```

The deferred cleanup ensures that the test server is shut down.

---

## Mistake 3: Testing only the body

A beginner may check:

```go
if recorder.Body.String() != "OK" {
    t.Error("wrong response")
}
```

but HTTP behavior includes much more than the body.

You should also consider testing:

```text
Status code
Headers
Cookies
Response body
Redirect behavior
HTTP method
Query parameters
Request body
Authentication behavior
```

For example:

```go
if recorder.Code != http.StatusUnauthorized {
    t.Errorf("expected 401")
}
```

and:

```go
if recorder.Header().Get("Content-Type") != "application/json" {
    t.Errorf("wrong content type")
}
```

---

# 8. Two Real-World Applications

## Application 1: Testing REST APIs

Suppose you build:

```text
GET /users
POST /users
GET /users/{id}
DELETE /users/{id}
```

You can use `httptest` to test every endpoint without deploying the application.

For example:

```text
Test request
    ↓
HTTP router
    ↓
Middleware
    ↓
Handler
    ↓
ResponseRecorder
    ↓
Assertions
```

You can verify:

- HTTP status codes
- JSON responses
- validation errors
- authentication
- authorization
- headers
- cookies
- middleware behavior

---

## Application 2: Testing HTTP API clients

Suppose your Go application communicates with:

```text
Payment API
Weather API
GitHub API
Internal microservice
Authentication service
```

Instead of making real API requests during tests, create:

```go
httptest.NewServer(...)
```

and simulate responses such as:

```text
200 OK
400 Bad Request
401 Unauthorized
404 Not Found
429 Too Many Requests
500 Internal Server Error
```

This makes your tests:

- deterministic
- fast
- independent of external services
- safe to run repeatedly

---

# 9. Three Progressively Challenging Exercises

## Exercise 1 — Basic Handler Testing

Create an HTTP handler:

```text
GET /hello
```

It should return:

```text
Hello, World!
```

with:

```text
HTTP 200 OK
```

Use `httptest.NewRequest` and `httptest.NewRecorder` to test:

1. The status code.
2. The response body.
3. The `Content-Type` header.

**Do not start a real HTTP server.**

---

## Exercise 2 — JSON API Testing

Create:

```text
GET /users?id=42
```

The handler should return a JSON user object.

Write tests using `httptest` that verify:

1. The request method.
2. The query parameter.
3. The status code.
4. The `Content-Type`.
5. The JSON response.
6. The behavior when `id` is missing.
7. The behavior when `id` is invalid.

Your tests should cover both successful and error responses.

---

## Exercise 3 — HTTP Client + Test Server

Create a Go HTTP client function that retrieves user information from:

```text
GET /users/{id}
```

The function should communicate with an HTTP server.

Then create an `httptest.NewServer` that simulates the remote API.

Your tests should verify how your client behaves when the server returns:

1. `200 OK` with valid JSON.
2. `400 Bad Request`.
3. `401 Unauthorized`.
4. `404 Not Found`.
5. `500 Internal Server Error`.
6. Invalid JSON.
7. A deliberately slow response.
8. A closed/unavailable server.

The goal is to test your HTTP client's error handling without communicating with a real external service.

---

# 10. Mental Model to Remember

| Tool | Main purpose |
|---|---|
| `NewRequest` | Create a test HTTP request |
| `NewRequestWithContext` | Create a test request with context |
| `NewRecorder` | Capture a handler's response |
| `ResponseRecorder` | Inspect recorded HTTP response |
| `NewServer` | Start a temporary HTTP server |
| `NewTLSServer` | Start a temporary HTTPS server |
| `NewUnstartedServer` | Create a server before starting it |
| `Server.Client()` | Get a client configured for the test server |
| `Server.Close()` | Shut down the test server |
| `Server.Start()` | Start an unstarted HTTP server |
| `Server.StartTLS()` | Start an unstarted HTTPS server |
| `Server.CloseClientConnections()` | Close active client connections |

The central distinction is:

> **Testing a handler? Use `NewRequest` + `NewRecorder`. Testing an HTTP client? Use `NewServer` or `NewTLSServer`.**

# Thought-Provoking Question

If your application has both an HTTP handler and an HTTP client that communicates with another service, **how would you design your tests so that they verify the behavior of your application without accidentally turning your unit tests into slow, network-dependent integration tests?**
