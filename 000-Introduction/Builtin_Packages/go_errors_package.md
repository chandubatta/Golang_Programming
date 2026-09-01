# Go `errors` Package

The Go standard-library `errors` package provides functions for **creating, comparing, inspecting, unwrapping, and combining errors**. It is one of the most important packages to understand because Go commonly handles failures by returning an `error` value rather than throwing exceptions.

> **Current Go note:** In Go 1.26, the standard `errors` package contains **6 exported functions**: `New`, `As`, `AsType`, `Is`, `Join`, and `Unwrap`. `AsType` was added in Go 1.26.

---

## 1. What is the `errors` package?

Import it with:

```go
import "errors"
```

The package is mainly used for:

- Creating errors
- Defining reusable/sentinel errors
- Comparing errors
- Inspecting the underlying type of an error
- Extracting wrapped errors
- Combining multiple errors
- Building error chains/trees

A typical Go function returns an error like this:

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }

    return a / b, nil
}
```

The caller then checks:

```go
result, err := divide(10, 0)

if err != nil {
    fmt.Println("Error:", err)
    return
}

fmt.Println(result)
```

The basic Go pattern is:

```text
function()
    ↓
value, error
    ↓
if err != nil
    ↓
handle or return error
```

---

# 2. Functions in the `errors` Package

Let's examine **every exported function** in the current standard-library `errors` package.

There are six:

| Function | Purpose |
|---|---|
| `errors.New()` | Create a new error |
| `errors.Is()` | Check whether an error matches another error |
| `errors.As()` | Find an error of a particular type |
| `errors.AsType()` | Find an error of a particular type using generics |
| `errors.Unwrap()` | Retrieve the directly wrapped error |
| `errors.Join()` | Combine multiple errors |

---

# 3. `errors.New()`

## Purpose

`errors.New()` creates a new error containing a text message.

Syntax:

```go
errors.New(text)
```

Example:

```go
package main

import (
    "errors"
    "fmt"
)

func main() {
    err := errors.New("file not found")

    fmt.Println(err)
}
```

Output:

```text
file not found
```

The returned value implements Go's `error` interface.

### Important point

Every call to `errors.New()` creates a **distinct error value**, even if the messages are identical.

For example:

```go
err1 := errors.New("file not found")
err2 := errors.New("file not found")

fmt.Println(err1 == err2)
```

Output:

```text
false
```

This is important when learning about `errors.Is()`.

---

## Common use

`errors.New()` is commonly used for:

- Validation errors
- Sentinel errors
- Simple application errors
- Returning an error from a function

Example:

```go
var ErrInvalidAge = errors.New("invalid age")
```

Then:

```go
func validateAge(age int) error {
    if age < 0 {
        return ErrInvalidAge
    }

    return nil
}
```

---

# 4. `errors.Is()`

## Purpose

`errors.Is()` determines whether an error **matches a target error**.

Syntax:

```go
errors.Is(err, target)
```

It returns:

```text
true
```

if the error matches the target somewhere in its error chain/tree.

Otherwise:

```text
false
```

Go recommends `errors.Is()` instead of directly comparing errors when errors may have been wrapped.

---

## Example

```go
package main

import (
    "errors"
    "fmt"
)

var ErrNotFound = errors.New("not found")

func findUser() error {
    return ErrNotFound
}

func main() {
    err := findUser()

    if errors.Is(err, ErrNotFound) {
        fmt.Println("User was not found")
    }
}
```

Output:

```text
User was not found
```

---

## Why not use `==`?

You might write:

```go
if err == ErrNotFound {
    // ...
}
```

Sometimes that works.

But suppose the error gets wrapped:

```go
err := fmt.Errorf("database lookup failed: %w", ErrNotFound)
```

Now:

```go
err == ErrNotFound
```

is false.

But:

```go
errors.Is(err, ErrNotFound)
```

is true.

This is one of the most important concepts in modern Go error handling.

---

# 5. `errors.As()`

## Purpose

`errors.As()` searches an error chain/tree for an error of a particular **type**.

Syntax:

```go
errors.As(err, &target)
```

It returns:

```text
true
```

if it finds a matching error.

It also places the matching error into `target`.

---

## Example

Suppose we have a custom error:

```go
type ValidationError struct {
    Field   string
    Message string
}

func (e *ValidationError) Error() string {
    return e.Field + ": " + e.Message
}
```

We can create it:

```go
err := &ValidationError{
    Field:   "email",
    Message: "invalid email address",
}
```

Now use `errors.As()`:

```go
var validationErr *ValidationError

if errors.As(err, &validationErr) {
    fmt.Println("Field:", validationErr.Field)
    fmt.Println("Message:", validationErr.Message)
}
```

Output:

```text
Field: email
Message: invalid email address
```

---

## `Is()` vs `As()`

This distinction is extremely important.

### `Is()`

Asks:

> "Is this error a particular error?"

```go
errors.Is(err, ErrNotFound)
```

### `As()`

Asks:

> "Does this error contain an error of this particular type?"

```go
var validationErr *ValidationError

errors.As(err, &validationErr)
```

Think:

```text
Is → identity/category
As → type/details
```

---

# 6. `errors.AsType()`

## Purpose

`errors.AsType()` is the newer generic alternative to `errors.As()`.

It was added in **Go 1.26**.

Syntax:

```go
errors.AsType[ErrorType](err)
```

It returns two values:

```text
matchingError, true
```

or:

```text
zeroValue, false
```

---

## Example

Using our `ValidationError`:

```go
package main

import (
    "errors"
    "fmt"
)

type ValidationError struct {
    Field   string
    Message string
}

func (e *ValidationError) Error() string {
    return e.Field + ": " + e.Message
}

func main() {
    err := &ValidationError{
        Field:   "email",
        Message: "invalid email",
    }

    validationErr, ok := errors.AsType[*ValidationError](err)

    if ok {
        fmt.Println("Field:", validationErr.Field)
        fmt.Println("Message:", validationErr.Message)
    }
}
```

Output:

```text
Field: email
Message: invalid email
```

---

## Why was `AsType()` introduced?

Traditional `As()` requires a target variable:

```go
var validationErr *ValidationError

if errors.As(err, &validationErr) {
    // use validationErr
}
```

`AsType()` lets you write:

```go
validationErr, ok := errors.AsType[*ValidationError](err)

if ok {
    // use validationErr
}
```

This can be easier to read and is particularly useful when working with generics.

---

# 7. `errors.Unwrap()`

## Purpose

`errors.Unwrap()` retrieves the error directly wrapped by another error.

Syntax:

```go
errors.Unwrap(err)
```

If the error doesn't wrap another error, it returns:

```text
nil
```

The standard `errors.Unwrap()` specifically handles an `Unwrap() error` method; it does **not** unwrap an `Unwrap() []error` returned by `errors.Join()`.

---

## Example

```go
package main

import (
    "errors"
    "fmt"
)

func main() {
    original := errors.New("database connection failed")

    wrapped := fmt.Errorf("service startup failed: %w", original)

    unwrapped := errors.Unwrap(wrapped)

    fmt.Println(unwrapped)
}
```

Output:

```text
database connection failed
```

Conceptually:

```text
wrapped error
     |
     | Unwrap()
     ↓
original error
```

---

## Important

`errors.Unwrap()` unwraps **one level**.

If you have:

```text
Error A
   ↓
Error B
   ↓
Error C
```

one call:

```go
errors.Unwrap(A)
```

gets:

```text
B
```

Another call gets:

```text
C
```

However, in normal application code, you generally don't manually repeatedly call `Unwrap()`. `errors.Is()` and `errors.As()` are usually more useful for inspecting an error chain.

---

# 8. `errors.Join()`

## Purpose

`errors.Join()` combines multiple errors into a single error.

It was introduced in **Go 1.20**.

Syntax:

```go
errors.Join(err1, err2, err3)
```

Example:

```go
package main

import (
    "errors"
    "fmt"
)

func main() {
    err1 := errors.New("username is required")
    err2 := errors.New("email is required")

    err := errors.Join(err1, err2)

    fmt.Println(err)
}
```

Output:

```text
username is required
email is required
```

The resulting error wraps multiple errors through `Unwrap() []error`.

---

## `errors.Is()` works with `Join()`

For example:

```go
var ErrUsername = errors.New("username is required")
var ErrEmail = errors.New("email is required")

err := errors.Join(ErrUsername, ErrEmail)

fmt.Println(errors.Is(err, ErrUsername))
fmt.Println(errors.Is(err, ErrEmail))
```

Output:

```text
true
true
```

This makes `Join()` particularly useful when multiple independent operations can fail.

---

# 9. Error Wrapping

Although there isn't an `errors.Wrap()` function in the standard package, **error wrapping is a major part of Go error handling**.

The usual approach is:

```go
fmt.Errorf("context: %w", err)
```

Example:

```go
err := errors.New("file not found")

wrappedErr := fmt.Errorf("loading configuration failed: %w", err)
```

Now you have:

```text
loading configuration failed
             |
             ↓
       file not found
```

You can still detect the original error:

```go
if errors.Is(wrappedErr, err) {
    fmt.Println("The original error was file not found")
}
```

Go's error model allows wrapped errors to preserve the underlying error while adding useful context.

---

# 10. The `error` Interface

To understand the `errors` package properly, you should also understand the built-in `error` interface:

```go
type error interface {
    Error() string
}
```

Any type implementing:

```go
Error() string
```

can be used as an error.

Example:

```go
type MyError struct {
    Message string
}

func (e MyError) Error() string {
    return e.Message
}
```

Now:

```go
err := MyError{
    Message: "something went wrong",
}
```

is an `error`.

This is how Go allows applications to create **custom error types**.

---

# 11. Complete Simple Example

Here's a small program combining several concepts:

```go
package main

import (
    "errors"
    "fmt"
)

var ErrUserNotFound = errors.New("user not found")

func findUser(id int) error {
    if id == 0 {
        return ErrUserNotFound
    }

    return nil
}

func getUser(id int) error {
    err := findUser(id)

    if err != nil {
        return fmt.Errorf("getUser failed: %w", err)
    }

    return nil
}

func main() {
    err := getUser(0)

    if err != nil {
        fmt.Println("Error:", err)

        if errors.Is(err, ErrUserNotFound) {
            fmt.Println("Handle missing user")
        }
    }
}
```

Output:

```text
Error: getUser failed: user not found
Handle missing user
```

Notice something important:

```text
getUser failed: user not found
       |
       ↓
ErrUserNotFound
```

The additional context is preserved, while `errors.Is()` can still identify the original error.

---

# 12. Three Common Beginner Mistakes

## Mistake 1: Comparing wrapped errors with `==`

Beginners often write:

```go
if err == ErrNotFound {
    // ...
}
```

This can fail when the error has been wrapped.

### Better:

```go
if errors.Is(err, ErrNotFound) {
    // ...
}
```

### Remember:

```text
==         → direct value comparison
errors.Is  → error-chain-aware comparison
```

---

## Mistake 2: Using `errors.As()` when they really need `errors.Is()`

For example, if you have:

```go
var ErrUnauthorized = errors.New("unauthorized")
```

you don't need:

```go
errors.As(...)
```

You need:

```go
errors.Is(err, ErrUnauthorized)
```

Use:

```text
Is → identify a particular error
As → retrieve a particular error type
```

---

## Mistake 3: Losing the original error while adding context

A beginner might do:

```go
return errors.New("database operation failed")
```

when the original error was:

```go
return fmt.Errorf("database operation failed: %w", err)
```

The first version loses the original error.

The second preserves it.

This matters because callers can then use:

```go
errors.Is(...)
```

or:

```go
errors.As(...)
```

to inspect the underlying failure.

---

# 13. Real-World Application #1: Database Applications

Imagine an application accessing a database:

```text
HTTP request
     ↓
Service
     ↓
Repository
     ↓
Database
```

The database might return:

```text
record not found
```

The repository can add context:

```go
return fmt.Errorf("finding user %d: %w", id, err)
```

The service can add more:

```go
return fmt.Errorf("loading user profile: %w", err)
```

The final error could look like:

```text
loading user profile:
finding user 42:
record not found
```

Yet the application can still do:

```go
if errors.Is(err, ErrNotFound) {
    // return HTTP 404
}
```

This gives you **human-readable context without losing machine-readable error information**.

---

# 14. Real-World Application #2: Validation of Multiple Fields

Suppose an API receives:

```json
{
    "username": "",
    "email": "",
    "age": -5
}
```

Instead of returning only the first validation failure, you could collect several errors:

```go
var errs []error

if username == "" {
    errs = append(errs, errors.New("username is required"))
}

if email == "" {
    errs = append(errs, errors.New("email is required"))
}

if age < 0 {
    errs = append(errs, errors.New("age cannot be negative"))
}

return errors.Join(errs...)
```

The caller can then receive all relevant validation problems at once.

This is a particularly good use case for `errors.Join()`.

---

# 15. Exercises

## Exercise 1 — Beginner: Basic Error Creation

Create a function:

```go
func checkAge(age int) error
```

Requirements:

- If `age` is less than `18`, return an error saying `"user must be at least 18"`.
- If `age` is 18 or greater, return `nil`.
- Use the `errors` package.
- In `main()`, call the function with several different ages.
- Print an appropriate message depending on whether an error occurred.

**Do not use `panic`.**

---

## Exercise 2 — Intermediate: Sentinel Error + Wrapping

Create a program for finding a user.

Requirements:

1. Define a package-level error:

```go
ErrUserNotFound
```

2. Create:

```go
func findUser(id int) error
```

3. If the ID doesn't exist, return `ErrUserNotFound`.

4. Create another function that calls `findUser()` and adds contextual information using error wrapping.

5. In `main()`, use `errors.Is()` to determine whether the final error originated from `ErrUserNotFound`.

6. Make sure your error message contains useful context while preserving the original error.

---

## Exercise 3 — Advanced: Custom Errors + `Join()` + `AsType()`

Build a user-registration validation system.

Create a custom error type:

```go
ValidationError
```

The error should contain information about:

- Field name
- Validation message

Your program should:

1. Validate username.
2. Validate email.
3. Validate age.
4. Create a separate `ValidationError` for every invalid field.
5. Collect all validation errors.
6. Combine them using `errors.Join()`.
7. Return the combined error.
8. In `main()`, determine whether a `ValidationError` exists using `errors.AsType()`.
9. Extract information from the matching custom error.
10. Also experiment with `errors.Is()` and explain through your own observations why `Is()` and `AsType()` serve different purposes.

**Do not provide yourself with a shortcut by simply printing the combined error. Try to inspect the error structure programmatically.**

---

# 16. Quick Function Cheat Sheet

| Function | What it does | Typical use |
|---|---|---|
| `errors.New()` | Creates an error | Simple/custom sentinel errors |
| `errors.Is()` | Checks for a matching error | Detect specific failures |
| `errors.As()` | Finds an error of a specific type | Extract custom error information |
| `errors.AsType()` | Generic version of type inspection | Modern Go 1.26+ code |
| `errors.Unwrap()` | Gets the directly wrapped error | Manually inspect one wrapping level |
| `errors.Join()` | Combines multiple errors | Multiple simultaneous failures |

### The most important mental model

Think about Go errors like this:

```text
                 error
                   │
          ┌────────┴────────┐
          ↓                 ↓
      Is this?           What type?
          │                 │
      errors.Is()       errors.As()
                           /
                    errors.AsType()
          │
          ↓
     Wrapped errors
          │
      errors.Unwrap()
          │
          ↓
     Multiple errors
          │
      errors.Join()
```

And remember:

> **`errors.New()` creates errors, `fmt.Errorf("%w")` adds context, `errors.Is()` identifies errors, `errors.As()` / `AsType()` identify error types, `errors.Unwrap()` accesses the underlying error, and `errors.Join()` combines multiple errors.**

---

# 17. Thought-Provoking Question

Imagine a large application where an error travels through **five layers**:

```text
Database
   ↓
Repository
   ↓
Service
   ↓
API Handler
   ↓
HTTP Response
```

**How would you design your error types and wrapping strategy so that developers get enough context to debug the problem, while the API layer can still reliably distinguish between errors such as “not found,” “invalid input,” and “temporary failure” without depending on fragile error-message strings?**

---

## Official Documentation

For authoritative API details, refer to the official Go `errors` package documentation:

https://pkg.go.dev/errors
