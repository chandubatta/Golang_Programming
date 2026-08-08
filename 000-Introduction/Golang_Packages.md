# Packages in Golang

## 1. What are Packages in Go?

A **package** in Go is a way of organizing and grouping related Go source files and code together.

A package can contain:

- Functions
- Variables
- Constants
- Types
- Structs
- Interfaces
- Methods

Every Go program belongs to a package. The most common package is:

```go
package main
```

### Purpose of Packages

Packages help you:

1. **Organize code** into logical units.
2. **Reuse code** across different programs.
3. **Avoid naming conflicts** by providing namespaces.
4. **Improve maintainability** of large applications.
5. **Control visibility** using exported and unexported identifiers.

### When are Packages Commonly Used?

Packages are especially useful when your application becomes larger.

For example, a backend application might be organized like:

```text
myapp/
├── main.go
├── user/
│   └── user.go
├── payment/
│   └── payment.go
├── database/
│   └── database.go
└── auth/
    └── auth.go
```

Each directory can represent a separate package.

---

# 2. Simple Package Example

Let's create a small application with two packages.

### Project Structure

```text
myapp/
├── go.mod
├── main.go
└── calculator/
    └── calculator.go
```

### `calculator/calculator.go`

```go
package calculator

func Add(a int, b int) int {
    return a + b
}

func Multiply(a int, b int) int {
    return a * b
}
```

Notice that `Add` and `Multiply` start with **capital letters**.

That means they are **exported** and can be accessed from another package.

### `main.go`

```go
package main

import (
    "fmt"

    "myapp/calculator"
)

func main() {
    result1 := calculator.Add(10, 20)
    result2 := calculator.Multiply(5, 4)

    fmt.Println("Addition:", result1)
    fmt.Println("Multiplication:", result2)
}
```

Output:

```text
Addition: 30
Multiplication: 20
```

### Important Concept

This:

```go
calculator.Add(10, 20)
```

means:

```text
package name + function name
```

So:

```text
calculator.Add
     │       │
     │       └── Function
     └────────── Package
```

---

## Exported vs Unexported

Go uses capitalization to control visibility.

### Exported

```go
func Add(a int, b int) int {
    return a + b
}
```

Because `Add` starts with a capital letter, other packages can use it.

### Unexported

```go
func add(a int, b int) int {
    return a + b
}
```

Because `add` starts with lowercase, it can only be accessed inside its package.

This is an important feature for **encapsulation**.

---

# 3. Three Common Beginner Mistakes

## Mistake 1: Thinking Every Function Must Be Exported

Beginners sometimes write:

```go
func Add()
func Calculate()
func Validate()
```

even when those functions are only needed internally.

### How to Avoid It

Use lowercase names for internal implementation details:

```go
func calculateTax() float64 {
    // internal logic
}
```

Use uppercase names only when other packages need access:

```go
func CalculateTax() float64 {
    // public API
}
```

---

## Mistake 2: Confusing Package Names with Folder Names

Suppose you have:

```text
project/
└── calculator/
    └── calculator.go
```

The directory is called `calculator`, and normally the package may also be:

```go
package calculator
```

But Go's package name is determined by the `package` declaration, not simply by the folder name.

For example:

```go
package mathutil
```

could exist inside a directory named `calculator`.

### How to Avoid It

Understand the distinction:

```text
Directory → location of package source files

Package name → declared using "package ..."
```

---

## Mistake 3: Forgetting the Go Module Import Path

A beginner may write:

```go
import "calculator"
```

for a local package.

But in a Go module, the import path normally comes from the module name in `go.mod`.

For example:

```go
module myapp
```

Then:

```go
import "myapp/calculator"
```

### How to Avoid It

Always understand your `go.mod`:

```go
module myapp
```

and build local imports from that module path:

```text
myapp/
    └── calculator/
```

```go
import "myapp/calculator"
```

---

# 4. Real-World Applications

## Application 1: Backend REST API

A large Go backend can separate responsibilities into packages:

```text
backend/
├── main.go
├── handlers/
├── services/
├── repository/
├── models/
├── middleware/
└── database/
```

For example:

```text
handlers
    ↓
services
    ↓
repository
    ↓
database
```

Each package has a specific responsibility.

This makes a large backend easier to:

- Maintain
- Test
- Debug
- Extend
- Reuse

This is particularly useful for Golang REST API and microservice applications.

---

## Application 2: Reusable Utility Library

Suppose you repeatedly need email validation.

You could create:

```text
myproject/
└── validator/
    └── email.go
```

with functionality such as:

```go
validator.IsValidEmail(...)
```

Then multiple parts of your application can reuse the same package instead of duplicating the validation logic.

---

# 5. Progressive Exercises

## Exercise 1 — Beginner: Calculator Package

Create a package named `calculator`.

Implement functions for:

- Addition
- Subtraction
- Multiplication
- Division

Create a separate `main` package that imports your calculator package and demonstrates all four operations.

**Requirement:** Make the calculator functions accessible from the `main` package.

---

## Exercise 2 — Intermediate: User Package

Create a `user` package containing a `User` struct with:

- ID
- Name
- Email
- Age

Add functions that allow the application to:

- Create a user
- Display user information
- Check whether the user is an adult
- Update user information

Then create a `main` package that imports and uses the `user` package.

**Challenge:** Decide which fields and functions should be exported and which should remain private.

---

## Exercise 3 — Advanced: Mini REST API Package Structure

Create a small Go REST API and organize it into multiple packages.

Suggested structure:

```text
myapi/
├── main.go
├── handlers/
├── services/
├── models/
├── repository/
└── database/
```

Your application should provide an API for managing users.

Design the packages so that:

```text
HTTP Request
     ↓
handlers
     ↓
services
     ↓
repository
     ↓
database
```

Requirements:

- `models` contains your user data structure.
- `handlers` handles HTTP requests and responses.
- `services` contains business logic.
- `repository` handles data access.
- `database` manages database-related functionality.
- `main` initializes the application and routes.

**Challenge:** Decide what should be exported from each package and what should remain private.

---

# 🧠 Think Deeper

Imagine you are building a **large e-commerce application with 50+ developers**.

You could put everything into one enormous `main.go` file, or divide the application into many packages.

**Question:**

> If packages make code more organized and reusable, can having *too many packages* actually make a Go application harder to understand and maintain? Where would you draw the boundary between "good separation" and "over-engineering"?
