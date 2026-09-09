# Go `crypto/ed25519` Package — Detailed Guide

> **Note:** The Go standard-library package is `crypto/ed25519`.

## 1. What is `crypto/ed25519`?

`crypto/ed25519` provides an implementation of **Ed25519**, a public-key digital signature algorithm.

Its primary purpose is to provide:

- Authentication
- Message/data integrity
- Digital signatures

It is **not an encryption package**.

A simplified flow is:

```text
Private Key
    │
    │ sign(message)
    ▼
Signature
    │
    │ verify(public key + message)
    ▼
Valid / Invalid
```

A sender signs a message with a private key. A receiver can use the corresponding public key to verify that the message was signed by the private-key holder and was not modified.

Ed25519 is commonly used for digital signatures, authentication protocols, software/package signing, secure tokens, and cryptographic identities.

### Key sizes

| Item | Size |
|---|---:|
| Seed | 32 bytes |
| Public key | 32 bytes |
| Private key | 64 bytes |
| Signature | 64 bytes |

---

# 2. Importing the package

```go
import "crypto/ed25519"
```

It is part of Go's standard library, so no third-party dependency is required.

---

# 3. Basic mental model

```text
                  GenerateKey()
                       │
             ┌─────────┴─────────┐
             ▼                   ▼
        Public Key           Private Key
             │                   │
             │                   │
             │              Sign(message)
             │                   │
             │                   ▼
             │              Signature
             │                   │
             └──────────┬────────┘
                        ▼
                 Verify(...)
                        │
                  true / false
```

The **private key signs**.

The **public key verifies**.

The public key does not allow someone to create valid signatures.

---

# 4. Simple complete example

```go
package main

import (
	"crypto/ed25519"
	"fmt"
	"log"
)

func main() {
	// Generate a new Ed25519 key pair.
	publicKey, privateKey, err := ed25519.GenerateKey(nil)
	if err != nil {
		log.Fatal(err)
	}

	// Message we want to sign.
	message := []byte("Hello, Ed25519!")

	// Sign the message with the private key.
	signature := ed25519.Sign(privateKey, message)

	fmt.Printf("Public key: %x\n", publicKey)
	fmt.Printf("Signature:  %x\n", signature)

	// Verify the signature.
	valid := ed25519.Verify(publicKey, message, signature)

	fmt.Println("Signature valid:", valid)
}
```

Typical output will contain a 32-byte public key, a 64-byte signature, and:

```text
Signature valid: true
```

If the message is modified:

```go
message = []byte("Hello, Hacker!")
```

and the original signature is reused, verification fails:

```text
false
```

---

# 5. Every function and method in `crypto/ed25519`

The package provides functions for key generation, signing, verification, key construction, and advanced Ed25519 variants. It also provides `Options`, `PrivateKey`, and `PublicKey` types with associated methods.

## 5.1 `GenerateKey`

### Signature

```go
func GenerateKey(rand io.Reader) (PublicKey, PrivateKey, error)
```

### Purpose

Generates a new Ed25519 public/private key pair.

```go
publicKey, privateKey, err := ed25519.GenerateKey(nil)
if err != nil {
	log.Fatal(err)
}
```

The result is:

```text
PublicKey  → 32 bytes
PrivateKey → 64 bytes
```

If `rand` is `nil`, Go uses `crypto/rand.Reader`, which provides cryptographically secure randomness.

### Why does it accept `io.Reader`?

This makes the function flexible and testable.

For production:

```go
ed25519.GenerateKey(nil)
```

uses secure randomness.

Conceptually:

```text
crypto/rand.Reader
       │
       ▼
random seed
       │
       ▼
Ed25519 key generation
       │
       ├── Public Key
       └── Private Key
```

### Important

Do not generate production cryptographic keys using `math/rand`.

Use cryptographically secure randomness.

---

# 5.2 `Sign`

### Signature

```go
func Sign(privateKey PrivateKey, message []byte) []byte
```

### Purpose

Creates an Ed25519 signature for a message.

```go
message := []byte("Hello")

signature := ed25519.Sign(privateKey, message)
```

The signature is 64 bytes.

Conceptually:

```text
Private Key
     +
 Message
     │
     ▼
Ed25519 algorithm
     │
     ▼
64-byte Signature
```

### Important warning

`Sign` panics if the private key does not have the expected `PrivateKeySize` of 64 bytes.

Do not blindly convert untrusted input into a `PrivateKey` and call `Sign`.

---

# 5.3 `Verify`

### Signature

```go
func Verify(publicKey PublicKey, message, sig []byte) bool
```

### Purpose

Checks whether a signature is valid for a message and public key.

```go
valid := ed25519.Verify(
	publicKey,
	message,
	signature,
)

if valid {
	fmt.Println("Valid signature")
} else {
	fmt.Println("Invalid signature")
}
```

It returns:

```text
true
```

or:

```text
false
```

### Verification model

```text
             Public Key
                  │
                  │
Message ──────────┤
                  │
Signature ────────┤
                  ▼
              Verify()
                  │
             ┌────┴────┐
             ▼         ▼
           true       false
```

### Tampering example

Original:

```go
message := []byte("Pay Alice $100")
signature := ed25519.Sign(privateKey, message)
```

Verification:

```go
ed25519.Verify(publicKey, message, signature)
```

returns:

```text
true
```

But if the message becomes:

```go
modified := []byte("Pay Alice $900")
```

then:

```go
ed25519.Verify(publicKey, modified, signature)
```

returns:

```text
false
```

`Verify` panics if the public key is not exactly `PublicKeySize` bytes.

---

# 5.4 `VerifyWithOptions`

### Signature

```go
func VerifyWithOptions(
	publicKey PublicKey,
	message []byte,
	sig []byte,
	opts *Options,
) error
```

Unlike:

```go
Verify(...)
```

which returns a `bool`, `VerifyWithOptions` returns an `error`.

A `nil` error means verification succeeded.

It supports advanced Ed25519 variants such as:

- Ed25519ctx
- Ed25519ph

Example:

```go
opts := &ed25519.Options{
	Context: "my-application",
}

err := ed25519.VerifyWithOptions(
	publicKey,
	message,
	signature,
	opts,
)

if err != nil {
	fmt.Println("Invalid signature")
}
```

### Why use a context?

Suppose one system uses signatures for:

```text
Payment authorization
File approval
Login challenge
API request
Document approval
```

A cryptographic context can provide **domain separation**, so signatures created for one purpose are not casually reused for another purpose.

Conceptually:

```text
"PAYMENT" + message
```

is separated from:

```text
"LOGIN" + message
```

when the protocol is designed to use different contexts.

The context is limited in size and should be deliberately designed as part of the protocol rather than added randomly.

---

# 5.5 `Options`

`Options` controls advanced Ed25519 behavior.

Its fields are:

```go
type Options struct {
	Hash    crypto.Hash
	Context string
}
```

## `Hash`

For ordinary Ed25519:

```go
Hash: crypto.Hash(0)
```

For Ed25519ph:

```go
Hash: crypto.SHA512
```

The latter is used when the protocol expects the message to have been prehashed with SHA-512.

## `Context`

A context provides domain separation.

Example:

```go
opts := &ed25519.Options{
	Context: "my-app-v1",
}
```

Use a context only when your protocol deliberately defines one and all parties agree on it.

---

# 5.6 `Options.HashFunc()`

### Signature

```go
func (o *Options) HashFunc() crypto.Hash
```

It returns the hash associated with the options.

Example:

```go
opts := &ed25519.Options{
	Hash: crypto.SHA512,
}

fmt.Println(opts.HashFunc())
```

This method allows `Options` to participate in Go's cryptographic signing interfaces.

---

# 5.7 `PrivateKey`

The type is:

```go
type PrivateKey []byte
```

It represents an Ed25519 private key and implements `crypto.Signer`.

A normal generated private key is 64 bytes.

Conceptually:

```text
PrivateKey
┌────────────────────────┬────────────────────────┐
│       32-byte seed     │   32-byte public key   │
└────────────────────────┴────────────────────────┘
```

The 32-byte seed is the secret seed representation, while Go's `PrivateKey` contains the representation used for signing and key handling.

---

# 5.8 `NewKeyFromSeed`

### Signature

```go
func NewKeyFromSeed(seed []byte) PrivateKey
```

Creates an Ed25519 private key from a 32-byte seed.

Example:

```go
seed := make([]byte, ed25519.SeedSize)

privateKey := ed25519.NewKeyFromSeed(seed)
```

The seed must be exactly 32 bytes. Otherwise the function panics.

### Why is this useful?

It is useful when you need deterministic reconstruction of a key from its seed.

Conceptually:

```text
32-byte seed
     │
     ▼
NewKeyFromSeed
     │
     ▼
same Ed25519 private key
     │
     ▼
same public key
```

A seed is highly sensitive secret material. Anyone who obtains it can reconstruct the private key.

---

# 5.9 `PrivateKey.Equal`

### Signature

```go
func (priv PrivateKey) Equal(x crypto.PrivateKey) bool
```

Compares the private key with another private-key value.

Example:

```go
same := privateKey1.Equal(privateKey2)
```

This is useful when determining whether two private-key values represent the same key.

---

# 5.10 `PrivateKey.Public`

### Signature

```go
func (priv PrivateKey) Public() crypto.PublicKey
```

Returns the public key corresponding to the private key.

Example:

```go
publicKey := privateKey.Public()
```

This is particularly useful when working with Go's generic `crypto.Signer` interface.

Conceptually:

```text
PrivateKey
    │
    │ Public()
    ▼
PublicKey
```

---

# 5.11 `PrivateKey.Seed`

### Signature

```go
func (priv PrivateKey) Seed() []byte
```

Extracts the 32-byte seed from a private key.

Example:

```go
seed := privateKey.Seed()

fmt.Println(len(seed))
```

Output:

```text
32
```

The returned seed is secret material and must be protected like the private key.

---

# 5.12 `PrivateKey.Sign`

`PrivateKey` implements Go's `crypto.Signer` interface.

The method is:

```go
func (priv PrivateKey) Sign(
	rand io.Reader,
	message []byte,
	opts crypto.SignerOpts,
) ([]byte, error)
```

This allows an Ed25519 private key to work with APIs expecting a generic `crypto.Signer`.

Conceptually:

```go
var signer crypto.Signer = privateKey
```

Then code can use:

```go
signer.Sign(...)
```

instead of depending directly on Ed25519.

This is valuable for reusable cryptographic infrastructure.

---

# 5.13 `PublicKey`

The type is:

```go
type PublicKey []byte
```

An Ed25519 public key is 32 bytes.

It is safe to distribute publicly.

```text
Server
  │
  ├── Private key → KEEP SECRET
  │
  └── Public key → distribute
```

Anyone can possess the public key without being able to generate signatures.

---

# 5.14 `PublicKey.Equal`

### Signature

```go
func (pub PublicKey) Equal(x crypto.PublicKey) bool
```

Compares a public key with another public-key value.

Example:

```go
if publicKey1.Equal(publicKey2) {
	fmt.Println("Same public key")
}
```

---

# 6. Package constants

The most important constants are:

```go
ed25519.PublicKeySize
ed25519.PrivateKeySize
ed25519.SignatureSize
ed25519.SeedSize
```

Their values are:

```text
PublicKeySize   = 32
PrivateKeySize  = 64
SignatureSize   = 64
SeedSize        = 32
```

Example:

```go
fmt.Println(ed25519.PublicKeySize)
fmt.Println(ed25519.PrivateKeySize)
fmt.Println(ed25519.SignatureSize)
fmt.Println(ed25519.SeedSize)
```

Output:

```text
32
64
64
32
```

---

# 7. `Sign` vs `Verify`

This distinction is fundamental.

### Signing

Only the private-key holder can do this:

```go
signature := ed25519.Sign(privateKey, message)
```

### Verification

Anyone with the public key can do this:

```go
valid := ed25519.Verify(
	publicKey,
	message,
	signature,
)
```

So:

```text
                 PRIVATE KEY
                      │
                      ▼
                   SIGN
                      │
                      ▼
                   SIGNATURE
                      │
          ┌───────────┴───────────┐
          │                       │
       MESSAGE                PUBLIC KEY
          │                       │
          └───────────┬───────────┘
                      ▼
                   VERIFY
                      │
                ┌─────┴─────┐
                ▼           ▼
              true        false
```

---

# 8. What Ed25519 does NOT do

A common misconception is:

> "Ed25519 encrypts messages."

No.

Ed25519 is primarily a **digital signature algorithm**.

It provides:

```text
Authentication
Integrity
Digital signatures
```

It does not provide ordinary message encryption.

If your goal is:

> "Only the recipient should be able to read this message."

you need encryption/key agreement, not Ed25519 signatures.

If your goal is:

> "The recipient should be able to verify that I signed this message."

Ed25519 is appropriate.

---

# 9. Three common beginner mistakes

## Mistake 1: Thinking Ed25519 encrypts messages

Incorrect:

```text
Ed25519 → encryption
```

Correct:

```text
Ed25519 → digital signatures
```

### Avoid it

Distinguish these two goals:

```text
Confidentiality → encryption
Authenticity/integrity → digital signature
```

---

## Mistake 2: Confusing the seed with the 64-byte private key

Beginners often assume:

```text
seed = private key
```

The relationship is more subtle.

Go's:

```go
ed25519.PrivateKey
```

is 64 bytes.

The RFC-style seed is:

```text
32 bytes
```

and:

```go
ed25519.NewKeyFromSeed(seed)
```

constructs the Go private-key representation from that seed.

### Remember

```text
Seed       → 32 bytes
PrivateKey → 64 bytes
PublicKey  → 32 bytes
Signature  → 64 bytes
```

---

## Mistake 3: Not protecting the private key

A signature algorithm is only as secure as the private-key handling around it.

Avoid putting private keys in:

```text
source code
Git repositories
logs
plain configuration files
client-side JavaScript
```

If an attacker obtains your private key, they can create valid signatures.

### Better approach

```text
Application
     │
     ▼
Secure key storage
     │
     ▼
Private key
     │
     ▼
Sign
```

Use an appropriate secret-management or key-management solution for production systems.

---

# 10. Two real-world applications

## Application 1: Software/package signing

A software vendor can sign software artifacts or metadata with an Ed25519 private key.

Users have the vendor's public key and can verify the signature.

```text
Vendor
  │
  │ private key
  ▼
Sign software
  │
  ▼
Signature
  │
  ▼
Internet
  │
  ▼
User
  │
  │ public key
  ▼
Verify
```

If an attacker modifies the artifact, verification fails.

This helps establish authenticity and integrity of distributed software.

---

## Application 2: Authentication / signed API requests

A device can possess an Ed25519 private key and sign API requests.

For example:

```text
POST /api/payment

amount=1000
timestamp=...
nonce=...
```

The device signs a canonical representation of the request.

The server knows the corresponding public key.

```text
Device
   │
   ├── request
   ├── timestamp
   ├── nonce
   │
   └── Ed25519 signature
             │
             ▼
           Server
             │
             ▼
        Verify signature
             │
        ┌────┴────┐
        ▼         ▼
      valid     invalid
```

A real protocol also needs protections such as canonicalization, replay prevention, key rotation, and secure private-key storage.

---

# 11. Three progressively challenging exercises

## Exercise 1 — Basic signing and verification

Write a Go program that:

1. Generates an Ed25519 key pair.
2. Creates a message.
3. Signs the message.
4. Prints the signature.
5. Verifies the signature.
6. Prints whether verification succeeded.
7. Changes one character of the message.
8. Verifies the original signature against the modified message.

Your program should demonstrate why changing even a small part of signed data causes verification to fail.

**Do not use a solution; implement it yourself.**

---

## Exercise 2 — Signed API request

Build a small Go program that simulates a client and server.

The client should:

1. Generate an Ed25519 key pair.
2. Construct a request containing:
   - HTTP method
   - path
   - timestamp
   - request body
3. Create a canonical representation of the request.
4. Sign the canonical request.
5. Send the request data and signature to a simulated server.

The server should:

1. Receive the request.
2. Reconstruct the exact canonical message.
3. Verify the signature.
4. Reject the request if any request component has been modified.

Then deliberately modify the request body and demonstrate that verification fails.

---

## Exercise 3 — Key persistence and protocol design

Build a small command-line application that implements an Ed25519-based identity system.

Support commands conceptually similar to:

```text
generate-key
show-public-key
sign
verify
```

Requirements:

1. Generate an Ed25519 key pair.
2. Persist the private-key material securely enough for your exercise.
3. Export the public key.
4. Sign arbitrary input data.
5. Verify signatures.
6. Use `NewKeyFromSeed` to demonstrate reconstruction from a 32-byte seed.
7. Demonstrate that reconstructing the private key from the same seed produces the same public key.
8. Add a protocol context using `ed25519.Options`.
9. Demonstrate that a signature generated for one context cannot simply be treated as a signature for another context.
10. Think carefully about what information your file format needs to store and what information must never be logged.

**Do not simply serialize everything blindly; design the format yourself.**

---

# 12. Useful learning map

If you're learning Go cryptography systematically, understand `ed25519` in this order:

```text
1. Public-key cryptography
        ↓
2. Digital signatures
        ↓
3. Ed25519 concept
        ↓
4. GenerateKey
        ↓
5. Sign
        ↓
6. Verify
        ↓
7. PublicKey / PrivateKey
        ↓
8. Seed vs PrivateKey
        ↓
9. NewKeyFromSeed
        ↓
10. crypto.Signer
        ↓
11. VerifyWithOptions
        ↓
12. Ed25519ctx / Ed25519ph
        ↓
13. Key storage & rotation
        ↓
14. Protocol design
```

The most important practical lesson is:

> **Using `ed25519.Sign` and `ed25519.Verify` is easy; designing a secure protocol around signatures is the harder part.**

---

# 13. Thought-provoking question

Imagine you use **one Ed25519 private key for your entire application** to sign:

- Login tokens
- Payment requests
- Software updates
- Administrative commands

The cryptography itself is secure.

**What could go wrong at the application/protocol level, and how could Ed25519 contexts, separate keys, key rotation, and domain separation help prevent a signature created for one purpose from being misused for another?**

This is the point where understanding `crypto/ed25519` moves from *"I know how to call the functions"* to *"I understand how to design systems with digital signatures."*
