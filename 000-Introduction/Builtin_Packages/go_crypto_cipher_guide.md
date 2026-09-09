# Go `crypto/cipher` Package — Detailed Guide

## 1. What is `crypto/cipher`?

The Go `crypto/cipher` package provides interfaces and implementations for using block ciphers in practical encryption modes.

It commonly works together with `crypto/aes`:

```text
crypto/aes
    │
    ▼
AES block cipher
    │
    ▼
crypto/cipher
    │
    ├── GCM  → authenticated encryption
    ├── CBC  → block chaining
    ├── CTR  → stream-like encryption
    ├── CFB  → stream-like encryption (deprecated)
    └── OFB  → stream-like encryption (deprecated)
```

`crypto/aes` answers:

> Which encryption algorithm?

`crypto/cipher` answers:

> How should that block cipher be used to encrypt data?

For new applications, authenticated encryption such as AES-GCM is generally the preferred choice.

---

## 2. Why do we need `crypto/cipher`?

A block cipher such as AES works on fixed-size blocks. AES has a block size of 16 bytes.

Real application data can be:

- JSON
- database records
- files
- messages
- HTTP payloads
- session information

Cipher modes provide ways to use a block cipher with such data.

---

## 3. Important concepts

### Plaintext

The original data:

```text
Hello World
```

### Ciphertext

The encrypted representation of the plaintext.

### Key

The secret value used by the cipher.

### IV

An Initialization Vector used by modes such as CBC.

### Nonce

A value used by modes such as GCM and CTR. For GCM, a nonce must be unique for a given key.

### Authentication tag

GCM produces an authentication tag that lets the receiver detect tampering.

---

# 4. `crypto/cipher` API

The major interfaces are:

```text
AEAD
Block
BlockMode
Stream
StreamReader
StreamWriter
```

Important constructors/functions include:

```text
NewGCM
NewGCMWithNonceSize
NewGCMWithRandomNonce
NewGCMWithTagSize

NewCBCEncrypter
NewCBCDecrypter

NewCFBEncrypter
NewCFBDecrypter

NewCTR

NewOFB
```

---

# 5. `cipher.Block`

The `Block` interface represents a low-level block cipher.

Conceptually:

```go
type Block interface {
    BlockSize() int
    Encrypt(dst, src []byte)
    Decrypt(dst, src []byte)
}
```

AES implements this interface.

Example:

```go
block, err := aes.NewCipher(key)
```

---

## `BlockSize()`

```go
BlockSize() int
```

Returns the cipher's block size in bytes.

For AES:

```go
block.BlockSize()
```

returns:

```text
16
```

Example:

```go
block, _ := aes.NewCipher(key)

fmt.Println(block.BlockSize())
```

Output:

```text
16
```

This is useful when a mode requires an IV of exactly one block.

---

## `Encrypt()`

```go
Encrypt(dst, src []byte)
```

Encrypts one block.

For AES:

```text
16-byte plaintext
      │
      ▼
     AES
      │
      ▼
16-byte ciphertext
```

Example:

```go
block.Encrypt(ciphertext, plaintext)
```

This is a low-level operation. Application code normally uses a cipher mode rather than manually calling `Encrypt()` repeatedly.

---

## `Decrypt()`

```go
Decrypt(dst, src []byte)
```

Decrypts one block.

Conceptually:

```text
ciphertext
    │
    ▼
  Decrypt
    │
    ▼
 plaintext
```

This is also a low-level operation.

---

# 6. `cipher.AEAD`

AEAD means:

> Authenticated Encryption with Associated Data

It is one of the most important modern cryptographic interfaces in Go.

Conceptually:

```go
type AEAD interface {
    NonceSize() int
    Overhead() int
    Seal(dst, nonce, plaintext, additionalData []byte) []byte
    Open(dst, nonce, ciphertext, additionalData []byte) ([]byte, error)
}
```

AES-GCM is a common implementation.

---

## `NonceSize()`

```go
NonceSize() int
```

Returns the required nonce size.

Example:

```go
gcm, _ := cipher.NewGCM(block)

fmt.Println(gcm.NonceSize())
```

Standard GCM commonly uses a 12-byte nonce.

You can allocate it using:

```go
nonce := make([]byte, gcm.NonceSize())
```

A nonce is not secret, but it must not be reused with the same key.

Good:

```text
KEY A + NONCE 1 → message 1
KEY A + NONCE 2 → message 2
KEY A + NONCE 3 → message 3
```

Bad:

```text
KEY A + NONCE 1 → message 1
KEY A + NONCE 1 → message 2
```

---

## `Overhead()`

```go
Overhead() int
```

Returns the maximum difference between plaintext and ciphertext lengths.

For standard GCM, this is typically the authentication tag size, commonly 16 bytes.

For example:

```text
plaintext = 100 bytes
GCM overhead = 16 bytes
ciphertext = 116 bytes
```

The nonce is normally stored separately or prepended by the application.

---

## `Seal()`

```go
Seal(dst, nonce, plaintext, additionalData []byte) []byte
```

Encrypts and authenticates plaintext.

Example:

```go
ciphertext := gcm.Seal(
    nil,
    nonce,
    plaintext,
    nil,
)
```

Conceptually:

```text
             plaintext
                 │
                 ▼
             ┌───────┐
additional → │  GCM  │
data         └───┬───┘
                 │
          ┌──────┴──────┐
          ▼             ▼
     ciphertext       tag
          └──────┬──────┘
                 ▼
             Seal result
```

---

## `Open()`

```go
Open(dst, nonce, ciphertext, additionalData []byte) ([]byte, error)
```

Authenticates and decrypts ciphertext.

Example:

```go
plaintext, err := gcm.Open(
    nil,
    nonce,
    ciphertext,
    nil,
)

if err != nil {
    // Authentication failed
}
```

If the ciphertext or associated data was modified, `Open()` returns an error.

---

# 7. Additional Authenticated Data (AAD)

AEAD allows data to be authenticated without encrypting it.

Suppose an application has:

```text
User ID: 12345
Role: admin
Message: hello
```

The message could be encrypted while metadata is supplied as additional authenticated data.

Example:

```go
aad := []byte("user=12345")

ciphertext := gcm.Seal(
    nil,
    nonce,
    plaintext,
    aad,
)
```

During decryption, the same AAD must be supplied:

```go
plaintext, err := gcm.Open(
    nil,
    nonce,
    ciphertext,
    aad,
)
```

If an attacker changes the authenticated metadata, authentication fails.

---

# 8. `cipher.NewGCM()`

```go
cipher.NewGCM(block)
```

Creates GCM around a block cipher.

Example:

```go
block, err := aes.NewCipher(key)
if err != nil {
    panic(err)
}

gcm, err := cipher.NewGCM(block)
if err != nil {
    panic(err)
}
```

For normal AES encryption in a new application, this is usually the most important `crypto/cipher` constructor to learn.

---

# 9. `NewGCMWithNonceSize()`

```go
NewGCMWithNonceSize(cipher Block, size int)
```

Creates GCM with a custom nonce size.

Example:

```go
gcm, err := cipher.NewGCMWithNonceSize(block, 12)
```

Generally, use `cipher.NewGCM(block)` unless an existing protocol requires a particular nonce size.

---

# 10. `NewGCMWithTagSize()`

```go
NewGCMWithTagSize(cipher Block, tagSize int)
```

Creates GCM with a custom authentication tag size.

Example:

```go
gcm, err := cipher.NewGCMWithTagSize(block, 16)
```

Use this primarily for interoperability with an existing cryptographic format or protocol.

For normal applications, prefer the standard `NewGCM()`.

---

# 11. `NewGCMWithRandomNonce()`

```go
cipher.NewGCMWithRandomNonce(block)
```

Creates a GCM AEAD that handles nonce generation internally when sealing.

This can reduce the risk of application-level nonce-management mistakes.

It is still important to understand the normal `NewGCM()` API and the nonce requirements.

---

# 12. `cipher.BlockMode`

`BlockMode` is used by block-based modes such as CBC.

Conceptually:

```go
type BlockMode interface {
    BlockSize() int
    CryptBlocks(dst, src []byte)
}
```

It provides:

```text
BlockSize()
CryptBlocks()
```

---

## `BlockMode.BlockSize()`

Returns the block size.

For AES:

```go
mode.BlockSize()
```

returns:

```text
16
```

---

## `BlockMode.CryptBlocks()`

```go
CryptBlocks(dst, src []byte)
```

Encrypts or decrypts multiple complete blocks.

Example:

```go
mode.CryptBlocks(ciphertext, plaintext)
```

For AES, input length must be a multiple of 16 bytes.

Valid:

```text
16 bytes
32 bytes
48 bytes
```

Invalid:

```text
50 bytes
```

CBC normally requires padding before `CryptBlocks()`.

---

# 13. `NewCBCEncrypter()`

```go
cipher.NewCBCEncrypter(block, iv)
```

Creates CBC encryption mode.

Conceptually:

```text
Plaintext Block 1
       │
       XOR IV
       │
       ▼
      AES
       │
       ▼
Ciphertext Block 1
       │
       ▼
Plaintext Block 2
       │
       XOR
Ciphertext Block 1
       │
       ▼
      AES
       │
       ▼
Ciphertext Block 2
```

Example:

```go
mode := cipher.NewCBCEncrypter(block, iv)

mode.CryptBlocks(ciphertext, plaintext)
```

### Security warning

CBC provides confidentiality but does not provide authentication.

Ciphertext should be authenticated separately, such as with an appropriate MAC construction.

For new applications, prefer AES-GCM.

---

# 14. `NewCBCDecrypter()`

```go
cipher.NewCBCDecrypter(block, iv)
```

Creates CBC decryption mode.

Example:

```go
mode := cipher.NewCBCDecrypter(block, iv)

mode.CryptBlocks(plaintext, ciphertext)
```

Encryption and decryption both use `CryptBlocks()`; the difference is whether the mode was created using `NewCBCEncrypter()` or `NewCBCDecrypter()`.

---

# 15. `cipher.Stream`

`Stream` represents a stream-like cipher.

Its main method is:

```go
XORKeyStream(dst, src []byte)
```

Conceptually:

```text
plaintext
    XOR
keystream
    =
ciphertext
```

For decryption:

```text
ciphertext
    XOR
same keystream
    =
plaintext
```

---

## `Stream.XORKeyStream()`

```go
XORKeyStream(dst, src []byte)
```

Encrypts or decrypts data using a keystream.

It can often operate in-place:

```go
stream.XORKeyStream(data, data)
```

The operation is based on XOR, so the same basic transformation is used for encryption and decryption.

A separately initialized stream with the same key and initial state is needed for the reverse operation.

---

# 16. `NewCTR()`

```go
cipher.NewCTR(block, iv)
```

Creates Counter Mode.

Conceptually:

```text
Counter 1 → AES → Keystream 1
Counter 2 → AES → Keystream 2
Counter 3 → AES → Keystream 3
```

Then:

```text
Plaintext XOR Keystream = Ciphertext
```

Example:

```go
stream := cipher.NewCTR(block, iv)

stream.XORKeyStream(ciphertext, plaintext)
```

CTR can process arbitrary-length data.

### Critical security rule

Never reuse the same key + IV/counter combination for different messages.

CTR also does not authenticate ciphertext, so it should not be treated as a complete authenticated encryption scheme.

---

# 17. `NewCFBEncrypter()` — deprecated

```go
cipher.NewCFBEncrypter(block, iv)
```

Creates CFB encryption.

CFB is deprecated in current Go.

It is useful mainly for:

- Understanding cryptography concepts
- Reading legacy code
- Compatibility with existing systems
- Exams/interviews

For new applications, use an AEAD mode such as GCM.

---

# 18. `NewCFBDecrypter()` — deprecated

```go
cipher.NewCFBDecrypter(block, iv)
```

Creates CFB decryption.

Example:

```go
stream := cipher.NewCFBDecrypter(block, iv)

stream.XORKeyStream(plaintext, ciphertext)
```

CFB should generally be avoided in new application designs.

---

# 19. `NewOFB()` — deprecated

```go
cipher.NewOFB(block, iv)
```

Creates Output Feedback mode.

It converts a block cipher into a stream-like cipher.

Example:

```go
stream := cipher.NewOFB(block, iv)

stream.XORKeyStream(ciphertext, plaintext)
```

OFB is deprecated in current Go.

For new applications, prefer authenticated encryption such as AES-GCM.

---

# 20. `StreamReader`

`StreamReader` connects a cipher stream to Go's `io.Reader`.

Conceptually:

```text
encrypted reader
       │
       ▼
 StreamReader
       │
       ▼
decrypted reader
```

The structure conceptually contains:

```go
type StreamReader struct {
    S Stream
    R io.Reader
}
```

This is useful for processing streams without loading the entire input into memory.

---

## `StreamReader.Read()`

```go
Read(dst []byte) (n int, err error)
```

Reads from the underlying reader and applies the stream cipher.

Example:

```go
reader := &cipher.StreamReader{
    S: stream,
    R: encryptedReader,
}

io.Copy(output, reader)
```

---

# 21. `StreamWriter`

`StreamWriter` connects a cipher stream to Go's `io.Writer`.

Conceptually:

```text
application
    │
    ▼
StreamWriter
    │
    ▼
encrypted output
```

It conceptually contains:

```go
type StreamWriter struct {
    S Stream
    W io.Writer
}
```

---

## `StreamWriter.Write()`

```go
Write(src []byte) (n int, err error)
```

Encrypts data and writes it to the underlying writer.

Example:

```go
writer := &cipher.StreamWriter{
    S: stream,
    W: encryptedFile,
}

writer.Write(data)
```

Useful for encrypting files and network streams.

---

## `StreamWriter.Close()`

```go
Close() error
```

Closes the underlying writer if it implements `io.Closer`.

This allows `StreamWriter` to integrate with Go's normal I/O abstractions.

---

# 22. Simple practical example — AES-GCM

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
    // 32-byte key = AES-256
    key := []byte("01234567890123456789012345678901")

    plaintext := []byte("Hello, this is secret data!")

    // Create AES block cipher.
    block, err := aes.NewCipher(key)
    if err != nil {
        panic(err)
    }

    // Wrap AES with GCM.
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        panic(err)
    }

    // Generate a unique random nonce.
    nonce := make([]byte, gcm.NonceSize())

    if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
        panic(err)
    }

    // Encrypt and authenticate.
    ciphertext := gcm.Seal(
        nil,
        nonce,
        plaintext,
        nil,
    )

    fmt.Printf("Nonce:      %x\n", nonce)
    fmt.Printf("Ciphertext: %x\n", ciphertext)

    // Decrypt and authenticate.
    decrypted, err := gcm.Open(
        nil,
        nonce,
        ciphertext,
        nil,
    )

    if err != nil {
        panic(err)
    }

    fmt.Printf("Plaintext: %s\n", decrypted)
}
```

The important flow is:

```text
KEY
 │
 ▼
aes.NewCipher(key)
 │
 ▼
AES Block
 │
 ▼
cipher.NewGCM(block)
 │
 ▼
GCM
 │
 ├── Seal()
 │      ↓
 │   ciphertext + authentication
 │
 └── Open()
        ↓
     verification + plaintext
```

---

# 23. Three common beginner mistakes

## Mistake 1: Reusing a GCM nonce

Bad:

```text
key = same

message 1 → nonce = 123
message 2 → nonce = 123
message 3 → nonce = 123
```

Better:

```text
message 1 → fresh nonce A
message 2 → fresh nonce B
message 3 → fresh nonce C
```

For GCM, nonce reuse with the same key can seriously compromise security.

---

## Mistake 2: Thinking encryption automatically provides authentication

Beginners sometimes think:

```text
encrypted = secure
```

Not necessarily.

Modes such as CTR and CBC provide encryption but do not inherently authenticate the ciphertext.

An attacker may potentially modify ciphertext.

AES-GCM provides:

```text
Encryption
+
Authentication
=
Authenticated encryption
```

---

## Mistake 3: Using deprecated modes for new applications

You may find code using:

```go
cipher.NewCFBEncrypter(...)
cipher.NewOFB(...)
```

Their presence in the standard library does not mean they are recommended for new applications.

Current Go documentation marks CFB and OFB as deprecated and recommends AEAD modes for new applications.

For new code, start with:

```text
AES-GCM
```

---

# 24. `crypto/cipher` vs `crypto/aes`

This distinction is essential.

### `crypto/aes`

Answers:

> Which encryption algorithm?

Answer:

```text
AES
```

### `crypto/cipher`

Answers:

> How should that block cipher be used?

Examples:

```text
GCM
CBC
CTR
CFB
OFB
```

So:

```go
block, _ := aes.NewCipher(key)
```

means:

> Give me AES.

While:

```go
gcm, _ := cipher.NewGCM(block)
```

means:

> Use that AES cipher in GCM mode.

---

# 25. Real-world application #1 — Encrypting sensitive database fields

Imagine an application storing:

```text
customer_id
name
email
phone
bank_account
```

Sensitive fields can be encrypted before storage.

Conceptually:

```text
Application
    │
    ▼
Sensitive data
    │
    ▼
AES-GCM
    │
    ▼
Ciphertext + nonce
    │
    ▼
Database
```

When retrieving:

```text
Database
    │
    ▼
Ciphertext + nonce
    │
    ▼
AES-GCM Open()
    │
    ▼
Original data
```

This can help protect sensitive data at rest.

---

# 26. Real-world application #2 — Encrypting files

A backup application might encrypt files before storing them:

```text
Original file
     │
     ▼
AES-GCM
     │
     ▼
Encrypted backup
```

For very large files, an application should carefully design a chunked encryption format rather than loading the entire file into memory. This involves decisions about:

- Chunk size
- Nonces
- Authentication tags
- File metadata
- Versioning
- Error handling
- Key management

`StreamReader` and `StreamWriter` demonstrate how stream ciphers can integrate with Go's `io.Reader` and `io.Writer` interfaces.

---

# 27. Important security distinction

Think about encryption as:

```text
Encryption
    ↓
"Can outsiders read it?"
```

Authentication as:

```text
Authentication
    ↓
"Can outsiders modify it without us detecting it?"
```

A modern secure design often wants:

```text
Confidentiality
+
Integrity
+
Authenticity
```

AEAD provides these properties together.

That is why:

```go
cipher.NewGCM(...)
```

is such an important API.

---

# 28. When should you use each mode?

| Mode | Current recommendation | Main idea |
|---|---|---|
| GCM | Recommended | Authenticated encryption |
| CTR | Specialized | Stream-like encryption, no authentication |
| CBC | Legacy/interoperability | Block chaining, no authentication |
| CFB | Deprecated | Stream-like mode |
| OFB | Deprecated | Stream-like mode |

As a beginner, prioritize:

```text
1. Block
2. AEAD
3. GCM
4. Seal
5. Open
6. Nonce
7. AdditionalData
8. CTR/CBC for understanding legacy systems
```

---

# 29. Exercise 1 — AES-GCM message encryption

Create a Go program that:

1. Generates or accepts a 32-byte AES key.
2. Creates an AES cipher.
3. Creates a GCM cipher.
4. Generates a random nonce.
5. Encrypts a user-provided message.
6. Decrypts the ciphertext.
7. Prints the original and decrypted messages.
8. Detects and reports an authentication error if the ciphertext is modified.

**Do not hard-code the nonce.**

---

# 30. Exercise 2 — Secure encrypted message format

Build a small encryption package that provides:

```text
Encrypt(plaintext)
Decrypt(encryptedData)
```

Design a byte format that stores everything necessary for decryption, such as:

```text
version | nonce | ciphertext
```

Requirements:

- Generate a fresh nonce for every encryption.
- Store the nonce together with the ciphertext.
- Use AES-GCM.
- Support arbitrary plaintext lengths.
- Detect modified ciphertext.
- Detect an incorrect key.
- Reject malformed encrypted data.

Think carefully about how your encrypted data format should be structured.

---

# 31. Exercise 3 — Encrypted file storage system

Build a command-line application:

```text
securefile encrypt input.txt output.enc
securefile decrypt output.enc restored.txt
```

Requirements:

- Use AES-GCM.
- Never reuse a nonce for separate encryption operations.
- Process the file without loading the entire file into memory.
- Design a file format containing the information necessary for decryption.
- Detect corrupted encrypted files.
- Detect an incorrect encryption key.
- Handle files larger than available RAM.
- Explain how your design prevents nonce reuse.
- Explain what happens if an attacker modifies one encrypted portion of the file.

### Bonus challenge

Design chunked encryption where each chunk has its own nonce and authentication tag.

---

# 32. Mental model to remember

```text
                    crypto/aes
                       │
                       ▼
                  AES Block
                       │
                 crypto/cipher
                       │
          ┌────────────┼────────────┐
          │            │            │
          ▼            ▼            ▼
         GCM          CBC          CTR
          │            │            │
      AEAD mode    Block mode   Stream mode
          │
     ┌────┴────┐
     ▼         ▼
   Seal       Open
     │         │
     ▼         ▼
 Encrypt     Verify +
 + Auth      Decrypt
```

The most important modern pattern is:

```go
key
 ↓
aes.NewCipher(key)
 ↓
cipher.NewGCM(block)
 ↓
nonce
 ↓
gcm.Seal(...)
 ↓
ciphertext
```

Then:

```go
ciphertext
 ↓
gcm.Open(...)
 ↓
authentication check
 ↓
plaintext
```

---

# 33. Thought-provoking question

If AES-GCM already provides encryption and authentication, why might an application still need to carefully decide which metadata belongs in `additionalData`?

What could go wrong if an attacker cannot read a message but can manipulate metadata that tells your application:

- who the message belongs to,
- what operation it represents,
- which resource it targets, or
- how it should be processed?

Think about this from the perspective of a real API or distributed system where **confidentiality alone is not enough**.
