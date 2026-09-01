# Go `log` Package

The standard-library `log` package provides a simple logging mechanism for Go programs. It is commonly used to record information about what an application is doing, especially while developing, debugging, diagnosing failures, or observing a running service.

The `log` package provides a predefined standard `Logger`, plus the ability to create custom `Logger` instances.

> **Important:** Go also has the newer `log/slog` package for structured logging. The traditional `log` package is still useful when you need straightforward text-based logging or compatibility with existing code.

## 1. What is the `log` package?

Import it with:

```go
import "log"
```

The `log` package helps you write messages such as:

```text
2026/09/01 13:20:15 Server started
2026/09/01 13:20:18 User logged in
2026/09/01 13:20:22 Database connection failed
```

By default, the standard logger writes to standard error (`stderr`) and includes the date and time.

### Why use logging?

Instead of scattering `fmt.Println()` throughout an application:

```go
fmt.Println("Server started")
fmt.Println("Connecting to database")
fmt.Println("Request received")
```

you can use:

```go
log.Println("Server started")
log.Println("Connecting to database")
log.Println("Request received")
```

The logger automatically adds useful information such as timestamps.

### When is `log` commonly used?

- Debugging applications
- Recording application events
- Reporting errors
- Monitoring servers
- Tracking program execution
- Diagnosing production problems
- Writing logs to files
- Creating multiple loggers for different components

---

# 2. Simple Example

```go
package main

import "log"

func main() {
	log.Println("Application started")

	username := "Chandu"

	log.Printf("User %s logged in", username)

	log.Println("Application finished")
}
```

Possible output:

```text
2026/09/01 13:20:15 Application started
2026/09/01 13:20:15 User Chandu logged in
2026/09/01 13:20:15 Application finished
```

The exact timestamp depends on when the program runs.

---

# 3. Understanding the `log` Package

The package has two main ways of working.

## Package-level functions

You can directly write:

```go
log.Println("Hello")
```

These use Go's standard logger.

## Custom `Logger`

You can create your own:

```go
logger := log.New(output, "APP: ", log.LstdFlags)
```

Then:

```go
logger.Println("Hello")
```

This is useful when you need different destinations, prefixes, or configurations.

---

# 4. Constants / Log Flags

The `log` package provides flags that control what appears before each message.

## `log.Ldate`

Displays the date.

```go
log.SetFlags(log.Ldate)
```

Example:

```text
2026/09/01 Application started
```

---

## `log.Ltime`

Displays the time.

```go
log.SetFlags(log.Ltime)
```

Example:

```text
13:20:15 Application started
```

---

## `log.Lmicroseconds`

Adds microsecond precision to the time.

```go
log.SetFlags(log.Ltime | log.Lmicroseconds)
```

Example:

```text
13:20:15.123456 Application started
```

It assumes `Ltime` is also enabled.

---

## `log.Llongfile`

Displays the complete source-file path and line number.

```go
log.SetFlags(log.Llongfile)
```

Example:

```text
C:/projects/myapp/main.go:15 Application started
```

---

## `log.Lshortfile`

Displays only the file name and line number.

```go
log.SetFlags(log.Lshortfile)
```

Example:

```text
main.go:15 Application started
```

If both `Llongfile` and `Lshortfile` are specified, `Lshortfile` takes precedence.

---

## `log.LUTC`

Uses UTC instead of local time when `Ldate` or `Ltime` is enabled.

```go
log.SetFlags(log.Ldate | log.Ltime | log.LUTC)
```

This is useful for distributed systems where servers may run in different time zones.

---

## `log.Lmsgprefix`

Moves the configured prefix from the beginning of the line to immediately before the message.

For example:

```go
logger := log.New(
	os.Stderr,
	"SERVER: ",
	log.Ldate|log.Ltime|log.Lmsgprefix,
)
```

---

## `log.LstdFlags`

This is the default combination:

```go
log.Ldate | log.Ltime
```

So:

```go
log.SetFlags(log.LstdFlags)
```

produces output similar to:

```text
2026/09/01 13:20:15 Application started
```

---

# 5. Every Function in the `log` Package

The `log` package exposes package-level functions plus methods on `Logger`.

## A. `log.Print()`

### Syntax

```go
log.Print(v ...any)
```

Prints a message using the standard logger.

It behaves similarly to `fmt.Print()`.

```go
log.Print("Hello")
log.Print("Hello ", "World")
```

Output:

```text
2026/09/01 13:20:15 Hello
2026/09/01 13:20:15 Hello World
```

Use it when you don't need formatted placeholders.

---

## B. `log.Printf()`

### Syntax

```go
log.Printf(format string, v ...any)
```

Useful when you need formatted output.

```go
name := "Chandu"
age := 25

log.Printf("Name: %s, Age: %d", name, age)
```

Output:

```text
2026/09/01 13:20:15 Name: Chandu, Age: 25
```

Think of it as:

```text
log + fmt.Printf-style formatting
```

---

## C. `log.Println()`

### Syntax

```go
log.Println(v ...any)
```

Prints a message and ensures that the log entry ends with a newline.

```go
log.Println("Server started")
log.Println("Listening on port", 8080)
```

This is one of the most commonly used basic logging functions for beginners.

---

## D. `log.Fatal()`

### Syntax

```go
log.Fatal(v ...any)
```

Logs the message and then terminates the program using:

```go
os.Exit(1)
```

Example:

```go
file, err := os.Open("config.txt")

if err != nil {
	log.Fatal("Could not open config file:", err)
}
```

The important point is:

```text
log.Fatal()
     ↓
write log message
     ↓
os.Exit(1)
     ↓
program terminates
```

### Important warning

Don't use `log.Fatal()` everywhere.

For example, this is usually bad inside reusable library code:

```go
func ReadConfig() {
	if err != nil {
		log.Fatal(err)
	}
}
```

The library shouldn't normally decide to terminate the entire application.

Instead, return the error:

```go
func ReadConfig() error {
	if err != nil {
		return err
	}

	return nil
}
```

Then the application can decide what to do.

---

## E. `log.Fatalf()`

### Syntax

```go
log.Fatalf(format string, v ...any)
```

This is the formatted version of `log.Fatal()`.

```go
port := 8080

if port < 1024 {
	log.Fatalf("Invalid port: %d", port)
}
```

Conceptually:

```text
log.Printf(...)
+
os.Exit(1)
```

---

## F. `log.Fatalln()`

### Syntax

```go
log.Fatalln(v ...any)
```

This is the `Println`-style version of `Fatal`.

```go
log.Fatalln("Database connection failed")
```

It:

1. Logs the message.
2. Terminates the program with exit status `1`.

---

## G. `log.Panic()`

### Syntax

```go
log.Panic(v ...any)
```

Logs the message and then calls:

```go
panic()
```

Example:

```go
func main() {
	log.Panic("Something went seriously wrong")
}
```

Output will contain the log message followed by a panic stack trace.

Unlike `Fatal`, `Panic` participates in Go's panic/recovery mechanism.

Conceptually:

```text
log.Panic()
     ↓
write log
     ↓
panic()
     ↓
defer/recover can potentially handle it
```

---

## H. `log.Panicf()`

### Syntax

```go
log.Panicf(format string, v ...any)
```

Formatted version of `Panic`.

```go
userID := 100

log.Panicf("Invalid user ID: %d", userID)
```

---

## I. `log.Panicln()`

### Syntax

```go
log.Panicln(v ...any)
```

`Println`-style version of `Panic`.

```go
log.Panicln("Critical application state detected")
```

It logs the message and then calls `panic()`.

---

## J. `log.SetFlags()`

### Syntax

```go
log.SetFlags(flag int)
```

Changes the flags used by the standard logger.

For example:

```go
log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

log.Println("Something happened")
```

Possible output:

```text
2026/09/01 13:20:15 main.go:10 Something happened
```

You combine flags using the bitwise OR operator:

```text
|
```

For example:

```go
log.Ldate | log.Ltime | log.Lshortfile
```

---

## K. `log.Flags()`

### Syntax

```go
log.Flags()
```

Returns the current flags configured for the standard logger.

Example:

```go
flags := log.Flags()

log.Println("Current flags:", flags)
```

You can use this when you need to inspect the current logging configuration.

---

## L. `log.SetPrefix()`

### Syntax

```go
log.SetPrefix(prefix string)
```

Adds a prefix to every message produced by the standard logger.

```go
log.SetPrefix("SERVER: ")

log.Println("Server started")
log.Println("Waiting for requests")
```

Output:

```text
2026/09/01 13:20:15 SERVER: Server started
2026/09/01 13:20:15 SERVER: Waiting for requests
```

This can be useful for identifying which component produced a message.

---

## M. `log.Prefix()`

### Syntax

```go
log.Prefix()
```

Returns the current prefix of the standard logger.

Example:

```go
log.SetPrefix("API: ")

fmt.Println(log.Prefix())
```

Output:

```text
API:
```

---

## N. `log.SetOutput()`

### Syntax

```go
log.SetOutput(w io.Writer)
```

Changes where the standard logger writes its output.

By default, the standard logger writes to standard error.

You can redirect it to a file:

```go
file, err := os.OpenFile(
	"application.log",
	os.O_CREATE|os.O_WRONLY|os.O_APPEND,
	0666,
)

if err != nil {
	log.Fatal(err)
}

defer file.Close()

log.SetOutput(file)

log.Println("Application started")
```

Now the log goes to:

```text
application.log
```

rather than the default output destination.

This is important because the function accepts `io.Writer`, allowing many different output destinations.

---

## O. `log.Writer()`

### Syntax

```go
log.Writer()
```

Returns the current output destination of the standard logger as an `io.Writer`.

Example:

```go
writer := log.Writer()

fmt.Println(writer)
```

This can be useful when another component needs access to the logger's current output destination.

---

## P. `log.Output()`

### Syntax

```go
log.Output(calldepth int, s string) error
```

This is a lower-level logging function.

```go
err := log.Output(1, "Custom log message")

if err != nil {
	fmt.Println("Logging failed:", err)
}
```

The `calldepth` parameter tells the logger how many stack frames to skip when determining the source file and line number.

This becomes particularly useful when creating your own logging helper functions.

For example:

```go
func myLog(message string) {
	log.Output(2, message)
}
```

The `2` helps the logger identify the caller of `myLog()` rather than incorrectly reporting the helper function itself.

---

# 6. The `Logger` Type

The package also provides:

```go
type Logger struct
```

A `Logger` represents an active logging object that writes messages to an `io.Writer`.

An important property is that a `Logger` can safely be used by multiple goroutines concurrently; it serializes access to its writer.

---

# 7. `log.New()`

### Syntax

```go
log.New(out io.Writer, prefix string, flag int) *log.Logger
```

Creates a new custom logger.

Example:

```go
logger := log.New(
	os.Stdout,
	"APP: ",
	log.Ldate|log.Ltime,
)

logger.Println("Application started")
```

Possible output:

```text
2026/09/01 13:20:15 APP: Application started
```

The three parameters are:

```text
out
 ↓
where logs go

prefix
 ↓
text added to log entries

flag
 ↓
date/time/file information
```

---

# 8. `log.Default()`

### Syntax

```go
log.Default()
```

Returns the standard logger used by the package-level functions.

For example:

```go
logger := log.Default()

logger.Println("Hello")
```

This is useful when you want to pass the standard logger to another function:

```go
func startServer(logger *log.Logger) {
	logger.Println("Server started")
}

func main() {
	startServer(log.Default())
}
```

---

# 9. `Logger.Print()`

### Syntax

```go
logger.Print(v ...any)
```

Same basic idea as:

```go
log.Print(...)
```

but operates on your custom logger.

```go
logger := log.New(os.Stdout, "APP: ", log.LstdFlags)

logger.Print("Hello")
```

---

# 10. `Logger.Printf()`

### Syntax

```go
logger.Printf(format string, v ...any)
```

Formatted output using the custom logger.

```go
logger.Printf("User ID: %d", 101)
```

---

# 11. `Logger.Println()`

### Syntax

```go
logger.Println(v ...any)
```

Prints using your custom logger.

```go
logger.Println("Server started")
```

---

# 12. `Logger.Fatal()`

### Syntax

```go
logger.Fatal(v ...any)
```

Logs through that particular logger and then terminates the process with exit code `1`.

```go
logger.Fatal("Unable to start server")
```

---

# 13. `Logger.Fatalf()`

### Syntax

```go
logger.Fatalf(format string, v ...any)
```

Formatted `Fatal`.

```go
logger.Fatalf("Unable to listen on port %d", 8080)
```

---

# 14. `Logger.Fatalln()`

### Syntax

```go
logger.Fatalln(v ...any)
```

`Println`-style fatal logging.

```go
logger.Fatalln("Unable to initialize database")
```

---

# 15. `Logger.Panic()`

### Syntax

```go
logger.Panic(v ...any)
```

Logs using the custom logger and then calls `panic()`.

```go
logger.Panic("Unexpected application state")
```

---

# 16. `Logger.Panicf()`

### Syntax

```go
logger.Panicf(format string, v ...any)
```

Formatted version of `Logger.Panic()`.

```go
logger.Panicf("Invalid value: %d", value)
```

---

# 17. `Logger.Panicln()`

### Syntax

```go
logger.Panicln(v ...any)
```

`Println`-style panic logging.

```go
logger.Panicln("Unexpected condition")
```

---

# 18. `Logger.SetFlags()`

### Syntax

```go
logger.SetFlags(flag int)
```

Changes the flags for a particular logger.

```go
logger.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
```

This differs from:

```go
log.SetFlags(...)
```

because the latter changes the standard logger, while this changes your specific custom logger.

---

# 19. `Logger.Flags()`

### Syntax

```go
logger.Flags()
```

Returns the flags currently configured for that logger.

```go
flags := logger.Flags()

fmt.Println(flags)
```

---

# 20. `Logger.SetPrefix()`

### Syntax

```go
logger.SetPrefix(prefix string)
```

Changes the prefix of a custom logger.

```go
logger.SetPrefix("DATABASE: ")

logger.Println("Connected")
```

---

# 21. `Logger.Prefix()`

### Syntax

```go
logger.Prefix()
```

Returns the logger's current prefix.

```go
prefix := logger.Prefix()

fmt.Println(prefix)
```

---

# 22. `Logger.SetOutput()`

### Syntax

```go
logger.SetOutput(w io.Writer)
```

Changes where that logger sends its output.

Example:

```go
logger := log.New(
	os.Stdout,
	"APP: ",
	log.LstdFlags,
)

file, err := os.OpenFile(
	"app.log",
	os.O_CREATE|os.O_WRONLY|os.O_APPEND,
	0666,
)

if err != nil {
	log.Fatal(err)
}

defer file.Close()

logger.SetOutput(file)

logger.Println("This goes into app.log")
```

---

# 23. `Logger.Writer()`

### Syntax

```go
logger.Writer()
```

Returns the current `io.Writer` used by that logger.

```go
writer := logger.Writer()
```

---

# 24. `Logger.Output()`

### Syntax

```go
logger.Output(calldepth int, s string) error
```

The custom-logger equivalent of the package-level `log.Output()`.

It is especially useful when writing helper functions.

For example:

```go
func logMessage(logger *log.Logger, message string) {
	logger.Output(2, message)
}
```

The call depth determines which caller information is reported when file/line flags are enabled.

---

# 25. Complete Example Using Several Features

```go
package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile(
		"application.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666,
	)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	logger := log.New(
		file,
		"APP: ",
		log.Ldate|log.Ltime|log.Lshortfile,
	)

	logger.Println("Application started")

	username := "Chandu"
	logger.Printf("User %s logged in", username)

	logger.Println("Processing request")

	logger.Println("Application finished")
}
```

This demonstrates:

```text
os.OpenFile()
      ↓
creates/opens log file
      ↓
log.New()
      ↓
creates custom Logger
      ↓
logger.Println()
      ↓
writes message
```

---

# 26. Three Common Beginner Mistakes

## Mistake 1: Using `log.Fatal()` for normal errors

Beginners sometimes write:

```go
if err != nil {
	log.Fatal(err)
}
```

everywhere.

The problem is that `log.Fatal()` terminates the entire program.

### Better approach

For functions that can recover or let their caller decide what to do:

```go
if err != nil {
	return err
}
```

Reserve `Fatal` for situations where continuing the application is genuinely impossible or inappropriate.

---

## Mistake 2: Confusing `Fatal`, `Panic`, and normal logging

These are very different:

```go
log.Println("message")
```

continues execution.

```go
log.Panic("message")
```

logs and invokes `panic()`.

```go
log.Fatal("message")
```

logs and calls `os.Exit(1)`.

A useful mental model is:

```text
Print
  ↓
log only

Panic
  ↓
log + panic()

Fatal
  ↓
log + os.Exit(1)
```

One important consequence: because `os.Exit(1)` terminates the process immediately, deferred functions are not executed by `log.Fatal()`.

---

## Mistake 3: Treating `log` as a complete structured logging system

The traditional `log` package primarily produces formatted text.

For example:

```go
log.Printf(
	"user=%d action=%s status=%s",
	userID,
	action,
	status,
)
```

This can work well for simple applications.

However, modern applications often benefit from structured logging, where fields are represented as structured key-value attributes.

Go's `log/slog` package was designed specifically for structured logging.

So understand the distinction:

```text
log
 ↓
simple text logging

log/slog
 ↓
structured logging
```

---

# 27. Real-World Application 1: Web Server Logging

Imagine you're developing a REST API.

You might want to record:

```text
2026/09/01 13:20:10 Server started on port 8080
2026/09/01 13:20:15 GET /users
2026/09/01 13:20:16 POST /users
2026/09/01 13:20:17 Database connection failed
```

Logging helps you understand what happened when something goes wrong.

A simple HTTP handler might use:

```go
func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf(
		"method=%s path=%s",
		r.Method,
		r.URL.Path,
	)

	w.Write([]byte("Hello"))
}
```

This is particularly useful during development and when diagnosing application behavior.

---

# 28. Real-World Application 2: Background Jobs

Suppose your Go application processes orders in the background.

You might log:

```text
2026/09/01 13:20:15 Starting order processing
2026/09/01 13:20:16 Processing order 1001
2026/09/01 13:20:16 Order 1001 completed
2026/09/01 13:20:17 Processing order 1002
2026/09/01 13:20:18 Order 1002 failed
```

This gives developers and operators a historical trail of what happened.

For a production system, think carefully about log volume, sensitive data, structured logging, and how logs will be collected and searched.

---

# 29. `log` vs `fmt.Println`

A common beginner question is:

> Why not just use `fmt.Println()`?

You can use `fmt.Println()` for simple console output.

But logging gives you additional capabilities:

| Feature | `fmt.Println()` | `log.Println()` |
|---|---:|---:|
| Simple output | Yes | Yes |
| Timestamp | No | Yes |
| Prefix | No | Yes |
| File/line information | No | Yes |
| Configurable destination | Limited/manual | Yes |
| Custom Logger | No | Yes |
| Fatal/Panic helpers | No | Yes |

So:

```go
fmt.Println("Hello")
```

is generally just output.

Whereas:

```go
log.Println("Hello")
```

is an application diagnostic/logging event.

---

# 30. Three Progressive Exercises

## Exercise 1 — Basic Application Logger

Create a Go program that:

1. Creates a custom `Logger`.
2. Uses the prefix `APP: `.
3. Includes the date and time.
4. Logs when the application starts.
5. Logs a user's name and age.
6. Logs when the application finishes.

**Do not use `fmt.Println()` for the application messages.**

---

## Exercise 2 — File-Based Server Logger

Create a small HTTP server that:

1. Listens on a configurable port.
2. Creates a log file named `server.log`.
3. Sends application logs to that file.
4. Logs every incoming request.
5. Records the HTTP method and URL path.
6. Uses `Lshortfile` so that the source file and line number appear in the logs.
7. Logs an appropriate message if the server fails to start.

Think carefully about which errors should be returned and which, if any, justify terminating the application.

---

## Exercise 3 — Build a Logging Helper

Create your own logging helper system around `log.Logger`.

Your program should have separate loggers for:

```text
INFO
ERROR
DATABASE
HTTP
```

Each logger should have:

- Its own prefix.
- Appropriate flags.
- A configurable output destination.

Then create helper functions such as:

```go
logInfo(...)
logError(...)
logDatabase(...)
logHTTP(...)
```

Use `Logger.Output()` and an appropriate `calldepth` so that when file/line information is enabled, the log points to the original caller rather than the helper function.

Finally, make your program safely handle several different error conditions without terminating unnecessarily.

**Do not use `log.Fatal()` as a substitute for proper error handling.**

---

# 31. Important Mental Model

If you're learning the package, remember this hierarchy:

```text
                    log package
                         │
            ┌────────────┴────────────┐
            │                         │
     Standard Logger            Custom Logger
            │                         │
       log.Println()          logger.Println()
       log.Printf()           logger.Printf()
       log.Fatal()            logger.Fatal()
       log.Panic()            logger.Panic()
       log.SetFlags()         logger.SetFlags()
       log.SetOutput()        logger.SetOutput()
            │                         │
            └────────────┬────────────┘
                         │
                     io.Writer
                         │
              ┌──────────┼──────────┐
              │          │          │
           stderr       file      custom
```

The biggest concepts to master are:

1. Standard logger vs custom logger
2. Output destination
3. Prefixes
4. Flags
5. `Print` vs `Printf` vs `Println`
6. `Fatal` vs `Panic`
7. Custom logging with `Logger.Output()`
8. Using `io.Writer` for flexible destinations
9. When simple `log` is sufficient
10. When structured `log/slog` is a better choice

---

# Thought-Provoking Question

Imagine your Go application is running on 100 servers and suddenly starts producing 50,000 log messages per minute.

**Would simply adding more `log.Println()` statements make the system easier to diagnose—or could logging itself become a performance, storage, security, and observability problem? What would you change in your logging design, and why?**

---

## Official Documentation

- Go `log` package: https://pkg.go.dev/log
- Go `log/slog` package: https://pkg.go.dev/log/slog
