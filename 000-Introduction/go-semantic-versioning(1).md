# Go Follows Semantic Versioning (SemVer)

## 1. Concise Explanation

**Semantic Versioning (SemVer)** is a standard way of numbering software versions so developers can understand the significance of changes between versions.

The general format is:

```text
MAJOR.MINOR.PATCH
```

For example:

```text
v2.5.3
```

means:

- **MAJOR (`2`)** → incompatible/breaking API changes
- **MINOR (`5`)** → new backward-compatible functionality
- **PATCH (`3`)** → backward-compatible bug fixes

### In Go

You commonly see versions such as:

```text
go 1.23
```

in `go.mod`, and module versions such as:

```text
github.com/example/mymodule v1.4.2
```

Go's module system uses semantic-version concepts to manage dependencies and determine whether a dependency update is compatible.

### Why is it useful?

Suppose your project uses:

```text
github.com/example/payment v1.4.2
```

A newer version becomes available:

```text
v1.4.3
```

The PATCH change generally indicates a bug fix without intentionally breaking the existing API.

But if the library releases:

```text
v2.0.0
```

you should expect potentially breaking changes and review the migration requirements.

---

## 2. Simple Go Example

Imagine you have a Go module with this `go.mod`:

```go
module example.com/myapp

go 1.23

require (
    github.com/example/payment v1.4.2
)
```

Here:

```text
module example.com/myapp
```

is your module path.

```text
go 1.23
```

specifies the Go language/toolchain version requirement for the module.

```text
github.com/example/payment v1.4.2
```

specifies a dependency and its version.

The dependency version follows:

```text
MAJOR.MINOR.PATCH
   ↓     ↓     ↓
   1     4     2
```

### Example version changes

```text
v1.4.2 → v1.4.3
```

Usually:

```text
PATCH change
```

Bug fixes.

```text
v1.4.3 → v1.5.0
```

Usually:

```text
MINOR change
```

New backward-compatible functionality.

```text
v1.5.0 → v2.0.0
```

Usually:

```text
MAJOR change
```

Potentially breaking changes.

You can inspect dependency versions with:

```bash
go list -m all
```

and update dependencies with:

```bash
go get -u
```

---

## 3. Three Common Mistakes

### Mistake 1: Thinking every new version is a breaking change

A beginner might see:

```text
v1.2.0 → v1.3.0
```

and assume the API has changed incompatibly.

**Avoid it:** Understand the SemVer levels:

```text
MAJOR → potentially breaking
MINOR → new compatible features
PATCH → bug fixes
```

---

### Mistake 2: Confusing Go's version with a module's version

For example:

```text
go 1.23
```

and:

```text
github.com/example/project v1.5.2
```

are not the same thing.

The first relates to the Go version/toolchain requirement declared by the module, while the second is the version of a dependency/module.

**Avoid it:** Think of them as two separate version numbers:

```text
Go version       → 1.23
Dependency       → v1.5.2
```

---

### Mistake 3: Assuming a major version can always be upgraded automatically

For example:

```text
v1.8.0 → v2.0.0
```

may introduce breaking API changes.

Go modules have special handling for major versions `v2` and above. A module's major version is normally reflected in its module path, such as:

```text
example.com/library/v2
```

**Avoid it:** When moving to a new major version, read the library's migration/change notes and check the new module path.

---

## 4. Two Real-World Applications

### Application 1: Managing Go dependencies

Imagine an e-commerce backend:

```text
my-ecommerce/
├── go.mod
├── go.sum
├── cmd/
├── product/
├── payment/
└── user/
```

Your `go.mod` might contain:

```go
require (
    github.com/gorilla/mux v1.8.1
)
```

When a newer compatible version is released, versioning helps you understand what type of change you're accepting.

This is particularly useful when maintaining production applications.

---

### Application 2: Publishing your own Go library

Suppose you create:

```text
github.com/chandu/mylogger
```

Your releases could be:

```text
v1.0.0
v1.0.1
v1.1.0
v2.0.0
```

You could interpret them as:

```text
v1.0.0 → First stable release

v1.0.1 → Bug fix

v1.1.0 → New backward-compatible feature

v2.0.0 → Breaking API change
```

This gives users a predictable way to decide whether they can safely upgrade.

---

## 5. Progressive Exercises

### Exercise 1 — Beginner

Create a Go module named:

```text
example.com/semver-demo
```

Use a `go.mod` file and document three hypothetical releases:

```text
v1.0.0
v1.0.1
v1.1.0
```

For each version, write a short description explaining what type of change it represents according to Semantic Versioning.

---

### Exercise 2 — Intermediate

Create a small Go library containing a function:

```go
CalculateTotal()
```

Pretend you publish it as:

```text
v1.0.0
```

Then design hypothetical changes for:

```text
v1.0.1
v1.1.0
v2.0.0
```

For each version, describe a change that would justify that particular version number.

Do **not** actually change the API for this exercise; focus on deciding which version number is appropriate for each type of change.

---

### Exercise 3 — Advanced

Create a small Go project with at least two dependencies in `go.mod`.

Design a dependency upgrade scenario involving:

```text
PATCH upgrade
MINOR upgrade
MAJOR upgrade
```

For each scenario:

1. Identify the old version.
2. Identify the new version.
3. Explain what kind of change you expect.
4. Decide whether you would upgrade immediately in a production application.
5. Explain what you would test before deploying the upgrade.

---

## Important Takeaway

Remember the basic pattern:

```text
        MAJOR.MINOR.PATCH
          ↓     ↓     ↓
          2     4     1
```

Think:

```text
MAJOR  → Breaking/incompatible changes
MINOR  → New compatible features
PATCH  → Bug fixes
```

And in Go, don't confuse:

```text
go 1.23
```

with:

```text
some-module v1.4.2
```

They represent different versioning concepts within the Go ecosystem.

---

## 🤔 Think Deeper

Suppose your production Go application depends on a library at **`v1.8.2`**, and the library releases **`v1.9.0`** with an important new feature.

**Would you automatically upgrade to `v1.9.0` just because Semantic Versioning says it should be backward-compatible, or would you still test it first? Why?**
