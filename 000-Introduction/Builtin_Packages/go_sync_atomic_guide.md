# `sync/atomic` Package in Go

The Go `sync/atomic` package provides **low-level atomic operations** for safely accessing and modifying shared variables when multiple goroutines run concurrently.

It is especially useful when you need to coordinate a **small piece of shared state**—for example, a counter, flag, pointer, or configuration value—without using a `sync.Mutex`.

The official documentation describes it as a package for low-level atomic memory primitives used to implement synchronization algorithms. It also recommends using channels or the higher-level `sync` package for many ordinary synchronization problems.

> **Important:** `sync/atomic` is not simply a "faster mutex." It solves a different class of problems. You should use it when the shared state can be manipulated atomically and the resulting design remains easy to reason about.

---

# 1. What is `sync/atomic`?

Imagine 10,000 goroutines incrementing the same counter:

```go
var count int

count++
```

It looks like one operation, but conceptually it is:

```text
READ count
ADD 1
WRITE count
```

Two goroutines can interleave these operations:

```text
Goroutine A: READ 100
Goroutine B: READ 100

Goroutine A: WRITE 101
Goroutine B: WRITE 101
```

You expected:

```text
102
```

but got:

```text
101
```

This is a **race condition**.

With `sync/atomic`:

```go
atomic.AddInt64(&count, 1)
```

the increment happens atomically.

The operation cannot be observed as a partially completed read-modify-write sequence by another atomic operation. Go's atomic operations also provide sequentially consistent ordering semantics.

---

# 2. When is `sync/atomic` commonly used?

Typical situations include:

- Concurrent counters
- Request/connection statistics
- Atomic boolean flags
- Lock-free state transitions
- Reference/state counters
- Updating a pointer safely
- Read-mostly configuration
- Feature flags
- Metrics
- Simple state machines
- Implementing low-level concurrent data structures

A common pattern looks like:

```text
HTTP requests
     │
     ├── Goroutine 1 ──┐
     ├── Goroutine 2 ──┤
     ├── Goroutine 3 ──┤
     ├── ...           ├──> atomic counter
     └── Goroutine N ──┘
```

---

# 3. The most important atomic concepts

Before learning every function, understand these five operations:

| Operation | Meaning |
|---|---|
| `Load` | Atomically read |
| `Store` | Atomically write |
| `Add` | Atomically increment/decrement |
| `Swap` | Atomically replace and get old value |
| `CompareAndSwap` | Change value only if it still equals an expected value |

Think of them as:

```text
Load             → "What is the value?"
Store            → "Set the value."
Add              → "Change the value."
Swap             → "Replace it and tell me the old value."
CompareAndSwap   → "Replace it only if nobody changed it."
```

---

# 4. Simple example — Atomic Counter

```go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64

	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			atomic.AddInt64(&counter, 1)
		}()
	}

	wg.Wait()

	fmt.Println("Counter:", atomic.LoadInt64(&counter))
}
```

Output:

```text
Counter: 1000
```

Without atomic operations, concurrently modifying `counter` would create a data race.

---

# 5. Every important function/type in `sync/atomic`

Modern Go provides both:

1. **Function-based APIs**
2. **Type-based APIs**

The type-based APIs are generally easier to use in new code. The typed APIs are more ergonomic and reduce some of the error-prone aspects of the older function-style API.

The package includes atomic types such as:

- `Bool`
- `Int32`
- `Int64`
- `Uint32`
- `Uint64`
- `Uintptr`
- `Pointer[T]`
- `Value`

---

# 6. `Add` operations

## `atomic.AddInt32`

```go
atomic.AddInt32(&counter, 1)
```

Atomically adds an `int32` value.

```go
var counter int32

atomic.AddInt32(&counter, 10)
```

If the counter was:

```text
50
```

it becomes:

```text
60
```

and the function returns:

```text
60
```

---

## `atomic.AddInt64`

```go
atomic.AddInt64(&counter, 1)
```

Same idea, but for `int64`.

Very common for:

```text
request counters
download counts
statistics
IDs
metrics
```

---

## `atomic.AddUint32`

```go
atomic.AddUint32(&counter, 1)
```

Used for `uint32`.

---

## `atomic.AddUint64`

```go
atomic.AddUint64(&counter, 1)
```

Used for `uint64`.

This is particularly common for counters that should never be negative.

---

## `atomic.AddUintptr`

```go
atomic.AddUintptr(&value, 1)
```

Operates on `uintptr`.

This is a lower-level type and is less commonly needed in ordinary application code.

---

## Modern alternative

Instead of:

```go
var counter atomic.Int64

counter.Add(1)
```

is generally preferable to manually manipulating an `int64` with the old function-style API.

The typed API was introduced in Go 1.19.

---

# 7. `Load` operations

`Load` means:

> "Read this value atomically."

## `LoadInt32`

```go
value := atomic.LoadInt32(&counter)
```

---

## `LoadInt64`

```go
value := atomic.LoadInt64(&counter)
```

---

## `LoadUint32`

```go
value := atomic.LoadUint32(&counter)
```

---

## `LoadUint64`

```go
value := atomic.LoadUint64(&counter)
```

---

## `LoadUintptr`

```go
value := atomic.LoadUintptr(&counter)
```

The important point is that the read itself is atomic.

Instead of:

```go
fmt.Println(counter)
```

for a variable concurrently modified by other goroutines, use the appropriate atomic load when that variable is intended to be accessed atomically.

---

# 8. `Store` operations

`Store` means:

> "Atomically replace the current value."

For example:

```go
var running int32

atomic.StoreInt32(&running, 1)
```

Another goroutine can safely perform:

```go
if atomic.LoadInt32(&running) == 1 {
	// server is running
}
```

---

## Available Store functions

```go
atomic.StoreInt32(...)
atomic.StoreInt64(...)
atomic.StoreUint32(...)
atomic.StoreUint64(...)
atomic.StoreUintptr(...)
```

There are also pointer-oriented forms and the typed APIs.

---

# 9. `Swap` operations

`Swap` performs two things atomically:

```text
1. Store new value
2. Return old value
```

For example:

```go
var state int32 = 1

old := atomic.SwapInt32(&state, 2)

fmt.Println(old)   // 1
fmt.Println(state) // 2
```

Conceptually:

```text
Before:

state = 1

Swap(state, 2)

After:

state = 2

Return:
old = 1
```

The important thing is that another atomic operation cannot sneak between the read of the old value and the write of the new value.

---

## Available Swap functions

```go
SwapInt32
SwapInt64
SwapUint32
SwapUint64
SwapUintptr
SwapPointer
```

The documentation describes `Swap` as the atomic equivalent of:

```go
old := *addr
*addr = new
return old
```

but performed atomically.

---

# 10. `CompareAndSwap` — CAS

This is probably the **most important advanced atomic operation** to understand.

`CompareAndSwap` means:

> "If the current value is exactly what I expect, replace it. Otherwise, do nothing."

Example:

```go
var state int32 = 0

success := atomic.CompareAndSwapInt32(
	&state,
	0,
	1,
)

fmt.Println(success)
```

The operation means:

```text
IF state == 0
    state = 1
    return true
ELSE
    don't change state
    return false
```

---

## Example

Suppose:

```text
state = 0
```

Goroutine A executes:

```go
atomic.CompareAndSwapInt32(&state, 0, 1)
```

Result:

```text
true

state = 1
```

Now Goroutine B executes:

```go
atomic.CompareAndSwapInt32(&state, 0, 1)
```

Result:

```text
false

state = 1
```

Why?

Because the expected old value was:

```text
0
```

but the actual value was already:

```text
1
```

---

# 11. Why CAS is powerful

CAS allows you to build algorithms like:

```text
if nobody changed the value:
       update it
else:
       retry / handle conflict
```

For example:

```go
for {
	old := atomic.LoadInt64(&value)

	newValue := old + 1

	if atomic.CompareAndSwapInt64(
		&value,
		old,
		newValue,
	) {
		break
	}
}
```

This is essentially:

```text
READ
  ↓
calculate new value
  ↓
CAS
  ↓
success? ── yes → done
  │
  no
  ↓
retry
```

This pattern is fundamental to many lock-free algorithms.

---

# 12. Bitwise `And` operations

Modern Go also provides atomic bitwise operations such as:

```go
AndInt32
AndInt64
AndUint32
AndUint64
AndUintptr
```

These were added in Go 1.23.

For example:

```go
var flags atomic.Uint32

flags.Store(0b1111)

old := flags.And(0b1100)
```

The operation performs:

```text
current = current AND mask
```

So:

```text
1111
AND 1100
---------
1100
```

The operation returns the **old value**.

This is useful when individual bits represent flags.

For example:

```text
bit 0 → logged in
bit 1 → admin
bit 2 → verified
bit 3 → premium
```

You can atomically set or clear individual flags.

---

# 13. Bitwise `Or` operations

Similarly:

```text
OrInt32
OrInt64
OrUint32
OrUint64
OrUintptr
```

perform an atomic OR.

Example:

```go
var flags atomic.Uint32

flags.Store(0)

flags.Or(0b0001)
flags.Or(0b0100)
```

The final value becomes:

```text
0101
```

This is useful for maintaining multiple boolean states inside one integer.

For example:

```text
bit 0 → logged in
bit 1 → admin
bit 2 → verified
bit 3 → premium
```

---

# 14. `atomic.Bool`

Modern Go provides:

```go
var running atomic.Bool
```

The zero value is:

```text
false
```

You can write:

```go
running.Store(true)
```

Read:

```go
if running.Load() {
	fmt.Println("Running")
}
```

Swap:

```go
old := running.Swap(false)
```

CAS:

```go
success := running.CompareAndSwap(false, true)
```

This is much clearer than manually maintaining:

```go
var running int32
```

with:

```go
atomic.StoreInt32(&running, 1)
```

---

# 15. `atomic.Int32`

Instead of:

```go
var counter int32
atomic.AddInt32(&counter, 1)
```

you can use:

```go
var counter atomic.Int32

counter.Add(1)
```

Read:

```go
value := counter.Load()
```

Write:

```go
counter.Store(100)
```

Swap:

```go
old := counter.Swap(200)
```

CAS:

```go
counter.CompareAndSwap(200, 300)
```

And:

```go
counter.And(mask)
```

Or:

```go
counter.Or(mask)
```

---

# 16. `atomic.Int64`

Same idea:

```go
var counter atomic.Int64
```

Operations:

```go
counter.Add(1)

counter.Load()

counter.Store(100)

counter.Swap(200)

counter.CompareAndSwap(200, 300)

counter.And(mask)

counter.Or(mask)
```

`Int64` is particularly useful for high-volume counters.

---

# 17. `atomic.Uint32`

```go
var counter atomic.Uint32
```

Operations:

```go
counter.Add(1)
counter.Load()
counter.Store(100)
counter.Swap(200)
counter.CompareAndSwap(200, 300)
counter.And(mask)
counter.Or(mask)
```

---

# 18. `atomic.Uint64`

```go
var requests atomic.Uint64

requests.Add(1)

fmt.Println(requests.Load())
```

This is an excellent choice for things like:

```text
total requests
total bytes
total jobs
total events
```

---

# 19. `atomic.Uintptr`

```go
var value atomic.Uintptr
```

It supports:

```go
Add()
Load()
Store()
Swap()
CompareAndSwap()
And()
Or()
```

`uintptr` is a low-level integer type associated with pointer-sized values, so ordinary application code usually doesn't need it.

---

# 20. `atomic.Pointer[T]`

This is extremely useful when you want to atomically replace a pointer.

Example:

```go
type Config struct {
	Port int
	Host string
}

var config atomic.Pointer[Config]

config.Store(&Config{
	Port: 8080,
	Host: "localhost",
})

current := config.Load()

fmt.Println(current.Port)
```

You can replace the entire configuration:

```go
config.Store(&Config{
	Port: 9090,
	Host: "example.com",
})
```

And readers can safely load the current pointer:

```go
current := config.Load()
```

It also supports:

```go
config.Swap(newConfig)
```

and:

```go
config.CompareAndSwap(oldConfig, newConfig)
```

The generic `Pointer[T]` type was introduced in Go 1.19.

---

# 21. `atomic.Value`

`atomic.Value` is different from the integer atomic types.

It allows you to atomically store and load a consistently typed Go value.

Example:

```go
var config atomic.Value

config.Store("production")

current := config.Load()

fmt.Println(current)
```

Output:

```text
production
```

You could store a configuration structure:

```go
type Config struct {
	Port int
	Debug bool
}

var config atomic.Value

config.Store(Config{
	Port: 8080,
	Debug: false,
})
```

Later:

```go
current := config.Load().(Config)

fmt.Println(current.Port)
```

This is particularly useful for **read-mostly configuration**.

The official documentation demonstrates using `atomic.Value` to periodically replace server configuration while worker goroutines read the latest configuration.

---

# 22. `atomic.Value` type restrictions

There is an important rule.

Once you store:

```go
config.Store(Config{})
```

you should continue storing the same concrete type.

Don't do:

```go
config.Store(Config{})
config.Store("hello")
```

That causes a panic because the concrete types are inconsistent.

Also:

```go
config.Store(nil)
```

is invalid.

`Load()` returns `nil` if nothing has been stored yet.

---

# 23. `atomic.Value.CompareAndSwap`

You can conditionally replace a value:

```go
var state atomic.Value

state.Store("idle")

success := state.CompareAndSwap(
	"idle",
	"running",
)
```

Conceptually:

```text
if current == "idle":
    current = "running"
```

It returns:

```text
true
```

if the swap occurred.

---

# 24. `atomic.Value.Swap`

You can replace a stored value and receive the old value:

```go
old := state.Swap("stopped")
```

If the previous value was:

```text
"running"
```

then:

```text
old = "running"
```

and the stored value becomes:

```text
"stopped"
```

---

# 25. Complete API mental model

A useful way to remember the package is:

```text
                    sync/atomic
                         │
       ┌─────────────────┼─────────────────┐
       │                 │                 │
      Read              Write           Modify
       │                 │                 │
     Load               Store          Add / Swap
       │                                   │
       └────────────── CAS ────────────────┘
                         │
                CompareAndSwap
```

For bit manipulation:

```text
              Atomic bit operations
                     │
                ┌────┴────┐
                │         │
               And       Or
```

For data types:

```text
Bool
Int32
Int64
Uint32
Uint64
Uintptr
Pointer[T]
Value
```

---

# 26. Three common beginner mistakes

## Mistake 1: Using normal `++` on shared data

Bad:

```go
var count int64

go func() {
	count++
}()
```

If multiple goroutines access `count`, this can create a data race.

Better:

```go
var count atomic.Int64

count.Add(1)
```

---

## Mistake 2: Thinking atomic operations make an entire data structure safe

Suppose:

```go
type Counter struct {
	value atomic.Int64
	name  string
}
```

Making `value` atomic does **not** automatically make every operation involving `Counter` atomic.

Atomicity applies to the particular atomic operation.

It does not magically make an entire multi-step transaction atomic.

For example:

```text
read A
read B
calculate
write A
write B
```

may require a mutex or another synchronization design.

---

## Mistake 3: Using atomics when a mutex or channel is clearer

Atomic code can become difficult to understand.

If your problem is:

```text
10 goroutines need to modify a complex map
```

a `sync.Mutex` may be much easier to reason about.

The Go documentation cautions that, except for special low-level applications, channels or facilities from `sync` are often preferable.

A good rule is:

```text
Simple shared value
       ↓
     atomic

Complex shared state
       ↓
   Mutex / RWMutex

Goroutine communication
       ↓
    Channels
```

---

# 27. Real-world application #1 — HTTP request counter

Imagine a web server handling thousands of requests per second.

You want:

```text
Total requests = 8,742,931
```

Every request can execute:

```go
requests.Add(1)
```

while a monitoring goroutine can execute:

```go
fmt.Println(requests.Load())
```

No lock is required just to maintain the counter.

This is an excellent atomic use case because the shared state is extremely simple.

---

# 28. Real-world application #2 — Hot configuration replacement

Suppose a server has:

```go
type Config struct {
	MaxConnections int
	TimeoutSeconds int
}
```

Workers continuously read configuration while an administrator periodically reloads it.

Using `atomic.Pointer[Config]`, you can replace the entire configuration:

```text
Old configuration
       │
       │ Store(new config)
       ▼
New configuration
```

Existing readers can continue using the configuration they already loaded, while new readers see the new configuration.

This is a classic **read-mostly / copy-on-write** pattern demonstrated in the official documentation.

---

# 29. Important rule: Don't copy atomic values

Atomic types such as:

```text
atomic.Int64
atomic.Bool
atomic.Uint64
atomic.Pointer[T]
atomic.Value
```

must not be copied after first use.

For example, be careful with:

```go
func process(x atomic.Int64) {
	// ...
}
```

because passing it by value copies it.

Prefer:

```go
func process(x *atomic.Int64) {
	// ...
}
```

or design your structure so the atomic field isn't copied.

---

# 30. Atomic operations and the Go memory model

This is the deeper part.

Atomic operations aren't merely about preventing:

```text
counter = wrong number
```

They also establish synchronization relationships.

Go specifies that when the effect of one atomic operation is observed by another atomic operation, the first operation synchronizes before the second. Atomic operations behave as though executed in a sequentially consistent order.

So atomic operations can be used not only for counters but also for coordinating visibility of state between goroutines.

---

# 31. Atomic vs Mutex

Consider a counter.

### Atomic

```go
var counter atomic.Int64

counter.Add(1)
```

### Mutex

```go
var (
	counter int64
	mu      sync.Mutex
)

mu.Lock()
counter++
mu.Unlock()
```

Both can provide correct synchronization.

But atomic is very natural for the simple operation:

```text
increment one number
```

A mutex becomes more attractive when the critical section contains multiple related operations.

For example:

```go
mu.Lock()

balance -= amount
transactionCount++
lastTransaction = transaction

mu.Unlock()
```

These operations collectively form one logical transaction. An atomic counter alone isn't enough to make that entire operation atomic.

---

# 32. Three progressively challenging exercises

As requested, **no solutions**.

## Exercise 1 — Concurrent Request Counter

Create a program that starts **1,000 goroutines**.

Each goroutine should:

```text
simulate one request
increment a shared request counter
```

After all goroutines finish, print the total number of requests.

Requirements:

- Use `atomic.Int64`
- Use `sync.WaitGroup`
- Do not use a mutex for the counter
- The final result must always be exactly `1000`

---

## Exercise 2 — Atomic Server State

Create a small server-state simulation with these states:

```text
STOPPED
RUNNING
PAUSED
```

Multiple goroutines will attempt to change the server state.

Requirements:

- Store the state using an atomic integer type.
- Use `Load()` to inspect the current state.
- Use `CompareAndSwap()` to perform a state transition.
- A transition should only succeed when the server is in the expected previous state.
- Print whether each transition succeeded or failed.

Example conceptual transitions:

```text
STOPPED → RUNNING
RUNNING → PAUSED
PAUSED  → RUNNING
RUNNING → STOPPED
```

The goal is to understand why **CAS is useful when multiple goroutines may try to change the same state simultaneously**.

---

## Exercise 3 — Lock-Free Configuration Updates

Create a program representing a production server whose configuration can be updated while many workers are processing requests.

Define:

```go
type Config struct {
	MaxConnections int
	TimeoutSeconds int
	Version        int
}
```

Requirements:

- Store the current configuration using `atomic.Pointer[Config]`.
- Start multiple worker goroutines.
- Workers should continuously load the current configuration.
- Another goroutine should periodically create a completely new `Config`.
- Atomically replace the current configuration.
- Workers must never mutate the configuration they loaded.
- Print which configuration version each worker observes.
- Experiment with many workers and frequent configuration updates.
- Use the race detector:

```bash
go run -race .
```

Your goal is to understand the **read-mostly / copy-on-write** pattern and why replacing an immutable configuration pointer can be safer than modifying a shared configuration object in place.

---

# 33. Quick cheat sheet

| Need | Use |
|---|---|
| Atomic boolean | `atomic.Bool` |
| Atomic 32-bit integer | `atomic.Int32` / `atomic.Uint32` |
| Atomic 64-bit integer | `atomic.Int64` / `atomic.Uint64` |
| Atomic pointer | `atomic.Pointer[T]` |
| Atomic arbitrary consistently typed value | `atomic.Value` |
| Read | `Load()` |
| Write | `Store()` |
| Increment/decrement | `Add()` |
| Replace + get old value | `Swap()` |
| Conditional replacement | `CompareAndSwap()` |
| Atomic bit clearing/masking | `And()` |
| Atomic bit setting | `Or()` |

The current standard library documents the typed APIs and the legacy function APIs together; `And`/`Or` are available in the newer API surface from Go 1.23 onward.

---

# 34. The deeper question

Suppose you are building a web server that receives **10,000 concurrent requests**.

You need to maintain:

```text
total requests
active requests
successful requests
failed requests
current configuration
server state
```

Some values can be represented by a single atomic integer, while others involve several related pieces of state.

**Thought-provoking question:**

> **How would you decide which of these pieces of state should use `sync/atomic`, which should use `sync.Mutex`, and which should be communicated through channels—and what could go wrong if you chose `atomic` simply because it appears faster?**

That question gets to the real purpose of `sync/atomic`: **not merely knowing its functions, but recognizing when atomicity is the correct synchronization model.**
