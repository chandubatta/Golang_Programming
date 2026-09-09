# Go `encoding/hex` Package

## 1. What is `encoding/hex`?

The Go `encoding/hex` package is part of the standard library and provides functions for converting binary data (`[]byte`) to hexadecimal text and converting hexadecimal text back into binary data. It also provides streaming encoders/decoders and utilities for producing human-readable hex dumps.

### When is it commonly used?

It is particularly useful when you need to:

- Represent binary data as readable text.
- Display binary data while debugging.
- Work with hashes such as SHA-256.
- Represent byte sequences in logs or diagnostic output.
- Convert hexadecimal strings received from APIs into bytes.
- Inspect network packets or binary files.
- Encode/decode identifiers or binary protocol fields.
- Produce traditional hexadecimal dumps.

**Important:** Hex encoding is not encryption. Anyone can decode hexadecimal text back into the original bytes.

---

# 2. Basic Example

The simplest example is encoding bytes into hexadecimal and then decoding them back.

```go
package main

import (
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	original := []byte("Hello Go")

	// Encode bytes to hexadecimal.
	encoded := hex.EncodeToString(original)

	fmt.Println("Encoded:", encoded)

	// Decode hexadecimal back to bytes.
	decoded, err := hex.DecodeString(encoded)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Decoded:", string(decoded))
}
```

Output:

```text
Encoded: 48656c6c6f20476f
Decoded: Hello Go
```

The important relationship is:

```text
[]byte
   |
   | Encode
   v
hexadecimal string
   |
   | Decode
   v
[]byte
```

---

# 3. Every Function in `encoding/hex`

The package provides these functions:

1. `AppendDecode`
2. `AppendEncode`
3. `Decode`
4. `DecodeString`
5. `DecodedLen`
6. `Dump`
7. `Dumper`
8. `Encode`
9. `EncodeToString`
10. `EncodedLen`
11. `NewDecoder`
12. `NewEncoder`

It also provides the `InvalidByteError` type and its `Error()` method, plus the `ErrLength` variable.

---

## 3.1 `hex.Encode`

```go
func Encode(dst, src []byte) int
```

`Encode` converts bytes from `src` into hexadecimal and writes the result into the preallocated `dst` buffer. It returns the number of bytes written, which is always `len(src) * 2`.

### Example

```go
package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	src := []byte("Go")

	dst := make([]byte, hex.EncodedLen(len(src)))

	n := hex.Encode(dst, src)

	fmt.Println("Bytes written:", n)
	fmt.Println("Hex:", string(dst))
}
```

Output:

```text
Bytes written: 4
Hex: 476f
```

### Why use `Encode`?

Use it when you want control over the destination buffer and want to avoid creating an additional string.

---

## 3.2 `hex.EncodeToString`

```go
func EncodeToString(src []byte) string
```

This is the easiest way to convert bytes into a hexadecimal string.

### Example

```go
package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	data := []byte("Hello")

	hexString := hex.EncodeToString(data)

	fmt.Println(hexString)
}
```

Output:

```text
48656c6c6f
```

### `Encode` vs `EncodeToString`

Use:

```go
hex.Encode(dst, src)
```

when you're managing byte buffers.

Use:

```go
hex.EncodeToString(src)
```

when you simply want a string.

For beginners, start with `EncodeToString`.

---

## 3.3 `hex.Decode`

```go
func Decode(dst, src []byte) (int, error)
```

`Decode` converts hexadecimal characters in `src` back into bytes and writes those bytes into `dst`.

### Example

```go
package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	src := []byte("48656c6c6f")

	dst := make([]byte, hex.DecodedLen(len(src)))

	n, err := hex.Decode(dst, src)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(dst[:n]))
}
```

Output:

```text
Hello
```

### Important

The source must contain:

- valid hexadecimal characters
- an even number of characters

For example:

```text
48656c6c6f    valid
```

but:

```text
48656c6c6     invalid
```

has an odd number of characters.

`Decode` returns the number of bytes successfully decoded before an error occurs.

---

## 3.4 `hex.DecodeString`

```go
func DecodeString(s string) ([]byte, error)
```

`DecodeString` is the convenient string-based version of `Decode`. It takes a hexadecimal string and returns the corresponding bytes.

### Example

```go
package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	hexString := "48656c6c6f"

	data, err := hex.DecodeString(hexString)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(data))
}
```

Output:

```text
Hello
```

### Beginner-friendly pattern

You'll frequently see:

```go
data, err := hex.DecodeString(input)
if err != nil {
	// handle error
}
```

For normal application code, this is often easier than manually allocating a destination buffer.

---

## 3.5 `hex.EncodedLen`

```go
func EncodedLen(n int) int
```

This tells you how many bytes are required to hex-encode `n` bytes.

The calculation is:

```text
n × 2
```

because one byte becomes two hexadecimal characters.

### Example

```go
package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	n := 10

	fmt.Println(hex.EncodedLen(n))
}
```

Output:

```text
20
```

### Why is it useful?

When using `Encode`, you need to allocate the destination:

```go
dst := make([]byte, hex.EncodedLen(len(src)))
```

Without `EncodedLen`, you might incorrectly allocate:

```go
dst := make([]byte, len(src)) // WRONG
```

because hex requires twice as many bytes.

---

## 3.6 `hex.DecodedLen`

```go
func DecodedLen(x int) int
```

This tells you how many bytes you'll get when decoding `x` hexadecimal characters.

The calculation is:

```text
x / 2
```

### Example

```go
package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	fmt.Println(hex.DecodedLen(10))
}
```

Output:

```text
5
```

For example:

```text
48656c6c6f
```

contains 10 hexadecimal characters and decodes into:

```text
Hello
```

which is 5 bytes.

---

## 3.7 `hex.AppendEncode`

```go
func AppendEncode(dst, src []byte) []byte
```

`AppendEncode` hex-encodes `src` and appends the result to `dst`, returning the extended slice.

### Example

```go
package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	dst := []byte("ID:")

	dst = hex.AppendEncode(dst, []byte("Go"))

	fmt.Println(string(dst))
}
```

Output:

```text
ID:476f
```

Conceptually:

```text
dst = "ID:"
src = "Go"

AppendEncode
      ↓

"ID:" + "476f"
      ↓
"ID:476f"
```

### Why is it useful?

It is convenient when constructing a larger byte buffer:

```go
dst = hex.AppendEncode(dst, data1)
dst = hex.AppendEncode(dst, data2)
```

rather than repeatedly creating separate intermediate strings.

---

## 3.8 `hex.AppendDecode`

```go
func AppendDecode(dst, src []byte) ([]byte, error)
```

`AppendDecode` decodes hexadecimal data from `src` and appends the resulting bytes to `dst`.

### Example

```go
package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	dst := []byte("Result: ")

	dst, err := hex.AppendDecode(dst, []byte("48656c6c6f"))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(dst))
}
```

Output:

```text
Result: Hello
```

### Important difference

`AppendDecode` appends.

So:

```go
dst := []byte("ABC")
```

doesn't get replaced. Decoded data is added after it.

---

## 3.9 `hex.Dump`

```go
func Dump(data []byte) string
```

`Dump` creates a human-readable hexadecimal dump of the supplied bytes. Its format matches the traditional `hexdump -C` format.

### Example

```go
package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	data := []byte("Hello Go")

	fmt.Print(hex.Dump(data))
}
```

Output looks like:

```text
00000000  48 65 6c 6c 6f 20 47 6f  |Hello Go|
```

This is extremely useful for debugging binary data.

For example, if you're processing:

- network packets
- binary files
- encrypted data
- protocol messages

you can inspect the raw bytes in a readable form.

---

## 3.10 `hex.Dumper`

```go
func Dumper(w io.Writer) io.WriteCloser
```

`Dumper` is similar to `Dump`, but instead of receiving all the data at once, it returns a writer that continuously produces a hex dump of whatever you write to it.

### Example

```go
package main

import (
	"encoding/hex"
	"os"
)

func main() {
	dumper := hex.Dumper(os.Stdout)
	defer dumper.Close()

	dumper.Write([]byte("Hello "))
	dumper.Write([]byte("Go"))
}
```

The data is written incrementally and displayed as a hex dump.

### `Dump` vs `Dumper`

Use:

```go
hex.Dump(data)
```

when you already have the entire data.

Use:

```go
hex.Dumper(writer)
```

when data is arriving as a stream.

For example:

```text
File
 ↓
Reader
 ↓
Dumper
 ↓
Terminal
```

---

## 3.11 `hex.NewEncoder`

```go
func NewEncoder(w io.Writer) io.Writer
```

`NewEncoder` returns a writer that converts bytes written to it into lowercase hexadecimal characters and sends those characters to the underlying writer.

### Example

```go
package main

import (
	"encoding/hex"
	"os"
)

func main() {
	encoder := hex.NewEncoder(os.Stdout)

	encoder.Write([]byte("Hello"))
}
```

Output:

```text
48656c6c6f
```

### Think of it as a pipeline

```text
Your program
    |
    | []byte("Hello")
    v
NewEncoder
    |
    | "48656c6c6f"
    v
os.Stdout
```

This is particularly useful with `io.Reader`/`io.Writer` based programs.

---

## 3.12 `hex.NewDecoder`

```go
func NewDecoder(r io.Reader) io.Reader
```

`NewDecoder` returns a reader that reads hexadecimal characters from another reader and decodes them into bytes.

### Example

```go
package main

import (
	"encoding/hex"
	"fmt"
	"io"
	"strings"
)

func main() {
	source := strings.NewReader("48656c6c6f")

	decoder := hex.NewDecoder(source)

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
Hello
```

### Think of it as

```text
Hex source
    |
    | "48656c6c6f"
    v
NewDecoder
    |
    | []byte("Hello")
    v
Your program
```

This is especially useful when working with streams, files, network connections, or other `io.Reader`s.

---

## 3.13 `hex.InvalidByteError`

```go
type InvalidByteError byte
```

`InvalidByteError` represents an error caused by an invalid character in a hexadecimal string.

For example:

```go
hex.DecodeString("48GG")
```

contains `G`, which isn't a hexadecimal digit.

You can inspect the error:

```go
package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	_, err := hex.DecodeString("48GG")

	fmt.Printf("%T\n", err)
	fmt.Println(err)
}
```

The error identifies the invalid byte.

---

## 3.14 `InvalidByteError.Error()`

```go
func (e InvalidByteError) Error() string
```

This method converts the `InvalidByteError` into its human-readable error message.

You normally don't call it yourself.

Go automatically calls `Error()` when you print an error:

```go
fmt.Println(err)
```

---

## 3.15 `hex.ErrLength`

Although it isn't a function, you should know this exported variable:

```go
var ErrLength
```

It represents an attempt to decode an odd-length hexadecimal string using `Decode` or `DecodeString`.

Example:

```go
_, err := hex.DecodeString("ABC")
```

There are three hexadecimal characters:

```text
A B C
```

But hexadecimal encoding requires two characters per byte, so the input is invalid.

You can check for this error:

```go
if err == hex.ErrLength {
	fmt.Println("Hex string has an odd length")
}
```

---

# 4. The Most Important Functions to Remember

You don't need to memorize everything immediately.

Start with these four:

| Function | Purpose |
|---|---|
| `EncodeToString` | `[]byte → hex string` |
| `DecodeString` | `hex string → []byte` |
| `Encode` | `[]byte → []byte` using a destination buffer |
| `Decode` | hex `[]byte → []byte` using a destination buffer |

A useful mental model is:

```text
              ENCODE
[]byte -----------------> Hex text
  ^                          |
  |                          |
  |                          |
  +---------- DECODE --------+
```

For most beginner programs:

```go
encoded := hex.EncodeToString(data)

decoded, err := hex.DecodeString(encoded)
```

will be enough.

---

# 5. Common Mistakes Beginners Make

## Mistake 1: Thinking hex encoding is encryption

This is probably the most important misconception.

Someone might see:

```text
password123
```

become:

```text
70617373776f7264313233
```

and assume the password is protected.

It isn't.

Anyone can decode it:

```go
decoded, _ := hex.DecodeString("70617373776f7264313233")
```

Hex is an encoding, not encryption.

### Avoid it

Use cryptographic algorithms when confidentiality or security is required.

Use hex when you simply need a textual representation of bytes.

---

## Mistake 2: Forgetting that one byte becomes two hex characters

Suppose:

```go
data := []byte("Hello")
```

`Hello` contains 5 bytes.

Its hex representation contains:

```text
5 × 2 = 10
```

characters:

```text
48656c6c6f
```

Therefore:

```go
dst := make([]byte, len(data))
```

is not large enough for `Encode`.

Instead:

```go
dst := make([]byte, hex.EncodedLen(len(data)))
```

Use `EncodedLen` and `DecodedLen` when working with buffers.

---

## Mistake 3: Forgetting that hexadecimal strings must contain valid pairs

This is valid:

```text
48656c6c6f
```

This is invalid:

```text
48656c6c6
```

because the second one contains an odd number of characters.

This is also invalid:

```text
48656cGG6f
```

because `G` isn't a hexadecimal digit.

Valid hexadecimal characters are:

```text
0 1 2 3 4 5 6 7 8 9
a b c d e f
A B C D E F
```

Always check the error returned by decoding functions.

---

# 6. Real-World Application #1: Hash Representation

One of the most common uses of hex is displaying a cryptographic hash.

For example, SHA-256 produces 32 bytes of binary data.

Those 32 bytes can be represented as 64 hexadecimal characters.

Conceptually:

```text
SHA-256
   |
   v
32 binary bytes
   |
   | hex encoding
   v
64-character hexadecimal string
```

In Go:

```go
hashBytes := sha256.Sum256([]byte("Hello"))

hexHash := hex.EncodeToString(hashBytes[:])

fmt.Println(hexHash)
```

This is why you frequently see hashes displayed like:

```text
a591a6d40bf420404a011733cfb7b190...
```

The hexadecimal string is simply a convenient textual representation of the binary hash.

---

# 7. Real-World Application #2: Debugging Binary Data

Suppose your application receives a binary network packet:

```go
packet := []byte{
	0x48,
	0x65,
	0x6c,
	0x6c,
	0x6f,
}
```

Printing it normally isn't very useful:

```go
fmt.Println(packet)
```

Instead:

```go
fmt.Print(hex.Dump(packet))
```

gives you a structured representation showing:

- byte offsets
- hexadecimal values
- ASCII representation

This is extremely useful when debugging:

- network protocols
- binary file formats
- device communication
- cryptographic protocols
- serialization formats
- low-level systems programming

---

# 8. Three Progressively Challenging Exercises

## Exercise 1 — Beginner: Hex Converter

Write a Go program that:

1. Reads a string from the user.
2. Converts it into `[]byte`.
3. Encodes the bytes using `hex.EncodeToString`.
4. Prints the hexadecimal representation.
5. Decodes the hexadecimal string back into bytes.
6. Prints the original string again.

Example interaction:

```text
Enter text: Hello Go

Hex: 48656c6c6f20476f

Decoded: Hello Go
```

**Do not use `hex.Encode` or `hex.Decode` for this exercise.**

---

## Exercise 2 — Intermediate: Hexadecimal File Processor

Create a program that:

1. Reads binary data from a file.
2. Converts the data to hexadecimal.
3. Writes the hexadecimal representation to another file.
4. Reads the hexadecimal file.
5. Decodes it back into the original binary data.
6. Writes the decoded data into a third file.
7. Verifies that the original and final files contain identical bytes.

Your program should correctly handle decoding errors.

Try using:

```text
Encode
Decode
EncodedLen
DecodedLen
```

rather than relying entirely on `EncodeToString` and `DecodeString`.

---

## Exercise 3 — Advanced: Streaming Hexadecimal Processor

Build a program that processes a large file without loading the entire file into memory.

Your program should:

1. Open a potentially large binary file.
2. Read it as a stream.
3. Convert the binary stream to hexadecimal.
4. Write the hexadecimal data to another stream/file.
5. Read the hexadecimal data back as a stream.
6. Decode it back into binary.
7. Write the decoded binary data to a new file.
8. Verify that the original and reconstructed files are identical.

Use the streaming APIs:

```text
hex.NewEncoder
hex.NewDecoder
io.Reader
io.Writer
```

The goal is to understand why streaming APIs can be preferable to loading a huge file into memory at once.

---

# 9. Useful Function Map

Keep this cheat sheet nearby while practicing:

```text
ENCODING
────────────────────────────────────

EncodeToString(src)
        ↓
[]byte → string

Encode(dst, src)
        ↓
[]byte → []byte

AppendEncode(dst, src)
        ↓
append encoded data to dst


DECODING
────────────────────────────────────

DecodeString(s)
        ↓
string → []byte

Decode(dst, src)
        ↓
hex []byte → []byte

AppendDecode(dst, src)
        ↓
append decoded data to dst


SIZE CALCULATIONS
────────────────────────────────────

EncodedLen(n)
        ↓
n bytes → n × 2 hex bytes

DecodedLen(n)
        ↓
n hex bytes → n / 2 decoded bytes


STREAMING
────────────────────────────────────

NewEncoder(writer)
        ↓
binary stream → hex stream

NewDecoder(reader)
        ↓
hex stream → binary stream


DEBUGGING
────────────────────────────────────

Dump(data)
        ↓
[]byte → formatted hex dump

Dumper(writer)
        ↓
stream → formatted hex dump


ERRORS
────────────────────────────────────

ErrLength
        ↓
odd-length hexadecimal input

InvalidByteError
        ↓
invalid hexadecimal character
```

---

# 10. Thought-Provoking Question

Imagine you're designing a system that stores SHA-256 hashes for **1 billion records**.

A SHA-256 hash is 32 bytes, but its hexadecimal representation requires 64 characters and therefore substantially more storage.

**If hex is so convenient for humans, when would you choose hex over storing the raw 32-byte hash—and when would you deliberately avoid hex? What trade-offs involving storage, database indexing, network bandwidth, debugging, interoperability, and human readability would influence your decision?**
