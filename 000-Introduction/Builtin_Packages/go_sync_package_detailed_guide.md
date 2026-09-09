# Go `sync` Package — Detailed Guide

The Go `sync` package provides **low-level synchronization primitives** for coordinating multiple goroutines that access shared data or need to work together safely.

It is commonly used when your program has **concurrent goroutines** and you need to prevent problems such as:

- Data races
- Multiple goroutines modifying the same data simultaneously
- One goroutine reading data while another is modifying it
- Waiting for a group of goroutines to finish
- Building thread-safe data structures
- Controlling access to shared resources

## Main types and functions

| Type / Function | Purpose |
|---|---|
| `sync.Mutex` | Mutual exclusion — only one goroutine enters a critical section |
| `sync.RWMutex` | Allows multiple readers but only one writer |
| `sync.WaitGroup` | Waits for a collection of goroutines to finish |
| `sync.Once` | Ensures an operation runs exactly once |
| `sync.Cond` | Allows goroutines to wait for/signaling a condition |
| `sync.Map` | Concurrent map designed for specific access patterns |
| `sync.Pool` | Temporary reuse of allocated objects |
| `sync.OnceFunc` | Creates a function that executes at most once |
| `sync.OnceValue` | Creates a function that computes and caches one value |
| `sync.OnceValues` | Creates a function that computes and caches multiple values |

> **Important:** `sync` primitives should generally be used to protect or coordinate shared state. If you can design your program so that goroutines communicate through channels instead of sharing mutable memory, that can often be simpler.

---

# 1. `sync.Mutex`

`Mutex` stands for **mutual exclusion**.

It ensures that only **one goroutine at a time** can execute a protected critical section.

```go
var mu sync.Mutex
```

The two most important methods are:

```go
mu.Lock()
mu.Unlock()
```

## `Lock()`

Acquires the mutex.

If another goroutine already owns the mutex, the calling goroutine waits until the mutex becomes available.

```go
mu.Lock()

// Critical section

mu.Unlock()
```

## `Unlock()`

Releases the mutex.

Once released, another waiting goroutine can acquire it.

### Example

```go
package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	mu      sync.Mutex
)

func increment() {
	mu.Lock()
	counter++
	mu.Unlock()
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			increment()
		}()
	}

	wg.Wait()

	fmt.Println("Counter:", counter)
}
```

Without the mutex, multiple goroutines could simultaneously read and modify `counter`.

## Best practice: `defer mu.Unlock()`

Usually write:

```go
mu.Lock()
defer mu.Unlock()

counter++
```

rather than:

```go
mu.Lock()

counter++

mu.Unlock()
```

The `defer` approach helps prevent accidentally forgetting `Unlock()` when the function has multiple return paths.

---

# 2. `sync.RWMutex`

`RWMutex` means **Read/Write Mutex**.

It is useful when:

- Many goroutines frequently **read**
- Relatively few goroutines **write**

Unlike `Mutex`, `RWMutex` distinguishes between readers and writers.

```go
var rwmu sync.RWMutex
```

It provides:

```go
Lock()
Unlock()

RLock()
RUnlock()
```

## `RLock()`

Acquires the **read lock**.

Multiple goroutines can hold the read lock simultaneously.

```go
rwmu.RLock()
defer rwmu.RUnlock()

// Read shared data
```

## `RUnlock()`

Releases the read lock.

```go
rwmu.RUnlock()
```

## `Lock()`

Acquires the **write lock**.

While a writer has the lock, other readers and writers cannot enter the protected section.

```go
rwmu.Lock()
defer rwmu.Unlock()

// Modify shared data
```

## `Unlock()`

Releases the write lock.

```go
rwmu.Unlock()
```

### Example

```go
package main

import (
	"fmt"
	"sync"
)

type Config struct {
	Port int
	Host string
}

var (
	config Config
	mu     sync.RWMutex
)

func getConfig() Config {
	mu.RLock()
	defer mu.RUnlock()

	return config
}

func updateConfig(port int, host string) {
	mu.Lock()
	defer mu.Unlock()

	config.Port = port
	config.Host = host
}

func main() {
	config = Config{
		Port: 8080,
		Host: "localhost",
	}

	fmt.Println(getConfig())

	updateConfig(9090, "example.com")

	fmt.Println(getConfig())
}
```

### Mental model

```text
RLock
 ├── Reader 1
 ├── Reader 2
 ├── Reader 3
 └── Reader 4

Many readers → allowed simultaneously
```

But:

```text
Lock
 └── Writer

Writer → exclusive access
```

---

# 3. `sync.WaitGroup`

`WaitGroup` is used when you want one goroutine to **wait for multiple goroutines to finish**.

The common methods are:

```go
Add()
Done()
Wait()
```

## `Add(delta int)`

Adds to the number of goroutines/tasks that the `WaitGroup` is waiting for.

Example:

```go
wg.Add(3)
```

means:

> "I am waiting for three tasks."

You can also use:

```go
wg.Add(1)
```

before starting each goroutine.

## `Done()`

Decreases the counter by one.

Usually:

```go
defer wg.Done()
```

is placed inside the goroutine.

## `Wait()`

Blocks until the counter becomes zero.

```go
wg.Wait()
```

### Example

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Worker", id, "started")

	time.Sleep(time.Second)

	fmt.Println("Worker", id, "finished")
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()

	fmt.Println("All workers finished")
}
```

### Execution concept

```text
main
 │
 ├── Worker 1
 ├── Worker 2
 ├── Worker 3
 ├── Worker 4
 └── Worker 5
       │
       ▼
    wg.Wait()
       │
       ▼
All workers finished
```

---

# 4. `sync.Once`

`sync.Once` guarantees that a particular operation is executed **at most once**, even if multiple goroutines attempt to execute it concurrently.

```go
var once sync.Once
```

The important method is:

```go
once.Do(function)
```

## `Do(f func())`

Runs `f` exactly once.

```go
once.Do(func() {
	fmt.Println("Initialize")
})
```

Even if 100 goroutines execute:

```go
once.Do(initialize)
```

`initialize()` will execute only once.

### Example

```go
package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func initialize() {
	fmt.Println("Initializing application...")
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			once.Do(initialize)
		}()
	}

	wg.Wait()
}
```

Output:

```text
Initializing application...
```

Only once.

---

# 5. `sync.Cond`

`sync.Cond` is a **condition variable**.

It allows goroutines to:

1. Wait until some condition becomes true.
2. Signal another waiting goroutine.
3. Signal all waiting goroutines.

It is less commonly needed in everyday Go code because channels often provide a simpler design.

A condition variable is associated with a locker:

```go
cond := sync.NewCond(&mu)
```

---

## `sync.NewCond(l Locker)`

Creates a new condition variable.

The locker normally is:

```go
&sync.Mutex{}
```

or:

```go
&sync.RWMutex{}
```

Example:

```go
mu := sync.Mutex{}
cond := sync.NewCond(&mu)
```

## `Cond.Wait()`

`Wait()`:

1. Releases the associated lock.
2. Puts the goroutine to sleep.
3. Waits for a signal.
4. Re-acquires the lock before returning.

Typical pattern:

```go
cond.L.Lock()

for !condition {
	cond.Wait()
}

cond.L.Unlock()
```

Notice the **`for`**, not `if`.

## `Cond.Signal()`

Wakes **one** waiting goroutine.

```go
cond.Signal()
```

Useful when one waiting goroutine should continue.

## `Cond.Broadcast()`

Wakes **all** waiting goroutines.

```go
cond.Broadcast()
```

Useful when a state change could allow multiple waiting goroutines to continue.

### Simple example

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)

	ready := false

	go func() {
		cond.L.Lock()

		for !ready {
			cond.Wait()
		}

		fmt.Println("Condition became true")

		cond.L.Unlock()
	}()

	mu.Lock()

	ready = true

	cond.Signal()

	mu.Unlock()
}
```

Conceptually:

```text
Worker
  │
  │ Wait()
  ▼
Sleeping
  │
  │ Signal()
  ▼
Continue
```

---

# 6. `sync.Map`

`sync.Map` is a concurrent map.

Normal Go maps are **not safe for concurrent access** when goroutines are modifying them.

You could protect a normal map with a mutex:

```go
var mu sync.Mutex
var data map[string]string
```

Or, for certain workloads, use:

```go
var data sync.Map
```

## `Map.Store(key, value)`

Stores a value.

```go
data.Store("name", "Chandu")
```

## `Map.Load(key)`

Loads a value.

```go
value, ok := data.Load("name")
```

`ok` tells you whether the key exists.

## `Map.LoadOrStore(key, value)`

Loads the existing value.

If the key doesn't exist, it stores the supplied value.

```go
actual, loaded := data.LoadOrStore("name", "Chandu")
```

If `loaded == true`, another value was already stored.

If `loaded == false`, your value was stored.

This is useful for concurrent initialization.

## `Map.LoadAndDelete(key)`

Loads a value and removes the key atomically.

```go
value, loaded := data.LoadAndDelete("name")
```

## `Map.Delete(key)`

Deletes a key.

```go
data.Delete("name")
```

## `Map.Swap(key, value)`

Atomically replaces the value associated with a key and returns the previous value.

```go
old, loaded := data.Swap("name", "Bob")
```

## `Map.CompareAndSwap(key, old, new)`

Atomically replaces `old` with `new` **only if the current value equals `old`**.

Conceptually:

```text
Current value = "Alice"

CompareAndSwap(
    "name",
    "Alice",
    "Bob",
)

Result:

"Alice" → "Bob"
```

If the current value isn't `"Alice"`, the replacement doesn't happen.

This is useful for atomic state transitions.

## `Map.CompareAndDelete(key, old)`

Deletes the key only if its current value equals the supplied value.

```go
data.CompareAndDelete("name", "Alice")
```

This is useful when you don't want to accidentally delete a value that another goroutine has already changed.

## `Map.Range(f func(key, value any) bool)`

Iterates over entries.

```go
data.Range(func(key, value any) bool {
	fmt.Println(key, value)
	return true
})
```

Returning:

```go
true
```

continues iteration.

Returning:

```go
false
```

stops iteration.

### Important `sync.Map` misconception

Don't automatically replace every:

```go
map + sync.Mutex
```

with:

```go
sync.Map
```

`sync.Map` is optimized for particular concurrent workloads. A normal map protected by a mutex is often easier to understand and can be the better choice.

---

# 7. `sync.Pool`

`sync.Pool` provides a mechanism for **temporarily reusing allocated objects**.

This can reduce allocations and garbage-collection pressure in some high-throughput programs.

Example:

```go
var pool sync.Pool

pool.New = func() any {
	return new(MyObject)
}
```

## `Pool.Get()`

Retrieves an object from the pool.

```go
obj := pool.Get()
```

If the pool doesn't have an available object and `New` is set, `New()` may be called.

## `Pool.Put(x)`

Returns an object to the pool.

```go
pool.Put(obj)
```

### Example

```go
package main

import (
	"fmt"
	"sync"
)

type Buffer struct {
	data []byte
}

var bufferPool = sync.Pool{
	New: func() any {
		return &Buffer{
			data: make([]byte, 0, 1024),
		}
	},
}

func main() {
	buf := bufferPool.Get().(*Buffer)

	buf.data = append(buf.data, "hello"...)

	fmt.Println(string(buf.data))

	buf.data = buf.data[:0]

	bufferPool.Put(buf)
}
```

### Critical rule

Objects in a `sync.Pool` are **temporary**.

The runtime may remove pooled objects during garbage collection.

Therefore, don't use `sync.Pool` as a persistent cache.

---

# 8. `sync.OnceFunc`

`OnceFunc` creates a function that executes its underlying function at most once.

```go
f := sync.OnceFunc(func() {
	fmt.Println("Executed")
})
```

Then:

```go
f()
f()
f()
```

The underlying function executes only once.

It is essentially a convenient modern alternative to manually creating a `sync.Once` around a function.

---

# 9. `sync.OnceValue`

`OnceValue` is useful when you want to execute a function once and **cache its returned value**.

Example:

```go
getConfig := sync.OnceValue(func() Config {
	return loadConfig()
})
```

Then:

```go
config1 := getConfig()
config2 := getConfig()
```

The initialization function executes once, and subsequent calls receive the same computed value.

This is particularly useful for **lazy initialization**.

---

# 10. `sync.OnceValues`

`OnceValues` is similar to `OnceValue`, but supports a function returning **two values**.

For example:

```go
load := sync.OnceValues(func() (*Config, error) {
	return loadConfig()
})
```

Then:

```go
config, err := load()
```

The function executes only once, and its returned values are reused on subsequent calls.

This is particularly convenient for lazy initialization where you need:

```go
(value, error)
```

---

# 11. `sync.Locker`

`Locker` is an interface used by synchronization primitives.

Its methods are:

```go
type Locker interface {
	Lock()
	Unlock()
}
```

Both:

```go
sync.Mutex
```

and:

```go
sync.RWMutex
```

can satisfy this interface.

This is why you can write:

```go
sync.NewCond(&mu)
```

where `mu` is a mutex.

---

# 12. `sync.Mutex` vs `sync.RWMutex`

A common beginner question is:

> "Should I always use `RWMutex` because reading is faster?"

**No.**

Use `Mutex` when you simply need exclusive access.

Use `RWMutex` when your workload genuinely benefits from allowing concurrent readers.

For example:

```text
Mutex

Reader 1 ──┐
Reader 2 ──┤
Reader 3 ──┤ → one at a time
Writer  ───┘
```

Whereas:

```text
RWMutex

Reader 1 ──┐
Reader 2 ──┤
Reader 3 ──┤ → simultaneously
Reader 4 ──┘

Writer → exclusive
```

Don't use `RWMutex` merely because your code has reads and writes.

---

# Three Common Beginner Mistakes

## Mistake 1: Forgetting `Unlock()`

Bad:

```go
mu.Lock()

if something {
	return
}

mu.Unlock()
```

If `something` is true, the mutex remains locked.

Prefer:

```go
mu.Lock()
defer mu.Unlock()

if something {
	return
}
```

---

## Mistake 2: Calling `WaitGroup.Add()` too late

Potentially problematic:

```go
go func() {
	wg.Add(1)
	defer wg.Done()

	// work
}()

wg.Wait()
```

The main goroutine may reach `Wait()` before the goroutine increments the counter.

Prefer:

```go
wg.Add(1)

go func() {
	defer wg.Done()

	// work
}()

wg.Wait()
```

The counter should generally be incremented **before starting the goroutine**.

---

## Mistake 3: Thinking `sync` automatically makes everything safe

For example:

```go
var mu sync.Mutex
var counter int

func increment() {
	mu.Lock()
	counter++
	mu.Unlock()
}
```

This protects `counter` **only when every access to `counter` follows the same synchronization discipline**.

If somewhere else you do:

```go
fmt.Println(counter)
```

without appropriate synchronization while another goroutine modifies it, you can still have a race.

Synchronization must be applied consistently to the shared state.

---

# Two Real-World Applications

## Application 1: Concurrent HTTP server

Imagine an HTTP server receiving thousands of requests.

Each request runs in its own goroutine, and all requests need to update:

```text
Total requests
Active users
Metrics
Counters
Shared configuration
```

A `sync.Mutex` or `sync.RWMutex` can protect those shared structures.

For example:

```text
10,000 requests
       │
       ├── goroutine 1 ──┐
       ├── goroutine 2 ──┤
       ├── goroutine 3 ──┤
       │                 │
       │             Mutex
       │                 │
       └── goroutine N ──┘
                         │
                     shared data
```

---

## Application 2: Worker pool

Suppose you have:

```text
1000 jobs
```

and:

```text
10 worker goroutines
```

You might use:

- `WaitGroup` → wait until workers finish
- `Mutex` → protect shared counters
- `Cond` → coordinate condition-based waiting
- `sync.Pool` → reuse temporary buffers

For example:

```text
                 Jobs
                  │
        ┌─────────┼─────────┐
        ▼         ▼         ▼
     Worker 1  Worker 2  Worker 3
        │         │         │
        └─────────┼─────────┘
                  ▼
             Shared Stats
                  │
               Mutex
```

---

# Progressive Exercises

## Exercise 1 — Beginner: Concurrent Counter

Create a program that starts **100 goroutines**.

Each goroutine should increment a shared counter **100 times**.

Requirements:

- Use `sync.WaitGroup`.
- Use `sync.Mutex` to protect the counter.
- Wait for all goroutines to finish.
- Print the final counter.
- The expected result should be `10,000`.

**Do not use `sync/atomic`.**

---

## Exercise 2 — Intermediate: Thread-Safe In-Memory Cache

Create a thread-safe cache:

```go
type Cache struct {
	// design this yourself
}
```

The cache should support:

```text
Set(key, value)
Get(key)
Delete(key)
```

Requirements:

- Multiple goroutines should be able to read concurrently.
- Writes must be protected.
- Use `sync.RWMutex`.
- Start multiple reader and writer goroutines.
- Use `sync.WaitGroup` to wait for them.
- Run the program with Go's race detector and make sure your design doesn't produce data races.

---

## Exercise 3 — Advanced: Concurrent Web Crawler

Build a small concurrent web crawler.

The crawler receives a starting URL and recursively discovers links.

Requirements:

- Multiple URLs should be processed concurrently.
- Don't process the same URL more than once.
- Maintain a shared set of visited URLs.
- Limit the number of concurrent workers.
- Wait until all work is completed.
- Protect shared state appropriately.
- Consider whether `sync.Mutex`, `sync.RWMutex`, `sync.Map`, `sync.Once`, or `WaitGroup` is appropriate for each part.
- Run the application with the race detector.
- Think carefully about what happens when one discovered URL generates several additional URLs.

**Do not use a third-party concurrency library.**

---

# A Useful Mental Model

When learning `sync`, think about the problem you're trying to solve:

```text
                What do I need?
                       │
        ┌──────────────┼──────────────┐
        │              │              │
        ▼              ▼              ▼
 Protect data      Wait for work    Run once
        │              │              │
        ▼              ▼              ▼
     Mutex        WaitGroup         Once
        │
        ▼
 Many readers?
        │
       Yes
        │
        ▼
    RWMutex
```

And:

```text
Need a concurrent map?
        │
        ▼
     sync.Map
```

```text
Need temporary object reuse?
        │
        ▼
     sync.Pool
```

```text
Need to wait for a condition?
        │
        ▼
     sync.Cond
```

```text
Need lazy one-time initialization?
        │
        ├── no return value → OnceFunc
        ├── one value       → OnceValue
        └── two values      → OnceValues
```

---

# Thought-Provoking Question

Imagine you have **10,000 HTTP requests arriving at the same time**, and every request needs to increment a shared counter.

You could protect the counter with a `sync.Mutex`, but that means goroutines may have to wait for one another.

**How would you redesign the architecture so that you don't need to put a mutex around every request's counter update—and what trade-offs would your design introduce?**

That question gets to the deeper Go concurrency principle:

> **Is it better to protect shared memory, or to design the system so that less memory needs to be shared in the first place?**
