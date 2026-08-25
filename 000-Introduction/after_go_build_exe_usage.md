# After `go build .` Command: How to Use the `.exe` File in Golang

## 1. What happens after `go build .`?

In Go, when you run:

```bash
go build .
```

Go **compiles your Go source code into an executable file**.

On **Windows**, you will normally get:

```text
your-project-name.exe
```

For example:

```text
myapp/
├── go.mod
├── main.go
└── myapp.exe
```

You can then run the `.exe` **without using `go run` or having the source code open**.

### Run the `.exe` on Windows

From PowerShell:

```powershell
.\myapp.exe
```

From Command Prompt:

```cmd
myapp.exe
```

The important idea is:

```text
Go source code
      ↓
go build .
      ↓
Executable (.exe)
      ↓
Run executable
      ↓
Program starts
```

### Why is this useful?

It is commonly used when you want to:

- Run a Go application in production.
- Distribute your application to another computer.
- Run a backend server without running `go run`.
- Create a deployable application.
- Package your Go application for Docker/AWS/server deployment.

---

## 2. Simple example

Create a `main.go`:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello from my Go executable!")
}
```

Initialize the project:

```bash
go mod init example.com/myapp
```

Build it:

```bash
go build .
```

On Windows, you may now see:

```text
myapp.exe
```

Run it:

```powershell
.\myapp.exe
```

Output:

```text
Hello from my Go executable!
```

### You can also choose the `.exe` name

Use:

```bash
go build -o server.exe .
```

Now you have:

```text
server.exe
```

Run it:

```powershell
.\server.exe
```

This is particularly useful for Go backend applications.

---

## 3. Three common mistakes beginners make

### Mistake 1: Using `myapp.exe` without `.\`

In PowerShell, this may not work:

```powershell
myapp.exe
```

Instead use:

```powershell
.\myapp.exe
```

`.\` means **"execute the file from the current directory."**

---

### Mistake 2: Thinking `go build .` always creates `.exe`

The output depends on the operating system.

On Windows:

```text
myapp.exe
```

On Linux:

```text
myapp
```

On macOS:

```text
myapp
```

So `.exe` is specifically associated with Windows executables.

---

### Mistake 3: Thinking the `.exe` requires Go to run

A compiled Go executable generally does **not require the Go compiler/runtime to be installed** on the target machine.

For example, you can build:

```bash
go build -o server.exe .
```

Then copy:

```text
server.exe
```

to another compatible Windows machine and run it there.

However, your application may still depend on external things such as:

- Configuration files
- Environment variables
- Database servers
- External APIs
- Required DLLs in special cases

So the `.exe` is portable, but **the entire application environment may not be**.

---

## 4. Real-world applications

### Application 1: Go REST API server

Suppose you create a REST API:

```text
main.go
   ↓
go build -o api-server.exe .
   ↓
api-server.exe
   ↓
HTTP server starts
   ↓
React / mobile app / other services call the API
```

You can deploy the executable to a Windows server and run:

```powershell
.\api-server.exe
```

For example, your Go application could listen on:

```text
http://localhost:8080
```

---

### Application 2: Deploying a Go application

A typical production workflow can look like:

```text
Developer writes Go code
        ↓
go test
        ↓
go build
        ↓
Executable
        ↓
Copy/deploy executable
        ↓
Run application
```

For example:

```bash
go build -o payment-service.exe .
```

Then the resulting executable can be deployed to the appropriate server/environment.

This is one reason Go is popular for backend and microservice applications: **the build process produces a standalone executable that is easy to distribute and deploy.**

---

## 5. Three exercises

### Exercise 1 — Beginner

Create a Go program that prints:

```text
Welcome to Go Executable!
```

Then:

1. Run it using `go run`.
2. Build it using `go build .`.
3. Find the generated `.exe`.
4. Run the `.exe` directly from PowerShell.
5. Compare the behavior of `go run` and the executable.

---

### Exercise 2 — Intermediate

Create a simple Go HTTP server that listens on port `8080`.

Then:

1. Build it as `server.exe`.
2. Run `server.exe`.
3. Open the API from a browser.
4. Stop the executable.
5. Start it again without using `go run`.

Your goal is to understand that the **compiled executable itself is the application you are running**.

---

### Exercise 3 — Advanced

Create a small Go REST API with multiple endpoints, for example:

```text
GET  /users
GET  /users/{id}
POST /users
```

Then:

1. Build the application with a custom executable name.
2. Run the executable.
3. Test all endpoints.
4. Move the executable to another directory.
5. Run it from there.
6. Determine what happens if the application expects a configuration file.
7. Modify the application so configuration is supplied through environment variables.
8. Build and run the executable again.

This will help you understand the difference between **the executable itself and the external resources your application depends on**.

---

## 🤔 Think deeper

If you build your Go REST API into `server.exe`, **why might you prefer deploying that executable instead of giving the production server your entire Go source-code project?**

Think about **security, deployment speed, dependencies, and production maintenance**.
