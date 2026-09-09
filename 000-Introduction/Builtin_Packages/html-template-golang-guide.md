# `html/template` Package in Go

The Go `html/template` package is the standard-library package for generating **HTML dynamically from data** while providing **context-aware automatic escaping** to protect against common code-injection and XSS problems. It has essentially the same template language/API as `text/template`, but is specifically designed for HTML and understands HTML, CSS, JavaScript, and URL contexts.

## 1. What is `html/template`?

Normally, an HTML page contains static content:

```html
<h1>Hello, Chandu</h1>
```

But web applications often need to generate HTML from dynamic Go data:

```go
user := "Chandu"
```

Instead of manually concatenating strings:

```go
fmt.Sprintf("<h1>Hello, %s</h1>", user)
```

you can create a template:

```html
<h1>Hello, {{.}}</h1>
```

and execute it with Go data.

The `html/template` package handles the connection between **HTML templates** and **Go data**.

### Why use `html/template`?

Its major purposes are:

- Generate dynamic HTML.
- Separate presentation from application logic.
- Insert Go data into HTML.
- Reuse HTML layouts and components.
- Iterate over slices and maps.
- Conditionally display content.
- Define reusable template blocks.
- Automatically escape untrusted data according to its HTML/JS/CSS/URL context.

That last point is particularly important.

For example, if a user submits:

```text
<script>alert("Hacked!")</script>
```

and you render it using:

```html
<p>{{.}}</p>
```

`html/template` escapes the value rather than allowing the browser to interpret it as executable HTML.

### When is it commonly used?

It is especially useful when building:

- Server-rendered web applications
- HTML pages served by `net/http`
- Admin dashboards
- Blogs
- E-commerce pages
- Documentation websites
- Server-side forms
- HTML email generation
- Go web applications that don't need a separate frontend framework

---

# 2. Basic Mental Model

A useful way to understand `html/template` is:

```text
Go Data
   ↓
Template
   ↓
html/template processing
   ↓
Context-aware escaping
   ↓
Generated HTML
   ↓
Browser
```

For example:

```text
Go struct
   ↓
{{.Name}}
   ↓
HTML template
   ↓
<h1>Chandu</h1>
```

The `.` is called the **dot** or **data context**.

If you execute:

```go
tmpl.Execute(w, user)
```

then:

```text
.
```

represents `user`.

If `user` has:

```go
type User struct {
    Name string
}
```

you can access:

```html
{{.Name}}
```

---

# 3. Simple Code Example

Here is a complete example:

```go
package main

import (
	"html/template"
	"log"
	"net/http"
)

type User struct {
	Name  string
	Email string
}

var tmpl = template.Must(template.New("user").Parse(`
<!DOCTYPE html>
<html>
<head>
	<title>User Profile</title>
</head>
<body>
	<h1>Hello, {{.Name}}!</h1>
	<p>Email: {{.Email}}</p>
</body>
</html>
`))

func userHandler(w http.ResponseWriter, r *http.Request) {
	user := User{
		Name:  "Chandu",
		Email: "chandu@example.com",
	}

	err := tmpl.Execute(w, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/", userHandler)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

Run:

```bash
go run main.go
```

Then open:

```text
http://localhost:8080
```

You would see:

```text
Hello, Chandu!

Email: chandu@example.com
```

### What happens?

The important line is:

```go
tmpl.Execute(w, user)
```

`user` becomes the template's `.` value.

Therefore:

```html
{{.Name}}
```

becomes:

```html
Chandu
```

and:

```html
{{.Email}}
```

becomes:

```html
chandu@example.com
```

---

# 4. Important Template Concepts

## `{{.}}`

Represents the current data.

```html
<p>{{.}}</p>
```

If the data is:

```go
"Hello"
```

the result is:

```html
<p>Hello</p>
```

---

## `{{.Name}}`

Accesses a field.

```html
<h1>{{.Name}}</h1>
```

---

## `{{if}}`

Conditional rendering:

```html
{{if .IsAdmin}}
	<p>Administrator</p>
{{else}}
	<p>Regular user</p>
{{end}}
```

---

## `{{range}}`

Iteration:

```html
<ul>
{{range .Users}}
	<li>{{.Name}}</li>
{{end}}
</ul>
```

---

## `{{with}}`

Changes the current dot:

```html
{{with .Address}}
	<p>{{.City}}</p>
{{end}}
```

---

## `{{define}}`

Defines a reusable template:

```html
{{define "header"}}
<header>
	<h1>My Website</h1>
</header>
{{end}}
```

Then:

```html
{{template "header" .}}
```

---

# 5. Every Function in `html/template`

The package-level functions include:

1. `HTMLEscape`
2. `HTMLEscapeString`
3. `HTMLEscaper`
4. `IsTrue`
5. `JSEscape`
6. `JSEscapeString`
7. `JSEscaper`
8. `URLQueryEscaper`
9. `Must`
10. `New`
11. `ParseFS`
12. `ParseFiles`
13. `ParseGlob`

The package also provides numerous methods on `Template`.

---

# 6. `HTMLEscape`

### Signature

```go
func HTMLEscape(w io.Writer, b []byte)
```

`HTMLEscape` writes an escaped version of HTML text to an `io.Writer`.

Example:

```go
package main

import (
	"html/template"
	"os"
)

func main() {
	data := []byte(`<script>alert("hello")</script>`)

	template.HTMLEscape(os.Stdout, data)
}
```

The dangerous HTML characters are escaped.

### When useful?

When you already have data as `[]byte` and want to safely write its HTML-escaped representation somewhere.

---

# 7. `HTMLEscapeString`

### Signature

```go
func HTMLEscapeString(s string) string
```

It returns the HTML-escaped representation of a string.

Example:

```go
package main

import (
	"fmt"
	"html/template"
)

func main() {
	input := `<script>alert("XSS")</script>`

	result := template.HTMLEscapeString(input)

	fmt.Println(result)
}
```

This is useful when you need the escaped result as a Go `string`.

### Difference

`HTMLEscape`:

```text
string/bytes → Writer
```

`HTMLEscapeString`:

```text
string → string
```

---

# 8. `HTMLEscaper`

### Signature

```go
func HTMLEscaper(args ...any) string
```

It HTML-escapes the textual representation of its arguments and returns the result.

Example:

```go
result := template.HTMLEscaper(
	"<h1>",
	"Hello",
	"</h1>",
)
```

It is useful when multiple values need to be converted into an escaped HTML string.

---

# 9. `IsTrue`

### Signature

```go
func IsTrue(val any) (truth, ok bool)
```

`IsTrue` determines whether a value has the "truth" value used by template `if` and similar actions.

For example, values such as:

```go
true
```

are considered true, while empty/nil/zero-like values can be false according to template truth rules.

This is primarily useful when implementing functionality around template execution rather than ordinary application code.

---

# 10. `JSEscape`

### Signature

```go
func JSEscape(w io.Writer, b []byte)
```

This writes a JavaScript-escaped representation of byte data to an `io.Writer`.

For example, when data needs to safely cross into a JavaScript context, escaping requirements differ from ordinary HTML text.

This distinction is one of the major reasons to use `html/template` rather than manually concatenating strings.

---

# 11. `JSEscapeString`

### Signature

```go
func JSEscapeString(s string) string
```

Returns the JavaScript-escaped representation of a string.

Conceptually:

```text
Go string
   ↓
JSEscapeString
   ↓
safe representation for JavaScript context
```

---

# 12. `JSEscaper`

### Signature

```go
func JSEscaper(args ...any) string
```

Escapes the textual representations of its arguments for JavaScript and returns the result.

This is lower-level functionality; most normal applications should allow `html/template`'s contextual escaping to handle template values automatically.

---

# 13. `URLQueryEscaper`

### Signature

```go
func URLQueryEscaper(args ...any) string
```

It returns the escaped representation of its arguments suitable for use in URL query components.

For example, a value containing spaces or special characters cannot simply be inserted into a URL without appropriate encoding.

`html/template` itself also performs contextual escaping when values are placed into URL contexts.

---

# 14. `Must`

### Signature

```go
func Must(t *Template, err error) *Template
```

`Must` is one of the most commonly used helpers.

It takes a template and an error:

```go
template.Must(template.New("page").Parse(...))
```

If the error is non-nil, `Must` panics.

If there is no error, it returns the template.

It is specifically intended for convenient initialization of templates.

### Without `Must`

```go
tmpl, err := template.New("page").Parse(html)

if err != nil {
	log.Fatal(err)
}
```

### With `Must`

```go
tmpl := template.Must(
	template.New("page").Parse(html),
)
```

### Important

Don't blindly use `Must` everywhere.

It is particularly appropriate when parsing a template during application startup and a malformed template represents a programmer/configuration error.

---

# 15. `New`

### Signature

```go
func New(name string) *Template
```

Creates a new HTML template with the specified name.

Example:

```go
tmpl := template.New("home")
```

Then:

```go
tmpl.Parse(...)
```

can be used to parse template contents.

Typical pattern:

```go
tmpl := template.Must(
	template.New("home").Parse(`
		<h1>{{.Title}}</h1>
	`),
)
```

---

# 16. `ParseFS`

### Signature

```go
func ParseFS(
	fs fs.FS,
	patterns ...string,
) (*Template, error)
```

`ParseFS` parses templates from an `fs.FS` filesystem. It is particularly useful with Go's `embed` package.

Example:

```go
//go:embed templates/*
var templateFS embed.FS

templates, err := template.ParseFS(
	templateFS,
	"templates/*.html",
)
```

This allows templates to be packaged into the compiled Go binary.

### Typical architecture

```text
templates/
    index.html
    about.html
    profile.html
        ↓
    embed.FS
        ↓
html/template.ParseFS
        ↓
Go application
```

This is extremely useful for production applications where you want templates bundled into the executable.

---

# 17. `ParseFiles`

### Signature

```go
func ParseFiles(filenames ...string) (*Template, error)
```

`ParseFiles` reads template files from the operating system's filesystem and parses them.

Suppose you have:

```text
templates/
    index.html
    profile.html
```

You can do:

```go
templates, err := template.ParseFiles(
	"templates/index.html",
	"templates/profile.html",
)
```

Then:

```go
templates.ExecuteTemplate(
	w,
	"profile.html",
	data,
)
```

### When useful?

During development, this is often the easiest approach.

---

# 18. `ParseGlob`

### Signature

```go
func ParseGlob(pattern string) (*Template, error)
```

`ParseGlob` parses all files matching a filesystem glob pattern.

For example:

```go
templates, err := template.ParseGlob(
	"templates/*.html",
)
```

Instead of manually specifying:

```text
templates/index.html
templates/about.html
templates/profile.html
```

the glob selects all matching files.

---

# 19. `Template.AddParseTree`

### Signature

```go
func (t *Template) AddParseTree(
	name string,
	tree *parse.Tree,
) (*Template, error)
```

This associates a parsed `text/template/parse` tree with a template.

This is an **advanced API** and is rarely necessary for beginners.

It is useful when you are programmatically manipulating parsed template structures or integrating template parsing infrastructure.

It returns an error if the template has already been executed.

---

# 20. `Template.Clone`

### Signature

```go
func (t *Template) Clone() (*Template, error)
```

`Clone` creates a copy of a template and its associated templates while providing an independent template namespace.

This is useful for creating template variants.

Conceptually:

```text
Base Template
      │
      ├── Clone A → customer layout
      │
      └── Clone B → admin layout
```

You can define common templates once and then customize the clones.

---

# 21. `Template.DefinedTemplates`

### Signature

```go
func (t *Template) DefinedTemplates() string
```

Returns a string describing the named templates defined within the template set.

For example, if you have:

```html
{{define "header"}}
...
{{end}}

{{define "footer"}}
...
{{end}}
```

you can inspect which named templates have been defined.

This is mainly useful for diagnostics and error reporting rather than ordinary application logic.

---

# 22. `Template.Delims`

### Signature

```go
func (t *Template) Delims(left, right string) *Template
```

Changes the delimiters used by templates.

Normally:

```text
{{ ... }}
```

You can change them:

```go
tmpl.Delims("<<", ">>")
```

Then:

```html
<h1><<.Name>></h1>
```

becomes valid template syntax.

This is useful when `{{ }}` conflicts with another templating system or language embedded in the page.

---

# 23. `Template.Execute`

### Signature

```go
func (t *Template) Execute(
	wr io.Writer,
	data any,
) error
```

This is one of the most important methods.

It executes the template and writes the generated output to an `io.Writer`.

In an HTTP server:

```go
err := tmpl.Execute(w, data)
```

Here:

```text
tmpl = template
w    = HTTP response writer
data = data supplied to template
```

This is the method you will use constantly when building Go web applications.

---

# 24. `Template.ExecuteTemplate`

### Signature

```go
func (t *Template) ExecuteTemplate(
	wr io.Writer,
	name string,
	data any,
) error
```

This executes a **specific named template** from a collection of templates.

For example:

```go
templates.ExecuteTemplate(
	w,
	"profile.html",
	user,
)
```

This becomes particularly useful when you have:

```text
templates/
    base.html
    home.html
    profile.html
    settings.html
```

and need to select which template to render.

---

# 25. `Template.Funcs`

### Signature

```go
func (t *Template) Funcs(
	funcMap FuncMap,
) *Template
```

`Funcs` adds custom Go functions that can be called from templates.

Example:

```go
func upper(s string) string {
	return strings.ToUpper(s)
}

tmpl := template.New("page").Funcs(
	template.FuncMap{
		"upper": upper,
	},
)
```

Then:

```html
<h1>{{upper .Name}}</h1>
```

If:

```text
.Name = "chandu"
```

the result is:

```text
CHANDU
```

### Important rule

Functions used in templates generally need to be registered **before parsing** the template.

---

# 26. `Template.Lookup`

### Signature

```go
func (t *Template) Lookup(name string) *Template
```

Finds a named template associated with another template.

Example:

```go
header := templates.Lookup("header")
```

If it doesn't exist:

```go
header == nil
```

This is useful when working with template sets and named templates.

---

# 27. `Template.Name`

### Signature

```go
func (t *Template) Name() string
```

Returns the template's name.

Example:

```go
tmpl := template.New("home")

fmt.Println(tmpl.Name())
```

Output:

```text
home
```

Simple, but useful when debugging template collections.

---

# 28. `Template.New`

### Signature

```go
func (t *Template) New(name string) *Template
```

Creates a new associated template.

For example:

```go
t := template.New("base")

t.New("header")
t.New("footer")
```

These templates become part of the same template set.

That allows named templates to call each other using:

```html
{{template "header" .}}
```

The association is transitive.

---

# 29. `Template.Option`

### Signature

```go
func (t *Template) Option(opt ...string) *Template
```

Sets template execution options.

One particularly useful option is:

```go
Option("missingkey=error")
```

Suppose you have:

```go
data := map[string]string{
	"name": "Chandu",
}
```

and your template tries:

```html
{{.email}}
```

With:

```go
tmpl.Option("missingkey=error")
```

the execution fails rather than silently continuing.

The supported `missingkey` modes include:

```text
missingkey=default
missingkey=invalid
missingkey=zero
missingkey=error
```

The default behavior is effectively to continue and produce `<no value>` when the missing map value is printed.

For development, a useful choice is:

```go
template.New("page").Option("missingkey=error")
```

because it helps reveal mistakes early.

---

# 30. `Template.Parse`

### Signature

```go
func (t *Template) Parse(
	text string,
) (*Template, error)
```

Parses a string as a template body.

Example:

```go
tmpl, err := template.New("hello").Parse(`
	<h1>Hello, {{.Name}}</h1>
`)
```

Then:

```go
tmpl.Execute(w, data)
```

This is convenient for small templates embedded directly in Go source.

For larger applications, template files are generally easier to maintain.

---

# 31. `Template.ParseFS`

### Signature

```go
func (t *Template) ParseFS(
	fs fs.FS,
	patterns ...string,
) (*Template, error)
```

This is the method version of `ParseFS`.

Example:

```go
tmpl := template.New("pages")

tmpl, err := tmpl.ParseFS(
	templateFS,
	"templates/*.html",
)
```

It is particularly useful with `embed.FS`.

---

# 32. `Template.ParseFiles`

### Signature

```go
func (t *Template) ParseFiles(
	filenames ...string,
) (*Template, error)
```

Parses named template files and associates them with the template set.

Example:

```go
tmpl := template.New("pages")

tmpl, err := tmpl.ParseFiles(
	"templates/header.html",
	"templates/home.html",
)
```

---

# 33. `Template.ParseGlob`

### Signature

```go
func (t *Template) ParseGlob(
	pattern string,
) (*Template, error)
```

Parses all files matching a glob pattern.

Example:

```go
tmpl, err := template.New("pages").
	ParseGlob("templates/*.html")
```

This is useful when you have many templates that should be loaded together.

---

# 34. `Template.Templates`

### Signature

```go
func (t *Template) Templates() []*Template
```

Returns the templates associated with a template, including itself.

For example:

```go
templates := tmpl.Templates()

for _, t := range templates {
	fmt.Println(t.Name())
}
```

This can be useful when inspecting or debugging a template collection.

---

# 35. Template Types You Should Know

The package also defines several important types, including:

```text
Template
FuncMap
HTML
HTMLAttr
JS
JSStr
CSS
URL
Srcset
Error
ErrorCode
```

The most important for beginners is `Template`.

Some specialized types are **trusted-content types**. For example:

```go
template.HTML
```

can tell the package:

> "This value is trusted HTML; don't escape it as ordinary text."

Example:

```go
template.HTML("<strong>Hello</strong>")
```

This can produce actual HTML rather than:

```html
&lt;strong&gt;Hello&lt;/strong&gt;
```

But this is extremely important:

**Do not convert arbitrary user input into `template.HTML`.**

Doing so can defeat the security protection that `html/template` provides. These trusted types introduce security risks and should only contain content from trusted/sanitized sources.

---

# 36. Automatic Context-Aware Escaping

This is arguably the most important feature of `html/template`.

Consider:

```html
<p>{{.}}</p>
```

The data is in an HTML-text context.

Now consider:

```html
<a href="{{.}}">
```

The data is in a URL/attribute context.

Or:

```html
<script>
	const name = {{.}};
</script>
```

The data is in a JavaScript context.

`html/template` analyzes the context and applies appropriate escaping.

This is why you should generally prefer:

```go
html/template
```

over:

```go
text/template
```

when your output is HTML.

---

# 37. `html/template` vs `text/template`

A common beginner question is:

> Why not just use `text/template` to generate HTML?

The answer is security.

`text/template` generates generic text and does **not** provide `html/template`'s contextual HTML escaping.

`html/template` is specifically designed for HTML and automatically protects interpolated data against relevant injection attacks.

### Rule of thumb

```text
HTML output
    ↓
html/template

Plain text / configuration
    ↓
text/template
```

---

# 38. Three Common Beginner Mistakes

## Mistake 1: Using `text/template` for HTML

Beginners sometimes write:

```go
import "text/template"
```

and use it to generate HTML.

### Problem

`text/template` doesn't provide `html/template`'s contextual HTML escaping.

### Better

```go
import "html/template"
```

---

## Mistake 2: Trusting `template.HTML`

A beginner may see:

```go
template.HTML(userInput)
```

and think:

> "This makes my HTML safe."

It doesn't.

It tells the template system that the content is **already trusted HTML**.

If `userInput` contains malicious HTML or JavaScript, you can undermine the protection provided by `html/template`.

### Avoid it

Only use trusted-content types when the source is genuinely trusted or has been appropriately sanitized.

---

## Mistake 3: Ignoring template errors

This is risky:

```go
tmpl.Execute(w, data)
```

without checking the error.

Execution can fail.

Better:

```go
if err := tmpl.Execute(w, data); err != nil {
	log.Println(err)
}
```

Also remember that template execution can produce **partial output before an error occurs**, so your HTTP error-handling strategy should account for that.

---

# 39. Two Real-World Applications

## Application 1: Server-Side Web Applications

Imagine an e-commerce website:

```text
Product database
       ↓
Go application
       ↓
Product struct
       ↓
html/template
       ↓
product.html
       ↓
Browser
```

A template could contain:

```html
<h1>{{.Name}}</h1>

<p>Price: ₹{{.Price}}</p>

<p>{{.Description}}</p>
```

The Go application retrieves product data and renders the HTML.

This approach works particularly well for traditional server-rendered applications.

---

## Application 2: Admin Dashboards

An internal dashboard might have:

```text
Users
Orders
Revenue
Reports
System status
```

Go can collect this information:

```go
type Dashboard struct {
	Users   int
	Orders  int
	Revenue float64
}
```

Then:

```html
<h1>Dashboard</h1>

<p>Users: {{.Users}}</p>
<p>Orders: {{.Orders}}</p>
<p>Revenue: ₹{{.Revenue}}</p>
```

`html/template` makes it straightforward to generate the HTML while keeping the presentation separate from Go business logic.

---

# 40. Three Progressive Exercises

## Exercise 1 — Beginner: User Profile

Create a Go HTTP server that uses `html/template` to render a user profile.

Your data should contain:

```text
Name
Email
Age
City
```

Create an HTML template that displays all four values.

### Requirements

- Use `html/template`.
- Use a Go struct.
- Use `template.Parse` or `template.New(...).Parse(...)`.
- Execute the template inside an HTTP handler.
- Serve the page on `localhost`.
- Properly handle template execution errors.

**Do not use hard-coded HTML values for the user information.**

---

## Exercise 2 — Intermediate: Product Catalog

Build a server-rendered product catalog.

Create a Go structure representing:

```text
Product
    Name
    Description
    Price
    InStock
```

Create a slice containing at least five products.

Your template should:

- Display all products using `range`.
- Display the price.
- Show `"Available"` when a product is in stock.
- Show `"Out of stock"` otherwise.
- Display a special message when there are no products.
- Use an appropriate HTML layout.

Also deliberately include product data containing characters such as:

```text
<
>
"
'
```

and observe how `html/template` handles them.

---

## Exercise 3 — Advanced: Reusable Template Layout

Build a small server-rendered website using multiple template files.

Create:

```text
templates/
    base.html
    header.html
    home.html
    profile.html
```

Implement:

- A reusable base layout.
- A reusable header.
- A home page.
- A profile page.
- Named templates using `define`.
- Template inheritance-like behavior using blocks/templates.
- A custom template function such as `formatDate`.
- Template loading using `ParseGlob` or `ParseFS`.
- `ExecuteTemplate` to select the correct page.
- `missingkey=error` during development.
- Proper error handling.

Finally, test the application with data containing malicious-looking HTML and URLs and verify that ordinary untrusted strings are escaped appropriately.

**Do not use `template.HTML` to bypass escaping.**

---

# 41. A Practical Learning Order

Since you're learning Go's standard library, I recommend learning `html/template` in this order:

```text
1. template.New()
       ↓
2. template.Parse()
       ↓
3. Template.Execute()
       ↓
4. {{.Field}}
       ↓
5. {{if}}
       ↓
6. {{range}}
       ↓
7. {{with}}
       ↓
8. {{define}} / {{template}}
       ↓
9. ParseFiles / ParseGlob
       ↓
10. ExecuteTemplate
       ↓
11. Funcs / FuncMap
       ↓
12. ParseFS + embed.FS
       ↓
13. Clone
       ↓
14. Option("missingkey=error")
       ↓
15. Contextual escaping & trusted types
```

The **security model and contextual escaping** are the concepts to pay special attention to—not just the syntax.

`html/template` is not merely an HTML formatting library; its central advantage over generic templating is that it treats the context in which data appears as a security boundary.

---

# Thought-Provoking Question

Suppose you are building a Go web application where **some HTML comes from your own trusted templates, while other content comes from users**.

**How would you design the boundary between trusted HTML and untrusted data so that developers cannot accidentally bypass `html/template`'s automatic escaping—and what could go wrong if that boundary were unclear?**
