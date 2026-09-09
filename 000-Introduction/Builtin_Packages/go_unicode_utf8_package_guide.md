# Go `unicode/utf8` Package

## 1. What is the `unicode/utf8` package?

Go's `unicode/utf8` package provides functions for working with **UTF-8 encoded Unicode text**.

UTF-8 is a variable-width encoding:

- ASCII characters use **1 byte**
- Many Latin characters use **2 bytes**
- Many other Unicode characters use **3 bytes**
- Some characters, including many emoji, use **4 bytes**

For example:

```text
A       → 1 byte
é       → 2 bytes
中      → 3 bytes
😀      → 4 bytes
```

Go's `string` type stores **bytes**, not Unicode characters. Therefore, when you need to inspect, validate, encode, decode, or count Unicode characters correctly, `unicode/utf8` is very useful.

Import it with:

```go
import "unicode/utf8"
```

### When is it commonly used?

You will commonly use `unicode/utf8` when:

- validating whether a string contains valid UTF-8
- counting Unicode characters
- decoding UTF-8 byte sequences
- encoding Unicode code points (`rune`) into UTF-8
- processing multilingual text
- working with text files or network protocols
- safely handling user-generated Unicode text

---

# 2. Important concepts before using `unicode/utf8`

Understanding these three concepts makes the package much easier.

## Byte

A byte is 8 bits.

```go
var b byte = 65
```

A Go string is essentially a sequence of bytes.

```go
s := "Hello"
fmt.Println(len(s))
```

Output:

```text
5
```

For ASCII, the number of bytes happens to equal the number of characters.

But consider:

```go
s := "世界"
fmt.Println(len(s))
```

The result is:

```text
6
```

There are only two Unicode characters, but each uses three bytes.

---

## Rune

A `rune` in Go is an alias for `int32`.

It represents a Unicode code point.

```go
var r rune = '世'
```

You can think of:

```text
byte → raw UTF-8 data
rune → Unicode code point
```

---

## UTF-8

UTF-8 converts Unicode code points into sequences of bytes.

For example:

```text
A    → 41
é    → C3 A9
世   → E4 B8 96
😀   → F0 9F 98 80
```

This is why:

```go
len("😀")
```

returns:

```text
4
```

while:

```go
utf8.RuneCountInString("😀")
```

returns:

```text
1
```

---

# 3. Functions and constants in `unicode/utf8`

The package provides several constants and functions for UTF-8 processing.

## Constants

| Constant | Meaning |
|---|---|
| `RuneError` | Replacement character used for invalid UTF-8 or invalid Unicode code points |
| `RuneSelf` | Maximum rune value that can be represented using one byte |
| `MaxRune` | Maximum valid Unicode code point |
| `UTFMax` | Maximum number of bytes required to encode a UTF-8 rune |

---

## `utf8.RuneError`

`RuneError` represents the Unicode replacement character:

```text
U+FFFD
```

You can use it like:

```go
fmt.Printf("%U\n", utf8.RuneError)
```

Output:

```text
U+FFFD
```

It is commonly used when invalid UTF-8 data is encountered.

For example:

```go
data := []byte{0xff}

r, size := utf8.DecodeRune(data)

fmt.Printf("Rune: %U\n", r)
fmt.Println("Size:", size)
```

The invalid byte is represented using `RuneError`.

### Important

`RuneError` does **not always mean that the original text contained the character `�`**.

It can indicate that invalid UTF-8 was encountered during decoding.

---

# `utf8.RuneSelf`

`RuneSelf` is:

```text
128
```

It represents the largest Unicode code point that can be encoded using a single UTF-8 byte.

In other words:

```go
utf8.RuneSelf
```

is equivalent to:

```text
0x80
```

ASCII characters range from:

```text
U+0000 → U+007F
```

Therefore, every ASCII character uses exactly one byte in UTF-8.

For example:

```go
fmt.Println(utf8.RuneSelf)
```

Output:

```text
128
```

This constant is particularly useful when implementing low-level UTF-8 processing.

---

# `utf8.MaxRune`

`MaxRune` represents the largest valid Unicode code point:

```text
U+10FFFF
```

You can check it with:

```go
fmt.Printf("%U\n", utf8.MaxRune)
```

Output:

```text
U+10FFFF
```

It is useful when checking whether a value can represent a valid Unicode code point.

---

# `utf8.UTFMax`

`UTFMax` is the maximum number of bytes needed to encode one UTF-8 encoded Unicode code point.

Its value is:

```text
4
```

For example:

```text
A    → 1 byte
é    → 2 bytes
世   → 3 bytes
😀   → 4 bytes
```

So:

```go
fmt.Println(utf8.UTFMax)
```

prints:

```text
4
```

---

# 4. `utf8.DecodeRune`

```go
utf8.DecodeRune(p []byte) (r rune, size int)
```

`DecodeRune` examines the beginning of a byte slice and decodes the first UTF-8 encoded rune.

It returns two values:

```text
rune
size
```

For example:

```go
data := []byte("Hello")

r, size := utf8.DecodeRune(data)

fmt.Printf("Rune: %c\n", r)
fmt.Println("Bytes:", size)
```

Output:

```text
Rune: H
Bytes: 1
```

With Unicode:

```go
data := []byte("😀")

r, size := utf8.DecodeRune(data)

fmt.Printf("Rune: %c\n", r)
fmt.Println("Bytes:", size)
```

Output:

```text
Rune: 😀
Bytes: 4
```

### Why is the size important?

UTF-8 is variable-width.

The function tells you exactly how many bytes belong to the decoded rune.

You can therefore manually process a UTF-8 byte slice:

```text
[byte byte byte byte] → rune + number of bytes
```

### Invalid input

If the beginning of `p` does not contain valid UTF-8, `DecodeRune` returns `RuneError` and consumes one byte.

---

# 5. `utf8.DecodeRuneInString`

```go
utf8.DecodeRuneInString(s string) (rune, size int)
```

This is similar to `DecodeRune`, but works directly on a string.

Example:

```go
s := "世界"

r, size := utf8.DecodeRuneInString(s)

fmt.Printf("Rune: %c\n", r)
fmt.Println("Bytes:", size)
```

Output:

```text
Rune: 世
Bytes: 3
```

The difference is simply the input type:

```text
DecodeRune          → []byte
DecodeRuneInString  → string
```

This function is useful when you're processing strings without first converting them to `[]byte`.

---

# 6. `utf8.EncodeRune`

```go
utf8.EncodeRune(p []byte, r rune) int
```

`EncodeRune` converts a Unicode code point into UTF-8 bytes. The destination slice must have enough room.

Example:

```go
buf := make([]byte, utf8.UTFMax)

n := utf8.EncodeRune(buf, '😀')

fmt.Println(n)
fmt.Println(buf[:n])
```

For the emoji `😀`, `n` will be:

```text
4
```

So four bytes were written.

### Why use this?

It is useful when working at the byte level and you need to manually construct UTF-8 data.

---

# 7. `utf8.RuneLen`

```go
utf8.RuneLen(r rune) int
```

`RuneLen` returns the number of bytes needed to encode a rune as UTF-8.

Example:

```go
fmt.Println(utf8.RuneLen('A'))
fmt.Println(utf8.RuneLen('é'))
fmt.Println(utf8.RuneLen('世'))
fmt.Println(utf8.RuneLen('😀'))
```

Conceptually:

```text
A    → 1
é    → 2
世   → 3
😀   → 4
```

If the rune is not a valid Unicode code point, the result is negative.

### Useful mental model

```text
RuneLen = "How many UTF-8 bytes would this rune require?"
```

---

# 8. `utf8.RuneStart`

```go
utf8.RuneStart(b byte) bool
```

This function determines whether a byte can be the **first byte of a UTF-8 encoded rune**.

Example:

```go
data := []byte("😀")

for _, b := range data {
    fmt.Printf("%02X %v\n", b, utf8.RuneStart(b))
}
```

For a four-byte UTF-8 sequence, the first byte is a rune start, while continuation bytes are not.

### Why is this useful?

It is useful for low-level UTF-8 parsing.

UTF-8 has a structure where continuation bytes have the form:

```text
10xxxxxx
```

while the initial byte has a different pattern.

You generally don't need `RuneStart` for ordinary text processing, but it can be valuable when implementing parsers or scanners.

---

# 9. `utf8.RuneCount`

```go
utf8.RuneCount(p []byte) int
```

`RuneCount` counts the number of UTF-8 encoded runes in a byte slice.

Example:

```go
data := []byte("Hello 世界 😀")

fmt.Println(utf8.RuneCount(data))
```

This counts Unicode code points rather than bytes.

Compare:

```go
fmt.Println(len(data))
fmt.Println(utf8.RuneCount(data))
```

These numbers can be different.

### Important

`RuneCount` counts runes/code points, **not necessarily what a human perceives as characters**.

For example, some visible characters can consist of multiple Unicode code points.

---

# 10. `utf8.RuneCountInString`

```go
utf8.RuneCountInString(s string) int
```

This is the string version of `RuneCount`.

Example:

```go
s := "Hello 世界 😀"

fmt.Println(utf8.RuneCountInString(s))
```

This is often more convenient than converting the string to `[]byte`.

Compare:

```go
utf8.RuneCount([]byte(s))
```

with:

```go
utf8.RuneCountInString(s)
```

The second version directly accepts the string.

---

# 11. `utf8.Valid`

```go
utf8.Valid(p []byte) bool
```

`Valid` checks whether an entire byte slice contains valid UTF-8.

Example:

```go
data := []byte("Hello 世界")

if utf8.Valid(data) {
    fmt.Println("Valid UTF-8")
} else {
    fmt.Println("Invalid UTF-8")
}
```

Output:

```text
Valid UTF-8
```

You can also construct invalid UTF-8:

```go
data := []byte{0xff, 0xfe}

fmt.Println(utf8.Valid(data))
```

This returns:

```text
false
```

### Real-world importance

This is particularly useful when receiving raw data from:

- files
- network connections
- external APIs
- databases
- user input

before assuming the bytes represent valid UTF-8 text.

---

# 12. `utf8.ValidString`

```go
utf8.ValidString(s string) bool
```

This checks whether a string contains valid UTF-8.

Example:

```go
s := "Hello 世界 😀"

fmt.Println(utf8.ValidString(s))
```

Output:

```text
true
```

You might wonder:

> "Why check a Go string for UTF-8 validity?"

Because a Go string can contain **arbitrary bytes**. It is not required to contain valid UTF-8.

For example:

```go
s := string([]byte{0xff, 0xfe})

fmt.Println(utf8.ValidString(s))
```

This can be:

```text
false
```

This distinction is important when processing external or binary-derived data.

---

# 13. `utf8.ValidRune`

```go
utf8.ValidRune(r rune) bool
```

This checks whether a rune represents a valid Unicode code point.

Example:

```go
r := '😀'

fmt.Println(utf8.ValidRune(r))
```

Output:

```text
true
```

You can also test an invalid value:

```go
r := rune(0x110000)

fmt.Println(utf8.ValidRune(r))
```

Since Unicode currently ends at:

```text
U+10FFFF
```

that value is invalid.

### Important distinction

`ValidRune` checks a **rune/code point**.

`Valid` checks a **byte sequence**.

`ValidString` checks the bytes contained in a **string**.

---

# 14. `utf8.FullRune`

```go
utf8.FullRune(p []byte) bool
```

`FullRune` determines whether the beginning of a byte slice contains a complete UTF-8 encoded rune.

This is particularly useful when data arrives in **chunks**.

Imagine receiving a UTF-8 character over a network:

```text
Chunk 1 → first bytes
Chunk 2 → remaining bytes
```

The first chunk might not contain the complete rune.

`FullRune` helps determine whether you have enough bytes to decode it.

Example concept:

```go
data := []byte{0xE4, 0xB8}

fmt.Println(utf8.FullRune(data))
```

The sequence is incomplete, so it reports that a full rune isn't available.

If the remaining byte is added:

```go
data := []byte{0xE4, 0xB8, 0x96}

fmt.Println(utf8.FullRune(data))
```

the sequence represents a complete rune.

### Real-world use

This is particularly relevant to:

- network protocols
- streaming parsers
- buffered readers
- incremental text processing

---

# 15. `utf8.FullRuneInString`

```go
utf8.FullRuneInString(s string) bool
```

This is the string equivalent of `FullRune`.

It determines whether the beginning of the string contains a complete UTF-8 encoded rune.

Example:

```go
s := "世"

fmt.Println(utf8.FullRuneInString(s))
```

This returns `true`.

The function becomes particularly useful when you're processing potentially incomplete UTF-8 strings.

---

# 16. `utf8.DecodeLastRune`

```go
utf8.DecodeLastRune(p []byte) (rune, size int)
```

This decodes the **last UTF-8 encoded rune** in a byte slice.

Example:

```go
data := []byte("Hello😀")

r, size := utf8.DecodeLastRune(data)

fmt.Printf("Rune: %c\n", r)
fmt.Println("Bytes:", size)
```

Output conceptually:

```text
Rune: 😀
Bytes: 4
```

This is useful when processing text backwards.

---

# 17. `utf8.DecodeLastRuneInString`

```go
utf8.DecodeLastRuneInString(s string) (rune, size int)
```

This is the string version of `DecodeLastRune`.

Example:

```go
s := "Hello😀"

r, size := utf8.DecodeLastRuneInString(s)

fmt.Printf("Rune: %c\n", r)
fmt.Println("Bytes:", size)
```

It returns the final Unicode code point and the number of bytes used by it.

---

# 18. Putting the important functions together

Here's a useful example that demonstrates several functions:

```go
package main

import (
    "fmt"
    "unicode/utf8"
)

func main() {
    text := "Hello, 世界! 😀"

    fmt.Println("Text:", text)

    // Number of bytes
    fmt.Println("Bytes:", len(text))

    // Number of Unicode code points
    fmt.Println("Runes:", utf8.RuneCountInString(text))

    // Check UTF-8 validity
    fmt.Println("Valid UTF-8:", utf8.ValidString(text))

    // Decode first rune
    r, size := utf8.DecodeRuneInString(text)

    fmt.Printf("First rune: %c\n", r)
    fmt.Println("First rune size:", size)

    // Decode last rune
    r, size = utf8.DecodeLastRuneInString(text)

    fmt.Printf("Last rune: %c\n", r)
    fmt.Println("Last rune size:", size)
}
```

The key lesson is:

```text
len(string)
    ↓
number of bytes

utf8.RuneCountInString(string)
    ↓
number of Unicode code points
```

---

# 19. `unicode/utf8` vs `unicode`

Don't confuse:

```go
unicode/utf8
```

with:

```go
unicode
```

They solve different problems.

## `unicode/utf8`

Deals primarily with **UTF-8 encoding and decoding**.

For example:

```go
utf8.DecodeRuneInString(...)
utf8.ValidString(...)
utf8.RuneCountInString(...)
```

## `unicode`

Deals primarily with **Unicode character properties**.

For example:

```go
unicode.IsLetter(...)
unicode.IsDigit(...)
unicode.IsUpper(...)
unicode.IsLower(...)
```

A real application may use both packages together.

---

# 20. `unicode/utf8` vs `len`

One of the most important beginner lessons is this:

```go
len(s)
```

does **not** return the number of Unicode characters.

It returns the number of bytes in the string.

For example:

```go
s := "😀"

fmt.Println(len(s))
```

Output:

```text
4
```

But:

```go
fmt.Println(utf8.RuneCountInString(s))
```

outputs:

```text
1
```

Therefore:

```text
len(s)
      → bytes

RuneCountInString(s)
      → Unicode code points
```

---

# 21. `unicode/utf8` and `range`

Go's `range` over a string automatically decodes UTF-8.

For example:

```go
package main

import "fmt"

func main() {
    s := "Hello 世界 😀"

    for i, r := range s {
        fmt.Printf("index=%d rune=%c\n", i, r)
    }
}
```

Notice that `i` is a **byte index**, not a character number.

This is another important concept:

```go
for i, r := range s
```

means:

```text
i → byte position
r → decoded rune
```

The `unicode/utf8` package becomes particularly useful when you need explicit control over this decoding process.

---

# 22. Common mistakes and misconceptions

## Mistake 1: Assuming `len()` counts characters

Beginners often write:

```go
if len(username) > 20 {
    // ...
}
```

and assume this means "20 characters."

It actually means **20 bytes**.

For Unicode text, these are not necessarily the same.

### Avoid it

If your requirement is specifically to count Unicode code points:

```go
utf8.RuneCountInString(username)
```

However, remember that even rune count isn't always identical to the number of user-perceived characters.

---

## Mistake 2: Assuming every rune is one byte

A beginner might think:

```text
one rune = one byte
```

That's only true for ASCII.

UTF-8 uses:

```text
1–4 bytes per Unicode code point
```

For example:

```text
A    → 1 byte
é    → 2 bytes
世   → 3 bytes
😀   → 4 bytes
```

### Avoid it

Use:

```go
utf8.RuneLen(r)
```

when you need to know how many bytes a particular rune requires.

---

## Mistake 3: Confusing Unicode code points with visible characters

This is a more advanced misconception.

A "character" from a user's perspective isn't always one Unicode code point.

For example, a displayed character can sometimes be constructed from:

```text
base character + combining mark
```

Therefore:

```go
utf8.RuneCountInString(...)
```

counts **code points**, not necessarily grapheme clusters (user-perceived characters).

### Avoid it

Understand what your application actually needs:

```text
bytes?
    ↓
len()

Unicode code points?
    ↓
utf8.RuneCountInString()

user-perceived characters/grapheme clusters?
    ↓
use Unicode grapheme-cluster-aware processing
```

---

# 23. Two real-world applications

## Application 1: Validating user input

Suppose a web service receives text from an external client.

Before processing raw bytes as UTF-8:

```go
if !utf8.Valid(data) {
    return fmt.Errorf("invalid UTF-8")
}
```

This can prevent invalid text from entering systems that expect valid UTF-8.

This is useful for:

- APIs
- web servers
- messaging systems
- database ingestion
- file processing

---

## Application 2: Processing multilingual text

Imagine building a text-processing application supporting:

```text
English
Hindi
Japanese
Chinese
Arabic
Emoji
```

You can't safely assume that:

```go
len(text)
```

equals the number of characters.

UTF-8-aware processing allows your program to correctly decode and count Unicode code points.

For example:

```go
count := utf8.RuneCountInString(text)
```

This is useful for:

- text editors
- messaging applications
- search systems
- text analysis
- input validation
- internationalized applications

---

# 24. Three progressively challenging exercises

## Exercise 1 — UTF-8 Inspector

Write a Go program that accepts a string such as:

```text
Hello, 世界! 😀
```

and prints:

1. The original string
2. The number of bytes
3. The number of Unicode code points
4. Whether the string contains valid UTF-8
5. The first rune
6. The last rune

Use functions from `unicode/utf8`.

**Do not use a library that directly provides character-counting functionality.**

---

## Exercise 2 — Manual UTF-8 Decoder

Create a program that receives a UTF-8 string and manually walks through its bytes.

For every Unicode code point, print:

```text
Byte position
Rune
Unicode code point
Number of bytes used
```

For example, the conceptual output should look like:

```text
Position: ...
Rune: ...
Code Point: ...
Bytes: ...
```

Use `DecodeRune` or `DecodeRuneInString` rather than using a `range` loop.

Your program should work correctly with:

- ASCII
- accented characters
- Asian characters
- emoji

---

## Exercise 3 — Streaming UTF-8 Processor

Create a program that simulates receiving UTF-8 text in small byte chunks rather than receiving the complete string at once.

Your program should:

1. Receive arbitrary byte chunks.
2. Determine whether enough bytes are available for a complete UTF-8 rune.
3. Handle a UTF-8 rune split across two or more chunks.
4. Detect invalid UTF-8 sequences.
5. Decode each complete rune.
6. Keep incomplete bytes until the next chunk arrives.
7. Report the total number of successfully decoded Unicode code points.

Use functions such as:

```text
FullRune
DecodeRune
Valid
```

The program should correctly handle input containing ASCII, multilingual text, and emoji.

---

# 25. Quick reference

| Function/constant | Purpose |
|---|---|
| `RuneError` | Unicode replacement/error rune |
| `RuneSelf` | Maximum one-byte UTF-8 rune value |
| `MaxRune` | Maximum valid Unicode code point |
| `UTFMax` | Maximum UTF-8 bytes per rune |
| `DecodeRune` | Decode first rune from `[]byte` |
| `DecodeRuneInString` | Decode first rune from `string` |
| `DecodeLastRune` | Decode last rune from `[]byte` |
| `DecodeLastRuneInString` | Decode last rune from `string` |
| `EncodeRune` | Encode a rune into UTF-8 bytes |
| `RuneLen` | Number of bytes needed for a rune |
| `RuneStart` | Tests whether a byte can start a UTF-8 rune |
| `RuneCount` | Count runes in `[]byte` |
| `RuneCountInString` | Count runes in a string |
| `Valid` | Check whether bytes are valid UTF-8 |
| `ValidString` | Check whether a string contains valid UTF-8 |
| `ValidRune` | Check whether a rune is a valid Unicode code point |
| `FullRune` | Check whether bytes contain a complete rune |
| `FullRuneInString` | Check whether a string starts with a complete rune |

---

# 26. Thought-provoking question

Suppose you're building a messaging application that limits messages to **100 characters**.

A user sends:

```text
😀😀😀😀😀
```

Another user sends:

```text
aaaaaaaaaa
```

And another sends a sequence containing combining Unicode characters.

**Should your application enforce the limit using bytes, Unicode code points, or user-perceived characters—and what problems could occur if you choose the wrong one?**

Think carefully about the difference between **storage representation, Unicode code points, and what the user actually sees on screen**.
