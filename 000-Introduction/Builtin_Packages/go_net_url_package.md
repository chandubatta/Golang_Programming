# Go `net/url` Package

The Go `net/url` package provides types and functions for **parsing, constructing, modifying, escaping, and resolving URLs/URIs**. It is especially useful when working with HTTP APIs, query parameters, redirects, links, and user-provided URLs. The package generally follows RFC 3986.

**Official documentation:** https://pkg.go.dev/net/url

---

## 1. What is the `net/url` package?

Import it with:

```go
import "net/url"
```

A URL can be thought of as:

```text
scheme://userinfo@host:port/path?query#fragment
```

For example:

```text
https://john:secret@example.com:8080/products/laptop?page=2#reviews
```

The `net/url` package lets you work with each component independently:

| Component | Example | Go field/function |
|---|---|---|
| Scheme | `https` | `URL.Scheme` |
| User info | `john:secret` | `URL.User` |
| Host | `example.com` | `URL.Host` |
| Port | `8080` | `URL.Port()` |
| Path | `/products/laptop` | `URL.Path` |
| Query | `page=2` | `URL.RawQuery`, `URL.Query()` |
| Fragment | `reviews` | `URL.Fragment` |

### When is it commonly used?

You will frequently use `net/url` when:

- Parsing URLs received from users or HTTP requests.
- Building API URLs.
- Adding or modifying query parameters.
- Encoding special characters.
- Decoding URL-encoded values.
- Extracting domains and ports.
- Resolving relative URLs.
- Building redirects.
- Working with pagination, filtering, and search parameters.

---

# 2. Simple Example

Here is a beginner-friendly example that parses a URL and modifies its query parameters:

```go
package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	// Parse a URL
	u, err := url.Parse("https://example.com/products?category=laptop")
	if err != nil {
		log.Fatal(err)
	}

	// Read URL components
	fmt.Println("Scheme:", u.Scheme)
	fmt.Println("Host:", u.Host)
	fmt.Println("Path:", u.Path)

	// Get query parameters
	query := u.Query()

	fmt.Println("Category:", query.Get("category"))

	// Add another query parameter
	query.Set("page", "2")

	// Put the modified query back into the URL
	u.RawQuery = query.Encode()

	fmt.Println("Final URL:", u.String())
}
```

Output:

```text
Scheme: https
Host: example.com
Path: /products
Category: laptop
Final URL: https://example.com/products?category=laptop&page=2
```

### Important pattern

One of the most useful patterns to remember is:

```go
u, err := url.Parse(rawURL)

query := u.Query()
query.Set("key", "value")

u.RawQuery = query.Encode()
```

This pattern appears constantly in real-world Go applications.

---

# 3. Functions and Methods in `net/url`

The current Go documentation lists package-level functions, `URL` methods, `Userinfo` methods, `Values` methods, and error-related methods.

---

# A. Package-Level Functions

## 1. `url.Parse`

```go
func Parse(rawURL string) (*URL, error)
```

Parses a URL string into a `*url.URL`.

Example:

```go
u, err := url.Parse("https://example.com/products?id=10")

if err != nil {
	log.Fatal(err)
}

fmt.Println(u.Scheme)
fmt.Println(u.Host)
fmt.Println(u.Path)
```

Output:

```text
https
example.com
/products
```

### When to use it

Use `Parse` when you have a URL represented as a string and need to inspect or modify it.

---

## 2. `url.ParseRequestURI`

```go
func ParseRequestURI(rawURL string) (*URL, error)
```

This is intended for URLs received as part of an HTTP request.

It interprets the input as either:

- an absolute URI, or
- an absolute path.

It assumes the input does not contain a fragment because browsers normally remove URL fragments before sending HTTP requests.

Example:

```go
u, err := url.ParseRequestURI("/products?page=2")
if err != nil {
	log.Fatal(err)
}

fmt.Println(u.Path)
fmt.Println(u.RawQuery)
```

Output:

```text
/products
page=2
```

### `Parse` vs `ParseRequestURI`

Think:

```text
url.Parse            → general URL parsing
url.ParseRequestURI  → parsing an HTTP request URI
```

---

## 3. `url.JoinPath`

```go
func JoinPath(base string, elem ...string) (result string, err error)
```

Joins URL path elements while cleaning `.` and `..` components and redundant slashes.

Example:

```go
result, err := url.JoinPath(
	"https://example.com/api",
	"users",
	"123",
	"profile",
)

if err != nil {
	log.Fatal(err)
}

fmt.Println(result)
```

Output:

```text
https://example.com/api/users/123/profile
```

This is useful when constructing API endpoints dynamically.

---

## 4. `url.PathEscape`

```go
func PathEscape(s string) string
```

Escapes a string so that it can safely be used as a URL **path segment**.

Example:

```go
name := "hello world"

escaped := url.PathEscape(name)

fmt.Println(escaped)
```

Output:

```text
hello%20world
```

For example:

```go
path := "/users/" + url.PathEscape("John Smith")
```

Result:

```text
/users/John%20Smith
```

### Important

`PathEscape` is for **path components**, not query parameters.

---

## 5. `url.PathUnescape`

```go
func PathUnescape(s string) (string, error)
```

Reverses path escaping.

```go
value, err := url.PathUnescape("John%20Smith")

if err != nil {
	log.Fatal(err)
}

fmt.Println(value)
```

Output:

```text
John Smith
```

It returns an error if the input contains invalid escaping.

---

## 6. `url.QueryEscape`

```go
func QueryEscape(s string) string
```

Escapes a string for use inside a URL query.

```go
value := "hello world"

fmt.Println(url.QueryEscape(value))
```

Output:

```text
hello+world
```

Another example:

```go
fmt.Println(url.QueryEscape("Go & HTTP"))
```

Possible output:

```text
Go+%26+HTTP
```

Notice that `&` is escaped because `&` has special meaning in query strings.

---

## 7. `url.QueryUnescape`

```go
func QueryUnescape(s string) (string, error)
```

Decodes a query-escaped string.

```go
value, err := url.QueryUnescape("hello+world")

if err != nil {
	log.Fatal(err)
}

fmt.Println(value)
```

Output:

```text
hello world
```

It also decodes percent-encoded characters.

---

# B. `URL` Type

The central type in `net/url` is:

```go
type URL struct {
	...
}
```

A `URL` represents a parsed URL/URI reference.

For example:

```go
u, _ := url.Parse(
	"https://example.com/products?page=2#details",
)
```

You can access:

```go
u.Scheme
u.Host
u.Path
u.RawQuery
u.Fragment
```

---

# C. `URL` Methods

## 8. `u.AppendBinary`

```go
func (u *URL) AppendBinary(b []byte) ([]byte, error)
```

Appends the encoded URL representation to an existing byte slice.

Example:

```go
u, _ := url.Parse("https://example.com")

data, err := u.AppendBinary(nil)
if err != nil {
	log.Fatal(err)
}

fmt.Println(string(data))
```

Output:

```text
https://example.com
```

This is useful when you are already working with byte buffers.

---

## 9. `u.Clone`

```go
func (u *URL) Clone() *URL
```

Creates a deep copy of a `URL`.

Example:

```go
original, _ := url.Parse("https://example.com/products")

copy := original.Clone()

copy.Path = "/users"

fmt.Println(original)
fmt.Println(copy)
```

The original URL remains unchanged.

This is useful when you want to create several variations of the same URL.

---

## 10. `u.EscapedFragment`

```go
func (u *URL) EscapedFragment() string
```

Returns the escaped representation of the URL fragment.

Example:

```go
u, _ := url.Parse("https://example.com/#hello%20world")

fmt.Println(u.Fragment)
fmt.Println(u.EscapedFragment())
```

Conceptually:

```text
Fragment:         hello world
EscapedFragment:  hello%20world
```

The package stores `Fragment` in decoded form while retaining information necessary to reproduce valid original escaping.

---

## 11. `u.EscapedPath`

```go
func (u *URL) EscapedPath() string
```

Returns the escaped version of the URL path.

Example:

```go
u, _ := url.Parse("https://example.com/foo%2Fbar")

fmt.Println(u.Path)
fmt.Println(u.EscapedPath())
```

The distinction is important because:

```text
/foo%2Fbar
```

can represent an encoded slash, while:

```text
/foo/bar
```

contains an actual path separator.

`Path` is decoded; `EscapedPath()` preserves the appropriate escaped representation.

---

## 12. `u.Hostname`

```go
func (u *URL) Hostname() string
```

Returns the hostname without the port.

```go
u, _ := url.Parse("https://example.com:8080/products")

fmt.Println(u.Hostname())
```

Output:

```text
example.com
```

It also handles IPv6 addresses:

```go
u, _ := url.Parse(
	"https://[2001:db8::1]:8080",
)

fmt.Println(u.Hostname())
```

The returned hostname doesn't include the IPv6 brackets.

---

## 13. `u.IsAbs`

```go
func (u *URL) IsAbs() bool
```

Checks whether the URL is absolute.

An absolute URL has a non-empty scheme.

```go
u, _ := url.Parse("https://example.com")

fmt.Println(u.IsAbs())
```

Output:

```text
true
```

Whereas:

```go
u, _ := url.Parse("/products")

fmt.Println(u.IsAbs())
```

returns:

```text
false
```

---

## 14. `u.JoinPath`

```go
func (u *URL) JoinPath(elem ...string) *URL
```

Adds path elements to an existing URL.

```go
u, _ := url.Parse("https://example.com/api")

result := u.JoinPath("users", "123")

fmt.Println(result)
```

Output:

```text
https://example.com/api/users/123
```

Unlike manually concatenating strings, this handles path cleaning.

---

## 15. `u.MarshalBinary`

```go
func (u *URL) MarshalBinary() ([]byte, error)
```

Returns a byte representation of the URL.

Example:

```go
u, _ := url.Parse("https://example.com")

data, err := u.MarshalBinary()
if err != nil {
	log.Fatal(err)
}

fmt.Println(string(data))
```

Output:

```text
https://example.com
```

This is useful when an API requires binary marshaling.

---

## 16. `u.Parse`

```go
func (u *URL) Parse(ref string) (*URL, error)
```

Parses a URL relative to another URL.

Example:

```go
base, _ := url.Parse("https://example.com")

result, err := base.Parse("/products/123")
if err != nil {
	log.Fatal(err)
}

fmt.Println(result)
```

Output:

```text
https://example.com/products/123
```

This is convenient for resolving links relative to a base URL.

---

## 17. `u.Port`

```go
func (u *URL) Port() string
```

Returns the port.

```go
u, _ := url.Parse("https://example.com:8080")

fmt.Println(u.Port())
```

Output:

```text
8080
```

If there is no valid numeric port, it returns an empty string.

---

## 18. `u.Query`

```go
func (u *URL) Query() Values
```

Parses the URL's query parameters.

Example:

```go
u, _ := url.Parse(
	"https://example.com/search?q=golang&page=2",
)

query := u.Query()

fmt.Println(query.Get("q"))
fmt.Println(query.Get("page"))
```

Output:

```text
golang
2
```

### Important warning

`Query()` silently discards malformed query pairs.

If you need to explicitly detect malformed query data, use:

```go
url.ParseQuery(...)
```

instead.

---

## 19. `u.Redacted`

```go
func (u *URL) Redacted() string
```

Returns the URL while hiding a password in URL user information.

Example:

```go
u, _ := url.Parse(
	"https://john:secret@example.com",
)

fmt.Println(u.Redacted())
```

Output:

```text
https://john:xxxxx@example.com
```

This is particularly useful when logging URLs.

### Security lesson

Avoid logging credentials directly.

Embedding authentication information in URLs is generally discouraged for security reasons.

---

## 20. `u.RequestURI`

```go
func (u *URL) RequestURI() string
```

Returns the path and query portion appropriate for an HTTP request.

```go
u, _ := url.Parse(
	"https://example.com/products?page=2",
)

fmt.Println(u.RequestURI())
```

Output:

```text
/products?page=2
```

Notice that it doesn't return:

```text
https://example.com
```

It returns the request-target portion.

---

## 21. `u.ResolveReference`

```go
func (u *URL) ResolveReference(ref *URL) *URL
```

Resolves a relative URL against a base URL according to RFC 3986.

Example:

```go
base, _ := url.Parse(
	"https://example.com/docs/",
)

ref, _ := url.Parse("guide.html")

result := base.ResolveReference(ref)

fmt.Println(result)
```

Output:

```text
https://example.com/docs/guide.html
```

This is extremely useful when processing links from HTML pages, APIs, or documentation.

---

## 22. `u.String`

```go
func (u *URL) String() string
```

Reconstructs the URL from its components.

Example:

```go
u := &url.URL{
	Scheme: "https",
	Host:   "example.com",
	Path:   "/products",
}

fmt.Println(u.String())
```

Output:

```text
https://example.com/products
```

You'll use `String()` constantly when converting a `URL` back to text.

Also note that:

```go
fmt.Println(u)
```

works because `URL` implements the relevant string representation.

---

## 23. `u.UnmarshalBinary`

```go
func (u *URL) UnmarshalBinary(text []byte) error
```

Reconstructs a `URL` from its byte representation.

Example:

```go
var u url.URL

err := u.UnmarshalBinary(
	[]byte("https://example.com/products"),
)

if err != nil {
	log.Fatal(err)
}

fmt.Println(u.String())
```

Output:

```text
https://example.com/products
```

This complements `MarshalBinary()`.

---

# D. `Userinfo`

`Userinfo` represents username/password information embedded in a URL.

For example:

```text
https://john:secret@example.com
```

The user information is:

```text
john:secret
```

The Go documentation describes this as an immutable representation of username/password information and recommends against using password-bearing URLs except for legacy situations.

---

## 24. `url.User`

```go
func User(username string) *Userinfo
```

Creates URL user information containing only a username.

```go
u := url.User("john")

fmt.Println(u)
```

Output:

```text
john
```

---

## 25. `url.UserPassword`

```go
func UserPassword(username, password string) *Userinfo
```

Creates URL user information containing username and password.

```go
u := url.UserPassword("john", "secret")

fmt.Println(u)
```

Output:

```text
john:secret
```

Again, embedding passwords in URLs is generally discouraged because credentials can be exposed through logs, browser history, monitoring systems, and other places.

---

## 26. `u.Username`

```go
func (u *Userinfo) Username() string
```

Returns the username.

```go
user := url.UserPassword("john", "secret")

fmt.Println(user.Username())
```

Output:

```text
john
```

---

## 27. `u.Password`

```go
func (u *Userinfo) Password() (string, bool)
```

Returns:

1. the password
2. whether a password exists

Example:

```go
user := url.UserPassword("john", "secret")

password, ok := user.Password()

fmt.Println(password)
fmt.Println(ok)
```

Output:

```text
secret
true
```

If there is no password:

```go
user := url.User("john")

password, ok := user.Password()

fmt.Println(password)
fmt.Println(ok)
```

Output:

```text

false
```

---

## 28. `u.String`

```go
func (u *Userinfo) String() string
```

Returns encoded user information.

```go
user := url.UserPassword("john", "secret")

fmt.Println(user.String())
```

Output:

```text
john:secret
```

---

# E. `Values`

One of the most important types in `net/url` is:

```go
type Values map[string][]string
```

It represents URL query parameters and form values.

For example:

```text
?name=John&age=30&tag=go&tag=http
```

can be represented as:

```go
url.Values{
	"name": {"John"},
	"age":  {"30"},
	"tag":  {"go", "http"},
}
```

The important thing to understand is:

> One key can have multiple values.

---

## 29. `url.ParseQuery`

```go
func ParseQuery(query string) (Values, error)
```

Parses a query string.

Example:

```go
values, err := url.ParseQuery(
	"name=John&age=30&tag=go&tag=http",
)

if err != nil {
	log.Fatal(err)
}

fmt.Println(values.Get("name"))
fmt.Println(values.Get("age"))
fmt.Println(values["tag"])
```

Output:

```text
John
30
[go http]
```

### `ParseQuery` vs `URL.Query()`

If you already have a URL:

```go
u.Query()
```

is convenient.

If you have a raw query string:

```go
url.ParseQuery(query)
```

is appropriate.

And importantly, `ParseQuery` gives you an error for malformed query input, whereas `URL.Query()` silently discards malformed value pairs.

---

# F. `Values` Methods

## 30. `v.Add`

```go
func (v Values) Add(key, value string)
```

Adds a value without replacing existing values.

```go
v := url.Values{}

v.Add("tag", "go")
v.Add("tag", "http")

fmt.Println(v["tag"])
```

Output:

```text
[go http]
```

Use `Add` when multiple values for the same key are expected.

---

## 31. `v.Clone`

```go
func (vs Values) Clone() Values
```

Creates a deep copy of `Values`.

Example:

```go
original := url.Values{}

original.Set("name", "John")

copy := original.Clone()

copy.Set("name", "Alice")

fmt.Println(original.Get("name"))
fmt.Println(copy.Get("name"))
```

Output:

```text
John
Alice
```

This is useful when you need an independent set of query/form parameters.

---

## 32. `v.Del`

```go
func (v Values) Del(key string)
```

Deletes all values associated with a key.

```go
v := url.Values{}

v.Set("name", "John")
v.Set("age", "30")

v.Del("age")

fmt.Println(v)
```

---

## 33. `v.Encode`

```go
func (v Values) Encode() string
```

Converts `Values` into URL-encoded query format.

```go
v := url.Values{}

v.Set("name", "John Smith")
v.Set("language", "Go")

fmt.Println(v.Encode())
```

Output will be similar to:

```text
language=Go&name=John+Smith
```

Keys are sorted during encoding.

This is one of the most commonly used methods in API development.

---

## 34. `v.Get`

```go
func (v Values) Get(key string) string
```

Returns the first value associated with a key.

```go
v := url.Values{}

v.Add("tag", "go")
v.Add("tag", "http")

fmt.Println(v.Get("tag"))
```

Output:

```text
go
```

If you need all values:

```go
fmt.Println(v["tag"])
```

Output:

```text
[go http]
```

---

## 35. `v.Has`

```go
func (v Values) Has(key string) bool
```

Checks whether a key exists.

```go
v := url.Values{}

v.Set("name", "John")

fmt.Println(v.Has("name"))
fmt.Println(v.Has("age"))
```

Output:

```text
true
false
```

This can be useful when you need to distinguish:

```text
parameter doesn't exist
```

from:

```text
parameter exists but has an empty value
```

---

## 36. `v.Set`

```go
func (v Values) Set(key, value string)
```

Sets a key to exactly one value.

```go
v := url.Values{}

v.Set("name", "John")
v.Set("name", "Alice")

fmt.Println(v.Get("name"))
```

Output:

```text
Alice
```

Unlike `Add`, `Set` replaces the existing values.

### Remember

```text
Set → replace
Add → append
```

This distinction is extremely important.

---

# G. Error Types

The package also defines error types.

## 37. `url.Error`

```go
type Error struct {
	Op  string
	URL string
	Err error
}
```

It provides information about an operation, URL, and underlying error.

It implements:

```go
Error()
Temporary()
Timeout()
Unwrap()
```

### `Error()`

```go
func (e *Error) Error() string
```

Returns the error message.

### `Temporary()`

```go
func (e *Error) Temporary() bool
```

Reports whether the underlying error is temporary.

### `Timeout()`

```go
func (e *Error) Timeout() bool
```

Reports whether the underlying error represents a timeout.

### `Unwrap()`

```go
func (e *Error) Unwrap() error
```

Returns the underlying error.

This allows standard Go error handling such as:

```go
errors.Is(...)
```

and:

```go
errors.As(...)
```

---

# H. `EscapeError`

## 38. `EscapeError.Error`

```go
func (e EscapeError) Error() string
```

`EscapeError` represents an invalid URL escape sequence.

For example, an incomplete percent escape can cause this type of error.

The `Error()` method converts it to a human-readable error string.

---

# I. `InvalidHostError`

## 39. `InvalidHostError.Error`

```go
func (e InvalidHostError) Error() string
```

Represents an invalid host-related URL error and provides its error message.

These error types are generally encountered while handling malformed URLs rather than being created manually.

---

# 4. Three Common Beginner Mistakes

## Mistake 1: Manually concatenating URLs

Beginners often write:

```go
url := baseURL + "?name=" + name
```

This becomes dangerous when `name` contains characters such as:

```text
&
?
=
space
#
%
```

### Better

Use `url.Values`:

```go
query := url.Values{}

query.Set("name", name)

finalURL := baseURL + "?" + query.Encode()
```

Or parse an existing URL and modify its query:

```go
u, _ := url.Parse(baseURL)

q := u.Query()
q.Set("name", name)

u.RawQuery = q.Encode()
```

---

## Mistake 2: Confusing `PathEscape` and `QueryEscape`

These functions serve different purposes.

```go
url.PathEscape(...)
```

is for a **path component**.

```go
url.QueryEscape(...)
```

is for a **query value**.

For example:

```text
/products/John%20Smith
```

and:

```text
/products?name=John+Smith
```

have different URL contexts.

Don't blindly use one escaping function everywhere.

---

## Mistake 3: Using `Query()` when you need query-validation errors

Consider:

```go
q := u.Query()
```

`Query()` is convenient, but malformed query pairs can be silently discarded.

If validation matters, use:

```go
q, err := url.ParseQuery(u.RawQuery)
```

and handle the error.

---

# 5. Two Real-World Applications

## Application 1: Building API URLs

Imagine an API:

```text
https://api.example.com/products
```

You need:

```text
category=laptop
page=2
limit=20
sort=price
```

Instead of manually constructing:

```text
...?category=laptop&page=2&limit=20&sort=price
```

you can use:

```go
q := url.Values{}

q.Set("category", "laptop")
q.Set("page", "2")
q.Set("limit", "20")
q.Set("sort", "price")

apiURL := "https://api.example.com/products?" + q.Encode()
```

This avoids many escaping bugs.

---

## Application 2: Processing links and redirects

Suppose a web crawler receives:

```text
https://example.com/blog/
```

and finds:

```text
../products/item.html
```

You can resolve the relative URL:

```go
base, _ := url.Parse("https://example.com/blog/")
ref, _ := url.Parse("../products/item.html")

absolute := base.ResolveReference(ref)

fmt.Println(absolute)
```

Result:

```text
https://example.com/products/item.html
```

This is useful in:

- web crawlers
- documentation systems
- link processors
- redirect systems
- HTML parsers
- web scraping applications

---

# 6. Three Progressively Challenging Exercises

## Exercise 1 — URL Inspector

Write a Go program that accepts this URL:

```text
https://example.com:8080/products/laptop?brand=lenovo&price=800#reviews
```

Use `net/url` to display:

- Scheme
- Host
- Hostname
- Port
- Path
- Query parameters
- Fragment
- Whether the URL is absolute

**Do not manually split the URL using string functions.**

---

## Exercise 2 — API URL Builder

Create a program that builds a product-search URL.

The program should accept:

- keyword
- category
- page number
- minimum price
- maximum price
- multiple product tags

Build the URL using:

```go
url.Values
```

Requirements:

- Properly encode special characters.
- Allow multiple values for the same `tag` parameter.
- Use `Set` where only one value should exist.
- Use `Add` for multiple tags.
- Print the final URL.

---

## Exercise 3 — Mini URL Resolver and Query Processor

Build a small program that receives:

### Base URL

```text
https://example.com/store/products/
```

### Relative links

```text
laptop.html
../phones/iphone.html
../accessories/?category=computer&sort=price
```

For each relative link:

1. Parse it.
2. Resolve it against the base URL.
3. Extract the hostname.
4. Extract the path.
5. Read its query parameters.
6. Add a new query parameter called `source=web`.
7. Print the resulting absolute URL.
8. Ensure all query parameters are correctly encoded.

Then extend your program so it can detect malformed query strings and report errors rather than silently ignoring them.

---

# 7. A Useful Mental Model

When learning `net/url`, think about it in five layers:

```text
                    net/url
                       │
       ┌───────────────┼────────────────┐
       │               │                │
    Parse            Modify           Encode
       │               │                │
       ▼               ▼                ▼
   URL struct       URL fields      Escaping
       │               │                │
       ├───────────────┼────────────────┤
       │               │                │
       ▼               ▼                ▼
    Path            Query            Values
```

The most important relationship to remember is:

```text
url.Parse()
     ↓
   *url.URL
     ↓
u.Query()
     ↓
url.Values
     ↓
q.Set() / q.Add() / q.Del()
     ↓
q.Encode()
     ↓
u.RawQuery
     ↓
u.String()
```

If you become comfortable with this flow, you'll understand a large portion of practical `net/url` usage.

---

# 8. Quick Function Cheat Sheet

| Function / Method | Main purpose |
|---|---|
| `url.Parse()` | Parse a URL |
| `url.ParseRequestURI()` | Parse an HTTP request URI |
| `url.JoinPath()` | Join URL paths |
| `url.PathEscape()` | Escape a path component |
| `url.PathUnescape()` | Decode a path component |
| `url.QueryEscape()` | Escape a query value |
| `url.QueryUnescape()` | Decode a query value |
| `u.AppendBinary()` | Append encoded URL to bytes |
| `u.Clone()` | Deep-copy a URL |
| `u.EscapedFragment()` | Get escaped fragment |
| `u.EscapedPath()` | Get escaped path |
| `u.Hostname()` | Get hostname |
| `u.IsAbs()` | Check whether URL is absolute |
| `u.JoinPath()` | Add path elements |
| `u.MarshalBinary()` | Convert URL to bytes |
| `u.Parse()` | Parse relative to a base URL |
| `u.Port()` | Get port |
| `u.Query()` | Get query values |
| `u.Redacted()` | Hide password |
| `u.RequestURI()` | Get HTTP request URI |
| `u.ResolveReference()` | Resolve relative URL |
| `u.String()` | Convert URL to string |
| `u.UnmarshalBinary()` | Reconstruct URL from bytes |
| `url.User()` | Create username information |
| `url.UserPassword()` | Create username/password information |
| `u.Username()` | Get username |
| `u.Password()` | Get password |
| `u.String()` | Encode user information |
| `url.ParseQuery()` | Parse raw query |
| `v.Add()` | Add another value |
| `v.Clone()` | Copy query values |
| `v.Del()` | Delete key |
| `v.Encode()` | Encode query |
| `v.Get()` | Get first value |
| `v.Has()` | Check key existence |
| `v.Set()` | Replace/set value |

**Official current API reference:** https://pkg.go.dev/net/url

---

# Thought-Provoking Question 🤔

Suppose your Go application receives this URL from an untrusted user:

```text
https://example.com/download?file=../../private/config.txt
```

**If you simply parse the URL successfully with `net/url`, does that mean the URL is safe to use?**

Think about the difference between **parsing a URL correctly** and **validating whether the parsed URL is safe for your application's intended operation**.

What additional security checks would you design before allowing your application to use that URL?
