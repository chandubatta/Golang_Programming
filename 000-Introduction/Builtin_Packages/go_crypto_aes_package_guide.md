# Go `crypto/aes` Package — Detailed Learning Guide

## 1. What is `crypto/aes`?

Go's `crypto/aes` package implements **AES (Advanced Encryption Standard)**, a symmetric block cipher.

**Symmetric encryption** means the same secret key is used for encryption and decryption.

AES supports three key sizes:

| AES | Key size | Block size |
|---|---:|---:|
| AES-128 | 16 bytes | 16 bytes |
| AES-192 | 24 bytes | 16 bytes |
| AES-256 | 32 bytes | 16 bytes |

The block size is **always 16 bytes**, regardless of the key size.

Typical architecture in Go:

```text
                 crypto/aes
                     │
                     ▼
              aes.NewCipher()
                     │
                     ▼
              cipher.Block
                     │
          ┌──────────┴──────────┐
          ▼                     ▼
     crypto/cipher          Direct blocks
          │
          ▼
       AES-GCM
       AES-CTR
       AES-CBC
       etc.
```

For most new applications, **AES-GCM** is the preferred choice because it provides both:

- confidentiality — hides the plaintext
- authentication/integrity — detects tampering

`crypto/aes` itself is a low-level block-cipher package. It does not provide a convenient "encrypt this arbitrary string" API.

---

# 2. Every public function/type in `crypto/aes`

The standard `crypto/aes` package has a very small public API:

```text
BlockSize

NewCipher()

KeySizeError
    Error()
```

---

## 2.1 `aes.BlockSize`

Definition:

```go
const BlockSize = 16
```

`BlockSize` represents the AES block size in **bytes**.

Therefore:

```text
16 bytes = 128 bits
```

Example:

```go
package main

import (
	"crypto/aes"
	"fmt"
)

func main() {
	fmt.Println(aes.BlockSize)
}
```

Output:

```text
16
```

### Important distinction

Don't confuse **block size** with **key size**.

```text
AES-128
    key   = 16 bytes
    block = 16 bytes

AES-192
    key   = 24 bytes
    block = 16 bytes

AES-256
    key   = 32 bytes
    block = 16 bytes
```

So:

```go
aes.BlockSize
```

is always:

```text
16
```

---

## 2.2 `aes.NewCipher()`

This is the **main function** in the package.

Signature:

```go
func NewCipher(key []byte) (cipher.Block, error)
```

It creates an AES cipher using the supplied key.

Example:

```go
key := []byte("1234567890123456")

block, err := aes.NewCipher(key)
if err != nil {
	panic(err)
}
```

Here the key is 16 bytes, so Go creates an **AES-128** cipher.

### Valid key lengths

```text
16 bytes → AES-128
24 bytes → AES-192
32 bytes → AES-256
```

For example:

```go
key128 := make([]byte, 16)
key192 := make([]byte, 24)
key256 := make([]byte, 32)
```

All three are valid.

But:

```go
key := []byte("hello")
```

is invalid because it is only 5 bytes.

`NewCipher` will return an error.

### What does `NewCipher()` return?

Notice:

```go
(cipher.Block, error)
```

It does **not** return an `aes.Cipher`.

It returns the interface:

```go
cipher.Block
```

The `cipher.Block` interface provides:

```go
type Block interface {
	BlockSize() int
	Encrypt(dst, src []byte)
	Decrypt(dst, src []byte)
}
```

This allows `crypto/cipher` to work with different block ciphers through a common interface.

For example:

```go
block, err := aes.NewCipher(key)
if err != nil {
	panic(err)
}

fmt.Println(block.BlockSize())
```

Output:

```text
16
```

---

## 2.3 `KeySizeError`

`KeySizeError` is the error type used when an invalid AES key size is supplied.

Conceptually:

```go
type KeySizeError int
```

Suppose you do:

```go
key := []byte("hello")

_, err := aes.NewCipher(key)

if err != nil {
	fmt.Println(err)
}
```

The error indicates that the key has an invalid length.

The valid sizes are:

```text
16
24
32
```

---

## 2.4 `KeySizeError.Error()`

The type has:

```go
func (k KeySizeError) Error() string
```

This implements Go's `error` interface.

You normally don't call it yourself.

Instead:

```go
_, err := aes.NewCipher(key)

if err != nil {
	fmt.Println(err)
}
```

When you print `err`, Go automatically uses its `Error()` method.

The normal beginner pattern is simply:

```go
if err != nil {
	return err
}
```

rather than manually calling:

```go
err.Error()
```

---

# 3. Simple AES example

Let's start with the **lowest-level example**.

AES encrypts one 16-byte block at a time when you use `cipher.Block` directly.

```go
package main

import (
	"crypto/aes"
	"fmt"
)

func main() {
	key := []byte("1234567890123456")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	plaintext := []byte("Hello AES World!")

	ciphertext := make([]byte, aes.BlockSize)

	block.Encrypt(ciphertext, plaintext)

	fmt.Printf("Plaintext : %s\n", plaintext)
	fmt.Printf("Encrypted : %x\n", ciphertext)

	decrypted := make([]byte, aes.BlockSize)

	block.Decrypt(decrypted, ciphertext)

	fmt.Printf("Decrypted : %s\n", decrypted)
}
```

Notice:

```text
"Hello AES World!"
```

is exactly **16 bytes**, which is required for direct `Block.Encrypt`.

The output will look roughly like:

```text
Plaintext : Hello AES World!
Encrypted : ...
Decrypted : Hello AES World!
```

### What happened?

```text
                 KEY
                  │
                  ▼
        ┌──────────────────┐
        │  aes.NewCipher() │
        └────────┬─────────┘
                 │
                 ▼
           cipher.Block
                 │
                 ▼
        ┌──────────────────┐
        │ block.Encrypt()  │
        └────────┬─────────┘
                 │
                 ▼
            Ciphertext
                 │
                 ▼
        ┌──────────────────┐
        │ block.Decrypt()  │
        └────────┬─────────┘
                 │
                 ▼
             Plaintext
```

**Important:** Don't use this direct-block approach for ordinary application encryption. For real applications, use an authenticated construction such as AES-GCM.

---

# 4. Recommended real-world approach: AES-GCM

A much more realistic example is:

```go
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func main() {
	key := []byte("1234567890123456")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err)
	}

	plaintext := []byte("Hello, AES-GCM!")

	ciphertext := gcm.Seal(nil, nonce, plaintext, nil)

	fmt.Printf("Nonce     : %x\n", nonce)
	fmt.Printf("Ciphertext: %x\n", ciphertext)

	decrypted, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Decrypted : %s\n", decrypted)
}
```

Here:

```text
aes.NewCipher()
       ↓
cipher.Block
       ↓
cipher.NewGCM()
       ↓
cipher.AEAD
       ↓
Seal() / Open()
```

This is much closer to how AES should normally be used in an application.

---

# 5. Understanding `Encrypt()` and `Decrypt()`

These aren't functions directly declared by the `crypto/aes` package API. They are methods provided through the `cipher.Block` interface returned by `NewCipher`.

## `Encrypt`

```go
block.Encrypt(dst, src)
```

It encrypts **one AES block**.

Both source and destination need enough space for one block:

```text
16 bytes
```

Example:

```go
plaintext := []byte("Hello AES World!")

ciphertext := make([]byte, aes.BlockSize)

block.Encrypt(ciphertext, plaintext)
```

Conceptually:

```text
Encrypt:

Plaintext ──AES + Key──> Ciphertext
```

---

## `Decrypt`

```go
block.Decrypt(dst, src)
```

It decrypts one AES block.

Example:

```go
plaintext := make([]byte, aes.BlockSize)

block.Decrypt(plaintext, ciphertext)
```

Conceptually:

```text
Decrypt:

Ciphertext ──AES + Key──> Plaintext
```

---

# 6. Why doesn't `crypto/aes` encrypt arbitrary strings directly?

This is a very important concept.

AES itself is a **block cipher**.

It works on:

```text
128-bit blocks
     ↓
16 bytes
```

Suppose you have:

```text
"Hello"
```

That's only 5 bytes.

Or:

```text
"This is a very long message..."
```

which may be hundreds of bytes.

You therefore need a **mode of operation** to safely process arbitrary-length data.

That's where:

```go
crypto/cipher
```

comes in.

For example:

```text
AES
 │
 ├── GCM
 ├── CTR
 ├── CBC
 └── other modes
```

For modern application encryption, AES-GCM is generally the easiest choice.

---

# 7. Three common beginner mistakes

## Mistake 1: Using an invalid key size

Beginners sometimes write:

```go
key := []byte("mypassword")
```

and expect AES to accept it.

It won't.

AES requires:

```text
16 bytes
24 bytes
32 bytes
```

### How to avoid it

Don't simply pad or truncate passwords.

A password/passphrase isn't automatically a cryptographically suitable AES key.

Instead, use an appropriate **password-based key derivation function (KDF)** when the key originates from a human password.

---

## Mistake 2: Thinking AES automatically provides authentication

Encryption answers:

> "Can someone read my data?"

But you also need to consider:

> "Can someone modify my encrypted data?"

Plain AES encryption alone does not automatically give you authenticated encryption.

For example, modes such as CTR provide confidentiality but don't inherently authenticate the ciphertext.

That's why **AES-GCM** is so useful: it provides authenticated encryption.

---

## Mistake 3: Reusing a nonce/IV incorrectly

When using AES-GCM, the nonce must be handled correctly.

A common mistake is:

```text
same key
+
same nonce
+
multiple messages
```

That can seriously compromise security.

A common safe pattern is:

```text
Generate random nonce
        ↓
Encrypt message
        ↓
Store/transmit:

nonce + ciphertext
```

The nonce does **not** need to be secret.

The key does.

---

# 8. Real-world application #1: Encrypting sensitive application data

Imagine an application storing sensitive information.

For example:

```text
Database
   │
   ├── User profile
   ├── Encrypted sensitive data
   └── Other application data
```

Instead of storing sensitive plaintext:

```text
Sensitive information
```

the application can encrypt it:

```text
plaintext
    ↓
AES-GCM
    ↓
ciphertext + authentication tag
```

The application keeps the AES key separately from the encrypted data.

This can be useful for protecting sensitive application-level data if the database contents are exposed.

---

# 9. Real-world application #2: Encrypting files/backups

AES can also be used as part of a system for protecting:

```text
Documents
Backups
Configuration data
Archives
Application exports
```

For example:

```text
original file
     │
     ▼
AES-GCM encryption
     │
     ▼
encrypted backup
```

When the backup needs to be restored:

```text
encrypted backup
       │
       ▼
AES-GCM decryption
       │
       ▼
original file
```

In a production design, you'd also need to carefully handle key management, nonce storage, authentication, key rotation, and recovery.

---

# 10. Exercises

## Exercise 1 — Basic AES block encryption

Write a Go program that:

1. Creates a valid 16-byte AES key.
2. Creates a 16-byte plaintext.
3. Calls `aes.NewCipher()`.
4. Encrypts the plaintext using `block.Encrypt()`.
5. Prints the ciphertext in hexadecimal.
6. Decrypts the ciphertext using `block.Decrypt()`.
7. Prints the original plaintext.
8. Verify that the decrypted data exactly matches the original plaintext.

**Do not use `crypto/cipher` yet.**

---

## Exercise 2 — AES-GCM message encryption

Build a small program that encrypts an arbitrary-length message using AES-GCM.

Your program should:

1. Generate or load a valid AES key.
2. Create an AES cipher using `aes.NewCipher()`.
3. Create an AES-GCM instance.
4. Generate a secure random nonce.
5. Encrypt a message of your choice.
6. Store/transmit the nonce together with the ciphertext.
7. Extract the nonce and ciphertext.
8. Decrypt the message.
9. Verify that the decrypted message equals the original message.
10. Modify one byte of the ciphertext and observe what happens during decryption.

---

## Exercise 3 — Build a reusable encrypted-message format

Design a small reusable encryption package with functions conceptually like:

```go
Encrypt(key, plaintext)
Decrypt(key, encryptedData)
```

Your design should:

1. Use AES-GCM.
2. Generate a fresh nonce for every encryption operation.
3. Define a format for storing the nonce and ciphertext together.
4. Validate malformed encrypted data.
5. Detect authentication failures.
6. Support arbitrary-length plaintext.
7. Avoid storing the encryption key alongside the encrypted data.
8. Write unit tests for successful encryption/decryption.
9. Write tests for corrupted ciphertext.
10. Write tests for an incorrect key.
11. Write tests for malformed input.
12. Explain how your design would handle key rotation in a production system.

---

# 11. The most important mental model

When learning `crypto/aes`, keep these layers separate:

```text
                 Your Application
                       │
                       ▼
              Encrypt sensitive data
                       │
                       ▼
                AES-GCM / AEAD
                 crypto/cipher
                       │
                       ▼
                 AES block cipher
                  crypto/aes
                       │
                       ▼
                    AES key
```

So don't think:

> "`crypto/aes` is a complete encryption system."

Think:

> **"`crypto/aes` provides the AES block cipher; `crypto/cipher` provides useful modes such as GCM that turn that primitive into an application-level encryption construction."**

That distinction will save you from many cryptography mistakes.

---

# 12. Thought-provoking question

Suppose you build a system that correctly uses AES-256-GCM, but an attacker steals the AES key from your server.

**Is your encrypted data still secure?**

If not:

**What architectural changes would you make so that compromising one key does not expose every piece of historical data?**

This leads into deeper topics such as:

- key management
- key rotation
- envelope encryption
- KMS/HSMs
- key hierarchy
- forward secrecy
- limiting the blast radius of key compromise

These concepts are just as important as knowing how to call `aes.NewCipher()`.

---

# Quick Reference

| API | Purpose |
|---|---|
| `aes.BlockSize` | Returns the AES block size: 16 bytes |
| `aes.NewCipher(key)` | Creates an AES block cipher |
| `aes.KeySizeError` | Represents an invalid AES key size |
| `KeySizeError.Error()` | Returns the error message |
| `block.BlockSize()` | Returns the block size through `cipher.Block` |
| `block.Encrypt()` | Encrypts one 16-byte block |
| `block.Decrypt()` | Decrypts one 16-byte block |

## Valid AES keys

```text
16 bytes → AES-128
24 bytes → AES-192
32 bytes → AES-256
```

## Recommended architecture

```text
Application
    ↓
AES-GCM / AEAD
    ↓
AES block cipher
    ↓
Securely managed AES key
```

## Key takeaway

**`crypto/aes` gives you the AES primitive. For practical application-level encryption, normally combine it with an authenticated mode such as AES-GCM from `crypto/cipher`.**
