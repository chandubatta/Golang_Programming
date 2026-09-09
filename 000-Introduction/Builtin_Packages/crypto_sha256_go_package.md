# Go `crypto/sha256` Package

## 1. Concise Explanation

The Go standard-library `crypto/sha256` package implements the SHA-224 and SHA-256 cryptographic hash algorithms.

SHA-256 produces a **256-bit (32-byte)** hash, commonly displayed as a **64-character hexadecimal string**.

A hash is a one-way fingerprint of data:

```text
Input data
    ↓
SHA-256
    ↓
256-bit hash / 32 bytes
    ↓
64 hexadecimal characters
```

For example:

```text
"hello"
    ↓
2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824
```

Even a tiny change to the input produces a completely different hash.

### When is it commonly used?

`crypto/sha256` is useful for:

- Checking file/data integrity
- Creating content fingerprints
- Detecting whether data has changed
- Digital signatures and certificate-related systems
- Merkle trees and blockchain-related data structures
- Hash-based identifiers
- Verifying downloaded files
- Building cryptographic protocols

**Important:** SHA-256 is a hash function, **not encryption**. You cannot decrypt a SHA-256 hash back into the original data.

The package exposes four public functions and three constants.

---

# 2. Package Constants

## `sha256.BlockSize`

```go
const BlockSize = 64
```

`BlockSize` is the internal block size of SHA-256 and SHA-224 in **bytes**.

Example:

```go
fmt.Println(sha256.BlockSize)
```

Output:

```text
64
```

This means the algorithm processes its input internally using 64-byte blocks.

It is mostly useful when writing lower-level cryptographic code or code that needs to work with the `hash.Hash` interface.

---

## `sha256.Size`

```go
const Size = 32
```

This is the size of a SHA-256 checksum in bytes.

```go
fmt.Println(sha256.Size)
```

Output:

```text
32
```

Because:

```text
32 bytes × 8 bits = 256 bits
```

Therefore SHA-256 = **256 bits**.

When represented as hexadecimal:

```text
32 bytes × 2 hex characters = 64 characters
```

---

## `sha256.Size224`

```go
const Size224 = 28
```

This is the size of a SHA-224 checksum in bytes.

```text
28 bytes × 8 = 224 bits
```

Therefore SHA-224 produces:

```text
224-bit hash
28 bytes
56 hexadecimal characters
```

---

# 3. Every Function in `crypto/sha256`

The package contains four public functions:

```text
sha256.New()
sha256.New224()
sha256.Sum224(data)
sha256.Sum256(data)
```

---

## Function 1: `sha256.New()`

### Definition

```go
func New() hash.Hash
```

`New()` creates and returns a new `hash.Hash` that calculates **SHA-256** checksums.

### Example

```go
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	h := sha256.New()

	h.Write([]byte("Hello, Go!"))

	hash := h.Sum(nil)

	fmt.Printf("%x\n", hash)
}
```

### How it works

```go
h := sha256.New()
```

Creates a SHA-256 hash object.

Then:

```go
h.Write([]byte("Hello, Go!"))
```

Feeds data into the hash.

Finally:

```go
h.Sum(nil)
```

Returns the resulting hash.

Conceptually:

```text
New()
 ↓
Create SHA-256 hash object
 ↓
Write(data)
 ↓
Write(more data)
 ↓
Sum(nil)
 ↓
Final hash
```

### Why use `New()` instead of `Sum256()`?

`New()` is particularly useful when data arrives **incrementally**.

For example:

```go
h := sha256.New()

h.Write([]byte("Hello "))
h.Write([]byte("world"))
h.Write([]byte("!"))

result := h.Sum(nil)
```

Conceptually, this hashes:

```text
Hello world!
```

This is particularly useful for large files or streams because you don't need to load the entire input into memory first.

---

# Function 2: `sha256.New224()`

### Definition

```go
func New224() hash.Hash
```

`New224()` creates a `hash.Hash` that calculates a **SHA-224** checksum.

### Example

```go
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	h := sha256.New224()

	h.Write([]byte("Hello, Go!"))

	hash := h.Sum(nil)

	fmt.Printf("%x\n", hash)
}
```

The major difference is:

```text
New()       → SHA-256 → 32 bytes
New224()    → SHA-224 → 28 bytes
```

Both belong to the SHA-2 family.

### Why does SHA-224 exist?

SHA-224 provides a shorter digest than SHA-256:

```text
SHA-224 → 224 bits
SHA-256 → 256 bits
```

Consequently, SHA-224 produces:

```text
28 bytes
```

while SHA-256 produces:

```text
32 bytes
```

For most modern applications where SHA-256 is expected, use `New()` or `Sum256()` rather than choosing SHA-224 simply because the output is shorter.

---

# Function 3: `sha256.Sum224()`

### Definition

```go
func Sum224(data []byte) [Size224]byte
```

`Sum224()` directly calculates the **SHA-224 checksum** of a byte slice.

### Example

```go
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	data := []byte("Hello, Go!")

	hash := sha256.Sum224(data)

	fmt.Printf("%x\n", hash)
}
```

### Return type

The function returns:

```go
[Size224]byte
```

Since:

```go
Size224 == 28
```

the actual return type is effectively:

```go
[28]byte
```

This is an **array**, not a slice.

Example:

```go
hash := sha256.Sum224([]byte("hello"))

fmt.Println(len(hash))
```

Output:

```text
28
```

---

# Function 4: `sha256.Sum256()`

This is probably the function you'll use most often when learning SHA-256.

### Definition

```go
func Sum256(data []byte) [Size]byte
```

`Sum256()` calculates the **SHA-256 checksum** of the supplied data.

### Example

```go
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	data := []byte("Hello, Go!")

	hash := sha256.Sum256(data)

	fmt.Printf("%x\n", hash)
}
```

### Return type

The return type is:

```go
[Size]byte
```

and:

```go
Size == 32
```

So the function returns:

```go
[32]byte
```

Example:

```go
hash := sha256.Sum256([]byte("hello"))

fmt.Println(len(hash))
```

Output:

```text
32
```

---

# 4. `Sum256()` vs `New()`

This is an important distinction for beginners.

| Situation | Prefer |
|---|---|
| You already have all the data | `Sum256()` |
| Data arrives in chunks | `New()` |
| Hashing a small string | `Sum256()` |
| Hashing a large file/stream | `New()` |
| Need SHA-224 | `Sum224()` / `New224()` |

### Simple data already in memory

```go
hash := sha256.Sum256([]byte("hello"))
```

### Streaming data

```go
h := sha256.New()

h.Write([]byte("hello"))
h.Write([]byte(" world"))

hash := h.Sum(nil)
```

Both approaches can calculate a SHA-256 digest, but they are useful in different situations.

---

# 5. Understanding `hash.Hash`

There is another important concept behind:

```go
func New() hash.Hash
```

`hash.Hash` is an interface from Go's `hash` package.

A simplified view is:

```go
type Hash interface {
	Write(p []byte) (n int, err error)
	Sum(b []byte) []byte
	Reset()
	Size() int
	BlockSize() int
}
```

Therefore, after:

```go
h := sha256.New()
```

you can use methods such as:

```go
h.Write(...)
h.Sum(...)
h.Reset()
h.Size()
h.BlockSize()
```

These are **methods of the returned `hash.Hash`**, not additional functions belonging directly to `crypto/sha256`.

---

## `Write()`

```go
h.Write(data)
```

Adds data to the hash.

Example:

```go
h := sha256.New()

h.Write([]byte("Hello "))
h.Write([]byte("World"))
```

The two writes are effectively hashing:

```text
Hello World
```

This is extremely useful for streaming data.

---

## `Sum()`

```go
result := h.Sum(nil)
```

Returns the current hash digest.

One subtle point:

```go
h.Sum(existingSlice)
```

does **not** reset the hash.

For example:

```go
h.Write([]byte("hello"))

hash1 := h.Sum(nil)
hash2 := h.Sum(nil)
```

The hash state remains available.

---

## `Reset()`

```go
h.Reset()
```

Resets the hash object so it can be reused.

Example:

```go
h := sha256.New()

h.Write([]byte("hello"))
first := h.Sum(nil)

h.Reset()

h.Write([]byte("world"))
second := h.Sum(nil)
```

After `Reset()`, the previous `"hello"` data is no longer part of the calculation.

---

## `Size()`

```go
h.Size()
```

Returns the size of the resulting hash in bytes.

For SHA-256:

```text
32
```

For SHA-224:

```text
28
```

---

## `BlockSize()`

```go
h.BlockSize()
```

Returns the hash's underlying block size.

For SHA-256:

```text
64 bytes
```

This corresponds to:

```go
sha256.BlockSize
```

---

# 6. Complete Beginner Example

Here's a useful program combining the concepts:

```go
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	message := "Learning Go is fun!"

	hash := sha256.Sum256([]byte(message))

	fmt.Println("Message:", message)
	fmt.Printf("SHA-256: %x\n", hash)
	fmt.Println("Hash size:", len(hash), "bytes")
}
```

Conceptually:

```text
"Learning Go is fun!"
          │
          ▼
    sha256.Sum256()
          │
          ▼
     [32]byte hash
          │
          ▼
64 hexadecimal characters
```

---

# 7. Three Common Mistakes Beginners Make

## Mistake 1: Thinking SHA-256 is encryption

A common misconception is:

```text
SHA-256 → encrypt data → decrypt later
```

That's incorrect.

SHA-256 is a **cryptographic hash function**, not an encryption algorithm.

Encryption:

```text
plaintext → encryption → ciphertext
ciphertext → decryption → plaintext
```

Hashing:

```text
data → SHA-256 → hash
```

There is no normal "decrypt SHA-256" operation.

### How to avoid it

Remember:

> **Encryption is reversible; hashing is designed to be one-way.**

---

## Mistake 2: Using plain SHA-256 to store passwords

A beginner might write:

```go
hash := sha256.Sum256([]byte(password))
```

and store the result in a database.

Although SHA-256 itself is a strong cryptographic hash, **plain SHA-256 is not a suitable password-storage scheme**.

Password hashing should use a deliberately slow, password-specific algorithm such as Argon2id, scrypt, or bcrypt, with appropriate salts and parameters.

### How to avoid it

Use a password-hashing algorithm designed for password storage rather than simply:

```go
sha256.Sum256([]byte(password))
```

---

## Mistake 3: Comparing hexadecimal strings unnecessarily

You might see code such as:

```go
hash := sha256.Sum256(data)

hexHash := fmt.Sprintf("%x", hash)
```

The hexadecimal representation is useful for display or textual storage, but internally the digest is binary data.

Don't assume SHA-256 fundamentally produces a "64-character string."

It fundamentally produces:

```text
32 bytes = 256 bits
```

The 64-character hexadecimal representation is just one way of displaying those bytes.

### How to avoid it

Understand the distinction:

```text
SHA-256 result
      ↓
32 raw bytes
      ↓
hex encoding
      ↓
64-character string
```

---

# 8. Two Real-World Applications

## Application 1: File Integrity Verification

Suppose you download:

```text
myprogram.zip
```

The publisher gives you an expected SHA-256 value.

You calculate the SHA-256 hash of your downloaded file.

If your calculated hash matches the expected hash:

```text
Downloaded file
      ↓
SHA-256
      ↓
ABC123
```

and:

```text
Expected
      ↓
ABC123
```

then you have strong evidence that the file's contents match the published file.

A typical Go approach is:

```go
file, err := os.Open("file.txt")
if err != nil {
	log.Fatal(err)
}
defer file.Close()

h := sha256.New()

if _, err := io.Copy(h, file); err != nil {
	log.Fatal(err)
}

fmt.Printf("%x\n", h.Sum(nil))
```

This approach is particularly appropriate for large files because the file doesn't need to be loaded entirely into memory.

---

## Application 2: Digital Signatures and Content-Addressed Systems

SHA-256 is also used as a building block in larger cryptographic systems.

For example:

```text
Message
   ↓
SHA-256
   ↓
Digest
   ↓
Digital signature algorithm
   ↓
Signature
```

It is also useful in structures where data is identified by its content, such as Merkle-tree-based systems and various distributed-data systems.

An important distinction is that **SHA-256 by itself doesn't authenticate a message**. If an attacker can modify both the message and its hash, simply comparing the two doesn't prove who created the data. Authentication generally requires additional mechanisms such as digital signatures or a keyed construction such as HMAC.

---

# 9. Three Progressively Challenging Exercises

## Exercise 1 — Basic SHA-256

Write a Go program that:

1. Reads a string from the user.
2. Converts the string into `[]byte`.
3. Calculates its SHA-256 hash using `sha256.Sum256()`.
4. Prints the hash in hexadecimal format.
5. Also prints the hash length in bytes.

**Do not use `sha256.New()` for this exercise.**

---

## Exercise 2 — File Integrity Checker

Create a program that:

1. Accepts a filename from the command line.
2. Opens the file.
3. Calculates its SHA-256 hash.
4. Prints the hash as a hexadecimal string.
5. Handles file-opening and reading errors properly.
6. Works correctly with a large file without loading the entire file into memory.

**Hint:** Think about how `sha256.New()` can work together with an `io.Reader`.

---

## Exercise 3 — Build a File Verification Tool

Create a command-line program that accepts:

```text
program <filename> <expected-sha256>
```

For example:

```text
verify myfile.zip abc123...
```

Your program should:

1. Open the specified file.
2. Calculate its SHA-256 digest incrementally.
3. Convert the calculated digest into a suitable representation for comparison.
4. Compare it with the expected SHA-256 value.
5. Print whether the file passes or fails verification.
6. Handle invalid input and file errors.
7. Avoid unnecessarily loading the complete file into memory.
8. Consider what kind of comparison is appropriate when comparing cryptographic values.

### Bonus challenge

Modify the program so that it can verify multiple files from a manifest containing filenames and expected SHA-256 hashes.

---

# 10. Quick Reference

| API | Purpose | Result |
|---|---|---|
| `sha256.New()` | Create streaming SHA-256 hash | `hash.Hash` |
| `sha256.New224()` | Create streaming SHA-224 hash | `hash.Hash` |
| `sha256.Sum256(data)` | Direct SHA-256 calculation | `[32]byte` |
| `sha256.Sum224(data)` | Direct SHA-224 calculation | `[28]byte` |
| `sha256.BlockSize` | SHA block size | `64` bytes |
| `sha256.Size` | SHA-256 digest size | `32` bytes |
| `sha256.Size224` | SHA-224 digest size | `28` bytes |

## Most important distinction

### Complete data already available

```go
hash := sha256.Sum256(data)
```

### Data arrives incrementally

```go
h := sha256.New()
h.Write(chunk1)
h.Write(chunk2)
h.Write(chunk3)

hash := h.Sum(nil)
```

If you understand **that distinction, the difference between SHA-224 and SHA-256, and the fact that hashing is not encryption**, you've got the core of `crypto/sha256`.

---

# 11. Thought-Provoking Question

**If SHA-256 produces a 32-byte fingerprint for any input, including inputs that may be gigabytes in size, why is it computationally difficult to find two different inputs that produce the same hash—and what would happen to systems that rely on SHA-256 if someone discovered a practical way to do that?**

---

## Official Documentation

Go package documentation:

https://pkg.go.dev/crypto/sha256
