# Go `crypto/rand` Package

## 1. What is `crypto/rand`?

The Go `crypto/rand` package provides a **cryptographically secure random number generator (CSPRNG)**.

```go
import "crypto/rand"
```

Its purpose is to generate random data that should be **difficult for an attacker to predict**.

It is commonly used for:

- Cryptographic keys
- Authentication/session tokens
- Password-reset tokens
- API keys
- Nonces
- Random IDs
- Security-sensitive random numbers
- Cryptographic prime numbers
- Security-sensitive protocols

The package's global `rand.Reader` obtains randomness from the operating system's secure random facilities.

### `crypto/rand` vs `math/rand`

| Package | Purpose | Security |
|---|---|---|
| `math/rand` | Simulations, games, randomized algorithms | ❌ Not cryptographically secure |
| `crypto/rand` | Keys, tokens, security-sensitive randomness | ✅ Cryptographically secure |

**Rule of thumb:**

> If predicting the random value could cause a security problem, use `crypto/rand`.

---

# 2. Every function in `crypto/rand`

The current standard library exposes these package-level functions:

```text
Int
Prime
Read
Text
```

It also exposes:

```go
rand.Reader
```

which is a package-level value implementing `io.Reader`.

---

## 2.1 `rand.Reader`

### Definition

```go
var Reader io.Reader
```

`Reader` is the shared cryptographically secure random source used by the package.

You will frequently see:

```go
rand.Reader
```

passed to functions such as:

```go
rand.Int(rand.Reader, max)
```

and:

```go
rand.Prime(rand.Reader, 2048)
```

The reader is safe for concurrent use.

### Think of it as

```text
Operating System
       ↓
 Secure randomness
       ↓
 crypto/rand.Reader
       ↓
 Your Go program
```

You normally **do not need to create your own random source**.

---

# 2.2 `rand.Read`

### Function signature

```go
func Read(b []byte) (n int, err error)
```

`Read` fills a byte slice with cryptographically secure random bytes. When it succeeds, the current implementation fills the entire slice.

### Example

```go
package main

import (
	"crypto/rand"
	"fmt"
)

func main() {
	data := make([]byte, 16)

	_, err := rand.Read(data)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%x\n", data)
}
```

Possible output:

```text
9a31c7e2b8450d8f1c6a92e37b5104aa
```

The output changes every time.

### What does `[]byte` mean?

Suppose:

```go
data := make([]byte, 16)
```

You created 16 bytes.

After:

```go
rand.Read(data)
```

you might get:

```text
9a 31 c7 e2 b8 45 0d 8f
1c 6a 92 e3 7b 51 04 aa
```

Each byte can contain values from:

```text
0 → 255
```

### Common use

Generating random keys or binary secrets:

```go
key := make([]byte, 32)

_, err := rand.Read(key)
if err != nil {
	panic(err)
}
```

A 32-byte value contains **256 bits** of random data.

---

# 2.3 `rand.Int`

### Function signature

```go
func Int(rand io.Reader, max *big.Int) (*big.Int, error)
```

It returns a **uniformly distributed cryptographically secure random integer** in:

```text
[0, max)
```

The upper bound is **exclusive**.

For example:

```go
max := big.NewInt(100)

n, err := rand.Int(rand.Reader, max)
```

`n` can be:

```text
0
1
2
...
98
99
```

but never:

```text
100
```

### Complete example

```go
package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	max := big.NewInt(100)

	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		panic(err)
	}

	fmt.Println(n)
}
```

### Why `*big.Int`?

Cryptographic applications frequently need integers much larger than Go's normal:

```text
int
int64
uint64
```

Cryptographic systems may work with hundreds or thousands of bits.

That's why `crypto/rand.Int` uses:

```go
*big.Int
```

---

# 2.4 `rand.Prime`

### Function signature

```go
func Prime(r io.Reader, bits int) (*big.Int, error)
```

`Prime` generates a cryptographically random number of the requested bit length that is prime with high probability. It returns an error if `bits < 2` or if randomness generation fails.

### Example

```go
package main

import (
	"crypto/rand"
	"fmt"
)

func main() {
	p, err := rand.Prime(rand.Reader, 128)
	if err != nil {
		panic(err)
	}

	fmt.Println(p)
}
```

You might get a value such as:

```text
309485009821345068724781371
```

The exact number varies.

### What does `128` mean?

It means you want a prime approximately **128 bits long**.

For example:

```go
rand.Prime(rand.Reader, 256)
```

or:

```go
rand.Prime(rand.Reader, 2048)
```

Prime generation is particularly relevant to cryptographic algorithms that rely on large prime numbers.

---

# 2.5 `rand.Text`

### Function signature

```go
func Text() string
```

`Text` was added in **Go 1.24**.

It generates a cryptographically random string using the standard RFC 4648 Base32 alphabet. The result contains at least **128 bits of randomness**, making it suitable for secrets, tokens, and other security-sensitive random text.

### Example

```go
package main

import (
	"crypto/rand"
	"fmt"
)

func main() {
	token := rand.Text()

	fmt.Println(token)
}
```

Example output:

```text
KRSXG5A7MFRWQ2L6N4ZQ====
```

The actual output will be different.

### Why is `Text()` convenient?

Without `Text`, you might generate random bytes and then encode them:

```go
bytes := make([]byte, 32)

_, err := rand.Read(bytes)
if err != nil {
	panic(err)
}
```

Then convert them to text using something like:

```go
hex.EncodeToString(bytes)
```

or Base64.

`rand.Text()` provides a convenient cryptographically secure random text value directly.

---

# 3. Simple complete example

Here is a small program demonstrating several functions together:

```go
package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	// 1. Generate random bytes
	data := make([]byte, 16)

	_, err := rand.Read(data)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Random bytes: %x\n", data)

	// 2. Generate a secure random integer from 0 to 99
	n, err := rand.Int(rand.Reader, big.NewInt(100))
	if err != nil {
		panic(err)
	}

	fmt.Println("Random number:", n)

	// 3. Generate a random prime
	prime, err := rand.Prime(rand.Reader, 128)
	if err != nil {
		panic(err)
	}

	fmt.Println("Random prime:", prime)

	// 4. Generate random text
	token := rand.Text()

	fmt.Println("Random text:", token)
}
```

This demonstrates the four current package functions:

- `Read`
- `Int`
- `Prime`
- `Text`

---

# 4. Three common beginner mistakes

## Mistake 1: Using `math/rand` for security

A beginner might write:

```go
import "math/rand"
```

and use:

```go
rand.Intn(100)
```

for a password-reset token or authentication token.

That's inappropriate because `math/rand` is intended for ordinary pseudorandom operations rather than cryptographic security.

### Better

```go
import "crypto/rand"
```

For example:

```go
n, err := rand.Int(rand.Reader, big.NewInt(100))
```

---

## Mistake 2: Thinking `crypto/rand` means "true randomness"

`crypto/rand` is commonly described as cryptographically secure randomness, but you shouldn't interpret that as magical or philosophically "perfect randomness."

It provides randomness through secure operating-system mechanisms and is designed to make generated values computationally unpredictable.

The important security property is:

> An attacker should not be able to practically predict future random values from the information available to them.

---

## Mistake 3: Ignoring errors

For example:

```go
data := make([]byte, 32)

rand.Read(data)
```

Although the current `crypto/rand.Read` implementation is designed to fill the buffer and the standard `Reader` is expected to provide secure randomness, beginners should still understand that the function returns an error:

```go
_, err := rand.Read(data)

if err != nil {
	// handle failure
}
```

For learning Go, **always pay attention to returned errors**.

---

# 5. Two real-world applications

## Application 1: Password-reset tokens

Imagine a website where a user clicks:

> "Forgot password?"

The server needs to generate a token that an attacker cannot guess.

Conceptually:

```go
token := rand.Text()
```

The server can associate that token with the user's password-reset request.

The important property is that an attacker shouldn't be able to predict the next valid token.

---

## Application 2: Cryptographic keys

Suppose an application needs a random 256-bit secret:

```go
key := make([]byte, 32)

_, err := rand.Read(key)
if err != nil {
	panic(err)
}
```

You now have 32 cryptographically random bytes.

This type of operation is useful when generating cryptographic key material or other high-entropy secrets.

**Important:** generating random bytes is only one part of secure cryptography. How the key is stored, used, rotated, and protected also matters.

---

# 6. Three exercises

## Exercise 1 — Random Security Token ⭐

Write a Go program that:

1. Creates a 32-byte slice.
2. Uses `crypto/rand.Read()` to fill it.
3. Converts the bytes into a readable hexadecimal string.
4. Prints the resulting token.

**Goal:** Become comfortable generating secure random bytes and converting them to text.

---

## Exercise 2 — Secure Random Number Challenge ⭐⭐

Create a program that:

1. Uses `crypto/rand.Int()`.
2. Generates a random number between `1` and `100`.
3. Asks the user to guess the number.
4. Continues until the user guesses correctly.
5. Counts the number of attempts.

**Goal:** Practice combining `crypto/rand.Int`, `math/big`, user input, loops, and error handling.

**Think carefully about:** the difference between `[0, max)` and an inclusive range such as `[1, 100]`.

---

## Exercise 3 — Cryptographic Prime Generator ⭐⭐⭐

Create a program that:

1. Asks the user for a bit size.
2. Uses `crypto/rand.Prime()` to generate a prime of that size.
3. Prints the generated prime.
4. Reports the number of bits in the generated value.
5. Validates that the generated number is actually prime using appropriate functionality from `math/big`.
6. Handles invalid bit sizes and errors gracefully.

**Goal:** Understand how secure randomness can be used as a building block for cryptographic mathematics.

---

# 7. Quick reference

| Function | Purpose | Result |
|---|---|---|
| `rand.Read(b)` | Generate secure random bytes | `[]byte` |
| `rand.Int(r, max)` | Generate secure random integer | `*big.Int` |
| `rand.Prime(r, bits)` | Generate secure random prime | `*big.Int` |
| `rand.Text()` | Generate secure random text | `string` |
| `rand.Reader` | Secure randomness source | `io.Reader` |

### One concept to remember

```text
crypto/rand
     │
     ├── Read()   → random bytes
     ├── Int()    → random integer
     ├── Prime()  → random prime
     └── Text()   → random secure text
```

And the most important distinction:

```text
math/rand   → ordinary randomness
crypto/rand → security-sensitive randomness
```

---

# 8. Thought-provoking question

Imagine you are building a password-reset system. You generate a token using `crypto/rand`, but then store that token in your database **exactly as generated**.

**Is the system automatically secure just because the token came from `crypto/rand`?**

Think about:

- What could happen if an attacker obtains the database?
- How should token expiration work?
- Should tokens be invalidated after use?
- Should tokens be stored in plaintext?
- What other parts of the system could undermine the security provided by `crypto/rand`?

---

## Official Documentation

Go `crypto/rand` package documentation:

https://pkg.go.dev/crypto/rand
