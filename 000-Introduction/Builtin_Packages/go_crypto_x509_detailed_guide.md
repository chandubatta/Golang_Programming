# Go `crypto/x509` Package — Detailed Guide

> A practical guide to understanding X.509 certificates, certificate chains, CSRs, CRLs, trust stores, key encodings, verification, common mistakes, real-world applications, and exercises in Go.

## 1. What is `crypto/x509`?

The Go `crypto/x509` package implements a substantial subset of the **X.509 standard**.

In simple terms:

> **`crypto/x509` is Go's toolbox for working with digital certificates, certificate chains, CSRs, CRLs, and standardized public/private-key encodings.**

It can:

- Parse X.509 certificates.
- Generate X.509 certificates.
- Verify certificate chains.
- Verify hostnames.
- Create and parse certificate signing requests (CSRs).
- Create and parse certificate revocation lists (CRLs).
- Create and parse RSA/ECDSA/Ed25519 and other supported keys.
- Manage trusted root certificates using `CertPool`.
- Work with certificate policies and OIDs.
- Inspect certificate properties such as:
  - Subject
  - Issuer
  - Serial number
  - Validity period
  - SANs
  - Key usage
  - Extended key usage
  - Public key
  - Signature algorithm
  - CA status

The package is heavily used by Go's TLS ecosystem.

---

## 2. Where does `x509` fit into HTTPS?

A simplified HTTPS certificate chain looks like:

```text
                    Root CA
                       │
                       │ signs
                       ▼
                Intermediate CA
                       │
                       │ signs
                       ▼
               Server Certificate
                       │
                       │
                       ▼
                 example.com
```

When your Go HTTPS client connects to:

```text
https://example.com
```

certificate validation needs to determine:

```text
Is this certificate genuine?
        │
        ▼
Was it signed by a trusted CA?
        │
        ▼
Does the certificate belong to example.com?
        │
        ▼
Is it currently valid?
        │
        ▼
Are its key usages appropriate?
        │
        ▼
Is the chain structurally valid?
        │
        ▼
YES → trust the certificate
```

`crypto/x509` performs much of this certificate validation work.

**Important:** `x509.Verify()` does **not** perform certificate revocation checking.

---

## 3. The three important encodings

Before learning the API, understand this:

```text
PEM
 │
 │ Base64 text wrapper
 ▼
DER
 │
 │ ASN.1 binary encoding
 ▼
Certificate / Key structure
```

For example:

```text
-----BEGIN CERTIFICATE-----
MIID...
...
-----END CERTIFICATE-----
```

That is **PEM**.

Inside it is Base64-encoded **DER**.

And DER represents an ASN.1 structure.

This distinction is extremely important.

`x509.ParseCertificate()` expects **DER**, not PEM.

Correct:

```go
block, _ := pem.Decode(pemBytes)

cert, err := x509.ParseCertificate(block.Bytes)
```

Not:

```go
x509.ParseCertificate(pemBytes) // wrong if pemBytes is PEM
```

---

# 4. Simple example — parse a certificate

Suppose you have:

```text
server.crt
```

containing a PEM certificate:

```go
package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("server.crt")
	if err != nil {
		panic(err)
	}

	// Convert PEM → DER.
	block, _ := pem.Decode(data)

	if block == nil {
		panic("failed to decode PEM")
	}

	// Convert DER → x509.Certificate.
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		panic(err)
	}

	fmt.Println("Subject:", cert.Subject)
	fmt.Println("Issuer:", cert.Issuer)
	fmt.Println("Serial:", cert.SerialNumber)
	fmt.Println("Valid from:", cert.NotBefore)
	fmt.Println("Valid until:", cert.NotAfter)
	fmt.Println("DNS names:", cert.DNSNames)
	fmt.Println("Is CA:", cert.IsCA)
	fmt.Println("Public key algorithm:", cert.PublicKeyAlgorithm)
	fmt.Println("Signature algorithm:", cert.SignatureAlgorithm)
}
```

Conceptually:

```text
server.crt
    │
    ▼
pem.Decode()
    │
    ▼
DER bytes
    │
    ▼
x509.ParseCertificate()
    │
    ▼
*Certificate
    │
    ├── Subject
    ├── Issuer
    ├── DNSNames
    ├── NotBefore
    ├── NotAfter
    ├── PublicKey
    ├── KeyUsage
    └── ...
```

---

# 5. All major exported functions in `crypto/x509`

The following sections explain the major exported functions and methods.

## Certificate creation

### `CreateCertificate`

```go
func CreateCertificate(
    rand io.Reader,
    template, parent *Certificate,
    pub, priv any,
) ([]byte, error)
```

Creates an X.509 v3 certificate.

Important parameters:

```text
rand
    randomness source

template
    describes the certificate being created

parent
    certificate that signs it

pub
    public key of the new certificate

priv
    private key used by the signer
```

For example:

```text
Root CA
  │
  │ private key signs
  ▼
CreateCertificate()
  │
  ▼
Server certificate
```

If:

```go
template == parent
```

the certificate is self-signed.

The returned certificate is DER encoded.

---

### `CreateCertificateRequest`

```go
func CreateCertificateRequest(
    rand io.Reader,
    template *CertificateRequest,
    priv any,
) ([]byte, error)
```

Creates a **CSR — Certificate Signing Request**.

A CSR essentially says:

> "I have this public key and I want a CA to issue me a certificate for these names."

Typical workflow:

```text
Generate private key
       │
       ▼
CreateCertificateRequest()
       │
       ▼
CSR
       │
       ▼
Certificate Authority
       │
       ▼
Certificate
```

Common fields include:

```go
DNSNames
EmailAddresses
IPAddresses
URIs
Subject
```

The returned data is DER encoded.

---

### `CreateRevocationList`

```go
func CreateRevocationList(
    rand io.Reader,
    template *RevocationList,
    issuer *Certificate,
    priv crypto.Signer,
) ([]byte, error)
```

Creates an X.509 v2 **CRL — Certificate Revocation List**.

Imagine a certificate was issued:

```text
Certificate #12345
```

but its private key was stolen.

The CA can revoke it.

A CRL can contain:

```text
12345 → revoked
12367 → revoked
12401 → revoked
```

`CreateRevocationList` creates such a list.

The issuer must be appropriate for CRL signing.

---

# 6. PEM encryption functions

There are three older PEM functions:

### `EncryptPEMBlock`

```go
func EncryptPEMBlock(
    rand io.Reader,
    blockType string,
    data, password []byte,
    alg PEMCipher,
) (*pem.Block, error)
```

Encrypts data using the old PEM encryption format.

### `DecryptPEMBlock`

```go
func DecryptPEMBlock(
    b *pem.Block,
    password []byte,
) ([]byte, error)
```

Decrypts that old format.

### `IsEncryptedPEMBlock`

```go
func IsEncryptedPEMBlock(b *pem.Block) bool
```

Checks whether the PEM block uses that legacy encryption format.

**Important:** all three are deprecated.

The legacy RFC 1423 PEM encryption format is insecure because it does not authenticate ciphertext and is vulnerable to padding-oracle attacks.

For new applications, do not build new security designs around these APIs.

---

# 7. Private-key encoding functions

## `MarshalECPrivateKey`

```go
func MarshalECPrivateKey(
    key *ecdsa.PrivateKey,
) ([]byte, error)
```

Converts an ECDSA private key to SEC 1 ASN.1 DER.

Usually represented as:

```text
-----BEGIN EC PRIVATE KEY-----
...
-----END EC PRIVATE KEY-----
```

For a more general private-key format, use PKCS#8.

---

## `MarshalPKCS1PrivateKey`

```go
func MarshalPKCS1PrivateKey(
    key *rsa.PrivateKey,
) []byte
```

Converts an RSA private key to PKCS#1 DER.

Usually:

```text
-----BEGIN RSA PRIVATE KEY-----
...
-----END RSA PRIVATE KEY-----
```

It is specifically for RSA.

---

## `MarshalPKCS8PrivateKey`

```go
func MarshalPKCS8PrivateKey(
    key any,
) ([]byte, error)
```

Converts a private key into the more general **PKCS#8** format.

Current Go supports several key types, including RSA, ECDSA, Ed25519, ML-DSA and ECDH/X25519 where supported.

Usually:

```text
-----BEGIN PRIVATE KEY-----
...
-----END PRIVATE KEY-----
```

This is often preferable when you want a generic private-key format rather than an algorithm-specific format.

---

# 8. Public-key encoding functions

## `MarshalPKCS1PublicKey`

```go
func MarshalPKCS1PublicKey(
    key *rsa.PublicKey,
) []byte
```

Encodes an RSA public key using PKCS#1.

Typically:

```text
-----BEGIN RSA PUBLIC KEY-----
...
-----END RSA PUBLIC KEY-----
```

---

## `MarshalPKIXPublicKey`

```go
func MarshalPKIXPublicKey(
    pub any,
) ([]byte, error)
```

Encodes a public key using the **PKIX SubjectPublicKeyInfo** format.

Typically:

```text
-----BEGIN PUBLIC KEY-----
...
-----END PUBLIC KEY-----
```

It supports several public-key algorithms.

---

# 9. Parsing private keys

## `ParseECPrivateKey`

```go
func ParseECPrivateKey(
    der []byte,
) (*ecdsa.PrivateKey, error)
```

Reads an ECDSA private key encoded as SEC 1 DER.

---

## `ParsePKCS1PrivateKey`

```go
func ParsePKCS1PrivateKey(
    der []byte,
) (*rsa.PrivateKey, error)
```

Reads a PKCS#1 RSA private key.

---

## `ParsePKCS8PrivateKey`

```go
func ParsePKCS8PrivateKey(
    der []byte,
) (any, error)
```

Reads a generic PKCS#8 private key.

Because different algorithms can be stored inside PKCS#8, it returns:

```go
any
```

You commonly use a type switch:

```go
key, err := x509.ParsePKCS8PrivateKey(der)
if err != nil {
	panic(err)
}

switch key := key.(type) {
case *rsa.PrivateKey:
	fmt.Println("RSA")
case *ecdsa.PrivateKey:
	fmt.Println("ECDSA")
case ed25519.PrivateKey:
	fmt.Println("Ed25519")
}
```

---

# 10. Parsing public keys

## `ParsePKCS1PublicKey`

```go
func ParsePKCS1PublicKey(
    der []byte,
) (*rsa.PublicKey, error)
```

Parses an RSA public key in PKCS#1 format.

---

## `ParsePKIXPublicKey`

```go
func ParsePKIXPublicKey(
    derBytes []byte,
) (any, error)
```

Parses a generic PKIX `SubjectPublicKeyInfo`.

Possible returned types include:

```text
*rsa.PublicKey
*dsa.PublicKey
*ecdsa.PublicKey
ed25519.PublicKey
*mldsa.PublicKey
*ecdh.PublicKey
```

Because the return type is `any`, a type switch is often required.

---

# 11. Certificate parsing

## `ParseCertificate`

```go
func ParseCertificate(
    der []byte,
) (*Certificate, error)
```

Converts:

```text
DER certificate
       ↓
*Certificate
```

Then you can inspect:

```go
cert.Subject
cert.Issuer
cert.DNSNames
cert.IPAddresses
cert.NotBefore
cert.NotAfter
cert.PublicKey
cert.IsCA
cert.KeyUsage
cert.ExtKeyUsage
```

This is probably the most commonly used `x509` parsing function.

It expects a single DER certificate.

---

## `ParseCertificates`

```go
func ParseCertificates(
    der []byte,
) ([]*Certificate, error)
```

Parses multiple concatenated DER certificates.

Useful for certificate chains.

---

# 12. Certificate verification functions

## `Certificate.Verify`

```go
func (c *Certificate) Verify(
    opts VerifyOptions,
) ([][]*Certificate, error)
```

This is the major certificate-chain verification API.

It attempts to construct:

```text
Leaf
 │
 ▼
Intermediate
 │
 ▼
Root CA
```

using:

```go
opts.Roots
opts.Intermediates
```

Example:

```go
opts := x509.VerifyOptions{
	DNSName: "api.example.com",
	Roots:   roots,
}

chains, err := cert.Verify(opts)

if err != nil {
	fmt.Println("Certificate invalid:", err)
	return
}

fmt.Println("Certificate is valid")
fmt.Println("Number of valid chains:", len(chains))
```

`Verify` checks things such as:

- certificate signatures
- validity periods
- CA constraints
- key usage
- extended key usage
- name constraints
- certificate policies
- hostname when `DNSName` is supplied
- chain construction

But:

```text
Verify()
   ≠
Revocation checking
```

It does **not** perform revocation checking.

---

## `Certificate.VerifyHostname`

```go
func (c *Certificate) VerifyHostname(
    h string,
) error
```

Checks whether a certificate is valid for a particular hostname.

Example:

```go
err := cert.VerifyHostname("api.example.com")

if err != nil {
	fmt.Println("Hostname mismatch")
}
```

It checks DNS SANs, IP SANs and wildcard rules.

The old certificate `Common Name` is ignored for modern hostname validation. Subject Alternative Name (SAN) is used.

---

# 13. Low-level signature functions

## `Certificate.CheckSignature`

```go
func (c *Certificate) CheckSignature(
    algo SignatureAlgorithm,
    signed, signature []byte,
) error
```

Verifies a raw signature using the certificate's public key.

It is a low-level operation.

It does not perform complete certificate validation.

Therefore:

```text
CheckSignature()
```

does **not** mean:

```text
"This certificate is trusted"
```

It means the public key correctly verifies the supplied signature under the specified algorithm.

---

## `Certificate.CheckSignatureFrom`

```go
func (c *Certificate) CheckSignatureFrom(
    parent *Certificate,
) error
```

Checks whether:

```text
parent
   │
   │ signed
   ▼
c
```

For example:

```go
err := child.CheckSignatureFrom(parent)
```

This is still **not complete certificate-chain verification**.

For normal certificate validation, prefer:

```go
child.Verify(...)
```

---

# 14. Certificate equality

## `Certificate.Equal`

```go
func (c *Certificate) Equal(
    other *Certificate,
) bool
```

Checks whether two certificates are equivalent.

Useful when comparing parsed certificate objects.

---

# 15. CRL-related certificate functions

## `Certificate.CheckCRLSignature`

```go
func (c *Certificate) CheckCRLSignature(
    crl *pkix.CertificateList,
) error
```

Checks whether a CRL was signed by the certificate.

However, this method is **deprecated**.

Use:

```go
RevocationList.CheckSignatureFrom(...)
```

instead.

---

## `Certificate.CreateCRL`

Creates a certificate revocation list.

However, this method is also **deprecated**.

The recommended modern API is:

```go
x509.CreateRevocationList(...)
```

---

# 16. `CertPool`

One of the most important concepts in `x509` is:

```go
type CertPool
```

Think of it as:

> **A collection of certificates that Go is allowed to trust as certificate authorities.**

For example:

```text
CertPool
 ├── Root CA A
 ├── Root CA B
 ├── Root CA C
 └── Root CA D
```

---

## `NewCertPool`

```go
roots := x509.NewCertPool()
```

Creates an empty certificate pool.

---

## `SystemCertPool`

```go
roots, err := x509.SystemCertPool()
```

Loads a copy of the system's trusted certificate pool.

The returned pool is a copy; modifying it does not modify the system certificate store.

---

## `AddCert`

```go
roots.AddCert(cert)
```

Adds one certificate to a certificate pool.

Example:

```go
roots := x509.NewCertPool()
roots.AddCert(myRootCA)
```

---

## `AppendCertsFromPEM`

```go
ok := roots.AppendCertsFromPEM(rootPEM)
```

Parses PEM-encoded certificates and adds them to the pool.

Example:

```go
roots := x509.NewCertPool()

ok := roots.AppendCertsFromPEM(rootPEM)

if !ok {
	panic("failed to load root certificate")
}
```

It returns `true` if at least one certificate was successfully parsed.

---

## `AddCertWithConstraint`

```go
pool.AddCertWithConstraint(cert, constraint)
```

An advanced API that lets you add a root certificate together with an additional custom constraint.

Conceptually:

```go
pool.AddCertWithConstraint(root, func(chain []*x509.Certificate) error {
	// inspect the entire chain
	return nil
})
```

Returning an error rejects that chain.

---

## `Clone`

```go
copy := roots.Clone()
```

Creates a copy of a certificate pool.

---

## `Equal`

```go
same := roots.Equal(other)
```

Determines whether two certificate pools are equal.

---

## `Subjects`

```go
subjects := roots.Subjects()
```

Returns DER-encoded subjects in the pool.

This method is deprecated.

---

# 17. `CertificateRequest`

A `CertificateRequest` represents a **PKCS#10 CSR**.

Think:

```text
Developer/server
       │
       │ creates
       ▼
     CSR
       │
       │ sends
       ▼
      CA
       │
       │ signs
       ▼
Certificate
```

---

## `ParseCertificateRequest`

```go
func ParseCertificateRequest(
    asn1Data []byte,
) (*CertificateRequest, error)
```

Parses a DER-encoded CSR.

---

## `CertificateRequest.CheckSignature`

```go
func (c *CertificateRequest) CheckSignature() error
```

Checks whether the CSR's signature is valid.

Important:

```text
CSR signature valid
        ≠
CA has approved the CSR
```

It only establishes that the CSR was correctly signed by the corresponding private key.

---

# 18. Extended Key Usage

`ExtKeyUsage` specifies what a certificate is intended to be used for.

Examples:

```go
x509.ExtKeyUsageServerAuth
x509.ExtKeyUsageClientAuth
x509.ExtKeyUsageCodeSigning
x509.ExtKeyUsageEmailProtection
```

For example:

```text
Server certificate
    ↓
ServerAuth

Client certificate
    ↓
ClientAuth
```

This becomes particularly important in mTLS.

---

## `ExtKeyUsage.OID`

```go
eku.OID()
```

Returns the OID associated with the extended key usage.

---

## `ExtKeyUsage.String`

```go
eku.String()
```

Returns a human-readable string representation.

---

# 19. `KeyUsage`

`KeyUsage` describes lower-level cryptographic operations permitted by a certificate.

Examples:

```go
x509.KeyUsageDigitalSignature
x509.KeyUsageKeyEncipherment
x509.KeyUsageKeyAgreement
x509.KeyUsageCertSign
x509.KeyUsageCRLSign
```

Because it is a bitmap, you can combine values:

```go
usage := x509.KeyUsageDigitalSignature |
	x509.KeyUsageKeyEncipherment
```

For a CA:

```go
usage := x509.KeyUsageCertSign |
	x509.KeyUsageCRLSign
```

---

## `KeyUsage.String`

```go
usage.String()
```

Returns a human-readable representation.

---

# 20. OID APIs

OID means:

> **Object Identifier**

X.509 uses OIDs extensively to identify algorithms, extensions, policies and usages.

Example:

```text
1.2.840.113549.1.1.11
```

represents a particular cryptographic algorithm identifier.

Go's modern `x509.OID` type provides APIs for manipulating these identifiers.

---

## `OIDFromInts`

```go
oid, err := x509.OIDFromInts(
	[]uint64{1, 2, 840, 113549},
)
```

Creates an OID from integer components.

---

## `OIDFromASN1OID`

```go
oid, err := x509.OIDFromASN1OID(asn1OID)
```

Converts an `encoding/asn1.ObjectIdentifier` into the newer `x509.OID` representation.

---

## `ParseOID`

```go
oid, err := x509.ParseOID("1.2.840.113549")
```

Parses a dotted OID string.

---

## `OID.String`

```go
fmt.Println(oid.String())
```

Converts the OID to its dotted string representation.

---

## `OID.Equal`

```go
oid1.Equal(oid2)
```

Tests whether two OIDs represent the same identifier.

---

## `OID.EqualASN1OID`

```go
oid.EqualASN1OID(other)
```

Compares an `x509.OID` with an ASN.1 OID.

---

## `OID.MarshalBinary`

```go
data, err := oid.MarshalBinary()
```

Encodes the OID into binary form.

---

## `OID.UnmarshalBinary`

```go
err := oid.UnmarshalBinary(data)
```

Decodes an OID from binary form.

---

## `OID.MarshalText`

```go
data, err := oid.MarshalText()
```

Encodes the OID as text.

---

## `OID.UnmarshalText`

```go
err := oid.UnmarshalText(data)
```

Parses an OID from text.

---

## `OID.AppendBinary`

```go
data, err := oid.AppendBinary(buffer)
```

Appends the binary OID representation to an existing byte slice.

---

## `OID.AppendText`

```go
data, err := oid.AppendText(buffer)
```

Appends the textual OID representation to an existing byte slice.

---

# 21. `RevocationList`

`RevocationList` represents an X.509 CRL.

Conceptually:

```text
CA
 │
 ├── Certificate 1001 → valid
 ├── Certificate 1002 → revoked
 ├── Certificate 1003 → valid
 └── Certificate 1004 → revoked
```

The CRL contains revoked certificate serial numbers.

---

## `ParseRevocationList`

```go
func ParseRevocationList(
    der []byte,
) (*RevocationList, error)
```

Parses a DER-encoded X.509 v2 CRL.

---

## `RevocationList.CheckSignatureFrom`

```go
func (rl *RevocationList) CheckSignatureFrom(
    parent *Certificate,
) error
```

Checks that the CRL was signed by the given issuer certificate.

Conceptually:

```text
CA certificate
      │
      │ signs
      ▼
     CRL
```

---

# 22. `PublicKeyAlgorithm`

Represents the algorithm used by a certificate's public key.

Possible values include:

```go
x509.RSA
x509.DSA
x509.ECDSA
x509.Ed25519
x509.MLDSA
```

There is also:

```go
x509.UnknownPublicKeyAlgorithm
```

---

## `PublicKeyAlgorithm.String`

```go
fmt.Println(cert.PublicKeyAlgorithm.String())
```

Converts the algorithm identifier into a human-readable string.

---

# 23. `SignatureAlgorithm`

Represents the algorithm used to sign a certificate or related object.

Examples include:

```text
SHA256WithRSA
SHA384WithRSA
SHA512WithRSA

ECDSAWithSHA256
ECDSAWithSHA384
ECDSAWithSHA512

SHA256WithRSAPSS
SHA384WithRSAPSS
SHA512WithRSAPSS

PureEd25519

MLDSA44
MLDSA65
MLDSA87
```

---

## `SignatureAlgorithm.String`

```go
fmt.Println(cert.SignatureAlgorithm.String())
```

Produces a readable algorithm name.

---

# 24. Error types

Understanding `x509` errors is useful in real applications.

## `CertificateInvalidError`

Returned when a certificate fails some validation rule.

It contains information such as:

```text
Cert
Reason
Detail
```

Possible reasons include:

- Expired
- Name mismatch
- Too many intermediates
- Incompatible usage

The `Error()` method converts it to a readable error message.

---

## `HostnameError`

Returned when the certificate doesn't match the requested hostname.

Example:

```text
Certificate:
    api.example.com

Requested:
    api.attacker.com
```

The `Error()` method returns a readable explanation.

---

## `UnknownAuthorityError`

Indicates that the certificate issuer isn't trusted or known.

Typical situation:

```text
Server certificate
       │
       ▼
Private CA
       │
       X
Not in trusted roots
```

Verification can return an `UnknownAuthorityError`.

---

## `ConstraintViolationError`

Indicates that a requested operation isn't allowed by the certificate.

For example, attempting to use a certificate for a purpose that its key usage doesn't permit.

---

## `InsecureAlgorithmError`

Indicates that an insecure signature algorithm was rejected.

---

## `UnhandledCriticalExtension`

Indicates that a certificate contains a critical extension the implementation doesn't understand.

Critical extensions must be understood; otherwise the certificate should be rejected.

---

## `SystemRootsError`

Indicates that Go couldn't load the operating system's trusted root certificates.

It also provides:

```go
Error()
Unwrap()
```

so the underlying error can be inspected with normal Go error handling.

---

# 25. `VerifyOptions`

This structure controls certificate verification.

Important fields include:

```go
type VerifyOptions struct {
	DNSName                   string
	Intermediates             *CertPool
	Roots                     *CertPool
	CurrentTime               time.Time
	KeyUsages                 []ExtKeyUsage
	MaxConstraintComparisions int
	CertificatePolicies       []OID
}
```

## `DNSName`

```go
DNSName: "api.example.com"
```

Checks hostname validity.

---

## `Roots`

```go
Roots: roots
```

Specifies trusted root CAs.

If `nil`, Go can use system roots/platform verification as appropriate.

---

## `Intermediates`

Contains intermediate CA certificates.

For example:

```text
Root CA
   │
   ▼
Intermediate CA
   │
   ▼
Server certificate
```

The intermediate belongs in:

```go
opts.Intermediates
```

not:

```go
opts.Roots
```

---

## `CurrentTime`

Allows deterministic certificate testing:

```go
CurrentTime: someTime
```

If zero, the current time is used.

---

## `KeyUsages`

Controls acceptable extended key usages.

Example:

```go
KeyUsages: []x509.ExtKeyUsage{
	x509.ExtKeyUsageServerAuth,
}
```

An empty list means server authentication by default. `ExtKeyUsageAny` can be used when any usage is acceptable.

---

# 26. `SetFallbackRoots`

```go
x509.SetFallbackRoots(roots)
```

Sets fallback roots for environments where system certificates aren't available or platform verification isn't available.

This can be useful in environments such as minimal containers.

Important restrictions:

```text
roots cannot be nil
+
SetFallbackRoots may only be called once
```

Calling it improperly can panic.

---

# 27. Useful mental model of the package

Think about `crypto/x509` as six major areas:

```text
                    crypto/x509
                         │
       ┌─────────────────┼──────────────────┐
       │                 │                  │
       ▼                 ▼                  ▼
 Certificates          Keys               CSRs
       │                 │                  │
       │                 │                  │
       ▼                 ▼                  ▼
 Verify               Marshal             Create
 Parse                Parse               Parse
 Create
       │
       ▼
 Certificate Chains
       │
       ▼
    CertPool
       │
       ▼
 Trusted Root CAs
       │
       ▼
      CRLs
       │
       ▼
 Revocation Lists
```

---

# 28. `x509` vs `crypto/rsa`, `crypto/ecdsa`, `crypto/ed25519`

This distinction is extremely important.

Suppose you're using RSA.

`crypto/rsa` handles the **cryptographic mathematics**:

```text
RSA key generation
RSA signing
RSA verification
RSA encryption/decryption
```

Whereas `crypto/x509` handles **standardized certificate/key representations and PKI**:

```text
X.509 certificate
certificate chain
CSR
PKCS#1
PKCS#8
PKIX
certificate verification
CA trust
```

Similarly:

```text
crypto/ecdsa
      ↓
ECDSA mathematical operations

crypto/ed25519
      ↓
Ed25519 signing

crypto/x509
      ↓
Put those keys into standardized certificates/encodings
```

---

# 29. Three common beginner mistakes

## Mistake 1: Confusing PEM with DER

Beginners often write:

```go
cert, err := x509.ParseCertificate(pemBytes)
```

when `pemBytes` contains:

```text
-----BEGIN CERTIFICATE-----
...
-----END CERTIFICATE-----
```

That's wrong.

You need:

```go
block, _ := pem.Decode(pemBytes)

cert, err := x509.ParseCertificate(block.Bytes)
```

Remember:

```text
PEM → pem.Decode()
DER → x509.ParseCertificate()
```

---

## Mistake 2: Thinking `Verify` checks revocation

You might think:

```go
cert.Verify(opts)
```

means:

> "This certificate has not been revoked."

It doesn't.

`Verify` validates the certificate chain and associated constraints, but it does **not** perform revocation checking.

---

## Mistake 3: Trusting any certificate that has a valid signature

Suppose:

```text
Certificate signature = valid
```

That doesn't automatically mean:

```text
Certificate = trusted
```

A certificate can have a perfectly valid cryptographic signature and still:

- be expired
- be issued by an untrusted CA
- have the wrong hostname
- have inappropriate key usage
- violate name constraints
- fail chain validation

That's why:

```go
CheckSignatureFrom()
```

and:

```go
Verify()
```

are very different concepts.

---

# 30. Real-world application #1 — HTTPS/TLS

One of the biggest applications is HTTPS.

Imagine your Go application calls:

```text
https://api.example.com
```

TLS needs to establish:

```text
Is api.example.com's certificate legitimate?
```

Certificate validation involves:

```text
Server certificate
       │
       ▼
Certificate chain
       │
       ▼
Trusted root
       │
       ▼
Signature verification
       │
       ▼
Validity period
       │
       ▼
Hostname validation
       │
       ▼
Key usage / constraints
```

`crypto/x509` provides the certificate machinery underlying this process.

Go's `crypto/tls` package integrates with X.509 certificates for TLS authentication.

---

# 31. Real-world application #2 — Mutual TLS / mTLS

mTLS is particularly interesting.

Normal TLS:

```text
Client ───────────────► Server
       Server proves
       its identity
```

mTLS:

```text
Client ◄──────────────► Server
       both prove
       their identities
```

For example:

```text
                    Root CA
                       │
             ┌─────────┴─────────┐
             ▼                   ▼
       Client Certificate   Server Certificate
             │                   │
             ▼                   ▼
          Client              Server
```

The server can validate the client's certificate using:

```go
cert.Verify(...)
```

and require:

```go
x509.ExtKeyUsageClientAuth
```

This is common in:

- microservices
- internal APIs
- service meshes
- banking systems
- enterprise networks
- zero-trust architectures

---

# 32. Exercise 1 — Beginner

## Certificate Inspector

Write a Go program that:

1. Reads a PEM certificate from a file.
2. Decodes the PEM.
3. Parses the certificate using `x509.ParseCertificate`.
4. Prints:
   - Subject
   - Issuer
   - Serial number
   - Valid-from time
   - Valid-until time
   - DNS SANs
   - IP SANs
   - Whether it is a CA
   - Public-key algorithm
   - Signature algorithm
5. Determines whether the certificate is currently expired.

**Do not use OpenSSL.**

---

# 33. Exercise 2 — Intermediate

## Build Your Own Certificate Chain Verifier

Create a program that receives:

```text
root.pem
intermediate.pem
server.pem
```

Your program must:

1. Parse all three certificates.
2. Create a `CertPool` containing the root CA.
3. Create another `CertPool` containing the intermediate CA.
4. Configure `VerifyOptions`.
5. Verify that the server certificate chains to the root CA.
6. Require `ServerAuth`.
7. Verify a specific hostname.
8. Print the complete verified chain.

Then deliberately modify the program so that:

- the hostname is wrong
- the root CA is missing
- the intermediate CA is missing
- the certificate is expired

Observe and classify the resulting errors.

---

# 34. Exercise 3 — Advanced

## Build a Mini Private PKI

Build a small private certificate authority system entirely in Go.

Your program should:

1. Generate a root CA key.
2. Create a self-signed root CA certificate.
3. Generate an intermediate CA key.
4. Create an intermediate CA certificate signed by the root.
5. Generate a server key.
6. Create a CSR for the server.
7. Sign the CSR with the intermediate CA.
8. Create a server certificate containing SANs such as:
   ```text
   api.internal.example
   10.0.0.10
   ```
9. Build the complete certificate chain.
10. Verify the server certificate using `x509.Verify`.
11. Require `ExtKeyUsageServerAuth`.
12. Test hostname validation.
13. Generate a client certificate with `ExtKeyUsageClientAuth`.
14. Verify that the server certificate cannot be used as a client certificate when appropriate EKU restrictions are enforced.
15. Create a CRL containing a revoked certificate.
16. Parse the CRL.
17. Verify the CRL's signature.
18. Design your application so that a revoked certificate is rejected during your own revocation-checking logic.

The goal is to understand the complete lifecycle:

```text
Private Key
    ↓
Root CA
    ↓
Intermediate CA
    ↓
CSR
    ↓
Certificate
    ↓
Certificate Chain
    ↓
Verification
    ↓
Revocation
```

Don't use OpenSSL for the implementation.

---

# 35. Important security note

Because `x509` is security-sensitive code, keeping Go updated matters.

Security vulnerabilities have affected certificate-chain verification, name constraints, and pathological verification workloads in different Go releases.

For production certificate-processing software:

```text
Don't pin yourself to an old Go version
without a security reason.
```

Always check the Go security advisories and current package documentation when deploying security-sensitive systems.

---

# 36. Functions and concepts to learn first

Although `crypto/x509` contains many APIs, don't try to memorize everything at once.

## Level 1 — Certificate inspection

```go
pem.Decode()
x509.ParseCertificate()
```

Then understand:

```go
Certificate.Subject
Certificate.Issuer
Certificate.DNSNames
Certificate.NotBefore
Certificate.NotAfter
Certificate.PublicKey
```

## Level 2 — Certificate verification

```go
x509.NewCertPool()
x509.SystemCertPool()
CertPool.AddCert()
CertPool.AppendCertsFromPEM()
Certificate.Verify()
Certificate.VerifyHostname()
```

## Level 3 — Key encoding

```go
MarshalPKCS8PrivateKey()
ParsePKCS8PrivateKey()

MarshalPKIXPublicKey()
ParsePKIXPublicKey()
```

## Level 4 — Certificate creation

```go
CreateCertificate()
CreateCertificateRequest()
ParseCertificateRequest()
CertificateRequest.CheckSignature()
```

## Level 5 — Revocation

```go
CreateRevocationList()
ParseRevocationList()
RevocationList.CheckSignatureFrom()
```

## Level 6 — Advanced PKI

Then learn:

```text
OID
PolicyMapping
name constraints
certificate policies
custom CertPool constraints
CRLs
custom verification
```

---

# 37. Final mental model

If you remember only one diagram, remember this:

```text
                    PRIVATE KEY
                         │
             ┌───────────┴───────────┐
             │                       │
             ▼                       ▼
          Signing                Key Encoding
             │                       │
             ▼                       ▼
            CSR                  PKCS#8 / PKIX
             │
             ▼
       Certificate Authority
             │
             ▼
       X.509 Certificate
             │
       ┌─────┴─────┐
       │            │
       ▼            ▼
 Certificate      Certificate
   Chain          Properties
       │
       ▼
    CertPool
       │
       ▼
 Trusted Root CA
       │
       ▼
 Certificate.Verify()
       │
       ├── Signature
       ├── Validity
       ├── Chain
       ├── Key Usage
       ├── Extended Key Usage
       ├── Name Constraints
       ├── Policies
       └── Hostname
       │
       ▼
      TRUST
```

And remember this distinction:

```text
crypto/rsa
crypto/ecdsa
crypto/ed25519
        │
        ▼
Cryptographic operations


crypto/x509
        │
        ▼
PKI + certificates + chains +
CSRs + CRLs + standardized
key/certificate encodings
```

---

# 38. Thought-provoking question

Suppose you build a microservice system where **every service has an X.509 certificate signed by your private CA**.

Now imagine an attacker steals the private key of one service.

**If your certificate verification only checks that the certificate chains to your trusted CA, what security problem could arise—and how would you design your certificate hierarchy, `KeyUsage`, `ExtKeyUsage`, certificate policies, and revocation strategy so that compromising one service does not automatically give the attacker the ability to impersonate every other service?**

---

## Further study

For the complete current API surface, consult the official Go `crypto/x509` package documentation:

- https://pkg.go.dev/crypto/x509
- https://pkg.go.dev/crypto/tls
- https://pkg.go.dev/encoding/pem
