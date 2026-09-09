# Go `context` Package — Detailed Guide

The Go standard library's `context` package is one of the most important packages to understand once you start writing concurrent programs, HTTP servers, database-backed applications, and distributed services.

The `context` package provides cancellation, deadlines/timeouts, request-scoped values, and cancellation-cause support.

---

## 1. What is the `context` package?

Import it with:

```go
import "context"
```

A `context.Context` carries three major kinds of information through a chain of function calls:

1. **Cancellation signals** — "Stop what you're doing."
2. **Deadlines/timeouts** — "Stop if this takes too long."
3. **Request-scoped values** — "Carry this request-related information along."

This is particularly important when one operation starts several other operations:

```text
HTTP Request
     │
     ▼
Handler
     │
     ├──► Service
     │       │
     │       ├──► Database
     │       │
     │       └──► External API
     │
     └──► Another Goroutine
```

The same context can flow through all of them.

If the request is canceled, derived contexts can be canceled too. This prevents unnecessary work from continuing after the operation that required it has disappeared.

### Commonly used in

- HTTP servers
- HTTP clients
- Database queries
- Goroutines
- Concurrent workers
- Microservices
- RPC calls
- External API calls
- Request tracing
- Request-scoped metadata
- Timeouts and cancellation

---

# 2. The `Context` interface

The central type is:

```go
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key any) any
}
```

There are **four methods** you should understand.

---

## `Deadline()`

```go
Deadline() (deadline time.Time, ok bool)
```

Returns the deadline associated with the context.

Example:

```go
ctx, cancel := context.WithTimeout(
    context.Background(),
    5*time.Second,
)
defer cancel()

deadline, ok := ctx.Deadline()

if ok {
    fmt.Println("Deadline:", deadline)
}
```

If the context has no deadline:

```go
deadline, ok := ctx.Deadline()

fmt.Println(ok) // false
```

### When useful

You might want to know how much time remains:

```go
if deadline, ok := ctx.Deadline(); ok {
    remaining := time.Until(deadline)
    fmt.Println("Remaining:", remaining)
}
```

---

# `Done()`

```go
Done() <-chan struct{}
```

Returns a channel that is closed when the context is canceled or its deadline expires.

This is one of the most important context mechanisms.

Example:

```go
func worker(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            fmt.Println("Worker stopped")
            return

        default:
            fmt.Println("Working...")
            time.Sleep(time.Second)
        }
    }
}
```

The important pattern is:

```go
select {
case <-ctx.Done():
    return
}
```

You can think of:

```go
ctx.Done()
```

as:

> "Tell me when I should stop."

The `Done` channel can be `nil` for contexts that cannot be canceled, such as `context.Background()` and `context.WithoutCancel()`.

---

# `Err()`

```go
Err() error
```

Returns the reason the context was canceled.

Typically you'll encounter:

```go
context.Canceled
```

or:

```go
context.DeadlineExceeded
```

Example:

```go
select {
case <-ctx.Done():
    fmt.Println(ctx.Err())
}
```

Possible output:

```text
context canceled
```

or:

```text
context deadline exceeded
```

You can compare errors safely:

```go
if errors.Is(ctx.Err(), context.Canceled) {
    fmt.Println("Request was canceled")
}
```

And:

```go
if errors.Is(ctx.Err(), context.DeadlineExceeded) {
    fmt.Println("Request timed out")
}
```

---

# `Value()`

```go
Value(key any) any
```

Retrieves a request-scoped value stored in the context.

Example:

```go
type contextKey string

const userKey contextKey = "user"

ctx := context.WithValue(
    context.Background(),
    userKey,
    "Chandu",
)

user := ctx.Value(userKey)

fmt.Println(user)
```

Output:

```text
Chandu
```

However, **do not use context values as a general-purpose parameter bag**.

Use context values primarily for request-scoped data that travels across API/process boundaries, rather than optional function parameters.

---

# 3. Creating contexts

There are several functions for creating derived contexts.

---

# `context.Background()`

```go
context.Background()
```

Creates an empty root context.

It:

- is never canceled
- has no deadline
- contains no values

Example:

```go
ctx := context.Background()
```

Commonly used in:

- `main()`
- tests
- initialization
- creating a root context

Example:

```go
func main() {
    ctx := context.Background()

    startApplication(ctx)
}
```

---

# `context.TODO()`

```go
context.TODO()
```

Creates an empty context just like `Background()`, but semantically communicates:

> "I don't know which context should be used here yet."

Example:

```go
func legacyFunction() {
    ctx := context.TODO()

    doSomething(ctx)
}
```

Use `Background()` when you know this is the root context and `TODO()` when you're temporarily unsure which context belongs there.

---

# `context.WithCancel()`

```go
context.WithCancel(parent)
```

Returns:

```text
child context
+
cancel function
```

Example:

```go
ctx, cancel := context.WithCancel(context.Background())

defer cancel()
```

Calling:

```go
cancel()
```

cancels the child context and contexts derived from it.

### Complete example

```go
package main

import (
    "context"
    "fmt"
    "time"
)

func worker(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            fmt.Println("Worker stopped:", ctx.Err())
            return

        default:
            fmt.Println("Working...")
            time.Sleep(500 * time.Millisecond)
        }
    }
}

func main() {
    ctx, cancel := context.WithCancel(context.Background())

    go worker(ctx)

    time.Sleep(2 * time.Second)

    fmt.Println("Canceling worker...")
    cancel()

    time.Sleep(500 * time.Millisecond)
}
```

The worker receives the cancellation signal and exits.

### Important

Normally call the returned cancellation function:

```go
ctx, cancel := context.WithCancel(parent)
defer cancel()
```

---

# `context.WithTimeout()`

```go
context.WithTimeout(parent, timeout)
```

Creates a context that automatically cancels after a duration.

Example:

```go
ctx, cancel := context.WithTimeout(
    context.Background(),
    2*time.Second,
)
defer cancel()
```

This means:

> "This operation gets at most two seconds."

Example:

```go
func main() {
    ctx, cancel := context.WithTimeout(
        context.Background(),
        2*time.Second,
    )
    defer cancel()

    select {
    case <-time.After(5 * time.Second):
        fmt.Println("Operation completed")

    case <-ctx.Done():
        fmt.Println("Operation timed out:", ctx.Err())
    }
}
```

Output:

```text
Operation timed out: context deadline exceeded
```

`WithTimeout` is essentially a convenience around `WithDeadline` using `time.Now().Add(timeout)`.

---

# `context.WithDeadline()`

```go
context.WithDeadline(parent, deadline)
```

Instead of saying:

> "Give this operation 5 seconds."

you specify an exact time:

```go
deadline := time.Now().Add(5 * time.Second)

ctx, cancel := context.WithDeadline(
    context.Background(),
    deadline,
)

defer cancel()
```

The context is canceled when:

1. the deadline arrives,
2. you call `cancel()`, or
3. the parent context is canceled.

The earliest applicable deadline wins.

---

# `context.WithValue()`

```go
context.WithValue(parent, key, value)
```

Stores request-scoped data.

Example:

```go
type contextKey string

const requestIDKey contextKey = "requestID"

ctx := context.WithValue(
    context.Background(),
    requestIDKey,
    "req-123",
)

requestID := ctx.Value(requestIDKey)

fmt.Println(requestID)
```

Output:

```text
req-123
```

### Don't do this

```go
context.WithValue(ctx, "userID", 123)
```

Using ordinary strings as keys can cause collisions between packages.

Prefer a custom key type:

```go
type contextKey string

const userIDKey contextKey = "userID"
```

---

# `context.WithCancelCause()`

This cancellation API allows you to specify **why** cancellation happened.

```go
context.WithCancelCause(parent)
```

Example:

```go
ctx, cancel := context.WithCancelCause(
    context.Background(),
)

err := errors.New("database connection lost")

cancel(err)

fmt.Println(ctx.Err())
fmt.Println(context.Cause(ctx))
```

Conceptually:

```text
ctx.Err()
    ↓
context canceled

context.Cause(ctx)
    ↓
database connection lost
```

The distinction is useful because:

```go
ctx.Err()
```

tells you the general cancellation state, while:

```go
context.Cause(ctx)
```

can tell you the specific reason.

---

# `context.Cause()`

```go
context.Cause(ctx)
```

Returns the reason a context was canceled.

Example:

```go
ctx, cancel := context.WithCancelCause(
    context.Background(),
)

cancel(errors.New("worker failed"))

fmt.Println(context.Cause(ctx))
```

Output:

```text
worker failed
```

If there is no custom cause, `Cause` falls back to the context's cancellation error.

---

# `context.WithDeadlineCause()`

```go
context.WithDeadlineCause(parent, deadline, cause)
```

Similar to `WithDeadline`, but lets you specify the error that should be reported through:

```go
context.Cause(ctx)
```

when the deadline expires.

Example:

```go
ctx, cancel := context.WithDeadlineCause(
    context.Background(),
    time.Now().Add(time.Second),
    errors.New("payment operation expired"),
)
defer cancel()
```

If the deadline expires:

```go
context.Cause(ctx)
```

can provide the more meaningful application-specific cause.

---

# `context.WithTimeoutCause()`

```go
context.WithTimeoutCause(parent, timeout, cause)
```

Similar to `WithTimeout`, but associates a custom cause with timeout cancellation.

Example:

```go
ctx, cancel := context.WithTimeoutCause(
    context.Background(),
    2*time.Second,
    errors.New("external payment service timed out"),
)
defer cancel()
```

Then:

```go
context.Cause(ctx)
```

can explain the application-level reason for the timeout.

---

# `context.WithoutCancel()`

```go
context.WithoutCancel(parent)
```

Creates a context derived from another context but **not canceled when the parent is canceled**.

Example:

```go
backgroundCtx := context.WithoutCancel(requestCtx)
```

This can be useful when some follow-up work should survive the lifetime of the original request.

Conceptually:

```text
HTTP Request
     │
     ├── Normal request work
     │
     └── Follow-up work
             │
             └── should survive request cancellation
```

`WithoutCancel` has no deadline, its `Done()` channel is `nil`, and its `Err()` returns `nil`.

### Be careful

This function deliberately breaks normal cancellation propagation. Don't use it simply to "make cancellation go away."

---

# `context.AfterFunc()`

```go
context.AfterFunc(ctx, f)
```

Registers a function that will run in its own goroutine when the context is canceled.

Example:

```go
stop := context.AfterFunc(ctx, func() {
    fmt.Println("Context was canceled")
})
```

The returned function:

```go
stop()
```

attempts to prevent the callback from running.

It returns:

```text
true
```

if it successfully stopped the association, or:

```text
false
```

if the callback has already started or was already stopped.

Important: `stop()` does **not** wait for an already-running callback to finish.

---

# `CancelFunc`

```go
type CancelFunc func()
```

A `CancelFunc` tells a context:

> "Cancel this context now."

Example:

```go
ctx, cancel := context.WithCancel(parent)

defer cancel()
```

Calling it multiple times is safe. After the first cancellation, subsequent calls have no effect.

It also does not wait for the canceled work itself to finish.

---

# `CancelCauseFunc`

```go
type CancelCauseFunc func(cause error)
```

This is the cancellation function returned by:

```go
context.WithCancelCause()
```

Example:

```go
ctx, cancel := context.WithCancelCause(parent)

cancel(errors.New("worker crashed"))
```

Then:

```go
context.Cause(ctx)
```

returns the supplied error.

---

# 4. The most important context pattern

You'll see this pattern constantly in real Go code:

```go
func DoSomething(ctx context.Context) error {
    select {
    case <-ctx.Done():
        return ctx.Err()

    default:
        // Do work
    }

    return nil
}
```

Function signatures generally look like:

```go
func GetUser(ctx context.Context, id int) (*User, error)
```

not:

```go
func GetUser(id int, ctx context.Context)
```

The `Context` is conventionally the **first parameter**, normally named `ctx`.

---

# 5. Simple complete example

Here is a small program combining timeout + cancellation + `Done()` + `Err()`:

```go
package main

import (
    "context"
    "fmt"
    "time"
)

func doWork(ctx context.Context) error {
    for i := 1; i <= 10; i++ {

        select {
        case <-ctx.Done():
            return ctx.Err()

        default:
            fmt.Println("Processing item", i)
            time.Sleep(500 * time.Millisecond)
        }
    }

    return nil
}

func main() {
    ctx, cancel := context.WithTimeout(
        context.Background(),
        2*time.Second,
    )
    defer cancel()

    err := doWork(ctx)

    if err != nil {
        fmt.Println("Work stopped:", err)
        return
    }

    fmt.Println("Work completed")
}
```

Because each iteration takes approximately 500 ms, the 2-second timeout will prevent all ten iterations from completing.

---

# 6. Three common beginner mistakes

## Mistake 1: Forgetting `cancel()`

Bad:

```go
ctx, cancel := context.WithTimeout(
    context.Background(),
    5*time.Second,
)

doSomething(ctx)
```

Better:

```go
ctx, cancel := context.WithTimeout(
    context.Background(),
    5*time.Second,
)
defer cancel()

doSomething(ctx)
```

Even if the operation finishes early, calling `cancel()` releases resources associated with the derived context.

---

## Mistake 2: Creating a context but ignoring `Done()`

This is a major conceptual mistake:

```go
func worker(ctx context.Context) {
    for {
        expensiveOperation()
    }
}
```

The caller can cancel `ctx`, but the worker never checks it.

Better:

```go
func worker(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            return
        default:
            expensiveOperation()
        }
    }
}
```

A context doesn't magically terminate your goroutine. **Your code must cooperate with cancellation.**

---

## Mistake 3: Using `context.WithValue()` for everything

Bad:

```go
ctx = context.WithValue(ctx, "username", username)
ctx = context.WithValue(ctx, "age", age)
ctx = context.WithValue(ctx, "theme", theme)
ctx = context.WithValue(ctx, "database", db)
```

Context values aren't intended to replace ordinary function parameters or application state.

Use them primarily for request-scoped information that needs to travel through API boundaries.

---

# 7. Two real-world applications

## Application 1: HTTP request cancellation

Imagine:

```text
Browser
   │
   ▼
HTTP Server
   │
   ▼
Handler
   │
   ├── Database
   │
   └── External API
```

Suppose the browser disconnects.

The server doesn't necessarily need to continue spending resources on:

```text
Database query
External API request
Expensive computation
```

If those operations receive the request's context and respect cancellation, they can stop their work.

For example:

```go
func handler(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()

    result, err := fetchData(ctx)

    if err != nil {
        return
    }

    fmt.Fprintln(w, result)
}
```

The context can propagate through the application's service and database layers.

---

# Application 2: Microservices and time budgets

Imagine:

```text
API Gateway
     │
     ▼
Order Service
     │
     ├── User Service
     ├── Inventory Service
     └── Payment Service
```

Suppose the original request has a 3-second budget.

You don't want the Payment Service to keep processing after the overall request has already timed out.

A context can carry the cancellation/deadline through the service chain.

Conceptually:

```text
3-second request deadline
          │
          ▼
     Order Service
       /       \
      /         \
 User Service   Payment Service
      │              │
      └──────┬───────┘
             ▼
       shared deadline
```

This is one of the reasons contexts are so important in distributed Go applications.

---

# 8. A useful mental model

Think of a context as a **control signal traveling alongside your actual data**.

Your normal data:

```text
User ID
Order ID
Product ID
```

Your context:

```text
Should I stop?
When must I stop?
What request am I part of?
Why was I canceled?
```

So:

```text
Function arguments
       +
Context
       │
       ▼
     Function
```

The context isn't your business data.

It is **control and request-scoped metadata**.

---

# 9. Important rules to remember

### Rule 1

Pass context explicitly:

```go
func Process(ctx context.Context, data Data)
```

rather than hiding it in a struct.

### Rule 2

Usually make it the first parameter:

```go
func Process(ctx context.Context, ...)
```

### Rule 3

Don't pass `nil` as a context.

If you're genuinely unsure what context to use, use:

```go
context.TODO()
```

### Rule 4

Don't store contexts in structs unless you have a very specific reason.

### Rule 5

Don't use context values as a general-purpose parameter container.

### Rule 6

When you create a cancellable context:

```go
ctx, cancel := context.WithTimeout(...)
defer cancel()
```

### Rule 7

Cancellation is cooperative.

This:

```go
cancel()
```

doesn't forcibly kill a goroutine.

Your goroutine needs to observe:

```go
<-ctx.Done()
```

and stop itself.

---

# 10. Three progressively challenging exercises

## Exercise 1 — Basic cancellation

Create a program with:

- A `worker(ctx context.Context)` function.
- The worker should continuously print a message every 500 ms.
- `main()` should create a context using `context.WithCancel()`.
- Run the worker in a goroutine.
- After 3 seconds, cancel the context.
- The worker must detect cancellation through `ctx.Done()` and exit gracefully.
- Print the value of `ctx.Err()` when the worker stops.

**Do not use `time.After()` for cancellation.**

---

## Exercise 2 — Timeout + multiple goroutines

Build a program that simulates an e-commerce order:

```text
Order Processing
      │
      ├── Check Inventory
      ├── Calculate Shipping
      └── Validate Payment
```

Requirements:

- Create a context with a 3-second timeout.
- Run the three operations concurrently.
- Each operation should periodically check `ctx.Done()`.
- Make each operation take a different amount of time.
- If the timeout occurs, all unfinished operations must stop.
- Use a channel to collect the results.
- Determine whether the overall operation completed successfully or timed out.

The goal is to understand **context propagation across multiple goroutines**.

---

## Exercise 3 — Mini microservice simulation

Build a simulated request pipeline:

```text
HTTP Request
     │
     ▼
Order Service
     │
     ├── User Service
     ├── Inventory Service
     └── Payment Service
```

Requirements:

- Create a root context representing the incoming request.
- Give the request a maximum deadline.
- Pass the context through every service function.
- Each service should run work in its own goroutine.
- Each service should respect cancellation.
- Add a request ID using `context.WithValue()`.
- Use a custom context key type.
- Simulate one service failing.
- Use `context.WithCancelCause()` to propagate the failure.
- At the top level, distinguish:
  - successful completion,
  - normal cancellation,
  - deadline exceeded,
  - cancellation caused by a specific service failure.
- Use `context.Cause()` to identify the actual cancellation reason.
- Ensure all goroutines eventually terminate.

This exercise combines **cancellation, deadlines, goroutines, values, error handling, and cancellation causes**.

---

# 🤔 Thought-provoking question

Imagine your Go API receives **10,000 concurrent requests**, and each request starts 5 goroutines that call different services.

**If the client disconnects after 100 ms, but your goroutines don't properly propagate and check `context.Context`, what could happen to your application's memory, goroutine count, database connections, and overall performance over time?**

And a deeper question:

> **Is `context.Context` primarily a mechanism for passing information, or is it actually a mechanism for controlling the lifetime of work? Why?**

That distinction becomes extremely important when you start building production-grade concurrent Go applications.
