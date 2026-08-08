# Run Without a Virtual Machine Like JVM in Go

## 1. Concise Explanation

In Go, **“Run Without a Virtual Machine, like JVM”** means that Go programs do not require a separate virtual machine such as the Java Virtual Machine (JVM) to execute the compiled application.

Go uses a **compiled-to-native-code** approach.

### Java execution model

```text
Java Source Code
       ↓
   javac compiler
       ↓
   Bytecode (.class)
       ↓
      JVM
       ↓
   Machine Code
       ↓
      CPU
```

### Go execution model

```text
Go Source Code
       ↓
   Go Compiler
       ↓
Native Machine Code
       ↓
      CPU
```

The Go compiler (`go build` or `go run`) compiles Go source code into native machine code for the target operating system and architecture.

For example:

```bash
go build main.go
```

This produces an executable:

```text
main.exe       # Windows
main           # Linux/macOS
```

You can then run the executable directly without installing a JVM.

### Purpose

The main purposes are:

- **Fast startup** — no JVM startup process is required.
- **Simple deployment** — the compiled executable can often be deployed as a single binary.
- **Lower runtime dependency** — applications don't require a separate VM such as the JVM.
- **Good performance** — Go produces native machine code.
- **Easy container deployment** — Go applications can be packaged efficiently in Docker containers.

### When is it commonly used?

This approach is particularly useful for:

- Backend services
- REST APIs
- Microservices
- CLI applications
- Cloud applications
- Docker/Kubernetes workloads
- Network services

---

## 2. Simple Code Example

Create a file called `main.go`:

```go
package main

import "fmt"

func main() {
    message := "Hello from Go!"

    fmt.Println(message)
}
```

Run it directly:

```bash
go run main.go
```

Output:

```text
Hello from Go!
```

`go run` compiles the program and executes it.

You can also create an executable:

```bash
go build -o hello main.go
```

Then run it:

### Windows

```bash
hello.exe
```

### Linux/macOS

```bash
./hello
```

The important point is that after compilation, the resulting executable contains the compiled Go program and its required Go runtime components; **it does not need a JVM to execute**.

Conceptually:

```text
main.go
   ↓
Go Compiler
   ↓
hello.exe
   ↓
Operating System
   ↓
CPU
```

---

## 3. Common Mistakes and Misconceptions

### Mistake 1: Thinking Go has absolutely no runtime

A common misconception is:

> “Go doesn't have a runtime at all.”

That's not quite correct.

Go does not require a **separate VM like the JVM**, but Go programs include/use the **Go runtime**, which provides functionality such as:

- Goroutine scheduling
- Garbage collection
- Memory management
- Stack management
- Some runtime support

So:

```text
Java → JVM + Java Runtime
Go   → Native executable + Go runtime support
```

**How to avoid it:**  
Remember that **“no JVM” does not mean “no runtime.”**

---

### Mistake 2: Thinking Go programs run identically on every OS

Go can compile programs for many operating systems and CPU architectures, but the executable is normally **platform-specific**.

For example:

```text
Windows → Windows executable
Linux   → Linux executable
macOS   → macOS executable
```

You can cross-compile:

```bash
GOOS=linux GOARCH=amd64 go build
```

**How to avoid it:**  
Understand the difference between **cross-platform source code** and a **platform-specific compiled binary**.

---

### Mistake 3: Thinking `go run` means Go is interpreted

A beginner might think:

> “I use `go run`, so Go must be interpreted.”

It isn't.

`go run` essentially performs compilation and then executes the resulting program.

Conceptually:

```text
go run main.go
       ↓
    Compile
       ↓
   Executable
       ↓
    Execute
```

**How to avoid it:**  
Distinguish between the **command you use to run a program** and the **way the language executes**.

---

## 4. Real-World Applications

### Application 1: Go Microservices

Suppose you create a payment microservice:

```text
payment-service
       ↓
     Go
       ↓
   go build
       ↓
payment-service binary
       ↓
     Docker
       ↓
   Kubernetes
```

You can package the compiled Go binary into a lightweight container without installing a JVM.

This is useful for:

- REST APIs
- gRPC services
- Payment services
- Authentication services
- Product services
- Order services

---

### Application 2: Cloud and CLI Applications

Go is also useful for command-line and cloud infrastructure software.

For example, a CLI application could be:

```text
mytool.exe
```

The user can execute:

```bash
mytool.exe
```

without installing:

```text
JVM
.NET Runtime
Python
```

This makes Go convenient for distributing standalone developer tools and infrastructure utilities.

---

## 5. Progressive Exercises

### Exercise 1 — Beginner

Create a simple Go program that:

1. Defines your name, age, and profession.
2. Prints them to the terminal.
3. Compile the program using `go build`.
4. Run the generated executable directly.
5. Verify that the program works without starting a JVM.

**Do not use any external libraries.**

---

### Exercise 2 — Intermediate

Create a Go program that:

1. Accepts two numbers from the user.
2. Performs addition, subtraction, multiplication, and division.
3. Prints the results.
4. Compile it into an executable.
5. Run the executable from the command line.
6. Identify which parts of the program are compiled before execution.
7. Explain why a JVM is not required.

---

### Exercise 3 — Advanced

Create a small Go HTTP server that:

1. Provides a `/hello` endpoint.
2. Returns a JSON response.
3. Uses goroutines to handle requests.
4. Builds the application into a standalone executable.
5. Runs the executable on your operating system.
6. Cross-compile the application for a different operating system/architecture.
7. Compare the generated binaries.
8. Explain which parts are provided by the Go runtime and which parts are native compiled application code.

**Do not use Docker or third-party frameworks for this exercise.**

---

## 🤔 Think Deeper

If Go can compile an application into a native executable without requiring a JVM, **why do you think Go still needs a runtime for things such as garbage collection and goroutine scheduling?**

And more importantly:

> **What advantages and disadvantages might this design have compared with Java's JVM-based approach when building large-scale microservices?**
