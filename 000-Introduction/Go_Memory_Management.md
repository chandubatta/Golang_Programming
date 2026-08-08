# Memory Management in Go

## 1. What is Memory Management?

**Memory Management** is the process of allocating, using, and releasing memory while a Go program is running.

Go provides **automatic memory management**, mainly through:

- **Stack** — stores local variables and function-call data.
- **Heap** — stores data that needs to live beyond a function call or cannot conveniently be kept on the stack.
- **Garbage Collector (GC)** — automatically identifies heap memory that is no longer reachable and reclaims it.
- **Escape analysis** — the Go compiler decides whether values can remain on the stack or need to escape to the heap.

### Purpose

Memory management helps your application:

- Use memory efficiently.
- Avoid unnecessary memory consumption.
- Prevent memory leaks caused by retaining objects unnecessarily.
- Maintain good application performance.
- Handle large amounts of data safely.

You commonly think about memory management when working with **large slices, maps, structs, files, network connections, caches, long-running services, and high-concurrency applications**.

---

## 2. Simple Code Example

Consider a function that creates a slice:

```go
package main

import "fmt"

func createNumbers() []int {
    numbers := make([]int, 5)

    for i := 0; i < len(numbers); i++ {
        numbers[i] = i + 1
    }

    return numbers
}

func main() {
    numbers := createNumbers()

    fmt.Println(numbers)
}
```

### What happens to memory?

```text
main()
   |
   | calls
   v
createNumbers()
   |
   |-- numbers slice created
   |
   |-- [1 2 3 4 5]
   |
   v
return numbers
   |
   v
main() receives numbers
```

The important point is that Go automatically determines how the underlying data should be stored.

The programmer normally **does not manually call `free()` or `delete()`** as in languages such as C/C++.

### Garbage collection example

```go
package main

import "fmt"

func createData() {
    data := make([]byte, 1024*1024) // 1 MB

    fmt.Println(len(data))
}

func main() {
    createData()

    // After createData returns, if data is no longer reachable,
    // the memory can eventually be reclaimed by Go's GC.
}
```

The garbage collector determines when unreachable heap objects can be reclaimed.

> **Important:** Garbage collection does not mean you can ignore memory usage. If your program continues holding references to objects, the GC cannot reclaim them.

---

# 3. Three Common Mistakes and Misconceptions

## Mistake 1: Thinking Go has no memory management

Some beginners think:

> "Go has garbage collection, so I don't need to care about memory."

That's incorrect.

You still need to avoid unnecessary allocations and unnecessary references.

For example:

```go
cache := make(map[string][]byte)
```

If you continuously add data to `cache` and never remove old entries, memory usage can keep increasing.

### How to avoid it

Understand:

- Stack vs heap
- Allocations
- Garbage collection
- References
- Slices and their backing arrays
- Maps
- Escape analysis

---

## Mistake 2: Thinking `runtime.GC()` should be called frequently

Go provides:

```go
runtime.GC()
```

Some beginners think:

```go
runtime.GC()
runtime.GC()
runtime.GC()
```

will automatically make their application faster.

Usually, manually forcing garbage collection is **not necessary**.

Go's garbage collector is designed to manage collection automatically.

### How to avoid it

Let Go's runtime normally manage GC. Investigate memory behavior using profiling tools before trying to control GC manually.

---

## Mistake 3: Accidentally retaining large memory

Consider:

```go
data := make([]byte, 100*1024*1024)

// Use data...

small := data[:10]
```

`small` contains only 10 bytes logically, but it can still keep the **large underlying array** reachable.

This can cause unexpectedly high memory usage.

### How to avoid it

When you need a small independent portion of a large object, consider copying the required data into a new smaller allocation.

---

# 4. Real-World Applications

## Application 1: High-traffic Web APIs

Imagine a Go REST API receiving thousands of requests.

Each request may create:

```text
Request
   ↓
JSON data
   ↓
Structs
   ↓
Slices / Maps
   ↓
Database results
   ↓
Response
```

Efficient memory usage becomes important because many requests may exist concurrently.

Poor memory management can lead to:

```text
More requests
     ↓
More allocations
     ↓
Higher memory usage
     ↓
More GC work
     ↓
Possible performance degradation
```

---

## Application 2: Large File/Data Processing

Suppose you process a **2 GB CSV file**.

A poor approach could attempt to load everything into memory:

```text
2 GB CSV
   ↓
Load entire file
   ↓
Very high memory usage
```

A better approach can process the file incrementally:

```text
CSV file
   ↓
Read chunk/line
   ↓
Process
   ↓
Discard
   ↓
Read next chunk
```

This can dramatically reduce the amount of memory required.

---

# 5. Practice Exercises

## Exercise 1 — Basic

Create a Go program that:

1. Creates a slice containing 1,000 integers.
2. Fills it with values.
3. Prints the first 10 values.
4. Passes the slice to a function.
5. Modifies some values inside the function.
6. Observes what happens to the original slice.

**Goal:** Understand how slices and their underlying memory behave when passed to functions.

---

## Exercise 2 — Intermediate

Create a program that processes a large collection of data.

Requirements:

1. Generate 1,000,000 integers.
2. Store them in a slice.
3. Calculate their sum.
4. Create another slice containing only the required values.
5. Observe the memory usage before and after the operations.
6. Use Go's memory statistics to investigate the behavior.

**Goal:** Understand allocations, heap usage, and how data structures affect memory consumption.

---

## Exercise 3 — Advanced

Build a small **in-memory cache**.

Requirements:

1. Create a cache using a Go map.
2. Store many key-value pairs.
3. Allow items to be retrieved.
4. Allow items to be removed.
5. Add a mechanism to limit the maximum number of cached items.
6. Investigate what happens to memory when the cache continuously grows.
7. Use Go memory profiling tools to investigate allocations.
8. Identify which parts of your program are responsible for the largest memory allocations.

**Goal:** Understand practical memory management in a long-running Go application.

---

# 🤔 Think Deeper

Suppose your Go API uses **2 GB of RAM** even though each individual request only needs a few megabytes.

**If the garbage collector is already automatic, what could be causing the application to keep so much memory—and how would you prove where the memory is being retained?**

Think about **slices, maps, goroutines, caches, references, heap allocations, and garbage collection** before answering.
