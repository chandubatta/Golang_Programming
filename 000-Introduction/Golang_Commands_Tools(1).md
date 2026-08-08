# Golang Commands (Tools)

## 1. Concise Explanation

In Go, **commands (tools)** are mainly the commands provided by the Go toolchain through the `go` command. They help you create, build, run, test, format, manage dependencies, inspect code, and maintain Go projects.

Some important commands:

| Command | Purpose |
|---|---|
| `go run` | Compile and run a Go program |
| `go build` | Compile a Go program/package |
| `go test` | Run tests |
| `go fmt` | Format Go source code |
| `go mod init` | Create a new Go module |
| `go mod tidy` | Clean and synchronize dependencies |
| `go get` | Add/update dependencies |
| `go install` | Compile and install a Go executable |
| `go vet` | Find suspicious code and common mistakes |
| `go doc` | Display documentation |
| `go version` | Display the installed Go version |
| `go env` | Display Go environment settings |

### Common Development Workflow

```text
Write Code
   ↓
go fmt
   ↓
go run
   ↓
go test
   ↓
go vet
   ↓
go build
   ↓
Deploy
```

The major advantage is that Go provides a **standard toolchain**, so developers can use the same basic commands across different Go projects and operating systems.

---

## 2. Simple Code Example

Let's create a small Go program and use several Go commands with it.

### Step 1: Create a project

```bash
mkdir hello
cd hello
```

Initialize a Go module:

```bash
go mod init example.com/hello
```

This creates:

```text
hello/
├── go.mod
└── main.go
```

### Step 2: Create `main.go`

```go
package main

import "fmt"

func main() {
    name := "Chandu"

    fmt.Println("Hello,", name)
}
```

### Step 3: Format the code

```bash
go fmt
```

Go automatically formats the source code according to Go's standard formatting rules.

### Step 4: Run the program

```bash
go run .
```

Output:

```text
Hello, Chandu
```

### Step 5: Build the program

```bash
go build
```

This compiles your program and creates an executable.

### Step 6: Check your Go version

```bash
go version
```

Example:

```text
go version go1.25.3 windows/amd64
```

---

## 3. Three Common Mistakes and Misconceptions

### Mistake 1: Thinking `go run` creates a normal executable for deployment

Beginners often think:

```bash
go run main.go
```

is the same as creating a production executable.

Actually, `go run` is primarily convenient for **development/testing**. For deployment, you will normally use:

```bash
go build
```

#### How to avoid it

Use:

```bash
go run .
```

while developing, and:

```bash
go build
```

when you need a compiled executable.

---

### Mistake 2: Forgetting `go fmt`

Beginners sometimes manually format Go code or use inconsistent formatting.

Go provides a standard formatter:

```bash
go fmt
```

For example:

```go
func main( ){
fmt.Println("Hello")
}
```

can be formatted automatically into Go's standard style.

#### How to avoid it

Run:

```bash
go fmt ./...
```

to format packages throughout the project.

---

### Mistake 3: Using `go get` without understanding modules

A beginner may see a package online and immediately run:

```bash
go get some-package
```

without understanding how dependencies are tracked.

Modern Go projects generally use **Go modules**, with dependency information maintained through `go.mod` and checksums through `go.sum`.

#### How to avoid it

Understand these commands:

```bash
go mod init
go get
go mod tidy
```

and regularly inspect your:

```text
go.mod
go.sum
```

files.

---

## 4. Real-World Applications

### Application 1: Backend API Development

Suppose you're developing a REST API in Go using Gin or Gorilla Mux.

During development, you might repeatedly use:

```bash
go run .
```

After adding functionality:

```bash
go test ./...
```

Then check for suspicious code:

```bash
go vet ./...
```

Finally, create the production executable:

```bash
go build
```

So the Go commands become part of your everyday backend development workflow.

---

### Application 2: CI/CD and Production Deployment

Go commands are particularly useful in automated CI/CD pipelines.

A pipeline might perform:

```bash
go fmt ./...
go test ./...
go vet ./...
go build
```

For example:

```text
Developer pushes code
        ↓
GitHub
        ↓
CI Pipeline
        ↓
go test ./...
        ↓
go vet ./...
        ↓
go build
        ↓
Docker image
        ↓
Kubernetes
        ↓
AWS
```

This allows a team to automatically verify and compile Go applications before deployment.

---

## 5. Three Progressively Challenging Exercises

### Exercise 1 — Beginner: Basic Go Commands

Create a Go program that prints:

```text
Name: Chandu
Language: Go
Role: Backend Developer
```

Then practice using:

1. `go mod init`
2. `go fmt`
3. `go run`
4. `go build`
5. `go version`

**Goal:** Become comfortable with the basic Go toolchain.

---

### Exercise 2 — Intermediate: Testing and Code Quality

Create a small calculator package containing functions for:

- Addition
- Subtraction
- Multiplication
- Division

Create unit tests for those functions.

Then use:

```bash
go fmt ./...
go test ./...
go vet ./...
```

**Goal:** Understand how Go commands can be combined to improve code quality.

---

### Exercise 3 — Advanced: Mini Go Project Workflow

Create a small REST API project with at least:

- A Go module
- Multiple packages
- Several API endpoints
- Unit tests
- At least one external dependency

Design a workflow where you use:

```bash
go mod init
go get
go mod tidy
go fmt ./...
go test ./...
go vet ./...
go build
```

Document what each command does and determine **at which stage of development each command should be executed**.

**Goal:** Simulate a real-world Go development and build workflow.

---

## Thought-Provoking Question

If `go build`, `go test`, `go fmt`, `go vet`, and `go mod tidy` can all be automated in a CI/CD pipeline, **what problems could occur if a development team skips some of these commands—and which ones would you consider essential before deploying a Go backend to production?**

---

## Follow-Up Challenge

Design your own production-ready Go CI/CD pipeline using only Go commands.

Think about:

```text
Code
 ↓
?
 ↓
?
 ↓
?
 ↓
Build
 ↓
Deploy
```

Explain the purpose of each command you choose and why you placed it at that stage.
