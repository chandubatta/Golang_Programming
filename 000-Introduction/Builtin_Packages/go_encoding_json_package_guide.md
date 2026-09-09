# Go `encoding/json` Package — Detailed Guide

## 1. What is `encoding/json`?

`encoding/json` is a Go standard-library package used to encode Go values into JSON and decode JSON into Go values.

In simple terms:

```text
Go data
   ↓
json.Marshal()
   ↓
JSON
```

and the reverse:

```text
JSON
   ↓
json.Unmarshal()
   ↓
Go data
```

JSON is commonly used for communication between applications, particularly HTTP APIs, configuration files, message queues, and data exchange between services.

Import it with:

```go
import "encoding/json"
```

There is no external dependency to install.

---

## 2. Encoding vs Decoding

### Encoding

**Go → JSON**

Example:

```go
user := User{
    Name: "Chandu",
    Age: 25,
}
```

becomes:

```json
{
    "Name": "Chandu",
    "Age": 25
}
```

This is called **marshaling**.

### Decoding

**JSON → Go**

Example:

```json
{
    "Name": "Chandu",
    "Age": 25
}
```

becomes:

```go
User{
    Name: "Chandu",
    Age: 25,
}
```

This is called **unmarshaling**.

A useful mental model:

```text
              encoding/json

        ┌──────────────────────┐
        │    Go application    │
        └──────────┬───────────┘
                   │
            Marshal │ Unmarshal
                   │
                   ▼
        ┌──────────────────────┐
        │         JSON         │
        │ {"name":"Chandu"}    │
        └──────────────────────┘
```

---

## 3. Simple Example

```go
package main

import (
    "encoding/json"
    "fmt"
    "log"
)

type User struct {
    Name  string `json:"name"`
    Age   int    `json:"age"`
    Email string `json:"email"`
}

func main() {

    // Go struct
    user := User{
        Name:  "Chandu",
        Age:   25,
        Email: "chandu@example.com",
    }

    // Go -> JSON
    data, err := json.Marshal(user)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(data))

    // JSON -> Go
    var decodedUser User

    err = json.Unmarshal(data, &decodedUser)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("%+v\n", decodedUser)
}
```

Output:

```text
{"name":"Chandu","age":25,"email":"chandu@example.com"}

{Name:Chandu Age:25 Email:chandu@example.com}
```

The two functions you should master first are:

```go
json.Marshal()
json.Unmarshal()
```

---

# 4. All Major Functions in `encoding/json`

The package-level functions include:

```text
Compact
HTMLEscape
Indent
Marshal
MarshalIndent
Unmarshal
Valid
```

There are also methods belonging to `Encoder`, `Decoder`, `Number`, and the package's error types.

---

## 4.1 `json.Marshal()`

### Signature

```go
func Marshal(v any) ([]byte, error)
```

This converts a Go value into JSON.

Example:

```go
user := User{
    Name: "Chandu",
    Age: 25,
}

data, err := json.Marshal(user)
if err != nil {
    log.Fatal(err)
}

fmt.Println(string(data))
```

Result:

```json
{"name":"Chandu","age":25}
```

### Why does it return `[]byte`?

Because JSON is ultimately text represented as bytes.

```go
data, err := json.Marshal(user)
```

`data` is:

```go
[]byte
```

To display it:

```go
fmt.Println(string(data))
```

### What can `Marshal` encode?

It handles many Go types:

```text
bool        → JSON boolean
int         → JSON number
float64     → JSON number
string      → JSON string
struct      → JSON object
slice       → JSON array
array       → JSON array
map         → JSON object
pointer     → pointed-to value
interface   → contained value
```

Some Go values cannot be represented directly in JSON, such as functions, channels, and complex numbers. These result in `UnsupportedTypeError`. Cyclic structures are also unsupported.

---

# 5. JSON Struct Tags

One of the most important concepts in `encoding/json` is the **struct tag**.

```go
type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}
```

Without tags:

```json
{
    "Name": "Chandu",
    "Age": 25
}
```

With tags:

```json
{
    "name": "Chandu",
    "age": 25
}
```

## `omitempty`

```go
type User struct {
    Name  string `json:"name"`
    Age   int    `json:"age,omitempty"`
    Email string `json:"email,omitempty"`
}
```

If:

```go
Age: 0
Email: ""
```

those fields are omitted under the legacy `encoding/json` semantics.

Result:

```json
{
    "name": "Chandu"
}
```

## Ignore a field

```go
Password string `json:"-"`
```

The field won't appear in JSON.

This is useful for preventing sensitive fields from accidentally being serialized.

---

# 6. `json.Unmarshal()`

### Signature

```go
func Unmarshal(data []byte, v any) error
```

This performs the reverse operation:

```text
JSON → Go
```

Example:

```go
data := []byte(`{
    "name": "Chandu",
    "age": 25
}`)

var user User

err := json.Unmarshal(data, &user)
if err != nil {
    log.Fatal(err)
}

fmt.Println(user.Name)
fmt.Println(user.Age)
```

Output:

```text
Chandu
25
```

## Why `&user`?

This is extremely important.

You normally need to pass a pointer:

```go
json.Unmarshal(data, &user)
```

not:

```go
json.Unmarshal(data, user)
```

`Unmarshal` needs to modify `user`.

Conceptually:

```text
JSON
 ↓
Unmarshal()
 ↓
&user
 ↓
user gets populated
```

If the second argument isn't a non-nil pointer, `Unmarshal` returns `InvalidUnmarshalError`.

---

# 7. Unmarshaling into `interface{}` / `any`

You can decode JSON without defining a struct:

```go
var data any

err := json.Unmarshal(
    []byte(`{"name":"Chandu","age":25}`),
    &data,
)
```

Default mappings are approximately:

```text
JSON object   → map[string]any
JSON array    → []any
JSON string   → string
JSON number   → float64
JSON boolean  → bool
JSON null     → nil
```

This is useful when JSON has an unknown or dynamic structure.

However, JSON numbers decoded into `interface{}` normally become `float64`, which can cause precision problems for large integers.

That's where `Decoder.UseNumber()` becomes useful.

---

# 8. `json.MarshalIndent()`

### Signature

```go
func MarshalIndent(v any, prefix, indent string) ([]byte, error)
```

It is similar to:

```go
json.Marshal()
```

but produces human-readable formatted JSON.

Example:

```go
data, err := json.MarshalIndent(user, "", "    ")
if err != nil {
    log.Fatal(err)
}

fmt.Println(string(data))
```

Output:

```json
{
    "name": "Chandu",
    "age": 25,
    "email": "chandu@example.com"
}
```

Useful for:

- configuration files
- debugging
- logs
- pretty API output
- generated JSON files

---

# 9. `json.Indent()`

### Signature

```go
func Indent(
    dst *bytes.Buffer,
    src []byte,
    prefix string,
    indent string,
) error
```

Unlike `MarshalIndent`, it doesn't take a Go value. It takes **existing JSON** and formats it.

Example:

```go
input := []byte(`{"name":"Chandu","age":25}`)

var output bytes.Buffer

err := json.Indent(&output, input, "", "    ")
if err != nil {
    log.Fatal(err)
}

fmt.Println(output.String())
```

Result:

```json
{
    "name": "Chandu",
    "age": 25
}
```

Think:

```text
MarshalIndent
Go → pretty JSON

Indent
JSON → pretty JSON
```

---

# 10. `json.Compact()`

### Signature

```go
func Compact(dst *bytes.Buffer, src []byte) error
```

It removes unnecessary whitespace from JSON.

Input:

```json
{
    "name": "Chandu",
    "age": 25
}
```

Output:

```json
{"name":"Chandu","age":25}
```

Example:

```go
input := []byte(`
{
    "name": "Chandu",
    "age": 25
}
`)

var output bytes.Buffer

err := json.Compact(&output, input)
if err != nil {
    log.Fatal(err)
}

fmt.Println(output.String())
```

Useful when you want smaller JSON payloads.

---

# 11. `json.HTMLEscape()`

### Signature

```go
func HTMLEscape(dst *bytes.Buffer, src []byte)
```

It escapes characters that can be problematic when JSON is embedded inside HTML `<script>` elements.

Characters such as:

```text
<
>
&
U+2028
U+2029
```

can be escaped.

Example:

```go
var output bytes.Buffer

input := []byte(`{"message":"<script>alert('hello')</script>"}`)

json.HTMLEscape(&output, input)

fmt.Println(output.String())
```

You may see:

```json
{"message":"\u003cscript\u003ealert('hello')\u003c/script\u003e"}
```

For normal REST APIs, you usually won't call this manually.

---

# 12. `json.Valid()`

### Signature

```go
func Valid(data []byte) bool
```

It checks whether a byte slice contains syntactically valid JSON.

Example:

```go
valid := json.Valid([]byte(`{"name":"Chandu"}`))

fmt.Println(valid)
```

Output:

```text
true
```

Invalid:

```go
valid := json.Valid([]byte(`{"name":}`))

fmt.Println(valid)
```

Output:

```text
false
```

## Important

`json.Valid()` checks **JSON syntax**, not whether the JSON is semantically correct for your application.

For example:

```json
{
    "age": "hello"
}
```

is valid JSON.

But if your Go struct expects:

```go
Age int
```

you'll encounter a type error when decoding.

---

# 13. `Decoder`

`json.Decoder` is a more advanced part of `encoding/json`.

A `Decoder` reads JSON from an `io.Reader`.

Examples of readers include:

```text
file
HTTP request body
network connection
buffer
string reader
```

The important difference is:

```text
json.Unmarshal
      ↓
works with []byte

json.Decoder
      ↓
works with io.Reader
```

The `Decoder` is designed for reading and decoding JSON values from an input stream.

---

# 14. `json.NewDecoder()`

### Signature

```go
func NewDecoder(r io.Reader) *Decoder
```

Example:

```go
decoder := json.NewDecoder(reader)
```

A very common HTTP-server example:

```go
func createUser(w http.ResponseWriter, r *http.Request) {

    var user User

    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    fmt.Println(user)
}
```

`NewDecoder` creates a decoder reading from an `io.Reader`.

---

# 15. `Decoder.Decode()`

### Signature

```go
func (dec *Decoder) Decode(v any) error
```

It reads the **next JSON value** from the stream and stores it in `v`.

Example:

```go
decoder := json.NewDecoder(strings.NewReader(`
    {"name":"Chandu"}
`))

var user User

err := decoder.Decode(&user)
if err != nil {
    log.Fatal(err)
}
```

Difference:

```go
json.Unmarshal(data, &user)
```

requires JSON data as `[]byte`.

Whereas:

```go
decoder.Decode(&user)
```

reads from a stream.

---

# 16. `Decoder.DisallowUnknownFields()`

### Signature

```go
func (dec *Decoder) DisallowUnknownFields()
```

By default, if your struct is:

```go
type User struct {
    Name string `json:"name"`
}
```

and incoming JSON is:

```json
{
    "name": "Chandu",
    "age": 25
}
```

the unknown `"age"` field is normally ignored.

But:

```go
decoder := json.NewDecoder(r.Body)
decoder.DisallowUnknownFields()
```

causes decoding to fail when the JSON contains an object field that doesn't correspond to an exported destination struct field.

This is useful when you want strict API request validation.

---

# 17. `Decoder.UseNumber()`

### Signature

```go
func (dec *Decoder) UseNumber()
```

Normally:

```go
var data any

decoder.Decode(&data)
```

causes JSON numbers to become:

```go
float64
```

Calling:

```go
decoder.UseNumber()
```

makes numbers become:

```go
json.Number
```

instead.

Example:

```go
decoder := json.NewDecoder(reader)

decoder.UseNumber()

var data any

err := decoder.Decode(&data)
```

This is particularly useful when preserving the textual representation of JSON numbers matters.

---

# 18. `Decoder.More()`

### Signature

```go
func (dec *Decoder) More() bool
```

This is mainly used when processing arrays or objects token-by-token.

Example:

```go
decoder := json.NewDecoder(reader)

token, _ := decoder.Token()

for decoder.More() {
    var user User
    decoder.Decode(&user)

    fmt.Println(user)
}

token, _ = decoder.Token()
```

Conceptually:

```text
[
    User 1   ← More()
    User 2   ← More()
    User 3   ← More()
]
         ↑
       stop
```

It tells you whether there are more elements in the current array/object context.

---

# 19. `Decoder.Token()`

### Signature

```go
func (dec *Decoder) Token() (json.Token, error)
```

This is a lower-level JSON parsing API.

Instead of decoding an entire struct:

```go
decoder.Decode(&user)
```

you can inspect JSON token-by-token.

For:

```json
{
    "name": "Chandu",
    "age": 25
}
```

tokens can include:

```text
{
"name"
"Chandu"
"age"
25
}
```

The `Token` type can represent delimiters, booleans, numbers, strings, or `nil`.

This becomes useful when processing very large or dynamically structured JSON.

---

# 20. `Decoder.Buffered()`

### Signature

```go
func (dec *Decoder) Buffered() io.Reader
```

This is an advanced function.

A decoder can read more bytes from the underlying reader than were needed for the current JSON value.

`Buffered()` gives you access to data that has already been read into the decoder's internal buffer but hasn't yet been consumed.

Conceptually:

```text
Underlying Reader
       ↓
   Decoder
       ↓
 internal buffer
       ↓
 Buffered()
```

It is mainly useful when combining JSON decoding with another protocol or when you need precise control over remaining input.

---

# 21. `Decoder.InputOffset()`

### Signature

```go
func (dec *Decoder) InputOffset() int64
```

It reports the approximate byte position in the input where the decoder currently is.

This can be useful for diagnostics when processing large JSON streams.

Conceptually:

```text
JSON:
012345678901234567890...
          ↑
       offset
```

---

# 22. `Encoder`

`Encoder` is essentially the streaming counterpart of `Marshal`.

Think:

```text
Marshal
Go → []byte

Encoder
Go → io.Writer
```

`Encoder` writes JSON values to an output stream.

---

# 23. `json.NewEncoder()`

### Signature

```go
func NewEncoder(w io.Writer) *Encoder
```

Example:

```go
encoder := json.NewEncoder(w)
```

Common HTTP usage:

```go
func getUser(w http.ResponseWriter, r *http.Request) {

    user := User{
        Name: "Chandu",
        Age: 25,
    }

    w.Header().Set("Content-Type", "application/json")

    json.NewEncoder(w).Encode(user)
}
```

This is extremely common in Go HTTP APIs.

---

# 24. `Encoder.Encode()`

### Signature

```go
func (enc *Encoder) Encode(v any) error
```

It writes JSON to the encoder's output stream.

```go
err := json.NewEncoder(w).Encode(user)
```

Difference from `Marshal`:

```go
data, _ := json.Marshal(user)
```

gives you:

```text
[]byte
```

while:

```go
json.NewEncoder(w).Encode(user)
```

writes directly to:

```text
w
```

`Encode` appends a newline after the JSON value.

---

# 25. `Encoder.SetEscapeHTML()`

### Signature

```go
func (enc *Encoder) SetEscapeHTML(on bool)
```

Default:

```go
true
```

You can disable HTML escaping:

```go
encoder.SetEscapeHTML(false)
```

This is useful when you aren't embedding JSON inside HTML and want output to be more readable.

---

# 26. `Encoder.SetIndent()`

### Signature

```go
func (enc *Encoder) SetIndent(prefix, indent string)
```

It tells the encoder to pretty-print subsequent JSON.

Example:

```go
encoder := json.NewEncoder(os.Stdout)

encoder.SetIndent("", "    ")

encoder.Encode(user)
```

Output:

```json
{
    "name": "Chandu",
    "age": 25
}
```

Calling:

```go
encoder.SetIndent("", "")
```

turns indentation off.

---

# 27. `json.Number`

`json.Number` is:

```go
type Number string
```

It represents a JSON number literal.

Example:

```go
n := json.Number("123456789.123456")
```

Unlike immediately converting everything to `float64`, it retains the number's textual representation.

---

## `Number.String()`

```go
func (n Number) String() string
```

Returns the original number text.

Example:

```go
n := json.Number("123.45")

fmt.Println(n.String())
```

Output:

```text
123.45
```

---

## `Number.Int64()`

```go
func (n Number) Int64() (int64, error)
```

Converts the JSON number to an `int64`.

```go
n := json.Number("12345")

value, err := n.Int64()
```

If the number isn't a valid integer or is outside the `int64` range, you'll get an error.

---

## `Number.Float64()`

```go
func (n Number) Float64() (float64, error)
```

Converts the number to a `float64`.

---

## Current-version note about additional Number methods

Current Go documentation also exposes methods related to the newer JSON implementation, such as `MarshalJSONTo` and `UnmarshalJSONFrom`. These are generally **not methods you need to call directly in ordinary `encoding/json` application code**.

For learning `encoding/json`, concentrate on:

```text
String()
Int64()
Float64()
```

first.

---

# 28. `json.Delim`

`json.Delim` represents JSON structural delimiters:

```text
[
]
{
}
```

For example:

```go
json.Delim('{')
json.Delim('[')
```

It is primarily used with:

```go
Decoder.Token()
```

Its method is:

```go
func (d Delim) String() string
```

which returns the delimiter as a string.

---

# 29. `json.RawMessage`

This is an **important advanced feature**.

```go
type RawMessage
```

It represents raw encoded JSON.

It allows you to **delay decoding part of a JSON document** or use precomputed JSON.

Example:

```go
type Event struct {
    Type string          `json:"type"`
    Data json.RawMessage `json:"data"`
}
```

Suppose:

```json
{
    "type": "user_created",
    "data": {
        "name": "Chandu",
        "age": 25
    }
}
```

You can first inspect:

```go
event.Type
```

and only then decide what type `event.Data` should be decoded into.

This is useful for:

- event-driven systems
- polymorphic JSON
- webhook processing
- message brokers
- APIs with different payload structures

---

# 30. Custom `MarshalJSON`

You can control how your own type becomes JSON by implementing:

```go
MarshalJSON() ([]byte, error)
```

through the `json.Marshaler` interface.

Example:

```go
type User struct {
    Name string
}

func (u User) MarshalJSON() ([]byte, error) {
    return json.Marshal(map[string]string{
        "username": u.Name,
    })
}
```

Now:

```go
user := User{Name: "Chandu"}

data, _ := json.Marshal(user)
```

produces:

```json
{
    "username": "Chandu"
}
```

The package recognizes types implementing `Marshaler` and calls their custom `MarshalJSON` method.

---

# 31. Custom `UnmarshalJSON`

Similarly, you can implement:

```go
UnmarshalJSON([]byte) error
```

through the `json.Unmarshaler` interface.

Example:

```go
type User struct {
    Name string
}

func (u *User) UnmarshalJSON(data []byte) error {
    // custom decoding
    return nil
}
```

Then:

```go
json.Unmarshal(data, &user)
```

will invoke your custom method.

This is useful when external JSON doesn't match your Go data model directly.

---

# 32. Error Types

The package provides several useful error types.

## `SyntaxError`

Indicates invalid JSON syntax.

Example:

```json
{"name":}
```

This is syntactically invalid.

`SyntaxError` contains an offset indicating where the problem occurred.

---

## `UnmarshalTypeError`

This occurs when JSON has a type that doesn't fit the destination Go type.

Example:

```go
type User struct {
    Age int `json:"age"`
}
```

JSON:

```json
{
    "age": "twenty five"
}
```

The JSON itself is valid, but the string cannot be decoded into `int`.

---

## `InvalidUnmarshalError`

Usually happens when calling:

```go
json.Unmarshal(data, user)
```

instead of:

```go
json.Unmarshal(data, &user)
```

The destination must be a non-nil pointer.

---

## `UnsupportedTypeError`

Returned when trying to marshal a Go type that JSON can't represent, such as:

```text
func
chan
complex
```

---

## `UnsupportedValueError`

Used when the value itself can't be represented properly, such as:

```go
math.NaN()
math.Inf(1)
```

when marshaling.

---

## `MarshalerError`

Wraps an error returned by a custom:

```go
MarshalJSON()
```

or text-marshaling method.

It also provides:

```go
Unwrap()
```

so you can access the underlying error.

---

# 33. Complete Function and Type Map

| Function / Method | Purpose |
|---|---|
| `json.Marshal` | Go → JSON bytes |
| `json.Unmarshal` | JSON bytes → Go |
| `json.MarshalIndent` | Go → pretty JSON |
| `json.Indent` | Format existing JSON |
| `json.Compact` | Remove unnecessary JSON whitespace |
| `json.HTMLEscape` | HTML-safe JSON escaping |
| `json.Valid` | Check JSON syntax |
| `json.NewDecoder` | Create streaming JSON decoder |
| `Decoder.Decode` | Read next JSON value |
| `Decoder.More` | Check for more array/object elements |
| `Decoder.Token` | Read JSON token |
| `Decoder.UseNumber` | Preserve numbers as `json.Number` |
| `Decoder.DisallowUnknownFields` | Reject unknown struct fields |
| `Decoder.Buffered` | Access decoder's unread buffer |
| `Decoder.InputOffset` | Get current input offset |
| `json.NewEncoder` | Create streaming JSON encoder |
| `Encoder.Encode` | Write JSON to `io.Writer` |
| `Encoder.SetIndent` | Pretty-print encoder output |
| `Encoder.SetEscapeHTML` | Control HTML escaping |
| `Number.String` | Get number text |
| `Number.Int64` | Convert JSON number to `int64` |
| `Number.Float64` | Convert JSON number to `float64` |
| `Delim.String` | Get delimiter as string |
| `Marshaler` | Customize Go → JSON |
| `Unmarshaler` | Customize JSON → Go |
| `RawMessage` | Delay/preserve JSON |
| `SyntaxError` | Invalid JSON syntax |
| `UnmarshalTypeError` | JSON/Go type mismatch |
| `InvalidUnmarshalError` | Invalid destination passed to `Unmarshal` |
| `UnsupportedTypeError` | Unsupported Go type |
| `UnsupportedValueError` | Unsupported value |
| `MarshalerError` | Error from custom marshaling |

---

# 34. Three Common Beginner Mistakes

## Mistake 1 — Forgetting the pointer in `Unmarshal`

Wrong:

```go
json.Unmarshal(data, user)
```

Correct:

```go
json.Unmarshal(data, &user)
```

Remember:

```text
Marshal:
Go → JSON

Unmarshal:
JSON → modifies Go value
       ↑
      &user
```

---

## Mistake 2 — Assuming JSON field names automatically match your API contract

You might write:

```go
type User struct {
    FirstName string
}
```

and expect:

```json
{
    "first_name": "Chandu"
}
```

But Go's JSON behavior doesn't automatically convert `FirstName` to `first_name`.

Use:

```go
type User struct {
    FirstName string `json:"first_name"`
}
```

Struct tags are essential when integrating with APIs whose naming conventions differ from Go's field names.

---

## Mistake 3 — Assuming `json.Unmarshal` validates business rules

Suppose:

```go
type User struct {
    Age int `json:"age"`
}
```

This JSON is valid:

```json
{
    "age": -500
}
```

`encoding/json` can successfully decode it.

But perhaps your application says:

```text
Age must be between 0 and 150
```

That's **business validation**, not JSON decoding.

So your pipeline should usually be:

```text
HTTP Request
     ↓
JSON decoding
     ↓
Go struct
     ↓
Validation
     ↓
Business logic
     ↓
Database
```

Don't confuse JSON parsing with application validation.

---

# 35. Two Real-World Applications

## Application 1 — REST API

Imagine your Go backend receives:

```http
POST /users
Content-Type: application/json
```

Body:

```json
{
    "name": "Chandu",
    "email": "chandu@example.com"
}
```

Your handler can use:

```go
func createUser(w http.ResponseWriter, r *http.Request) {

    var user User

    err := json.NewDecoder(r.Body).Decode(&user)

    if err != nil {
        http.Error(
            w,
            "Invalid JSON",
            http.StatusBadRequest,
        )
        return
    }

    // validate user
    // save user
    // return response
}
```

And for the response:

```go
json.NewEncoder(w).Encode(user)
```

This Go pattern is extremely common in HTTP services.

---

# 36. Application 2 — Event-Driven / Microservice Systems

Suppose your system publishes events:

```json
{
    "event": "payment_completed",
    "data": {
        "payment_id": "P1001",
        "amount": 5000
    }
}
```

Another service receives it.

You might initially decode:

```go
type Event struct {
    Event string          `json:"event"`
    Data  json.RawMessage `json:"data"`
}
```

Then:

```text
Receive event
      ↓
Decode Event
      ↓
Inspect Event field
      ↓
Choose appropriate Go type
      ↓
Decode RawMessage
      ↓
Process event
```

`RawMessage` is particularly useful for this kind of polymorphic payload.

---

# 37. Three Progressive Exercises

## Exercise 1 — Beginner: User JSON

Create a Go program with:

```go
type User struct {
    Name     string
    Age      int
    Email    string
}
```

Your program should:

1. Create a `User`.
2. Convert the `User` into JSON.
3. Print the JSON.
4. Create a JSON string representing another user.
5. Decode that JSON back into a `User`.
6. Print the resulting Go struct.

**Goal:** Practice `Marshal` and `Unmarshal`.

---

## Exercise 2 — Intermediate: REST API Request

Create an HTTP server with:

```text
POST /users
```

The request body should look like:

```json
{
    "name": "Chandu",
    "age": 25,
    "email": "chandu@example.com"
}
```

Your server should:

1. Decode the request body using `json.Decoder`.
2. Reject malformed JSON.
3. Reject unknown fields.
4. Validate that `name` and `email` aren't empty.
5. Return an appropriate JSON response.
6. Set the correct `Content-Type`.
7. Use `json.Encoder` to generate the response.

**Goal:** Practice `NewDecoder`, `Decode`, `DisallowUnknownFields`, `NewEncoder`, and `Encode`.

---

## Exercise 3 — Advanced: Polymorphic Event Processor

Build an event-processing program that receives JSON events with this general structure:

```json
{
    "type": "user_created",
    "data": {}
}
```

Support at least three event types:

```text
user_created
order_created
payment_completed
```

Each event must have a different `data` structure.

Your program should:

1. Decode the outer event.
2. Keep `data` as `json.RawMessage`.
3. Inspect the event type.
4. Select the appropriate Go struct.
5. Decode `data` into that struct.
6. Reject unsupported event types.
7. Properly handle malformed JSON.
8. Demonstrate how your design could process a stream containing multiple events.
9. Consider what happens when a new event type is added in the future.

**Goal:** Practice `RawMessage`, `Decoder`, custom structures, error handling, and designing an extensible JSON-processing system.

---

# 38. One Important Modern Go Point

Since you're learning this package now, there is one important thing to know.

Current Go documentation distinguishes:

```text
encoding/json
      ↓
v1 / legacy behavior

encoding/json/v2
      ↓
newer JSON implementation
```

The older `encoding/json` package remains supported, but Go's documentation recommends `encoding/json/v2` for **new usage**, because v2 has cleaner and safer defaults.

For example, there are differences involving:

- case-sensitive field matching
- duplicate JSON object names
- `omitempty`
- nil slices/maps
- invalid UTF-8
- map ordering
- array length handling
- `time.Duration`
- error reporting

The existing `encoding/json` package preserves its historical behavior for compatibility.

For a learning path:

```text
First:
encoding/json fundamentals
        ↓
Marshal / Unmarshal
        ↓
Struct tags
        ↓
Encoder / Decoder
        ↓
RawMessage
        ↓
Custom Marshaler / Unmarshaler
        ↓
Error handling
        ↓
Then:
encoding/json/v2
```

This gives you both real-world compatibility knowledge and an understanding of where modern Go JSON handling is heading.

---

# 39. Thought-Provoking Question

Imagine you're building a Go payment API that receives JSON from **10,000 different clients**, and some clients send unexpected fields, duplicate JSON keys, very large numbers, or slightly different capitalization of field names.

**Would you prefer your JSON decoder to be forgiving and "do its best," or strict and reject anything unexpected—and what could go wrong in a payment system if you choose the wrong behavior?**

This gets to a deeper issue behind `encoding/json`:

> **JSON parsing isn't just about converting data—it is also about defining the trust boundary between your application and the outside world.**
