# Built-in Packages in Go

> **Important terminology:** Go does not officially call its standard-library packages “built-in packages.” Usually, learners use **“built-in packages”** to mean Go's **standard library packages**, which are provided with the Go installation and can be imported without installing third-party libraries.

Examples include `fmt`, `os`, `strings`, `strconv`, `time`, `net/http`, `encoding/json`, and many others.

---

## 1. What are Built-in Packages?

Go's **standard library** is a collection of packages that provide ready-to-use functionality for common programming tasks.

Instead of writing everything yourself, you can import a package and use the functionality it provides.

### Examples

| Package | Purpose |
|---|---|
| `fmt` | Printing and formatted input/output |
| `os` | Operating-system functionality |
| `strings` | String manipulation |
| `strconv` | Converting strings and numbers |
| `math` | Mathematical operations |
| `time` | Dates, times, and durations |
| `net/http` | HTTP clients and servers |
| `encoding/json` | JSON encoding/decoding |
| `io` | Reading and writing data |
| `path/filepath` | Working with file paths |
| `errors` | Creating and working with errors |

### Why are they useful?

They help you:

- Avoid rewriting common functionality.
- Write programs faster.
- Use well-tested standard functionality.
- Build applications without depending on third-party libraries for basic tasks.
- Keep your project simpler when the standard library is sufficient.

### When are they commonly used?

Almost every Go application uses standard-library packages.

For example:

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    name := "golang programming"

    fmt.Println(strings.ToUpper(name))
}
```

Output:

```text
GOLANG PROGRAMMING
```

Here:

```go
import (
    "fmt"
    "strings"
)
```

imports two standard-library packages.

---

# 2. Simple Code Example

Let's create a small program that uses several standard-library packages.

```go
package main

import (
    "fmt"
    "strconv"
    "strings"
)

func main() {
    name := "  chandu  "
    ageText := "25"

    // strings package
    name = strings.TrimSpace(name)
    name = strings.ToUpper(name)

    // strconv package
    age, err := strconv.Atoi(ageText)

    if err != nil {
        fmt.Println("Invalid age")
        return
    }

    // fmt package
    fmt.Println("Name:", name)
    fmt.Println("Age:", age)
}
```

Output:

```text
Name: CHANDU
Age: 25
```

### What happened?

We used three standard-library packages.

### `fmt`

```go
fmt.Println("Name:", name)
```

Used for formatted input/output.

### `strings`

```go
strings.TrimSpace(name)
strings.ToUpper(name)
```

Used for string manipulation.

### `strconv`

```go
age, err := strconv.Atoi(ageText)
```

Converts a string into an integer.

So instead of implementing our own functions for these operations, we use functionality already provided by Go.

---

# 3. Three Common Mistakes

## Mistake 1: Thinking every package is a built-in package

A beginner may think:

> "If I can import it, it must be a built-in package."

Not necessarily.

For example:

```go
import "fmt"
```

`fmt` is part of Go's standard library.

But:

```go
import "github.com/someuser/somepackage"
```

is a third-party package.

### How to avoid it

Learn to distinguish:

```text
Go standard library
        ↓
fmt
os
strings
time
net/http
encoding/json
...
```

from external dependencies:

```text
Third-party packages
        ↓
github.com/...
```

---

## Mistake 2: Thinking standard-library packages need to be installed separately

You generally don't install packages such as `fmt` or `strings` using:

```bash
go get fmt
```

They come with the Go distribution.

You simply import them:

```go
import "fmt"
```

### How to avoid it

Remember:

```text
Standard library → comes with Go
Third-party package → usually added as a dependency
```

---

## Mistake 3: Importing a package but not using it

For example:

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    fmt.Println("Hello")
}
```

Here `strings` is imported but never used.

Go will report an error similar to:

```text
"strings" imported and not used
```

Go intentionally prevents unused imports.

### How to avoid it

Only import packages you actually need.

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello")
}
```

---

# 4. Real-World Applications

## Application 1: Building a REST API

The standard library provides `net/http`, which can be used to create HTTP servers.

For example:

```go
package main

import (
    "fmt"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello from Go API")
}

func main() {
    http.HandleFunc("/hello", helloHandler)

    http.ListenAndServe(":8080", nil)
}
```

You can then access:

```text
http://localhost:8080/hello
```

This demonstrates how standard-library packages can be used to build actual web services.

---

## Application 2: Reading and processing files

The `os` package can be used to work with files.

For example:

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    data, err := os.ReadFile("data.txt")

    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println(string(data))
}
```

This type of functionality is useful for:

- Configuration files
- Log processing
- Data import/export
- File-based applications
- Scripts and automation tools

---

# 5. Progressive Exercises

## Exercise 1 — Beginner

### String Information Tool

Create a Go program that:

1. Accepts a sentence from the user.
2. Removes leading and trailing spaces.
3. Converts the sentence to uppercase.
4. Converts the sentence to lowercase.
5. Prints the number of characters in the sentence.

**Requirement:** Use appropriate Go standard-library packages.

**Do not write your own string manipulation functions.**

---

## Exercise 2 — Intermediate

### Simple File Analyzer

Create a program that reads a text file and analyzes its contents.

The program should:

1. Open a text file.
2. Read its contents.
3. Count the number of:
   - Lines
   - Words
   - Characters
4. Display the results.
5. Handle the situation where the file does not exist.

**Requirement:** Use Go standard-library packages for file handling and string processing.

---

## Exercise 3 — Advanced

### Mini REST API

Build a small HTTP API using Go's standard library.

Create the following endpoints:

```text
GET /users
GET /users/{id}
POST /users
```

Your program should:

1. Start an HTTP server.
2. Maintain a collection of users in memory.
3. Return users as JSON.
4. Accept JSON when creating a user.
5. Validate incoming data.
6. Return appropriate HTTP status codes.
7. Handle invalid JSON.
8. Handle a user ID that doesn't exist.
9. Use standard-library packages rather than a third-party web framework.

This exercise will make you combine several standard-library packages, such as HTTP handling, JSON processing, string/number conversion, and error handling.

---

# A Useful Mental Model

Think of Go's standard library like a **toolbox**:

```text
                 Go
                  │
          ┌───────┴────────┐
          │                │
      Language         Standard Library
      features              │
                         Packages
                            │
       ┌────────┬───────────┼───────────┐
       ↓        ↓           ↓           ↓
      fmt     strings       os       net/http
       │        │           │           │
     I/O     Strings      Files       HTTP
```

You don't need to build every tool yourself.

You select the appropriate package:

```text
Need printing?            → fmt
Need string operations?   → strings
Need files?               → os
Need JSON?                → encoding/json
Need HTTP?                → net/http
Need time?                → time
Need number conversion?  → strconv
```

That is one of the major strengths of Go's standard library.

---

# 🤔 Thought-provoking Question

Suppose you are building a **production REST API** in Go.

The Go standard library already provides `net/http`, `encoding/json`, `context`, error handling, and many other useful packages.

**At what point would you choose to introduce a third-party package or framework instead of using only the standard library—and what trade-offs would you consider before making that decision?**

This is an important question because becoming a good Go developer isn't just about knowing **which package to use**, but also knowing **when not to introduce another dependency**.
