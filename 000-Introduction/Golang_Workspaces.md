# Golang Workspaces (`go.work`)

## Table of Contents

1.  Introduction
2.  What is a Workspace?
3.  Purpose of Workspaces
4.  When Should You Use a Workspace?
5.  Module vs Workspace
6.  How Workspaces Work
7.  Workspace Directory Structure
8.  Creating a Workspace
9.  Understanding the `go.work` File
10. Useful Workspace Commands
11. Simple Example
12. Common Beginner Mistakes
13. Real-World Applications
14. Practice Exercises
15. Summary
16. Thought-Provoking Question

------------------------------------------------------------------------

# 1. Introduction

As Go applications become larger, developers often split their projects
into multiple **Go Modules**. These modules may represent:

-   Microservices
-   Shared libraries
-   SDKs
-   Internal tools
-   APIs

Managing multiple modules during development can become difficult
because every change in one module usually requires updating dependency
versions in another.

To solve this problem, Go introduced **Workspaces** in **Go 1.18**.

------------------------------------------------------------------------

# 2. What is a Workspace?

A **Workspace** is a collection of one or more Go modules that are
developed together.

It is managed by a file named:

``` text
go.work
```

Instead of downloading dependencies from GitHub or another repository,
Go uses the **local modules** listed inside the workspace.

Think of it like this:

    Workspace
    ├── Module A
    ├── Module B
    ├── Module C
    └── Module D

All modules can communicate using the latest local code.

------------------------------------------------------------------------

# 3. Purpose of Workspaces

The main purpose of Workspaces is to simplify local development of
multiple Go modules.

Without Workspaces:

``` text
Change Library
        │
Commit → Push → Version → go get
```

With Workspaces:

``` text
Change Library
        │
Run Application Immediately
```

No commits, version updates, or publishing are required for local
testing.

------------------------------------------------------------------------

# 4. When Should You Use a Workspace?

Use Workspaces when developing:

-   Multiple microservices
-   Shared libraries
-   Internal company packages
-   SDKs
-   Monorepositories
-   Backend services that depend on one another

------------------------------------------------------------------------

# 5. Module vs Workspace

  Feature      Module (`go.mod`)     Workspace (`go.work`)
  ------------ --------------------- ----------------------------
  Purpose      Defines one project   Connects multiple projects
  Required     Yes                   Optional
  File Name    `go.mod`              `go.work`
  Contains     Dependencies          List of Modules
  Introduced   Go 1.11               Go 1.18

> Every Workspace contains Modules, and every Module has its own
> `go.mod`.

------------------------------------------------------------------------

# 6. How Workspaces Work

Without Workspace:

``` text
Application
    │
GitHub Module
```

With Workspace:

``` text
go.work
   │
App Module ───── Shared Module
```

Go automatically uses the local version.

------------------------------------------------------------------------

# 7. Workspace Directory Structure

``` text
projects/
├── go.work
├── user-service/
│   ├── go.mod
│   └── main.go
├── payment-service/
│   ├── go.mod
│   └── main.go
└── shared/
    ├── go.mod
    └── utils.go
```

------------------------------------------------------------------------

# 8. Creating a Workspace

``` bash
go work init ./user-service ./payment-service ./shared
```

------------------------------------------------------------------------

# 9. Understanding the `go.work` File

``` go
go 1.25

use (
    ./user-service
    ./payment-service
    ./shared
)
```

The `use` block tells Go which local modules belong to the workspace.

------------------------------------------------------------------------

# 10. Useful Workspace Commands

  Command                   Purpose
  ------------------------- --------------------------
  `go work init`            Create a workspace
  `go work use`             Add a module
  `go work edit -dropuse`   Remove a module
  `go work sync`            Synchronize dependencies
  `go work vendor`          Vendor dependencies
  `go env GOWORK`           Show active workspace

------------------------------------------------------------------------

# 11. Simple Example

## shared/math.go

``` go
package shared

func Add(a, b int) int {
    return a + b
}
```

## app/main.go

``` go
package main

import (
    "fmt"
    "github.com/chandu/shared"
)

func main() {
    fmt.Println(shared.Add(10, 20))
}
```

Create the workspace:

``` bash
go work init ./app ./shared
```

Run:

``` bash
go run .
```

Output:

``` text
30
```

------------------------------------------------------------------------

# 12. Common Beginner Mistakes

1.  **Thinking `go.work` replaces `go.mod`.**
    -   Every module still requires its own `go.mod`.
2.  **Forgetting to add new modules.**
    -   Run:

        ``` bash
        go work use ./new-module
        ```
3.  **Committing `go.work` without team agreement.**
    -   Follow your team's repository policy.

------------------------------------------------------------------------

# 13. Real-World Applications

## Microservices

-   User Service
-   Order Service
-   Payment Service
-   Inventory Service
-   Shared Library

## Shared Company Libraries

-   Logger
-   Authentication
-   Database
-   Configuration
-   Utilities

------------------------------------------------------------------------

# 14. Practice Exercises

## Exercise 1 (Beginner)

Create two modules:

-   calculator
-   app

Use a workspace so the app imports the local calculator module.

## Exercise 2 (Intermediate)

Create:

-   auth
-   database
-   api

Connect them with a workspace and verify local changes are reflected
immediately.

## Exercise 3 (Advanced)

Build a workspace containing:

-   user-service
-   order-service
-   payment-service
-   shared

Organize them with a single `go.work`.

------------------------------------------------------------------------

# 15. Summary

-   Workspace groups multiple Go modules.
-   Uses the `go.work` file.
-   Every module still requires a `go.mod`.
-   Ideal for microservices, monorepos, and shared libraries.
-   Enables local development without publishing dependencies.

------------------------------------------------------------------------

# 16. Thought-Provoking Question

If your organization has 10 microservices maintained by different teams,
would you use one large workspace or multiple smaller workspaces?
Consider development speed, dependency management, team collaboration,
and maintainability before deciding.
