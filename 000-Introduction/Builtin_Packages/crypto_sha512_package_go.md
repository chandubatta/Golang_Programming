# `crypto/sha512` Package in Go

The Go `crypto/sha512` package provides implementations of the **SHA-512 family of cryptographic hash functions**. It is part of Go's standard library and is commonly used when you need to produce a fixed-size, one-way cryptographic digest of data.

## 1. What is `crypto/sha512`?

You import it with:

```go
import "crypto/sha512"
```

The package implements:

- **SHA-512** → 512-bit (64-byte) hash
- **SHA-384** → 384-bit (48-byte) hash
- **SHA-512/224** → 224-bit (28-byte) hash
- **SHA-512/256** → 256-bit (32-byte) hash

For example:

```text
Input:
Hello, Go!

        ↓ SHA-512

Output:
512-bit digest
64 bytes
128 hexadecimal characters
```

A cryptographic hash has several important properties:

1. The same input produces the same hash.
2. A tiny input change produces a dramatically different hash.
3. You cannot practically reconstruct the original input from the hash.
4. It is computationally difficult to find two different inputs with the same hash.
5. The output has a fixed size regardless of input size.

### When is it commonly used?

`crypto/sha512` is useful for:

- File integrity verification
- Digital-signature-related operations
- Content identification
- Generating cryptographic digests
- Hash-based protocols
- Verifying that downloaded data hasn't changed
- Building cryptographic constructions that specifically require SHA-512

**Important:** SHA-512 is a hash function, **not encryption**. You cannot decrypt a SHA-512 hash to recover the original data.

---

# 2. Important Functions and Types in `crypto/sha512`

The package is small, but understanding its API is important.

## `sha512.New()`

### Purpose

Creates a new SHA-512 hash object.

```go
h := sha512.New()
```

The returned value implements the `hash.Hash` interface.

You can then provide data using `Write()` and retrieve the final digest using `Sum()`.

### Example

```go
package main

import (
	"crypto/sha512"
	"fmt"
)

func main() {
	h := sha512.New()

	h.Write([]byte("Hello, Go!"))

	hash := h.Sum(nil)

	fmt.Printf("%x\n", hash)
}
```

### Why use `New()`?

`New()` is particularly useful when you want to hash data **incrementally**.

For example, suppose you have a very large file:

```text
Large File
   ↓
Read chunk 1 → Write()
Read chunk 2 → Write()
Read chunk 3 → Write()
...
   ↓
Sum()
```

You don't need to load the entire file into memory.

---

# `sha512.Sum512()`

This is one of the simplest functions in the package.

```go
func Sum512(data []byte) [Size]byte
```

It calculates the SHA-512 checksum of the supplied data.

### Example

```go
package main

import (
	"crypto/sha512"
	"fmt"
)

func main() {
	data := []byte("Hello, Go!")

	hash := sha512.Sum512(data)

	fmt.Printf("%x\n", hash)
}
```

The result is:

```go
[64]byte
```

because SHA-512 produces:

```text
512 bits
÷ 8
= 64 bytes
```

### When should you use it?

Use `Sum512()` when you already have all the data in memory and simply want its SHA-512 digest.

For example:

```go
hash := sha512.Sum512([]byte("hello"))
```

It's considerably simpler than creating a hash object manually.

---

# `sha512.New384()`

```go
func New384() hash.Hash
```

Creates a new **SHA-384** hash.

SHA-384 belongs to the SHA-2 family and internally uses the SHA-512 construction but produces a shorter digest.

Output:

```text
384 bits
= 48 bytes
= 96 hexadecimal characters
```

### Example

```go
package main

import (
	"crypto/sha512"
	"fmt"
)

func main() {
	h := sha512.New384()

	h.Write([]byte("Hello, Go!"))

	hash := h.Sum(nil)

	fmt.Printf("%x\n", hash)
}
```

### When is it useful?

Use SHA-384 when a system specifically requires SHA-384 or when its shorter digest is desirable while remaining part of the SHA-2 family.

---

# `sha512.Sum384()`

```go
func Sum384(data []byte) [Size384]byte
```

Calculates the SHA-384 digest of data.

### Example

```go
package main

import (
	"crypto/sha512"
	"fmt"
)

func main() {
	hash := sha512.Sum384([]byte("Hello, Go!"))

	fmt.Printf("%x\n", hash)
}
```

The result is:

```go
[48]byte
```

---

# `sha512.New512_224()`

```go
func New512_224() hash.Hash
```

Creates a SHA-512/224 hash object.

It produces:

```text
224 bits
= 28 bytes
```

Although the output is only 224 bits, it is based on the SHA-512 algorithm.

### Example

```go
package main

import (
	"crypto/sha512"
	"fmt"
)

func main() {
	h := sha512.New512_224()

	h.Write([]byte("Hello, Go!"))

	hash := h.Sum(nil)

	fmt.Printf("%x\n", hash)
}
```

### Why does it exist?

SHA-512/224 can be useful when a protocol requires a **224-bit digest** but specifically wants the SHA-512 family.

---

# `sha512.Sum512_224()`

```go
func Sum512_224(data []byte) [Size224]byte
```

Calculates a SHA-512/224 digest directly.

### Example

```go
package main

import (
	"crypto/sha512"
	"fmt"
)

func main() {
	hash := sha512.Sum512_224([]byte("Hello, Go!"))

	fmt.Printf("%x\n", hash)
}
```

The returned array contains:

```text
28 bytes
```

---

# `sha512.New512_256()`

```go
func New512_256() hash.Hash
```

Creates a SHA-512/256 hash object.

It produces:

```text
256 bits
= 32 bytes
```

### Example

```go
package main

import (
	"crypto/sha512"
	"fmt"
)

func main() {
	h := sha512.New512_256()

	h.Write([]byte("Hello, Go!"))

	hash := h.Sum(nil)

	fmt.Printf("%x\n", hash)
}
```

---

# `sha512.Sum512_256()`

```go
func Sum512_256(data []byte) [Size256]byte
```

Calculates the SHA-512/256 digest directly.

### Example

```go
package main

import (
	"crypto/sha512"
	"fmt"
)

func main() {
	hash := sha512.Sum512_256([]byte("Hello, Go!"))

	fmt.Printf("%x\n", hash)
}
```

The result contains:

```text
32 bytes
```

---

# Constants in the Package

The package also provides constants describing digest sizes.

## `sha512.Size`

```go
const Size = 64
```

This represents the size of a SHA-512 digest in bytes.

```go
fmt.Println(sha512.Size)
```

Output:

```text
64
```

---

## `sha512.Size224`

```go
const Size224 = 28
```

Size of a SHA-512/224 digest.

```text
224 bits = 28 bytes
```

---

## `sha512.Size256`

```go
const Size256 = 32
```

Size of a SHA-512/256 digest.

```text
256 bits = 32 bytes
```

---

## `sha512.Size384`

```go
const Size384 = 48
```

Size of a SHA-384 digest.

```text
384 bits = 48 bytes
```

---

# Understanding `hash.Hash`

Functions such as:

```go
sha512.New()
sha512.New384()
sha512.New512_224()
sha512.New512_256()
```

return objects implementing Go's `hash.Hash` interface.

That gives you methods such as:

## `Write()`

Adds data to the hash.

```go
h.Write([]byte("Hello"))
h.Write([]byte(" World"))
```

The package effectively hashes:

```text
Hello World
```

You can therefore process data in chunks.

---

## `Sum()`

Returns the current digest.

```go
result := h.Sum(nil)
```

A subtle but important point:

```go
h.Sum(nil)
```

does **not** reset the hash.

For example:

```go
h.Write([]byte("Hello"))

hash1 := h.Sum(nil)

h.Write([]byte(" World"))

hash2 := h.Sum(nil)
```

`hash2` represents:

```text
Hello World
```

not merely:

```text
 World
```

---

## `Reset()`

Resets the hash object so it can be reused.

```go
h.Reset()
```

After:

```go
h.Write([]byte("Hello"))
h.Reset()
h.Write([]byte("Goodbye"))
```

the resulting digest corresponds only to:

```text
Goodbye
```

---

## `Size()`

Returns the digest size in bytes.

For SHA-512:

```go
h := sha512.New()

fmt.Println(h.Size())
```

Output:

```text
64
```

---

## `BlockSize()`

Returns the hash function's underlying block size.

For SHA-512:

```go
h := sha512.New()

fmt.Println(h.BlockSize())
```

The block size is:

```text
128 bytes
```

This is different from the digest size.

### Important distinction

```text
SHA-512

Block size  = 128 bytes
Digest size = 64 bytes
```

Beginners often confuse these two.

---

# A Complete Basic Example

Here is a simple example using the most convenient API:

```go
package main

import (
	"crypto/sha512"
	"fmt"
)

func main() {
	message := "Hello, Go!"

	hash := sha512.Sum512([]byte(message))

	fmt.Printf("Message: %s\n", message)
	fmt.Printf("SHA-512: %x\n", hash)
}
```

The `%x` formatting converts the binary digest into hexadecimal.

Remember:

```text
64 bytes
↓
2 hexadecimal characters per byte
↓
128 hexadecimal characters
```

---

# Incremental Hashing Example

For larger data, use `New()`:

```go
package main

import (
	"crypto/sha512"
	"fmt"
)

func main() {
	h := sha512.New()

	h.Write([]byte("Hello"))
	h.Write([]byte(" "))
	h.Write([]byte("Go!"))

	hash := h.Sum(nil)

	fmt.Printf("SHA-512: %x\n", hash)
}
```

This is equivalent to hashing:

```text
Hello Go!
```

The ability to process data incrementally becomes particularly useful when dealing with files, network streams, or other large inputs.

---

# 3. Three Common Beginner Mistakes

## Mistake 1: Thinking SHA-512 is encryption

A beginner might think:

```text
Original data
    ↓
SHA-512
    ↓
Encrypted data
    ↓
Decrypt
```

That's incorrect.

SHA-512 is a **one-way cryptographic hash**.

Correct mental model:

```text
Data
 ↓
SHA-512
 ↓
Digest
```

There is no normal "decrypt SHA-512" operation.

### Avoid it

If you need confidentiality, use an appropriate encryption algorithm.

If you need a cryptographic digest, use SHA-512.

---

## Mistake 2: Using SHA-512 directly for password storage

You might see:

```go
hash := sha512.Sum512([]byte(password))
```

and think this is a secure password-storage solution.

It isn't.

Passwords should generally be processed with a **password-hashing/key-derivation algorithm designed for password storage**, such as Argon2id, bcrypt, scrypt, or PBKDF2, with appropriate parameters and a salt.

SHA-512 is designed as a general-purpose cryptographic hash, not as a deliberately expensive password-hashing function.

### Avoid it

Don't build a password database like:

```text
password
   ↓
SHA-512
   ↓
store hash
```

without using an appropriate password hashing scheme.

---

## Mistake 3: Confusing bytes with hexadecimal text

Suppose:

```go
hash := sha512.Sum512([]byte("hello"))
```

The result is:

```go
[64]byte
```

It is **binary data**, not a 128-character string.

When you do:

```go
fmt.Printf("%x", hash)
```

you're displaying those 64 bytes as hexadecimal.

So:

```text
64 bytes
```

becomes:

```text
128 hexadecimal characters
```

### Avoid it

Understand whether your application needs:

- raw digest bytes, or
- hexadecimal representation of the digest.

For APIs, files, or databases, the appropriate representation depends on the protocol/design.

---

# 4. Two Real-World Applications

## Application 1: File Integrity Verification

Suppose you download a large software file.

The publisher gives you a SHA-512 checksum:

```text
Expected:
ABC123...
```

You calculate:

```text
Downloaded file
      ↓
   SHA-512
      ↓
Calculated hash
```

Then compare:

```text
Expected hash == Calculated hash
```

If they match, you have strong evidence that the file contents are identical to the data represented by that checksum.

This is useful for:

- Software downloads
- Backups
- Large datasets
- Build artifacts
- Archive verification

---

## Application 2: Cryptographic Protocols and Digital Signatures

SHA-512 can be used as part of cryptographic systems where a message needs to be represented by a fixed-size digest.

For example:

```text
Large message
      ↓
   SHA-512
      ↓
   64-byte digest
      ↓
Cryptographic operation
```

Instead of operating directly on an arbitrarily large message, a cryptographic system can work with its digest.

The exact hash algorithm and signature scheme should always be determined by the protocol/library specification rather than chosen arbitrarily.

---

# 5. Progressive Exercises

## Exercise 1 — Basic SHA-512 Hash

Write a Go program that:

1. Accepts a string from the user.
2. Calculates its SHA-512 hash.
3. Prints the original string.
4. Prints the SHA-512 digest in hexadecimal.
5. Prints the number of bytes in the digest.

**Hint:** Investigate `sha512.Sum512()` and `sha512.Size`.

---

## Exercise 2 — Hash a Large File

Create a Go program that accepts a file path and calculates the file's SHA-512 checksum.

Requirements:

1. Open the file.
2. Read it incrementally rather than loading the entire file into memory.
3. Feed the data into a SHA-512 hash.
4. Print the final digest in hexadecimal.
5. Handle file-opening and reading errors correctly.
6. Make sure the file is closed properly.

Your program should work even when the file is several gigabytes in size.

---

## Exercise 3 — Build a File Integrity Checker

Build a small command-line integrity verification tool.

The program should accept:

```text
program <filename> <expected-sha512>
```

It should:

1. Open the specified file.
2. Calculate its SHA-512 digest incrementally.
3. Convert the digest into a representation suitable for comparison with the supplied expected value.
4. Compare the calculated digest with the expected digest.
5. Report whether the file passes or fails the integrity check.
6. Handle invalid arguments.
7. Handle files that don't exist.
8. Handle invalid digest input.
9. Avoid loading the entire file into memory.

**Bonus challenge:** Make the comparison resistant to timing side channels where appropriate.

---

# Quick API Cheat Sheet

| API | Purpose | Output |
|---|---|---|
| `sha512.New()` | Create SHA-512 hash | `hash.Hash` |
| `sha512.Sum512()` | Hash data directly | `[64]byte` |
| `sha512.New384()` | Create SHA-384 hash | `hash.Hash` |
| `sha512.Sum384()` | Hash data directly with SHA-384 | `[48]byte` |
| `sha512.New512_224()` | Create SHA-512/224 hash | `hash.Hash` |
| `sha512.Sum512_224()` | Hash data directly | `[28]byte` |
| `sha512.New512_256()` | Create SHA-512/256 hash | `hash.Hash` |
| `sha512.Sum512_256()` | Hash data directly | `[32]byte` |
| `sha512.Size` | SHA-512 digest size | `64` |
| `sha512.Size224` | SHA-512/224 digest size | `28` |
| `sha512.Size256` | SHA-512/256 digest size | `32` |
| `sha512.Size384` | SHA-384 digest size | `48` |

And the `hash.Hash` returned by the `New...` functions provides:

```text
Write()      → add data
Sum()        → obtain digest
Reset()      → clear current state
Size()       → digest size
BlockSize()  → internal block size
```

# One Deeper Question

Imagine you're designing a system that verifies **10 GB files uploaded by 10,000 users simultaneously**. Why might using:

```go
sha512.Sum512(fileBytes)
```

be a poor design even though SHA-512 itself is perfectly capable of hashing the file—and how would the streaming:

```go
sha512.New()
```

approach change the application's **memory usage and scalability**?

That question gets you beyond simply learning the API and into understanding **why the API is designed the way it is**.
