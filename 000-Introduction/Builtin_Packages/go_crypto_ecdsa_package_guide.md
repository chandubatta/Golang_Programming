# Go `crypto/ecdsa` Package

## 1. What is `crypto/ecdsa`?

The Go `crypto/ecdsa` package implements **ECDSA — Elliptic Curve Digital Signature Algorithm**.

Its primary purpose is **digital signatures**.

```text
Message
   │
   ▼
Hash (SHA-256, SHA-384, ...)
   │
   ▼
ECDSA Sign with Private Key
   │
   ▼
Signature
```

Another party can then use your **public key**:

```text
Message
   │
   ▼
Hash
   │
   ├──────────────┐
   ▼              ▼
Signature      Public Key
   │              │
   └──────┬───────┘
          ▼
       Verify
          │
      true / false
```

ECDSA provides **authentication and integrity**, but it does **not encrypt the message**.

The standard-library implementation follows FIPS 186-5. For the standard NIST curves (`P224`, `P256`, `P384`, `P521`), operations involving private keys use constant-time algorithms.

### Common uses

ECDSA is commonly used for:

- TLS/HTTPS certificates
- digitally signed software
- authentication tokens
- API request signatures
- secure document signing
- certificate infrastructure
- blockchain systems that specifically use ECDSA

---

# 2. Important Concepts Before Using ECDSA

You should understand four things.

## Private key

The private key is secret.

```text
Private Key
    │
    └── used to SIGN
```

## Public key

The public key can be shared.

```text
Public Key
    │
    └── used to VERIFY
```

## Hash

ECDSA normally signs a **hash of the message**, not the raw message.

```text
"Hello"
   ↓
SHA-256
   ↓
32-byte digest
   ↓
ECDSA
```

## Signature

An ECDSA signature contains two mathematical values:

```text
(r, s)
```

Go can expose these directly with `Sign`/`Verify`, or encode them as ASN.1 DER using `SignASN1`/`VerifyASN1`.

---

# 3. Simple Example

```go
package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

func main() {
	// 1. Generate an ECDSA private key.
	privateKey, err := ecdsa.GenerateKey(
		elliptic.P256(),
		rand.Reader,
	)
	if err != nil {
		panic(err)
	}

	// 2. Message we want to sign.
	message := []byte("Hello, ECDSA!")

	// 3. Hash the message.
	hash := sha256.Sum256(message)

	// 4. Sign the hash.
	signature, err := ecdsa.SignASN1(
		rand.Reader,
		privateKey,
		hash[:],
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Signature: %x\n", signature)

	// 5. Verify using the public key.
	valid := ecdsa.VerifyASN1(
		&privateKey.PublicKey,
		hash[:],
		signature,
	)

	fmt.Println("Signature valid:", valid)
}
```

Output will look conceptually like:

```text
Signature: 3045022100...
Signature valid: true
```

The signature bytes will normally be different each time because ECDSA signing is randomized.

---

# 4. Every Important Function in `crypto/ecdsa`

The package provides package-level signing/verification functions and methods on its key types.

---

## Function 1: `GenerateKey`

```go
func GenerateKey(
    c elliptic.Curve,
    r io.Reader,
) (*PrivateKey, error)
```

### Purpose

Generates a new ECDSA private/public key pair for a specified elliptic curve.

Example:

```go
privateKey, err := ecdsa.GenerateKey(
	elliptic.P256(),
	rand.Reader,
)
```

You receive:

```text
PrivateKey
   │
   ├── D       → private scalar
   │
   └── PublicKey
          ├── Curve
          ├── X
          └── Y
```

### Common curves

```go
elliptic.P224()
elliptic.P256()
elliptic.P384()
elliptic.P521()
```

For most applications, **P-256 is a common practical choice**, unless a protocol specifies another curve.

### Important

Never generate production private keys using:

```go
math/rand
```

Use cryptographically secure randomness.

---

# Function 2: `Sign`

```go
func Sign(
    rand io.Reader,
    priv *PrivateKey,
    hash []byte,
) (r, s *big.Int, err error)
```

This performs ECDSA signing but returns the signature as two integers:

```text
(r, s)
```

Example:

```go
r, s, err := ecdsa.Sign(
	rand.Reader,
	privateKey,
	hash[:],
)
```

You could print:

```go
fmt.Println("r:", r)
fmt.Println("s:", s)
```

### Why `r` and `s`?

ECDSA mathematically represents its signature as:

```text
Signature = (r, s)
```

If you want to transmit the signature over a network, you need to serialize those integers.

That's why most applications should prefer:

```go
ecdsa.SignASN1(...)
```

rather than manually handling `r` and `s`.

---

# Function 3: `SignASN1`

```go
func SignASN1(
    r io.Reader,
    priv *PrivateKey,
    hash []byte,
) ([]byte, error)
```

This is generally the more convenient signing API.

Instead of:

```text
r
s
```

you receive an encoded signature:

```text
ASN.1 DER signature
```

Example:

```go
signature, err := ecdsa.SignASN1(
	rand.Reader,
	privateKey,
	hash[:],
)
```

Conceptually:

```text
(r, s)
   ↓
ASN.1 DER encoding
   ↓
[]byte
```

This is useful when the signature needs to be transmitted or stored.

---

# Function 4: `Verify`

```go
func Verify(
    pub *PublicKey,
    hash []byte,
    r, s *big.Int,
) bool
```

This verifies a signature represented by `r` and `s`.

Example:

```go
valid := ecdsa.Verify(
	&privateKey.PublicKey,
	hash[:],
	r,
	s,
)

fmt.Println(valid)
```

Result:

```text
true
```

if the signature is valid.

If the message has been modified, the hash changes and verification fails:

```text
false
```

Verification inputs are public; don't design systems that assume verification hides those values.

---

# Function 5: `VerifyASN1`

```go
func VerifyASN1(
    pub *PublicKey,
    hash, sig []byte,
) bool
```

This is the counterpart to `SignASN1`.

You give it:

```text
Public Key
Hash
ASN.1 Signature
```

and it returns:

```text
true
```

or:

```text
false
```

Example:

```go
valid := ecdsa.VerifyASN1(
	&privateKey.PublicKey,
	hash[:],
	signature,
)
```

Typical application flow:

```text
SignASN1
    ↓
[]byte signature
    ↓
Network / Database / File
    ↓
VerifyASN1
```

This is generally preferable to manually transporting `r` and `s`.

---

# 5. `PrivateKey`

The package provides a private-key type containing a public key and private scalar:

```go
type PrivateKey struct {
    PublicKey
    D *big.Int
}
```

Conceptually:

```text
PrivateKey
│
├── PublicKey
│   ├── Curve
│   ├── X
│   └── Y
│
└── D
    └── private scalar
```

`D` is the private scalar, so it must be protected carefully.

Avoid directly manipulating the internal private-key value when dedicated APIs are available.

---

# 6. `PrivateKey.Bytes()`

```go
func (priv *PrivateKey) Bytes() ([]byte, error)
```

Added in Go 1.25.

It converts the private key into a fixed-length big-endian raw representation according to SEC 1.

Example:

```go
raw, err := privateKey.Bytes()
if err != nil {
	panic(err)
}

fmt.Printf("%x\n", raw)
```

Conceptually:

```text
PrivateKey
    ↓
Bytes()
    ↓
raw private-key bytes
```

This is useful when a protocol specifically requires the raw SEC 1 representation.

For normal application key storage, formats such as PKCS#8 are often more appropriate.

---

# 7. `ParseRawPrivateKey`

```go
func ParseRawPrivateKey(
    curve elliptic.Curve,
    data []byte,
) (*PrivateKey, error)
```

This does the reverse of `Bytes()`.

```text
raw bytes
    ↓
ParseRawPrivateKey()
    ↓
PrivateKey
```

Example:

```go
privateKey, err := ecdsa.ParseRawPrivateKey(
	elliptic.P256(),
	raw,
)
```

The raw value must be valid for the selected NIST curve and cannot be zero or an unreduced value.

---

# 8. `PrivateKey.Public()`

```go
func (priv *PrivateKey) Public() crypto.PublicKey
```

Returns the public key associated with the private key.

Example:

```go
publicKey := privateKey.Public()
```

This is useful because `crypto.PublicKey` is a generic type used by Go's crypto APIs.

You can also directly access:

```go
privateKey.PublicKey
```

when you specifically need an `*ecdsa.PublicKey`.

---

# 9. `PrivateKey.Equal()`

```go
func (priv *PrivateKey) Equal(
    x crypto.PrivateKey,
) bool
```

Checks whether two private keys represent the same key.

Conceptually:

```text
PrivateKey A
     │
     ├── compare
     │
PrivateKey B
     │
     ▼
 true / false
```

This is preferable to manually comparing internal cryptographic fields.

---

# 10. `PrivateKey.Sign()`

```go
func (priv *PrivateKey) Sign(
    random io.Reader,
    digest []byte,
    opts crypto.SignerOpts,
) ([]byte, error)
```

This method makes `ecdsa.PrivateKey` implement Go's `crypto.Signer` interface.

Example:

```go
signature, err := privateKey.Sign(
	rand.Reader,
	hash[:],
	nil,
)
```

Why is this useful?

It allows code to work with different signing implementations through a common interface.

For example:

```text
Application
     │
     ▼
crypto.Signer
     │
 ┌───┴────┐
 ▼        ▼
ECDSA   Hardware
        Security
        Module
```

This abstraction is useful in systems where private keys might eventually live in hardware rather than application memory.

---

# 11. `PrivateKey.ECDH()`

```go
func (priv *PrivateKey) ECDH() (*ecdh.PrivateKey, error)
```

Added in Go 1.20.

This converts a compatible ECDSA private key into an `ecdh.PrivateKey`.

ECDSA and ECDH are different cryptographic operations:

```text
ECDSA
  ↓
Digital signatures

ECDH
  ↓
Shared-secret/key agreement
```

So you should not think:

> "ECDSA encrypts data."

It doesn't.

For key agreement, Go provides the dedicated `crypto/ecdh` package. The ECDSA package provides this conversion for compatible curves.

---

# 12. `PublicKey`

The public key contains:

```go
type PublicKey struct {
    elliptic.Curve
    X, Y *big.Int
}
```

Conceptually:

```text
PublicKey
│
├── Curve
├── X
└── Y
```

The `(X, Y)` values represent a point on the elliptic curve.

---

# 13. `PublicKey.Bytes()`

```go
func (pub *PublicKey) Bytes() ([]byte, error)
```

Added in Go 1.25.

It serializes the public key into the raw uncompressed format defined by SEC 1.

Conceptually:

```text
Public Key
   │
   ▼
Bytes()
   │
   ▼
04 || X || Y
```

The `04` indicates an uncompressed elliptic-curve point.

This is useful when a protocol explicitly requires raw SEC 1 encoding.

---

# 14. `ParseUncompressedPublicKey`

```go
func ParseUncompressedPublicKey(
    curve elliptic.Curve,
    data []byte,
) (*PublicKey, error)
```

Converts raw uncompressed public-key bytes into an ECDSA public key.

Conceptually:

```text
04 || X || Y
      │
      ▼
ParseUncompressedPublicKey()
      │
      ▼
PublicKey
```

This is particularly useful when interoperating with another cryptographic system that provides an SEC 1 uncompressed public key.

---

# 15. `PublicKey.Equal()`

```go
func (pub *PublicKey) Equal(
    x crypto.PublicKey,
) bool
```

Determines whether two public keys represent the same key.

Example:

```go
if publicKey.Equal(otherPublicKey) {
	fmt.Println("Same public key")
}
```

It compares the key values and curve appropriately.

---

# 16. `PublicKey.ECDH()`

```go
func (pub *PublicKey) ECDH() (*ecdh.PublicKey, error)
```

Converts a compatible ECDSA public key to an ECDH public key.

Conceptually:

```text
ECDSA PublicKey
       │
       ▼
    ECDH()
       │
       ▼
ECDH PublicKey
```

This is useful when an application has an EC key and needs to use it for an ECDH operation supported by `crypto/ecdh`.

---

# 17. Function Summary

| Function / Method | Purpose |
|---|---|
| `GenerateKey()` | Generate ECDSA key pair |
| `Sign()` | Sign hash → `(r, s)` |
| `SignASN1()` | Sign hash → DER/ASN.1 bytes |
| `Verify()` | Verify `(r, s)` |
| `VerifyASN1()` | Verify DER/ASN.1 signature |
| `PrivateKey.Bytes()` | Private key → raw bytes |
| `ParseRawPrivateKey()` | Raw bytes → private key |
| `PrivateKey.Public()` | Get public key |
| `PrivateKey.Equal()` | Compare private keys |
| `PrivateKey.Sign()` | Implement `crypto.Signer` |
| `PrivateKey.ECDH()` | Convert to ECDH private key |
| `PublicKey.Bytes()` | Public key → raw bytes |
| `ParseUncompressedPublicKey()` | Raw public key → public key |
| `PublicKey.Equal()` | Compare public keys |
| `PublicKey.ECDH()` | Convert to ECDH public key |

---

# 18. The Most Important Workflow

For a beginner, remember this:

```text
                 KEY GENERATION
                      │
                      ▼
             ecdsa.GenerateKey()
                      │
             ┌────────┴────────┐
             ▼                 ▼
        Private Key        Public Key
             │                 │
             │                 │
             ▼                 │
          SIGN                 │
             │                 │
Message → SHA-256              │
             │                 │
             ▼                 │
       SignASN1()              │
             │                 │
             ▼                 │
        Signature              │
             │                 │
             └───────┬─────────┘
                     ▼
                  VERIFY
                     │
                     ▼
               VerifyASN1()
                     │
                ┌────┴────┐
                ▼         ▼
              true       false
```

This is the core of ECDSA.

---

# 19. Three Common Beginner Mistakes

## Mistake 1: Thinking ECDSA encrypts data

Wrong:

```text
ECDSA → encryption
```

Correct:

```text
ECDSA → digital signatures
```

If you need encryption or key agreement, use an appropriate cryptographic primitive such as `crypto/ecdh` for ECDH key agreement and then an authenticated symmetric cipher for data encryption.

---

## Mistake 2: Signing the raw message

A beginner might conceptually try:

```go
ecdsa.SignASN1(rand.Reader, privateKey, message)
```

ECDSA's API expects a **hash/digest**, not arbitrary plaintext. The usual pattern is:

```go
hash := sha256.Sum256(message)

signature, err := ecdsa.SignASN1(
	rand.Reader,
	privateKey,
	hash[:],
)
```

---

## Mistake 3: Treating the private key like ordinary data

Don't casually print:

```go
fmt.Println(privateKey.D)
```

or store the private scalar in an ordinary database column without proper key-management protections.

The private key is the credential that allows somebody to impersonate the signer.

Think:

```text
Public Key  → safe to distribute

Private Key → MUST remain secret
```

Also avoid directly modifying `D`; use dedicated key APIs.

---

# 20. Real-World Application #1 — TLS/HTTPS

ECDSA is widely used in public-key infrastructure.

For example:

```text
Web Server
    │
    ├── ECDSA Private Key
    │
    └── ECDSA Certificate/Public Key
```

During TLS authentication, cryptographic signatures help the server prove possession of the private key associated with its certificate.

This is one reason elliptic-curve cryptography is important in modern HTTPS infrastructure.

---

# 21. Real-World Application #2 — Signed Documents/API Requests

Imagine an API request:

```text
POST /transfer

amount=1000
account=12345
```

The client can hash the request:

```text
request
   ↓
SHA-256
   ↓
digest
   ↓
ECDSA Sign
   ↓
signature
```

The server has the client's public key:

```text
public key + request + signature
                 │
                 ▼
             VerifyASN1
                 │
            ┌────┴────┐
            ▼         ▼
          valid     invalid
```

The server can therefore determine whether the request was signed by whoever controls the corresponding private key and whether the signed content was altered.

---

# 22. Exercises

## Exercise 1 — Beginner: Sign and Verify

Create a Go program that:

1. Generates an ECDSA P-256 key pair.
2. Creates a message such as `"I am learning ECDSA"`.
3. Calculates its SHA-256 hash.
4. Signs the hash using `SignASN1`.
5. Verifies the signature using `VerifyASN1`.
6. Prints whether verification succeeded.

Then modify the message **after signing** and attempt verification again.

Your program should demonstrate that the modified message fails verification.

---

## Exercise 2 — Intermediate: Signature Storage and Key Serialization

Build a small program that:

1. Generates an ECDSA P-256 private key.
2. Converts the private key to raw bytes using `PrivateKey.Bytes()`.
3. Reconstructs the private key using `ParseRawPrivateKey()`.
4. Generates the corresponding public key.
5. Signs a message using the reconstructed key.
6. Serializes the public key using `PublicKey.Bytes()`.
7. Parses that public key using `ParseUncompressedPublicKey()`.
8. Uses the parsed public key to verify the signature.

Your goal is to demonstrate that a key can travel through a serialization/deserialization process without changing its cryptographic identity.

---

## Exercise 3 — Advanced: Build a Signed API Request System

Create a miniature client/server authentication system.

### Client

The client should:

1. Generate an ECDSA key pair.
2. Construct a request containing:
   - user ID
   - timestamp
   - HTTP-like method
   - request body
3. Create a canonical representation of the request.
4. Hash the canonical request.
5. Sign the hash using ECDSA.
6. Send the request and signature to the server.

### Server

The server should:

1. Store the client's public key.
2. Receive the request and signature.
3. Reconstruct the exact canonical request.
4. Hash it.
5. Verify the signature.
6. Reject the request if:
   - the signature is invalid,
   - the request has been modified,
   - the signature belongs to a different key,
   - or the timestamp is outside an allowed window.

Then investigate what happens when an attacker modifies **one byte** of the request after it has been signed.

---

# 23. Thought-Provoking Question

Suppose you build an API where every request is authenticated with ECDSA.

You have successfully solved:

> **"Was this request signed by the holder of the private key?"**

But now consider this:

> **What prevents an attacker from copying a completely valid signed request and sending that exact same request to your server 1,000 times?**

Think about how **digital signatures, timestamps, nonces, request IDs, replay protection, and key management** need to work together.

That question is where ECDSA moves from being just a cryptography API you can call into a component of a real security system.
