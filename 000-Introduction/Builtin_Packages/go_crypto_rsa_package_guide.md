# Go `crypto/rsa` Package — Detailed Guide

The Go `crypto/rsa` package implements **RSA public-key cryptography**, including encryption/decryption and digital signatures. It follows the RSA mechanisms specified by **PKCS #1 / RFC 8017**.

In modern applications, **RSA-OAEP** is preferred for encryption and **RSA-PSS** is preferred for signatures; older PKCS #1 v1.5 encryption APIs are deprecated.

Official documentation: https://pkg.go.dev/crypto/rsa

---

# 1. What is the `crypto/rsa` package?

RSA is an **asymmetric/public-key cryptographic algorithm**.

It uses two mathematically related keys:

```text
             RSA Key Pair
          ┌───────────────┐
          │               │
     Public Key       Private Key
          │               │
          ▼               ▼
      Encrypt          Decrypt
      Verify           Sign
```

The important idea is:

- **Public key** → can be shared with everybody.
- **Private key** → must remain secret.
- Data encrypted with the public key can be decrypted with the private key.
- Data signed with the private key can be verified with the public key.

The package supports both:

## Encryption

```text
Message
   ↓
RSA-OAEP
   ↓
Ciphertext
   ↓
RSA-OAEP
   ↓
Message
```

## Digital signatures

```text
Message
   ↓
SHA-256
   ↓
Digest
   ↓
RSA-PSS + Private Key
   ↓
Signature
   ↓
RSA-PSS + Public Key
   ↓
Valid / Invalid
```

---

# 2. When is `crypto/rsa` commonly used?

Typical uses include:

- Digital signatures
- Verifying signed data
- Encrypting small secrets
- Protecting symmetric encryption keys
- Authentication protocols
- Certificates and PKI
- Secure exchange of AES keys
- JWT/JWS ecosystems that use RSA signatures
- TLS/certificate-related cryptographic operations

One important limitation is that **RSA is not designed for encrypting large amounts of data directly**.

For example, instead of:

```text
10 MB file
    ↓
RSA encryption
```

a real application normally uses:

```text
                 Random AES key
                       │
             ┌─────────┴─────────┐
             ↓                   ↓
       Encrypt file          RSA-OAEP
       using AES-GCM       encrypt AES key
             │                   │
             └─────────┬─────────┘
                       ↓
                Store/transmit
```

This is called **hybrid encryption**.

---

# 3. Basic RSA example

Let's start with a simple example that:

1. Generates an RSA private key.
2. Gets the public key.
3. Encrypts a message using RSA-OAEP.
4. Decrypts it using the private key.

```go
package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

func main() {
	// Generate a 2048-bit RSA private key.
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	// Get the corresponding public key.
	publicKey := &privateKey.PublicKey

	message := []byte("Hello RSA!")

	// Encrypt using the public key.
	ciphertext, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		publicKey,
		message,
		nil,
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Encrypted: %x\n", ciphertext)

	// Decrypt using the private key.
	plaintext, err := rsa.DecryptOAEP(
		sha256.New(),
		nil,
		privateKey,
		ciphertext,
		nil,
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("Decrypted:", string(plaintext))
}
```

Output will conceptually look like:

```text
Encrypted: 8a7f...lots of hexadecimal data...
Decrypted: Hello RSA!
```

Notice something important:

```go
rsa.EncryptOAEP(...)
```

uses the **public key**, while:

```go
rsa.DecryptOAEP(...)
```

uses the **private key**.

Also, because OAEP uses randomness, encrypting the same plaintext twice will normally produce different ciphertexts.

---

# 4. RSA key concepts

Before going through the APIs, understand these structures.

## Private key

Conceptually:

```text
PrivateKey
 ├── PublicKey
 │    ├── N
 │    └── E
 │
 ├── D
 ├── Primes
 └── Precomputed
```

The private key contains sensitive mathematical values.

## Public key

Conceptually:

```text
PublicKey
 ├── N
 └── E
```

`N` is the RSA modulus.

`E` is the public exponent.

The public key can safely be distributed.

---

# 5. Every important function and method in `crypto/rsa`

The package contains encryption, decryption, signing, verification, key-generation, and key-management APIs. Some older APIs are deprecated.

Let's go through them systematically.

---

# A. `GenerateKey`

```go
func GenerateKey(random io.Reader, bits int) (*PrivateKey, error)
```

Generates an RSA private key.

Example:

```go
privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
```

### Parameters

### `random`

Traditionally specifies the source of randomness.

Commonly:

```go
rand.Reader
```

from:

```go
crypto/rand
```

### `bits`

Specifies RSA key size.

For example:

```go
2048
```

or:

```go
3072
```

or:

```go
4096
```

For new applications, 2048 bits is a common baseline, while larger keys may be chosen for longer-term requirements.

Example:

```go
privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
if err != nil {
	log.Fatal(err)
}
```

---

# B. `GenerateMultiPrimeKey`

```go
func GenerateMultiPrimeKey(
	random io.Reader,
	nprimes int,
	bits int,
) (*PrivateKey, error)
```

Generates a multi-prime RSA key.

Traditional RSA uses two primes:

```text
p × q
```

Multi-prime RSA can use:

```text
p × q × r
```

or more.

However, this function is **deprecated**.

The Go documentation recommends normal two-prime RSA using:

```go
rsa.GenerateKey(...)
```

for new code.

So for new applications:

```go
rsa.GenerateKey(rand.Reader, 2048)
```

is preferred.

---

# C. `EncryptOAEP`

```go
func EncryptOAEP(
	hash hash.Hash,
	random io.Reader,
	pub *PublicKey,
	msg []byte,
	label []byte,
) ([]byte, error)
```

This is one of the most important functions in the package.

It encrypts data using **RSA-OAEP**.

Example:

```go
ciphertext, err := rsa.EncryptOAEP(
	sha256.New(),
	rand.Reader,
	publicKey,
	message,
	nil,
)
```

### Arguments

### 1. `hash`

Example:

```go
sha256.New()
```

OAEP uses this hash as part of its padding mechanism.

Encryption and decryption must use compatible parameters.

### 2. `random`

```go
rand.Reader
```

Provides cryptographically secure randomness.

### 3. `pub`

The recipient's public key:

```go
publicKey
```

### 4. `msg`

The plaintext:

```go
[]byte("Secret message")
```

### 5. `label`

An optional OAEP label.

Most applications use:

```go
nil
```

or:

```go
[]byte("my-application")
```

The same label must be supplied during decryption.

### Important limitation

RSA can only encrypt messages that fit within the RSA modulus after OAEP padding.

Therefore:

```go
rsa.EncryptOAEP(...)
```

should **not** be used to encrypt arbitrary large files.

Use hybrid encryption instead.

---

# D. `DecryptOAEP`

```go
func DecryptOAEP(
	hash hash.Hash,
	random io.Reader,
	priv *PrivateKey,
	ciphertext []byte,
	label []byte,
) ([]byte, error)
```

Decrypts ciphertext created using RSA-OAEP.

Example:

```go
plaintext, err := rsa.DecryptOAEP(
	sha256.New(),
	nil,
	privateKey,
	ciphertext,
	nil,
)
```

The important relationship is:

```text
EncryptOAEP
    ↓
public key
    ↓
ciphertext
    ↓
private key
    ↓
DecryptOAEP
```

The hash and label need to match the encryption parameters.

---

# E. `EncryptOAEPWithOptions`

```go
func EncryptOAEPWithOptions(
	random io.Reader,
	pub *PublicKey,
	msg []byte,
	opts *OAEPOptions,
) ([]byte, error)
```

This is a more configurable OAEP encryption function.

It can be used when you specifically need to control OAEP and MGF1 hash configuration.

For most applications:

```go
rsa.EncryptOAEP(
	sha256.New(),
	rand.Reader,
	publicKey,
	message,
	nil,
)
```

is easier and preferable.

---

# F. `EncryptPKCS1v15` — Deprecated

```go
func EncryptPKCS1v15(
	random io.Reader,
	pub *PublicKey,
	msg []byte,
) ([]byte, error)
```

Encrypts using the older:

```text
RSA PKCS #1 v1.5
```

padding scheme.

Example:

```go
ciphertext, err := rsa.EncryptPKCS1v15(
	rand.Reader,
	publicKey,
	message,
)
```

### Don't use this for new applications.

For new code, use:

```go
rsa.EncryptOAEP(...)
```

instead.

---

# G. `DecryptPKCS1v15` — Deprecated

```go
func DecryptPKCS1v15(
	random io.Reader,
	priv *PrivateKey,
	ciphertext []byte,
) ([]byte, error)
```

Decrypts PKCS #1 v1.5 ciphertext.

It exists primarily for compatibility with older protocols and systems.

For new applications:

```go
rsa.DecryptOAEP(...)
```

should normally be preferred.

---

# H. `DecryptPKCS1v15SessionKey` — Deprecated

```go
func DecryptPKCS1v15SessionKey(
	random io.Reader,
	priv *PrivateKey,
	ciphertext []byte,
	key []byte,
) error
```

This is a specialized API for safely handling RSA PKCS #1 v1.5 encrypted **session keys**.

It is designed around protections against **Bleichenbacher-style chosen-ciphertext attacks**.

The idea is approximately:

```text
Random session key
       ↓
fallback buffer

RSA ciphertext
       ↓
attempt decryption
       ↓
if valid → replace buffer
if invalid → keep random key
```

The function is deprecated for new designs; OAEP is preferred.

---

# I. `SignPSS`

```go
func SignPSS(
	random io.Reader,
	priv *PrivateKey,
	hash crypto.Hash,
	digest []byte,
	opts *PSSOptions,
) ([]byte, error)
```

Creates an RSA-PSS digital signature.

This is one of the preferred RSA signing mechanisms for modern applications.

First hash your message:

```go
digest := sha256.Sum256(message)
```

Then sign:

```go
signature, err := rsa.SignPSS(
	rand.Reader,
	privateKey,
	crypto.SHA256,
	digest[:],
	nil,
)
```

Conceptually:

```text
Message
   ↓
SHA-256
   ↓
Digest
   ↓
RSA-PSS
   ↓
Signature
```

PSS uses randomness, so signing the same message multiple times can produce different signatures.

---

# J. `VerifyPSS`

```go
func VerifyPSS(
	pub *PublicKey,
	hash crypto.Hash,
	digest []byte,
	sig []byte,
	opts *PSSOptions,
) error
```

Verifies an RSA-PSS signature.

Example:

```go
digest := sha256.Sum256(message)

err := rsa.VerifyPSS(
	publicKey,
	crypto.SHA256,
	digest[:],
	signature,
	nil,
)

if err != nil {
	fmt.Println("Invalid signature")
} else {
	fmt.Println("Valid signature")
}
```

A `nil` error means the signature is valid.

---

# K. `SignPKCS1v15`

```go
func SignPKCS1v15(
	random io.Reader,
	priv *PrivateKey,
	hash crypto.Hash,
	hashed []byte,
) ([]byte, error)
```

Creates an RSA PKCS #1 v1.5 signature.

Example:

```go
digest := sha256.Sum256(message)

signature, err := rsa.SignPKCS1v15(
	nil,
	privateKey,
	crypto.SHA256,
	digest[:],
)
```

Unlike PKCS #1 v1.5 **encryption**, PKCS #1 v1.5 signatures are not simply the same thing and remain widely used for interoperability.

For new systems, RSA-PSS is generally the preferred modern RSA signature scheme.

---

# L. `VerifyPKCS1v15`

```go
func VerifyPKCS1v15(
	pub *PublicKey,
	hash crypto.Hash,
	hashed []byte,
	sig []byte,
) error
```

Verifies an RSA PKCS #1 v1.5 signature.

Example:

```go
digest := sha256.Sum256(message)

err := rsa.VerifyPKCS1v15(
	publicKey,
	crypto.SHA256,
	digest[:],
	signature,
)
```

If:

```go
err == nil
```

the signature is valid.

Otherwise it is invalid.

---

# M. `PrivateKey.Decrypt`

```go
func (priv *PrivateKey) Decrypt(
	rand io.Reader,
	ciphertext []byte,
	opts crypto.DecrypterOpts,
) ([]byte, error)
```

This is the interface-oriented decryption method.

It allows an RSA private key to implement Go's generic:

```go
crypto.Decrypter
```

interface.

Depending on `opts`, it can perform different RSA decryption operations.

This is useful when you want code to work with an abstract cryptographic key rather than explicitly depending on RSA everywhere.

---

# N. `PrivateKey.Sign`

```go
func (priv *PrivateKey) Sign(
	rand io.Reader,
	digest []byte,
	opts crypto.SignerOpts,
) ([]byte, error)
```

This allows an RSA private key to implement the generic:

```go
crypto.Signer
```

interface.

Example:

```go
signature, err := privateKey.Sign(
	rand.Reader,
	digest[:],
	opts,
)
```

If `opts` specifies PSS options, PSS is used.

Otherwise PKCS #1 v1.5 signing is used.

---

# O. `PrivateKey.Public`

```go
func (priv *PrivateKey) Public() crypto.PublicKey
```

Returns the public key corresponding to the private key.

Example:

```go
privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)

publicKey := privateKey.Public()
```

Usually you'll see:

```go
publicKey := &privateKey.PublicKey
```

when you specifically need `*rsa.PublicKey`.

---

# P. `PrivateKey.Validate`

```go
func (priv *PrivateKey) Validate() error
```

Checks whether the RSA private key is mathematically valid.

Example:

```go
err := privateKey.Validate()

if err != nil {
	fmt.Println("Invalid RSA key")
}
```

This can detect problems with the mathematical relationships inside the RSA key.

It is useful when loading keys from external sources.

---

# Q. `PrivateKey.Precompute`

```go
func (priv *PrivateKey) Precompute()
```

Precomputes values used to speed up future private-key operations.

Example:

```go
privateKey.Precompute()
```

This can improve performance for repeated operations.

Important:

> `Precompute()` modifies the private key.

It should not be called concurrently with another method on that key.

---

# R. `PrivateKey.Equal`

```go
func (priv *PrivateKey) Equal(
	x crypto.PrivateKey,
) bool
```

Determines whether two private keys have equivalent values.

The precomputed values are ignored when determining equality.

Example:

```go
same := privateKey.Equal(otherPrivateKey)
```

---

# S. `PublicKey.Size`

```go
func (pub *PublicKey) Size() int
```

Returns the RSA modulus size in **bytes**.

For example, a 2048-bit RSA key has:

```text
2048 bits ÷ 8 = 256 bytes
```

So:

```go
fmt.Println(publicKey.Size())
```

would normally produce:

```text
256
```

This is useful when dealing with RSA ciphertext or signature sizes.

---

# T. `PublicKey.Equal`

```go
func (pub *PublicKey) Equal(
	x crypto.PublicKey,
) bool
```

Checks whether two public keys have the same value.

Example:

```go
if publicKey.Equal(otherPublicKey) {
	fmt.Println("Same public key")
}
```

---

# 6. `PSSOptions`

```go
type PSSOptions struct {
	SaltLength int
	Hash       crypto.Hash
}
```

Controls RSA-PSS signatures.

For example:

```go
opts := &rsa.PSSOptions{
	SaltLength: rsa.PSSSaltLengthEqualsHash,
	Hash:       crypto.SHA256,
}
```

Then:

```go
signature, err := rsa.SignPSS(
	rand.Reader,
	privateKey,
	crypto.SHA256,
	digest[:],
	opts,
)
```

---

# 7. PSS salt constants

The package provides:

```go
rsa.PSSSaltLengthAuto
```

and:

```go
rsa.PSSSaltLengthEqualsHash
```

## `PSSSaltLengthAuto`

```go
rsa.PSSSaltLengthAuto
```

uses a salt length that is as large as possible when signing and automatically detected during verification.

## `PSSSaltLengthEqualsHash`

```go
rsa.PSSSaltLengthEqualsHash
```

makes the salt length equal to the hash length.

For SHA-256:

```text
SHA-256 = 32 bytes
```

so the salt is 32 bytes.

---

# 8. `PSSOptions.HashFunc`

```go
func (opts *PSSOptions) HashFunc() crypto.Hash
```

Returns the hash function specified in the options.

This method is useful because `PSSOptions` can implement the necessary hashing-options interface used by Go's cryptographic APIs.

---

# 9. `OAEPOptions`

The package also provides:

```go
type OAEPOptions struct {
	Hash    crypto.Hash
	MGFHash crypto.Hash
	Label   []byte
}
```

It allows more explicit control over OAEP parameters.

Conceptually:

```text
OAEPOptions
 ├── Hash
 ├── MGFHash
 └── Label
```

Most applications don't need this level of control.

The simpler:

```go
rsa.EncryptOAEP(...)
```

is usually preferable.

---

# 10. `CRTValue`

```go
type CRTValue struct {
	Exp   *big.Int
	Coeff *big.Int
	R     *big.Int
}
```

`CRTValue` represents values used for **Chinese Remainder Theorem (CRT)** optimization of RSA private-key operations.

You generally don't construct these manually.

The purpose is performance.

Instead of doing one large mathematical operation, RSA can use smaller operations involving the prime factors and then combine the results.

Conceptually:

```text
RSA private operation
        │
        ▼
   CRT optimization
      /       \
     /         \
smaller op   smaller op
     \         /
      \       /
       combine
          │
          ▼
       result
```

---

# 11. `PrecomputedValues`

RSA has:

```go
type PrecomputedValues struct {
	Dp, Dq *big.Int
	Qinv   *big.Int

	CRTValues []CRTValue
}
```

These values are generated for faster private-key operations.

You normally don't manipulate this structure yourself.

Instead:

```go
privateKey.Precompute()
```

does the work.

---

# 12. Errors provided by the package

The package defines several important errors.

## `rsa.ErrMessageTooLong`

```go
rsa.ErrMessageTooLong
```

Occurs when a message is too large for the RSA operation.

For example:

```go
ciphertext, err := rsa.EncryptOAEP(...)
```

can fail if the plaintext exceeds the maximum size permitted by the RSA key and OAEP parameters.

This is one reason RSA isn't suitable for bulk data encryption.

---

## `rsa.ErrDecryption`

```go
rsa.ErrDecryption
```

Represents a decryption failure.

The error is intentionally vague.

That's important for security because giving attackers detailed information about why RSA decryption failed can contribute to adaptive attacks.

---

## `rsa.ErrVerification`

```go
rsa.ErrVerification
```

Represents a signature verification failure.

Again, cryptographic APIs intentionally avoid exposing unnecessary details.

---

# 13. Important distinction: encryption vs signing

This is probably the **most important concept** to understand.

## Encryption

Goal:

> Keep information secret.

```text
Sender
  │
  │ plaintext
  ▼
Recipient's PUBLIC KEY
  │
  ▼
ciphertext
  │
  ▼
Recipient's PRIVATE KEY
  │
  ▼
plaintext
```

Use:

```go
rsa.EncryptOAEP(...)
rsa.DecryptOAEP(...)
```

---

## Digital signature

Goal:

> Prove authenticity and integrity.

```text
Message
   │
   ▼
Private Key
   │
   ▼
Signature
```

Anyone possessing the public key can verify:

```text
Message + Signature + Public Key
               │
               ▼
          Valid / Invalid
```

Use:

```go
rsa.SignPSS(...)
rsa.VerifyPSS(...)
```

---

# 14. Complete signing example

Here's a useful beginner example:

```go
package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

func main() {
	// Generate RSA private key.
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	publicKey := &privateKey.PublicKey

	message := []byte("Important transaction")

	// Hash the message.
	digest := sha256.Sum256(message)

	// Sign the digest.
	signature, err := rsa.SignPSS(
		rand.Reader,
		privateKey,
		crypto.SHA256,
		digest[:],
		nil,
	)
	if err != nil {
		panic(err)
	}

	// Verify the signature.
	err = rsa.VerifyPSS(
		publicKey,
		crypto.SHA256,
		digest[:],
		signature,
		nil,
	)

	if err != nil {
		fmt.Println("Signature is INVALID")
		return
	}

	fmt.Println("Signature is VALID")
}
```

Notice that we don't normally sign a large message directly.

Instead:

```text
Large message
     ↓
SHA-256
     ↓
32-byte digest
     ↓
RSA-PSS
     ↓
signature
```

---

# 15. Three common beginner mistakes

## Mistake 1 — Encrypting large data directly with RSA

Beginners sometimes try:

```go
rsa.EncryptOAEP(..., hugeFile, ...)
```

This is wrong.

RSA has a strict message-size limitation.

### Better approach

Use hybrid encryption:

```text
                 Random AES key
                      │
          ┌───────────┴───────────┐
          ↓                       ↓
     AES-GCM                  RSA-OAEP
          ↓                       ↓
     encrypted data          encrypted AES key
```

RSA protects the key.

AES-GCM protects the actual data.

---

## Mistake 2 — Using PKCS #1 v1.5 encryption in new code

You may see:

```go
rsa.EncryptPKCS1v15(...)
```

and assume it's the normal RSA encryption function.

It isn't recommended for new designs.

Use:

```go
rsa.EncryptOAEP(...)
```

instead.

---

## Mistake 3 — Thinking RSA signatures encrypt the message

A signature does **not** provide confidentiality.

This:

```go
rsa.SignPSS(...)
```

does not mean:

> "Encrypt this message so nobody can read it."

It means:

> "Create cryptographic evidence that this message was signed by the holder of the private key."

For confidentiality, use encryption.

For authenticity/integrity, use signatures.

Often you need **both**.

---

# 16. Two real-world applications

## Application 1 — Hybrid encryption

Imagine you're building a secure file-sharing application.

You have a 500 MB file.

You wouldn't RSA-encrypt the entire file.

Instead:

```text
500 MB file
    │
    ▼
AES-GCM
    │
    ├── encrypted file
    │
    └── AES key
           │
           ▼
       RSA-OAEP
           │
           ▼
     encrypted AES key
```

The recipient uses their RSA private key to recover the AES key and then uses AES-GCM to decrypt the file.

This is one of the most important practical applications of RSA encryption.

---

## Application 2 — Digital signatures

Imagine a software company distributes an executable:

```text
myprogram.exe
```

The company can calculate a digest and sign it using its RSA private key.

Users can then verify the signature using the company's public key.

Conceptually:

```text
Developer

program
   ↓
SHA-256
   ↓
RSA-PSS + private key
   ↓
signature
```

User:

```text
program + signature
        │
        ▼
RSA-PSS + public key
        │
        ▼
    VALID / INVALID
```

This helps detect modification and authenticate the signer.

---

# 17. Three progressively challenging exercises

## Exercise 1 — RSA-OAEP basics

Write a Go program that:

1. Generates a 2048-bit RSA key pair.
2. Extracts the public key.
3. Encrypts the message:

```text
Learning RSA in Go
```

using RSA-OAEP and SHA-256.
4. Decrypts the ciphertext.
5. Prints the original plaintext.
6. Demonstrates that the decrypted message exactly matches the original message.

**Do not use PKCS #1 v1.5 encryption.**

---

## Exercise 2 — RSA-PSS digital signatures

Build a program that:

1. Generates a 2048-bit RSA key pair.
2. Creates a message representing a transaction.
3. Calculates its SHA-256 digest.
4. Signs the digest using RSA-PSS.
5. Verifies the signature using the public key.
6. Changes one byte of the original message.
7. Attempts verification again.
8. Observe and explain why verification succeeds initially but fails after modification.

---

## Exercise 3 — Hybrid encryption system

Build a small secure-message system that:

1. Generates an RSA key pair.
2. Generates a random AES key.
3. Uses AES-GCM to encrypt a larger message.
4. Uses RSA-OAEP with SHA-256 to encrypt the AES key.
5. Produces a structure containing:
   - RSA-encrypted AES key
   - AES-GCM nonce
   - encrypted message
6. Decrypts the AES key using the RSA private key.
7. Uses the recovered AES key to decrypt the message.
8. Adds an RSA-PSS signature over the appropriate data.
9. Verifies the signature before accepting the decrypted message.
10. Test what happens when the ciphertext, nonce, encrypted AES key, or signature is modified.

The goal is to understand why real systems generally combine **asymmetric cryptography + symmetric cryptography + authentication**, rather than attempting to use RSA for everything.

---

# 18. A useful mental model

If you're learning `crypto/rsa`, remember this table:

| Task | Preferred RSA API |
|---|---|
| Generate key | `rsa.GenerateKey` |
| Encrypt small secret | `rsa.EncryptOAEP` |
| Decrypt OAEP | `rsa.DecryptOAEP` |
| Sign | `rsa.SignPSS` |
| Verify | `rsa.VerifyPSS` |
| Legacy encryption | `EncryptPKCS1v15` — deprecated |
| Legacy decryption | `DecryptPKCS1v15` — deprecated |
| Legacy session-key decryption | `DecryptPKCS1v15SessionKey` — deprecated |
| Legacy signing | `SignPKCS1v15` |
| Legacy verification | `VerifyPKCS1v15` |
| Validate key | `PrivateKey.Validate` |
| Optimize private operations | `PrivateKey.Precompute` |
| Get public key | `PrivateKey.Public` |
| Get RSA key size | `PublicKey.Size` |
| Compare public keys | `PublicKey.Equal` |
| Compare private keys | `PrivateKey.Equal` |

The key distinction to memorize is:

```text
                RSA
                 │
       ┌─────────┴─────────┐
       │                   │
   Encryption           Signature
       │                   │
      OAEP                 PSS
       │                   │
Confidentiality      Authenticity
                    + Integrity
```

And the most important production rule is:

> **Don't design a system around "RSA alone."** Use RSA for the small asymmetric part of a protocol, such as protecting a symmetric key or creating a signature, and use appropriate symmetric/authenticated encryption for the actual data.

---

# 19. Thought-provoking question

Suppose you build a messaging application where **Alice encrypts a message with Bob's RSA public key** and sends the ciphertext to Bob.

**Question:** If an attacker changes the ciphertext while it is being transmitted, what does RSA-OAEP alone tell Bob about **who created the message** or whether it was **intentionally sent by Alice**?

What additional cryptographic mechanism would you need, and why?

---

## Key takeaway

Remember these four APIs first:

```go
rsa.GenerateKey(...)
rsa.EncryptOAEP(...)
rsa.DecryptOAEP(...)
rsa.SignPSS(...)
rsa.VerifyPSS(...)
```

A good learning path is:

```text
RSA key generation
       ↓
RSA-OAEP encryption/decryption
       ↓
RSA-PSS signatures
       ↓
Key serialization
       ↓
Hybrid encryption
       ↓
Authenticated secure protocols
```

Do not start by memorizing every field in the RSA implementation. First understand **what problem each primitive solves**, then learn the API details needed to implement it safely.
