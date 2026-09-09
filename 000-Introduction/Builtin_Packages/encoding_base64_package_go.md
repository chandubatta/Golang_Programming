# `encoding/base64` Package in Go

The Go `encoding/base64` package provides **Base64 encoding and decoding**. It is commonly used when binary data needs to be represented as text—for example, in JSON, HTTP headers, URLs, configuration files, or email.

> **Important:** Base64 is **encoding, not encryption**. Anyone who has the encoded value can decode it.

---

## 1. What is `encoding/base64`?

The package is imported as:

```go
import "encoding/base64"
```

It provides implementations of several Base64 standards, including:

- **Standard Base64** — uses `A-Z`, `a-z`, `0-9`, `+`, `/`
- **URL-safe Base64** — uses `-`, `_` instead of `+`, `/`
- **Raw Base64** — omits `=` padding
- Streaming encoders/decoders for working with `io.Reader`/`io.Writer`

A quick example:

```text
Hello
   ↓ Base64 Encode
SGVsbG8=
   ↓ Base64 Decode
Hello
```

### When is it commonly used?

Typical situations include:

- Sending binary data through text-based protocols
- Encoding data inside JSON
- Representing small files/images as text
- Encoding values for HTTP headers
- Encoding data for URLs using URL-safe Base64
- Creating text representations of binary identifiers or tokens

---

# 2. Simple Example

```go
package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	message := "Hello, Go!"

	// Encode
	encoded := base64.StdEncoding.EncodeToString([]byte(message))
	fmt.Println("Encoded:", encoded)

	// Decode
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Decode error:", err)
		return
	}

	fmt.Println("Decoded:", string(decoded))
}
```

Output:

```text
Encoded: SGVsbG8sIEdvIQ==
Decoded: Hello, Go!
```

### What's happening?

```go
[]byte(message)
```

converts the string into bytes.

Then:

```go
base64.StdEncoding.EncodeToString(...)
```

converts those bytes into a Base64 string.

And:

```go
base64.StdEncoding.DecodeString(...)
```

converts the Base64 string back into bytes.

---

# 3. Important Types and Functions

The central type in this package is:

```go
type Encoding struct
```

You normally work with an `Encoding` value such as:

```go
base64.StdEncoding
```

or:

```go
base64.URLEncoding
```

---

# 4. `Encoding` Type — Detailed Functions

## `Encode`

```go
func (enc *Encoding) Encode(dst, src []byte)
```

Encodes the bytes in `src` into Base64 and writes the result into `dst`.

Example:

```go
src := []byte("Hello")
dst := make([]byte, base64.StdEncoding.EncodedLen(len(src)))

base64.StdEncoding.Encode(dst, src)

fmt.Println(string(dst))
```

Output:

```text
SGVsbG8=
```

### Important

You must allocate enough space for `dst`.

Use:

```go
base64.StdEncoding.EncodedLen(len(src))
```

to calculate the required size.

---

## `EncodeToString`

```go
func (enc *Encoding) EncodeToString(src []byte) string
```

Encodes bytes and directly returns a string.

Example:

```go
encoded := base64.StdEncoding.EncodeToString([]byte("Hello"))

fmt.Println(encoded)
```

Output:

```text
SGVsbG8=
```

This is usually the **simplest option** when you already have all the data in memory.

---

## `Decode`

```go
func (enc *Encoding) Decode(dst, src []byte) (n int, err error)
```

Decodes Base64 data from `src` and writes the decoded bytes into `dst`.

Example:

```go
encoded := []byte("SGVsbG8=")

dst := make([]byte, base64.StdEncoding.DecodedLen(len(encoded)))

n, err := base64.StdEncoding.Decode(dst, encoded)
if err != nil {
	fmt.Println("Error:", err)
	return
}

fmt.Println(string(dst[:n]))
```

Output:

```text
Hello
```

Notice:

```go
dst[:n]
```

Only the first `n` bytes contain the decoded data.

---

## `DecodeString`

```go
func (enc *Encoding) DecodeString(s string) ([]byte, error)
```

Decodes a Base64 string.

Example:

```go
decoded, err := base64.StdEncoding.DecodeString("SGVsbG8=")
if err != nil {
	fmt.Println("Error:", err)
	return
}

fmt.Println(string(decoded))
```

Output:

```text
Hello
```

This is generally the easiest decoding function to use.

---

## `EncodedLen`

```go
func (enc *Encoding) EncodedLen(n int) int
```

Returns the number of bytes required to encode `n` bytes.

Example:

```go
size := base64.StdEncoding.EncodedLen(5)

fmt.Println(size)
```

Output:

```text
8
```

Why?

```text
Hello
5 bytes
   ↓
SGVsbG8=
8 Base64 characters
```

This is especially useful when manually allocating destination buffers.

---

## `DecodedLen`

```go
func (enc *Encoding) DecodedLen(n int) int
```

Returns the maximum number of bytes that may be produced when decoding `n` Base64 characters.

Example:

```go
size := base64.StdEncoding.DecodedLen(8)

fmt.Println(size)
```

Output:

```text
6
```

For:

```text
SGVsbG8=
```

the maximum decoded size is calculated from its Base64 length.

---

## `Strict`

```go
func (enc *Encoding) Strict() *Encoding
```

Returns a new encoding that requires trailing padding bits to be zero during decoding.

This is useful when you want stricter validation of Base64 input.

Base64 represents data using groups of 6 bits. Some bits in the final encoded character can be padding bits. A strict encoding checks that those unused bits contain the expected zero values.

You normally don't need this for everyday Base64 usage, but it can be useful when **input validation matters**.

---

## `WithPadding`

```go
func (enc *Encoding) WithPadding(padding rune) *Encoding
```

Returns an encoding with a specified padding character.

Standard Base64 uses:

```text
=
```

as padding.

Example:

```go
encoding := base64.StdEncoding.WithPadding('*')

encoded := encoding.EncodeToString([]byte("Hello"))

fmt.Println(encoded)
```

You can also disable padding using:

```go
base64.NoPadding
```

For example:

```go
encoding := base64.StdEncoding.WithPadding(base64.NoPadding)

encoded := encoding.EncodeToString([]byte("Hello"))

fmt.Println(encoded)
```

Instead of:

```text
SGVsbG8=
```

you get:

```text
SGVsbG8
```

---

# 5. Streaming Functions

The package also provides functions for working with streams.

These are particularly useful when you're processing **large amounts of data** rather than one small string.

---

## `NewEncoder`

```go
func NewEncoder(enc *Encoding, w io.Writer) io.WriteCloser
```

Creates a streaming Base64 encoder.

Instead of:

```text
Entire input
     ↓
Base64 encode everything
     ↓
Entire output
```

you can encode data progressively.

Example:

```go
package main

import (
	"encoding/base64"
	"os"
)

func main() {
	encoder := base64.NewEncoder(base64.StdEncoding, os.Stdout)

	encoder.Write([]byte("Hello, Go!"))

	encoder.Close()
}
```

Output:

```text
SGVsbG8sIEdvIQ==
```

### Why `Close()` matters

Always close the encoder:

```go
encoder.Close()
```

The encoder may have buffered data that needs to be written as final output.

---

## `NewDecoder`

```go
func NewDecoder(enc *Encoding, r io.Reader) io.Reader
```

Creates a streaming Base64 decoder.

Example:

```go
package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"strings"
)

func main() {
	input := strings.NewReader("SGVsbG8sIEdvIQ==")

	decoder := base64.NewDecoder(base64.StdEncoding, input)

	data, err := io.ReadAll(decoder)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(data))
}
```

Output:

```text
Hello, Go!
```

The decoder reads Base64 data from the underlying `io.Reader` and provides decoded bytes through another `io.Reader`.

---

# 6. Predefined Encodings

The package provides several predefined encodings.

## `StdEncoding`

```go
base64.StdEncoding
```

This is the standard Base64 encoding defined by RFC 4648.

Alphabet:

```text
ABCDEFGHIJKLMNOPQRSTUVWXYZ
abcdefghijklmnopqrstuvwxyz
0123456789+/
```

It uses:

```text
=
```

for padding.

Example:

```go
encoded := base64.StdEncoding.EncodeToString([]byte("Hello"))
```

---

## `URLEncoding`

```go
base64.URLEncoding
```

This is the URL-safe Base64 variant.

The important difference is:

```text
Standard       URL-safe

+              -
/              _
```

For example, standard Base64 can produce:

```text
+
/
```

while URL-safe Base64 uses:

```text
-
_
```

This makes it better suited for URLs and similar contexts.

---

## `RawStdEncoding`

```go
base64.RawStdEncoding
```

This is standard Base64 **without padding**.

For example:

```go
base64.StdEncoding.EncodeToString([]byte("Hello"))
```

produces:

```text
SGVsbG8=
```

while:

```go
base64.RawStdEncoding.EncodeToString([]byte("Hello"))
```

produces:

```text
SGVsbG8
```

---

## `RawURLEncoding`

```go
base64.RawURLEncoding
```

This is URL-safe Base64 without padding.

It combines the characteristics of:

- `URLEncoding`
- no `=` padding

It is frequently useful for compact URL-safe representations.

---

# 7. Padding

One concept you should understand well is **padding**.

Base64 works with groups of 3 input bytes and converts them into 4 encoded characters.

For example:

```text
3 bytes → 4 Base64 characters
```

If the input isn't divisible by 3, padding may be added.

Example:

```text
Hello
```

becomes:

```text
SGVsbG8=
```

The `=` is padding.

With raw encoding:

```text
SGVsbG8
```

The padding is omitted.

---

# 8. Three Common Beginner Mistakes

## Mistake 1: Thinking Base64 is encryption

A common misconception is:

```text
password123
    ↓ Base64
cGFzc3dvcmQxMjM=
```

and assuming the password is now secure.

It isn't.

Anyone can decode it:

```go
decoded, _ := base64.StdEncoding.DecodeString("cGFzc3dvcmQxMjM=")
```

### Avoid it

Use encryption or password hashing when security is required.

Think:

```text
Base64 = encoding
Encryption = confidentiality
Hashing = one-way transformation
```

---

## Mistake 2: Using the wrong Base64 variant

You may encode using:

```go
base64.StdEncoding
```

and later try to decode using:

```go
base64.RawURLEncoding
```

Depending on the data, this can fail or produce unexpected behavior.

### Avoid it

Use matching encodings:

```text
StdEncoding → StdEncoding
RawStdEncoding → RawStdEncoding
URLEncoding → URLEncoding
RawURLEncoding → RawURLEncoding
```

---

## Mistake 3: Forgetting that Base64 increases size

Base64 isn't a compression algorithm.

It generally increases data size by about **33%** for sufficiently large inputs.

For example:

```text
3 bytes
   ↓
4 Base64 characters
```

### Avoid it

Don't Base64-encode data simply because you think it will make it smaller.

If your goal is compression, use something such as:

```text
compress/gzip
compress/zlib
```

and Base64-encode afterward only when you need a textual representation.

---

# 9. Real-World Application #1 — JSON APIs

Suppose an API needs to send binary data.

JSON is text-based, so binary data can be represented using Base64:

```json
{
    "filename": "photo.jpg",
    "data": "/9j/4AAQSkZJRgABAQ..."
}
```

In Go:

```go
encoded := base64.StdEncoding.EncodeToString(fileBytes)
```

The API can then transmit the Base64 string inside JSON.

The receiver decodes it:

```go
data, err := base64.StdEncoding.DecodeString(encoded)
```

This is useful for **small binary payloads**, although for large files, multipart uploads or object storage are generally preferable.

---

# 10. Real-World Application #2 — URL-Safe Tokens

Sometimes binary data needs to be represented inside a URL.

Standard Base64 contains:

```text
+
/
=
```

which can be inconvenient in some URL contexts.

URL-safe Base64 replaces:

```text
+ → -
/ → _
```

and you can omit padding with:

```go
base64.RawURLEncoding
```

Example:

```go
token := base64.RawURLEncoding.EncodeToString([]byte("user-123"))
```

This produces a compact URL-safe representation.

---

# 11. Three Practice Exercises

## Exercise 1 — Basic Encoding and Decoding

Write a Go program that:

1. Accepts a string from the user.
2. Encodes it using `base64.StdEncoding`.
3. Prints the encoded value.
4. Decodes the encoded value.
5. Prints the original string.

**Goal:** Practice `EncodeToString()` and `DecodeString()`.

---

## Exercise 2 — URL-Safe File Identifier

Create a program that:

1. Creates a byte slice containing some binary data.
2. Encodes it using URL-safe Base64.
3. Removes padding.
4. Prints the resulting value.
5. Decodes the value back to the original bytes.
6. Verifies that the decoded data matches the original.

**Goal:** Practice `RawURLEncoding` and understand the difference between standard, URL-safe, and padded Base64.

---

## Exercise 3 — Streaming Large Data

Create a program that:

1. Opens a large input file.
2. Uses `base64.NewEncoder()` to encode the file as a stream.
3. Writes the encoded data to another file.
4. Uses `base64.NewDecoder()` to decode the encoded file.
5. Writes the decoded data to a third file.
6. Verifies that the original and decoded files contain identical data.

**Goal:** Understand why streaming APIs are useful when the entire input shouldn't be loaded into memory at once.

---

# 12. Quick Function Reference

| Function / Value | Purpose |
|---|---|
| `StdEncoding` | Standard padded Base64 |
| `URLEncoding` | URL-safe padded Base64 |
| `RawStdEncoding` | Standard Base64 without padding |
| `RawURLEncoding` | URL-safe Base64 without padding |
| `NoPadding` | Special value for disabling padding |
| `Encode()` | Encode bytes into a destination buffer |
| `EncodeToString()` | Encode bytes directly to a string |
| `Decode()` | Decode Base64 into a destination buffer |
| `DecodeString()` | Decode a Base64 string |
| `EncodedLen()` | Calculate encoded size |
| `DecodedLen()` | Calculate maximum decoded size |
| `Strict()` | Create stricter decoding behavior |
| `WithPadding()` | Customize/disable padding |
| `NewEncoder()` | Create a streaming Base64 encoder |
| `NewDecoder()` | Create a streaming Base64 decoder |

---

# 13. Useful Mental Model

Think about the package in four layers:

```text
                    encoding/base64
                           │
             ┌─────────────┴─────────────┐
             │                           │
         ENCODING                    DECODING
             │                           │
    ┌────────┴────────┐          ┌───────┴────────┐
    │                 │          │                │
ToString           Buffer    FromString        Buffer
    │                 │          │                │
EncodeToString     Encode    DecodeString       Decode
```

And then there are the **streaming APIs**:

```text
io.Reader  ──→ NewDecoder() ──→ decoded bytes

bytes ──→ NewEncoder() ──→ io.Writer
```

The key distinction to remember is:

> **Base64 changes representation; it does not make data secret.**

---

# 14. Thought-Provoking Question

If Base64 makes binary data larger and provides **no confidentiality**, why do you think modern systems still use Base64 so extensively instead of simply sending the original binary bytes everywhere?

What problem is Base64 actually solving?
