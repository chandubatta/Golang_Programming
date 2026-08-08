# Static Typing and Dynamic Typing in Go

## 1. What are Static Typing and Dynamic Typing?

### Static Typing

A **statically typed language** determines the type of a variable at **compile time**.

Go is a **statically typed language**.

```go
var age int = 25
var name string = "Chandu"
var salary float64 = 50000.50
```

Once `age` is an `int`, you cannot assign a `string` to it:

```go
age = "25" // Compile-time error
```

Go can also infer the type:

```go
age := 25
name := "Chandu"
```

The compiler knows these types before the program runs.

### Dynamic Typing

In a **dynamically typed language**, the type of a value is generally determined and checked **at runtime**.

For example, languages such as Python and JavaScript allow a variable to hold values of different types during execution:

```python
x = 10
x = "Hello"
```

Go itself is **not dynamically typed**.

However, Go provides mechanisms that can hold values of different concrete types, particularly `any` (an alias for `interface{}`):

```go
var value any

value = 10
value = "Hello"
value = true
```

The variable `value` can contain different types at different times, but the actual type is inspected at runtime when you need to work with the contained value.

### Quick Comparison

| Feature | Static Typing | Dynamic Typing |
|---|---|---|
| Type checking | Compile time | Runtime |
| Go | Yes | No |
| Python | No | Yes |
| Java | Yes | No |
| Type errors | Often caught earlier | Often discovered during execution |
| Flexibility | More controlled | More flexible |
| Performance predictability | Generally better | Can have more runtime overhead |

### Purpose

**Static typing** helps you:

- Catch many errors before running the program.
- Make code easier to understand and maintain.
- Get better IDE/compiler support.
- Make large applications more predictable.

**Dynamic typing** is useful when:

- Data structures need to handle different types.
- Rapid experimentation is important.
- The exact type of data isn't known until runtime.

---

## 2. Simple Go Example

Because Go is statically typed, this is the normal approach:

```go
package main

import "fmt"

func main() {
    var age int = 25
    var name string = "Chandu"

    fmt.Println(age)
    fmt.Println(name)

    // age = "25" // Compile-time error
}
```

The compiler knows:

```text
age  -> int
name -> string
```

### Go's `any` Example

Go can also store different types in an `any` variable:

```go
package main

import "fmt"

func main() {
    var value any

    value = 100
    fmt.Println(value)

    value = "Hello"
    fmt.Println(value)

    value = true
    fmt.Println(value)
}
```

Output:

```text
100
Hello
true
```

But remember:

> **Using `any` does NOT make Go a dynamically typed language.**

Go is still fundamentally **statically typed**.

For example:

```go
var value any = 100

// fmt.Println(value + 10) // Not directly allowed
```

You may need a type assertion:

```go
number, ok := value.(int)

if ok {
    fmt.Println(number + 10)
}
```

Here, the concrete type contained inside `any` is checked at runtime.

---

## 3. Three Common Beginner Mistakes

### Mistake 1: Thinking `:=` Makes Go Dynamically Typed

Beginners sometimes see:

```go
x := 10
x = "Hello"
```

and think Go is dynamically typed.

It isn't.

```go
x := 10
x = "Hello" // Compile-time error
```

`:=` simply allows Go to **infer the type**.

Remember:

```go
x := 10
```

means approximately:

```go
var x int = 10
```

---

### Mistake 2: Thinking `any` Makes Go Dynamically Typed

Another common misconception is:

```go
var data any

data = 10
data = "Hello"
data = true
```

This doesn't change Go's type system.

`any` is an interface type that can hold a value of different concrete types.

Remember:

> **Go = statically typed, even when using `any`.**

---

### Mistake 3: Confusing Type Inference with Dynamic Typing

Consider:

```go
name := "Chandu"
```

Go automatically determines:

```text
name -> string
```

The compiler still knows the type.

Therefore:

```go
name = 100 // Compile-time error
```

Remember:

```text
Explicit typing:
var age int = 25

Type inference:
age := 25

Both are statically typed.
```

---

## 4. Real-World Applications

### Application 1: Large Backend Systems

Go's static typing is particularly useful for backend systems such as:

- REST APIs
- Microservices
- Payment systems
- Banking applications
- E-commerce applications
- Authentication systems

For example:

```go
type Payment struct {
    ID     int
    Amount float64
    Status string
}
```

The compiler can catch many incorrect uses of these fields before deployment.

This is especially valuable when multiple developers are working on a large codebase.

### Application 2: Processing Dynamic External Data

Sometimes an application receives data whose exact structure or type isn't known beforehand.

For example, an API might return different JSON structures.

You might temporarily decode such data into:

```go
var data any
```

and inspect the actual value at runtime.

This can be useful for:

- Generic JSON processing
- Configuration data
- External APIs
- Logging systems
- Message-processing systems
- Flexible metadata

However, for application/business logic, **strongly typed structs are generally preferable when the data structure is known**.

---

## 5. Progressive Exercises

### Exercise 1 — Beginner

Create a Go program that declares variables for:

- Name
- Age
- Salary
- IsEmployee

Use both:

```go
var
```

and:

```go
:=
```

Then try assigning a value of the wrong type to each variable.

**Your goal:** Observe and understand the compile-time errors.

---

### Exercise 2 — Intermediate

Create a program using:

```go
var data any
```

Store the following values in it one at a time:

- integer
- string
- float
- boolean

Then use a **type switch** to identify which type is currently stored in `data`.

Your program should print a meaningful message for each type.

---

### Exercise 3 — Advanced

Create a function that accepts:

```go
any
```

as an argument.

The function should receive different types of values such as:

- `int`
- `string`
- `float64`
- `bool`

and perform an appropriate operation depending on the actual type.

For example:

- `int` → perform a mathematical calculation
- `string` → perform a string operation
- `float64` → perform a numerical calculation
- `bool` → display a meaningful message

Use a **type switch** rather than writing separate functions for every type.

**Challenge:** Also decide what your function should do when it receives a type that you don't explicitly support.

---

## Thought-Provoking Question

Suppose you are designing a **Golang microservices payment system**.

You can use strongly typed structures such as:

```go
type Payment struct {
    ID     int
    Amount float64
    Status string
}
```

But you could also use:

```go
var payment any
```

**Question:**

> If `any` gives you more flexibility, why might using too much `any` make a large Go application harder to maintain, debug, and secure?
