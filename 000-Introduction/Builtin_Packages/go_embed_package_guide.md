# Go `embed` Package

The Go `embed` package provides a way to **include files and directories directly inside a compiled Go binary**. It was introduced in **Go 1.16** and is especially useful for applications that need static assets without requiring those assets to be present as separate files at runtime.

A key idea is:

> **The files are included at build time, then accessed at runtime.**

This makes `embed` particularly useful for web applications, CLI tools, configuration templates, SQL files, HTML/CSS/JavaScript assets, and other resources that should travel with the executable.

---

## 1. What is the `embed` package?

The package is imported as:

```go
import "embed"
```

or, when you only need the `//go:embed` directive for a `string` or `[]byte`:

```go
import _ "embed"
```

The actual embedding is performed by the special compiler directive:

```go
//go:embed filename
```

The directive can initialize one of three kinds of variables:

1. `string`
2. `[]byte`
3. `embed.FS`

For example:

```go
//go:embed message.txt
var message string
```

The contents of `message.txt` become the value of `message`.

For multiple files:

```go
//go:embed templates/*
var templates embed.FS
```

Now `templates` behaves like a **read-only filesystem** containing the embedded files.

### Why use `embed`?

Without `embed`, you might have:

```text
myapp/
├── myapp
├── index.html
├── style.css
└── script.js
```

Your program depends on those external files.

With `embed`:

```text
myapp
```

The HTML, CSS, JavaScript, etc. can be compiled into the executable itself.

This is particularly convenient when distributing applications.

---

# 2. The three ways to embed data

## A. Embed into a `string`

Use this when you have **one text file**.

```go
package main

import (
	_ "embed"
	"fmt"
)

//go:embed message.txt
var message string

func main() {
	fmt.Println(message)
}
```

Suppose `message.txt` contains:

```text
Hello from an embedded file!
```

The compiled program can access the contents without opening `message.txt` from the operating system.

A `string` variable can only be used with **one pattern matching one file**.

---

## B. Embed into `[]byte`

This is useful when you want the raw bytes.

```go
package main

import (
	_ "embed"
	"fmt"
)

//go:embed data.bin
var data []byte

func main() {
	fmt.Println(len(data))
}
```

This can be useful for binary data such as:

- images
- fonts
- certificates
- binary resources
- other non-text files

Again, a `[]byte` embedding uses a single pattern matching a single file.

---

## C. Embed into `embed.FS`

This is the most powerful option when you have **multiple files or directories**.

```go
package main

import (
	"embed"
	"fmt"
)

//go:embed templates/*
var templates embed.FS

func main() {
	data, err := templates.ReadFile("templates/index.html")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
```

The embedded filesystem might look like:

```text
templates/
├── index.html
├── about.html
└── contact.html
```

`embed.FS` implements the standard `io/fs.FS` interface, meaning it can work with other Go packages that understand filesystem abstractions.

---

# 3. Important: `//go:embed` is a compiler directive

This:

```go
//go:embed hello.txt
var message string
```

is **not an ordinary comment**.

The Go toolchain recognizes it specially.

The file is read during the build process and its contents are included in the executable.

The directive must immediately precede the variable declaration:

```go
//go:embed hello.txt
var message string
```

You cannot put arbitrary code between them.

Also, embedded variables must be declared at **package scope**, not inside a function.

Invalid:

```go
func main() {
	//go:embed hello.txt
	var message string
}
```

Correct:

```go
//go:embed hello.txt
var message string

func main() {
	fmt.Println(message)
}
```

---

# 4. `embed.FS`

`embed.FS` represents a **read-only collection of embedded files**.

Conceptually:

```text
embed.FS
   │
   ├── templates/
   │     ├── index.html
   │     └── about.html
   │
   ├── static/
   │     ├── style.css
   │     └── app.js
   │
   └── config.json
```

You can access those files using the methods provided by `embed.FS`.

The package itself is intentionally small. It has **no standalone package-level functions**; its main API is the `FS` type and its three methods: `Open`, `ReadDir`, and `ReadFile`.

Let's examine **every function/method provided by the package**.

---

# 5. Every function/method in the `embed` package

## `FS.Open`

Signature:

```go
func (f FS) Open(name string) (fs.File, error)
```

`Open` opens an embedded file and returns an `fs.File`.

Example:

```go
package main

import (
	"embed"
	"fmt"
)

//go:embed hello.txt
var files embed.FS

func main() {
	file, err := files.Open("hello.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		panic(err)
	}

	fmt.Println("Name:", info.Name())
}
```

### What does it do?

Think of:

```go
files.Open("hello.txt")
```

as conceptually similar to:

```go
os.Open("hello.txt")
```

The major difference is that `embed.FS` accesses the **embedded copy**, rather than opening the file from the operating system's filesystem.

The returned value is an `fs.File`, which provides operations such as:

```go
Read()
Close()
Stat()
```

For regular files, the returned file also implements `io.Seeker`.

### When should you use `Open`?

Use `Open` when you want a file-like interface rather than simply getting the entire contents at once.

For example, it can be useful when passing the file to code that expects an `fs.File`.

---

# 6. `FS.ReadFile`

Signature:

```go
func (f FS) ReadFile(name string) ([]byte, error)
```

This reads the **entire contents** of an embedded file and returns them as `[]byte`.

Example:

```go
package main

import (
	"embed"
	"fmt"
)

//go:embed message.txt
var files embed.FS

func main() {
	data, err := files.ReadFile("message.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
```

If `message.txt` contains:

```text
Hello, Go!
```

the output is:

```text
Hello, Go!
```

### Why is `ReadFile` convenient?

Instead of:

```go
file, err := files.Open("message.txt")
if err != nil {
	panic(err)
}
defer file.Close()

data, err := io.ReadAll(file)
if err != nil {
	panic(err)
}
```

you can simply write:

```go
data, err := files.ReadFile("message.txt")
```

### When should you use it?

Use `ReadFile` when you want the **complete contents** of an embedded file.

Common examples:

```text
config.json
schema.sql
email.html
template.txt
README.md
```

---

# 7. `FS.ReadDir`

Signature:

```go
func (f FS) ReadDir(name string) ([]fs.DirEntry, error)
```

`ReadDir` reads an embedded directory and returns its entries.

Suppose:

```text
assets/
├── index.html
├── style.css
└── app.js
```

You could do:

```go
package main

import (
	"embed"
	"fmt"
)

//go:embed assets/*
var files embed.FS

func main() {
	entries, err := files.ReadDir("assets")
	if err != nil {
		panic(err)
	}

	for _, entry := range entries {
		fmt.Println(entry.Name())
	}
}
```

Possible output:

```text
app.js
index.html
style.css
```

Each `fs.DirEntry` provides information about the directory entry.

For example:

```go
entry.Name()
entry.IsDir()
entry.Type()
entry.Info()
```

### When should you use `ReadDir`?

Use it when you need to discover what files exist inside an embedded directory.

For example:

- listing embedded templates
- discovering plugins/resources
- inspecting static assets
- building an index of embedded files

---

# 8. There are no package-level functions

This is an important beginner point.

You might expect:

```go
embed.ReadFile(...)
```

or:

```go
embed.Open(...)
```

But those **do not exist**.

The package provides the `FS` type:

```go
embed.FS
```

and its methods:

```text
FS.Open()
FS.ReadFile()
FS.ReadDir()
```

The package documentation's function section is empty; these three operations belong to the `FS` type.

Also note that `io/fs` provides general filesystem functions such as:

```go
fs.ReadFile()
fs.ReadDir()
fs.WalkDir()
```

These are functions from **`io/fs`**, not from `embed`.

---

# 9. Embedding an entire directory

You can embed a directory:

```go
package main

import (
	"embed"
	"fmt"
)

//go:embed templates
var templates embed.FS

func main() {
	data, err := templates.ReadFile("templates/index.html")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
```

A directory pattern recursively embeds its files.

There is an important detail: when embedding a directory normally, files beginning with `.` or `_` are excluded.

For example:

```text
templates/
├── index.html
├── _partial.html
└── .hidden.html
```

The normal directory embedding does not include the latter two.

If you specifically need those files, the `all:` prefix can change this behavior:

```go
//go:embed all:templates
var templates embed.FS
```

---

# 10. Real-world example: Embedded web application

One of the most useful applications is serving a web application's static files directly from the executable.

Project:

```text
myapp/
├── main.go
└── static/
    ├── index.html
    ├── style.css
    └── app.js
```

`main.go`:

```go
package main

import (
	"embed"
	"log"
	"net/http"
)

//go:embed static
var staticFiles embed.FS

func main() {
	fs := http.FS(staticFiles)

	handler := http.FileServer(fs)

	http.Handle("/", handler)

	log.Println("Server running on http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
```

Here:

```go
embed.FS
```

implements the standard filesystem abstraction, and:

```go
http.FS(staticFiles)
```

adapts it for `net/http`.

This means your executable can contain the HTML, CSS, JavaScript, images, etc.

---

# 11. Another powerful use: templates

Because `embed.FS` implements `io/fs.FS`, it can work with `html/template` and `text/template`.

For example:

```go
package main

import (
	"embed"
	"html/template"
	"os"
)

//go:embed templates/*
var templates embed.FS

func main() {
	tmpl, err := template.ParseFS(
		templates,
		"templates/*.html",
	)
	if err != nil {
		panic(err)
	}

	data := struct {
		Name string
	}{
		Name: "Chandu",
	}

	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
```

This is a very common combination:

```text
embed
  ↓
embed.FS
  ↓
template.ParseFS
  ↓
HTML output
```

---

# 12. Three common beginner mistakes

## Mistake 1: Forgetting to import `embed`

You might write:

```go
//go:embed hello.txt
var message string
```

and get an error because the package isn't imported.

When using a `string` or `[]byte`:

```go
import _ "embed"
```

is enough.

When using `embed.FS`:

```go
import "embed"
```

is appropriate.

The blank import is necessary because the Go toolchain requires the `embed` package to be imported for `//go:embed`, even when the package name isn't directly referenced.

---

## Mistake 2: Thinking `embed` reads files at runtime

This is a major conceptual mistake.

Consider:

```go
//go:embed config.json
var config []byte
```

The file is incorporated during the **build**.

The application isn't dependent on `config.json` being present beside the executable afterward.

This is one of the main reasons `embed` is useful for distributing self-contained applications.

However, it also means changing `config.json` after compilation won't change the embedded data. You must rebuild the application.

---

## Mistake 3: Using filesystem paths incorrectly

Suppose:

```text
templates/
└── index.html
```

and:

```go
//go:embed templates
var files embed.FS
```

You read the file using:

```go
files.ReadFile("templates/index.html")
```

not:

```go
files.ReadFile("./templates/index.html")
```

and not:

```go
files.ReadFile("/templates/index.html")
```

Embedded filesystem paths follow the `io/fs` path conventions.

Also, embedding patterns are interpreted relative to the package containing the source file, and they use `/` as the path separator even on Windows.

---

# 13. Two real-world applications

## Application 1: Self-contained web servers

Imagine distributing:

```text
my-server
```

as one executable.

Instead of asking users to install:

```text
index.html
style.css
app.js
images/
```

you can embed all of them.

```text
                Go executable
              ┌───────────────┐
              │ Go application│
              │               │
              │ HTML          │
              │ CSS           │
              │ JavaScript    │
              │ Images        │
              └────────┬──────┘
                       │
                       ▼
                  HTTP server
```

This is excellent for internal tools, dashboards, admin applications, APIs with documentation UIs, and small production services.

---

## Application 2: CLI tools with bundled resources

A command-line program may need:

- SQL schemas
- default configuration
- templates
- migration files
- help documents
- certificates
- static assets

Instead of requiring:

```text
mytool
schema.sql
templates/
config/
```

you can package the resources into the executable.

That simplifies distribution and reduces the number of external files users need to manage.

---

# 14. `embed` vs normal filesystem access

| Feature | `os` filesystem | `embed` |
|---|---|---|
| Data location | External files | Inside executable |
| Read at build time | No | Yes |
| Read at runtime | Yes | Access embedded copy |
| Can modify embedded data | N/A | No |
| Survives moving executable | Depends on external files | Yes |
| Good for static assets | Yes | Excellent |
| Good for user-generated files | Yes | No |
| Good for mutable configuration | Yes | Usually no |

A useful rule is:

> **Embed resources that are part of your application; use the normal filesystem for data that users or the application need to modify.**

---

# 15. Important limitations

`embed` is not a replacement for a normal filesystem.

Embedded files are **read-only**. `embed.FS` is designed as a read-only filesystem and is safe to use concurrently.

For example, you cannot do:

```go
files.WriteFile(...)
```

because `embed.FS` doesn't provide writing functionality.

If your program needs to modify a file:

```text
database
user uploads
logs
cache
generated reports
runtime configuration
```

you generally need an ordinary filesystem or another storage mechanism.

---

# 16. Exercises

Here are three progressively challenging exercises. **No solutions are provided.**

## Exercise 1 — Embed a text file

Create a Go program with a file named:

```text
message.txt
```

containing a short message.

Use `//go:embed` to embed the file into a `string` variable.

Your program should:

1. Embed the file.
2. Print its contents.
3. Print the number of characters in the embedded string.
4. Verify that the program still works after the original `message.txt` is no longer present next to the executable.

---

## Exercise 2 — Build an embedded template system

Create this project:

```text
myapp/
├── main.go
└── templates/
    ├── home.html
    ├── about.html
    └── contact.html
```

Use `embed.FS` to embed the entire `templates` directory.

Your program should:

1. Read the templates using `embed.FS`.
2. Use `html/template` to parse them.
3. Create a struct containing data such as a user's name.
4. Execute one of the templates using that data.
5. Use `ReadDir` to list all embedded templates.

Try to make your design work without accessing the original template files at runtime.

---

## Exercise 3 — Build a self-contained web application

Build a small web server with:

```text
webapp/
├── main.go
└── static/
    ├── index.html
    ├── css/
    │   └── style.css
    └── js/
        └── app.js
```

Use `embed.FS` to embed the entire `static` directory.

Your application should:

1. Serve the embedded HTML through `net/http`.
2. Correctly serve CSS and JavaScript files.
3. Add at least two HTTP routes.
4. Use an embedded HTML template for one route.
5. Enumerate the embedded files using `ReadDir`.
6. Walk through the embedded filesystem using functionality from `io/fs`.
7. Make the final application distributable as a single executable without requiring the `static` directory at runtime.

As an additional challenge, experiment with how directory prefixes affect the paths used by your HTTP server.

---

# 17. Mental model to remember

The easiest way to remember the package is:

```text
             BUILD TIME
                 │
                 ▼
        //go:embed files
                 │
                 ▼
       ┌─────────────────┐
       │ Compiled binary │
       │                 │
       │ embedded data   │
       └────────┬────────┘
                │
                ▼
             RUNTIME
                │
        ┌───────┴────────┐
        ▼                ▼
     string           embed.FS
     []byte               │
                          ├── Open()
                          ├── ReadFile()
                          └── ReadDir()
```

And the most important distinction is:

```text
string       → one text file
[]byte       → one binary/file resource
embed.FS     → collection/tree of files
```

The `embed` package is deliberately small: its primary API is `embed.FS` with `Open`, `ReadFile`, and `ReadDir`; the `//go:embed` directive is the mechanism that causes resources to be included at build time.

---

# 18. Thought-provoking question

If you were designing a production Go web application, **which resources would you choose to embed and which would you deliberately keep outside the binary—and how would those choices affect configuration management, security, binary size, and the ability to update resources without recompiling?**
