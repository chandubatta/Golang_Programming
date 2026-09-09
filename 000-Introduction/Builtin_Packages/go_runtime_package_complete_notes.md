# Go `runtime` Package

The Go `runtime` package is a low-level and important package because it provides access to information and controls related to Go's runtime system: goroutines, scheduling, CPUs, garbage collection, memory statistics, stack traces, profiling, OS threads, and more.

```go
import "runtime"
```

## 1. What is the `runtime` package?

The `runtime` package provides functions and types that interact directly with the Go runtime system.

The Go runtime is responsible for:

- Running goroutines
- Scheduling goroutines onto OS threads
- Managing memory
- Garbage collection
- Managing OS threads
- Providing stack information
- CPU/memory profiling
- Runtime tracing
- Getting information about the machine
- Controlling the number of CPUs used for parallel execution
- Working with cgo
- Getting information about the Go version

### When is `runtime` commonly used?

You generally don't need `runtime` for ordinary Go programs. It becomes useful when doing:

1. Performance tuning
2. Concurrency diagnostics
3. Memory analysis
4. CPU/memory profiling
5. Debugging
6. Runtime monitoring
7. Low-level system programming
8. Working with OS threads
9. cgo
10. Building developer/debugging tools

Think of it as:

```text
Your Go application
       ↓
Go standard library
       ↓
runtime package
       ↓
Go runtime
       ↓
OS / CPU / memory
```

---

# 2. Simple Example

This example shows the Go version, operating system, architecture, CPU count, and goroutine count.

```go
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Go version:", runtime.Version())
	fmt.Println("Operating system:", runtime.GOOS)
	fmt.Println("Architecture:", runtime.GOARCH)
	fmt.Println("Number of CPUs:", runtime.NumCPU())
	fmt.Println("Number of goroutines:", runtime.NumGoroutine())

	go func() {
		time.Sleep(2 * time.Second)
	}()

	fmt.Println("Goroutines after starting one:", runtime.NumGoroutine())
}
```

Possible output:

```text
Go version: go1.27.1
Operating system: windows
Architecture: amd64
Number of CPUs: 8
Number of goroutines: 1
Goroutines after starting one: 2
```

The exact output depends on your computer.

---

# 3. Important Functions in `runtime`

## 3.1 `BlockProfile`

```go
func BlockProfile(p []BlockProfileRecord) (n int, ok bool)
```

Returns information about goroutines that were blocked waiting for synchronization events.

For example:

```text
goroutine
   ↓
waiting for mutex
   ↓
blocked
```

It can help investigate synchronization bottlenecks.

For normal profiling, `runtime/pprof` or testing-package profiling facilities are generally preferable.

**Beginner takeaway:** `BlockProfile` helps investigate where goroutines are spending time blocked.

---

## 3.2 `Breakpoint`

```go
func Breakpoint()
```

Triggers a breakpoint trap.

Example:

```go
runtime.Breakpoint()
```

This is primarily useful for debugging/runtime-level work and is rarely needed in ordinary application code.

---

## 3.3 `CPUProfile` — Deprecated

```go
func CPUProfile() []byte
```

This function is deprecated. The old raw CPU profiling mechanism has been removed.

For new programs, use:

```text
runtime/pprof
```

or:

```text
net/http/pprof
```

Do not use `CPUProfile` in new application code.

---

## 3.4 `Caller`

```go
func Caller(skip int) (
	pc uintptr,
	file string,
	line int,
	ok bool,
)
```

Retrieves information about a caller in the call stack.

Example:

```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	printLocation()
}

func printLocation() {
	pc, file, line, ok := runtime.Caller(1)

	fmt.Println(pc)
	fmt.Println(file)
	fmt.Println(line)
	fmt.Println(ok)
}
```

Possible output resembles:

```text
4821234
C:/project/main.go
10
true
```

### What does `skip` mean?

`skip` tells Go how many stack frames to skip.

Conceptually:

```text
main()
   ↓
printLocation()
   ↓
runtime.Caller()
```

With:

```go
runtime.Caller(0)
```

you examine the current caller frame.

With:

```go
runtime.Caller(1)
```

you move one frame upward.

### Real-world use

Logging frameworks can use call-stack information to identify where a log message originated.

---

## 3.5 `Callers`

```go
func Callers(skip int, pc []uintptr) int
```

Retrieves the program counters of multiple stack frames.

Example:

```go
pcs := make([]uintptr, 10)

n := runtime.Callers(0, pcs)

fmt.Println("Frames:", n)
```

Unlike `Caller`, which returns information for one frame, `Callers` collects multiple frames.

It is commonly combined with:

```go
runtime.CallersFrames()
```

to turn program counters into useful file/function/line information.

---

## 3.6 `GC`

```go
func GC()
```

Requests that the garbage collector perform a garbage collection.

Example:

```go
runtime.GC()
```

Important: this does not mean:

> "The garbage collector will definitely instantly free all memory."

The runtime manages memory and garbage collection automatically.

You normally should not manually call `runtime.GC()` in ordinary application code.

### When might it be useful?

Testing and specialized performance experiments can sometimes benefit from explicitly triggering GC.

---

## 3.7 `GOMAXPROCS`

```go
func GOMAXPROCS(n int) int
```

Controls the maximum number of CPUs that can execute Go code simultaneously.

Example:

```go
old := runtime.GOMAXPROCS(2)

fmt.Println("Previous:", old)
```

You can query the current setting with:

```go
current := runtime.GOMAXPROCS(0)

fmt.Println(current)
```

Modern Go automatically chooses an appropriate default based on factors such as logical CPUs, CPU affinity, and container CPU limits.

### Important distinction

`GOMAXPROCS` does **not** mean maximum number of goroutines.

You can have:

```text
GOMAXPROCS = 4
Goroutines = 10,000
```

That is perfectly possible.

Think:

```text
10,000 goroutines
        ↓
scheduler
        ↓
up to 4 CPUs executing Go code simultaneously
```

---

## 3.8 `GOROOT` — Deprecated

```go
func GOROOT() string
```

Returns the root directory of the Go installation.

Example:

```go
fmt.Println(runtime.GOROOT())
```

The function is deprecated. For tooling, prefer mechanisms such as:

```bash
go env GOROOT
```

rather than building application logic around `runtime.GOROOT()`.

---

## 3.9 `Goexit`

```go
func Goexit()
```

Terminates the current goroutine.

Example:

```go
go func() {
	fmt.Println("Before")

	runtime.Goexit()

	fmt.Println("After")
}()
```

Output:

```text
Before
```

The second statement is not reached.

### Important difference from `return`

`return` returns from the current function.

`runtime.Goexit()` terminates the entire current goroutine.

Deferred functions still execute.

---

## 3.10 `GoroutineProfile`

```go
func GoroutineProfile(p []StackRecord) (n int, ok bool)
```

Gets stack information about active goroutines.

It can help diagnose goroutine behavior, such as:

```text
How many goroutines exist?
What are they doing?
Where are they blocked?
```

For ordinary profiling, `runtime/pprof` is generally preferred.

---

## 3.11 `Gosched`

```go
func Gosched()
```

Yields the processor so another goroutine can run.

Example:

```go
runtime.Gosched()
```

Conceptually:

```text
Goroutine A
     ↓
Gosched()
     ↓
"Let another goroutine run"
     ↓
Goroutine B
```

Important: `Gosched()` does not:

- Sleep for a specific duration
- Block until another goroutine finishes
- Guarantee that another goroutine will run

For ordinary synchronization, use channels, `sync`, or other appropriate mechanisms.

---

## 3.12 `KeepAlive`

```go
func KeepAlive(x any)
```

Ensures that a value is considered reachable until a particular point in the program.

This becomes important in low-level code involving:

- Pointers
- Finalizers
- System calls
- File descriptors
- cgo
- Unsafe operations

Conceptually:

```go
resource := acquireResource()

useResource(resource)

runtime.KeepAlive(resource)
```

The purpose is to prevent the compiler/GC from considering the object unreachable too early.

### Beginner warning

You usually don't need `KeepAlive`. It becomes important when writing low-level code where garbage collection and external resources interact.

---

## 3.13 `LockOSThread`

```go
func LockOSThread()
```

Locks the current goroutine to its current OS thread.

Typical pattern:

```go
runtime.LockOSThread()
defer runtime.UnlockOSThread()

// thread-sensitive operation
```

Conceptually:

```text
Goroutine
    │
    └──── permanently associated
             ↓
          OS thread
```

This is relevant to:

- GUI libraries
- Certain system APIs
- cgo
- Thread-local OS state

---

## 3.14 `MemProfile`

```go
func MemProfile(
	p []MemProfileRecord,
	inuseZero bool,
) (n int, ok bool)
```

Provides memory allocation profiling information.

It can help investigate:

```text
Where is memory being allocated?
Which allocation sites remain in use?
How many objects are involved?
```

For normal application profiling, `runtime/pprof` is generally easier and preferable.

---

## 3.15 `MutexProfile`

```go
func MutexProfile(
	p []BlockProfileRecord,
) (n int, ok bool)
```

Provides mutex contention profiling information.

This can help identify places where goroutines spend significant time waiting for mutexes.

Conceptually:

```text
Goroutine 1 ──┐
              ↓
           Mutex
              ↑
Goroutine 2 ──┘
              ↑
          contention
```

For profiling, `runtime/pprof` is generally preferred.

---

## 3.16 `NumCPU`

```go
func NumCPU() int
```

Returns the number of logical CPUs available to the Go process.

Example:

```go
fmt.Println(runtime.NumCPU())
```

Possible output:

```text
8
```

Useful for understanding available CPU resources.

### Important

`NumCPU()` and `GOMAXPROCS()` answer different questions:

```text
NumCPU()
    ↓
How many logical CPUs are available?

GOMAXPROCS(0)
    ↓
How many CPUs may execute Go code simultaneously?
```

---

## 3.17 `NumCgoCall`

```go
func NumCgoCall() int64
```

Returns the number of cgo calls made by the current process.

Useful when investigating applications that interact with C code through cgo.

For pure Go programs, this usually isn't very interesting.

---

## 3.18 `NumGoroutine`

```go
func NumGoroutine() int
```

Returns the number of currently existing goroutines.

Example:

```go
fmt.Println("Goroutines:", runtime.NumGoroutine())
```

Example:

```go
for i := 0; i < 100; i++ {
	go func() {
		time.Sleep(time.Second)
	}()
}

fmt.Println(runtime.NumGoroutine())
```

This can help you observe goroutine growth.

### Important

A high goroutine count is not automatically a problem. Go is designed to support many goroutines.

---

## 3.19 `ReadMemStats`

```go
func ReadMemStats(m *MemStats)
```

Reads detailed memory statistics into a `runtime.MemStats` structure.

Example:

```go
var stats runtime.MemStats

runtime.ReadMemStats(&stats)

fmt.Println("Allocated:", stats.Alloc)
fmt.Println("Total allocated:", stats.TotalAlloc)
fmt.Println("System memory:", stats.Sys)
fmt.Println("Number of mallocs:", stats.Mallocs)
fmt.Println("Number of frees:", stats.Frees)
```

Useful fields include:

```text
Alloc
TotalAlloc
Sys
Mallocs
Frees
HeapAlloc
HeapObjects
NumGC
```

Example:

```go
fmt.Println("GC cycles:", stats.NumGC)
```

---

## 3.20 `ReadTrace`

```go
func ReadTrace() []byte
```

Reads data produced by the Go execution tracer.

This is a low-level API.

For normal tracing, use Go's supported tracing tools rather than treating this as a beginner-level application API.

---

## 3.21 `SetBlockProfileRate`

```go
func SetBlockProfileRate(rate int)
```

Controls the fraction of blocking events that are recorded in the block profile.

Conceptually:

```text
rate = 0
    ↓
disable block profiling

rate > 0
    ↓
record blocking events
```

Useful when investigating synchronization/blocking behavior.

---

## 3.22 `SetCPUProfileRate`

```go
func SetCPUProfileRate(hz int)
```

Sets the CPU profiling sampling rate.

For normal CPU profiling, you should generally use:

```text
runtime/pprof
```

rather than directly manipulating the runtime's CPU profiler.

---

## 3.23 `SetCgoTraceback`

```go
func SetCgoTraceback(
	version int,
	traceback unsafe.Pointer,
	context unsafe.Pointer,
	symbolizer unsafe.Pointer,
)
```

An advanced function used to configure traceback support for cgo code.

It allows the runtime to obtain stack traces from C/C++ code participating in a Go program.

### Beginner recommendation

Do not use this until you are comfortable with:

- cgo
- pointers
- `unsafe`
- C stack frames
- Go runtime internals

---

## 3.24 `SetDefaultGOMAXPROCS`

```go
func SetDefaultGOMAXPROCS()
```

Restores the runtime's automatic/default `GOMAXPROCS` behavior.

This is particularly relevant because modern Go can automatically determine an appropriate value based on CPU availability, affinity, and container CPU limits.

---

## 3.25 `SetFinalizer`

```go
func SetFinalizer(obj any, finalizer any)
```

Associates a finalizer with an object.

A finalizer is a function that may be called when an object becomes unreachable and is about to be collected.

Conceptually:

```text
object becomes unreachable
          ↓
GC notices it
          ↓
finalizer may execute
```

Example:

```go
type Resource struct {
	name string
}

func cleanup(r *Resource) {
	fmt.Println("Cleaning:", r.name)
}

r := &Resource{name: "resource-1"}

runtime.SetFinalizer(r, cleanup)
```

### Important warning

Do not treat finalizers as deterministic cleanup.

You cannot reliably say:

```text
"I created the object at 10:00,
therefore finalizer runs at 10:01."
```

For resources such as files, sockets, database connections, and locks, explicit cleanup is normally much better:

```go
defer file.Close()
```

---

## 3.26 `SetMutexProfileFraction`

```go
func SetMutexProfileFraction(rate int) int
```

Controls mutex contention profiling.

It returns the previous setting.

Useful when investigating mutex contention and synchronization performance.

---

## 3.27 `Stack`

```go
func Stack(buf []byte, all bool) int
```

Writes stack information into a byte buffer.

Example:

```go
buf := make([]byte, 4096)

n := runtime.Stack(buf, false)

fmt.Println(string(buf[:n]))
```

Output may resemble:

```text
goroutine 1 [running]:
main.main()
    /project/main.go:10 +0x...
```

If:

```go
all = true
```

the runtime includes stack information for all goroutines.

This is useful for debugging goroutine behavior.

---

## 3.28 `StartTrace`

```go
func StartTrace() error
```

Starts execution tracing.

Example:

```go
err := runtime.StartTrace()

if err != nil {
	fmt.Println(err)
}
```

Tracing provides information about runtime events such as:

- Goroutine scheduling
- Blocking
- System calls
- Garbage collection
- Processor activity

For modern applications, Go's tracing tooling is generally preferred over manually handling low-level runtime trace APIs.

---

## 3.29 `StopTrace`

```go
func StopTrace()
```

Stops an execution trace previously started with:

```go
runtime.StartTrace()
```

Typical pattern:

```go
runtime.StartTrace()

// program activity

runtime.StopTrace()
```

---

## 3.30 `ThreadCreateProfile`

```go
func ThreadCreateProfile(
	p []StackRecord,
) (n int, ok bool)
```

Provides information about OS threads created by the Go runtime.

Useful when diagnosing unusual thread creation or cgo/system-level behavior.

Most application developers will rarely need it.

---

## 3.31 `UnlockOSThread`

```go
func UnlockOSThread()
```

Releases the current goroutine from its locked OS thread.

It normally pairs with:

```go
runtime.LockOSThread()
```

Typical pattern:

```go
runtime.LockOSThread()
defer runtime.UnlockOSThread()

// thread-sensitive operation
```

---

## 3.32 `Version`

```go
func Version() string
```

Returns the Go version used to build the running program.

Example:

```go
fmt.Println(runtime.Version())
```

Possible output:

```text
go1.27.1
```

Useful for diagnostics and runtime reporting.

---

# 4. Important `runtime` Types

Although the focus is on functions, several exported types make the functions easier to understand.

The package exposes types such as:

- `MemStats`
- `StackRecord`
- `Frame`
- `Frames`
- `Func`
- `Pinner`
- `Cleanup`
- `BlockProfileRecord`

## `MemStats`

Used with:

```go
runtime.ReadMemStats()
```

Contains detailed memory and garbage-collection statistics.

## `StackRecord`

Represents stack information.

Used by functions such as:

```go
runtime.GoroutineProfile()
runtime.ThreadCreateProfile()
```

## `BlockProfileRecord`

Contains information about blocking events.

Used by:

```go
runtime.BlockProfile()
runtime.MutexProfile()
```

## `Func`

Represents information about a Go function.

You can obtain one using:

```go
runtime.FuncForPC(pc)
```

Then:

```go
f.Name()
```

can provide the function name.

## `Frames`

Used for turning program counters into readable stack frames.

Typical flow:

```text
runtime.Callers()
       ↓
[]uintptr
       ↓
runtime.CallersFrames()
       ↓
Frames
       ↓
Next()
       ↓
Frame
```

---

# 5. More Practical Example

This example combines several runtime functions.

```go
package main

import (
	"fmt"
	"runtime"
	"time"
)

func worker() {
	time.Sleep(2 * time.Second)
}

func main() {
	fmt.Println("Go version:", runtime.Version())
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("Architecture:", runtime.GOARCH)
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	for i := 0; i < 5; i++ {
		go worker()
	}

	fmt.Println("After creating goroutines:")
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)

	fmt.Println("Heap allocation:", stats.HeapAlloc)
	fmt.Println("Heap objects:", stats.HeapObjects)
	fmt.Println("GC cycles:", stats.NumGC)

	time.Sleep(3 * time.Second)

	fmt.Println("Final goroutine count:",
		runtime.NumGoroutine())
}
```

This demonstrates:

```text
runtime.Version()
       ↓
Go version

runtime.GOOS
       ↓
Operating system

runtime.GOARCH
       ↓
CPU architecture

runtime.NumCPU()
       ↓
logical CPUs

runtime.GOMAXPROCS(0)
       ↓
current execution parallelism

runtime.NumGoroutine()
       ↓
goroutine count

runtime.ReadMemStats()
       ↓
memory statistics
```

---

# 6. Three Common Beginner Mistakes

## Mistake 1: Thinking `runtime.NumCPU()` equals the number of goroutines

Wrong:

```text
8 CPUs → maximum 8 goroutines
```

Correct:

```text
8 CPUs
+
thousands of goroutines
```

is completely normal.

`NumCPU()` reports logical CPU availability, while goroutines are lightweight concurrent units managed by the Go runtime.

---

## Mistake 2: Calling `runtime.GC()` everywhere

A beginner may think:

```go
runtime.GC()
```

means:

> "Free my unused memory immediately."

That's not how Go's garbage collector should normally be managed.

The Go runtime automatically performs garbage collection.

Do not add:

```go
runtime.GC()
```

randomly to production code in an attempt to improve performance.

Use profiling to determine whether garbage collection is actually a problem.

---

## Mistake 3: Using `runtime.Gosched()` for synchronization

A beginner might write:

```go
runtime.Gosched()
```

thinking:

> "Wait until the other goroutine finishes."

That is incorrect.

`Gosched()` simply gives the scheduler an opportunity to run another goroutine.

It does not establish synchronization.

For synchronization, use tools such as:

```go
sync.WaitGroup
```

or:

```text
channels
```

or:

```go
sync.Mutex
```

depending on the problem.

---

# 7. Two Real-World Applications

## Application 1: Production Performance Monitoring

Imagine you are running a Go web service.

You observe:

```text
CPU usage ↑
Memory usage ↑
Latency ↑
```

You can investigate runtime behavior using:

```text
runtime.NumGoroutine()
runtime.ReadMemStats()
runtime.GOMAXPROCS()
runtime/pprof
runtime/trace
```

For example, a continuously increasing goroutine count could indicate a **goroutine leak**.

```text
Requests
   ↓
goroutines created
   ↓
goroutines never finish
   ↓
goroutine count ↑
   ↓
memory/resource usage ↑
```

This is a practical use of runtime diagnostics.

---

## Application 2: High-Performance Server / Container Tuning

Suppose your Go application runs inside a container with limited CPU resources.

Modern Go automatically considers CPU availability, CPU affinity, and container CPU limits when choosing the default `GOMAXPROCS`.

You can inspect:

```go
runtime.NumCPU()
```

and:

```go
runtime.GOMAXPROCS(0)
```

to understand how your program is configured.

This is useful when tuning:

- Web servers
- Microservices
- High-throughput APIs
- Worker systems
- Containerized applications

---

# 8. Three Progressive Exercises

## Exercise 1 — Beginner: Runtime Information

Write a Go program that prints:

1. Go version
2. Operating system
3. CPU architecture
4. Number of logical CPUs
5. Current `GOMAXPROCS`
6. Number of goroutines

### Requirements

Use the appropriate functions/constants from the `runtime` package.

Do not use external packages for obtaining this information.

---

## Exercise 2 — Intermediate: Goroutine Monitor

Create a program that starts **100 goroutines**.

Each goroutine should:

1. Print or record that it started.
2. Sleep for a short amount of time.
3. Finish normally.

Your main function should periodically report:

```text
Current goroutines: X
```

using the `runtime` package.

### Your goal

Observe how the goroutine count changes:

```text
Before starting goroutines
        ↓
While goroutines are running
        ↓
After goroutines finish
```

Try to explain why the number changes.

---

## Exercise 3 — Advanced: Runtime Diagnostic Tool

Build a small command-line diagnostic program that periodically reports:

- Go version
- OS
- Architecture
- CPU count
- Current `GOMAXPROCS`
- Current goroutine count
- Heap allocation
- Heap object count
- Total allocations
- Total frees
- Number of GC cycles

Use:

```text
runtime.NumCPU()
runtime.GOMAXPROCS()
runtime.NumGoroutine()
runtime.ReadMemStats()
runtime.Version()
```

### Extra challenge

Add a mechanism that allows you to create a large number of goroutines and observe how:

```text
goroutine count
memory allocation
heap objects
GC cycles
```

change over time.

Do **not** manually call `runtime.GC()` as part of the experiment initially. First observe the runtime's natural behavior.

---

# 9. `runtime` vs `runtime/debug` vs `runtime/pprof`

This distinction is useful:

| Package | Main purpose |
|---|---|
| `runtime` | Low-level interaction with the Go runtime |
| `runtime/debug` | Runtime debugging and GC/memory controls |
| `runtime/pprof` | CPU, memory, goroutine, mutex and block profiling |
| `runtime/trace` | Execution tracing |
| `runtime/cgo` | Runtime support for cgo |

For example, don't automatically reach for:

```go
runtime.ReadMemStats()
```

just because you are interested in performance.

For serious performance investigations, `runtime/pprof` is often the more appropriate tool.

---

# 10. Important Modern-Go Note

The runtime evolves between Go releases.

For example, Go 1.26 introduced the **Green Tea garbage collector as the default**, after it had been experimental in Go 1.25.

Therefore, avoid assuming that an implementation detail of the runtime will remain identical forever.

The best mindset is:

> **Use the documented runtime API, not assumptions about how the runtime happens to be implemented internally.**

This is particularly important with:

- Garbage collection
- Scheduling
- `GOMAXPROCS`
- Memory management
- Stack behavior
- Profiling internals

---

# 11. A Useful Mental Model

Organize the runtime functions into categories rather than trying to memorize them individually.

## Goroutines

```text
NumGoroutine()
Goexit()
Gosched()
GoroutineProfile()
Stack()
```

## CPUs / Scheduling

```text
NumCPU()
GOMAXPROCS()
SetDefaultGOMAXPROCS()
Gosched()
```

## Memory / Garbage Collection

```text
GC()
ReadMemStats()
MemProfile()
SetFinalizer()
KeepAlive()
```

## OS Threads

```text
LockOSThread()
UnlockOSThread()
ThreadCreateProfile()
```

## Stack / Debugging

```text
Caller()
Callers()
CallersFrames()
Stack()
FuncForPC()
```

## Profiling

```text
BlockProfile()
MemProfile()
MutexProfile()
SetBlockProfileRate()
SetCPUProfileRate()
SetMutexProfileFraction()
```

## Tracing

```text
StartTrace()
StopTrace()
ReadTrace()
```

## Runtime Information

```text
Version()
NumCPU()
GOOS
GOARCH
```

---

# 12. The Big Picture

Think about Go execution like this:

```text
                    Go Program
                        │
              ┌─────────┴─────────┐
              │                   │
          Goroutines          Memory
              │                   │
              ▼                   ▼
         Scheduler               GC
              │                   │
              └─────────┬─────────┘
                        │
                    Go Runtime
                        │
        ┌───────────────┼────────────────┐
        │               │                │
       CPU           OS Threads       Memory
        │               │                │
        └───────────────┼────────────────┘
                        │
                       OS
```

The `runtime` package gives you a window into this machinery.

You normally write Go code at a higher level, but when you need to understand why a program is behaving a certain way, `runtime` can expose what is happening underneath.

The package provides exported functions for profiling, goroutines, memory, CPU, tracing, OS-thread, and diagnostic operations, along with related types such as `MemStats`, `Frames`, `Func`, `Pinner`, and `Cleanup`.

---

# 13. Thought-Provoking Question 🤔

Suppose you have a Go web server with **50,000 goroutines**, but only **8 logical CPUs**. The server is experiencing high latency.

**Would you immediately increase `GOMAXPROCS` to 50,000? Why or why not?**

Think carefully about the difference between:

- Goroutines
- OS threads
- CPUs
- Scheduling
- Blocking
- Actual parallelism

The goal is to understand why **50,000 goroutines and 8 CPUs can coexist normally**, and why increasing `GOMAXPROCS` blindly may not solve a latency problem.
