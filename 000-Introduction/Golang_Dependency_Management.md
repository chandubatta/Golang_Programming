# Dependency Management in Go

## 1. What is Dependency Management?

**Dependency Management** is the process of adding, downloading,
updating, and controlling external libraries/packages that your Go
application depends on.

For example, instead of writing your own HTTP router, you can use an
existing package such as `github.com/gorilla/mux`.

Go primarily uses **Go Modules** for dependency management.

### Main purpose

Dependency Management helps you:

-   Reuse existing libraries.
-   Keep track of package versions.
-   Make builds reproducible.
-   Avoid manually downloading packages.
-   Update dependencies safely.
-   Share projects with other developers.

### Common Go dependency-management files and commands

``` text
go.mod       → Defines your module and dependencies
go.sum       → Stores checksums for dependency verification
go get       → Adds/updates dependencies
go mod tidy  → Cleans and synchronizes dependencies
go list      → Inspects dependencies
```

### When is it commonly used?

Almost every real-world Go project uses dependency management,
especially when working with:

-   REST APIs
-   Databases
-   Authentication
-   Web frameworks
-   gRPC
-   Kafka
-   Redis
-   Testing libraries
-   Cloud services
-   Microservices

------------------------------------------------------------------------

# 2. Simple Code Example

Let's create a small Go project that uses the external package
`github.com/google/uuid` to generate a UUID.

## Step 1: Create a project

``` bash
mkdir dependency-demo
cd dependency-demo
```

Initialize a Go module:

``` bash
go mod init dependency-demo
```

This creates:

``` text
dependency-demo/
│
├── go.mod
└── main.go
```

## Step 2: Add a dependency

``` bash
go get github.com/google/uuid
```

Go downloads the package and records it in `go.mod`.

## Step 3: Write the program

``` go
package main

import (
    "fmt"

    "github.com/google/uuid"
)

func main() {
    id := uuid.New()

    fmt.Println("Generated ID:", id)
}
```

## Step 4: Run

``` bash
go run .
```

Example output:

``` text
Generated ID: 550e8400-e29b-41d4-a716-446655440000
```

Your `go.mod` will contain a dependency similar to:

``` go
module dependency-demo

go 1.XX

require github.com/google/uuid vX.X.X
```

The `go.sum` file contains checksums used to verify downloaded module
content.

### Important concept

Think about the flow like this:

``` text
Your Go Application
        ↓
     go.mod
        ↓
External Dependency
        ↓
Download specific version
        ↓
go.sum verifies dependency
        ↓
Application builds/runs
```

------------------------------------------------------------------------

# 3. Three Common Mistakes

## Mistake 1: Manually copying external packages

A beginner might download a library manually and put it inside the
project.

### Why this is a problem

It becomes difficult to:

-   Track versions.
-   Update the library.
-   Reproduce the project on another computer.
-   Know which dependencies the project actually needs.

### How to avoid it

Use Go Modules:

``` bash
go mod init
go get
go mod tidy
```

------------------------------------------------------------------------

## Mistake 2: Thinking `go.mod` and `go.sum` are the same

They have different purposes.

### `go.mod`

Defines your module and its dependencies.

``` text
go.mod
   ↓
Which dependencies does my project use?
```

### `go.sum`

Contains cryptographic checksums for module versions.

``` text
go.sum
   ↓
Can I verify the downloaded dependency?
```

**Remember:**

> `go.mod` → dependency information\
> `go.sum` → dependency verification information

------------------------------------------------------------------------

## Mistake 3: Updating every dependency without thinking

A beginner may run dependency updates frequently without checking
compatibility.

For example:

``` bash
go get -u ./...
```

A newer dependency version might introduce:

-   Breaking API changes
-   Behavioral changes
-   New bugs
-   Compatibility problems

### How to avoid it

Before upgrading an important dependency:

1.  Check the current version.
2.  Check the new version.
3.  Read its release notes.
4.  Run your tests.
5.  Verify your application.

Use:

``` bash
go list -m -u all
```

to inspect available updates.

------------------------------------------------------------------------

# 4. Real-World Applications

## Scenario 1: REST API Development

Suppose you're developing a Go REST API.

Your application might depend on:

``` text
Go REST API
│
├── Gin
├── MySQL driver
├── JWT library
├── UUID library
└── Validation library
```

Instead of manually managing all these libraries, Go Modules records
them in `go.mod`.

Another developer can clone the project and run:

``` bash
go mod download
```

or simply:

``` bash
go run .
```

Go can download the required modules automatically.

------------------------------------------------------------------------

## Scenario 2: Microservices

Imagine an e-commerce system:

``` text
                React
                  │
                  ▼
             API Gateway
                  │
        ┌─────────┼─────────┐
        ▼         ▼         ▼
    User Service Product   Payment
        │         │         │
       MySQL     MongoDB   MySQL
```

Each Go microservice can have its own dependencies.

For example:

``` text
User Service
 ├── JWT
 ├── MySQL driver
 └── UUID

Product Service
 ├── MongoDB driver
 ├── gRPC
 └── Protobuf

Payment Service
 ├── MySQL driver
 ├── Kafka
 └── gRPC
```

Dependency Management allows each service to control its own dependency
versions without requiring every service to use exactly the same
libraries.

------------------------------------------------------------------------

# 5. Progressive Exercises

## Exercise 1 --- Beginner

Create a Go project called:

``` text
calculator-app
```

Use Go Modules and add an external dependency that provides a useful
mathematical or utility function.

Your program should:

1.  Initialize a Go module.
2.  Add the external dependency.
3.  Import and use it.
4.  Run the program.
5.  Inspect the generated `go.mod` and `go.sum`.
6.  Remove the dependency from your code and use `go mod tidy`.
7.  Observe what happens to the dependency information.

**Do not manually edit `go.mod` or `go.sum`.**

------------------------------------------------------------------------

## Exercise 2 --- Intermediate

Create a small REST API project called:

``` text
user-api
```

Add at least **two external dependencies**.

Your API should:

1.  Create a Go module.
2.  Add a web framework.
3.  Add a UUID-generation package.
4.  Create a `POST /users` endpoint.
5.  Generate a unique ID for every user.
6.  Return the user information as JSON.
7.  Check which dependencies are recorded in `go.mod`.
8.  Use `go mod tidy`.
9.  Inspect the dependency graph using an appropriate Go command.

------------------------------------------------------------------------

## Exercise 3 --- Advanced

Create a small project with the following structure:

``` text
ecommerce/
│
├── user-service/
├── product-service/
└── payment-service/
```

Each service should be a separate Go module.

Requirements:

-   `user-service` uses a UUID library.
-   `product-service` uses a MongoDB driver.
-   `payment-service` uses a MySQL driver.
-   Each service must manage its dependencies independently.
-   Use different dependency versions where appropriate.
-   Run `go mod tidy` for every service.
-   Inspect the dependency graph of each service.
-   Identify direct and indirect dependencies.
-   Investigate what happens when one dependency is upgraded.
-   Run tests after dependency changes.

**Do not provide yourself with a solution---design the dependency
structure yourself.**

------------------------------------------------------------------------

# 6. Think Deeper: Dependency Management in Production

Imagine you are responsible for a large Go application with **100+
dependencies**.

One of those dependencies suddenly releases a new major version
containing an important security fix, but upgrading it could potentially
break your application.

### Thought-provoking question

**Would you upgrade it immediately, keep the old version, or first
create a controlled upgrade process---and why?**

Consider these points:

1.  Would you upgrade immediately or test first?
2.  How would you verify that the new version doesn't break your
    application?
3.  Would you update `go.mod` manually or use Go commands?
4.  What role would automated tests and CI/CD play?
5.  If the upgrade breaks production, how would you roll back?

Try answering as if you were the **Golang developer responsible for a
production microservice**.

### The deeper question

> **How should a production team balance security, stability,
> reproducibility, and the need to keep dependencies up to date?**
