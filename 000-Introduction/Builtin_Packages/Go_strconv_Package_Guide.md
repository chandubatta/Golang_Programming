# Go `strconv` Package

The Go `strconv` package is part of the standard library and is used primarily for converting values between their string representations and basic Go data types. It is especially useful when working with user input, command-line arguments, configuration values, environment variables, HTTP parameters, file formats, and serialization.

## 1. What is the `strconv` package?

`strconv` means **string conversion**.

Its main job is to perform conversions such as:

- string → int
- string → float
- string → bool
- string → uint
- string → complex
- int → string
- float → string
- bool → string
- uint → string
- complex → string

It also provides functions for:

- Formatting numbers.
- Parsing numbers.
- Quoting and unquoting Go strings.
- Working with quoted runes.
- Checking whether Unicode characters are printable/graphic.
- Efficiently appending formatted values to `[]byte`.

Import it with:

```go
import "strconv"
```

---

## 2. Why do we need `strconv`?

Go does not automatically convert a string into a number.

For example:

```go
age := "25"
```

Here `age` is a `string`, not an `int`.

You cannot simply do:

```go
result := age + 5
```

You need to explicitly convert it:

```go
ageInt, err := strconv.Atoi(age)
```

Now:

```go
ageInt + 5
```

is valid.

This explicit conversion is important because Go is a strongly typed language.

---

## 3. Main categories of `strconv`

| Category | Main functions |
|---|---|
| String → number/bool | `Atoi`, `ParseInt`, `ParseUint`, `ParseFloat`, `ParseComplex`, `ParseBool` |
| Number/bool → string | `Itoa`, `FormatInt`, `FormatUint`, `FormatFloat`, `FormatComplex`, `FormatBool` |
| Append conversions | `AppendInt`, `AppendUint`, `AppendFloat`, `AppendBool` |
| Quoting/unquoting | `Quote`, `Unquote`, `QuoteRune`, `UnquoteChar`, etc. |
| Character/string checks | `IsPrint`, `IsGraphic`, `CanBackquote` |

---

# 4. String → Integer conversions

## 4.1 `Atoi()`

```go
strconv.Atoi(s string) (int, error)
```

`Atoi` means **ASCII to integer**.

It converts a decimal string into an `int`.

Example:

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	age, err := strconv.Atoi("25")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(age)
	fmt.Printf("%T\n", age)
}
```

Output:

```text
25
int
```

### When to use it

Use `Atoi` when you simply need:

```text
decimal string → int
```

For example:

```text
"25"  → 25
"100" → 100
"-50" → -50
```

`Atoi` is a convenient form of `ParseInt` using base 10 and the native `int` size.

---

# 5. `ParseInt()`

```go
strconv.ParseInt(s, base, bitSize)
```

This is the more powerful integer parser.

Example:

```go
number, err := strconv.ParseInt("101", 2, 64)
```

Here:

- `"101"` → input string
- `2` → binary
- `64` → result should fit within 64 bits

Result:

```text
5
```

### Bases

`ParseInt` supports bases from 2 through 36, and `base == 0` allows Go-style prefixes to determine the base.

Examples:

```go
strconv.ParseInt("1010", 2, 64) // 10
strconv.ParseInt("42", 10, 64)  // 42
strconv.ParseInt("2A", 16, 64)  // 42
```

You can also use:

```go
strconv.ParseInt("0x2A", 0, 64)
```

With base `0`, Go determines the base from the prefix.

### `bitSize`

You can specify:

```text
0
8
16
32
64
```

For example:

```go
value, err := strconv.ParseInt("127", 10, 8)
```

This means the value must fit within an `int8`-sized signed integer range.

The returned type is still `int64`; `bitSize` tells `strconv` what range to validate.

---

# 6. `ParseUint()`

```go
strconv.ParseUint(s, base, bitSize)
```

`ParseUint` is used for **unsigned integers**.

Example:

```go
value, err := strconv.ParseUint("255", 10, 8)
```

Result:

```text
255
```

The return type is:

```text
uint64
```

### Difference between `ParseInt` and `ParseUint`

```go
strconv.ParseInt("-25", 10, 64)
```

is valid.

But:

```go
strconv.ParseUint("-25", 10, 64)
```

is invalid because unsigned integers cannot have a negative sign.

Use:

```text
ParseInt  → signed integers
ParseUint → unsigned integers
```

---

# 7. String → floating-point conversion

## `ParseFloat()`

```go
strconv.ParseFloat(s string, bitSize int) (float64, error)
```

It converts a string into a floating-point number.

Example:

```go
price, err := strconv.ParseFloat("99.95", 64)

if err != nil {
	fmt.Println("Error:", err)
	return
}

fmt.Println(price)
```

Output:

```text
99.95
```

### `bitSize`

You normally use:

```text
32
```

or:

```text
64
```

For example:

```go
strconv.ParseFloat("3.14159", 32)
```

or:

```go
strconv.ParseFloat("3.14159", 64)
```

The function returns `float64` in both cases. The `bitSize` controls the precision/range to which the input is interpreted.

### Special values

`ParseFloat` also recognizes values such as:

```text
NaN
Inf
Infinity
```

and signed forms where appropriate.

---

# 8. String → complex number

## `ParseComplex()`

```go
strconv.ParseComplex(s string, bitSize int) (complex128, error)
```

It converts a string into a complex number.

Example:

```go
value, err := strconv.ParseComplex("3+4i", 128)

if err != nil {
	fmt.Println("Error:", err)
	return
}

fmt.Println(value)
```

Output:

```text
(3+4i)
```

It can parse forms such as:

```text
3
4i
3+4i
3-4i
(3+4i)
```

The `bitSize` can be:

```text
64
128
```

The result is `complex128`, even when `bitSize` is 64.

---

# 9. String → Boolean

## `ParseBool()`

```go
strconv.ParseBool(str string) (bool, error)
```

It converts accepted textual representations into a Boolean.

Accepted true forms include:

```text
1
t
T
TRUE
true
True
```

Accepted false forms include:

```text
0
f
F
FALSE
false
False
```

Example:

```go
enabled, err := strconv.ParseBool("true")

if err != nil {
	fmt.Println("Invalid boolean")
	return
}

fmt.Println(enabled)
```

Output:

```text
true
```

This is useful when configuration values arrive as strings.

---

# 10. Integer → String conversions

## `Itoa()`

```go
strconv.Itoa(i int) string
```

`Itoa` means **integer to ASCII**.

Example:

```go
age := 25

text := strconv.Itoa(age)

fmt.Println(text)
fmt.Printf("%T\n", text)
```

Output:

```text
25
string
```

It is the simplest choice when converting an `int` to a decimal string.

Conceptually:

```go
strconv.Itoa(25)
```

is equivalent to:

```go
strconv.FormatInt(int64(25), 10)
```

---

# 11. `FormatInt()`

```go
strconv.FormatInt(i int64, base int) string
```

This converts an `int64` to a string using a specified base.

Example:

```go
fmt.Println(strconv.FormatInt(42, 10))
fmt.Println(strconv.FormatInt(42, 2))
fmt.Println(strconv.FormatInt(42, 16))
```

Output:

```text
42
101010
2a
```

Supported bases are:

```text
2 through 36
```

Characters above 9 are represented using lowercase letters.

---

# 12. `FormatUint()`

```go
strconv.FormatUint(i uint64, base int) string
```

This is the unsigned equivalent of `FormatInt`.

Example:

```go
number := uint64(255)

fmt.Println(strconv.FormatUint(number, 10))
fmt.Println(strconv.FormatUint(number, 16))
```

Output:

```text
255
ff
```

Use:

```text
FormatInt  → int64
FormatUint → uint64
```

---

# 13. Floating-point → String

## `FormatFloat()`

```go
strconv.FormatFloat(f, fmt, prec, bitSize)
```

This is one of the most important functions in `strconv`.

Example:

```go
value := 3.1415926535

text := strconv.FormatFloat(value, 'f', 2, 64)

fmt.Println(text)
```

Output:

```text
3.14
```

### Parameters

```go
FormatFloat(
	value,
	format,
	precision,
	bitSize,
)
```

For example:

```go
strconv.FormatFloat(3.14159, 'f', 2, 64)
```

means:

```text
value     = 3.14159
format    = 'f'
precision = 2
bitSize   = 64
```

### Common formats

| Format | Meaning |
|---|---|
| `'b'` | Binary exponent |
| `'e'` | Scientific notation |
| `'E'` | Scientific notation with uppercase E |
| `'f'` | Decimal without exponent |
| `'g'` | Compact general format |
| `'G'` | Compact general format with uppercase E |
| `'x'` | Hexadecimal floating-point |
| `'X'` | Uppercase hexadecimal floating-point |

### Precision `-1`

A particularly useful value is:

```go
strconv.FormatFloat(3.1415926535, 'f', -1, 64)
```

This requests the smallest number of digits necessary for `ParseFloat` to recover the value exactly.

---

# 14. Complex number → String

## `FormatComplex()`

```go
strconv.FormatComplex(c, fmt, prec, bitSize)
```

It converts a complex number to a string.

Example:

```go
value := complex(3.14, 2.71)

text := strconv.FormatComplex(value, 'f', 2, 128)

fmt.Println(text)
```

Output:

```text
(3.14+2.71i)
```

The supported `bitSize` values are:

```text
64   → complex64 precision
128  → complex128 precision
```

---

# 15. Boolean → String

## `FormatBool()`

```go
strconv.FormatBool(b bool) string
```

Example:

```go
result := strconv.FormatBool(true)

fmt.Println(result)
```

Output:

```text
true
```

It produces either:

```text
true
```

or:

```text
false
```

---

# 16. Append functions

The `Append...` functions are similar to the `Format...` functions, but instead of creating a new string, they append the formatted representation to an existing `[]byte`.

This can be useful when building output incrementally.

## `AppendInt()`

```go
strconv.AppendInt(dst []byte, i int64, base int) []byte
```

Example:

```go
data := []byte("ID=")

data = strconv.AppendInt(data, 123, 10)

fmt.Println(string(data))
```

Output:

```text
ID=123
```

## `AppendUint()`

```go
strconv.AppendUint(dst []byte, i uint64, base int) []byte
```

Example:

```go
data := []byte("Count=")

data = strconv.AppendUint(data, 500, 10)

fmt.Println(string(data))
```

Output:

```text
Count=500
```

## `AppendFloat()`

```go
strconv.AppendFloat(dst []byte, f float64, fmt byte, prec int, bitSize int) []byte
```

Example:

```go
data := []byte("Price=")

data = strconv.AppendFloat(data, 99.95, 'f', 2, 64)

fmt.Println(string(data))
```

Output:

```text
Price=99.95
```

## `AppendBool()`

```go
strconv.AppendBool(dst []byte, b bool) []byte
```

Example:

```go
data := []byte("Enabled=")

data = strconv.AppendBool(data, true)

fmt.Println(string(data))
```

Output:

```text
Enabled=true
```

---

# 17. Quoting functions

`strconv` is not only about numbers. It also provides functions for working with **Go string and character literal representations**.

## `Quote()`

```go
strconv.Quote(s string) string
```

It creates a double-quoted Go string literal.

Example:

```go
text := `Hello "Go"`

quoted := strconv.Quote(text)

fmt.Println(quoted)
```

The result contains appropriate escaping.

---

# 18. `QuoteToASCII()`

```go
strconv.QuoteToASCII(s string) string
```

This is similar to `Quote`, but non-ASCII characters are escaped.

Example:

```go
text := "Hello 世界"

fmt.Println(strconv.QuoteToASCII(text))
```

The Unicode characters are represented using Go escape sequences such as:

```text
\u...
```

This is useful when an ASCII-only representation is desired.

---

# 19. `QuoteToGraphic()`

```go
strconv.QuoteToGraphic(s string) string
```

This produces a quoted Go string while preserving characters considered **graphic** by Unicode where appropriate.

Example:

```go
text := "Hello ☺"

fmt.Println(strconv.QuoteToGraphic(text))
```

This is useful when you want readable Unicode graphic characters rather than escaping every non-ASCII character.

---

# 20. `QuoteRune()`

```go
strconv.QuoteRune(r rune) string
```

This creates a quoted Go character literal.

Example:

```go
result := strconv.QuoteRune('A')

fmt.Println(result)
```

Output:

```text
'A'
```

Another example:

```go
fmt.Println(strconv.QuoteRune('☺'))
```

Output:

```text
'☺'
```

---

# 21. `QuoteRuneToASCII()`

```go
strconv.QuoteRuneToASCII(r rune) string
```

This creates a quoted rune literal while escaping non-ASCII characters.

Example:

```go
result := strconv.QuoteRuneToASCII('☺')

fmt.Println(result)
```

It produces an ASCII escape representation such as:

```text
'\u263a'
```

---

# 22. `QuoteRuneToGraphic()`

```go
strconv.QuoteRuneToGraphic(r rune) string
```

This creates a quoted rune literal while preserving graphic Unicode characters where possible.

For example, a graphic Unicode character can remain readable rather than being unnecessarily escaped.

---

# 23. Append quoting functions

## `AppendQuote()`

```go
strconv.AppendQuote(dst []byte, s string) []byte
```

Appends a quoted string to an existing byte slice.

## `AppendQuoteToASCII()`

```go
strconv.AppendQuoteToASCII(dst []byte, s string) []byte
```

Appends an ASCII-oriented quoted string.

## `AppendQuoteToGraphic()`

```go
strconv.AppendQuoteToGraphic(dst []byte, s string) []byte
```

Appends a graphic-oriented quoted string.

## `AppendQuoteRune()`

```go
strconv.AppendQuoteRune(dst []byte, r rune) []byte
```

Appends a quoted rune.

## `AppendQuoteRuneToASCII()`

```go
strconv.AppendQuoteRuneToASCII(dst []byte, r rune) []byte
```

Appends an ASCII-oriented quoted rune.

## `AppendQuoteRuneToGraphic()`

```go
strconv.AppendQuoteRuneToGraphic(dst []byte, r rune) []byte
```

Appends a graphic-oriented quoted rune.

These functions are useful when constructing a `[]byte` buffer and avoiding repeated intermediate strings.

---

# 24. `Unquote()`

```go
strconv.Unquote(s string) (string, error)
```

`Unquote` performs the opposite operation of quoting.

Example:

```go
text, err := strconv.Unquote(`"Hello\nWorld"`)

if err != nil {
	fmt.Println(err)
	return
}

fmt.Println(text)
```

The escape sequence `\n` is interpreted as a newline.

`Unquote` understands Go-style double-quoted strings, single-quoted character literals, and backquoted strings according to its parsing rules.

---

# 25. `UnquoteChar()`

```go
strconv.UnquoteChar(
	s string,
	quote byte,
) (value rune, multibyte bool, tail string, err error)
```

This is a lower-level function.

It decodes **one character** from an escaped string or character literal.

It returns:

```text
value
multibyte
tail
error
```

Conceptually:

```text
input:
\"Hello

decoded character:
"

remaining:
Hello
```

The `quote` parameter tells the function whether it is parsing a single-quoted or double-quoted literal.

This function is more advanced and is generally useful when implementing your own parser.

---

# 26. `QuotedPrefix()`

```go
strconv.QuotedPrefix(s string) (string, error)
```

This function looks at the **beginning of a string** and extracts the first valid quoted Go string/character literal.

For example, given:

```text
"hello" remaining text
```

it can return:

```text
"hello"
```

without requiring the entire input to be one quoted value.

This is useful when parsing input that begins with a Go quoted literal.

---

# 27. Unicode-related functions

## `IsPrint()`

```go
strconv.IsPrint(r rune) bool
```

It checks whether a rune is considered printable according to Go's definition.

Example:

```go
fmt.Println(strconv.IsPrint('A'))
fmt.Println(strconv.IsPrint('\n'))
```

The first is printable; the newline is not.

This can be useful when processing text that should contain printable characters.

---

# 28. `IsGraphic()`

```go
strconv.IsGraphic(r rune) bool
```

This checks whether a Unicode character is considered **graphic**.

Graphic characters include Unicode categories such as:

- letters
- marks
- numbers
- punctuation
- symbols
- certain spaces

Example:

```go
fmt.Println(strconv.IsGraphic('A'))
fmt.Println(strconv.IsGraphic('☺'))
```

Both are graphic characters.

### `IsPrint` vs `IsGraphic`

```text
IsPrint()
    ↓
Go's printable-character definition

IsGraphic()
    ↓
Unicode graphic-character definition
```

They overlap significantly, but they are not identical concepts.

---

# 29. `CanBackquote()`

```go
strconv.CanBackquote(s string) bool
```

It determines whether a string can safely be represented as a Go raw string literal using backquotes:

```go
`some text`
```

Example:

```go
text := "Hello World"

if strconv.CanBackquote(text) {
	fmt.Println("Can use backquotes")
}
```

This is particularly useful when generating Go source code or working with Go literal representations.

---

# 30. `NumError`

`NumError` is an error type used by numeric parsing functions.

Its structure contains information such as:

```go
type NumError struct {
	Func string
	Num  string
	Err  error
}
```

For example:

```go
value, err := strconv.Atoi("abc")
```

The error contains information about:

- which function failed
- what input failed
- why it failed

You can inspect the underlying error with:

```go
var numErr *strconv.NumError

if errors.As(err, &numErr) {
	fmt.Println(numErr.Func)
	fmt.Println(numErr.Num)
	fmt.Println(numErr.Err)
}
```

This is useful for sophisticated error handling.

---

# 31. `ErrSyntax`

```go
strconv.ErrSyntax
```

This represents invalid syntax during conversion.

For example:

```go
_, err := strconv.Atoi("hello")
```

The string isn't a valid integer, so the underlying error can be `ErrSyntax`.

---

# 32. `ErrRange`

```go
strconv.ErrRange
```

This represents a value that is outside the permitted range.

For example, trying to put a huge value into an 8-bit integer can result in a range error.

Conceptually:

```text
"999999999999999999999999"
          ↓
      ParseInt
          ↓
      ErrRange
```

The distinction is:

```text
ErrSyntax
    ↓
Input isn't valid syntax

ErrRange
    ↓
Input is valid, but value is too large/small
```

---

# 33. Complete simple example

Here is a small program demonstrating several commonly used functions:

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {

	// String → int
	age, err := strconv.Atoi("25")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// String → float
	price, err := strconv.ParseFloat("99.99", 64)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// String → bool
	active, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Age:", age)
	fmt.Println("Price:", price)
	fmt.Println("Active:", active)

	// int → string
	ageText := strconv.Itoa(age)

	// float → string
	priceText := strconv.FormatFloat(price, 'f', 2, 64)

	// bool → string
	activeText := strconv.FormatBool(active)

	fmt.Println("Age text:", ageText)
	fmt.Println("Price text:", priceText)
	fmt.Println("Active text:", activeText)
}
```

The fundamental pattern is:

```text
INPUT STRING
     ↓
strconv.Parse...
     ↓
GO VALUE
```

And the reverse:

```text
GO VALUE
    ↓
strconv.Format...
    ↓
OUTPUT STRING
```

---

# 34. A useful mental model

Remember these pairs:

```text
String → int
    Atoi

int → String
    Itoa
```

```text
String → int64
    ParseInt

int64 → String
    FormatInt
```

```text
String → uint64
    ParseUint

uint64 → String
    FormatUint
```

```text
String → float64
    ParseFloat

float64 → String
    FormatFloat
```

```text
String → complex128
    ParseComplex

complex128 → String
    FormatComplex
```

```text
String → bool
    ParseBool

bool → String
    FormatBool
```

This makes the package much easier to remember.

---

# 35. Common mistake #1 — Ignoring the error

A beginner may write:

```go
age, _ := strconv.Atoi(input)
```

This throws away the error.

That's dangerous because the user might enter:

```text
abc
```

instead of:

```text
25
```

A better approach is:

```go
age, err := strconv.Atoi(input)

if err != nil {
	fmt.Println("Invalid age")
	return
}
```

### Rule

**Always consider the `error` returned by parsing functions.**

---

# 36. Common mistake #2 — Confusing `ParseInt`'s `base` and `bitSize`

Consider:

```go
strconv.ParseInt("1010", 2, 64)
```

The `2` means **binary**.

The `64` means the value must fit within a **64-bit signed integer range**.

They have completely different purposes.

```text
base
 ↓
How is the string written?

bitSize
 ↓
How large can the resulting number be?
```

---

# 37. Common mistake #3 — Assuming `ParseFloat(..., 32)` returns `float32`

This is a subtle beginner misconception.

Consider:

```go
value, err := strconv.ParseFloat("3.14", 32)
```

The return type is still:

```text
float64
```

The `32` controls the precision/range used for parsing; it does **not** change the function's declared return type.

If you actually need a `float32`, convert it explicitly:

```go
value32 := float32(value)
```

---

# 38. Real-world application #1 — Processing HTTP/query parameters

HTTP parameters commonly arrive as strings.

For example:

```text
/api/products?page=10&limit=20
```

Your Go application may receive:

```go
pageString := "10"
limitString := "20"
```

You can convert them:

```go
page, err := strconv.Atoi(pageString)
if err != nil {
	// invalid page
}

limit, err := strconv.Atoi(limitString)
if err != nil {
	// invalid limit
}
```

Now the application can perform numeric calculations.

This pattern is common in:

- REST APIs
- web applications
- pagination
- filtering
- search parameters

---

# 39. Real-world application #2 — Configuration and environment variables

Environment variables are strings.

For example:

```text
PORT=8080
DEBUG=true
TIMEOUT=30
```

A Go application may read:

```go
portString := os.Getenv("PORT")
```

which gives:

```text
"8080"
```

You can then use:

```go
port, err := strconv.Atoi(portString)
```

Similarly:

```go
debug, err := strconv.ParseBool(os.Getenv("DEBUG"))
```

This makes `strconv` useful in:

- Docker applications
- cloud deployments
- configuration systems
- command-line tools
- server applications

---

# 40. Three progressively challenging exercises

## Exercise 1 — User Age Converter

Write a Go program that:

1. Accepts an age as a string.
2. Converts it into an `int` using `strconv`.
3. Checks whether the conversion succeeded.
4. Prints the age.
5. Calculates the age after 10 years.
6. Converts the result back into a string.

**Do not use `fmt.Sscanf` or other parsing packages.**

---

## Exercise 2 — Configuration Parser

Create a program that receives these configuration values as strings:

```text
port = "8080"
timeout = "30"
debug = "true"
maxUsers = "500"
```

Your program should:

1. Convert `port` into an integer.
2. Convert `timeout` into an integer.
3. Convert `debug` into a Boolean.
4. Convert `maxUsers` into an unsigned integer.
5. Handle invalid input gracefully.
6. Display the resulting values and their Go types.

---

## Exercise 3 — Multi-Base Number Converter

Build a command-line program that accepts a number as a string and a target base.

The program should:

1. Accept numbers represented in decimal, binary, hexadecimal, or another supported base.
2. Parse the number using `strconv.ParseInt`.
3. Handle negative numbers.
4. Detect invalid digits.
5. Detect values outside the requested bit size.
6. Convert the resulting number into:
   - decimal
   - binary
   - hexadecimal
   - the user-selected target base
7. Use `strconv.FormatInt` for the output.
8. Provide meaningful error messages for invalid input.

This exercise will force you to understand the difference between **string representation, base, and integer bit size**.

---

# 41. Quick `strconv` cheat sheet

| Function | Purpose |
|---|---|
| `Atoi` | string → `int` |
| `Itoa` | `int` → string |
| `ParseBool` | string → bool |
| `ParseInt` | string → signed integer |
| `ParseUint` | string → unsigned integer |
| `ParseFloat` | string → floating-point |
| `ParseComplex` | string → complex number |
| `FormatBool` | bool → string |
| `FormatInt` | `int64` → string |
| `FormatUint` | `uint64` → string |
| `FormatFloat` | float → string |
| `FormatComplex` | complex → string |
| `AppendBool` | append bool representation |
| `AppendInt` | append signed integer |
| `AppendUint` | append unsigned integer |
| `AppendFloat` | append float |
| `Quote` | quote a string |
| `QuoteToASCII` | quote with ASCII escapes |
| `QuoteToGraphic` | quote preserving graphic characters |
| `QuoteRune` | quote a rune |
| `QuoteRuneToASCII` | ASCII-escaped quoted rune |
| `QuoteRuneToGraphic` | graphic quoted rune |
| `AppendQuote` | append quoted string |
| `AppendQuoteToASCII` | append ASCII-quoted string |
| `AppendQuoteToGraphic` | append graphic-quoted string |
| `AppendQuoteRune` | append quoted rune |
| `AppendQuoteRuneToASCII` | append ASCII-quoted rune |
| `AppendQuoteRuneToGraphic` | append graphic quoted rune |
| `Unquote` | decode a quoted Go literal |
| `UnquoteChar` | decode one quoted character |
| `QuotedPrefix` | extract a quoted literal at the beginning |
| `CanBackquote` | check raw-string-literal suitability |
| `IsPrint` | check printable rune |
| `IsGraphic` | check Unicode graphic rune |

---

# 42. The most important functions to master first

You do **not** need to memorize all of `strconv` immediately.

For a beginner, learn these first.

### Level 1

```go
strconv.Atoi()
strconv.Itoa()
```

### Level 2

```go
strconv.ParseInt()
strconv.ParseUint()
strconv.ParseFloat()
strconv.ParseBool()
```

### Level 3

```go
strconv.FormatInt()
strconv.FormatUint()
strconv.FormatFloat()
strconv.FormatBool()
```

### Level 4

```go
strconv.AppendInt()
strconv.AppendUint()
strconv.AppendFloat()
strconv.AppendBool()
```

### Advanced

```go
strconv.ParseComplex()
strconv.FormatComplex()

strconv.Quote()
strconv.Unquote()

strconv.UnquoteChar()
strconv.QuotedPrefix()

strconv.IsPrint()
strconv.IsGraphic()
strconv.CanBackquote()
```

That progression provides a manageable learning path.

---

# Thought-provoking question 🤔

Suppose your Go web application receives **all configuration and HTTP input as strings**.

Why do you think Go deliberately makes you explicitly parse those strings instead of automatically converting `"123"` into `123`?

Consider what could go wrong if automatic conversion were allowed—for example with:

- invalid input
- integer overflow
- different numeric bases
- Boolean representations
- loss of floating-point precision

**What advantages does explicit conversion with `strconv` give you as a programmer?**

---

## Summary

The `strconv` package is fundamentally about **converting values between Go data types and their textual representations**.

The most important concepts to remember are:

```text
Parsing:
String → Go value
```

```text
Formatting:
Go value → String
```

And the most commonly used functions are:

```text
Atoi
Itoa
ParseInt
ParseUint
ParseFloat
ParseBool
FormatInt
FormatUint
FormatFloat
FormatBool
```

Once these become comfortable, move on to the append and quoting functions.
