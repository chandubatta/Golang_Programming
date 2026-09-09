# Go `encoding/pem` Package — Detailed Guide

## 1. What is `encoding/pem`?

The Go standard-library `encoding/pem` package implements **PEM (Privacy-Enhanced Mail) encoding and decoding**.

PEM is a textual representation commonly used to store or transmit cryptographic material such as:

- TLS certificates
- Public keys
- Private keys
- Certificate chains
- Certificate signing requests

A typical PEM block looks like:

```text
-----BEGIN CERTIFICATE-----
MIID...
...
-----END CERTIFICATE-----
```

The important idea is:

> **PEM is a container/text format; it is not itself an encryption algorithm.**

The binary data inside a PEM block is commonly Base64-encoded DER data. The `encoding/pem` package handles the PEM wrapper and Base64 conversion; packages such as `crypto/x509`, `crypto/rsa`, `crypto/ecdsa`, and `crypto/ed25519` interpret the decoded cryptographic bytes.

Import it with:

```go
import "encoding/pem"
```

### Typical workflow

```text
RSA/ECDSA/Ed25519 key
        ↓
binary representation
        ↓
DER / ASN.1
        ↓
PEM encoding
        ↓
-----BEGIN ...-----
Base64 text
-----END ...-----
```

When reading:

```text
PEM text
   ↓
pem.Decode()
   ↓
binary DER bytes
   ↓
x509.Parse...
   ↓
actual Go key/certificate
```

So `encoding/pem` is generally the **formatting/container layer**, not the cryptographic parsing layer.

---

# 2. Functions and types inside `encoding/pem`

The package is intentionally small. Its main public API consists of:

| API | Purpose |
|---|---|
| `pem.Encode` | Writes a PEM block to an `io.Writer` |
| `pem.EncodeToMemory` | Converts a PEM block directly into `[]byte` |
| `pem.Decode` | Extracts the next PEM block from input |
| `pem.Block` | Represents a PEM block |

---

# 3. `pem.Block`

Before understanding the functions, understand `pem.Block`.

```go
type Block struct {
    Type    string
    Headers map[string]string
    Bytes   []byte
}
```

## `Type`

Identifies what the PEM block contains.

Examples:

```go
Type: "CERTIFICATE"
```

```go
Type: "PUBLIC KEY"
```

```go
Type: "RSA PRIVATE KEY"
```

Common PEM labels include:

```text
CERTIFICATE
PUBLIC KEY
PRIVATE KEY
RSA PRIVATE KEY
EC PRIVATE KEY
CERTIFICATE REQUEST
```

### Important

The `Type` field does **not** cause Go to interpret the data.

For example:

```go
block.Type == "PUBLIC KEY"
```

doesn't automatically turn `block.Bytes` into an `*rsa.PublicKey`.

You might subsequently do:

```go
pub, err := x509.ParsePKIXPublicKey(block.Bytes)
```

---

## `Headers`

Headers are optional metadata.

Example:

```go
block := &pem.Block{
    Type: "MESSAGE",
    Headers: map[string]string{
        "Animal": "Gopher",
    },
    Bytes: []byte("hello"),
}
```

Conceptually, this can produce:

```text
-----BEGIN MESSAGE-----
Animal: Gopher

aGVsbG8=
-----END MESSAGE-----
```

Most modern certificate/key PEM files don't require custom headers.

---

## `Bytes`

This is the actual binary payload.

For example:

```go
block := &pem.Block{
    Type:  "MESSAGE",
    Bytes: []byte("hello"),
}
```

`Bytes` is **not the Base64 text**.

You provide:

```text
hello
```

and PEM encoding produces its Base64 representation:

```text
aGVsbG8=
```

When decoding a PEM block, `block.Bytes` contains the decoded binary contents.

---

# 4. `pem.Encode`

Signature:

```go
func Encode(out io.Writer, b *Block) error
```

`Encode` writes a PEM representation of `b` to an `io.Writer`.

## Simple example

```go
package main

import (
    "encoding/pem"
    "os"
    "log"
)

func main() {
    block := &pem.Block{
        Type:  "MESSAGE",
        Bytes: []byte("Hello, Gopher!"),
    }

    err := pem.Encode(os.Stdout, block)
    if err != nil {
        log.Fatal(err)
    }
}
```

Output:

```text
-----BEGIN MESSAGE-----
SGVsbG8sIEdvcGhlciE=
-----END MESSAGE-----
```

## Why does `Encode` use `io.Writer`?

You can write directly to:

- a file
- `os.Stdout`
- a network connection
- a buffer
- another writer

Example:

```go
file, err := os.Create("message.pem")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

err = pem.Encode(file, block)
if err != nil {
    log.Fatal(err)
}
```

This is useful when you don't need the complete encoded PEM in memory.

---

# 5. `pem.EncodeToMemory`

Signature:

```go
func EncodeToMemory(b *Block) []byte
```

Instead of writing to an `io.Writer`, it returns the encoded PEM as a byte slice.

Example:

```go
package main

import (
    "encoding/pem"
    "fmt"
)

func main() {
    block := &pem.Block{
        Type:  "MESSAGE",
        Bytes: []byte("Hello, Gopher!"),
    }

    data := pem.EncodeToMemory(block)

    fmt.Println(string(data))
}
```

Output:

```text
-----BEGIN MESSAGE-----
SGVsbG8sIEdvcGhlciE=
-----END MESSAGE-----
```

## When should you use it?

Use:

```go
pemData := pem.EncodeToMemory(block)
```

when you want the PEM data in memory.

For example:

```go
pemData := pem.EncodeToMemory(block)

err := os.WriteFile("message.pem", pemData, 0600)
```

## Difference between `Encode` and `EncodeToMemory`

`Encode`:

```go
pem.Encode(writer, block)
```

writes directly to a writer.

`EncodeToMemory`:

```go
data := pem.EncodeToMemory(block)
```

returns the complete PEM representation.

`EncodeToMemory` returns `nil` if the block cannot be encoded because of invalid headers. If you need details about an encoding error, use `Encode`.

---

# 6. `pem.Decode`

Signature:

```go
func Decode(data []byte) (p *Block, rest []byte)
```

This is one of the most important functions when **reading certificates and keys**.

It searches the input for the next PEM block.

Example:

```go
data := []byte(`
-----BEGIN MESSAGE-----
SGVsbG8=
-----END MESSAGE-----
`)

block, rest := pem.Decode(data)
```

You can then inspect:

```go
fmt.Println(block.Type)
fmt.Println(string(block.Bytes))
fmt.Println(string(rest))
```

Conceptually:

```text
Input:

PEM BLOCK
PEM BLOCK
some other data

       ↓ pem.Decode()

block = first PEM block
rest  = everything after first PEM block
```

If no PEM data is found, `block` is `nil` and the entire input is returned as `rest`.

---

# 7. Why does `pem.Decode` return two values?

This is useful when a file contains multiple PEM blocks.

For example:

```text
-----BEGIN CERTIFICATE-----
...
-----END CERTIFICATE-----

-----BEGIN CERTIFICATE-----
...
-----END CERTIFICATE-----
```

You can decode them one at a time:

```go
rest := data

for {
    block, remaining := pem.Decode(rest)

    if block == nil {
        break
    }

    fmt.Println("Type:", block.Type)

    rest = remaining
}
```

This is especially useful for certificate chains containing multiple certificates.

---

# 8. Simple complete example

This example demonstrates both encoding and decoding.

```go
package main

import (
    "encoding/pem"
    "fmt"
    "log"
)

func main() {
    // Create a PEM block.
    block := &pem.Block{
        Type:  "MESSAGE",
        Bytes: []byte("Hello from Go!"),
    }

    // Encode it into memory.
    encoded := pem.EncodeToMemory(block)

    fmt.Println("Encoded PEM:")
    fmt.Println(string(encoded))

    // Decode the PEM.
    decoded, rest := pem.Decode(encoded)

    if decoded == nil {
        log.Fatal("failed to decode PEM")
    }

    fmt.Println("Type:", decoded.Type)
    fmt.Println("Data:", string(decoded.Bytes))
    fmt.Println("Remaining:", len(rest))
}
```

Output will be approximately:

```text
Encoded PEM:
-----BEGIN MESSAGE-----
SGVsbG8gZnJvbSBHbyE=
-----END MESSAGE-----

Type: MESSAGE
Data: Hello from Go!
Remaining: 0
```

The round trip is:

```text
original data
     ↓
pem.EncodeToMemory
     ↓
PEM text
     ↓
pem.Decode
     ↓
original data
```

---

# 9. `encoding/pem` with `crypto/x509`

This is where `encoding/pem` becomes particularly important in real Go applications.

Consider a certificate:

```text
-----BEGIN CERTIFICATE-----
MIID...
...
-----END CERTIFICATE-----
```

You first decode the PEM:

```go
block, _ := pem.Decode(certPEM)
```

Then parse the DER bytes using `crypto/x509`:

```go
cert, err := x509.ParseCertificate(block.Bytes)
```

So:

```text
PEM
 ↓
pem.Decode()
 ↓
DER bytes
 ↓
x509.ParseCertificate()
 ↓
*x509.Certificate
```

This demonstrates the separation of responsibilities:

- `encoding/pem` handles the PEM container/text encoding.
- `crypto/x509` interprets certificate/key structures encoded inside the PEM block.

---

# 10. Real-world example: reading an RSA public key

Suppose you have:

```text
public.pem
```

containing:

```text
-----BEGIN PUBLIC KEY-----
...
-----END PUBLIC KEY-----
```

You could do:

```go
package main

import (
    "crypto/x509"
    "encoding/pem"
    "fmt"
    "log"
    "os"
)

func main() {
    data, err := os.ReadFile("public.pem")
    if err != nil {
        log.Fatal(err)
    }

    block, _ := pem.Decode(data)

    if block == nil {
        log.Fatal("failed to decode PEM")
    }

    if block.Type != "PUBLIC KEY" {
        log.Fatalf("unexpected PEM type: %s", block.Type)
    }

    publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Public key type: %T\n", publicKey)
}
```

The responsibility is divided like this:

```text
os
 ↓
read file

encoding/pem
 ↓
remove PEM wrapper + Base64 decode

crypto/x509
 ↓
interpret DER/ASN.1 public-key structure
```

---

# 11. Three common beginner mistakes

## Mistake 1: Thinking PEM is encryption

A common misconception is:

> "PEM protects my private key."

Not necessarily.

PEM is primarily an **encoding/container format**.

For example:

```text
-----BEGIN PRIVATE KEY-----
...
-----END PRIVATE KEY-----
```

doesn't automatically mean the private key is encrypted.

PEM commonly wraps Base64-encoded binary data.

### Avoid it

Think:

```text
PEM ≠ encryption
PEM ≈ textual container/encoding
```

If sensitive data needs encryption, you need an appropriate cryptographic mechanism rather than simply PEM-encoding it.

---

## Mistake 2: Assuming `pem.Decode` parses the key

Beginners sometimes write:

```go
block, _ := pem.Decode(data)
```

and think they now have an RSA/ECDSA/Ed25519 key.

They don't.

They have:

```go
*pem.Block
```

The actual key material is still inside:

```go
block.Bytes
```

You generally need another parser:

```text
pem.Decode()
     ↓
block.Bytes
     ↓
x509.ParsePKCS8PrivateKey()
```

or another appropriate parser depending on the key format.

### Avoid it

Always ask:

> "What format are the bytes inside this PEM block?"

Then choose the appropriate parser.

---

## Mistake 3: Ignoring `block == nil`

This is dangerous:

```go
block, _ := pem.Decode(data)

fmt.Println(block.Type)
```

If no PEM block exists, `block` can be `nil`.

Instead:

```go
block, rest := pem.Decode(data)

if block == nil {
    return
}
```

Also don't blindly assume:

```go
block.Type == "PRIVATE KEY"
```

Always validate the expected type when your application requires a specific kind of PEM block.

---

# 12. Two real-world applications

## Application 1: TLS certificates

One of the most common uses of PEM is TLS.

For example:

```text
server.crt
server.key
```

may contain:

```text
-----BEGIN CERTIFICATE-----
...
-----END CERTIFICATE-----
```

and:

```text
-----BEGIN PRIVATE KEY-----
...
-----END PRIVATE KEY-----
```

Go's TLS ecosystem frequently works with PEM-formatted certificates and keys.

Example:

```go
certPEM, err := os.ReadFile("server.crt")
if err != nil {
    log.Fatal(err)
}

keyPEM, err := os.ReadFile("server.key")
if err != nil {
    log.Fatal(err)
}
```

These PEM representations can then be used as part of TLS configuration.

---

## Application 2: Public/private key exchange and configuration

Applications frequently store cryptographic keys in PEM files.

For example:

```text
config/
    public.pem
    private.pem
```

An application can process them using:

```text
Read file
   ↓
pem.Decode
   ↓
block.Bytes
   ↓
x509 parser
   ↓
Go crypto key
```

This is useful for:

- JWT signing keys
- Digital signatures
- Certificate authorities
- SSH-related workflows
- Service-to-service authentication
- TLS infrastructure

---

# 13. Three progressively challenging exercises

## Exercise 1 — Basic PEM round trip

Create a Go program that:

1. Creates a `pem.Block`.
2. Uses `"MESSAGE"` as the block type.
3. Stores a short message inside `Bytes`.
4. Encodes it using `pem.EncodeToMemory`.
5. Prints the resulting PEM.
6. Decodes it using `pem.Decode`.
7. Prints the decoded type and original message.
8. Verifies that there is no remaining data.

**Goal:** Become comfortable with `Block`, `EncodeToMemory`, and `Decode`.

---

## Exercise 2 — Multiple PEM blocks

Create a program containing **three PEM blocks** in one byte slice:

```text
MESSAGE
MESSAGE
MESSAGE
```

Each block should contain a different message.

Your program should:

1. Decode the first block.
2. Process it.
3. Use the returned `rest` value.
4. Continue decoding until no blocks remain.
5. Print each block's `Type` and decoded contents.
6. Correctly handle the case where the input contains no PEM block.

**Goal:** Understand why `pem.Decode` returns both `block` and `rest`.

---

## Exercise 3 — Certificate inspection

Create a program that reads a PEM-encoded X.509 certificate from a file.

Your program should:

1. Read the certificate file.
2. Use `pem.Decode` to extract the PEM block.
3. Verify that the block is not `nil`.
4. Verify that the PEM type is `"CERTIFICATE"`.
5. Pass `block.Bytes` to `x509.ParseCertificate`.
6. Print useful certificate information such as:
   - Subject
   - Issuer
   - DNS names
   - NotBefore
   - NotAfter
7. Handle malformed PEM data.
8. Handle malformed certificate DER data.
9. Extend the program so that it can process multiple certificates from a certificate chain.

**Goal:** Understand how `encoding/pem` works together with `crypto/x509` in a realistic application.

---

# 14. The most important mental model

When learning `encoding/pem`, remember this pipeline:

```text
                  WRITING
                    │
                    ▼
        Go cryptographic object
                    │
                    ▼
             DER / ASN.1 bytes
                    │
                    ▼
            pem.Block{Bytes: ...}
                    │
                    ▼
          pem.Encode / EncodeToMemory
                    │
                    ▼
       -----BEGIN SOMETHING-----
             Base64 data
       -----END SOMETHING-----


                  READING
                    │
                    ▼
       -----BEGIN SOMETHING-----
             Base64 data
       -----END SOMETHING-----
                    │
                    ▼
              pem.Decode()
                    │
                    ▼
                pem.Block
                    │
                    ▼
              block.Bytes
                    │
                    ▼
          x509 / crypto parser
                    │
                    ▼
          Go cryptographic object
```

The key distinction is:

> **`encoding/pem` answers "How do I package these bytes as PEM text?"**

while packages such as **`crypto/x509` answer "What do these cryptographic bytes actually represent?"**

This distinction is especially useful when learning related Go cryptography packages such as `crypto/rsa`, `crypto/ecdsa`, `crypto/ed25519`, and `crypto/x509`.

---

# 15. Thought-provoking question

Suppose an application receives a PEM block labeled:

```text
-----BEGIN PUBLIC KEY-----
...
-----END PUBLIC KEY-----
```

**Should your application trust the key simply because the PEM label says `"PUBLIC KEY"`? Why or why not—and what additional validation would you perform before using that key for authentication or signature verification?**

---

## Official documentation

Go package documentation:

https://pkg.go.dev/encoding/pem

Related package:

https://pkg.go.dev/crypto/x509
