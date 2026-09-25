# Go `int()`, `string()`, and Related Conversions in APIs

## 1. What are `int()`, `string()`, etc.?

In Go, type conversions allow you to convert a value from one data type to another.

Common examples include:

```go
int()
string()
float64()
bool()
```

For example:

```go
var price float64 = 99.50
var amount int = int(price)
```

Here, `99.50` becomes `99`.

### Why are they useful in APIs?

API data often arrives in different types. For example:

- URL parameters are usually strings.
- JSON fields can contain numbers, strings, booleans, etc.
- Database values may need conversion before being returned.
- Query parameters such as `?id=100` arrive as strings.

### Important distinction

There are two different concepts.

**Type conversion:**

```go
age := int(25.5)
```

**String-to-integer parsing:**

```go
age, err := strconv.Atoi("25")
```

`int("25")` is **not valid Go**.

---

# 2. Simple API Example

Suppose we have an API:

```text
GET /users/25
```

The ID comes from the URL as a string.

Using `strconv.Atoi()`:

```go
package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func userHandler(w http.ResponseWriter, r *http.Request) {

	// Example URL:
	// /users/25

	parts := strings.Split(r.URL.Path, "/")

	idString := parts[2]

	// Convert string → int
	id, err := strconv.Atoi(idString)

	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "User ID: %d", id)
}

func main() {
	http.HandleFunc("/users/", userHandler)

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
```

Request:

```text
GET http://localhost:8080/users/25
```

The flow is:

```text
URL
 ↓
"25"  ← string
 ↓
strconv.Atoi()
 ↓
25    ← int
 ↓
Business logic
```

### Other common conversions

```go
// int → string
s := strconv.Itoa(100)

// string → int
n, err := strconv.Atoi("100")

// string → int64
n64, err := strconv.ParseInt("100", 10, 64)

// string → float64
price, err := strconv.ParseFloat("99.50", 64)

// string → bool
active, err := strconv.ParseBool("true")

// int → float64
x := float64(100)

// float64 → int
y := int(99.99)
```

For API development, the `strconv` package is particularly important because API input commonly arrives as strings and must be validated and converted before business logic uses it.

---

# 3. Three Common Mistakes

## Mistake 1: Thinking `int("100")` converts a string to an integer

This is incorrect:

```go
id := int("100") // ❌
```

Use:

```go
id, err := strconv.Atoi("100")
```

because converting a string containing digits requires **parsing**, not a normal Go type conversion.

---

## Mistake 2: Ignoring conversion errors

Beginners sometimes write:

```go
id, _ := strconv.Atoi(input)
```

This throws away the error.

For API requests, that can be dangerous because users can send:

```text
/users/abc
```

instead of:

```text
/users/123
```

Prefer:

```go
id, err := strconv.Atoi(input)

if err != nil {
	// return HTTP 400
}
```

Validate API input before using it.

---

## Mistake 3: Forgetting that numeric conversion can lose information

For example:

```go
price := 99.99

amount := int(price)

fmt.Println(amount)
```

Output:

```text
99
```

The fractional part is discarded.

Therefore, don't blindly convert numeric values just because the target type is convenient. Choose the type based on what your application actually needs.

---

# 4. Real-World API Applications

## Application 1: URL Parameters

Consider:

```text
GET /products/500
```

The product ID extracted from the URL is:

```go
"500"
```

Your service may need:

```go
500
```

So the flow is:

```text
URL parameter
      ↓
string
      ↓
strconv.Atoi()
      ↓
int
      ↓
Database query
```

For example:

```sql
SELECT * FROM products WHERE id = 500;
```

This pattern is extremely common in REST APIs.

---

## Application 2: Query Parameters

Suppose your API supports:

```text
GET /products?page=2&limit=20
```

Query parameters arrive as strings:

```text
page  = "2"
limit = "20"
```

Your Go API can parse them:

```go
page, err := strconv.Atoi(r.URL.Query().Get("page"))
limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
```

Then your application can use:

```text
page = 2
limit = 20
```

for pagination.

A typical flow is:

```text
HTTP Request
     ↓
Query Parameters
     ↓
Strings
     ↓
Parse / Validate
     ↓
Integer values
     ↓
Business Logic
     ↓
Database
```

---

# 5. Exercises

## 🟢 Exercise 1 — Basic Conversion

Create a Go program that receives the following values:

```text
"25"
"100"
"500"
```

Convert them from strings to integers and calculate:

- Sum
- Average
- Maximum value
- Minimum value

Handle invalid input properly.

---

## 🟡 Exercise 2 — Product API

Create a REST API endpoint:

```text
GET /products/{id}
```

For example:

```text
GET /products/101
```

Your API should:

1. Extract the product ID from the URL.
2. Convert the ID from string to integer.
3. Validate that the ID is positive.
4. Return HTTP `400 Bad Request` if the ID is invalid.
5. Return the ID as JSON when the input is valid.

Example expected response:

```json
{
  "product_id": 101
}
```

---

## 🔴 Exercise 3 — Pagination API

Create an API:

```text
GET /products?page=2&limit=10
```

Requirements:

1. Read `page` and `limit`.
2. Convert both values from strings to integers.
3. Validate that both are greater than zero.
4. Apply reasonable maximum limits.
5. Calculate the database offset:

```text
offset = (page - 1) * limit
```

6. Return `page`, `limit`, and `offset` as JSON.
7. Properly handle missing, non-numeric, negative, and excessively large values.

---

# 🤔 Think Deeper

If an API receives:

```text
GET /products/00100
```

and you convert `"00100"` to the integer `100`:

**What information has been lost during the conversion, and can you think of an API scenario where preserving the original string would be more important than converting it to an integer?**
