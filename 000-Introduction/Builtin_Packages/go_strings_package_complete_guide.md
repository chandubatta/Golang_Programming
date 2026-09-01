# Go `strings` Package — Complete Learning Guide

> A detailed beginner-friendly guide to the Go `strings` package, including its purpose, functions, types, examples, common mistakes, real-world applications, and exercises.

---

## 1. What is the `strings` package?

The Go `strings` package is part of Go's **standard library** and provides functions for searching, splitting, joining, comparing, modifying, trimming, and building strings.

Import it with:

```go
import "strings"
```

The package is designed for manipulating **UTF-8 encoded strings**.

### Common uses

- Searching inside strings
- Checking prefixes and suffixes
- Splitting strings
- Joining string slices
- Replacing text
- Converting case
- Removing unwanted characters
- Comparing strings
- Processing Unicode text
- Building large strings efficiently
- Reading a string through an `io.Reader`
- Performing multiple replacements

### Simple example

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "  Chandu Batta  "

	name = strings.TrimSpace(name)
	name = strings.ToUpper(name)

	fmt.Println(name)
}
```

Output:

```text
CHANDU BATTA
```

---

# 2. Important concept before learning `strings`

Go strings are **immutable**.

For example:

```go
name := "hello"

name = strings.ToUpper(name)
```

`strings.ToUpper()` does not modify the original string in place. It returns a new string value.

Also remember:

```go
len("hello")
```

returns the number of **bytes**, not necessarily the number of Unicode characters.

For example:

```go
fmt.Println(len("hello"))
fmt.Println(len("世界"))
```

The second value is larger than 2 because `"世界"` uses multiple UTF-8 bytes.

When working with Unicode characters, understand the difference between:

- byte
- rune
- Unicode code point
- string

This distinction becomes important with functions such as `ContainsRune`, `IndexRune`, `Map`, and the `Func` variants.

---

# 3. Complete `strings` package function guide

## A. `Clone`

### Syntax

```go
strings.Clone(s)
```

### Purpose

Creates a fresh copy of a string.

```go
original := "Hello"
copy := strings.Clone(original)

fmt.Println(copy)
```

Output:

```text
Hello
```

### When useful

Usually, you **do not need** `Clone`.

It can be useful in memory-sensitive situations where you want to ensure a retained small string does not keep a much larger underlying string allocation alive.

```go
small := strings.Clone(largeString[:10])
```

The Go documentation recommends using `Clone` rarely and generally only when profiling indicates that it is useful.

---

## B. `Compare`

### Syntax

```go
strings.Compare(a, b)
```

### Purpose

Performs a lexicographical comparison.

Returns:

| Result | Meaning |
|---:|---|
| `-1` | `a < b` |
| `0` | `a == b` |
| `1` | `a > b` |

Example:

```go
fmt.Println(strings.Compare("apple", "banana"))
fmt.Println(strings.Compare("apple", "apple"))
fmt.Println(strings.Compare("banana", "apple"))
```

Output:

```text
-1
0
1
```

### Important

For normal comparisons, simply use:

```go
a == b
a < b
a > b
```

Those operators are generally clearer and faster. `Compare` is useful when you specifically need a three-way comparison result.

---

## C. `Contains`

### Syntax

```go
strings.Contains(s, substr)
```

### Purpose

Checks whether one string exists inside another.

```go
text := "I am learning Golang"

fmt.Println(strings.Contains(text, "Golang"))
```

Output:

```text
true
```

It is case-sensitive:

```go
strings.Contains("Hello", "hello")
```

returns:

```text
false
```

### Common use

```go
if strings.Contains(message, "error") {
	fmt.Println("Error found")
}
```

---

## D. `ContainsAny`

### Syntax

```go
strings.ContainsAny(s, chars)
```

### Purpose

Checks whether `s` contains **at least one Unicode code point** from `chars`.

```go
fmt.Println(strings.ContainsAny("hello", "xyz"))
```

Output:

```text
false
```

```go
fmt.Println(strings.ContainsAny("hello", "aeiou"))
```

Output:

```text
true
```

### Difference from `Contains`

```go
strings.Contains("hello", "ae")
```

asks:

> Does `"ae"` occur as a sequence?

Whereas:

```go
strings.ContainsAny("hello", "ae")
```

asks:

> Does the string contain either `a` or `e`?

---

## E. `ContainsFunc`

### Syntax

```go
strings.ContainsFunc(s, func(rune) bool)
```

### Purpose

Checks whether **any Unicode character** satisfies a condition.

Example:

```go
hasDigit := strings.ContainsFunc("hello123", func(r rune) bool {
	return r >= '0' && r <= '9'
})

fmt.Println(hasDigit)
```

Output:

```text
true
```

Useful when your condition cannot be expressed simply using `Contains` or `ContainsAny`.

---

## F. `ContainsRune`

### Syntax

```go
strings.ContainsRune(s, r)
```

### Purpose

Checks whether a specific Unicode code point exists in a string.

```go
fmt.Println(strings.ContainsRune("Golang", 'G'))
```

Output:

```text
true
```

Unicode example:

```go
fmt.Println(strings.ContainsRune("Hello 世界", '世'))
```

Output:

```text
true
```

### `ContainsRune` vs `Contains`

Use:

```go
Contains(s, "abc")
```

for a substring.

Use:

```go
ContainsRune(s, 'a')
```

for one Unicode character.

---

## G. `Count`

### Syntax

```go
strings.Count(s, substr)
```

### Purpose

Counts **non-overlapping** occurrences of a substring.

```go
count := strings.Count("banana", "a")

fmt.Println(count)
```

Output:

```text
3
```

Another example:

```go
fmt.Println(strings.Count("aaaa", "aa"))
```

Output:

```text
2
```

It does not count overlapping occurrences.

---

## H. `Cut`

### Syntax

```go
before, after, found := strings.Cut(s, sep)
```

### Purpose

Splits a string around the **first occurrence** of a separator.

```go
before, after, found := strings.Cut("name=Chandu", "=")

fmt.Println(before)
fmt.Println(after)
fmt.Println(found)
```

Output:

```text
name
Chandu
true
```

If the separator doesn't exist:

```go
before, after, found := strings.Cut("Hello", ":")

fmt.Println(before)
fmt.Println(after)
fmt.Println(found)
```

Output:

```text
Hello

false
```

### Why `Cut` is useful

It often produces cleaner code than `Split` when you only care about the first separator.

---

## I. `CutLast`

### Syntax

```go
before, after, found := strings.CutLast(s, sep)
```

### Purpose

Splits around the **last occurrence** of a separator.

```go
before, after, found := strings.CutLast("a/b/c", "/")

fmt.Println(before)
fmt.Println(after)
fmt.Println(found)
```

Output:

```text
a/b
c
true
```

Useful for extracting a filename from a path or processing the final delimiter.

---

## J. `CutPrefix`

### Syntax

```go
strings.CutPrefix(s, prefix)
```

### Purpose

Removes a prefix **only if it exists**.

```go
result, found := strings.CutPrefix("Bearer abc123", "Bearer ")

fmt.Println(result)
fmt.Println(found)
```

Output:

```text
abc123
true
```

If the prefix doesn't exist, the original string is returned and `found` is false.

---

## K. `CutSuffix`

### Syntax

```go
strings.CutSuffix(s, suffix)
```

### Purpose

Removes a suffix if it exists.

```go
result, found := strings.CutSuffix("file.txt", ".txt")

fmt.Println(result)
fmt.Println(found)
```

Output:

```text
file
true
```

---

## L. `EqualFold`

### Syntax

```go
strings.EqualFold(s, t)
```

### Purpose

Performs Unicode-aware case-insensitive comparison.

```go
fmt.Println(strings.EqualFold("GoLang", "golang"))
```

Output:

```text
true
```

Compare:

```go
"Go" == "go"
```

which is:

```text
false
```

while:

```go
strings.EqualFold("Go", "go")
```

is:

```text
true
```

`EqualFold` performs Unicode simple case folding. It should not be confused with every possible locale-sensitive form of case conversion.

---

## M. `Fields`

### Syntax

```go
strings.Fields(s)
```

### Purpose

Splits a string around runs of whitespace.

```go
text := "  Go   is   awesome  "

words := strings.Fields(text)

fmt.Println(words)
```

Output:

```text
[Go is awesome]
```

It handles multiple spaces, tabs, and other Unicode whitespace.

### Difference from `Split`

```go
strings.Split("Go   is   awesome", " ")
```

can produce empty elements.

`Fields` is usually better when you want words separated by whitespace.

---

## N. `FieldsFunc`

### Syntax

```go
strings.FieldsFunc(s, function)
```

### Purpose

Splits a string whenever your function says a character is a separator.

```go
text := "Go,Java;Python|Rust"

result := strings.FieldsFunc(text, func(r rune) bool {
	return r == ',' || r == ';' || r == '|'
})

fmt.Println(result)
```

Output:

```text
[Go Java Python Rust]
```

This is useful for custom tokenization.

---

## O. `FieldsFuncSeq`

### Syntax

```go
strings.FieldsFuncSeq(s, f)
```

### Purpose

Similar to `FieldsFunc`, but returns an iterator instead of constructing a slice.

```go
for word := range strings.FieldsFuncSeq("Go,Java;Python", func(r rune) bool {
	return r == ',' || r == ';'
}) {
	fmt.Println(word)
}
```

Useful when processing potentially large strings where you want to iterate over fields instead of immediately creating a complete slice.

---

## P. `FieldsSeq`

### Syntax

```go
strings.FieldsSeq(s)
```

### Purpose

Iterator-based equivalent of `Fields`.

```go
text := "Go is fast"

for word := range strings.FieldsSeq(text) {
	fmt.Println(word)
}
```

Output:

```text
Go
is
fast
```

---

## Q. `HasPrefix`

### Syntax

```go
strings.HasPrefix(s, prefix)
```

### Purpose

Checks whether a string starts with a particular prefix.

```go
fmt.Println(strings.HasPrefix("Golang", "Go"))
```

Output:

```text
true
```

Common example:

```go
if strings.HasPrefix(url, "https://") {
	fmt.Println("Secure URL")
}
```

---

## R. `HasSuffix`

### Syntax

```go
strings.HasSuffix(s, suffix)
```

### Purpose

Checks whether a string ends with a particular suffix.

```go
fmt.Println(strings.HasSuffix("main.go", ".go"))
```

Output:

```text
true
```

Common use:

```go
if strings.HasSuffix(filename, ".json") {
	fmt.Println("JSON file")
}
```

---

## S. `Index`

### Syntax

```go
strings.Index(s, substr)
```

### Purpose

Returns the byte index of the **first occurrence** of `substr`.

```go
index := strings.Index("Hello World", "World")

fmt.Println(index)
```

Output:

```text
6
```

If not found:

```text
-1
```

### Important

The result is a **byte index**, not necessarily a character/rune position.

This matters for Unicode.

---

## T. `IndexAny`

### Syntax

```go
strings.IndexAny(s, chars)
```

### Purpose

Returns the byte index of the first Unicode code point from `chars`.

```go
index := strings.IndexAny("hello", "aeiou")

fmt.Println(index)
```

Output:

```text
1
```

Because `e` occurs first.

---

## U. `IndexByte`

### Syntax

```go
strings.IndexByte(s, c)
```

### Purpose

Finds the first occurrence of a particular **byte**.

```go
fmt.Println(strings.IndexByte("golang", 'a'))
```

Output:

```text
4
```

Use this when you specifically want byte-level searching.

---

## V. `IndexFunc`

### Syntax

```go
strings.IndexFunc(s, f)
```

### Purpose

Finds the first position where your function returns `true`.

```go
index := strings.IndexFunc("abc123", func(r rune) bool {
	return r >= '0' && r <= '9'
})

fmt.Println(index)
```

Output:

```text
3
```

Useful for searching according to custom Unicode-aware rules.

---

## W. `IndexRune`

### Syntax

```go
strings.IndexRune(s, r)
```

### Purpose

Finds the byte index of the first occurrence of a Unicode code point.

```go
fmt.Println(strings.IndexRune("Hello 世界", '世'))
```

Useful when searching for one Unicode character.

---

## X. `Join`

### Syntax

```go
strings.Join(slice, separator)
```

### Purpose

Combines a slice of strings into one string.

```go
names := []string{"Go", "Java", "Python"}

result := strings.Join(names, ", ")

fmt.Println(result)
```

Output:

```text
Go, Java, Python
```

### Common applications

- CSV-like data
- URLs
- file paths
- SQL fragments
- log messages
- comma-separated lists

---

## Y. `LastIndex`

### Syntax

```go
strings.LastIndex(s, substr)
```

### Purpose

Finds the last occurrence of a substring.

```go
fmt.Println(strings.LastIndex("go/go/main.go", "/"))
```

Useful for finding the final path separator.

---

## Z. `LastIndexAny`

### Syntax

```go
strings.LastIndexAny(s, chars)
```

### Purpose

Finds the last occurrence of any Unicode code point from `chars`.

```go
fmt.Println(strings.LastIndexAny("hello world", "aeiou"))
```

---

## AA. `LastIndexByte`

### Syntax

```go
strings.LastIndexByte(s, c)
```

### Purpose

Finds the last occurrence of a byte.

```go
fmt.Println(strings.LastIndexByte("banana", 'a'))
```

Output:

```text
5
```

---

## AB. `LastIndexFunc`

### Syntax

```go
strings.LastIndexFunc(s, f)
```

### Purpose

Finds the last position where a function returns `true`.

```go
index := strings.LastIndexFunc("abc123xyz456", func(r rune) bool {
	return r >= '0' && r <= '9'
})

fmt.Println(index)
```

Useful for custom searches.

---

## AC. `Lines`

### Syntax

```go
strings.Lines(s)
```

### Purpose

Returns an iterator over the lines of a string.

```go
text := "Hello\nWorld\nGo"

for line := range strings.Lines(text) {
	fmt.Printf("%q\n", line)
}
```

The iterator preserves terminating newline characters where they exist.

This is useful for processing multiline text without first creating a complete slice of lines.

---

## AD. `Map`

### Syntax

```go
strings.Map(mapping, s)
```

### Purpose

Transforms every Unicode character using a function.

```go
result := strings.Map(func(r rune) rune {
	if r >= 'a' && r <= 'z' {
		return r - 32
	}

	return r
}, "hello")

fmt.Println(result)
```

Output:

```text
HELLO
```

A particularly useful feature is that returning a negative value removes the character.

For example:

```go
strings.Map(func(r rune) rune {
	if r == ' ' {
		return -1
	}

	return r
}, "hello world")
```

produces:

```text
helloworld
```

---

## AE. `Repeat`

### Syntax

```go
strings.Repeat(s, count)
```

### Purpose

Repeats a string.

```go
fmt.Println(strings.Repeat("Go ", 3))
```

Output:

```text
Go Go Go
```

Useful for:

- indentation
- separators
- test data
- formatting

A negative count causes a panic, as does a result-size overflow.

---

## AF. `Replace`

### Syntax

```go
strings.Replace(s, old, new, n)
```

### Purpose

Replaces up to `n` occurrences.

```go
text := "cat cat cat"

result := strings.Replace(text, "cat", "dog", 2)

fmt.Println(result)
```

Output:

```text
dog dog cat
```

If `n < 0`, all occurrences are replaced.

---

## AG. `ReplaceAll`

### Syntax

```go
strings.ReplaceAll(s, old, new)
```

### Purpose

Replaces every non-overlapping occurrence.

```go
text := "cat cat cat"

result := strings.ReplaceAll(text, "cat", "dog")

fmt.Println(result)
```

Output:

```text
dog dog dog
```

### Difference

```text
Replace       → control how many replacements
ReplaceAll    → replace everything
```

---

## AH. `Split`

### Syntax

```go
strings.Split(s, sep)
```

### Purpose

Splits a string around a separator.

```go
text := "Go,Java,Python"

languages := strings.Split(text, ",")

fmt.Println(languages)
```

Output:

```text
[Go Java Python]
```

### Important

`Split` does not remove empty elements.

```go
strings.Split("a,,b", ",")
```

produces:

```text
[a  b]
```

There is an empty element between the two commas.

---

## AI. `SplitAfter`

### Syntax

```go
strings.SplitAfter(s, sep)
```

### Purpose

Splits the string but **keeps the separator attached to each resulting part**.

```go
result := strings.SplitAfter("a,b,c", ",")

fmt.Printf("%q\n", result)
```

Conceptually:

```text
["a," "b," "c"]
```

---

## AJ. `SplitAfterN`

### Syntax

```go
strings.SplitAfterN(s, sep, n)
```

### Purpose

Same idea as `SplitAfter`, but limits the number of pieces.

```go
strings.SplitAfterN("a,b,c,d", ",", 2)
```

produces approximately:

```text
["a," "b,c,d"]
```

---

## AK. `SplitAfterSeq`

### Syntax

```go
strings.SplitAfterSeq(s, sep)
```

### Purpose

Iterator version of `SplitAfter`.

```go
for part := range strings.SplitAfterSeq("a,b,c", ",") {
	fmt.Printf("%q\n", part)
}
```

Useful when processing pieces lazily.

---

## AL. `SplitN`

### Syntax

```go
strings.SplitN(s, sep, n)
```

### Purpose

Splits into at most `n` pieces.

```go
result := strings.SplitN("name=Chandu=Batta", "=", 2)

fmt.Println(result)
```

Output:

```text
[name Chandu=Batta]
```

This is very useful for key-value parsing.

---

## AM. `SplitSeq`

### Syntax

```go
strings.SplitSeq(s, sep)
```

### Purpose

Iterator-based version of `Split`.

```go
for part := range strings.SplitSeq("Go,Java,Python", ",") {
	fmt.Println(part)
}
```

This lets you process pieces one at a time.

---

## AN. `Title`

### Syntax

```go
strings.Title(s)
```

### Important: Deprecated

`strings.Title` is **deprecated** in current Go documentation.

Older code may contain:

```go
strings.Title("hello world")
```

For new applications, don't build new code around it. For sophisticated Unicode-aware title casing, use the `golang.org/x/text/cases` package.

---

## AO. `ToLower`

### Syntax

```go
strings.ToLower(s)
```

### Purpose

Converts Unicode letters to lowercase.

```go
fmt.Println(strings.ToLower("HELLO WORLD"))
```

Output:

```text
hello world
```

Very common when normalizing user input:

```go
input := strings.ToLower(userInput)
```

---

## AP. `ToLowerSpecial`

### Syntax

```go
strings.ToLowerSpecial(unicode.SpecialCase, s)
```

### Purpose

Performs lowercase conversion using a specific Unicode special-case mapping.

Example:

```go
result := strings.ToLowerSpecial(unicode.TurkishCase, "I")

fmt.Println(result)
```

You need:

```go
import "unicode"
```

This is an advanced Unicode feature.

---

## AQ. `ToTitle`

### Syntax

```go
strings.ToTitle(s)
```

### Purpose

Converts Unicode letters to title case.

```go
fmt.Println(strings.ToTitle("hello world"))
```

This is different from manually implementing English title formatting.

---

## AR. `ToTitleSpecial`

### Syntax

```go
strings.ToTitleSpecial(unicode.SpecialCase, s)
```

### Purpose

Performs title-case conversion using a specified Unicode special-case mapping.

This is mainly useful for specialized Unicode processing.

---

## AS. `ToUpper`

### Syntax

```go
strings.ToUpper(s)
```

### Purpose

Converts Unicode letters to uppercase.

```go
fmt.Println(strings.ToUpper("hello world"))
```

Output:

```text
HELLO WORLD
```

Common application:

```go
country := strings.ToUpper(input)
```

---

## AT. `ToUpperSpecial`

### Syntax

```go
strings.ToUpperSpecial(unicode.SpecialCase, s)
```

### Purpose

Converts characters to uppercase using a specified Unicode special-case mapping.

This is an advanced function and is generally unnecessary for ordinary English text.

---

## AU. `ToValidUTF8`

### Syntax

```go
strings.ToValidUTF8(s, replacement)
```

### Purpose

Replaces invalid UTF-8 byte sequences with a replacement string.

This can be useful when processing external or potentially malformed byte data that has been converted into a Go string.

Conceptually:

```go
clean := strings.ToValidUTF8(input, "?")
```

The result contains valid UTF-8.

---

## AV. `Trim`

### Syntax

```go
strings.Trim(s, cutset)
```

### Purpose

Removes characters contained in `cutset` from **both ends** of a string.

```go
result := strings.Trim("!!!Hello!!!", "!")

fmt.Println(result)
```

Output:

```text
Hello
```

### Important misconception

`cutset` is a **set of characters**, not a substring.

For example:

```go
strings.Trim("abcHelloabc", "abc")
```

does not mean:

> Remove `"abc"` as a whole string.

It means:

> Keep removing `a`, `b`, or `c` from the beginning/end while they occur there.

---

## AW. `TrimFunc`

### Syntax

```go
strings.TrimFunc(s, f)
```

### Purpose

Removes leading and trailing Unicode characters for which the supplied function returns `true`.

```go
result := strings.TrimFunc("123Hello456", func(r rune) bool {
	return r >= '0' && r <= '9'
})

fmt.Println(result)
```

Output:

```text
Hello
```

---

## AX. `TrimLeft`

### Syntax

```go
strings.TrimLeft(s, cutset)
```

### Purpose

Removes characters from the **beginning** of a string.

```go
fmt.Println(strings.TrimLeft("!!!Hello!!!", "!"))
```

Output:

```text
Hello!!!
```

### Important

Don't confuse this with:

```go
strings.TrimPrefix()
```

`TrimLeft` removes characters from a set.

`TrimPrefix` removes one specific prefix.

---

## AY. `TrimLeftFunc`

### Syntax

```go
strings.TrimLeftFunc(s, f)
```

### Purpose

Removes leading Unicode characters satisfying a function.

```go
result := strings.TrimLeftFunc("123Hello", func(r rune) bool {
	return r >= '0' && r <= '9'
})

fmt.Println(result)
```

Output:

```text
Hello
```

---

## AZ. `TrimPrefix`

### Syntax

```go
strings.TrimPrefix(s, prefix)
```

### Purpose

Removes a specific prefix.

```go
fmt.Println(strings.TrimPrefix("HelloWorld", "Hello"))
```

Output:

```text
World
```

If the prefix doesn't exist, the original string is returned.

---

## BA. `TrimRight`

### Syntax

```go
strings.TrimRight(s, cutset)
```

### Purpose

Removes characters from the end of a string.

```go
fmt.Println(strings.TrimRight("Hello!!!", "!"))
```

Output:

```text
Hello
```

This operates on a **set of characters**, not a whole suffix.

---

## BB. `TrimRightFunc`

### Syntax

```go
strings.TrimRightFunc(s, f)
```

### Purpose

Removes trailing Unicode characters satisfying a custom condition.

```go
result := strings.TrimRightFunc("Hello123", func(r rune) bool {
	return r >= '0' && r <= '9'
})

fmt.Println(result)
```

Output:

```text
Hello
```

---

## BC. `TrimSpace`

### Syntax

```go
strings.TrimSpace(s)
```

### Purpose

Removes leading and trailing whitespace.

```go
input := "   Hello World   "

result := strings.TrimSpace(input)

fmt.Println(result)
```

Output:

```text
Hello World
```

This is one of the most commonly used functions in the package.

Especially useful when processing user input:

```go
name := strings.TrimSpace(userInput)
```

---

## BD. `TrimSuffix`

### Syntax

```go
strings.TrimSuffix(s, suffix)
```

### Purpose

Removes a specific suffix.

```go
result := strings.TrimSuffix("main.go", ".go")

fmt.Println(result)
```

Output:

```text
main
```

If `.go` isn't present at the end, the original string remains unchanged.

---

# 4. `strings.Builder`

The package also contains the `Builder` type.

```go
var builder strings.Builder
```

It is designed for efficiently building strings.

Instead of repeatedly doing:

```go
result := ""

result += "Hello "
result += "World "
result += "Go"
```

you can use:

```go
var b strings.Builder

b.WriteString("Hello ")
b.WriteString("World ")
b.WriteString("Go")

fmt.Println(b.String())
```

Output:

```text
Hello World Go
```

---

## `Builder.Cap()`

Returns the capacity of the builder's underlying buffer.

```go
var b strings.Builder

fmt.Println(b.Cap())
```

Useful mainly for performance-related code.

---

## `Builder.Grow(n)`

Ensures that the builder has enough capacity for at least `n` additional bytes.

```go
var b strings.Builder

b.Grow(100)

b.WriteString("Hello")

fmt.Println(b.String())
```

Use this when you know approximately how large the resulting string will be.

---

## `Builder.Len()`

Returns the current length in bytes.

```go
var b strings.Builder

b.WriteString("Hello")

fmt.Println(b.Len())
```

Output:

```text
5
```

---

## `Builder.Reset()`

Clears the builder.

```go
var b strings.Builder

b.WriteString("Hello")
b.Reset()

b.WriteString("World")

fmt.Println(b.String())
```

Output:

```text
World
```

---

## `Builder.String()`

Returns the accumulated string.

```go
var b strings.Builder

b.WriteString("Hello")
b.WriteString(" Go")

fmt.Println(b.String())
```

Output:

```text
Hello Go
```

---

## `Builder.Write()`

Writes bytes into the builder.

```go
var b strings.Builder

b.Write([]byte("Hello"))

fmt.Println(b.String())
```

---

## `Builder.WriteByte()`

Writes one byte.

```go
var b strings.Builder

b.WriteByte('A')
b.WriteByte('B')

fmt.Println(b.String())
```

Output:

```text
AB
```

---

## `Builder.WriteRune()`

Writes a Unicode code point.

```go
var b strings.Builder

b.WriteRune('世')
b.WriteRune('界')

fmt.Println(b.String())
```

Output:

```text
世界
```

---

## `Builder.WriteString()`

Writes a string.

```go
var b strings.Builder

b.WriteString("Hello")
b.WriteString(" World")

fmt.Println(b.String())
```

This is probably the most commonly used `Builder` method.

---

# 5. `strings.Reader`

`strings.Reader` allows a string to behave like a reader implementing interfaces such as `io.Reader`, `io.ReaderAt`, `io.ByteReader`, `io.RuneReader`, `io.Seeker`, and `io.WriterTo`.

Create one using:

```go
reader := strings.NewReader("Hello World")
```

---

## `NewReader`

```go
reader := strings.NewReader("Hello World")
```

Creates a reader backed by a string.

---

## `Reader.Len()`

Returns the number of **unread bytes**.

```go
reader := strings.NewReader("Hello")

fmt.Println(reader.Len())
```

---

## `Reader.Read()`

Reads bytes from the string.

```go
reader := strings.NewReader("Hello")

buffer := make([]byte, 5)

n, err := reader.Read(buffer)

fmt.Println(n)
fmt.Println(string(buffer))
fmt.Println(err)
```

This allows ordinary string data to work with APIs expecting an `io.Reader`.

---

## `Reader.ReadAt()`

Reads bytes from a specific offset without changing the normal reader position.

```go
reader := strings.NewReader("Hello World")

buffer := make([]byte, 5)

n, err := reader.ReadAt(buffer, 6)

fmt.Println(n)
fmt.Println(string(buffer))
fmt.Println(err)
```

---

## `Reader.ReadByte()`

Reads one byte.

```go
reader := strings.NewReader("ABC")

b, err := reader.ReadByte()

fmt.Println(string(b))
fmt.Println(err)
```

---

## `Reader.ReadRune()`

Reads one Unicode code point.

```go
reader := strings.NewReader("世界")

r, size, err := reader.ReadRune()

fmt.Println(string(r))
fmt.Println(size)
fmt.Println(err)
```

A Unicode character can occupy multiple bytes.

---

## `Reader.Reset()`

Resets the reader to a new string.

```go
reader := strings.NewReader("Hello")

reader.Reset("World")
```

Now the reader reads from:

```text
World
```

---

## `Reader.Seek()`

Moves the reader's current position.

```go
reader := strings.NewReader("Hello World")

position, err := reader.Seek(6, io.SeekStart)

fmt.Println(position)
fmt.Println(err)
```

---

## `Reader.Size()`

Returns the original size of the string in bytes.

Difference:

```text
Size() → original total size
Len()  → unread remaining size
```

---

## `Reader.UnreadByte()`

Moves back one byte after a successful byte read.

Useful when a parser reads a byte but needs to process it again.

---

## `Reader.UnreadRune()`

Moves back one Unicode code point after a successful `ReadRune()`.

Useful when parsing text character by character.

---

## `Reader.WriteTo()`

Writes the unread portion of the string to an `io.Writer`.

For example, it can write to:

```go
os.Stdout
```

or a file.

---

# 6. `strings.Replacer`

`Replacer` is useful when you need to perform **multiple string replacements** efficiently.

Create one using:

```go
r := strings.NewReplacer(
	"<", "&lt;",
	">", "&gt;",
	"&", "&amp;",
)
```

Then:

```go
result := r.Replace("<hello>")

fmt.Println(result)
```

The `Replacer` type is safe for concurrent use by multiple goroutines.

---

## `NewReplacer`

### Syntax

```go
strings.NewReplacer(old1, new1, old2, new2, ...)
```

Example:

```go
r := strings.NewReplacer(
	"Go", "Golang",
	"Java", "Java Language",
)
```

### Important

Arguments must come in pairs:

```text
old, new
old, new
old, new
```

An odd number of arguments causes a panic.

---

## `Replacer.Replace`

```go
result := r.Replace("I like Go and Java")
```

Performs all configured replacements.

---

## `Replacer.WriteString`

```go
r.WriteString(writer, "Hello <world>")
```

Performs replacements while writing the result to an `io.Writer`.

---

# 7. Quick classification of the `strings` package

## Searching

```text
Contains
ContainsAny
ContainsFunc
ContainsRune
Index
IndexAny
IndexByte
IndexFunc
IndexRune
LastIndex
LastIndexAny
LastIndexByte
LastIndexFunc
```

## Checking

```text
HasPrefix
HasSuffix
EqualFold
```

## Counting

```text
Count
```

## Splitting

```text
Fields
FieldsFunc
FieldsFuncSeq
FieldsSeq
Split
SplitN
SplitSeq
SplitAfter
SplitAfterN
SplitAfterSeq
Cut
CutLast
CutPrefix
CutSuffix
Lines
```

## Joining/building

```text
Join
Builder
```

## Replacing

```text
Replace
ReplaceAll
Replacer
```

## Case conversion

```text
ToLower
ToLowerSpecial
ToUpper
ToUpperSpecial
ToTitle
ToTitleSpecial
Title
```

## Trimming

```text
Trim
TrimFunc
TrimLeft
TrimLeftFunc
TrimPrefix
TrimRight
TrimRightFunc
TrimSpace
TrimSuffix
```

## Character transformation

```text
Map
```

## Repetition

```text
Repeat
```

## UTF-8

```text
ToValidUTF8
```

## String readers

```text
Reader
```

## Memory-related

```text
Clone
```

---

# 8. One practical example using multiple functions

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "   Go, Java, Python, Go   "

	// Remove surrounding whitespace.
	input = strings.TrimSpace(input)

	// Convert to lowercase.
	input = strings.ToLower(input)

	// Split into individual languages.
	languages := strings.Split(input, ",")

	// Clean each language.
	for i := range languages {
		languages[i] = strings.TrimSpace(languages[i])
	}

	fmt.Println("Languages:")

	for _, language := range languages {
		fmt.Println(language)
	}

	// Search.
	if strings.Contains(input, "go") {
		fmt.Println("Go was found")
	}

	// Count.
	fmt.Println("Go count:", strings.Count(input, "go"))

	// Join.
	result := strings.Join(languages, " | ")

	fmt.Println("Final:", result)
}
```

Output:

```text
Languages:
go
java
python
go
Go was found
Go count: 2
Final: go | java | python | go
```

Typical string-processing workflow:

```text
input
  ↓
Trim
  ↓
Normalize
  ↓
Split
  ↓
Process
  ↓
Search / Count
  ↓
Join
  ↓
output
```

---

# 9. Three common beginner mistakes

## Mistake 1: Thinking strings are modified in place

Beginners sometimes write:

```go
name := "hello"

strings.ToUpper(name)

fmt.Println(name)
```

They expect:

```text
HELLO
```

But `strings.ToUpper` returns a new string.

Correct:

```go
name = strings.ToUpper(name)
```

### Rule

Always check whether the function returns a new string and assign it when needed.

---

## Mistake 2: Confusing bytes with characters

Beginners often assume:

```go
len("世界") == 2
```

because there are two visible characters.

But `len(string)` counts bytes.

For Unicode-aware processing:

```go
runes := []rune("世界")

fmt.Println(len(runes))
```

This produces:

```text
2
```

Don't blindly assume a byte index is a character index.

---

## Mistake 3: Confusing `Trim` with `TrimPrefix`

Consider:

```go
strings.Trim("abcHelloabc", "abc")
```

`"abc"` is interpreted as a **set of characters**.

It does not mean:

> Remove the exact string `"abc"`.

For an exact prefix:

```go
strings.TrimPrefix("abcHello", "abc")
```

For an exact suffix:

```go
strings.TrimSuffix("Helloabc", "abc")
```

### Rule

```text
Trim        → character set
TrimPrefix  → exact prefix
TrimSuffix  → exact suffix
```

---

# 10. Two real-world applications

## Application 1: Processing HTTP/API input

Suppose an API receives:

```text
  BEARER abc123xyz
```

You might process it using:

```go
authorization := strings.TrimSpace(header)

token, found := strings.CutPrefix(authorization, "Bearer ")

if found {
	fmt.Println("Token:", token)
}
```

The `strings` package is useful for:

- headers
- query parameters
- tokens
- user input
- command-line arguments
- text preprocessing

---

## Application 2: Log and text processing

Imagine processing:

```text
2026-09-01 ERROR database connection failed
```

You could use:

```go
if strings.Contains(line, "ERROR") {
	fmt.Println("Error log found")
}
```

Then:

```go
fields := strings.Fields(line)
```

to obtain individual components.

You can combine:

```text
Fields
Contains
HasPrefix
Split
Cut
TrimSpace
ToLower
ToUpper
Count
```

to build:

- log processors
- configuration parsers
- command processors
- text-analysis tools

---

# 11. Three progressively challenging exercises

## Exercise 1 — Beginner: User Input Cleaner

Write a Go program that accepts a person's name containing:

- leading spaces
- trailing spaces
- inconsistent capitalization

Example input:

```text
"   chandu batta   "
```

Expected form:

```text
Chandu Batta
```

### Requirements

Use the `strings` package to:

- remove unnecessary surrounding whitespace
- perform appropriate case conversion
- split the name into words
- reconstruct the final name

**Do not use a regular-expression package.**

---

## Exercise 2 — Intermediate: Log Analyzer

Create a program that processes a multiline log string such as:

```text
INFO User logged in
ERROR Database connection failed
INFO Request completed
WARNING Memory usage is high
ERROR File not found
```

Your program should determine:

1. Number of `INFO` messages
2. Number of `ERROR` messages
3. Number of `WARNING` messages
4. Total number of log lines
5. Whether any database-related error occurred
6. The individual words contained in each error message

Try to solve the problem primarily using the `strings` package.

---

## Exercise 3 — Advanced: Command Parser

Build a command parser for input such as:

```text
"  CREATE USER name=Chandu email=chandu@example.com role=admin  "
```

Your program should:

1. Remove unnecessary whitespace.
2. Identify the command.
3. Extract the command name.
4. Extract all key-value pairs.
5. Handle arbitrary whitespace.
6. Handle multiple key-value pairs.
7. Detect malformed key-value pairs.
8. Handle values containing spaces when appropriate.
9. Make command matching case-insensitive.
10. Produce a structured representation of the command.

Conceptually:

```text
Command: CREATE
name: Chandu
email: chandu@example.com
role: admin
```

Try to solve it using combinations of:

```text
TrimSpace
Fields
EqualFold
HasPrefix
Contains
Cut
SplitN
Join
```

and other `strings` APIs where appropriate.

**Do not use regular expressions for this exercise.**

---

# 12. Important functions to learn first

Although the package is large, don't try to memorize everything at once.

## Level 1 — Beginner

```go
strings.TrimSpace()
strings.ToUpper()
strings.ToLower()
strings.Contains()
strings.HasPrefix()
strings.HasSuffix()
```

## Level 2 — Intermediate

```go
strings.Split()
strings.Join()
strings.Replace()
strings.ReplaceAll()
strings.Count()
strings.Index()
strings.LastIndex()
```

## Level 3 — More practical parsing

```go
strings.Cut()
strings.CutPrefix()
strings.CutSuffix()
strings.Fields()
strings.FieldsFunc()
strings.EqualFold()
```

## Level 4 — Advanced

```go
strings.Builder
strings.Reader
strings.Replacer
strings.Map()
strings.ContainsFunc()
strings.IndexFunc()
```

## Level 5 — Modern Go

```go
strings.Lines()
strings.FieldsSeq()
strings.FieldsFuncSeq()
strings.SplitSeq()
strings.SplitAfterSeq()
```

These iterator APIs are part of newer Go releases.

---

# 13. The big picture

The most important idea is that `strings` isn't just a collection of unrelated functions.

Think of it as a **string-processing toolkit**:

```text
                    strings
                       │
       ┌───────────────┼────────────────┐
       │               │                │
    Search          Modify           Split/Join
       │               │                │
 Contains          Replace          Split
 Index             ToUpper           Fields
 HasPrefix         ToLower           Join
 Count             Map               Cut
       │               │                │
       └───────────────┼────────────────┘
                       │
                  Advanced APIs
                       │
             ┌─────────┼─────────┐
             │         │         │
          Builder    Reader   Replacer
```

Once you understand these categories, the package becomes much easier to remember.

The current Go API includes numerous standalone functions plus methods on `Builder`, `Reader`, and `Replacer`, with newer iterator-oriented APIs added in recent Go releases.

---

# 14. Thought-provoking question

Suppose you are building a **high-traffic Go API** that receives millions of user-submitted strings every day.

You could normalize every string using operations such as `TrimSpace`, `ToLower`, `ReplaceAll`, `Split`, and `Map`.

**Would you apply all of these transformations automatically to every input, or would you carefully choose which transformations are necessary? Why?**

Think about:

- correctness
- Unicode
- memory allocations
- performance
- user intent
- whether transformations can change the meaning of the original data
- whether normalization belongs at input, storage, or comparison time

This question is important because effective string processing is not just about knowing functions—it is about knowing **when a transformation is appropriate**.
