# Go `fmt` Package — Detailed Explanation

The `fmt` package is one of the most commonly used packages in Go. It provides functions for **formatted input and output**, including printing values, reading user input, creating formatted strings, and constructing errors.

Official documentation: https://pkg.go.dev/fmt

---

## 1. What is the `fmt` package?

`fmt` stands for **format**.

```go
import "fmt"
```

It is part of Go's **standard library** and is mainly used for:

- Printing output to the terminal
- Taking input from the keyboard
- Formatting values
- Creating formatted strings
- Writing formatted data to files, buffers, network connections, etc.
- Creating formatted errors
- Parsing values from strings

The package can be viewed as:

```text
                 fmt PACKAGE
                     |
        +------------+------------+
        |            |            |
      Output        Input       Formatting
        |            |            |
     Print*        Scan*       Sprint*
     Fprint*       Fscan*      Sprintf*
```

### When do we commonly use it?

For example:

```go
fmt.Println("Hello World")
```

For user input:

```go
fmt.Scan(&name)
```

For formatted output:

```go
fmt.Printf("Name: %s, Age: %d\n", name, age)
```

For creating a string:

```go
message := fmt.Sprintf("Hello %s", name)
```

For creating an error:

```go
err := fmt.Errorf("user %s not found", name)
```

---

# 2. Simple Example

```go
package main

import "fmt"

func main() {

    name := "Chandu"
    age := 25
    salary := 45000.50

    // Print
    fmt.Print("Name: ")
    fmt.Print(name)

    // Println
    fmt.Println()

    // Printf
    fmt.Printf("Age: %d\n", age)

    // Formatted floating-point value
    fmt.Printf("Salary: %.2f\n", salary)

    // Create a string
    message := fmt.Sprintf(
        "Hello %s, you are %d years old.",
        name,
        age,
    )

    fmt.Println(message)
}
```

Output:

```text
Name: Chandu
Age: 25
Salary: 45000.50
Hello Chandu, you are 25 years old.
```

---

# 3. `fmt` Functions — Complete List

The `fmt` package contains these major functions:

### Printing

1. `Print`
2. `Printf`
3. `Println`

### String formatting

4. `Sprint`
5. `Sprintf`
6. `Sprintln`

### Writing to an `io.Writer`

7. `Fprint`
8. `Fprintf`
9. `Fprintln`

### Scanning from standard input

10. `Scan`
11. `Scanf`
12. `Scanln`

### Scanning from an `io.Reader`

13. `Fscan`
14. `Fscanf`
15. `Fscanln`

### Scanning from a string

16. `Sscan`
17. `Sscanf`
18. `Sscanln`

### Error creation

19. `Errorf`

### Byte-slice formatting

20. `Append`
21. `Appendf`
22. `Appendln`

### Formatting support

23. `FormatString`

---

# Part A — Printing Functions

## 4. `fmt.Print()`

### Purpose

`Print()` prints values to **standard output**, normally your terminal.

Syntax:

```go
fmt.Print(a ...any)
```

Example:

```go
package main

import "fmt"

func main() {

    name := "Chandu"
    age := 25

    fmt.Print(name)
    fmt.Print(age)
}
```

Output:

```text
Chandu25
```

Notice that `Print()` does **not automatically add a newline**.

### Multiple values

```go
fmt.Print("Name:", name, "Age:", age)
```

Output:

```text
Name:ChanduAge:25
```

`Print` has specific spacing behavior depending on its operands. For predictable readable lines, `Println` or `Printf` is often clearer.

### When useful?

Use `Print()` when you want simple output without automatically adding a newline.

---

# 5. `fmt.Println()`

`Println()` prints values and **always adds a newline**.

Syntax:

```go
fmt.Println(a ...any)
```

Example:

```go
fmt.Println("Hello")
fmt.Println("Go")
```

Output:

```text
Hello
Go
```

Multiple values:

```go
name := "Chandu"
age := 25

fmt.Println("Name:", name)
fmt.Println("Age:", age)
```

Output:

```text
Name: Chandu
Age: 25
```

### Difference

```go
fmt.Print("Hello")
fmt.Print("World")
```

Output:

```text
HelloWorld
```

Whereas:

```go
fmt.Println("Hello")
fmt.Println("World")
```

Output:

```text
Hello
World
```

---

# 6. `fmt.Printf()`

`Printf()` is one of the **most important functions in `fmt`**.

It allows you to control exactly how values are displayed.

Syntax:

```go
fmt.Printf(format string, a ...any)
```

Example:

```go
name := "Chandu"
age := 25

fmt.Printf("Name: %s, Age: %d\n", name, age)
```

Output:

```text
Name: Chandu, Age: 25
```

Here:

```text
%s → string
%d → decimal integer
```

### Example

```go
price := 1250.5678

fmt.Printf("Price: %.2f\n", price)
```

Output:

```text
Price: 1250.57
```

`%.2f` means:

```text
%   → formatting starts
.2  → 2 digits after decimal
f   → floating-point number
```

---

# 7. Important `Printf` Format Verbs

These are extremely important when learning `fmt`.

## General

| Verb | Meaning |
|---|---|
| `%v` | Default value |
| `%+v` | Value with field names for structs |
| `%#v` | Go-syntax representation |
| `%T` | Type |
| `%%` | Literal `%` |

Example:

```go
name := "Chandu"

fmt.Printf("%v\n", name)
fmt.Printf("%T\n", name)
fmt.Printf("%#v\n", name)
```

Output:

```text
Chandu
string
"Chandu"
```

---

## Integer verbs

| Verb | Meaning |
|---|---|
| `%b` | Binary |
| `%c` | Character/rune |
| `%d` | Decimal |
| `%o` | Octal |
| `%O` | Octal with `0o` |
| `%x` | Hexadecimal lowercase |
| `%X` | Hexadecimal uppercase |
| `%U` | Unicode format |

Example:

```go
number := 255

fmt.Printf("%d\n", number)
fmt.Printf("%b\n", number)
fmt.Printf("%o\n", number)
fmt.Printf("%x\n", number)
fmt.Printf("%X\n", number)
```

Output:

```text
255
11111111
377
ff
FF
```

---

## Floating-point verbs

| Verb | Meaning |
|---|---|
| `%f` | Decimal |
| `%e` | Scientific notation |
| `%E` | Scientific notation uppercase |
| `%g` | Compact representation |
| `%G` | Compact uppercase representation |
| `%x` | Hexadecimal floating point |
| `%X` | Uppercase hexadecimal floating point |

Example:

```go
value := 123.456789

fmt.Printf("%f\n", value)
fmt.Printf("%.2f\n", value)
fmt.Printf("%e\n", value)
```

Output:

```text
123.456789
123.46
1.234568e+02
```

---

## String verbs

| Verb | Meaning |
|---|---|
| `%s` | String |
| `%q` | Quoted string |
| `%x` | Hexadecimal bytes |
| `%X` | Uppercase hexadecimal bytes |

Example:

```go
name := "Chandu"

fmt.Printf("%s\n", name)
fmt.Printf("%q\n", name)
```

Output:

```text
Chandu
"Chandu"
```

---

## Boolean

```go
value := true

fmt.Printf("%t\n", value)
```

Output:

```text
true
```

---

## Type

```go
age := 25

fmt.Printf("%T\n", age)
```

Output:

```text
int
```

This is particularly useful while debugging.

---

# Part B — String Formatting Functions

These functions are similar to `Print`, `Println`, and `Printf`, except that they **return a string instead of directly printing it**.

---

# 8. `fmt.Sprint()`

`Sprint()` formats values and returns the result as a string.

```go
message := fmt.Sprint("Hello ", "Chandu")

fmt.Println(message)
```

Output:

```text
Hello Chandu
```

Think:

```text
Print  → displays
Sprint → creates string
```

---

# 9. `fmt.Sprintln()`

`Sprintln()` works like `Println()`, but returns a string.

```go
message := fmt.Sprintln("Hello", "Chandu")

fmt.Print(message)
```

Output:

```text
Hello Chandu
```

It also includes a newline.

---

# 10. `fmt.Sprintf()`

`Sprintf()` is extremely useful in real-world Go applications.

It formats values and **returns a formatted string**.

```go
name := "Chandu"
age := 25

message := fmt.Sprintf(
    "Name: %s, Age: %d",
    name,
    age,
)

fmt.Println(message)
```

Output:

```text
Name: Chandu, Age: 25
```

### Real-world example

Suppose you need to construct an API URL:

```go
userID := 101

url := fmt.Sprintf(
    "https://example.com/users/%d",
    userID,
)
```

Now `url` contains:

```text
https://example.com/users/101
```

So:

```text
Printf  → format + print
Sprintf → format + return string
```

---

# Part C — Writer Functions

The `F` in:

```text
Fprint
Fprintf
Fprintln
```

means you provide a destination implementing `io.Writer`.

The destination could be:

- Terminal
- File
- Buffer
- Network connection
- HTTP response
- Custom writer

---

# 11. `fmt.Fprint()`

Syntax:

```go
fmt.Fprint(w, a ...)
```

Example:

```go
package main

import (
    "fmt"
    "os"
)

func main() {

    fmt.Fprint(os.Stdout, "Hello Go")
}
```

Here:

```go
os.Stdout
```

is the destination.

---

# 12. `fmt.Fprintln()`

```go
fmt.Fprintln(os.Stdout, "Hello", "Go")
```

Output:

```text
Hello Go
```

It behaves like `Println`, but writes to the specified `io.Writer`.

---

# 13. `fmt.Fprintf()`

`Fprintf()` is essentially:

```text
Printf + destination
```

Example:

```go
fmt.Fprintf(
    os.Stdout,
    "Name: %s, Age: %d\n",
    "Chandu",
    25,
)
```

This becomes especially useful when writing to files.

Example:

```go
file, err := os.Create("output.txt")

if err != nil {
    return
}

defer file.Close()

fmt.Fprintf(file, "Name: %s\n", "Chandu")
fmt.Fprintf(file, "Age: %d\n", 25)
```

Now the formatted data is written to the file.

---

# Part D — Input Functions

The `Scan*` family reads input.

---

# 14. `fmt.Scan()`

`Scan()` reads values from **standard input**.

Example:

```go
package main

import "fmt"

func main() {

    var name string
    var age int

    fmt.Print("Enter name and age: ")

    fmt.Scan(&name, &age)

    fmt.Println("Name:", name)
    fmt.Println("Age:", age)
}
```

Input:

```text
Chandu 25
```

Output:

```text
Name: Chandu
Age: 25
```

Notice:

```go
&name
&age
```

The `&` is important because `Scan()` needs to store the input into those variables.

---

# 15. `fmt.Scanln()`

`Scanln()` is similar to `Scan()`, but it stops at a newline.

```go
var name string
var age int

fmt.Scanln(&name, &age)
```

Input:

```text
Chandu 25
```

It scans those values until the line ends.

---

# 16. `fmt.Scanf()`

`Scanf()` allows you to specify a format.

Example:

```go
var name string
var age int

fmt.Scanf("%s %d", &name, &age)
```

Input:

```text
Chandu 25
```

Now:

```text
name = "Chandu"
age  = 25
```

### Important distinction

```text
Scan   → simple scanning
Scanln → scanning until newline
Scanf  → format-controlled scanning
```

---

# Part E — `Fscan` Functions

These are the input equivalents of `Fprint`.

Instead of reading from standard input, they read from an `io.Reader`.

---

# 17. `fmt.Fscan()`

```go
fmt.Fscan(reader, &name, &age)
```

Example:

```go
file, err := os.Open("data.txt")

if err != nil {
    return
}

defer file.Close()

var name string
var age int

fmt.Fscan(file, &name, &age)
```

Here the input comes from the file.

---

# 18. `fmt.Fscanln()`

```go
fmt.Fscanln(reader, &name, &age)
```

It behaves similarly to `Scanln`, but reads from the specified `io.Reader`.

---

# 19. `fmt.Fscanf()`

```go
fmt.Fscanf(
    reader,
    "%s %d",
    &name,
    &age,
)
```

This provides formatted scanning from an `io.Reader`.

---

# Part F — Scanning Strings

Sometimes you don't want to read from the keyboard or a file.

You already have a string containing data.

That's where `Sscan*` functions are useful.

---

# 20. `fmt.Sscan()`

```go
input := "Chandu 25"

var name string
var age int

fmt.Sscan(input, &name, &age)

fmt.Println(name)
fmt.Println(age)
```

Output:

```text
Chandu
25
```

Think:

```text
Sscan = Scan from String
```

---

# 21. `fmt.Sscanln()`

```go
input := "Chandu 25"

var name string
var age int

fmt.Sscanln(input, &name, &age)
```

It behaves like `Scanln`, but the source is a string.

---

# 22. `fmt.Sscanf()`

This is formatted scanning from a string.

```go
input := "Chandu 25"

var name string
var age int

fmt.Sscanf(
    input,
    "%s %d",
    &name,
    &age,
)
```

After execution:

```text
name = Chandu
age  = 25
```

### Very useful for parsing text

```go
input := "101 Chandu"

var id int
var name string

fmt.Sscanf(input, "%d %s", &id, &name)
```

Now:

```text
id   = 101
name = Chandu
```

---

# Part G — Error Creation

# 23. `fmt.Errorf()`

`Errorf()` creates an `error` using formatting.

```go
err := fmt.Errorf(
    "user %d not found",
    101,
)
```

The resulting error represents:

```text
user 101 not found
```

Example:

```go
package main

import "fmt"

func findUser(id int) error {

    if id == 101 {
        return fmt.Errorf("user %d not found", id)
    }

    return nil
}
```

### Why is this important?

Go frequently uses:

```go
return value, error
```

So `fmt.Errorf()` is commonly used when an error needs dynamic information.

For example:

```go
return fmt.Errorf(
    "failed to open file: %s",
    filename,
)
```

### `%w`

`Errorf()` also supports `%w` for wrapping an error:

```go
return fmt.Errorf("database operation failed: %w", err)
```

This allows the original error to remain available for error inspection using Go's error-handling mechanisms.

---

# Part H — Byte-Slice Functions

These functions are useful when you already work with `[]byte`.

---

# 24. `fmt.Append()`

`Append()` formats values and appends the formatted result to an existing byte slice.

```go
buffer := []byte("Hello ")

buffer = fmt.Append(buffer, "Go")

fmt.Println(string(buffer))
```

Output:

```text
Hello Go
```

Think:

```text
Sprint → creates string
Append → appends formatted data to []byte
```

---

# 25. `fmt.Appendln()`

Similar to `Append()`, but adds spaces and a newline.

```go
buffer := []byte("Hello")

buffer = fmt.Appendln(buffer, "Go")

fmt.Print(string(buffer))
```

Output:

```text
Hello Go
```

---

# 26. `fmt.Appendf()`

`Appendf()` is the formatted version.

```go
buffer := []byte("User: ")

buffer = fmt.Appendf(
    buffer,
    "%s, Age: %d",
    "Chandu",
    25,
)

fmt.Println(string(buffer))
```

Output:

```text
User: Chandu, Age: 25
```

This can be useful when you're already working with `[]byte` and don't want to first create a separate formatted string.

---

# Part I — `fmt.FormatString()`

This is an advanced function.

```go
fmt.FormatString(state, verb)
```

It is primarily useful when implementing the `fmt.Formatter` interface.

You normally **do not need this function as a beginner**.

Its purpose is to reconstruct the formatting directive being used by `fmt`.

For example, when implementing custom formatting, the formatter receives information such as:

```text
flags
width
precision
verb
```

`FormatString()` helps reconstruct that formatting directive.

This is an **advanced `fmt` feature**, rather than a function you would normally use in everyday Go programs.

---

# 27. Important `fmt` Interfaces

The package also defines several interfaces/types that become important when you move beyond beginner-level formatting.

Important ones include:

```text
Formatter
GoStringer
Scanner
ScanState
State
Stringer
```

## `fmt.Stringer`

One of the most useful is:

```go
type Stringer interface {
    String() string
}
```

Suppose you have:

```go
type User struct {
    Name string
    Age  int
}
```

You can define:

```go
func (u User) String() string {
    return fmt.Sprintf(
        "User{Name: %s, Age: %d}",
        u.Name,
        u.Age,
    )
}
```

Then:

```go
user := User{
    Name: "Chandu",
    Age:  25,
}

fmt.Println(user)
```

can use your custom `String()` representation.

This is powerful because you can control how your own types are represented when formatted.

---

# 28. `fmt.Formatter`

This is more advanced.

```go
type Formatter interface {
    Format(f State, verb rune)
}
```

It allows a type to completely control its formatting behavior.

For example, you could make a custom type behave differently for:

```text
%v
%s
%d
%x
```

This is generally something you'll learn after understanding:

- structs
- methods
- interfaces
- `Stringer`
- pointers
- formatting verbs

---

# 29. `fmt.GoStringer`

`GoStringer` provides a custom representation for `%#v`.

```go
type GoStringer interface {
    GoString() string
}
```

Example conceptually:

```go
fmt.Printf("%#v", value)
```

If the value implements `GoStringer`, its `GoString()` method can control that representation.

---

# 30. `fmt.Scanner`

`Scanner` is related to custom scanning behavior.

```go
type Scanner interface {
    Scan(state ScanState, verb rune) error
}
```

It allows a custom type to define how it should be scanned from formatted input.

This is an advanced feature.

---

# 31. The Most Important `fmt` Family Pattern

A very useful way to remember the package is this:

| Purpose | Default | Formatted | Newline |
|---|---|---|---|
| Terminal output | `Print` | `Printf` | `Println` |
| String output | `Sprint` | `Sprintf` | `Sprintln` |
| Writer output | `Fprint` | `Fprintf` | `Fprintln` |
| Terminal input | `Scan` | `Scanf` | `Scanln` |
| Reader input | `Fscan` | `Fscanf` | `Fscanln` |
| String input | `Sscan` | `Sscanf` | `Sscanln` |
| Byte slice | `Append` | `Appendf` | `Appendln` |

This table is worth memorizing.

---

# 32. Three Common Beginner Mistakes

## Mistake 1 — Confusing `%d`, `%f`, and `%s`

Incorrect:

```go
name := "Chandu"

fmt.Printf("%d\n", name)
```

The format verb doesn't match the value.

Better:

```go
fmt.Printf("%s\n", name)
```

### Avoid it

Remember:

```text
%d → integer
%f → floating point
%s → string
%t → boolean
%T → type
%v → general/default value
```

---

## Mistake 2 — Forgetting `&` with `Scan`

Incorrect:

```go
var age int

fmt.Scan(age)
```

Correct:

```go
var age int

fmt.Scan(&age)
```

Why?

`Scan()` needs a location where it can store the value.

---

## Mistake 3 — Thinking `Print`, `Println`, and `Printf` are identical

They aren't.

```go
fmt.Print("Hello")
```

doesn't automatically add a newline.

```go
fmt.Println("Hello")
```

adds a newline.

```go
fmt.Printf("Hello %s\n", name)
```

uses formatting directives.

A good rule:

```text
Print   → simple output
Println → simple output + newline
Printf  → formatted output
```

---

# 33. Two Real-World Applications

## Application 1 — CLI Applications

Suppose you're creating a command-line application:

```go
fmt.Println("===== User Registration =====")

fmt.Print("Enter username: ")
fmt.Scan(&username)

fmt.Print("Enter age: ")
fmt.Scan(&age)

fmt.Printf(
    "Welcome %s! Your age is %d.\n",
    username,
    age,
)
```

The `fmt` package handles both:

```text
Input
  ↓
Scan
  ↓
Application logic
  ↓
Printf / Println
  ↓
Output
```

---

## Application 2 — Application Errors and Logging

Suppose an application cannot find a customer:

```go
customerID := 5001

err := fmt.Errorf(
    "customer with ID %d was not found",
    customerID,
)
```

Or you might format information for logs:

```go
fmt.Printf(
    "Request completed: method=%s status=%d\n",
    method,
    status,
)
```

This makes `fmt` useful throughout backend applications, CLI tools, utilities, testing/debugging code, and many other Go programs.

---

# 34. Three Progressively Challenging Exercises

## Exercise 1 — Student Information

Create a Go program that:

1. Declares variables for:
   - Student name
   - Age
   - Course
   - Marks
2. Prints the information using `fmt.Println()`.
3. Prints the same information using `fmt.Printf()`.
4. Display the marks with exactly two digits after the decimal point.

**Do not use solutions from the explanation above directly; write the program yourself.**

---

## Exercise 2 — Employee Salary Calculator

Create a program that:

1. Takes an employee's:
   - Name
   - Basic salary
   - Bonus
   - Tax percentage
2. Uses `fmt.Scan()` to receive the input.
3. Calculates:
   - Gross salary
   - Tax amount
   - Net salary
4. Uses `fmt.Printf()` to display a formatted salary report.
5. Uses `fmt.Sprintf()` to create a summary message and then prints that message.

Your output should be formatted like a small professional salary report.

---

## Exercise 3 — Text Data Parser

Create a program that receives a string in this format:

```text
101 Chandu 25 45000.50
```

Use `fmt.Sscanf()` to extract:

```text
ID
Name
Age
Salary
```

Then:

1. Store each value in the appropriate Go type.
2. Validate that all expected values were successfully scanned.
3. Create a formatted employee summary using `fmt.Sprintf()`.
4. Create an error using `fmt.Errorf()` when the input is invalid.
5. Display the final employee information using formatted output.

This exercise combines several important parts of `fmt` rather than using only one function.

---

# 35. What You Should Learn First

Although `fmt` contains many functions, **don't try to memorize all of them immediately**.

I'd recommend learning them in this order:

### Level 1 — Must know

```text
fmt.Print()
fmt.Println()
fmt.Printf()

fmt.Scan()
fmt.Scanln()
fmt.Scanf()
```

### Level 2 — Very important

```text
fmt.Sprint()
fmt.Sprintf()
fmt.Sprintln()

fmt.Errorf()
```

### Level 3 — File/reader/writer programming

```text
fmt.Fprint()
fmt.Fprintf()
fmt.Fprintln()

fmt.Fscan()
fmt.Fscanf()
fmt.Fscanln()
```

### Level 4 — Parsing strings

```text
fmt.Sscan()
fmt.Sscanf()
fmt.Sscanln()
```

### Level 5 — Advanced

```text
fmt.Append()
fmt.Appendf()
fmt.Appendln()
fmt.FormatString()

Formatter
Stringer
GoStringer
Scanner
```

That progression will make the package much easier to understand.

---

# 36. One Mental Model to Remember

The entire `fmt` package becomes much easier if you remember these four questions:

```text
Where is the data coming FROM?
Where is the data going TO?
Do I want FORMATTING?
Do I want a NEWLINE?
```

For example:

```text
Keyboard → Terminal
    ↓
  Scan → Println
```

```text
String → Variables
    ↓
  Sscan
```

```text
Variables → String
    ↓
  Sprintf
```

```text
Variables → File
    ↓
  Fprintf
```

```text
Variables → []byte
    ↓
  Appendf
```

Once you understand this pattern, the large number of `fmt` functions becomes much less confusing.

---

# 37. Thought-Provoking Question

Suppose you are building a large Go backend application.

**Would you use `fmt.Printf()` everywhere for logging and debugging, or would you choose another package/mechanism in some situations?**

Why might separating:

- human-readable terminal output,
- application logs,
- errors, and
- machine-readable data

be important as the application grows?

---

## Quick Summary

```text
fmt.Print()      → print values
fmt.Println()    → print values + newline
fmt.Printf()     → formatted output

fmt.Sprint()     → create string
fmt.Sprintln()   → create string + newline
fmt.Sprintf()    → create formatted string

fmt.Fprint()     → write to io.Writer
fmt.Fprintln()   → write + newline
fmt.Fprintf()    → formatted write

fmt.Scan()       → read standard input
fmt.Scanln()     → read input until newline
fmt.Scanf()      → formatted input

fmt.Fscan()      → read from io.Reader
fmt.Fscanln()    → read from reader until newline
fmt.Fscanf()     → formatted reader input

fmt.Sscan()      → scan from string
fmt.Sscanln()    → scan string until newline
fmt.Sscanf()     → formatted scanning from string

fmt.Errorf()     → create formatted error

fmt.Append()     → append formatted data to []byte
fmt.Appendln()   → append + newline
fmt.Appendf()    → formatted append

fmt.FormatString() → advanced formatting support
```

**Key idea:**

> `fmt` is primarily Go's standard-library package for formatting, printing, scanning, and creating formatted strings/errors.
