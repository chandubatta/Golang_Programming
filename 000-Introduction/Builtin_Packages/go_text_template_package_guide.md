# Go `text/template` Package

The `text/template` package is part of Go's standard library and provides a powerful way to generate **textual output from templates and data**. It is commonly used for configuration files, emails, reports, source-code generation, CLI output, and other situations where the structure of the output is mostly fixed but the actual values change.

> **Important:** `text/template` does **not** automatically escape output. If you are generating HTML, prefer `html/template`, which provides contextual HTML escaping.

---

## 1. What is `text/template`?

A template is essentially a piece of text containing **placeholders and instructions**.

For example:

```text
Hello, {{.Name}}!
You are {{.Age}} years old.
```

Given:

```go
User{
    Name: "Chandu",
    Age: 25,
}
```

the template produces:

```text
Hello, Chandu!
You are 25 years old.
```

The `.` is called **dot**. It represents the current data being processed by the template.

### Basic workflow

Most template programs follow this pattern:

```text
Create template
      ↓
Parse template
      ↓
Provide data
      ↓
Execute template
      ↓
Generate output
```

---

# 2. Simple Example

```go
package main

import (
	"os"
	"text/template"
)

type User struct {
	Name  string
	Age   int
	Email string
}

func main() {
	const templateText = `
User Information
----------------
Name:  {{.Name}}
Age:   {{.Age}}
Email: {{.Email}}
`

	tmpl, err := template.New("user").Parse(templateText)
	if err != nil {
		panic(err)
	}

	user := User{
		Name:  "Chandu",
		Age:   25,
		Email: "chandu@example.com",
	}

	err = tmpl.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
```

### Output

```text
User Information
----------------
Name:  Chandu
Age:   25
Email: chandu@example.com
```

Here:

- `template.New("user")` creates a template.
- `Parse()` converts the template text into an executable template.
- `{{.Name}}` accesses the `Name` field.
- `Execute()` supplies the `User` value and writes the generated output to `os.Stdout`.

---

# 3. Important Template Concepts

Before looking at every API function, you should understand these concepts.

## Actions

Anything inside:

```text
{{ ... }}
```

is a template **action**.

Example:

```text
Hello {{.Name}}
```

The text outside the action is copied directly to the output.

---

## Fields

Given:

```go
type User struct {
	Name string
}
```

you can access the field with:

```text
{{.Name}}
```

Nested fields can be chained:

```text
{{.Address.City}}
```

---

## `if`

```text
{{if .IsAdmin}}
    Administrator
{{else}}
    Regular user
{{end}}
```

---

## `range`

Used for iterating over collections:

```text
{{range .Users}}
    {{.Name}}
{{end}}
```

---

## `with`

Changes the current `.` value:

```text
{{with .Address}}
City: {{.City}}
Country: {{.Country}}
{{end}}
```

Inside the `with`, `.` represents `.Address`.

---

## Pipelines

You can pass the result of one operation into another:

```text
{{.Name | printf "%q"}}
```

Conceptually:

```text
printf("%q", .Name)
```

Pipelines become particularly useful with custom template functions.

---

# 4. Every Public Function in `text/template`

The package exposes several **package-level functions**, plus methods on `Template`.

The main package-level functions are:

- `HTMLEscape`
- `HTMLEscapeString`
- `HTMLEscaper`
- `IsTrue`
- `JSEscape`
- `JSEscapeString`
- `JSEscaper`
- `Must`
- `New`
- `ParseFS`
- `ParseFiles`
- `ParseGlob`
- `URLQueryEscaper`

It also provides the `Template` type and its methods.

---

# 5. `template.New`

### Signature

```go
func New(name string) *Template
```

Creates a new, undefined template with the specified name.

Example:

```go
tmpl := template.New("welcome")
```

You normally follow it with `Parse()`:

```go
tmpl, err := template.New("welcome").Parse("Hello {{.Name}}")
```

### When to use it

Use `New()` when you want to explicitly create and name a template.

---

# 6. `template.Must`

### Signature

```go
func Must(t *Template, err error) *Template
```

`Must()` is a convenience function that returns the template if `err == nil`, but **panics if the error is non-nil**.

Instead of:

```go
tmpl, err := template.New("test").Parse("Hello {{.Name}}")
if err != nil {
	panic(err)
}
```

you can write:

```go
tmpl := template.Must(
	template.New("test").Parse("Hello {{.Name}}"),
)
```

### When should you use it?

It is particularly useful when a template is defined by the programmer and a parsing error represents a programming/configuration error.

Avoid using `Must()` for templates containing user-controlled input.

---

# 7. `template.ParseFiles`

### Signature

```go
func ParseFiles(filenames ...string) (*Template, error)
```

Reads one or more template files from the operating system and parses them.

Example:

```go
tmpl, err := template.ParseFiles(
	"header.txt",
	"body.txt",
	"footer.txt",
)
if err != nil {
	panic(err)
}
```

This is useful when your templates are stored as separate files rather than embedded in Go source code.

For example:

```text
templates/
    header.txt
    body.txt
    footer.txt
```

You can keep your presentation/output structure separate from your Go logic.

---

# 8. `template.ParseGlob`

### Signature

```go
func ParseGlob(pattern string) (*Template, error)
```

Parses template files matching a filesystem glob pattern.

Example:

```go
tmpl, err := template.ParseGlob("templates/*.txt")
if err != nil {
	panic(err)
}
```

If you have:

```text
templates/
    home.txt
    about.txt
    contact.txt
```

the pattern:

```text
templates/*.txt
```

can load all of them.

### When useful

This is especially convenient for applications with many template files.

---

# 9. `template.ParseFS`

### Signature

```go
func ParseFS(
	fsys fs.FS,
	patterns ...string,
) (*Template, error)
```

`ParseFS()` parses templates from an `fs.FS` filesystem abstraction. It was added in Go 1.16.

This becomes particularly interesting with Go's `embed` package.

Example:

```go
package main

import (
	"embed"
	"os"
	"text/template"
)

//go:embed templates/*
var templateFS embed.FS

func main() {
	tmpl, err := template.ParseFS(
		templateFS,
		"templates/*.txt",
	)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, "Hello")
	if err != nil {
		panic(err)
	}
}
```

### Why is this useful?

You can package your templates **inside your compiled Go application** rather than requiring separate template files at runtime.

---

# 10. `template.HTMLEscape`

### Signature

```go
func HTMLEscape(w io.Writer, b []byte)
```

Writes an HTML-escaped representation of a byte slice to an `io.Writer`.

For example, text containing:

```text
<script>alert("hello")</script>
```

can be escaped so that it is represented as HTML text rather than interpreted as HTML.

### Important distinction

Although these HTML escaping helpers exist in `text/template`, **you generally should not use `text/template` to build HTML**.

Use:

```go
html/template
```

instead.

`html/template` is specifically designed to provide safer contextual HTML escaping.

---

# 11. `template.HTMLEscapeString`

### Signature

```go
func HTMLEscapeString(s string) string
```

Returns the HTML-escaped version of a string.

Example:

```go
input := `<b>Hello</b>`

escaped := template.HTMLEscapeString(input)

fmt.Println(escaped)
```

The result represents the HTML characters safely rather than treating the input as HTML markup.

---

# 12. `template.HTMLEscaper`

### Signature

```go
func HTMLEscaper(args ...any) string
```

Converts its arguments to textual form and HTML-escapes the resulting representation.

It is also available as a predefined template function named `html`.

For example:

```text
{{html .Content}}
```

However, for HTML generation you should normally use `html/template`.

---

# 13. `template.JSEscape`

### Signature

```go
func JSEscape(w io.Writer, b []byte)
```

Writes a JavaScript-escaped representation of the byte slice to a writer.

This can be useful when you specifically need to escape textual data for JavaScript contexts.

---

# 14. `template.JSEscapeString`

### Signature

```go
func JSEscapeString(s string) string
```

Returns a JavaScript-escaped representation of a string.

Example:

```go
safe := template.JSEscapeString(`Hello "World"`)
```

---

# 15. `template.JSEscaper`

### Signature

```go
func JSEscaper(args ...any) string
```

Returns a JavaScript-escaped representation of the textual representation of its arguments.

It corresponds to the predefined template function:

```text
js
```

For example:

```text
{{js .Message}}
```

---

# 16. `template.URLQueryEscaper`

### Signature

```go
func URLQueryEscaper(args ...any) string
```

Escapes the textual representation of its arguments into a form suitable for use in a URL query.

Example:

```go
value := template.URLQueryEscaper("hello world")
```

The result is conceptually related to URL query escaping such as:

```text
hello%20world
```

The template package also exposes this functionality through the predefined template function:

```text
urlquery
```

---

# 17. `template.IsTrue`

### Signature

```go
func IsTrue(val any) (truth, ok bool)
```

Reports whether a value is considered **true/non-empty** according to the template package's notion of truth.

This is particularly useful when implementing functionality around template evaluation.

For example, template conditionals use this concept:

```text
{{if .Value}}
    Value exists
{{end}}
```

The boolean behavior is based on zero/empty values and non-zero/non-empty values.

---

# 18. `Template.Parse`

### Signature

```go
func (t *Template) Parse(text string) (*Template, error)
```

This is one of the most important methods.

It parses template text and associates the resulting template with `t`.

Example:

```go
tmpl, err := template.New("greeting").Parse(
	"Hello {{.Name}}!",
)
```

Then:

```go
tmpl.Execute(os.Stdout, user)
```

### Important

Parsing and executing are separate operations:

```text
Parse → understand/compile template
Execute → apply data and produce output
```

This separation is useful because you can parse once and execute many times.

---

# 19. `Template.Execute`

### Signature

```go
func (t *Template) Execute(wr io.Writer, data any) error
```

Executes the template using `data` and writes the resulting output to `wr`.

Example:

```go
var output bytes.Buffer

err := tmpl.Execute(&output, user)
if err != nil {
	panic(err)
}

fmt.Println(output.String())
```

The writer can be:

- `os.Stdout`
- a file
- `bytes.Buffer`
- an HTTP response writer
- another `io.Writer`

### Important

Execution can produce **partial output before an error occurs**, so you should not automatically assume that an error means nothing was written.

---

# 20. `Template.ExecuteTemplate`

### Signature

```go
func (t *Template) ExecuteTemplate(
	wr io.Writer,
	name string,
	data any,
) error
```

Executes a specific named template associated with `t`.

This becomes important when you have multiple templates.

Example:

```go
const text = `
{{define "welcome"}}
Welcome, {{.Name}}!
{{end}}

{{define "goodbye"}}
Goodbye, {{.Name}}!
{{end}}
`

tmpl := template.Must(
	template.New("messages").Parse(text),
)

tmpl.ExecuteTemplate(os.Stdout, "welcome", user)
```

---

# 21. `Template.New`

### Signature

```go
func (t *Template) New(name string) *Template
```

Creates a new template associated with an existing template.

Example:

```go
root := template.New("root")

header := root.New("header")
```

Associated templates can invoke each other using:

```text
{{template "header" .}}
```

This is useful for building reusable template structures.

---

# 22. `Template.Lookup`

### Signature

```go
func (t *Template) Lookup(name string) *Template
```

Finds an associated template by name.

Example:

```go
header := tmpl.Lookup("header")

if header == nil {
	fmt.Println("header template not found")
}
```

If the template doesn't exist, it returns `nil`.

---

# 23. `Template.Name`

### Signature

```go
func (t *Template) Name() string
```

Returns the name of the template.

Example:

```go
tmpl := template.New("invoice")

fmt.Println(tmpl.Name())
```

Output:

```text
invoice
```

---

# 24. `Template.Funcs`

### Signature

```go
func (t *Template) Funcs(funcMap FuncMap) *Template
```

Adds custom functions that can be called from templates.

This is one of the most powerful features of `text/template`.

Example:

```go
func upper(s string) string {
	return strings.ToUpper(s)
}

tmpl := template.New("test").Funcs(
	template.FuncMap{
		"upper": upper,
	},
)

tmpl = template.Must(
	tmpl.Parse("Name: {{upper .Name}}"),
)
```

Then:

```text
Name: CHANDU
```

### Critical rule

Functions needed while parsing must be registered **before parsing**:

```go
template.New("test").
	Funcs(funcMap).
	Parse(text)
```

The function map values must be functions with supported return signatures. Invalid function registrations can cause a panic.

---

# 25. `template.FuncMap`

`FuncMap` is technically a **type**, not a function:

```go
type FuncMap map[string]any
```

It maps names used in templates to Go functions.

Example:

```go
funcMap := template.FuncMap{
	"upper": strings.ToUpper,
	"add": func(a, b int) int {
		return a + b
	},
}
```

Template:

```text
{{upper .Name}}
{{add .X .Y}}
```

A custom function can return either:

```go
func(...) T
```

or:

```go
func(...) (T, error)
```

If the second return value is a non-nil error, template execution stops.

---

# 26. `Template.Option`

### Signature

```go
func (t *Template) Option(opt ...string) *Template
```

Configures template execution behavior.

One particularly important option is:

```text
missingkey
```

For example:

```go
tmpl.Option("missingkey=error")
```

### Available behaviors

```text
missingkey=default
missingkey=invalid
missingkey=zero
missingkey=error
```

With:

```go
missingkey=error
```

an attempt to access a missing map key causes execution to stop with an error.

This is often useful in production applications because silently missing values can otherwise be difficult to detect.

---

# 27. `Template.Delims`

### Signature

```go
func (t *Template) Delims(left, right string) *Template
```

Changes the delimiters used for template actions.

Normally:

```text
{{.Name}}
```

But you could change them:

```go
tmpl.Delims("[[", "]]")
```

Then:

```text
Hello [[.Name]]
```

The new delimiters apply to subsequent parsing operations.

This can be useful if `{{` and `}}` already have meaning in the content being generated.

---

# 28. `Template.Clone`

### Signature

```go
func (t *Template) Clone() (*Template, error)
```

Creates a copy of a template and its associated templates.

This is useful when you want to create variants of a common template without modifying the original.

Conceptually:

```text
Base template
      │
      ├── Clone A → configuration A
      │
      └── Clone B → configuration B
```

Templates can then be customized independently.

---

# 29. `Template.Templates`

### Signature

```go
func (t *Template) Templates() []*Template
```

Returns the templates associated with `t`.

Example:

```go
for _, tmpl := range root.Templates() {
	fmt.Println(tmpl.Name())
}
```

This is useful when working with collections of named templates.

---

# 30. `Template.DefinedTemplates`

### Signature

```go
func (t *Template) DefinedTemplates() string
```

Returns a string listing the templates that have definitions.

It is mainly useful for diagnostics and error messages.

---

# 31. `Template.ParseFiles`

### Signature

```go
func (t *Template) ParseFiles(
	filenames ...string,
) (*Template, error)
```

This is the method version of the package-level `ParseFiles`.

Example:

```go
tmpl := template.New("report")

tmpl, err := tmpl.ParseFiles(
	"header.txt",
	"report.txt",
	"footer.txt",
)
```

It allows additional files to be parsed into an existing template association.

---

# 32. `Template.ParseGlob`

### Signature

```go
func (t *Template) ParseGlob(
	pattern string,
) (*Template, error)
```

Parses all files matching a glob pattern into an existing template association.

Example:

```go
tmpl := template.New("pages")

tmpl, err := tmpl.ParseGlob("templates/*.txt")
```

Useful when your application has a directory containing many related templates.

---

# 33. `Template.ParseFS`

### Signature

```go
func (t *Template) ParseFS(
	fsys fs.FS,
	patterns ...string,
) (*Template, error)
```

Parses templates from an `fs.FS`.

For example, this works nicely with `embed.FS`:

```go
//go:embed templates/*
var templateFS embed.FS

tmpl := template.Must(
	template.New("pages").ParseFS(
		templateFS,
		"templates/*.txt",
	),
)
```

This is a modern and convenient way to distribute templates with a Go application.

---

# 34. Built-in Template Functions

`text/template` also provides predefined functions available directly inside templates.

Some of the most important ones are:

| Function | Purpose |
|---|---|
| `and` | Logical AND / returns first empty argument or last argument |
| `or` | Logical OR / returns first non-empty argument or last argument |
| `not` | Boolean negation |
| `eq` | Equality comparison |
| `ne` | Not equal |
| `lt` | Less than |
| `le` | Less than or equal |
| `gt` | Greater than |
| `ge` | Greater than or equal |
| `index` | Index a map, slice, array, etc. |
| `slice` | Slice a string, slice, or array |
| `len` | Get length |
| `call` | Call a function-valued value |
| `print` | Similar to `fmt.Sprint` |
| `printf` | Similar to `fmt.Sprintf` |
| `println` | Similar to `fmt.Sprintln` |
| `html` | HTML escaping |
| `js` | JavaScript escaping |
| `urlquery` | URL query escaping |

For example:

```text
{{if eq .Role "admin"}}
    Administrator
{{else}}
    Regular User
{{end}}
```

And:

```text
{{range .Users}}
    {{.Name}}
{{end}}
```

---

# 35. Three Common Beginner Mistakes

## Mistake 1: Using `text/template` for untrusted HTML

A beginner might write:

```go
template.New("page").Parse(`
<h1>{{.Title}}</h1>
`)
```

and assume the package automatically makes the HTML safe.

It doesn't.

`text/template` is designed for general text and assumes template authors are trusted. It does **not** automatically escape output.

### Avoid it

For HTML:

```go
import "html/template"
```

instead of:

```go
import "text/template"
```

---

## Mistake 2: Registering custom functions too late

Incorrect:

```go
tmpl, _ := template.New("test").Parse(
	"{{upper .Name}}",
)

tmpl.Funcs(template.FuncMap{
	"upper": strings.ToUpper,
})
```

The function should generally be registered before parsing:

```go
tmpl := template.Must(
	template.New("test").
		Funcs(template.FuncMap{
			"upper": strings.ToUpper,
		}).
		Parse("{{upper .Name}}"),
)
```

`Funcs()` is specifically intended to establish the functions available during template parsing.

---

## Mistake 3: Confusing `.` with the original data

Consider:

```text
{{with .Address}}
City: {{.City}}
{{end}}
```

Inside the `with`, `.` is no longer the entire original object.

It now represents:

```text
.Address
```

Similarly, inside:

```text
{{range .Users}}
```

`.` represents each individual user.

### Avoid it

Think of `.` as:

> **"What data am I currently looking at?"**

This mental model makes `with`, `range`, nested templates, and pipelines much easier to understand.

---

# 36. Two Real-World Applications

## Application 1: Configuration generation

Suppose you have an application that needs to generate configuration files:

```text
server:
  host: {{.Host}}
  port: {{.Port}}
  workers: {{.Workers}}
```

You could provide different data for development, staging, and production.

For example:

```text
development → development configuration
staging     → staging configuration
production  → production configuration
```

The template stays mostly unchanged while the data changes.

This is useful for deployment automation and infrastructure tooling.

---

## Application 2: Automated reports and emails

Imagine a system that generates a daily report:

```text
Daily Sales Report
==================

Date: {{.Date}}

Total Orders: {{.Orders}}
Revenue:      {{.Revenue}}

Top Products:
{{range .Products}}
- {{.Name}}: {{.Sales}}
{{end}}
```

The Go program supplies the data, while the template controls the presentation.

This separation makes it easier to modify the report format without changing the business logic.

---

# 37. Three Progressively Challenging Exercises

## Exercise 1 — Basic User Profile

Create a `User` struct containing:

- Name
- Age
- Email
- City

Create a `text/template` template that produces a formatted user profile.

Requirements:

- Use `template.New()`.
- Parse the template with `Parse()`.
- Use field expressions such as `{{.Name}}`.
- Execute the template using `Execute()`.
- Write the result to standard output.

**Do not hard-code the user's values into the template.**

---

## Exercise 2 — Employee Report with Conditions and Loops

Create an employee-report generator.

Each employee should contain:

- Name
- Department
- Salary
- IsManager

Your program should generate a report containing all employees.

Requirements:

- Use `range` to iterate over employees.
- Use `if` to identify managers.
- Display a different message for managers and non-managers.
- Add a custom template function that formats salary.
- Register the function using `FuncMap`.
- Store the template in a separate `.txt` file.
- Load it using `ParseFiles()`.

---

## Exercise 3 — Reusable Multi-Template Reporting System

Build a reusable report-generation system containing several named templates:

```text
header
employee
footer
report
```

Requirements:

- Use `{{define}}` to create named templates.
- Use `{{template}}` to invoke other templates.
- Load multiple templates from a directory.
- Use `ParseGlob()` or `ParseFS()`.
- Add at least three custom template functions.
- Use `ExecuteTemplate()` to execute a specific named template.
- Add `missingkey=error`.
- Handle template parsing and execution errors properly.
- Use `embed.FS` so the templates can be packaged into the compiled Go application.
- Generate reports for multiple datasets without reparsing the templates for every report.

**Do not provide a solution until you have attempted the exercise yourself.**

---

# 38. A Useful Mental Model

When learning `text/template`, remember this pipeline:

```text
                 TEMPLATE
                    │
                    ▼
              ┌───────────┐
              │   Parse   │
              └─────┬─────┘
                    │
                    ▼
             Parsed Template
                    │
                    │ + Data
                    ▼
              ┌───────────┐
              │  Execute  │
              └─────┬─────┘
                    │
                    ▼
              Generated Text
```

And for more complex applications:

```text
                    Template
                       │
        ┌──────────────┼──────────────┐
        ▼              ▼              ▼
      Fields          if/range       Functions
        │              │              │
        └──────────────┼──────────────┘
                       ▼
                    Execute
                       │
                       ▼
                    Output
```

The most important distinction to internalize is:

**Parsing determines what the template means; execution applies actual data to that template.**

Once you understand that distinction, `Parse`, `Execute`, `ExecuteTemplate`, `Funcs`, `FuncMap`, `range`, `if`, `with`, pipelines, and named templates become much easier to reason about.

---

# Thought-Provoking Question

**Suppose you were designing a large Go application where non-programmers are allowed to edit templates. What capabilities would you allow inside the templates, and what capabilities would you deliberately keep in Go code—and how would your decision affect security, maintainability, and the boundary between "presentation" and "business logic"?**
