# Go `math/rand` Package — Detailed Guide

## 1. What is the `math/rand` package?

The Go **`math/rand`** package provides functions and types for generating **pseudo-random numbers**.

"Pseudo-random" means that the numbers look random, but they are actually produced by a deterministic algorithm.

### Simple example

```go
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(rand.Intn(100))
}
```

Possible output:

```text
73
```

### When is `math/rand` commonly used?

It is useful for:

- Simulations
- Games
- Randomized algorithms
- Generating test data
- Shuffling data
- Selecting random elements
- Statistical experiments
- Modeling random events
- Load/testing scenarios

The package is **not appropriate for passwords, authentication tokens, encryption keys, security codes, or other security-sensitive randomness**. For those situations, use `crypto/rand`.

> **Version note:** `math/rand` and `math/rand/v2` are different packages. This guide focuses on the classic `math/rand` API, while noting the modern Go direction where relevant.

---

# 2. Basic Example

```go
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Random integer from 0 through 99
	number := rand.Intn(100)

	fmt.Println("Random number:", number)

	// Random floating-point number from 0.0 up to but not including 1.0
	value := rand.Float64()

	fmt.Println("Random float:", value)

	// Select a random item
	colors := []string{"Red", "Green", "Blue", "Yellow"}

	index := rand.Intn(len(colors))

	fmt.Println("Random color:", colors[index])
}
```

Notice this pattern:

```go
rand.Intn(len(colors))
```

If there are 4 elements, `Intn(4)` produces:

```text
0
1
2
3
```

So it can safely be used as a slice index.

---

# 3. Important Concept: `[0, n)`

One of the most important things to understand about `math/rand` is the notation:

```text
[0, n)
```

It means:

- `0` **is included**
- `n` **is excluded**

For example:

```go
rand.Intn(10)
```

can produce:

```text
0 1 2 3 4 5 6 7 8 9
```

but **never 10**.

This is extremely useful when working with arrays and slices.

```go
names := []string{
	"Alice",
	"Bob",
	"Charlie",
	"David",
}

index := rand.Intn(len(names))

fmt.Println(names[index])
```

`len(names)` is `4`, so the generated index is always:

```text
0 <= index < 4
```

---

# 4. Functions in `math/rand`

There are quite a few functions and methods in this package. They can be grouped by purpose.

> **Version note:** Go's random-number APIs have evolved. Modern Go also provides `math/rand/v2`, which has a redesigned API. The classic `math/rand` package remains important when working with existing Go code and older APIs.

---

## A. `rand.Int()`

### Syntax

```go
rand.Int()
```

### Purpose

Returns a non-negative pseudo-random `int`.

### Example

```go
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	n := rand.Int()

	fmt.Println(n)
}
```

Possible output:

```text
384729182
```

The exact value depends on the random generator's state.

### When to use it

Use `Int()` when you need a random integer and **don't care about a specific upper limit**.

---

## B. `rand.Intn(n)`

This is probably the function you'll use most often as a beginner.

### Syntax

```go
rand.Intn(n)
```

### Purpose

Returns a random integer in:

```text
[0, n)
```

### Example

```go
number := rand.Intn(10)

fmt.Println(number)
```

Possible results:

```text
0
1
2
...
9
```

### Important

This is invalid:

```go
rand.Intn(0)
```

The argument must be greater than zero.

---

## C. `rand.Int31()`

### Syntax

```go
rand.Int31()
```

Returns a non-negative pseudo-random 31-bit integer as `int32`.

Example:

```go
n := rand.Int31()

fmt.Println(n)
```

Use this when you specifically need an `int32`-compatible random value.

---

## D. `rand.Int31n(n)`

### Syntax

```go
rand.Int31n(n)
```

Returns a random `int32` in:

```text
[0, n)
```

Example:

```go
n := rand.Int31n(100)

fmt.Println(n)
```

Possible values:

```text
0 through 99
```

The `n` argument must be positive.

---

## E. `rand.Int63()`

### Syntax

```go
rand.Int63()
```

Returns a non-negative pseudo-random 63-bit integer as `int64`.

Example:

```go
n := rand.Int63()

fmt.Println(n)
```

This is useful when working with larger integer values.

---

## F. `rand.Int63n(n)`

### Syntax

```go
rand.Int63n(n)
```

Returns a random `int64` in:

```text
[0, n)
```

Example:

```go
n := rand.Int63n(1000000)

fmt.Println(n)
```

Possible values:

```text
0 through 999999
```

---

## G. `rand.Float32()`

### Syntax

```go
rand.Float32()
```

Returns a pseudo-random `float32` in:

```text
[0.0, 1.0)
```

Example:

```go
value := rand.Float32()

fmt.Println(value)
```

Possible result:

```text
0.37482965
```

It can approach `1.0`, but does not return `1.0`.

### Useful for

- Probabilities
- Simulations
- Random percentages
- Game mechanics

For example:

```go
if rand.Float32() < 0.25 {
	fmt.Println("25% event happened")
}
```

---

## H. `rand.Float64()`

### Syntax

```go
rand.Float64()
```

Returns a random `float64` in:

```text
[0.0, 1.0)
```

Example:

```go
value := rand.Float64()

fmt.Println(value)
```

`Float64()` provides more precision than `Float32()`.

---

## I. `rand.Perm(n)`

This function is very useful.

### Syntax

```go
rand.Perm(n)
```

It returns a randomly shuffled permutation of:

```text
0, 1, 2, ..., n-1
```

Example:

```go
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(rand.Perm(5))
}
```

Possible output:

```text
[3 0 4 1 2]
```

Another execution could produce:

```text
[1 4 0 3 2]
```

Every number from `0` through `4` appears exactly once.

### Real-world use

Suppose you have:

```go
students := []string{
	"Alice",
	"Bob",
	"Charlie",
	"David",
	"Eve",
}
```

You could use a permutation to create a random ordering.

---

## J. `rand.Shuffle()`

### Syntax

```go
rand.Shuffle(n, swap)
```

This randomly rearranges elements.

Example:

```go
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	names := []string{
		"Alice",
		"Bob",
		"Charlie",
		"David",
	}

	rand.Shuffle(len(names), func(i, j int) {
		names[i], names[j] = names[j], names[i]
	})

	fmt.Println(names)
}
```

Possible output:

```text
[Charlie Alice David Bob]
```

### Understanding the callback

This part:

```go
func(i, j int) {
	names[i], names[j] = names[j], names[i]
}
```

tells `Shuffle` how to exchange two elements.

### Important

`Shuffle()` changes the original slice.

It does **not** create a new shuffled slice.

---

## K. `rand.Seed()`

Historically, `Seed()` was commonly used to initialize the package's global random generator.

Example from older Go code:

```go
rand.Seed(42)
```

A fixed seed makes the sequence reproducible.

For example:

```go
rand.Seed(42)

fmt.Println(rand.Intn(100))
fmt.Println(rand.Intn(100))
fmt.Println(rand.Intn(100))
```

Running the same program with the same seed produces the same sequence.

### Why is that useful?

Especially for:

- Tests
- Simulations
- Debugging
- Reproducing bugs

### Important modern-Go note

You will encounter older tutorials recommending:

```go
rand.Seed(time.Now().UnixNano())
```

Be careful when following older tutorials. The random APIs have changed over recent Go releases, and `math/rand/v2` is now the modern redesigned package.

---

## L. `rand.NewSource()`

`NewSource()` creates a new pseudo-random source using a seed.

Conceptually:

```go
source := rand.NewSource(42)
```

You can then create a `Rand` from it:

```go
r := rand.New(source)
```

Then:

```go
fmt.Println(r.Intn(100))
```

### Why create your own generator?

It gives you an independent random sequence.

This can be useful when different parts of your program need separate random generators.

---

## M. `rand.New()`

### Syntax

```go
rand.New(source)
```

It creates a new `*rand.Rand` using the supplied random source.

Example:

```go
source := rand.NewSource(42)

r := rand.New(source)

fmt.Println(r.Intn(100))
fmt.Println(r.Intn(100))
```

Now `r` has its own random-generator state.

---

## N. `Rand` methods

`math/rand` provides a `Rand` type whose methods correspond to many of the package-level functions.

For example:

```go
r := rand.New(rand.NewSource(42))

r.Int()
r.Intn(100)
r.Int31()
r.Int31n(100)
r.Int63()
r.Int63n(100)
r.Float32()
r.Float64()
r.Perm(10)
```

The major idea is:

```text
package-level function
        ↓
uses package/global generator

Rand method
        ↓
uses a specific Rand instance
```

This distinction becomes important in larger programs.

---

## O. `Rand.Seed()`

A `Rand` object can also be reseeded:

```go
r := rand.New(rand.NewSource(10))

r.Seed(42)
```

This resets the generator's sequence.

This is useful when you deliberately want deterministic/reproducible behavior.

---

## P. `rand.ExpFloat64()`

This is more advanced.

It returns a random `float64` following an **exponential distribution**.

```go
value := rand.ExpFloat64()

fmt.Println(value)
```

The default exponential distribution has:

```text
mean = 1
lambda = 1
```

This is useful for modeling things such as:

- Time between events
- Arrival processes
- Queueing simulations
- Reliability modeling

---

## Q. `rand.NormFloat64()`

Returns a random value following a **standard normal distribution**.

That means approximately:

```text
mean = 0
standard deviation = 1
```

Example:

```go
value := rand.NormFloat64()

fmt.Println(value)
```

You will often see values such as:

```text
-0.42
0.18
1.23
-1.07
```

Unlike `Intn()`, the result isn't limited to a small integer range.

### Creating another normal distribution

Suppose you want:

```text
mean = 100
standard deviation = 15
```

You can transform the result:

```go
value := rand.NormFloat64()*15 + 100
```

This is useful for simulations involving naturally distributed measurements.

---

# 5. Randomness vs Security

This is one of the **most important concepts** to learn.

Don't do this for a password:

```go
rand.Intn(...)
```

Don't use `math/rand` for:

- Password generation
- Authentication tokens
- Session tokens
- API secrets
- Encryption keys
- Security-sensitive IDs

Instead:

```go
import "crypto/rand"
```

The `crypto/rand` package is specifically designed to provide cryptographically secure randomness.

Think:

```text
math/rand
    ↓
simulation / games / testing

crypto/rand
    ↓
security / cryptography
```

---

# 6. Three Common Beginner Mistakes

## Mistake 1: Thinking `math/rand` is truly random

A common misconception is:

> "The computer doesn't know the next number."

Not exactly.

`math/rand` generates **pseudo-random** values using an algorithm.

The sequence is deterministic given the generator's state/seed.

### Avoid it

Understand the difference between:

```text
random
```

and:

```text
pseudo-random
```

This distinction becomes particularly important when writing tests and simulations.

---

## Mistake 2: Getting the range wrong

Beginners often write:

```go
rand.Intn(10)
```

and expect:

```text
1–10
```

But the result is:

```text
0–9
```

Remember:

```text
Intn(n) → [0,n)
```

If you specifically want 1–10:

```go
rand.Intn(10) + 1
```

---

## Mistake 3: Using `math/rand` for security

For example, someone might generate a password like:

```go
passwordNumber := rand.Intn(1000000)
```

This is **not appropriate for security**.

A predictable pseudo-random generator can potentially allow an attacker to predict values.

### Avoid it

Use:

```go
crypto/rand
```

for security-sensitive randomness.

---

# 7. Real-World Application #1 — Games

Randomness is everywhere in games.

For example, imagine a dice roll:

```go
dice := rand.Intn(6) + 1

fmt.Println("You rolled:", dice)
```

Possible results:

```text
1
2
3
4
5
6
```

You can use the same concept for:

- Enemy behavior
- Loot drops
- Random maps
- Card games
- Random encounters
- Critical hits
- NPC decisions

For example:

```go
if rand.Intn(100) < 20 {
	fmt.Println("Critical hit!")
}
```

This gives an event with an intended 20% probability.

---

# 8. Real-World Application #2 — Simulations

Suppose you're simulating customers arriving at a store.

You can generate random values to model:

```text
customer arrivals
service times
waiting times
demand
failures
network traffic
```

For example, you could use a random distribution to simulate how long customers wait.

This is one reason the package is useful for simulations.

---

# 9. Progressive Exercises

**No solutions are provided.**

## Exercise 1 — Random Dice Simulator ⭐

Create a program that simulates rolling a six-sided dice 10 times.

Requirements:

- Generate a value between 1 and 6.
- Print every roll.
- Count how many times each number occurred.
- At the end, display the frequency of each number.

Example structure:

```text
Roll 1: 4
Roll 2: 2
Roll 3: 6
...

1 appeared: X times
2 appeared: X times
...
6 appeared: X times
```

---

## Exercise 2 — Random Student Selector ⭐⭐

Create a program containing a list of at least 10 students.

Your program should:

1. Randomly select one student.
2. Print the selected student's name.
3. Randomly shuffle the entire student list.
4. Display the shuffled list.
5. Select three different students for a team.

Try to make sure the same student isn't selected twice for the three-person team.

---

## Exercise 3 — Monte Carlo Probability Simulation ⭐⭐⭐

Build a simulation that estimates the probability of getting **at least one six when rolling a six-sided die multiple times**.

Your program should:

1. Ask the user how many dice rolls should happen in each experiment.
2. Run thousands of experiments.
3. In each experiment, roll the die the requested number of times.
4. Determine whether at least one six occurred.
5. Count successful experiments.
6. Calculate the estimated probability.
7. Display the result as a percentage.

Then experiment with different numbers of rolls and observe how the probability changes.

This exercise will take you beyond simply calling random functions and make you think about **probability, simulation, and statistical convergence**.

---

# 10. Useful Mental Model

When learning `math/rand`, think of it like this:

```text
                 math/rand
                     │
        ┌────────────┼─────────────┐
        │            │             │
     Numbers       Floats       Ordering
        │            │             │
    Int / Intn   Float32/64   Perm / Shuffle
        │
        ├── Int31 / Int31n
        │
        └── Int63 / Int63n

        Advanced distributions
                 │
          ┌──────┴──────┐
          │             │
     NormFloat64   ExpFloat64
```

And remember the fundamental distinction:

```text
math/rand
    ↓
pseudo-random
    ↓
simulation / games / testing

crypto/rand
    ↓
cryptographically secure
    ↓
security-sensitive applications
```

The current Go ecosystem also has **`math/rand/v2`**, which modernizes the API and adds newer generators such as PCG and ChaCha8. The modern package includes functions such as `IntN`, `N`, `Perm`, `Shuffle`, `NormFloat64`, and `ExpFloat64`, plus `Rand`, `PCG`, `ChaCha8`, and `Zipf` types.

## Official Documentation

- Go `math/rand`: https://pkg.go.dev/math/rand
- Go `math/rand/v2`: https://pkg.go.dev/math/rand/v2
- Go `crypto/rand`: https://pkg.go.dev/crypto/rand

---

# 11. Thought-Provoking Question

Suppose you're building an online game where a player receives a rare item with a **1% drop probability**.

If you use `math/rand` to generate that probability, **how would you prove that the item really has approximately a 1% chance of dropping over millions of attempts—and how might a biased random-number implementation affect the fairness of the game?**

That question gets you from simply *using* `math/rand` to understanding **randomness, probability distributions, statistical testing, and fairness**.
