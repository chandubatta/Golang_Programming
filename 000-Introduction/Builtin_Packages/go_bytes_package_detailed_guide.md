# Go `bytes` Package — Detailed Guide

## 1. What is the `bytes` package?

The Go `bytes` package provides functions and types for **working with byte slices (`[]byte`)**.

```go
import "bytes"
```

Its purpose is similar to the `strings` package, but instead of primarily operating on strings, it operates on **raw byte data**.

It is commonly used to:

- search inside byte slices
- compare byte slices
- split byte data
- replace portions of byte data
- trim unwanted bytes
- change case
- join byte slices
- validate/transform UTF-8
- efficiently build byte data
- read byte data like a stream
- write byte data into buffers

It is particularly useful when working with **files, network protocols, HTTP bodies, binary formats, serialization, parsers, and I/O**.

---

# 2. Why `[]byte` instead of `string`?

A Go string is immutable:

```go
name := "Chandu"
```

A byte slice can be modified:

```go
name := []byte("Chandu")
name[0] = 'X'

fmt.Println(string(name))
```

Output:

```text
Xhandu
```

Therefore, `[]byte` is commonly used when data needs to be **read, modified, assembled, or transmitted at the byte level**.

---

# 3. Simple example

```go
package main

import (
	"bytes"
	"fmt"
)

func main() {
	data := []byte("Hello, Go Programming!")

	// Check whether data contains "Go"
	fmt.Println(bytes.Contains(data, []byte("Go")))

	// Find the position of "Programming"
	fmt.Println(bytes.Index(data, []byte("Programming")))

	// Replace "Go" with "Golang"
	result := bytes.ReplaceAll(
		data,
		[]byte("Go"),
		[]byte("Golang"),
	)

	fmt.Println(string(result))

	// Convert to uppercase
	fmt.Println(string(bytes.ToUpper(data)))
}
```

Output:

```text
true
7
Hello, Golang Programming!
HELLO, GO PROGRAMMING!
```

---

# 4. Complete `bytes` package API

The `bytes` package contains package-level functions plus the important `bytes.Buffer` and `bytes.Reader` types.

## A. Comparison Functions

### 4.1 `bytes.Equal`

```go
func Equal(a, b []byte) bool
```

Checks whether two byte slices contain exactly the same bytes.

```go
a := []byte("hello")
b := []byte("hello")

fmt.Println(bytes.Equal(a, b))
```

Output:

```text
true
```

Important: `nil` and an empty byte slice are treated as equivalent by `bytes.Equal`.

---

### 4.2 `bytes.Compare`

```go
func Compare(a, b []byte) int
```

Lexicographically compares two byte slices.

Results:

```text
-1   a < b
 0   a == b
+1   a > b
```

Example:

```go
fmt.Println(bytes.Compare(
	[]byte("apple"),
	[]byte("banana"),
))
```

Output:

```text
-1
```

---

### 4.3 `bytes.EqualFold`

```go
func EqualFold(s, t []byte) bool
```

Compares byte slices without considering Unicode case differences.

```go
a := []byte("HELLO")
b := []byte("hello")

fmt.Println(bytes.EqualFold(a, b))
```

Output:

```text
true
```

Useful when comparison should be case-insensitive.

---

# B. Searching Functions

## 4.4 `bytes.Contains`

```go
func Contains(b, subslice []byte) bool
```

Checks whether one byte slice contains another.

```go
data := []byte("Hello Golang")

fmt.Println(bytes.Contains(data, []byte("Go")))
```

Output:

```text
true
```

---

## 4.5 `bytes.ContainsAny`

```go
func ContainsAny(b []byte, chars string) bool
```

Checks whether the byte slice contains **any character from a string**.

```go
data := []byte("hello")

fmt.Println(bytes.ContainsAny(data, "xyz"))
fmt.Println(bytes.ContainsAny(data, "ae"))
```

Output:

```text
false
true
```

---

## 4.6 `bytes.ContainsRune`

```go
func ContainsRune(b []byte, r rune) bool
```

Checks whether a byte slice contains a particular Unicode code point.

```go
data := []byte("Golang")

fmt.Println(bytes.ContainsRune(data, 'G'))
```

Output:

```text
true
```

---

## 4.7 `bytes.ContainsFunc`

```go
func ContainsFunc(b []byte, f func(rune) bool) bool
```

Checks whether any Unicode character satisfies a function.

```go
data := []byte("hello123")

result := bytes.ContainsFunc(data, func(r rune) bool {
	return r >= '0' && r <= '9'
})

fmt.Println(result)
```

Output:

```text
true
```

Useful for questions such as:

- Does this data contain a digit?
- Does it contain uppercase characters?
- Does it contain a particular class of Unicode characters?

---

# C. Counting

## 4.8 `bytes.Count`

```go
func Count(s, sep []byte) int
```

Counts non-overlapping occurrences of `sep`.

```go
data := []byte("banana")

fmt.Println(bytes.Count(data, []byte("a")))
```

Output:

```text
3
```

Another example:

```go
data := []byte("Go Go Go")

fmt.Println(bytes.Count(data, []byte("Go")))
```

Output:

```text
3
```

Special case: if `sep` is empty, `Count` returns one more than the number of UTF-8 code points in `s`.

---

# D. Cutting Data

## 4.9 `bytes.Cut`

```go
func Cut(s, sep []byte) (before, after []byte, found bool)
```

Splits data around the **first occurrence** of a separator.

```go
data := []byte("name=Chandu")

before, after, found := bytes.Cut(data, []byte("="))

fmt.Println(string(before))
fmt.Println(string(after))
fmt.Println(found)
```

Output:

```text
name
Chandu
true
```

Useful for parsing:

```text
key=value
header:value
username:password
```

---

## 4.10 `bytes.CutPrefix`

```go
func CutPrefix(s, prefix []byte) (after []byte, found bool)
```

Removes a prefix if it exists.

```go
data := []byte("Bearer abc123")

result, found := bytes.CutPrefix(
	data,
	[]byte("Bearer "),
)

fmt.Println(string(result))
fmt.Println(found)
```

Output:

```text
abc123
true
```

---

## 4.11 `bytes.CutSuffix`

```go
func CutSuffix(s, suffix []byte) (before []byte, found bool)
```

Removes a suffix if it exists.

```go
data := []byte("photo.jpg")

result, found := bytes.CutSuffix(
	data,
	[]byte(".jpg"),
)

fmt.Println(string(result))
fmt.Println(found)
```

Output:

```text
photo
true
```

---

## 4.12 `bytes.CutLast`

```go
func CutLast(s, sep []byte) (before, after []byte, found bool)
```

Splits around the **last occurrence** of a separator.

```go
data := []byte("a/b/c")

before, after, found := bytes.CutLast(
	data,
	[]byte("/"),
)

fmt.Println(string(before))
fmt.Println(string(after))
fmt.Println(found)
```

Output:

```text
a/b
c
true
```

Useful for file paths or hierarchical identifiers.

---

# E. Prefix and Suffix

## 4.13 `bytes.HasPrefix`

```go
func HasPrefix(s, prefix []byte) bool
```

Checks whether data starts with a specific prefix.

```go
data := []byte("https://example.com")

fmt.Println(bytes.HasPrefix(data, []byte("https://")))
```

Output:

```text
true
```

---

## 4.14 `bytes.HasSuffix`

```go
func HasSuffix(s, suffix []byte) bool
```

Checks whether data ends with a specific suffix.

```go
data := []byte("image.png")

fmt.Println(bytes.HasSuffix(data, []byte(".png")))
```

Output:

```text
true
```

---

# F. Finding Positions

## 4.15 `bytes.Index`

```go
func Index(s, sep []byte) int
```

Returns the index of the first occurrence.

```go
data := []byte("Hello Go")

fmt.Println(bytes.Index(data, []byte("Go")))
```

Output:

```text
6
```

If not found:

```text
-1
```

---

## 4.16 `bytes.IndexByte`

```go
func IndexByte(b []byte, c byte) int
```

Finds the first occurrence of a single byte.

```go
data := []byte("hello")

fmt.Println(bytes.IndexByte(data, 'l'))
```

Output:

```text
2
```

---

## 4.17 `bytes.IndexAny`

```go
func IndexAny(s []byte, chars string) int
```

Returns the index of the first occurrence of any character in `chars`.

```go
data := []byte("hello123")

fmt.Println(bytes.IndexAny(data, "0123456789"))
```

Output:

```text
5
```

---

## 4.18 `bytes.IndexRune`

```go
func IndexRune(s []byte, r rune) int
```

Finds the first occurrence of a Unicode code point.

```go
data := []byte("Golang")

fmt.Println(bytes.IndexRune(data, 'a'))
```

---

## 4.19 `bytes.IndexFunc`

```go
func IndexFunc(s []byte, f func(rune) bool) int
```

Finds the first Unicode character for which a function returns `true`.

```go
data := []byte("abc123")

index := bytes.IndexFunc(data, func(r rune) bool {
	return r >= '0' && r <= '9'
})

fmt.Println(index)
```

Output:

```text
3
```

---

# G. Finding From the End

## 4.20 `bytes.LastIndex`

```go
func LastIndex(s, sep []byte) int
```

Finds the last occurrence of a byte sequence.

```go
data := []byte("go-go-go")

fmt.Println(bytes.LastIndex(data, []byte("go")))
```

---

## 4.21 `bytes.LastIndexByte`

```go
func LastIndexByte(s []byte, c byte) int
```

Finds the last occurrence of a byte.

```go
data := []byte("banana")

fmt.Println(bytes.LastIndexByte(data, 'a'))
```

---

## 4.22 `bytes.LastIndexAny`

```go
func LastIndexAny(s []byte, chars string) int
```

Finds the last occurrence of any character from `chars`.

---

## 4.23 `bytes.LastIndexFunc`

```go
func LastIndexFunc(s []byte, f func(rune) bool) int
```

Finds the last Unicode character satisfying a function.

These functions are useful for parsing things such as:

```text
/path/to/file.txt
```

where you might need the final `/` or `.`.

---

# H. Splitting Byte Data

## 4.24 `bytes.Split`

```go
func Split(s, sep []byte) [][]byte
```

Splits a byte slice around every occurrence of `sep`.

```go
data := []byte("Go,Python,Java")

parts := bytes.Split(data, []byte(","))

for _, part := range parts {
	fmt.Println(string(part))
}
```

Output:

```text
Go
Python
Java
```

---

## 4.25 `bytes.SplitN`

```go
func SplitN(s, sep []byte, n int) [][]byte
```

Splits into at most `n` pieces.

```go
data := []byte("a:b:c:d")

parts := bytes.SplitN(data, []byte(":"), 2)

for _, p := range parts {
	fmt.Println(string(p))
}
```

Output:

```text
a
b:c:d
```

---

## 4.26 `bytes.SplitAfter`

```go
func SplitAfter(s, sep []byte) [][]byte
```

Like `Split`, but keeps the separator at the end of each piece.

For:

```text
a,b,c
```

splitting by `,` gives conceptually:

```text
a,
b,
c
```

---

## 4.27 `bytes.SplitAfterN`

```go
func SplitAfterN(s, sep []byte, n int) [][]byte
```

Combines the behavior of `SplitAfter` with a maximum number of pieces.

---

## 4.28 `bytes.SplitSeq`

```go
func SplitSeq(s, sep []byte) iter.Seq[[]byte]
```

Returns an iterator sequence over pieces separated by `sep`.

Example:

```go
for part := range bytes.SplitSeq(data, []byte(",")) {
	fmt.Println(string(part))
}
```

Useful when you want to process pieces without first constructing a complete `[][]byte`.

---

## 4.29 `bytes.SplitAfterSeq`

```go
func SplitAfterSeq(s, sep []byte) iter.Seq[[]byte]
```

Iterator-based equivalent of `SplitAfter`.

The separator remains attached to each piece where applicable.

---

# I. Splitting on Whitespace

## 4.30 `bytes.Fields`

```go
func Fields(s []byte) [][]byte
```

Splits data around runs of Unicode whitespace.

```go
data := []byte("  Go   is   awesome  ")

words := bytes.Fields(data)

for _, word := range words {
	fmt.Println(string(word))
}
```

Output:

```text
Go
is
awesome
```

Leading and trailing whitespace is ignored.

---

## 4.31 `bytes.FieldsFunc`

```go
func FieldsFunc(s []byte, f func(rune) bool) [][]byte
```

Splits according to a custom function.

```go
data := []byte("Go,Python;Java")

parts := bytes.FieldsFunc(data, func(r rune) bool {
	return r == ',' || r == ';'
})

for _, part := range parts {
	fmt.Println(string(part))
}
```

Output:

```text
Go
Python
Java
```

---

## 4.32 `bytes.FieldsSeq`

```go
func FieldsSeq(s []byte) iter.Seq[[]byte]
```

Iterator-based version of `Fields`.

---

## 4.33 `bytes.FieldsFuncSeq`

```go
func FieldsFuncSeq(s []byte, f func(rune) bool) iter.Seq[[]byte]
```

Iterator-based version of `FieldsFunc`.

---

# J. Joining

## 4.34 `bytes.Join`

```go
func Join(s [][]byte, sep []byte) []byte
```

Joins multiple byte slices using a separator.

```go
parts := [][]byte{
	[]byte("Go"),
	[]byte("is"),
	[]byte("fun"),
}

result := bytes.Join(parts, []byte(" "))

fmt.Println(string(result))
```

Output:

```text
Go is fun
```

---

# K. Replacing Data

## 4.35 `bytes.Replace`

```go
func Replace(s, old, new []byte, n int) []byte
```

Replaces up to `n` occurrences.

```go
data := []byte("cat cat cat")

result := bytes.Replace(
	data,
	[]byte("cat"),
	[]byte("dog"),
	2,
)

fmt.Println(string(result))
```

Output:

```text
dog dog cat
```

---

## 4.36 `bytes.ReplaceAll`

```go
func ReplaceAll(s, old, new []byte) []byte
```

Replaces every occurrence.

```go
data := []byte("cat cat cat")

result := bytes.ReplaceAll(
	data,
	[]byte("cat"),
	[]byte("dog"),
)

fmt.Println(string(result))
```

Output:

```text
dog dog dog
```

---

# L. Trimming

## 4.37 `bytes.Trim`

```go
func Trim(s, cutset []byte) []byte
```

Removes characters from both ends that belong to `cutset`.

```go
data := []byte("...hello...")

result := bytes.Trim(data, []byte("."))

fmt.Println(string(result))
```

Output:

```text
hello
```

---

## 4.38 `bytes.TrimSpace`

```go
func TrimSpace(s []byte) []byte
```

Removes leading and trailing Unicode whitespace.

```go
data := []byte("   hello world   ")

fmt.Println(string(bytes.TrimSpace(data)))
```

Output:

```text
hello world
```

This is one of the most commonly used functions.

---

## 4.39 `bytes.TrimPrefix`

```go
func TrimPrefix(s, prefix []byte) []byte
```

Removes a prefix if present.

```go
data := []byte("Hello Go")

result := bytes.TrimPrefix(data, []byte("Hello "))

fmt.Println(string(result))
```

Output:

```text
Go
```

---

## 4.40 `bytes.TrimSuffix`

```go
func TrimSuffix(s, suffix []byte) []byte
```

Removes a suffix if present.

```go
data := []byte("hello.txt")

result := bytes.TrimSuffix(data, []byte(".txt"))

fmt.Println(string(result))
```

Output:

```text
hello
```

---

## 4.41 `bytes.TrimLeft`

```go
func TrimLeft(s, cutset []byte) []byte
```

Removes matching bytes from the left side.

---

## 4.42 `bytes.TrimRight`

```go
func TrimRight(s, cutset []byte) []byte
```

Removes matching bytes from the right side.

---

## 4.43 `bytes.TrimLeftFunc`

```go
func TrimLeftFunc(s []byte, f func(rune) bool) []byte
```

Removes characters from the beginning while the function returns `true`.

---

## 4.44 `bytes.TrimRightFunc`

```go
func TrimRightFunc(s []byte, f func(rune) bool) []byte
```

Removes characters from the end while the function returns `true`.

---

## 4.45 `bytes.TrimFunc`

```go
func TrimFunc(s []byte, f func(rune) bool) []byte
```

Trims from both sides according to a function.

---

# M. Case Conversion

## 4.46 `bytes.ToUpper`

```go
func ToUpper(s []byte) []byte
```

Converts Unicode letters to uppercase.

```go
fmt.Println(string(
	bytes.ToUpper([]byte("hello go")),
))
```

Output:

```text
HELLO GO
```

---

## 4.47 `bytes.ToLower`

```go
func ToLower(s []byte) []byte
```

Converts Unicode letters to lowercase.

---

## 4.48 `bytes.ToTitle`

```go
func ToTitle(s []byte) []byte
```

Converts letters to Unicode title case.

---

## 4.49 `bytes.ToUpperSpecial`

```go
func ToUpperSpecial(c unicode.SpecialCase, s []byte) []byte
```

Uppercases using a specific Unicode special-case mapping.

---

## 4.50 `bytes.ToLowerSpecial`

```go
func ToLowerSpecial(c unicode.SpecialCase, s []byte) []byte
```

Lowercases using a specified Unicode special-case mapping.

---

## 4.51 `bytes.ToTitleSpecial`

```go
func ToTitleSpecial(c unicode.SpecialCase, s []byte) []byte
```

Title-cases using a specified Unicode special-case mapping.

---

# N. UTF-8 Handling

## 4.52 `bytes.Runes`

```go
func Runes(s []byte) []rune
```

Converts UTF-8 encoded bytes into Unicode code points.

```go
data := []byte("Go ❤️")

runes := bytes.Runes(data)

fmt.Println(runes)
```

Useful when you need to work at the Unicode character level rather than individual bytes.

---

## 4.53 `bytes.ToValidUTF8`

```go
func ToValidUTF8(s, replacement []byte) []byte
```

Replaces invalid UTF-8 byte sequences with a replacement sequence.

```go
data := []byte{0xff, 0xfe, 'G', 'o'}

result := bytes.ToValidUTF8(
	data,
	[]byte("?"),
)

fmt.Println(string(result))
```

Useful when processing potentially malformed external data.

---

# O. Mapping Characters

## 4.54 `bytes.Map`

```go
func Map(mapping func(rune) rune, s []byte) []byte
```

Applies a function to each Unicode code point.

For example, you can remove digits:

```go
data := []byte("Go123")

result := bytes.Map(func(r rune) rune {
	if r >= '0' && r <= '9' {
		return -1
	}
	return r
}, data)

fmt.Println(string(result))
```

Output:

```text
Go
```

Returning `-1` removes the character.

---

# P. Cloning

## 4.55 `bytes.Clone`

```go
func Clone(b []byte) []byte
```

Creates an independent copy of a byte slice.

```go
original := []byte("hello")

copy := bytes.Clone(original)

copy[0] = 'H'

fmt.Println(string(original))
fmt.Println(string(copy))
```

Output:

```text
hello
Hello
```

This is important when you need to prevent two slices from sharing the same underlying data.

---

# Q. Repeating Data

## 4.56 `bytes.Repeat`

```go
func Repeat(b []byte, count int) []byte
```

Repeats a byte slice `count` times.

```go
result := bytes.Repeat([]byte("Go"), 3)

fmt.Println(string(result))
```

Output:

```text
GoGoGo
```

---

# R. `bytes.Buffer`

One of the most important parts of the package is:

```go
bytes.Buffer
```

A `Buffer` is a growable byte buffer that implements common `io.Reader` and `io.Writer` interfaces.

Example:

```go
var buffer bytes.Buffer

buffer.WriteString("Hello ")
buffer.WriteString("Golang")

fmt.Println(buffer.String())
```

Output:

```text
Hello Golang
```

---

## 4.57 `bytes.NewBuffer`

```go
func NewBuffer(buf []byte) *Buffer
```

Creates a buffer using an existing byte slice.

```go
data := []byte("Hello")

buffer := bytes.NewBuffer(data)

fmt.Println(buffer.String())
```

Be aware that the buffer can use the supplied slice as its underlying storage.

---

## 4.58 `bytes.NewBufferString`

```go
func NewBufferString(s string) *Buffer
```

Creates a buffer initialized with a string.

```go
buffer := bytes.NewBufferString("Hello Go")

fmt.Println(buffer.String())
```

---

# Buffer Methods

## `Available()`

```go
buffer.Available()
```

Returns the unused capacity available for writing.

---

## `AvailableBuffer()`

```go
buffer.AvailableBuffer()
```

Provides a byte slice representing available capacity, useful for efficient writes.

This is an advanced performance-oriented method.

---

## `Bytes()`

```go
buffer.Bytes()
```

Returns the unread portion of the buffer.

```go
buffer := bytes.NewBufferString("Hello")

data := buffer.Bytes()

fmt.Println(string(data))
```

Important: the returned slice refers to the buffer's underlying storage. Treat it accordingly.

---

## `Cap()`

```go
buffer.Cap()
```

Returns the buffer's capacity.

---

## `Grow()`

```go
buffer.Grow(n)
```

Ensures that the buffer can accommodate at least `n` additional bytes without another allocation.

```go
var b bytes.Buffer

b.Grow(1000)
```

Useful when you know approximately how much data you will write.

---

## `Len()`

```go
buffer.Len()
```

Returns the number of unread bytes.

```go
b := bytes.NewBufferString("hello")

fmt.Println(b.Len())
```

Output:

```text
5
```

---

## `Write()`

```go
buffer.Write(p []byte)
```

Writes bytes into the buffer.

```go
var b bytes.Buffer

b.Write([]byte("Hello"))

fmt.Println(b.String())
```

---

## `WriteString()`

```go
buffer.WriteString(s string)
```

Writes a string.

```go
var b bytes.Buffer

b.WriteString("Hello")
b.WriteString(" Go")

fmt.Println(b.String())
```

---

## `WriteByte()`

```go
buffer.WriteByte(c byte)
```

Writes one byte.

```go
var b bytes.Buffer

b.WriteByte('A')
b.WriteByte('B')

fmt.Println(b.String())
```

---

## `WriteRune()`

```go
buffer.WriteRune(r rune)
```

Writes a Unicode code point as UTF-8.

```go
var b bytes.Buffer

b.WriteRune('♥')

fmt.Println(b.String())
```

---

## `Read()`

```go
buffer.Read(p []byte)
```

Reads bytes from the buffer.

```go
b := bytes.NewBufferString("Hello")

data := make([]byte, 2)

n, err := b.Read(data)

fmt.Println(string(data[:n]))
fmt.Println(err)
```

Output begins with:

```text
He
```

---

## `ReadByte()`

Reads one byte.

```go
b := bytes.NewBufferString("Go")

c, _ := b.ReadByte()

fmt.Println(string(c))
```

---

## `ReadRune()`

Reads one Unicode code point.

```go
b := bytes.NewBufferString("Go")

r, _, _ := b.ReadRune()

fmt.Println(string(r))
```

---

## `ReadBytes()`

```go
buffer.ReadBytes(delim)
```

Reads until a specified delimiter.

```go
b := bytes.NewBufferString("hello,world")

data, _ := b.ReadBytes(',')

fmt.Println(string(data))
```

Output:

```text
hello,
```

---

## `ReadString()`

```go
buffer.ReadString(delim)
```

Same basic idea as `ReadBytes`, but returns a string.

```go
b := bytes.NewBufferString("hello,world")

result, _ := b.ReadString(',')

fmt.Println(result)
```

---

## `Next()`

```go
buffer.Next(n)
```

Returns the next `n` bytes and advances the buffer.

```go
b := bytes.NewBufferString("Hello")

data := b.Next(2)

fmt.Println(string(data))
```

Output:

```text
He
```

---

## `Peek()`

```go
buffer.Peek(n)
```

Looks at the next `n` bytes without advancing the buffer.

This is useful when parsing protocols.

---

## `Reset()`

```go
buffer.Reset()
```

Empties the buffer while allowing it to reuse its allocated storage.

```go
b := bytes.NewBufferString("Hello")

b.Reset()

fmt.Println(b.Len())
```

Output:

```text
0
```

---

## `String()`

```go
buffer.String()
```

Returns the unread contents as a string.

```go
b := bytes.NewBufferString("Hello Go")

fmt.Println(b.String())
```

---

## `Truncate()`

```go
buffer.Truncate(n)
```

Keeps only the first `n` unread bytes.

---

## `UnreadByte()`

```go
buffer.UnreadByte()
```

Causes the most recently read byte to become available again.

Useful when a parser reads one byte too far.

---

## `UnreadRune()`

```go
buffer.UnreadRune()
```

Unreads the most recently read Unicode rune.

It must generally be used immediately after `ReadRune`.

---

## `ReadFrom()`

```go
buffer.ReadFrom(r io.Reader)
```

Reads data from an `io.Reader` until EOF and appends it to the buffer.

Useful for:

```text
file → buffer
network connection → buffer
HTTP body → buffer
```

---

## `WriteTo()`

```go
buffer.WriteTo(w io.Writer)
```

Writes the buffer's unread data to an `io.Writer`.

Conceptually:

```text
Buffer → File
Buffer → Network connection
Buffer → HTTP response
```

---

# S. `bytes.Reader`

The second major type is:

```go
bytes.Reader
```

It lets you treat a byte slice as an `io.Reader`, `io.ReaderAt`, `io.Seeker`, etc.

Example:

```go
data := []byte("Hello Go")

reader := bytes.NewReader(data)

buffer := make([]byte, 5)

n, _ := reader.Read(buffer)

fmt.Println(string(buffer[:n]))
```

Output:

```text
Hello
```

---

## `bytes.NewReader`

```go
func NewReader(b []byte) *Reader
```

Creates a reader over a byte slice.

---

## `Reader.Len()`

```go
reader.Len()
```

Returns the number of unread bytes.

---

## `Reader.Size()`

```go
reader.Size()
```

Returns the original size of the underlying byte slice.

This differs from `Len()`.

```text
Size = original size
Len  = remaining size
```

---

## `Reader.Read()`

```go
reader.Read(p []byte)
```

Reads bytes from the reader.

---

## `Reader.ReadAt()`

```go
reader.ReadAt(p []byte, off int64)
```

Reads from a specific offset without changing the normal reader position.

Useful for random-access reading.

---

## `Reader.ReadByte()`

Reads one byte.

---

## `Reader.ReadRune()`

Reads one Unicode code point.

---

## `Reader.Seek()`

```go
reader.Seek(offset, whence)
```

Moves the reader position.

For example:

```go
reader.Seek(0, io.SeekStart)
```

moves back to the beginning.

---

## `Reader.Reset()`

```go
reader.Reset(b)
```

Reuses the reader with a different byte slice.

---

## `Reader.UnreadByte()`

Moves the most recently read byte back.

---

## `Reader.UnreadRune()`

Moves the most recently read rune back.

---

## `Reader.WriteTo()`

```go
reader.WriteTo(w io.Writer)
```

Copies the remaining reader contents to a writer.

---

# 5. Three common beginner mistakes

## Mistake 1: Thinking `[]byte` and `string` are identical

These are different types:

```go
var s string = "hello"
var b []byte = []byte("hello")
```

Conversion:

```go
b := []byte(s)
s := string(b)
```

Repeated conversions can have performance implications.

### Avoid it

Understand whether your API expects:

```text
string
```

or:

```text
[]byte
```

---

## Mistake 2: Forgetting that byte indexes are not necessarily character indexes

Consider:

```go
text := []byte("café")
```

The character `é` uses multiple bytes in UTF-8.

Therefore:

```go
len(text)
```

counts **bytes**, not human-visible characters.

For Unicode-aware processing, consider:

```go
bytes.Runes()
bytes.IndexRune()
bytes.Fields()
```

---

## Mistake 3: Assuming functions always modify the original slice

For example:

```go
data := []byte("hello")

result := bytes.ToUpper(data)
```

Use the returned value:

```go
data = bytes.ToUpper(data)
```

Similarly:

```go
result := bytes.ReplaceAll(...)
```

doesn't mean you can ignore the returned slice.

### Avoid it

Always check the function's return value and understand whether it returns a new slice, a subslice, or exposes underlying storage.

---

# 6. Two real-world applications

## Application 1: Network protocol parsing

Network data frequently arrives as:

```go
[]byte
```

For example:

```text
GET /users HTTP/1.1
Host: example.com
Content-Type: application/json
```

You can use:

```go
bytes.Index()
bytes.Contains()
bytes.Cut()
bytes.TrimSpace()
bytes.Split()
bytes.HasPrefix()
```

to efficiently inspect and parse the data.

This is useful in HTTP servers, TCP protocols, message brokers, and custom binary/text protocols.

---

## Application 2: Efficient file/data processing

Suppose you are processing a large file.

Instead of repeatedly converting data between:

```text
[]byte → string → []byte
```

you can operate directly on byte slices.

For example:

```go
data := make([]byte, 4096)

n, err := file.Read(data)

if err != nil {
	// handle error
}

data = bytes.TrimSpace(data[:n])
```

This approach is useful for:

- log processing
- CSV-like data
- configuration files
- binary files
- network payloads
- parsers
- serialization formats

---

# 7. `bytes.Buffer` vs `bytes.Reader`

This distinction is very important.

| Type | Main purpose |
|---|---|
| `bytes.Buffer` | Build/read/write growing byte data |
| `bytes.Reader` | Read existing byte data |
| `[]byte` | Store raw bytes |
| `string` | Store immutable text |

Think of them like this:

```text
[]byte
  │
  ├── bytes.Reader
  │       ↓
  │    READ data
  │
  └── bytes.Buffer
          ↓
       BUILD / READ / WRITE data
```

---

# 8. Three progressively challenging exercises

## Exercise 1 — Beginner: Byte Analyzer

Write a Go program that receives:

```text
"Go is fast and Go is simple"
```

as a `[]byte`.

Your program should:

1. Count how many times `"Go"` occurs.
2. Check whether the data contains `"fast"`.
3. Find the position of `"simple"`.
4. Convert the data to uppercase.
5. Trim leading and trailing whitespace.

**Do not use the `strings` package.**

---

## Exercise 2 — Intermediate: Log Parser

Create a program that processes this byte data:

```text
INFO:User logged in
ERROR:Database connection failed
WARN:Retrying connection
INFO:User logged out
ERROR:Timeout occurred
```

Using the `bytes` package:

1. Split the input into lines.
2. Determine the log level of each line.
3. Count `INFO`, `WARN`, and `ERROR` entries.
4. Extract the message after `:`.
5. Print only the error messages.

Try to solve the problem using byte operations rather than converting the entire input into a string first.

---

## Exercise 3 — Advanced: Mini Protocol Parser

Design a parser for messages in this format:

```text
METHOD=POST;PATH=/users;AUTH=Bearer123;BODY=name=Chandu
```

Your program should:

1. Parse each field.
2. Extract `METHOD`.
3. Extract `PATH`.
4. Extract `AUTH`.
5. Extract `BODY`.
6. Detect missing fields.
7. Handle fields appearing in different orders.
8. Handle malformed fields.
9. Avoid unnecessary string conversions.
10. Use `bytes.Buffer` where appropriate to construct the final normalized message.

Try to design the parser so that it could eventually process thousands of messages efficiently.

**Do not provide a solution; only implement the problem statement yourself.**

---

# 9. The deeper idea behind `bytes`

The important thing isn't memorizing 50+ functions.

The important question is:

> **At what level should I process my data?**

You have several levels:

```text
Unicode characters
       ↓
     string
       ↓
    []byte
       ↓
   individual bytes
       ↓
   binary representation
```

The `bytes` package gives you tools for operating around the `[]byte` level while still providing Unicode-aware operations where appropriate.

A useful learning strategy is to first learn:

```text
Equal
Compare
Contains
Index
Count
Cut
HasPrefix
HasSuffix
Split
Fields
Join
Replace
ReplaceAll
TrimSpace
ToUpper
ToLower
Clone
Buffer
Reader
```

Then learn the specialized functions.

The API also includes iterator-oriented functions such as `SplitSeq`, `SplitAfterSeq`, `FieldsSeq`, and `FieldsFuncSeq`, which are worth learning after understanding the traditional slice-returning functions.

---

# 10. Thought-provoking question

Imagine you're building a **high-performance Go HTTP server** that receives millions of requests per day.

Each request contains headers and a JSON body.

**Would you process everything as `string`, everything as `[]byte`, or use a combination of both?**

More importantly:

> **At what points would converting between `string` and `[]byte` be worth the cost, and when could avoiding that conversion make your program significantly more efficient?**

That question gets to the heart of why the `bytes` package exists.
