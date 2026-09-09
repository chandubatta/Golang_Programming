# Go `path/filepath` Package

Official documentation: https://pkg.go.dev/path/filepath

## 1. What is the `path/filepath` package?

The `path/filepath` package provides functions for creating, manipulating, analyzing, and traversing file-system paths.

Its most important feature is that it understands the operating system's path conventions.

For example:

```text
Linux/macOS:
/home/chandu/projects/app/main.go

Windows:
C:\Users\Chandu\projects\app\main.go
```

Instead of manually writing `/` or `\`, use:

```go
filepath.Join("projects", "app", "main.go")
```

Go uses the appropriate path separator for the operating system.

### `path` vs `path/filepath`

| Package | Intended for |
|---|---|
| `path` | Slash-separated paths such as URLs |
| `path/filepath` | Operating-system file-system paths |

Use `path/filepath` when working with actual files and directories on the local operating system.

---

# 2. Important constants and variables

## `filepath.Separator`

Represents the operating system's path separator.

```go
fmt.Println(filepath.Separator)
```

Typically:

```text
Linux/macOS: /
Windows:     \
```

## `filepath.ListSeparator`

Separates multiple paths in an environment path list.

Typically:

```text
Linux/macOS: :
Windows:     ;
```

## `filepath.ErrBadPattern`

An error returned for malformed patterns supplied to functions such as `Match` and `Glob`.

Example:

```go
matches, err := filepath.Glob("[abc")
if err != nil {
    fmt.Println("Invalid pattern:", err)
}
```

## `filepath.SkipDir`

Used with `Walk`/`WalkDir` to tell the walker not to descend into a directory.

## `filepath.SkipAll`

Used to stop a directory walk completely.

---

# 3. Every function in `path/filepath`

## 3.1 `Abs()`

### Signature

```go
filepath.Abs(path string) (string, error)
```

Converts a path into an absolute path.

```go
package main

import (
    "fmt"
    "path/filepath"
)

func main() {
    absolutePath, err := filepath.Abs("documents/report.txt")
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(absolutePath)
}
```

If the current directory is:

```text
/home/chandu/project
```

the result could be:

```text
/home/chandu/project/documents/report.txt
```

`Abs` cleans the resulting path.

### Common uses

- Finding the absolute location of a configuration file
- Logging complete file locations
- Converting relative paths into absolute paths

---

## 3.2 `Base()`

### Signature

```go
filepath.Base(path string) string
```

Returns the last element of a path.

```go
p := "/home/chandu/documents/report.pdf"
fmt.Println(filepath.Base(p))
```

Output:

```text
report.pdf
```

Think:

> "What is the final component of this path?"

---

## 3.3 `Clean()`

### Signature

```go
filepath.Clean(path string) string
```

Cleans a path using lexical rules.

```go
fmt.Println(filepath.Clean("/home/chandu/./documents/../file.txt"))
```

Result:

```text
/home/chandu/file.txt
```

It handles:

- `.`
- `..`
- repeated separators

For example:

```go
filepath.Clean("a//b///c")
```

produces an equivalent cleaned path such as:

```text
a/b/c
```

### Important

`Clean()` performs lexical processing. It does **not** inspect the actual filesystem and does not resolve symbolic links.

---

## 3.4 `Dir()`

### Signature

```go
filepath.Dir(path string) string
```

Returns the directory portion of a path.

```go
p := "/home/chandu/documents/report.pdf"
fmt.Println(filepath.Dir(p))
```

Output:

```text
/home/chandu/documents
```

---

## 3.5 `EvalSymlinks()`

### Signature

```go
filepath.EvalSymlinks(path string) (string, error)
```

Evaluates symbolic links in a path.

Suppose:

```text
real-file.txt
current.txt -> real-file.txt
```

Then:

```go
resolved, err := filepath.EvalSymlinks("current.txt")
```

can resolve the link to its target.

### Important distinction

`Clean()`:

```text
purely lexical
```

`EvalSymlinks()`:

```text
interacts with the filesystem
```

Therefore `EvalSymlinks` can return an error when a path cannot be resolved.

---

## 3.6 `Ext()`

### Signature

```go
filepath.Ext(path string) string
```

Returns the file-name extension.

```go
fmt.Println(filepath.Ext("report.pdf"))
```

Output:

```text
.pdf
```

For:

```text
archive.tar.gz
```

the result is:

```text
.gz
```

The returned extension includes the `.`.

---

## 3.7 `FromSlash()`

### Signature

```go
filepath.FromSlash(path string) string
```

Converts `/` characters into the operating system's path separator.

For example, on Windows:

```go
p := filepath.FromSlash("documents/reports/2026/report.pdf")
```

can produce:

```text
documents\reports\2026\report.pdf
```

Useful when receiving slash-separated paths and converting them to local filesystem paths.

---

## 3.8 `Glob()`

### Signature

```go
filepath.Glob(pattern string) ([]string, error)
```

Finds files whose names match a pattern.

```go
files, err := filepath.Glob("*.go")
if err != nil {
    fmt.Println(err)
    return
}

for _, file := range files {
    fmt.Println(file)
}
```

Possible matching files:

```text
main.go
server.go
```

Common patterns:

```text
*.go
*.json
*.txt
image?.png
```

### Important

`Glob()` does not automatically recursively search every subdirectory.

Malformed patterns produce an error such as `ErrBadPattern`.

---

## 3.9 `HasPrefix()` — Deprecated

### Signature

```go
filepath.HasPrefix(p, prefix string) bool
```

This function is deprecated.

It does not correctly handle path boundaries and platform-specific case rules.

Do not use it for security-sensitive path validation.

Depending on your requirement, consider `filepath.IsLocal()` or appropriate path-boundary logic instead.

---

## 3.10 `IsAbs()`

### Signature

```go
filepath.IsAbs(path string) bool
```

Determines whether a path is absolute.

```go
fmt.Println(filepath.IsAbs("/home/chandu/file.txt"))
```

On Unix-like systems:

```text
true
```

And:

```go
fmt.Println(filepath.IsAbs("documents/file.txt"))
```

returns:

```text
false
```

### Important

Windows has different path semantics from Unix-like systems, so don't assume all absolute-path rules are identical across operating systems.

---

## 3.11 `IsLocal()`

### Signature

```go
filepath.IsLocal(path string) bool
```

Determines whether a path is lexically local.

A local path:

- is not absolute
- is not empty
- does not escape its starting directory using `..`
- is not a reserved Windows name such as `NUL`

Example:

```go
fmt.Println(filepath.IsLocal("documents/report.txt"))
```

Result:

```text
true
```

But:

```go
fmt.Println(filepath.IsLocal("../secret.txt"))
```

returns:

```text
false
```

And:

```go
fmt.Println(filepath.IsLocal("/etc/passwd"))
```

returns:

```text
false
```

### Security relevance

`IsLocal` can be useful when validating user-supplied paths.

### Important limitation

`IsLocal` checks lexical locality. It does **not** account for symbolic links. A symbolic link can introduce additional filesystem-level security concerns.

---

## 3.12 `Join()`

### Signature

```go
filepath.Join(elem ...string) string
```

Joins multiple path components together.

```go
path := filepath.Join(
    "home",
    "chandu",
    "documents",
    "report.pdf",
)

fmt.Println(path)
```

On Unix:

```text
home/chandu/documents/report.pdf
```

On Windows:

```text
home\chandu\documents\report.pdf
```

`Join()` also cleans the resulting path.

For example:

```go
filepath.Join("home", "chandu", "..", "documents")
```

produces a cleaned equivalent of:

```text
home/documents
```

### Why use `Join()`?

Avoid:

```go
path := directory + "/" + filename
```

Prefer:

```go
path := filepath.Join(directory, filename)
```

This makes filesystem paths portable across operating systems.

---

## 3.13 `Localize()`

### Signature

```go
filepath.Localize(path string) (string, error)
```

`Localize` converts a slash-separated path into an operating-system-specific local path.

Example:

```go
localPath, err := filepath.Localize("documents/report.txt")
```

On Unix:

```text
documents/report.txt
```

On Windows:

```text
documents\report.txt
```

The input must be a valid `io/fs.ValidPath`. The function can reject paths that cannot safely be represented by the target operating system.

### Why useful?

It is useful when working with APIs such as `io/fs`, where paths conventionally use `/`, but local filesystem operations require OS-native paths.

---

## 3.14 `Match()`

### Signature

```go
filepath.Match(pattern, name string) (bool, error)
```

Checks whether a filename matches a shell-style pattern.

```go
matched, err := filepath.Match("*.go", "main.go")

fmt.Println(matched)
fmt.Println(err)
```

Output:

```text
true
<nil>
```

Common pattern characters:

```text
*     any sequence of non-separator characters
?     exactly one non-separator character
[ ]   character class
```

Examples:

```go
filepath.Match("*.go", "main.go")   // true
filepath.Match("*.go", "main.txt")  // false
```

`Match()` checks the whole name.

Malformed patterns return an error such as `ErrBadPattern`.

---

## 3.15 `Rel()`

### Signature

```go
filepath.Rel(basePath, targPath string) (string, error)
```

Calculates a path relative to another path.

Suppose:

```text
base:
    /home/chandu/project

target:
    /home/chandu/project/images/logo.png
```

Then:

```go
relative, err := filepath.Rel(
    "/home/chandu/project",
    "/home/chandu/project/images/logo.png",
)
```

produces:

```text
images/logo.png
```

Another example:

```text
base:
    /home/chandu/project

target:
    /home/chandu/config.json
```

produces:

```text
../config.json
```

### Useful for

- Generating relative links
- Comparing locations
- Build tools
- Project file organization
- Backup software

---

## 3.16 `Split()`

### Signature

```go
filepath.Split(path string) (dir, file string)
```

Splits a path into directory and file components.

```go
dir, file := filepath.Split(
    "/home/chandu/documents/report.pdf",
)

fmt.Println("Directory:", dir)
fmt.Println("File:", file)
```

Conceptually:

```text
Directory: /home/chandu/documents/
File:      report.pdf
```

`Split()` is useful when you need both components at once.

---

## 3.17 `SplitList()`

### Signature

```go
filepath.SplitList(path string) []string
```

Splits a list of paths using the operating system's path-list separator.

This is useful for environment variables such as `PATH`.

Unix example:

```text
/usr/bin:/usr/local/bin:/home/chandu/bin
```

Windows example:

```text
C:\Windows;C:\Program Files
```

Example:

```go
paths := filepath.SplitList(os.Getenv("PATH"))

for _, p := range paths {
    fmt.Println(p)
}
```

---

## 3.18 `ToSlash()`

### Signature

```go
filepath.ToSlash(path string) string
```

Converts OS-specific path separators into `/`.

For example, on Windows:

```text
C:\Users\Chandu\project\main.go
```

becomes:

```text
C:/Users/Chandu/project/main.go
```

Useful when creating:

- URLs
- archive entries
- configuration formats
- slash-separated filesystem representations

---

## 3.19 `VolumeName()`

### Signature

```go
filepath.VolumeName(path string) string
```

Returns the volume portion of a path.

On Windows:

```go
filepath.VolumeName(`C:\Users\Chandu`)
```

returns:

```text
C:
```

For a UNC path:

```text
\\server\share\folder
```

it can return:

```text
\\server\share
```

On Unix-like systems, the result is normally:

```text
""
```

This is mainly useful for code that needs to understand Windows filesystem paths.

---

## 3.20 `Walk()`

### Signature

```go
filepath.Walk(root string, fn filepath.WalkFunc) error
```

Walks through a directory tree recursively.

Suppose:

```text
project/
├── main.go
├── go.mod
├── internal/
│   ├── auth.go
│   └── database.go
└── tests/
    └── api_test.go
```

You can traverse it:

```go
err := filepath.Walk("project", func(
    path string,
    info os.FileInfo,
    err error,
) error {

    if err != nil {
        return err
    }

    fmt.Println(path)

    return nil
})

if err != nil {
    fmt.Println(err)
}
```

The callback receives:

```text
path
info
err
```

`info` can provide:

```go
info.IsDir()
info.Name()
info.Size()
info.Mode()
```

### Skipping directories

```go
if info.IsDir() && info.Name() == ".git" {
    return filepath.SkipDir
}
```

### Important

`Walk`:

- walks recursively
- visits the root
- does not follow symbolic links
- walks in lexical order
- is generally less efficient than `WalkDir` because `WalkDir` avoids an `os.Lstat` call for every visited item

---

## 3.21 `WalkDir()`

### Signature

```go
filepath.WalkDir(root string, fn fs.WalkDirFunc) error
```

`WalkDir` is the newer and generally preferred approach when you don't need the exact behavior of `Walk`.

Example:

```go
package main

import (
    "fmt"
    "io/fs"
    "path/filepath"
)

func main() {
    err := filepath.WalkDir(".", func(
        path string,
        d fs.DirEntry,
        err error,
    ) error {

        if err != nil {
            return err
        }

        fmt.Println(path)

        return nil
    })

    if err != nil {
        fmt.Println(err)
    }
}
```

You can inspect entries efficiently:

```go
if d.IsDir() {
    fmt.Println("Directory:", path)
} else {
    fmt.Println("File:", path)
}
```

### `Walk` vs `WalkDir`

| Feature | `Walk` | `WalkDir` |
|---|---|---|
| Recursive traversal | Yes | Yes |
| Root included | Yes | Yes |
| Follows symlinks | No | No |
| Callback information | `fs.FileInfo` | `fs.DirEntry` |
| Efficiency | Lower | Better |
| Introduced | Older | Go 1.16 |
| Usually preferred for new code | Sometimes | Yes |

---

# 4. `WalkFunc`

`WalkFunc` is a function type, not a function itself.

```go
type WalkFunc func(
    path string,
    info fs.FileInfo,
    err error,
) error
```

You provide a function matching this signature to `filepath.Walk`.

Example:

```go
func visit(
    path string,
    info fs.FileInfo,
    err error,
) error {

    if err != nil {
        return err
    }

    fmt.Println(path)

    return nil
}
```

Then:

```go
filepath.Walk(".", visit)
```

---

# 5. Simple complete example

This example combines several important `path/filepath` functions:

```go
package main

import (
    "fmt"
    "path/filepath"
)

func main() {
    directory := "documents"
    filename := "report.pdf"

    // Build a platform-independent filesystem path.
    fullPath := filepath.Join(directory, filename)

    fmt.Println("Full path:", fullPath)

    // Get the directory.
    fmt.Println("Directory:", filepath.Dir(fullPath))

    // Get the filename.
    fmt.Println("Base:", filepath.Base(fullPath))

    // Get the extension.
    fmt.Println("Extension:", filepath.Ext(fullPath))

    // Check whether the path is absolute.
    fmt.Println("Absolute:", filepath.IsAbs(fullPath))

    // Convert to an absolute path.
    absolute, err := filepath.Abs(fullPath)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Absolute path:", absolute)
}
```

The conceptual flow is:

```text
"documents"
      +
"report.pdf"
      |
      v
filepath.Join()
      |
      v
documents/report.pdf
      |
      +------------------+
      |                  |
    Dir()              Base()
      |                  |
  documents          report.pdf
                         |
                        Ext()
                         |
                        .pdf
```

---

# 6. Three common beginner mistakes

## Mistake 1: Manually constructing paths

Beginners often write:

```go
path := folder + "/" + filename
```

Prefer:

```go
path := filepath.Join(folder, filename)
```

### Rule

> Use `filepath.Join()` for filesystem paths.

---

## Mistake 2: Confusing `path` with `path/filepath`

Use:

```go
import "path"
```

for slash-separated paths such as URL-style paths.

Use:

```go
import "path/filepath"
```

for actual filesystem paths.

They are related but serve different purposes.

---

## Mistake 3: Assuming `Clean()` makes a path safe

A beginner may think:

```go
clean := filepath.Clean(userInput)
```

means the path is now safe.

Not necessarily.

For example:

```text
../../secret.txt
```

can still escape the intended directory.

For lexical local-path validation, `filepath.IsLocal()` is useful. However, `IsLocal` does not account for symbolic links.

For serious security requirements, path validation must be designed together with the filesystem access model.

---

# 7. Two real-world applications

## Application 1: File upload system

Imagine a web server receiving:

```text
profile.jpg
```

from a user.

You might construct:

```go
uploadPath := filepath.Join(
    "uploads",
    userID,
    filename,
)
```

You can use:

```go
filepath.Ext(filename)
```

to determine the extension.

And:

```go
filepath.IsLocal(filename)
```

can help reject obviously unsafe lexical paths.

This kind of path handling is important because accepting arbitrary filesystem paths from users can create path-traversal vulnerabilities.

---

## Application 2: Build tools / backup programs

Imagine:

```text
project/
├── main.go
├── config.json
├── internal/
│   ├── auth.go
│   └── database.go
└── README.md
```

A backup program can use:

```go
filepath.WalkDir()
```

to discover files.

Then:

```go
filepath.Ext()
```

can identify file types.

And:

```go
filepath.Rel()
```

can calculate each file's location relative to the project root.

This pattern is common in:

- backup tools
- compilers
- code generators
- file indexers
- CLI tools
- deployment tools

---

# 8. Three progressively challenging exercises

## Exercise 1 — Beginner: File Path Analyzer

Write a Go program that accepts a file path from the user and prints:

```text
Original path:
Base name:
Directory:
Extension:
Is absolute:
Clean path:
```

Use appropriate functions from `path/filepath`.

Test your program with:

```text
documents/reports/../2026/report.pdf
```

and with an absolute path.

Do not use manual `/` or `\` manipulation.

---

## Exercise 2 — Intermediate: File Extension Scanner

Create a program that receives a directory path and searches for files with a user-selected extension.

For example:

```text
Directory: project
Extension: .go
```

The program should recursively search the directory and print every matching Go source file.

### Requirements

- Use `filepath.WalkDir()`.
- Ignore directories named `.git`.
- Only print regular files.
- Use `filepath.Ext()` to determine the extension.
- Count the number of matching files.
- Handle filesystem errors properly.

Example conceptual output:

```text
Found:
project/main.go
project/server.go
project/internal/auth.go

Total: 3
```

Do not use `filepath.Glob()` for the recursive traversal.

---

## Exercise 3 — Advanced: Secure File Server Path Resolver

Build a program that simulates resolving user-requested files inside a fixed directory:

```text
server_files/
```

The user supplies a path such as:

```text
documents/report.pdf
```

Your program should determine whether the requested path can safely be treated as a local path before constructing the final filesystem path.

Test it with:

```text
documents/report.pdf
images/logo.png
../secret.txt
../../etc/passwd
/absolute/path/file.txt
```

### Requirements

1. Validate the supplied path.
2. Reject paths that are not local.
3. Convert/construct the filesystem path correctly.
4. Prevent the requested path from escaping the intended root.
5. Print the final resolved path when accepted.
6. Explain why a rejected path was rejected.

Then think carefully about symbolic links and whether lexical validation alone is sufficient for your security requirements.

---

# 9. Recommended learning order

Instead of memorizing every function independently, learn them in groups.

## Building paths

```text
Join()
Clean()
Abs()
Rel()
```

## Extracting information

```text
Base()
Dir()
Ext()
Split()
VolumeName()
```

## Checking paths

```text
IsAbs()
IsLocal()
```

## Converting paths

```text
ToSlash()
FromSlash()
Localize()
```

## Finding/matching

```text
Match()
Glob()
```

## Filesystem resolution

```text
EvalSymlinks()
```

## Directory traversal

```text
Walk()
WalkDir()
```

## Environment path lists

```text
SplitList()
```

## Deprecated

```text
HasPrefix()
```

### The most important everyday functions

Start with:

```go
filepath.Join()
filepath.Base()
filepath.Dir()
filepath.Ext()
filepath.Clean()
filepath.Abs()
filepath.WalkDir()
```

For security-conscious code, also learn:

```go
filepath.IsLocal()
```

---

# 10. Thought-provoking question

Suppose you're building a Go web server that lets users download files from a directory called `uploads/`.

A user sends:

```text
../../secret/passwords.txt
```

You use:

```go
filepath.Clean(userInput)
```

and get:

```text
../../secret/passwords.txt
```

Then you use:

```go
filepath.Join("uploads", userInput)
```

### Deeper question

Why isn't `filepath.Clean()` enough to make this secure, and how would symbolic links make the problem even more interesting—even if `filepath.IsLocal()` says that a path is safe?

Think about:

> **lexical paths vs. actual filesystem objects**

before answering.
