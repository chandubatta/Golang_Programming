# Cross-Platform Support in Go

## 1. What is Cross-Platform Support?

**Cross-platform support** means writing a Go program that can run on multiple operating systems and CPU architectures with little or no change to the source code.

Go provides excellent cross-platform support because the Go compiler can compile the same Go source code for different target platforms.

For example, you can build an application for:

- Windows
- Linux
- macOS
- Android-related environments
- Different CPU architectures such as AMD64 and ARM64

### Purpose

The main purpose is to **write code once and build it for multiple platforms**.

```text
Go Source Code
      |
      +----> Windows  → .exe
      |
      +----> Linux    → executable
      |
      +----> macOS    → executable
      |
      +----> ARM64    → ARM executable
```

### When is it commonly used?

Cross-platform support is particularly useful when:

- Developing CLI tools
- Building backend services
- Creating DevOps tools
- Deploying applications to different servers
- Supporting Windows, Linux, and macOS developers
- Building applications for different CPU architectures
- Creating Docker/container images for different architectures

---

# 2. Simple Go Example

Let's create a program that prints the operating system and CPU architecture on which the program is running.

```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Operating System:", runtime.GOOS)
	fmt.Println("Architecture:", runtime.GOARCH)
}
```

### Run it

On your current machine:

```bash
go run main.go
```

You might see:

```text
Operating System: windows
Architecture: amd64
```

The same source code can be compiled for Linux.

### Build for Linux

```bash
GOOS=linux GOARCH=amd64 go build -o myapp
```

For macOS:

```bash
GOOS=darwin GOARCH=amd64 go build -o myapp
```

For Windows:

```bash
GOOS=windows GOARCH=amd64 go build -o myapp.exe
```

For ARM64 Linux:

```bash
GOOS=linux GOARCH=arm64 go build -o myapp
```

The important idea is:

> **The Go source code remains the same; the target operating system and architecture can change during compilation.**

---

# 3. Three Common Beginner Mistakes

## Mistake 1: Assuming all operating-system APIs are the same

A beginner might write:

```go
os.Remove("C:\\temp\\file.txt")
```

This works with a Windows-style path but is not portable to Linux or macOS.

### Better approach

Use Go's `filepath` package:

```go
package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	path := filepath.Join("data", "files", "example.txt")

	fmt.Println(path)
}
```

Go automatically uses the appropriate path separator.

**Avoid it:** Don't hard-code OS-specific paths when you can use Go's standard library.

---

## Mistake 2: Thinking cross-platform means "every program works identically everywhere"

Go makes compilation for different platforms easy, but your application may still depend on:

- Operating-system APIs
- File-system behavior
- Environment variables
- External commands
- Hardware
- Network configuration

For example:

```go
exec.Command("cmd", "/C", "dir")
```

is Windows-specific.

Linux commonly uses:

```go
exec.Command("ls")
```

**Avoid it:** Minimize OS-specific dependencies and isolate platform-specific code when necessary.

---

## Mistake 3: Forgetting CPU architecture

A program isn't only dependent on the operating system.

For example:

```text
Linux + AMD64
Linux + ARM64
```

are both Linux, but they use different CPU architectures.

Building for the wrong architecture can result in an executable that won't run on the target machine.

**Avoid it:** Know your deployment environment and explicitly choose the appropriate `GOOS` and `GOARCH`.

---

# 4. Real-World Applications

## Application 1: DevOps and CLI Tools

Go is well suited for command-line tools because you can compile one project into binaries for different environments.

For example:

```text
mytool.exe       → Windows
mytool           → Linux
mytool           → macOS
```

A development team can distribute the same tool to developers using different operating systems.

---

## Application 2: Cloud and Container Deployment

Suppose you develop a Go backend on Windows but deploy it to a Linux server.

You can build the Linux binary from your development machine:

```bash
GOOS=linux GOARCH=amd64 go build -o server
```

Then deploy:

```text
Windows Development Machine
          |
          | Go Build
          ↓
     Linux Binary
          |
          ↓
     Docker Container
          |
          ↓
        AWS
```

This is especially useful for backend services, microservices, Docker containers, and cloud deployments.

---

# 5. Progressive Exercises

## Exercise 1 — Beginner ⭐

Create a Go program that displays:

```text
Operating System: <OS>
Architecture: <Architecture>
Go Version: <Go Version>
```

Use Go's `runtime` package.

Your program should work without modifying the source code when executed on different operating systems.

---

## Exercise 2 — Intermediate ⭐⭐

Create a **cross-platform file-management CLI application**.

The application should:

1. Create a `data` directory.
2. Create a text file inside it.
3. Write some text into the file.
4. Read the file.
5. Display the file contents.
6. Delete the file.

The program should work on Windows, Linux, and macOS.

**Requirement:** Do not hard-code `/` or `\` as path separators.

---

## Exercise 3 — Advanced ⭐⭐⭐

Create a cross-platform **system information CLI tool**.

The tool should display:

```text
===== System Information =====

Operating System:
Architecture:
Go Version:
Number of CPUs:
Current Working Directory:

==============================
```

Then add a feature that creates a platform-specific output file:

```text
Windows → system-info.txt
Linux   → system-info.txt
macOS   → system-info.txt
```

Finally, compile your program for at least three different combinations of `GOOS` and `GOARCH` and verify that the generated binaries work on their target platforms.

**Do not use OS-specific commands such as `dir`, `ls`, or `pwd`.**

---

# 🧠 Think Deeper

Suppose you are building a **Golang microservice that must run on Windows during development, Linux in production, and ARM64 inside some cloud environments**.

**What parts of your Go application would you design differently to make the code truly cross-platform, and what kinds of dependencies would you try to avoid?**
