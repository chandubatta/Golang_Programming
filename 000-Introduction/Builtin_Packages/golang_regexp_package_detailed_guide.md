# Go `regexp` Package — Detailed Guide

## 1. What is the `regexp` package?

The Go standard library's `regexp` package provides support for **regular expressions (regex)** in Go.

A regular expression is a pattern used to describe text that you want to:

- Search for
- Validate
- Extract
- Replace
- Split

For example:

```text
Pattern: \d+
Text:    "My order number is 12345"
Match:   "12345"
```

In Go:

```go
package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`\d+`)

	result := re.FindString("My order number is 12345")

	fmt.Println(result)
}
```

Output:

```text
12345
```

### When is `regexp` commonly used?

Common applications include:

- Validating input formats
- Extracting information from text
- Searching log files
- Parsing structured strings
- Finding URLs, IP addresses, or IDs
- Cleaning or transforming text
- Replacing sensitive information
- Processing configuration files
- Building search/filter functionality

Go's `regexp` package operates on UTF-8 text and uses RE2-style regular-expression syntax.

---

# 2. Basic Example

Let's build a small example that extracts email addresses from a string.

```go
package main

import (
	"fmt"
	"regexp"
)

func main() {
	text := "Contact us at support@example.com or sales@example.org"

	re := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)

	emails := re.FindAllString(text, -1)

	fmt.Println(emails)
}
```

Output:

```text
[support@example.com sales@example.org]
```

Here:

```go
regexp.MustCompile(...)
```

creates a compiled regular expression.

And:

```go
re.FindAllString(text, -1)
```

finds all matching strings.

The `-1` means **return all matches**.

---

# 3. Understanding Regex Syntax

| Pattern | Meaning | Example |
|---|---|---|
| `.` | Any character | `a.c` → `abc`, `axc` |
| `*` | Zero or more | `ab*` |
| `+` | One or more | `ab+` |
| `?` | Zero or one | `colou?r` |
| `^` | Beginning of text/line | `^Hello` |
| `$` | End of text/line | `world$` |
| `[abc]` | One of a, b, c | `[aeiou]` |
| `[^abc]` | Anything except a, b, c | `[^0-9]` |
| `[0-9]` | Digit | `123` |
| `[a-z]` | Lowercase letter | `hello` |
| `\d` | Digit | `123` |
| `\w` | Word character | `hello_123` |
| `\s` | Whitespace | space/tab |
| `()` | Capturing group | `(hello)` |
| `|` | OR | `cat|dog` |

For example:

```text
^[0-9]{5}$
```

means:

> The entire string must contain exactly five digits.

---

# 4. Every Function in the `regexp` Package

The package contains several package-level functions plus the `Regexp` type and its methods.

## A. `regexp.Match`

```go
func Match(pattern string, b []byte) (matched bool, err error)
```

`Match` checks whether a byte slice contains a match for a regular expression.

Example:

```go
package main

import (
	"fmt"
	"regexp"
)

func main() {
	matched, err := regexp.Match(`Go`, []byte("I love Go"))

	fmt.Println(matched)
	fmt.Println(err)
}
```

Output:

```text
true
<nil>
```

If the regex itself is invalid, an error is returned.

```go
matched, err := regexp.Match(`(`, []byte("hello"))

fmt.Println(matched)
fmt.Println(err)
```

Use `Match` for **simple one-off checks**.

If you're going to reuse the pattern many times, use `Compile` instead.

---

## B. `regexp.MatchString`

```go
func MatchString(pattern string, s string) (matched bool, err error)
```

This is similar to `Match`, but accepts a `string`.

```go
matched, err := regexp.MatchString(`Go`, "I love Go")

fmt.Println(matched)
fmt.Println(err)
```

Output:

```text
true
<nil>
```

A common use:

```go
matched, err := regexp.MatchString(`^[0-9]+$`, "12345")

if err != nil {
	fmt.Println("Regex error:", err)
	return
}

fmt.Println(matched)
```

---

## C. `regexp.MatchReader`

```go
func MatchReader(pattern string, r io.RuneReader) (matched bool, err error)
```

This checks whether text read from an `io.RuneReader` contains a match.

It is useful when your input comes through a reader rather than directly as a string or byte slice.

---

## D. `regexp.QuoteMeta`

```go
func QuoteMeta(s string) string
```

This is useful when user-provided text should be treated as **literal text rather than regex syntax**.

For example:

```go
text := "hello.world"

pattern := regexp.QuoteMeta(text)

fmt.Println(pattern)
```

Result:

```text
hello\.world
```

Without `QuoteMeta`, the `.` means "any character."

With `QuoteMeta`, it means a literal period.

Example:

```go
userInput := "file.txt"

pattern := regexp.QuoteMeta(userInput)

re := regexp.MustCompile(pattern)

fmt.Println(re.MatchString("file.txt"))
```

---

## E. `regexp.Compile`

```go
func Compile(expr string) (*Regexp, error)
```

`Compile` converts a regex pattern into a compiled `*Regexp`.

```go
re, err := regexp.Compile(`\d+`)

if err != nil {
	fmt.Println("Invalid regex:", err)
	return
}

fmt.Println(re.MatchString("Order 123"))
```

Output:

```text
true
```

### Why compile?

If you repeatedly use the same regex, compile it once:

```go
re, err := regexp.Compile(`\d+`)
```

Then reuse it:

```go
re.MatchString(...)
re.FindString(...)
re.FindAllString(...)
```

---

## F. `regexp.MustCompile`

```go
func MustCompile(str string) *Regexp
```

`MustCompile` is similar to `Compile`, except it **panics if the regex is invalid**.

Example:

```go
re := regexp.MustCompile(`\d+`)
```

This is convenient when the regex is a fixed pattern known by the programmer.

Example:

```go
var emailRegex = regexp.MustCompile(
	`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`,
)
```

### Don't use it blindly with untrusted input

Avoid:

```go
re := regexp.MustCompile(userProvidedPattern)
```

because an invalid user-provided pattern will panic.

Use:

```go
re, err := regexp.Compile(userProvidedPattern)
```

instead.

---

## G. `regexp.CompilePOSIX`

```go
func CompilePOSIX(expr string) (*Regexp, error)
```

This compiles a regular expression using **POSIX ERE syntax** and uses **leftmost-longest** matching semantics instead of Go's normal leftmost-first behavior.

Example:

```go
re, err := regexp.CompilePOSIX(`a|ab`)

if err != nil {
	panic(err)
}

fmt.Println(re.FindString("ab"))
```

This is primarily useful when compatibility with POSIX regular-expression behavior is required.

For most Go applications, ordinary:

```go
regexp.Compile()
```

is what you'll want.

---

## H. `regexp.MustCompilePOSIX`

```go
func MustCompilePOSIX(str string) *Regexp
```

This is the panic-on-error version of `CompilePOSIX`.

```go
re := regexp.MustCompilePOSIX(`a|ab`)
```

Use it only when the expression is known to be valid and POSIX semantics are specifically desired.

---

# The `Regexp` Type

The central type in this package is:

```go
type Regexp struct
```

A `Regexp` represents a **compiled regular expression**.

For example:

```go
re := regexp.MustCompile(`\d+`)
```

Here:

```text
re
│
└── *Regexp
```

You can then use the compiled expression to perform many operations.

A `Regexp` is safe for concurrent use by multiple goroutines, except configuration methods such as `Longest`.

---

# I. `Match`

```go
func (re *Regexp) Match(b []byte) bool
```

Checks whether a byte slice contains a match.

```go
re := regexp.MustCompile(`Go`)

fmt.Println(re.Match([]byte("I love Go")))
```

Output:

```text
true
```

---

# J. `MatchString`

```go
func (re *Regexp) MatchString(s string) bool
```

Checks whether a string contains a match.

```go
re := regexp.MustCompile(`\d+`)

fmt.Println(re.MatchString("Order 123"))
fmt.Println(re.MatchString("Order ABC"))
```

Output:

```text
true
false
```

This is one of the most frequently used methods.

---

# K. `MatchReader`

```go
func (re *Regexp) MatchReader(r io.RuneReader) bool
```

Searches text obtained from an `io.RuneReader`.

This is useful when processing text through a reader rather than having the entire input as a string.

---

# L. `Find`

```go
func (re *Regexp) Find(b []byte) []byte
```

Returns the **first matching portion** of a byte slice.

```go
re := regexp.MustCompile(`\d+`)

result := re.Find([]byte("Order 123 and 456"))

fmt.Println(string(result))
```

Output:

```text
123
```

If there is no match, `nil` is returned.

---

# M. `FindString`

```go
func (re *Regexp) FindString(s string) string
```

Returns the first matching substring.

```go
re := regexp.MustCompile(`\d+`)

result := re.FindString("Order 123 and 456")

fmt.Println(result)
```

Output:

```text
123
```

Think:

```text
FindString → first matching string
```

---

# N. `FindAll`

```go
func (re *Regexp) FindAll(b []byte, n int) [][]byte
```

Finds multiple matches in a byte slice.

```go
re := regexp.MustCompile(`\d+`)

matches := re.FindAll([]byte("123 abc 456 xyz 789"), -1)

for _, match := range matches {
	fmt.Println(string(match))
}
```

Output:

```text
123
456
789
```

The `n` parameter controls the number of matches.

```text
n = 2  → at most two matches
n = -1 → all matches
```

---

# O. `FindAllString`

```go
func (re *Regexp) FindAllString(s string, n int) []string
```

Returns multiple matching strings.

```go
re := regexp.MustCompile(`\d+`)

matches := re.FindAllString("123 abc 456 xyz 789", -1)

fmt.Println(matches)
```

Output:

```text
[123 456 789]
```

This is very useful for extraction.

---

# P. `FindIndex`

```go
func (re *Regexp) FindIndex(b []byte) []int
```

Instead of returning the matching text, it returns the **byte indexes** of the first match.

```go
re := regexp.MustCompile(`Go`)

index := re.FindIndex([]byte("I love Go"))

fmt.Println(index)
```

You might get:

```text
[7 9]
```

meaning:

```text
start = 7
end   = 9
```

The indexes are byte positions, not necessarily Unicode character positions.

---

# Q. `FindStringIndex`

```go
func (re *Regexp) FindStringIndex(s string) []int
```

Same concept as `FindIndex`, but accepts a string.

```go
re := regexp.MustCompile(`Go`)

fmt.Println(re.FindStringIndex("I love Go"))
```

This is useful when you need to know **where** a match occurs.

---

# R. `FindAllIndex`

```go
func (re *Regexp) FindAllIndex(b []byte, n int) [][]int
```

Returns indexes for multiple matches.

Conceptually:

```text
input
 ↓
regex
 ↓
[ [start,end], [start,end], ... ]
```

Useful for text scanners and syntax/highlighting tools.

---

# S. `FindAllStringIndex`

```go
func (re *Regexp) FindAllStringIndex(s string, n int) [][]int
```

Finds all matching strings and returns their byte indexes.

Example:

```go
re := regexp.MustCompile(`\d+`)

indexes := re.FindAllStringIndex(
	"123 abc 456 xyz 789",
	-1,
)

fmt.Println(indexes)
```

---

# T. `FindSubmatch`

```go
func (re *Regexp) FindSubmatch(b []byte) [][]byte
```

Returns the first match plus its capturing groups.

Example:

```go
re := regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)

matches := re.FindSubmatch(
	[]byte("Date: 2026-09-09"),
)

for _, match := range matches {
	fmt.Println(string(match))
}
```

Output:

```text
2026-09-09
2026
09
09
```

The first element is the complete match.

---

# U. `FindStringSubmatch`

```go
func (re *Regexp) FindStringSubmatch(s string) []string
```

Same idea, but returns strings.

```go
re := regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)

result := re.FindStringSubmatch("Date: 2026-09-09")

fmt.Println(result)
```

Output:

```text
[2026-09-09 2026 09 09]
```

---

# V. `FindSubmatchIndex`

```go
func (re *Regexp) FindSubmatchIndex(b []byte) []int
```

Returns byte indexes for the full match and each capturing group.

This is useful when you need the **location** of each captured part.

---

# W. `FindStringSubmatchIndex`

```go
func (re *Regexp) FindStringSubmatchIndex(s string) []int
```

Same concept, but works with a string.

For:

```text
(\d{4})-(\d{2})-(\d{2})
```

you receive indexes for:

```text
full match
group 1
group 2
group 3
```

---

# X. `FindAllSubmatch`

```go
func (re *Regexp) FindAllSubmatch(b []byte, n int) [][][]byte
```

Finds multiple matches and their capturing groups.

For example, you could extract multiple dates:

```text
2026-09-09
2027-01-15
2028-12-31
```

along with their year, month, and day groups.

---

# Y. `FindAllStringSubmatch`

```go
func (re *Regexp) FindAllStringSubmatch(s string, n int) [][]string
```

String version.

Example:

```go
re := regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)

results := re.FindAllStringSubmatch(
	"2026-09-09 and 2027-01-15",
	-1,
)

for _, result := range results {
	fmt.Println(result)
}
```

Conceptually:

```text
[
    [2026-09-09 2026 09 09],
    [2027-01-15 2027 01 15]
]
```

---

# Z. `FindAllSubmatchIndex`

```go
func (re *Regexp) FindAllSubmatchIndex(b []byte, n int) [][]int
```

Returns the indexes of multiple matches and their capturing groups.

Useful when building a parser and needing exact byte locations.

---

# AA. `FindAllStringSubmatchIndex`

```go
func (re *Regexp) FindAllStringSubmatchIndex(s string, n int) [][]int
```

String version.

It returns positions for:

```text
complete match
group 1
group 2
group 3
...
```

---

# AB. `ReplaceAll`

```go
func (re *Regexp) ReplaceAll(src, repl []byte) []byte
```

Replaces every match with replacement bytes.

Example:

```go
re := regexp.MustCompile(`\d+`)

result := re.ReplaceAll(
	[]byte("User 123 bought 456 items"),
	[]byte("XXX"),
)

fmt.Println(string(result))
```

Output:

```text
User XXX bought XXX items
```

---

# AC. `ReplaceAllString`

```go
func (re *Regexp) ReplaceAllString(src, repl string) string
```

String version of `ReplaceAll`.

```go
re := regexp.MustCompile(`\d+`)

result := re.ReplaceAllString(
	"User 123 bought 456 items",
	"XXX",
)

fmt.Println(result)
```

Output:

```text
User XXX bought XXX items
```

---

# AD. `ReplaceAllLiteral`

```go
func (re *Regexp) ReplaceAllLiteral(src, repl []byte) []byte
```

Replaces matches using the replacement **literally**.

This matters when your replacement string contains `$`.

Normal regex replacement templates can give `$` special meaning. `ReplaceAllLiteral` treats the replacement as literal text.

---

# AE. `ReplaceAllLiteralString`

```go
func (re *Regexp) ReplaceAllLiteralString(src, repl string) string
```

String version of `ReplaceAllLiteral`.

Use it when you want the replacement string treated literally.

---

# AF. `ReplaceAllFunc`

```go
func (re *Regexp) ReplaceAllFunc(
	src []byte,
	repl func([]byte) []byte,
) []byte
```

This is useful when the replacement is generated dynamically.

Example:

```go
re := regexp.MustCompile(`\d+`)

result := re.ReplaceAllFunc(
	[]byte("10 20 30"),
	func(match []byte) []byte {
		return []byte("[" + string(match) + "]")
	},
)

fmt.Println(string(result))
```

Output:

```text
[10] [20] [30]
```

This is useful when replacement depends on the matched value.

---

# AG. `ReplaceAllStringFunc`

```go
func (re *Regexp) ReplaceAllStringFunc(
	src string,
	repl func(string) string,
) string
```

String version of `ReplaceAllFunc`.

Example:

```go
re := regexp.MustCompile(`[a-z]+`)

result := re.ReplaceAllStringFunc(
	"hello world",
	func(s string) string {
		return "[" + s + "]"
	},
)

fmt.Println(result)
```

Output:

```text
[hello] [world]
```

---

# AH. `Split`

```go
func (re *Regexp) Split(s string, n int) []string
```

Splits a string wherever the regex matches.

Example:

```go
re := regexp.MustCompile(`[,;]+`)

result := re.Split("apple,banana;orange", -1)

fmt.Println(result)
```

Output:

```text
[apple banana orange]
```

This is useful when multiple delimiters are possible.

---

# AI. `Expand`

```go
func (re *Regexp) Expand(
	dst []byte,
	template []byte,
	src []byte,
	match []int,
) []byte
```

`Expand` is an advanced function for constructing replacement text from capturing groups.

The `match` indexes should generally come from:

```go
re.FindSubmatchIndex(...)
```

For example, a regex could capture:

```text
first name
last name
```

and an expansion template could rearrange them.

This is useful when you need fine-grained control over replacement processing and allocations.

---

# AJ. `ExpandString`

```go
func (re *Regexp) ExpandString(
	dst []byte,
	template string,
	src string,
	match []int,
) []byte
```

This is the string-oriented version of `Expand`.

For example:

```text
Regex:
(\w+) (\w+)

Input:
John Smith

Template:
$2, $1

Output:
Smith, John
```

The important concept is:

```text
capturing groups → match indexes → template → expanded output
```

---

# AK. `LiteralPrefix`

```go
func (re *Regexp) LiteralPrefix() (
	prefix string,
	complete bool,
)
```

Returns a literal prefix that every match must begin with.

Example:

```go
re := regexp.MustCompile(`hello[0-9]+`)

prefix, complete := re.LiteralPrefix()

fmt.Println(prefix)
fmt.Println(complete)
```

The prefix is:

```text
hello
```

`complete` tells you whether the literal prefix represents the entire expression.

---

# AL. `Longest`

```go
func (re *Regexp) Longest()
```

Changes matching behavior so that the regex prefers the **leftmost-longest** match.

Normally, Go uses leftmost-first matching.

Example:

```go
re := regexp.MustCompile(`a(|b)`)

fmt.Println(re.FindString("ab"))

re.Longest()

fmt.Println(re.FindString("ab"))
```

Conceptually:

```text
Normal:
a

After Longest():
ab
```

Important: `Longest()` changes the `Regexp`, so it should not be called concurrently with other methods on that same `Regexp`.

---

# AM. `NumSubexp`

```go
func (re *Regexp) NumSubexp() int
```

Returns the number of capturing groups.

Example:

```go
re := regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)

fmt.Println(re.NumSubexp())
```

Output:

```text
3
```

---

# AN. `SubexpNames`

```go
func (re *Regexp) SubexpNames() []string
```

Returns the names of capturing groups.

Go supports named groups such as:

```go
(?P<year>\d{4})
```

Example:

```go
re := regexp.MustCompile(
	`(?P<year>\d{4})-(?P<month>\d{2})-(?P<day>\d{2})`,
)

fmt.Println(re.SubexpNames())
```

You can use the names to understand which group corresponds to which piece of data.

---

# AO. `SubexpIndex`

```go
func (re *Regexp) SubexpIndex(name string) int
```

Returns the index of a named capturing group.

Example:

```go
re := regexp.MustCompile(
	`(?P<year>\d{4})-(?P<month>\d{2})-(?P<day>\d{2})`,
)

fmt.Println(re.SubexpIndex("year"))
fmt.Println(re.SubexpIndex("month"))
fmt.Println(re.SubexpIndex("day"))
```

This is useful with named groups because you don't need to hard-code numeric group indexes.

---

# AP. `String`

```go
func (re *Regexp) String() string
```

Returns the original regular-expression text used to create the `Regexp`.

Example:

```go
re := regexp.MustCompile(`\d+`)

fmt.Println(re.String())
```

Output:

```text
\d+
```

Useful for logging or debugging.

---

# AQ. `Copy` — Deprecated

```go
func (re *Regexp) Copy() *Regexp
```

`Copy` creates another `Regexp`.

However, this method is **deprecated** for normal concurrent use. Modern Go allows a `Regexp` to be safely shared among goroutines.

The main remaining reason to use it is when you specifically need separate copies with different `Longest()` settings.

For normal applications, simply share the compiled `Regexp`.

---

# AR. `AppendText`

```go
func (re *Regexp) AppendText(b []byte) ([]byte, error)
```

This method appends the textual representation of the regex to a byte slice and implements `encoding.TextAppender`.

For normal beginner-level regex work, you won't use this often.

---

# AS. `MarshalText`

```go
func (re *Regexp) MarshalText() ([]byte, error)
```

Returns the textual representation of the regular expression.

This allows a `Regexp` to participate in Go's text-marshaling mechanisms.

Conceptually:

```text
Regexp
  ↓
MarshalText()
  ↓
[]byte containing regex text
```

---

# AT. `UnmarshalText`

```go
func (re *Regexp) UnmarshalText(text []byte) error
```

The opposite direction of `MarshalText`.

It reconstructs a regular expression from textual data.

Conceptually:

```text
regex text
   ↓
UnmarshalText()
   ↓
Regexp
```

This can be useful when regex configuration needs to be loaded from text-based formats.

---

# 5. Understanding the `Find*` Family

The large number of methods can initially look confusing.

A useful way to remember them is:

```text
Find
│
├── FindString
│
├── FindAll
│   └── FindAllString
│
├── FindIndex
│   └── FindStringIndex
│
├── FindSubmatch
│   └── FindStringSubmatch
│
└── FindAllSubmatch
    └── FindAllStringSubmatch
```

Think about the suffixes:

### `String`

Works with strings.

```go
FindString()
```

### `All`

Finds multiple matches.

```go
FindAllString()
```

### `Index`

Returns positions.

```go
FindStringIndex()
```

### `Submatch`

Returns capturing groups.

```go
FindStringSubmatch()
```

### `All + Submatch`

Finds multiple matches and their groups.

```go
FindAllStringSubmatch()
```

This naming system makes the API much easier to learn.

---

# 6. Common Mistakes Beginners Make

## Mistake 1: Recompiling the same regex repeatedly

Bad:

```go
for _, value := range values {
	re, _ := regexp.Compile(`\d+`)
	fmt.Println(re.MatchString(value))
}
```

You're repeatedly compiling the same pattern.

Better:

```go
re := regexp.MustCompile(`\d+`)

for _, value := range values {
	fmt.Println(re.MatchString(value))
}
```

Compile once and reuse the compiled `Regexp`.

---

## Mistake 2: Assuming `MatchString` means "the entire string matches"

Consider:

```go
re := regexp.MustCompile(`\d+`)

fmt.Println(re.MatchString("abc123xyz"))
```

The result is:

```text
true
```

Why?

Because `\d+` occurs somewhere inside the string.

If you want the **entire string** to consist of digits, anchor the expression:

```go
re := regexp.MustCompile(`^\d+$`)
```

Now:

```text
123      → true
abc123   → false
123abc   → false
```

This distinction is extremely important for validation.

---

## Mistake 3: Treating regex as a complete parser

Regex is powerful, but it isn't appropriate for every parsing problem.

Trying to parse complex nested programming-language syntax with a giant regex is usually a poor design.

Prefer dedicated parsers when the input has complicated structure.

Also remember that Go's regexp engine intentionally does **not** provide every feature found in backtracking regex engines; its RE2-based design prioritizes predictable linear-time behavior.

---

# 7. Two Real-World Applications

## Application 1: Log processing

Imagine a server produces:

```text
2026-09-09 INFO User 123 logged in
2026-09-09 ERROR User 456 failed authentication
2026-09-09 INFO User 789 logged out
```

You could use:

```go
re := regexp.MustCompile(`User (\d+)`)
```

to extract user IDs.

You could also search for:

```text
ERROR
```

to identify error messages.

A log-processing system can therefore use regex to extract:

```text
timestamp
log level
user ID
message
```

---

## Application 2: Data extraction from unstructured text

Suppose an application receives:

```text
Please contact support@example.com or call +91-9876543210.
```

Regex can extract:

```text
support@example.com
+91-9876543210
```

This is useful in:

- Web scraping
- Document processing
- Email processing
- Customer-support systems
- Data-cleaning pipelines
- Search systems

---

# 8. Three Progressive Exercises

## Exercise 1 — Beginner: Extract Numbers

Write a Go program that:

1. Creates a string containing several numbers.
2. Uses `regexp` to find every integer.
3. Prints all the numbers found.
4. Prints the number of matches.

Example input:

```text
"I bought 3 apples, 12 oranges, and 25 bananas."
```

Expected conceptual result:

```text
3
12
25
```

**Do not use `strconv` to identify the numbers; use `regexp`.**

---

## Exercise 2 — Intermediate: Parse Log Entries

Create a program that processes log entries such as:

```text
2026-09-09 ERROR user=123 action=login
2026-09-09 INFO user=456 action=logout
2026-09-10 ERROR user=789 action=payment
```

Your program should use a regular expression with **capturing groups** to extract:

- Date
- Log level
- User ID
- Action

Then print each extracted field separately.

Try to design your regex so that it can process multiple log entries.

---

## Exercise 3 — Advanced: Sensitive Data Redaction

Build a program that receives a block of text containing:

- Email addresses
- Phone numbers
- Credit-card-like numbers

For example:

```text
Customer: alice@example.com
Phone: +91-9876543210
Card: 4111-1111-1111-1111
```

Use the `regexp` package to replace sensitive values with safe placeholders.

The resulting text should conceptually look like:

```text
Customer: [EMAIL]
Phone: [PHONE]
Card: [CARD]
```

Requirements:

1. Use compiled regular expressions.
2. Use capturing groups where appropriate.
3. Process multiple occurrences.
4. Use replacement functions for at least one category.
5. Consider what should happen when the input contains malformed data.

**Do not provide or look for a solution until you have attempted the exercise yourself.**

---

# 9. A Practical Mental Model

When learning `regexp`, don't try to memorize every method individually.

Think about the workflow:

```text
             Regex Pattern
                   │
                   ▼
            regexp.Compile()
                   │
                   ▼
             *regexp.Regexp
                   │
       ┌───────────┼───────────┐
       ▼           ▼           ▼
     Match        Find       Replace
       │           │           │
       │       ┌───┼───┐       │
       │       ▼   ▼   ▼       │
       │      All Index Groups │
       │                       │
       └───────────┬───────────┘
                   ▼
                Result
```

The most important methods to master first are:

```go
regexp.MustCompile()
re.MatchString()
re.FindString()
re.FindAllString()
re.FindStringSubmatch()
re.FindAllStringSubmatch()
re.ReplaceAllString()
re.ReplaceAllStringFunc()
re.Split()
```

Once these become comfortable, the rest of the API becomes much easier to understand.

---

# 10. One Important Go-Specific Point

Go's regex engine deliberately favors **predictable performance**. The package guarantees linear-time execution with respect to input size, unlike many regex implementations that can suffer from catastrophic backtracking.

That makes `regexp` particularly attractive for applications processing potentially large or untrusted text.

---

# 11. Thought-Provoking Question

**Suppose you were building a web API that accepts user-provided regular expressions and uses them to search millions of lines of server logs. Why might Go's `regexp` implementation be a safer architectural choice than a regex engine that supports more advanced features such as arbitrary backtracking—and what functionality might you have to give up in exchange for that safety?**

---

## Quick Learning Checklist

Before moving on, make sure you can explain:

- What a regular expression is
- What the `regexp` package does
- The difference between `Compile` and `MustCompile`
- The difference between `MatchString` and `FindString`
- How `FindAllString` works
- What capturing groups are
- How `FindStringSubmatch` works
- How `FindStringIndex` works
- How regex replacement works
- The difference between `ReplaceAllString` and `ReplaceAllStringFunc`
- How `Split` works
- Why `QuoteMeta` is useful
- Why `^` and `$` matter for validation
- Why compiling once and reusing a `Regexp` is useful
- Why Go's RE2-based engine prioritizes predictable execution time
