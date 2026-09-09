# Go `net/http` Package

The Go `net/http` package is one of the most important standard-library packages for building **web servers, REST APIs, HTTP clients, proxies, web applications, and HTTP-based services**.

The package provides both **HTTP client and HTTP server implementations**. It also includes support for HTTP/1.x and HTTP/2, with higher-level APIs that hide most protocol details.

## 1. What is the `net/http` package?

Import it with:

```go
import "net/http"
```

The package provides APIs for two major jobs:

### HTTP Server

You can create a server that receives requests:

```text
Browser / Client
      |
      | HTTP Request
      v
Go HTTP Server
      |
      | Handler
      v
Response
```

For example:

```go
http.HandleFunc("/hello", helloHandler)
http.ListenAndServe(":8080", nil)
```

### HTTP Client

You can also use Go to communicate with other web servers:

```go
resp, err := http.Get("https://example.com")
```

Conceptually:

```text
             net/http
                |
       +--------+--------+
       |                 |
    Server              Client
       |                 |
 Receives HTTP       Sends HTTP
 requests            requests
```

Common uses include:

- REST APIs
- Web applications
- Microservices
- HTTP clients
- Calling third-party APIs
- File servers
- Authentication
- Cookies
- Redirects
- HTTPS/TLS
- HTTP middleware
- Proxies
- Health-check endpoints

---

# 2. Simple Example

Let's create a tiny web server.

```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from Go!")
}

func main() {
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Server running at http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

Run:

```bash
go run main.go
```

Then open:

```text
http://localhost:8080/hello
```

You should see:

```text
Hello from Go!
```

### What happens?

```go
http.HandleFunc("/hello", helloHandler)
```

Registers a function to handle requests to `/hello`.

```go
func helloHandler(w http.ResponseWriter, r *http.Request)
```

The two parameters are extremely important:

- `w` — used to send the HTTP response.
- `r` — contains information about the incoming request.

For example:

```go
fmt.Fprintln(w, "Hello from Go!")
```

writes the response body.

Finally:

```go
http.ListenAndServe(":8080", nil)
```

starts an HTTP server on port `8080`.

---

# 3. Important Concepts Before Learning the Functions

You should understand these four types/interfaces first.

## `http.Request`

Represents an HTTP request.

For example:

```text
GET /users/10 HTTP/1.1
Host: example.com
Authorization: ...
```

Go represents this information using:

```go
*http.Request
```

You can access:

```go
r.Method
r.URL
r.Header
r.Body
r.Cookies()
```

---

## `http.Response`

Represents an HTTP response received by an HTTP client.

For example:

```go
resp, err := http.Get("https://example.com")
```

You can inspect:

```go
resp.StatusCode
resp.Status
resp.Header
resp.Body
```

---

## `http.ResponseWriter`

Used by a server handler to construct the response.

```go
func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello"))
}
```

Think:

```text
Request ---> Handler ---> ResponseWriter
                              |
                              v
                          HTTP Response
```

---

## `http.Handler`

The fundamental server-side interface:

```go
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
```

Any type implementing `ServeHTTP` can act as an HTTP handler.

---

# 4. Important `net/http` Functions — Detailed Reference

The package contains many APIs. The following are the major exported package-level functions, followed by important methods on its types.

---

## A. `http.Get`

```go
func Get(url string) (*Response, error)
```

Performs an HTTP `GET` request.

Example:

```go
resp, err := http.Get("https://example.com")
if err != nil {
	log.Fatal(err)
}
defer resp.Body.Close()

body, err := io.ReadAll(resp.Body)
if err != nil {
	log.Fatal(err)
}

fmt.Println(string(body))
```

### Important

Always close the response body:

```go
defer resp.Body.Close()
```

Also remember:

**A HTTP 404 or 500 does not automatically produce an error from `Get`.**

You must inspect:

```go
resp.StatusCode
```

`Get` follows redirects according to the default client's redirect behavior.

---

# B. `http.Head`

```go
func Head(url string) (*Response, error)
```

Sends an HTTP `HEAD` request.

`HEAD` is useful when you want response metadata without downloading the resource body.

Example:

```go
resp, err := http.Head("https://example.com/file.zip")
if err != nil {
	log.Fatal(err)
}
defer resp.Body.Close()

fmt.Println(resp.StatusCode)
fmt.Println(resp.Header.Get("Content-Length"))
```

Useful for:

- Checking whether a resource exists
- Checking file size
- Checking modification information
- Inspecting headers

---

# C. `http.Post`

```go
func Post(
	url string,
	contentType string,
	body io.Reader,
) (*Response, error)
```

Sends an HTTP `POST` request.

Example:

```go
body := strings.NewReader(`{"name":"Chandu"}`)

resp, err := http.Post(
	"https://example.com/users",
	"application/json",
	body,
)

if err != nil {
	log.Fatal(err)
}

defer resp.Body.Close()
```

Commonly used for sending:

- JSON
- XML
- text
- form data
- files

---

# D. `http.PostForm`

```go
func PostForm(
	url string,
	data url.Values,
) (*Response, error)
```

Sends form-encoded data.

Example:

```go
data := url.Values{}

data.Set("username", "chandu")
data.Set("password", "secret")

resp, err := http.PostForm(
	"https://example.com/login",
	data,
)

if err != nil {
	log.Fatal(err)
}

defer resp.Body.Close()
```

The content type is:

```text
application/x-www-form-urlencoded
```

---

# E. `http.HandleFunc`

```go
func HandleFunc(
	pattern string,
	handler func(ResponseWriter, *Request),
)
```

Registers a function as an HTTP handler.

Example:

```go
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Home page")
}

func main() {
	http.HandleFunc("/", home)

	http.ListenAndServe(":8080", nil)
}
```

This is one of the easiest ways to create a server.

---

# F. `http.Handle`

```go
func Handle(pattern string, handler Handler)
```

Registers a `Handler`.

Unlike `HandleFunc`, it expects an object implementing:

```go
ServeHTTP(...)
```

Example:

```go
type MyHandler struct{}

func (h MyHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	fmt.Fprintln(w, "Hello!")
}

func main() {
	http.Handle("/hello", MyHandler{})

	http.ListenAndServe(":8080", nil)
}
```

---

# G. `http.ListenAndServe`

```go
func ListenAndServe(
	addr string,
	handler Handler,
) error
```

Starts an HTTP server.

Example:

```go
http.ListenAndServe(":8080", nil)
```

The first argument is the address:

```text
:8080
```

The second argument is the handler.

If the handler is `nil`, Go uses the default `ServeMux`.

The function blocks while the server runs.

---

# H. `http.ListenAndServeTLS`

```go
func ListenAndServeTLS(
	addr string,
	certFile string,
	keyFile string,
	handler Handler,
) error
```

Starts an HTTPS server.

Example:

```go
http.ListenAndServeTLS(
	":8443",
	"server.crt",
	"server.key",
	nil,
)
```

HTTPS provides encryption through TLS.

---

# I. `http.Serve`

```go
func Serve(
	l net.Listener,
	handler Handler,
) error
```

Runs an HTTP server using an existing network listener.

This gives you more control than:

```go
http.ListenAndServe()
```

Conceptually:

```text
net.Listener
     |
     v
http.Serve()
     |
     v
HTTP Server
```

---

# J. `http.ServeTLS`

```go
func ServeTLS(
	l net.Listener,
	handler Handler,
	certFile string,
	keyFile string,
) error
```

Similar to `Serve`, but provides TLS/HTTPS.

Useful when you already have a customized `net.Listener`.

---

# K. `http.Redirect`

```go
func Redirect(
	w ResponseWriter,
	r *Request,
	url string,
	code int,
)
```

Sends a redirect response.

Example:

```go
func oldPage(w http.ResponseWriter, r *http.Request) {
	http.Redirect(
		w,
		r,
		"/new-page",
		http.StatusMovedPermanently,
	)
}
```

Common status codes:

```go
http.StatusMovedPermanently // 301
http.StatusFound            // 302
http.StatusSeeOther         // 303
```

---

# L. `http.Error`

```go
func Error(
	w ResponseWriter,
	error string,
	code int,
)
```

Sends an HTTP error response.

Example:

```go
http.Error(
	w,
	"User not found",
	http.StatusNotFound,
)
```

The client receives:

```text
404 Not Found
```

---

# M. `http.NotFound`

```go
func NotFound(
	w ResponseWriter,
	r *Request,
)
```

Responds with:

```text
404 Not Found
```

Example:

```go
func handler(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}
```

---

# N. `http.ServeFile`

```go
func ServeFile(
	w ResponseWriter,
	r *Request,
	name string,
)
```

Serves a file to the client.

Example:

```go
func download(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "files/report.pdf")
}
```

Useful for:

- PDFs
- Images
- HTML
- Downloads
- Static resources

Be careful when constructing `name` from user input. Improper file-path handling can create security vulnerabilities.

---

# O. `http.ServeFileFS`

```go
func ServeFileFS(
	w ResponseWriter,
	r *Request,
	fsys fs.FS,
	name string,
)
```

Serves a file from an `fs.FS`.

This works particularly well with Go's filesystem abstractions such as `embed.FS`.

---

# P. `http.FileServer`

```go
func FileServer(root FileSystem) Handler
```

Creates a handler that serves files from a filesystem.

Example:

```go
files := http.FileServer(http.Dir("./public"))

http.Handle("/", files)

http.ListenAndServe(":8080", nil)
```

Now files inside `./public` can be served through HTTP.

---

# Q. `http.StripPrefix`

```go
func StripPrefix(
	prefix string,
	h Handler,
) Handler
```

Removes a prefix from the request URL before passing the request to another handler.

Example:

```go
files := http.FileServer(http.Dir("./public"))

http.Handle(
	"/static/",
	http.StripPrefix("/static/", files),
)
```

A request:

```text
/static/index.html
```

can be passed to the file server as:

```text
/index.html
```

---

# R. `http.RedirectHandler`

```go
func RedirectHandler(
	url string,
	code int,
) Handler
```

Returns a handler that redirects requests.

Example:

```go
handler := http.RedirectHandler(
	"https://example.com",
	http.StatusMovedPermanently,
)

http.Handle("/old", handler)
```

---

# S. `http.TimeoutHandler`

```go
func TimeoutHandler(
	h Handler,
	dt time.Duration,
	msg string,
) Handler
```

Wraps a handler with a timeout.

Example:

```go
handler := http.TimeoutHandler(
	myHandler,
	2*time.Second,
	"Request timed out",
)
```

If the handler exceeds the specified duration, the timeout handler returns a `503 Service Unavailable` response.

---

# T. `http.SetCookie`

```go
func SetCookie(
	w ResponseWriter,
	cookie *Cookie,
)
```

Adds a cookie to the HTTP response.

Example:

```go
cookie := &http.Cookie{
	Name:  "username",
	Value: "chandu",
}

http.SetCookie(w, cookie)
```

The browser receives a:

```text
Set-Cookie
```

header.

---

# U. `http.DetectContentType`

```go
func DetectContentType(data []byte) string
```

Attempts to determine the MIME type of some data.

Example:

```go
data := []byte("<html>Hello</html>")

contentType := http.DetectContentType(data)

fmt.Println(contentType)
```

It examines at most the first 512 bytes and returns a valid MIME type; if it cannot identify something more specific, it returns `application/octet-stream`.

---

# V. `http.CanonicalHeaderKey`

```go
func CanonicalHeaderKey(s string) string
```

Converts an HTTP header name into canonical form.

Example:

```go
fmt.Println(
	http.CanonicalHeaderKey("content-type"),
)
```

Result:

```text
Content-Type
```

This is mostly useful when dealing with HTTP headers directly.

---

# W. `http.StatusText`

```go
func StatusText(code int) string
```

Returns the text associated with an HTTP status code.

Example:

```go
fmt.Println(http.StatusText(http.StatusNotFound))
```

Output:

```text
Not Found
```

---

# X. `http.MaxBytesReader`

```go
func MaxBytesReader(
	w ResponseWriter,
	r io.ReadCloser,
	n int64,
) io.ReadCloser
```

Limits how much data can be read from an incoming request body.

This is extremely useful for protecting servers from unexpectedly large requests.

Example:

```go
r.Body = http.MaxBytesReader(
	w,
	r.Body,
	1<<20, // 1 MB
)
```

Now the request body is limited to approximately 1 MB.

This is particularly important for APIs that accept uploaded data.

---

# Y. `http.ParseTime`

```go
func ParseTime(text string) (time.Time, error)
```

Parses an HTTP date.

Useful when working with HTTP headers such as:

```text
Date
Last-Modified
If-Modified-Since
```

Example:

```go
t, err := http.ParseTime(
	"Mon, 02 Jan 2006 15:04:05 GMT",
)
```

---

# Z. `http.ProxyFromEnvironment`

```go
func ProxyFromEnvironment(
	req *Request,
) (*url.URL, error)
```

Determines whether an HTTP proxy should be used based on environment variables.

Useful in:

- Corporate networks
- Enterprise environments
- CI/CD systems
- Controlled network environments

---

# AA. `http.ProxyURL`

```go
func ProxyURL(
	fixedURL *url.URL,
) func(*Request) (*url.URL, error)
```

Creates a proxy function that always returns the specified proxy URL.

This is normally used with an HTTP `Transport`.

---

# 5. Important `http.Client` Methods

`http.Client` represents an HTTP client.

Example:

```go
client := &http.Client{}

resp, err := client.Get("https://example.com")
```

The important methods are:

## `Client.Do`

```go
client.Do(req)
```

Sends a custom HTTP request.

This is the most flexible client method.

For example:

```go
req, err := http.NewRequest(
	"GET",
	"https://example.com",
	nil,
)

if err != nil {
	log.Fatal(err)
}

req.Header.Set("Authorization", "Bearer token")

resp, err := client.Do(req)
```

A non-2xx HTTP status does **not** itself cause `Client.Do` to return an error.

Use `Client.Do` when you need control over method, headers, body, authentication, context, or other request properties.

---

## `Client.Get`

```go
client.Get(url)
```

Convenient GET request.

---

## `Client.Head`

```go
client.Head(url)
```

Convenient HEAD request.

---

## `Client.Post`

```go
client.Post(url, contentType, body)
```

Convenient POST request.

---

## `Client.PostForm`

```go
client.PostForm(url, data)
```

Convenient form POST request.

---

## `Client.CloseIdleConnections`

```go
client.CloseIdleConnections()
```

Closes idle keep-alive connections that are no longer being used.

This is generally relevant to long-running applications that need explicit connection cleanup.

---

# 6. `http.NewRequest`

```go
func NewRequest(
	method string,
	url string,
	body io.Reader,
) (*Request, error)
```

Creates an HTTP request.

Example:

```go
req, err := http.NewRequest(
	"GET",
	"https://example.com",
	nil,
)
```

You can then modify it:

```go
req.Header.Set("Accept", "application/json")
```

and send it:

```go
resp, err := http.DefaultClient.Do(req)
```

---

# 7. `http.NewRequestWithContext`

```go
func NewRequestWithContext(
	ctx context.Context,
	method string,
	url string,
	body io.Reader,
) (*Request, error)
```

Creates a request associated with a context.

This is extremely important in production applications.

Example:

```go
ctx, cancel := context.WithTimeout(
	context.Background(),
	5*time.Second,
)

defer cancel()

req, err := http.NewRequestWithContext(
	ctx,
	"GET",
	"https://example.com",
	nil,
)

if err != nil {
	log.Fatal(err)
}

resp, err := http.DefaultClient.Do(req)
```

If the timeout expires, the request can be cancelled.

---

# 8. `http.ReadRequest`

```go
func ReadRequest(
	b *bufio.Reader,
) (*Request, error)
```

Reads an HTTP/1.x request from a buffered reader.

This is a **low-level API**.

Most application developers should not use it directly; normally the Go HTTP server handles request parsing for you.

---

# 9. Important `http.Request` Methods

The `Request` type has many useful methods.

## `r.Context()`

Gets the request's context.

```go
ctx := r.Context()
```

Useful for cancellation, deadlines, and passing request-scoped values.

---

## `r.WithContext(ctx)`

Returns a shallow copy of the request with a new context.

```go
r = r.WithContext(ctx)
```

---

## `r.Clone(ctx)`

Creates a deeper copy of the request with a new context.

```go
newRequest := r.Clone(ctx)
```

---

## `r.Cookie(name)`

Gets a particular cookie.

```go
cookie, err := r.Cookie("session")
```

---

## `r.Cookies()`

Gets all cookies.

```go
cookies := r.Cookies()
```

---

## `r.CookiesNamed(name)`

Gets all cookies with a specified name.

---

## `r.AddCookie(cookie)`

Adds a cookie to an outgoing request.

---

## `r.BasicAuth()`

Extracts HTTP Basic Authentication credentials.

```go
username, password, ok := r.BasicAuth()
```

---

## `r.SetBasicAuth()`

Sets Basic Authentication credentials on an outgoing request.

```go
req.SetBasicAuth("admin", "password")
```

**Important:** Basic Auth should normally be used over HTTPS because credentials are otherwise exposed in transit.

---

## `r.ParseForm()`

Parses URL query parameters and form data.

Example:

```go
err := r.ParseForm()
```

Then:

```go
value := r.Form.Get("name")
```

---

## `r.FormValue(name)`

Conveniently retrieves a form value.

```go
name := r.FormValue("name")
```

---

## `r.PostFormValue(name)`

Retrieves a POST form value.

```go
name := r.PostFormValue("name")
```

---

## `r.ParseMultipartForm(maxMemory)`

Parses multipart form data, commonly used for file uploads.

```go
err := r.ParseMultipartForm(10 << 20)
```

---

## `r.FormFile(key)`

Retrieves an uploaded file.

```go
file, header, err := r.FormFile("photo")
```

---

## `r.MultipartReader()`

Returns a multipart reader for processing multipart request data.

This is useful when you need more direct control over multipart data processing.

---

## `r.UserAgent()`

Returns the client's `User-Agent`.

```go
fmt.Println(r.UserAgent())
```

---

## `r.Referer()`

Returns the HTTP `Referer` header.

---

## `r.PathValue(name)`

Retrieves a wildcard path value from the `ServeMux`.

For example, with a pattern such as:

```text
/users/{id}
```

you can retrieve:

```go
id := r.PathValue("id")
```

Modern Go's `ServeMux` supports method-aware patterns and path wildcards.

---

# 10. `http.Header`

`Header` represents HTTP headers:

```go
type Header map[string][]string
```

For example:

```go
w.Header().Set("Content-Type", "application/json")
```

Important methods include:

### `Set`

```go
w.Header().Set("Content-Type", "application/json")
```

Replaces the value.

### `Add`

```go
w.Header().Add("X-Custom", "value")
```

Adds another value.

### `Get`

```go
value := w.Header().Get("Content-Type")
```

### `Del`

```go
w.Header().Del("X-Custom")
```

### `Values`

```go
values := w.Header().Values("Accept")
```

Returns all values for a header.

### `Clone`

```go
copy := w.Header().Clone()
```

Creates a copy of the header map.

### `Write`

Writes headers to an `io.Writer`.

### `WriteSubset`

Writes headers while excluding selected headers.

---

# 11. `http.HandlerFunc`

`HandlerFunc` adapts an ordinary function into an `http.Handler`.

For example:

```go
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello")
}
```

This function has the correct signature and can be converted into:

```go
http.HandlerFunc(hello)
```

That's essentially what makes this convenient:

```go
http.HandleFunc("/hello", hello)
```

---

# 12. `http.Server`

For production servers, you will often want an explicit `http.Server`.

Example:

```go
server := &http.Server{
	Addr:         ":8080",
	ReadTimeout:  10 * time.Second,
	WriteTimeout: 10 * time.Second,
	IdleTimeout:  60 * time.Second,
}

log.Fatal(server.ListenAndServe())
```

A custom `Server` allows you to configure things such as:

- Address
- Handler
- Read timeout
- Write timeout
- Idle timeout
- Maximum header size
- TLS configuration
- Graceful shutdown behavior

The official documentation recommends a custom `Server` when more control over server behavior is needed.

## Important `Server` methods

### `server.ListenAndServe()`

Starts the HTTP server.

```go
server.ListenAndServe()
```

### `server.ListenAndServeTLS()`

Starts an HTTPS server.

### `server.Serve(listener)`

Serves HTTP connections from a listener.

### `server.ServeTLS(listener, certFile, keyFile)`

Serves HTTPS connections from a listener.

### `server.Shutdown(ctx)`

Gracefully shuts down the server.

This is extremely important for production applications.

Conceptually:

```text
Receive shutdown signal
        |
        v
Stop accepting new requests
        |
        v
Allow active requests to finish
        |
        v
Shutdown
```

### `server.Close()`

Immediately closes active connections rather than performing graceful shutdown.

### `server.RegisterOnShutdown(f)`

Registers a function to run when shutdown begins.

---

# 13. `http.ServeMux`

`ServeMux` is Go's HTTP request multiplexer/router.

Example:

```go
mux := http.NewServeMux()

mux.HandleFunc("/users", usersHandler)
mux.HandleFunc("/products", productsHandler)

server := &http.Server{
	Addr:    ":8080",
	Handler: mux,
}

log.Fatal(server.ListenAndServe())
```

Modern Go's `ServeMux` supports patterns involving:

- HTTP methods
- Hosts
- Path segments
- Wildcards

For example:

```go
mux.HandleFunc("GET /users/{id}", getUser)
```

Then:

```go
id := r.PathValue("id")
```

can retrieve the wildcard value.

---

# 14. Three Common Beginner Mistakes

## Mistake 1: Forgetting to close `resp.Body`

Bad:

```go
resp, err := http.Get(url)

if err != nil {
	log.Fatal(err)
}

body, _ := io.ReadAll(resp.Body)
```

Better:

```go
resp, err := http.Get(url)

if err != nil {
	log.Fatal(err)
}

defer resp.Body.Close()

body, err := io.ReadAll(resp.Body)
```

Why?

HTTP connections can be reused, and failing to properly consume/close response bodies can cause resource and connection-management problems.

---

## Mistake 2: Assuming HTTP 404/500 is a Go `error`

Beginners sometimes write:

```go
resp, err := http.Get(url)

if err != nil {
	// "The server returned 404!"
}
```

But:

```text
HTTP 404
```

is not necessarily a Go-level `error`.

You should check:

```go
if resp.StatusCode != http.StatusOK {
	// Handle HTTP status
}
```

For example:

```go
if resp.StatusCode >= 400 {
	fmt.Println("HTTP error:", resp.Status)
}
```

A non-2xx response does not itself cause `Get` or `Client.Do` to return an error.

---

## Mistake 3: Creating a new HTTP client for every request

You may see:

```go
for _, url := range urls {
	client := &http.Client{}

	client.Get(url)
}
```

This isn't generally the best design.

Instead, create a client and reuse it:

```go
client := &http.Client{}

for _, url := range urls {
	resp, err := client.Get(url)
	// ...
}
```

Reusing clients allows the underlying transport to manage persistent connections efficiently.

---

# 15. Two Real-World Applications

## Application 1: REST API / Microservice

Imagine you're building an e-commerce backend:

```text
Client
   |
   | GET /products
   v
Go HTTP Server
   |
   +---- Database
   |
   +---- Product Service
   |
   +---- Authentication
   |
   v
JSON Response
```

`net/http` can handle:

```text
GET     /products
GET     /products/{id}
POST    /products
PUT     /products/{id}
DELETE  /products/{id}
```

This makes `net/http` particularly useful for Go-based REST APIs and microservices.

---

## Application 2: Calling External APIs

Suppose your application needs information from another service:

```text
Your Go Application
        |
        | HTTP request
        v
Payment API
        |
        v
JSON response
```

You can use:

```go
req, err := http.NewRequest(
	"GET",
	"https://api.example.com/payments",
	nil,
)

req.Header.Set("Authorization", "Bearer token")

resp, err := client.Do(req)
```

This pattern is extremely common when integrating:

- Payment services
- Cloud APIs
- Authentication services
- Shipping APIs
- Weather APIs
- Internal microservices

---

# 16. Three Practice Exercises

## Exercise 1 — Beginner: Simple HTTP Server

Create a Go HTTP server on port `8080`.

Requirements:

1. Create a `/hello` endpoint.
2. Respond with a greeting.
3. Create a `/about` endpoint.
4. Return a different message from `/about`.
5. Print the HTTP method and requested path to the terminal.
6. Handle an unknown route with a `404 Not Found` response.

**Do not use a third-party router.**

---

## Exercise 2 — Intermediate: JSON REST API

Build a small in-memory REST API for books.

Each book should contain:

```text
ID
Title
Author
Price
```

Your API should support:

```text
GET    /books
GET    /books/{id}
POST   /books
PUT    /books/{id}
DELETE /books/{id}
```

Requirements:

- Use `net/http`.
- Return JSON responses.
- Set appropriate `Content-Type` headers.
- Return appropriate HTTP status codes.
- Handle invalid JSON.
- Handle nonexistent book IDs.
- Validate required fields.
- Do not use a database.

---

## Exercise 3 — Advanced: Production-Style API Client and Server

Build a small service consisting of:

```text
Client
   |
   v
Go API Server
   |
   v
External HTTP API
```

Requirements:

1. Create a Go HTTP server.
2. Implement an endpoint such as:
   ```text
   GET /weather
   ```
3. The handler should make an outgoing HTTP request to another API.
4. Use `http.Client`.
5. Create outgoing requests using `http.NewRequestWithContext`.
6. Configure a request timeout.
7. Forward useful errors to the client.
8. Return JSON.
9. Add appropriate HTTP status codes.
10. Limit incoming request body sizes where appropriate.
11. Implement graceful server shutdown.
12. Add server timeouts.
13. Reuse the HTTP client.
14. Log important request information.
15. Make sure response bodies are properly closed.

The goal is to combine the major concepts you've learned rather than simply creating another basic HTTP server.

---

# 17. A Useful Mental Model

When learning `net/http`, think of it as two sides:

```text
                 net/http
                    |
          +---------+---------+
          |                   |
       SERVER               CLIENT
          |                   |
    http.Server          http.Client
          |                   |
     ServeMux              Request
          |                   |
      Handler                Do()
          |                   |
     Request ------------> Response
          |
    ResponseWriter
```

For a **server**, remember:

```text
Request
   ↓
ServeMux
   ↓
Handler
   ↓
ResponseWriter
   ↓
Client
```

For a **client**, remember:

```text
Request
   ↓
http.Client
   ↓
HTTP Server
   ↓
Response
```

Once this mental model is clear, most of the `net/http` API becomes much easier to understand.

---

# 18. What You Should Learn First

Because `net/http` is large, I recommend learning it in this order:

## Level 1 — Fundamentals

```text
http.HandleFunc
http.ListenAndServe
http.Request
http.ResponseWriter
http.Response
```

## Level 2 — Client

```text
http.Get
http.Post
http.NewRequest
http.Client
client.Do
```

## Level 3 — Routing and HTTP data

```text
ServeMux
Request.URL
Request.Header
Request.Body
Request.FormValue
Request.PathValue
ResponseWriter.Header
```

## Level 4 — Production HTTP

```text
http.Server
timeouts
context.Context
NewRequestWithContext
TLS
graceful shutdown
MaxBytesReader
```

## Level 5 — Advanced

```text
Transport
RoundTripper
HTTP/2
proxies
connection management
streaming
hijacking
custom middleware
```

The `net/http` API is considerably larger than just the functions above; it also contains types such as `Transport`, `Cookie`, `ServeMux`, `Server`, `Client`, `Response`, `Request`, `ResponseWriter`, `RoundTripper`, `Flusher`, `Hijacker`, `Pusher`, and HTTP protocol configuration APIs.

---

# 🤔 Thought-provoking question

Imagine you have a Go API that receives **10,000 requests per second**, and each request calls a slow external API.

**Would simply increasing the server's timeout make the application more reliable, or could it actually make the system worse?**

Think about:

- Goroutines
- Connection reuse
- Context cancellation
- Timeouts
- Memory
- Backpressure
- What happens when the external API becomes unavailable

That question gets to the heart of why understanding `net/http` is about much more than simply learning how to create a web server.
