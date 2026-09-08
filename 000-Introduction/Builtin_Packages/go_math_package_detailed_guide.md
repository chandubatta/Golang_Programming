# Go `math` Package — Detailed Guide

## 1. What is the `math` package?

Go's `math` package provides mathematical operations primarily for `float64` values.

Import it with:

```go
import "math"
```

It is useful for:

- Square roots
- Powers
- Logarithms
- Trigonometry
- Rounding
- Absolute values
- Exponential calculations
- Floating-point manipulation
- `NaN` and infinity handling
- Gamma/Bessel functions
- Floating-point bit manipulation

The package is part of Go's standard library, so no external dependency is required.

Official documentation: https://pkg.go.dev/math

---

# 2. Simple Example

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	radius := 5.0

	area := math.Pi * math.Pow(radius, 2)
	circumference := 2 * math.Pi * radius

	fmt.Println("Area:", area)
	fmt.Println("Circumference:", circumference)

	fmt.Println("Square root:", math.Sqrt(25))
	fmt.Println("Absolute value:", math.Abs(-10.5))
	fmt.Println("Rounded:", math.Round(10.6))
}
```

Output:

```text
Area: 78.53981633974483
Circumference: 31.41592653589793
Square root: 5
Absolute value: 10.5
Rounded: 11
```

Important:

```go
math.Sqrt(25)
```

returns a `float64`, even though `25` looks like an integer.

---

# 3. Important `math` Constants

## Mathematical Constants

| Constant | Meaning |
|---|---|
| `math.Pi` | π ≈ 3.14159 |
| `math.E` | Euler's number ≈ 2.71828 |
| `math.Phi` | Golden ratio ≈ 1.61803 |
| `math.Sqrt2` | √2 |
| `math.SqrtE` | √e |
| `math.SqrtPi` | √π |
| `math.SqrtPhi` | √φ |
| `math.Ln2` | ln(2) |
| `math.Ln10` | ln(10) |
| `math.Log2E` | log₂(e) |
| `math.Log10E` | log₁₀(e) |

## Floating-Point Limits

```go
math.MaxFloat32
math.SmallestNonzeroFloat32

math.MaxFloat64
math.SmallestNonzeroFloat64
```

For example:

```go
fmt.Println(math.MaxFloat64)
```

is approximately:

```text
1.7976931348623157e+308
```

## Integer Limits

The package also exposes integer limits such as:

```go
math.MaxInt
math.MinInt

math.MaxInt8
math.MinInt8

math.MaxInt16
math.MinInt16

math.MaxInt32
math.MinInt32

math.MaxInt64
math.MinInt64

math.MaxUint
math.MaxUint8
math.MaxUint16
math.MaxUint32
math.MaxUint64
```

These constants describe the limits of the corresponding integer types.

---

# 4. Every Function in the `math` Package

The functions are grouped below by purpose.

---

## A. Absolute Value

### `math.Abs()`

```go
math.Abs(x)
```

Returns the absolute value.

```go
fmt.Println(math.Abs(-25.5))
```

Output:

```text
25.5
```

Think:

```text
-25 → 25
 25 → 25
```

Useful for calculating distances or differences where the sign does not matter.

---

# B. Trigonometric Functions

Go's trigonometric functions use **radians**, not degrees.

## `math.Sin()`

```go
math.Sin(x)
```

Returns the sine of `x`.

```go
math.Sin(math.Pi / 2)
```

Approximately:

```text
1
```

---

## `math.Cos()`

```go
math.Cos(x)
```

Returns cosine.

```go
math.Cos(0)
```

Output:

```text
1
```

---

## `math.Tan()`

```go
math.Tan(x)
```

Returns tangent.

```go
math.Tan(math.Pi / 4)
```

Approximately:

```text
1
```

---

# C. Inverse Trigonometric Functions

## `math.Asin()`

```go
math.Asin(x)
```

Returns inverse sine in radians.

```go
angle := math.Asin(1)
```

Result:

```text
π/2
```

The input normally needs to be between `-1` and `+1`. Otherwise the result is `NaN`.

---

## `math.Acos()`

```go
math.Acos(x)
```

Returns inverse cosine in radians.

```go
math.Acos(1)
```

Result:

```text
0
```

---

## `math.Atan()`

```go
math.Atan(x)
```

Returns inverse tangent in radians.

```go
math.Atan(1)
```

Result:

```text
π/4
```

---

## `math.Atan2()`

```go
math.Atan2(y, x)
```

Calculates an angle from an `(x, y)` coordinate while correctly determining the quadrant.

```go
angle := math.Atan2(10, 20)
```

Common applications:

- Robotics
- Games
- Navigation
- GPS calculations
- Direction calculations
- Graphics

`Atan2` is usually preferable to simply doing:

```go
math.Atan(y / x)
```

because `Atan2` uses the signs of both coordinates to determine the correct quadrant.

---

# D. Hyperbolic Functions

## `math.Sinh()`

Returns the hyperbolic sine.

```go
math.Sinh(x)
```

## `math.Cosh()`

Returns the hyperbolic cosine.

```go
math.Cosh(x)
```

## `math.Tanh()`

Returns the hyperbolic tangent.

```go
math.Tanh(x)
```

## `math.Asinh()`

Returns inverse hyperbolic sine.

```go
math.Asinh(x)
```

## `math.Acosh()`

Returns inverse hyperbolic cosine.

```go
math.Acosh(x)
```

For real-valued results, the input must be at least `1`.

## `math.Atanh()`

Returns inverse hyperbolic tangent.

```go
math.Atanh(x)
```

For real results, `x` must be within `-1` and `1`.

---

# E. Powers and Roots

## `math.Sqrt()`

Calculates the square root.

```go
math.Sqrt(25)
```

Result:

```text
5
```

Example:

```go
distance := math.Sqrt(3*3 + 4*4)
```

Result:

```text
5
```

---

## `math.Cbrt()`

Calculates the cube root.

```go
math.Cbrt(27)
```

Result:

```text
3
```

It can also handle negative values:

```go
math.Cbrt(-27)
```

Result:

```text
-3
```

---

## `math.Pow()`

Calculates:

```text
x^y
```

Example:

```go
result := math.Pow(2, 10)
fmt.Println(result)
```

Output:

```text
1024
```

---

## `math.Pow10()`

Calculates:

```text
10^n
```

Example:

```go
fmt.Println(math.Pow10(3))
```

Output:

```text
1000
```

---

# F. Exponential Functions

## `math.Exp()`

Calculates:

```text
e^x
```

Example:

```go
fmt.Println(math.Exp(1))
```

Approximately:

```text
2.71828
```

Applications include:

- Probability
- Statistics
- Finance
- Scientific computing
- Machine learning

---

## `math.Exp2()`

Calculates:

```text
2^x
```

Example:

```go
math.Exp2(10)
```

Result:

```text
1024
```

---

## `math.Expm1()`

Calculates:

```text
e^x - 1
```

Conceptually similar to:

```go
math.Exp(x) - 1
```

but designed to provide better accuracy when `x` is very close to zero.

---

# G. Logarithmic Functions

## `math.Log()`

Calculates the natural logarithm:

```text
ln(x)
```

Example:

```go
math.Log(math.E)
```

Approximately:

```text
1
```

---

## `math.Log2()`

Calculates the base-2 logarithm.

```go
math.Log2(8)
```

Result:

```text
3
```

Because:

```text
2³ = 8
```

Useful for:

- Computer science
- Binary trees
- Algorithms
- Bit-related calculations
- Information theory

---

## `math.Log10()`

Calculates base-10 logarithm.

```go
math.Log10(1000)
```

Result:

```text
3
```

---

## `math.Log1p()`

Calculates:

```text
ln(1 + x)
```

It is especially useful when `x` is very small because it maintains better numerical accuracy than:

```go
math.Log(1 + x)
```

---

## `math.Logb()`

Returns the binary exponent of a floating-point number.

```go
math.Logb(8)
```

Approximately:

```text
3
```

This is a lower-level floating-point operation and is less commonly needed in ordinary application development.

---

# H. Rounding Functions

## `math.Floor()`

Rounds toward negative infinity.

```go
math.Floor(3.9)   // 3
math.Floor(-3.9)  // -4
```

---

## `math.Ceil()`

Rounds toward positive infinity.

```go
math.Ceil(3.1)   // 4
math.Ceil(-3.1)  // -3
```

---

## `math.Round()`

Rounds to the nearest integer, with halfway cases rounded away from zero.

```go
math.Round(10.5)   // 11
math.Round(-10.5)  // -11
```

---

## `math.RoundToEven()`

Rounds to the nearest integer, with halfway values going toward the nearest even integer.

```go
math.RoundToEven(10.5) // 10
math.RoundToEven(11.5) // 12
```

This is also known as **banker's rounding**.

---

## `math.Trunc()`

Removes the fractional portion.

```go
math.Trunc(10.99)   // 10
math.Trunc(-10.99)  // -10
```

This differs from `Floor()` for negative numbers.

---

# I. Minimum, Maximum, and Difference

## `math.Max()`

Returns the larger value.

```go
math.Max(10, 20)
```

Result:

```text
20
```

---

## `math.Min()`

Returns the smaller value.

```go
math.Min(10, 20)
```

Result:

```text
10
```

---

## `math.Dim()`

Returns:

```text
max(x-y, 0)
```

Example:

```go
math.Dim(10, 4) // 6
math.Dim(4, 10) // 0
```

Useful when you need a non-negative difference.

---

# J. Floating-Point Remainder

## `math.Mod()`

Returns the floating-point remainder.

```go
math.Mod(10.5, 3)
```

Result:

```text
1.5
```

It is conceptually similar to `%`, but works with floating-point numbers.

---

## `math.Remainder()`

Returns the IEEE 754 floating-point remainder.

```go
math.Remainder(10, 3)
```

This differs from `math.Mod()` in how the quotient/remainder is determined.

Do not assume `Mod()` and `Remainder()` are interchangeable.

---

## `math.Modf()`

Splits a floating-point number into:

1. Integer portion
2. Fractional portion

Example:

```go
integer, fraction := math.Modf(12.345)

fmt.Println(integer)
fmt.Println(fraction)
```

Conceptually:

```text
12
0.345
```

---

# K. Hypotenuse

## `math.Hypot()`

Calculates:

```text
sqrt(x² + y²)
```

Example:

```go
distance := math.Hypot(3, 4)

fmt.Println(distance)
```

Output:

```text
5
```

It is useful for Euclidean distance and geometry calculations.

---

# L. Fused Multiply-Add

## `math.FMA()`

FMA means **Fused Multiply-Add**.

```go
math.FMA(x, y, z)
```

It calculates:

```text
x*y + z
```

with a fused operation designed to reduce intermediate rounding.

Example:

```go
result := math.FMA(2, 3, 4)
```

Conceptually:

```text
2 × 3 + 4
= 10
```

Useful in:

- Scientific computing
- Numerical algorithms
- Simulations
- High-precision floating-point calculations

---

# M. Floating-Point Decomposition

## `math.Frexp()`

Breaks a floating-point number into:

```text
fraction × 2^exponent
```

Example:

```go
frac, exp := math.Frexp(8)

fmt.Println(frac)
fmt.Println(exp)
```

Conceptually:

```text
8 = 0.5 × 2⁴
```

Therefore:

```text
frac = 0.5
exp = 4
```

---

## `math.Ldexp()`

Performs the reverse operation.

```go
math.Ldexp(frac, exp)
```

It calculates:

```text
frac × 2^exp
```

Example:

```go
math.Ldexp(0.5, 4)
```

Result:

```text
8
```

So:

```text
Frexp ↔ Ldexp
```

are roughly inverse operations.

---

## `math.Ilogb()`

Returns the binary exponent of a floating-point number as an integer.

```go
math.Ilogb(8)
```

Approximately:

```text
3
```

Primarily useful for low-level floating-point work.

---

# N. Floating-Point Bit Conversion

These functions allow you to inspect IEEE-754 bit representations.

## `math.Float64bits()`

Converts a `float64` to its raw `uint64` representation.

```go
bits := math.Float64bits(3.14)
```

---

## `math.Float64frombits()`

Performs the reverse conversion:

```go
value := math.Float64frombits(bits)
```

Conceptually:

```text
float64
   ↓
Float64bits
   ↓
uint64
```

and:

```text
uint64
   ↓
Float64frombits
   ↓
float64
```

---

## `math.Float32bits()`

Converts:

```text
float32 → uint32
```

Example:

```go
bits := math.Float32bits(3.14)
```

---

## `math.Float32frombits()`

Converts:

```text
uint32 → float32
```

Useful for:

- Serialization
- Binary protocols
- Floating-point analysis
- Low-level systems programming

---

# O. Special Floating-Point Values

## `math.Inf()`

Creates positive or negative infinity.

```go
positive := math.Inf(1)
negative := math.Inf(-1)
```

---

## `math.IsInf()`

Checks whether a value is infinity.

```go
x := math.Inf(1)

fmt.Println(math.IsInf(x, 0))
```

Output:

```text
true
```

The second argument controls the sign:

```text
 1 → positive infinity
-1 → negative infinity
 0 → either infinity
```

---

## `math.NaN()`

Creates a NaN value.

NaN means:

```text
Not a Number
```

Example:

```go
x := math.NaN()
```

---

## `math.IsNaN()`

Checks whether a value is NaN.

```go
x := math.NaN()

fmt.Println(math.IsNaN(x))
```

Output:

```text
true
```

Important:

Do not use:

```go
x == math.NaN()
```

to detect NaN.

Use:

```go
math.IsNaN(x)
```

instead.

---

## `math.Signbit()`

Checks whether a floating-point value has a negative sign.

```go
math.Signbit(-10) // true
math.Signbit(10)  // false
```

This is particularly useful because floating-point values can distinguish between positive and negative zero.

---

## `math.Copysign()`

Copies the sign of one number onto another.

```go
math.Copysign(10, -1)  // -10
math.Copysign(-10, 1)  // 10
```

---

# P. Floating-Point Neighboring Values

## `math.Nextafter()`

Returns the next representable `float64` value moving from `x` toward `y`.

```go
next := math.Nextafter(1.0, 2.0)
```

Useful for:

- Floating-point precision
- Numerical algorithms
- Boundary conditions
- Testing floating-point behavior

---

## `math.Nextafter32()`

Same concept for `float32`.

```go
next := math.Nextafter32(float32(1.0), float32(2.0))
```

---

# Q. Gamma Functions

## `math.Gamma()`

Returns the Gamma function:

```text
Γ(x)
```

For positive integers:

```text
Γ(n) = (n-1)!
```

Example:

```go
math.Gamma(5)
```

returns:

```text
24
```

because:

```text
4! = 24
```

Applications:

- Probability distributions
- Statistics
- Combinatorics
- Bayesian mathematics
- Scientific computing

---

## `math.Lgamma()`

Returns the natural logarithm of the absolute value of the Gamma function and its sign.

```go
lgamma, sign := math.Lgamma(x)
```

Useful when Gamma values become extremely large and calculating Gamma directly could overflow.

---

# R. Error Functions

## `math.Erf()`

Calculates the error function:

```text
erf(x)
```

It appears in probability distributions and statistics.

---

## `math.Erfc()`

Calculates the complementary error function:

```text
erfc(x) = 1 - erf(x)
```

---

## `math.Erfinv()`

Calculates the inverse error function.

Conceptually:

```text
Erf(Erfinv(x)) ≈ x
```

---

## `math.Erfcinv()`

Calculates the inverse complementary error function.

These functions are generally used in statistical and scientific software.

---

# S. Bessel Functions

Bessel functions occur in problems involving cylindrical symmetry, waves, heat transfer, vibration, electromagnetics, and physics.

## `math.J0()`

Bessel function of the first kind, order 0.

```go
math.J0(x)
```

---

## `math.J1()`

Bessel function of the first kind, order 1.

```go
math.J1(x)
```

---

## `math.Jn()`

Bessel function of the first kind, order `n`.

```go
math.Jn(n, x)
```

---

## `math.Y0()`

Bessel function of the second kind, order 0.

```go
math.Y0(x)
```

---

## `math.Y1()`

Bessel function of the second kind, order 1.

```go
math.Y1(x)
```

---

## `math.Yn()`

Bessel function of the second kind, order `n`.

```go
math.Yn(n, x)
```

These are advanced scientific-computing functions rather than everyday backend-development functions.

---

# 5. Function Cheat Sheet

| Function | Purpose |
|---|---|
| `Abs` | Absolute value |
| `Acos` | Inverse cosine |
| `Acosh` | Inverse hyperbolic cosine |
| `Asin` | Inverse sine |
| `Asinh` | Inverse hyperbolic sine |
| `Atan` | Inverse tangent |
| `Atan2` | Angle from x/y coordinates |
| `Atanh` | Inverse hyperbolic tangent |
| `Cbrt` | Cube root |
| `Ceil` | Round upward |
| `Copysign` | Copy sign |
| `Cos` | Cosine |
| `Cosh` | Hyperbolic cosine |
| `Dim` | Positive difference |
| `Erf` | Error function |
| `Erfc` | Complementary error function |
| `Erfcinv` | Inverse complementary error function |
| `Erfinv` | Inverse error function |
| `Exp` | `e^x` |
| `Exp2` | `2^x` |
| `Expm1` | `e^x - 1` |
| `FMA` | Fused multiply-add |
| `Float32bits` | float32 → bits |
| `Float32frombits` | bits → float32 |
| `Float64bits` | float64 → bits |
| `Float64frombits` | bits → float64 |
| `Floor` | Round downward |
| `Frexp` | Decompose float |
| `Gamma` | Gamma function |
| `Hypot` | Hypotenuse |
| `Ilogb` | Binary exponent |
| `Inf` | Create infinity |
| `IsInf` | Check infinity |
| `IsNaN` | Check NaN |
| `J0` | Bessel J, order 0 |
| `J1` | Bessel J, order 1 |
| `Jn` | Bessel J, order n |
| `Ldexp` | Reconstruct float |
| `Lgamma` | Log Gamma |
| `Log` | Natural logarithm |
| `Log1p` | `ln(1+x)` |
| `Log2` | Base-2 logarithm |
| `Log10` | Base-10 logarithm |
| `Logb` | Binary exponent |
| `Max` | Maximum |
| `Min` | Minimum |
| `Mod` | Floating remainder |
| `Modf` | Integer/fraction split |
| `NaN` | Create NaN |
| `Nextafter` | Next float64 |
| `Nextafter32` | Next float32 |
| `Pow` | Power |
| `Pow10` | `10^n` |
| `Remainder` | IEEE floating remainder |
| `Round` | Nearest integer |
| `RoundToEven` | Banker's rounding |
| `Signbit` | Check negative sign |
| `Sin` | Sine |
| `Sincos` | Sine + cosine together |
| `Sinh` | Hyperbolic sine |
| `Sqrt` | Square root |
| `Tan` | Tangent |
| `Tanh` | Hyperbolic tangent |
| `Trunc` | Remove fractional part |
| `Y0` | Bessel Y, order 0 |
| `Y1` | Bessel Y, order 1 |
| `Yn` | Bessel Y, order n |

---

# 6. `math.Sincos()` — Important Function

```go
math.Sincos(x)
```

It returns both sine and cosine:

```go
sinValue, cosValue := math.Sincos(angle)
```

Instead of:

```go
sinValue := math.Sin(angle)
cosValue := math.Cos(angle)
```

you can request both values together.

This is useful for algorithms that need both sine and cosine values.

---

# 7. Three Common Beginner Mistakes

## Mistake 1: Thinking Trigonometric Functions Use Degrees

Beginners may write:

```go
math.Sin(90)
```

thinking this means 90 degrees.

Go's trigonometric functions use **radians**.

Convert degrees to radians:

```go
degrees := 90.0

radians := degrees * math.Pi / 180

result := math.Sin(radians)
```

Result:

```text
1
```

---

## Mistake 2: Confusing `Floor`, `Ceil`, `Round`, and `Trunc`

For:

```text
3.7
```

you get:

```text
Floor → 3
Ceil  → 4
Round → 4
Trunc → 3
```

For negative numbers:

```text
-3.7

Floor → -4
Ceil  → -3
Round → -4
Trunc → -3
```

Choose the rounding function based on the mathematical behavior you actually need.

---

## Mistake 3: Ignoring Floating-Point Precision

Do not assume:

```go
0.1 + 0.2 == 0.3
```

will necessarily behave like exact decimal arithmetic.

Floating-point numbers use binary representation, so some decimal fractions cannot be represented exactly.

Avoid raw floating-point equality when exact decimal behavior is required.

For financial applications, consider whether `float64` is appropriate instead of automatically using the `math` package.

---

# 8. Two Real-World Applications

## Application 1: GPS, Navigation, and Maps

Suppose you have two coordinate points and need to calculate a direction.

You might use:

```go
angle := math.Atan2(deltaY, deltaX)
```

Then convert the result to degrees:

```go
degrees := angle * 180 / math.Pi
```

`Atan2`, `Sin`, `Cos`, and `Hypot` can all appear in navigation and geometry calculations.

---

## Application 2: Financial and Scientific Calculations

Suppose you calculate compound growth:

```text
A = P × e^(rt)
```

Go can implement the exponential portion with:

```go
growth := math.Exp(rate * time)
```

Other mathematical applications include:

- Statistical models
- Probability calculations
- Simulations
- Scientific measurements
- Engineering calculations
- Signal processing
- Machine-learning algorithms

---

# 9. Three Progressively Challenging Exercises

## Exercise 1 — Beginner: Circle Calculator

Create a Go program that asks the user for the radius of a circle and calculates:

- Diameter
- Circumference
- Area

Use appropriate functions/constants from the `math` package.

Your program should display the results rounded to two decimal places.

**Do not use a solution from this document; implement it yourself.**

---

## Exercise 2 — Intermediate: Distance and Direction

Create a program that receives two points:

```text
Point A: (x1, y1)
Point B: (x2, y2)
```

Calculate:

1. The horizontal difference.
2. The vertical difference.
3. The distance between the points.
4. The direction/angle from Point A to Point B.

Use:

- `math.Hypot()`
- `math.Atan2()`

Display the angle in **degrees**, even though Go's trigonometric functions work with radians.

---

## Exercise 3 — Advanced: Floating-Point Numerical Analysis

Create a program that investigates floating-point precision.

Your program should:

1. Calculate a mathematical expression involving very small floating-point values.
2. Compare `math.Log(1+x)` against `math.Log1p(x)` for increasingly small values of `x`.
3. Compare `math.Exp(x)-1` against `math.Expm1(x)`.
4. Display the differences between the approaches.
5. Experiment with `math.Nextafter()` to determine neighboring representable floating-point values around a chosen number.

The goal is not simply to get an answer. The goal is to understand **why numerical functions such as `Log1p`, `Expm1`, and `Nextafter` exist**.

---

# 10. Recommended Learning Strategy

You do **not** need to memorize every function.

For normal Go backend development, prioritize these first:

```text
Abs
Min
Max
Sqrt
Pow
Cbrt

Floor
Ceil
Round
RoundToEven
Trunc

Sin
Cos
Tan
Atan2

Exp
Log
Log2
Log10

Mod
Hypot

IsNaN
IsInf

Pi
E
```

Then learn the floating-point internals and specialized mathematical functions when your project actually needs them.

---

# 11. Thought-Provoking Question

Imagine you're building a **payment system in Go** where you calculate prices, discounts, taxes, and interest.

**Would you choose `float64` + the `math` package for all of those calculations? Why or why not?**

Think especially about this:

> If a computer's floating-point representation cannot exactly represent some decimal values, what could happen when you repeatedly calculate money using `math.Pow()`, `math.Round()`, and `float64`?

This leads naturally into an important Go topic:

**`float64` vs `math/big` vs integer-based money calculations.**
