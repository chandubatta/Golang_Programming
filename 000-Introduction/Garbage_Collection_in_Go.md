# Garbage Collection in Go

## 1. What is Garbage Collection?

**Garbage Collection (GC)** is the automatic process of finding and removing memory that a Go program no longer needs.

When a program creates objects, variables, slices, maps, structs, and other values, they may require memory. When those objects are no longer reachable, Go's **Garbage Collector** can reclaim that memory so it can be reused.

### Purpose of Garbage Collection

The main purposes are:

- Automatically reclaim unused memory.
- Reduce the risk of memory-management problems caused by forgotten deallocation.
- Make memory management easier for developers.
- Allow developers to focus more on application logic rather than manually freeing memory.

Go uses a **concurrent, tri-color, mark-and-sweep garbage collector** designed to keep GC pauses relatively small.

### When is it commonly used?

Garbage collection happens automatically while a Go program is running. You normally **do not manually perform GC**.

It is especially important in applications that continuously allocate memory, such as:

- Web servers
- REST APIs
- Microservices
- Network applications
- Cloud applications
- Large data-processing systems

---

# 2. Simple Garbage Collection Example

```go
package main

import (
	"fmt"
	"runtime"
)

type User struct {
	Name string
	Age  int
}

func main() {

	var m runtime.MemStats

	// Allocate memory
	for i := 0; i < 100000; i++ {
		user := &User{
			Name: "Chandu",
			Age:  25,
		}

		_ = user
	}

	// Memory statistics before GC
	runtime.ReadMemStats(&m)
	fmt.Println("Before GC:")
	fmt.Println("Allocated memory:", m.Alloc)

	// Request garbage collection
	runtime.GC()

	// Memory statistics after GC
	runtime.ReadMemStats(&m)
	fmt.Println("After GC:")
	fmt.Println("Allocated memory:", m.Alloc)
}
```

### What happens here?

A `User` object is created during each loop iteration:

```go
user := &User{
	Name: "Chandu",
	Age:  25,
}
```

Many objects are allocated during the loop.

Once an object is no longer reachable, it becomes **eligible for garbage collection**.

We then call:

```go
runtime.GC()
```

This **requests** that Go perform a garbage collection cycle.

> **Important:** You normally should **not** call `runtime.GC()` in production code just to "clean memory." Go's runtime automatically decides when GC should run.

---

# Basic GC Flow

```text
Program creates objects
        ↓
Memory is allocated
        ↓
Objects are used
        ↓
Objects become unreachable
        ↓
Garbage Collector identifies them
        ↓
Unused memory is reclaimed
        ↓
Memory can be reused
```

### Example

```go
func createUser() {
	user := &User{
		Name: "Chandu",
	}

	fmt.Println(user.Name)
}
```

When `createUser()` finishes, if nothing else references `user`, that object may eventually become eligible for garbage collection.

---

# 3. Three Common Mistakes and Misconceptions

## Mistake 1: Thinking GC immediately deletes unused objects

Beginners sometimes think:

> "As soon as an object becomes unused, Go immediately deletes it."

That's not how it works.

An object can become unreachable, but the garbage collector may run later.

### Avoid it

Think:

```text
Object becomes unreachable
        ↓
Eligible for GC
        ↓
GC runs later
        ↓
Memory is reclaimed
```

---

## Mistake 2: Calling `runtime.GC()` everywhere

A beginner might write:

```go
runtime.GC()
```

after every large allocation.

This is usually unnecessary and can negatively affect performance.

### Avoid it

Let Go's runtime manage garbage collection automatically.

Use `runtime.GC()` mainly for specialized situations such as experiments, testing, or controlled runtime behavior—not as a normal memory-management technique.

---

## Mistake 3: Thinking GC means memory leaks are impossible

Garbage collection helps prevent many types of manual-memory-management problems, but **Go programs can still have memory leaks**.

For example, if your program accidentally keeps references to objects that it no longer logically needs:

```go
var cache []*User

func addUser() {
	user := &User{Name: "Chandu"}

	cache = append(cache, user)
}
```

If `cache` continuously grows and never removes old users, those users are still reachable.

Therefore, the GC considers them **in use**.

### Avoid it

Look for:

- Unbounded slices
- Growing maps
- Caches without expiration
- Goroutines that never terminate
- Objects accidentally kept alive by long-lived references

---

# 4. Real-World Applications

## Application 1: REST API / Web Server

Imagine a Go REST API receiving thousands of requests:

```text
Client
   ↓
HTTP Request
   ↓
Go REST API
   ↓
Create request objects
   ↓
Process request
   ↓
Send response
   ↓
Objects no longer needed
   ↓
GC reclaims memory
```

For example, an API might create temporary structs while processing JSON:

```go
type UserRequest struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Phone string `json:"phone"`
}
```

After processing the request, many temporary objects may become unreachable.

Go's GC helps manage that memory automatically.

---

## Application 2: Microservices

Consider a Go payment microservice:

```text
Payment Request
      ↓
Create request data
      ↓
Validate payment
      ↓
Call database
      ↓
Call another service
      ↓
Return response
      ↓
Temporary objects become unreachable
      ↓
Garbage Collector
      ↓
Memory reused
```

This is particularly useful because microservices may process **large numbers of short-lived requests**.

---

# 5. Progressive Exercises

## Exercise 1 — Beginner: Observe Garbage Collection

Create a Go program that:

1. Allocates many objects inside a loop.
2. Stores them temporarily.
3. Removes the references to those objects.
4. Uses `runtime.MemStats` to observe memory usage.
5. Calls `runtime.GC()`.
6. Compares the memory statistics before and after GC.

**Goal:** Understand the relationship between object references and garbage collection.

---

## Exercise 2 — Intermediate: Find the Memory Retention Problem

Create a program that maintains a large in-memory cache:

```go
type User struct {
	ID   int
	Name string
}
```

The program should:

1. Continuously add users to a slice or map.
2. Process users.
3. Attempt to remove users that are no longer needed.
4. Observe how memory behaves.
5. Identify what happens if references to old users are accidentally retained.

**Goal:** Understand why garbage collection cannot reclaim objects that are still reachable.

---

## Exercise 3 — Advanced: GC Behavior in a High-Load Program

Build a small Go HTTP server that:

1. Provides an endpoint such as `/users`.
2. Creates temporary objects for every request.
3. Generates many concurrent requests.
4. Records memory allocation statistics.
5. Records the number of garbage collections.
6. Uses goroutines to simulate high traffic.
7. Compares memory behavior with different allocation patterns.

Investigate:

- How allocation rate affects GC.
- How many GC cycles occur.
- What happens when objects have longer lifetimes.
- How excessive allocations affect performance.

**Goal:** Understand the relationship between **allocation, object lifetime, GC frequency, and application performance**.

---

# 🧠 Think Deeper

Suppose you are building a **high-traffic Go microservice** that processes 100,000 requests per second.

> **Would you simply rely on Go's Garbage Collector, or would you also redesign your code to create fewer temporary objects? Why?**

Think about the trade-off between:

- Developer simplicity
- Memory usage
- GC frequency
- Latency
- Application performance
