# `crypto/md5` Package in Go

## 1. What is `crypto/md5`?

The Go `crypto/md5` package implements the **MD5 (Message-Digest Algorithm 5)** hash algorithm, standardized in RFC 1321.

MD5 accepts arbitrary input and produces a fixed-size **128-bit (16-byte) digest**, commonly represented as a **32-character hexadecimal string**.

```text
Input
  ↓
MD5
  ↓
128-bit / 16-byte digest
  ↓
32 hexadecimal characters
```

For example:

```text
MD5("hello")
    ↓
5d41402abc4b2a76b9719d911017c592
```

The same input always produces the same digest.

### Important security warning

MD5 is **cryptographically broken** because practical collision attacks exist.

Do **not** use MD5 for:

- Password hashing
- Digital signatures
- Certificates
- Authentication
- Security-sensitive integrity checks where an attacker can deliberately manipulate data

For new security-sensitive applications, prefer SHA-256/SHA-512 or an appropriate modern construction.

### When is MD5 still used?

MD5 can still be useful when:

- Interoperating with a legacy system that requires MD5
- Working with legacy protocols or file formats
- Computing non-security-oriented checksums where deliberate collision resistance is not required
- Comparing data against an MD5 checksum supplied by a trusted source

---

# 2. Functions and Constants in `crypto/md5`

The public API of `crypto/md5` is intentionally small. It provides constants and functions for creating and calculating MD5 hashes.

## `md5.Size`

```go
const Size = 16
```

`Size` is the size of an MD5 digest in **bytes**.

Example:

```go
package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	fmt.Println(md5.Size)
}
```

Output:

```text
16
```

Since:

```text
16 bytes × 8 = 128 bits
```

MD5 produces a 128-bit digest.

---

## `md5.BlockSize`

```go
const BlockSize = 64
```

`BlockSize` is the internal block size used by MD5, measured in bytes.

```text
64 bytes = 512 bits
```

Example:

```go
package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	fmt.Println(md5.BlockSize)
}
```

Output:

```text
64
```

Do not confuse:

```text
Digest size = 16 bytes
Block size  = 64 bytes
```

The digest size describes the final MD5 result, while the block size describes the chunks used internally by the algorithm.

---

# `md5.New()`

```go
func New() hash.Hash
```

`New()` creates and returns a new MD5 hash object implementing Go's `hash.Hash` interface.

Example:

```go
h := md5.New()
```

You can feed data into it:

```go
h.Write([]byte("hello"))
```

Then obtain the digest:

```go
result := h.Sum(nil)
```

Convert it to hexadecimal:

```go
fmt.Printf("%x\n", result)
```

### Complete example

```go
package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	h := md5.New()

	h.Write([]byte("hello"))

	hash := h.Sum(nil)

	fmt.Printf("%x\n", hash)
}
```

Output:

```text
5d41402abc4b2a76b9719d911017c592
```

### Why does `New()` return `hash.Hash`?

`hash.Hash` is an interface shared by many Go hash algorithms.

For example:

```go
md5.New()
sha256.New()
sha1.New()
sha512.New()
```

can all be used through a common interface.

The important `hash.Hash` methods are:

```go
Write(p []byte) (n int, err error)
Sum(b []byte) []byte
Reset()
Size() int
BlockSize() int
```

---

# 3. Understanding the `hash.Hash` Methods

Because `md5.New()` returns a `hash.Hash`, understanding these methods is important.

## `Write()`

```go
h.Write(data)
```

`Write()` adds data to the current hash calculation.

Example:

```go
h := md5.New()

h.Write([]byte("Hello"))
h.Write([]byte(" World"))
```

Conceptually, MD5 sees:

```text
Hello World
```

rather than two independent hashes.

You can therefore process data in pieces:

```go
h.Write([]byte("part 1"))
h.Write([]byte("part 2"))
h.Write([]byte("part 3"))
```

This is equivalent to hashing:

```text
part 1part 2part 3
```

### Why is `Write()` useful?

It allows streaming large data without loading everything into memory at once.

For example, a large file can be read in chunks and supplied to the hash.

---

## `Sum()`

```go
h.Sum(nil)
```

`Sum()` returns the current digest.

Important: **`Sum()` does not reset the hash.**

Example:

```go
h.Write([]byte("hello"))

sum1 := h.Sum(nil)
sum2 := h.Sum(nil)
```

Both represent the same current hash state.

`Sum()` also accepts a byte slice to which the digest is appended:

```go
prefix := []byte("hash: ")

result := h.Sum(prefix)
```

Conceptually:

```text
result = "hash: " + MD5(...)
```

---

## `Reset()`

`Reset()` clears the current hash state.

Example:

```go
h := md5.New()

h.Write([]byte("hello"))

fmt.Printf("%x\n", h.Sum(nil))

h.Reset()

h.Write([]byte("world"))

fmt.Printf("%x\n", h.Sum(nil))
```

After `Reset()`, `"hello"` is no longer part of the calculation.

This can be useful when reusing a hash object.

---

## `Size()`

`Size()` returns the number of bytes in the digest.

For MD5:

```go
h := md5.New()

fmt.Println(h.Size())
```

Output:

```text
16
```

That corresponds to:

```text
128 bits
```

---

## `BlockSize()`

`BlockSize()` returns the internal block size of the hash algorithm.

For MD5:

```go
h := md5.New()

fmt.Println(h.BlockSize())
```

Output:

```text
64
```

So:

```text
Digest size = 16 bytes
Block size  = 64 bytes
```

---

# 4. A Simple One-Shot MD5 Example

For small data that is already available in memory, `md5.Sum()` is convenient.

```go
package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	data := []byte("hello")

	hash := md5.Sum(data)

	fmt.Printf("%x\n", hash)
}
```

Output:

```text
5d41402abc4b2a76b9719d911017c592
```

`md5.Sum()` returns a `[16]byte` digest.

---

# 5. `md5.New()` vs `md5.Sum()`

## Use `md5.Sum()`

Use it when you already have the entire input:

```go
data := []byte("hello")

hash := md5.Sum(data)
```

This is simple and convenient.

## Use `md5.New()`

Use it when you want to stream data:

```go
h := md5.New()

h.Write(part1)
h.Write(part2)
h.Write(part3)

hash := h.Sum(nil)
```

Streaming is particularly useful for large files because you don't have to load the entire file into memory.

### Large-file example

```go
package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("large-file.iso")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	h := md5.New()

	_, err = io.Copy(h, file)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%x\n", h.Sum(nil))
}
```

Conceptually:

```text
File
 ↓
io.Copy
 ↓
MD5 hash
 ↓
16-byte digest
```

---

# 6. Three Common Beginner Mistakes

## Mistake 1: Thinking MD5 is Encryption

MD5 is a **hash**, not encryption.

Encryption is designed to be reversible with a key:

```text
plaintext → encryption → ciphertext
ciphertext → decryption → plaintext
```

Hashing works differently:

```text
data → hash → digest
```

There is no normal "decrypt MD5" operation.

### How to avoid it

Remember:

> Encryption is designed to be reversible with a key; hashing produces a digest and is not normally reversible.

---

## Mistake 2: Using MD5 for Passwords

A beginner might write:

```go
passwordHash := md5.Sum([]byte(password))
```

and store the result.

**Do not do this for password storage.**

MD5 is extremely fast and has known cryptographic weaknesses, making it inappropriate for password hashing.

For passwords, use a password-specific algorithm such as:

- Argon2id
- bcrypt
- scrypt
- PBKDF2

The important principle is:

> Password hashing should deliberately be expensive and resistant to brute-force attacks.

---

## Mistake 3: Assuming Hashes Cannot Collide

A collision occurs when two different inputs produce the same hash.

You might initially imagine:

```text
A → hash A
B → hash B
```

with every possible input producing a unique result.

But MD5 produces only:

```text
2^128
```

possible digests while accepting arbitrarily large inputs.

More importantly, practical attacks have demonstrated that MD5's collision resistance is broken.

Therefore:

```text
different data
      ↓
potentially same MD5
```

is a real security concern.

### How to avoid it

Never use MD5 when an attacker being able to construct a collision could cause a security problem.

---

# 7. Two Real-World Applications

## Application 1: Legacy File Checksum Verification

Suppose an old software distribution provides:

```text
application.zip
MD5: 8d777f385d3dfec8815d20f7496026dc
```

You can calculate the MD5 of your downloaded file and compare it.

This can detect accidental corruption during transfer.

However, **MD5 should not be treated as a strong security guarantee against a malicious attacker**, because collision attacks exist.

For new systems, SHA-256 is generally a better choice.

---

## Application 2: Legacy Protocol/System Compatibility

You may need to communicate with an existing application that explicitly specifies MD5.

For example:

```text
Legacy application
       ↓
expects MD5
       ↓
Go application
       ↓
crypto/md5
```

In this situation, replacing MD5 with SHA-256 may break interoperability.

The reason to use MD5 here isn't that MD5 is secure; it is that the external protocol requires it.

---

# 8. Three Progressively Challenging Exercises

## Exercise 1 — Basic MD5 Calculator

Write a Go program that:

1. Accepts a string from the user.
2. Calculates its MD5 hash.
3. Prints the digest as a lowercase hexadecimal string.
4. Also prints the digest size in bytes.

Do not use a solution from this explanation. Build it yourself using `crypto/md5`.

---

## Exercise 2 — MD5 File Checksum Tool

Create a command-line program:

```text
go run main.go myfile.txt
```

The program should:

1. Open the specified file.
2. Calculate its MD5 checksum.
3. Process the file without loading the entire file into memory.
4. Print the resulting MD5 checksum.
5. Handle file-opening and reading errors properly.

### Challenge

Use the streaming approach with `md5.New()`.

---

## Exercise 3 — Build a Checksum Verification Tool

Create a command-line program that accepts:

```text
go run main.go myfile.iso expected-md5
```

It should:

1. Open the file.
2. Calculate its MD5 checksum while streaming the file.
3. Convert the calculated digest to hexadecimal.
4. Compare it with the expected checksum.
5. Print whether the checksums match.
6. Handle invalid input and file errors.
7. Avoid accidentally comparing checksums in a way that depends on the strings' lengths or causes confusing behavior.

### Extra challenge

Modify the program so that it can verify multiple files from a checksum manifest.

---

# 9. The Big Picture

The most important parts to remember are:

```text
crypto/md5
│
├── Size       → 16 bytes
│
├── BlockSize  → 64 bytes
│
├── New()      → creates streaming MD5 hash
│
└── Sum()      → computes MD5 directly from []byte
```

When using `New()`:

```text
md5.New()
    ↓
Write(...)
    ↓
Write(...)
    ↓
Write(...)
    ↓
Sum(nil)
    ↓
16-byte MD5 digest
```

## Security Rule

> MD5 is useful for legacy compatibility and some non-adversarial checksum scenarios, but it should not be selected for new security-sensitive designs.

---

# 10. Thought-Provoking Question

If MD5 is known to be cryptographically broken, **why do you think Go continues to include `crypto/md5` in the standard library instead of removing it completely?**

Consider the difference between:

1. An algorithm being unsafe for new security designs, and
2. An algorithm still being necessary for compatibility with existing systems.

What problems could developers face if an old but widely used algorithm were suddenly removed from the standard library?
