# Object-Oriented or Not in Golang

## 1. What does “Object-Oriented or Not” mean in Go?

Go is **not a traditional Object-Oriented Programming (OOP) language** like Java, C++, or C#.

However, Go supports many **OOP-like concepts** through:

- **Structs** → represent data/objects
- **Methods** → behavior associated with a type
- **Interfaces** → abstraction and polymorphism
- **Embedding** → composition and code reuse
- **Encapsulation** → controlled access using exported/unexported names

Go intentionally avoids some traditional OOP features such as:

- Classes
- Class inheritance
- Constructors as a language feature
- Method overloading
- Operator overloading
- Explicit `public`, `private`, and `protected` keywords

### Main idea

Instead of thinking:

> “Everything must be an object.”

Go encourages:

> **“Combine data and behavior when useful, and prefer composition over inheritance.”**

This makes Go suitable for backend services, APIs, distributed systems, networking applications, CLI tools, and many other systems.

---

## 2. Simple Code Example

Here is a simple example showing how Go achieves an **object-oriented style without classes**:

```go
package main

import "fmt"

// Struct represents data
type User struct {
    Name string
    Age  int
}

// Method represents behavior
func (u User) Introduce() {
    fmt.Println("Hello, my name is", u.Name)
    fmt.Println("I am", u.Age, "years old")
}

func main() {
    // Create a value of User
    user := User{
        Name: "Chandu",
        Age: 25,
    }

    // Call the method
    user.Introduce()
}
```

### Output

```text
Hello, my name is Chandu
I am 25 years old
```

Here:

```go
type User struct {
    Name string
    Age  int
}
```

acts somewhat like a class's **data definition**.

And:

```go
func (u User) Introduce()
```

is a **method associated with `User`**.

### Traditional OOP vs Go

| Traditional OOP | Go |
|---|---|
| Class | Struct |
| Object | Struct value |
| Method | Method |
| Interface | Interface |
| Inheritance | Composition/Embedding |
| Encapsulation | Exported/unexported names |
| Polymorphism | Interfaces |

---

## 3. Three Common Mistakes and Misconceptions

### Mistake 1: “Go is completely OOP”

This is incorrect.

Go supports several OOP concepts, but it is **not a traditional class-based OOP language**.

For example, Go doesn't have:

```go
class User {
    ...
}
```

Instead, it uses:

```go
type User struct {
    Name string
}
```

### How to avoid it

Think of Go as:

> **A multi-paradigm language that supports object-oriented techniques without requiring traditional OOP.**

---

### Mistake 2: Trying to implement inheritance everywhere

A beginner coming from Java or C++ might try to create inheritance hierarchies.

For example:

```text
Animal
  |
  +-- Dog
  |
  +-- Cat
```

Go generally encourages **composition** instead.

For example:

```go
type Animal struct {
    Name string
}

type Dog struct {
    Animal
    Breed string
}
```

Here `Dog` embeds `Animal`.

### How to avoid it

Prefer:

> **Composition over inheritance**

Ask:

> “Can I build this type by combining smaller types?”

instead of:

> “Which class should this inherit from?”

---

### Mistake 3: Thinking methods can only use pointers

Go supports methods with both **value receivers** and **pointer receivers**.

Value receiver:

```go
func (u User) GetName() string {
    return u.Name
}
```

Pointer receiver:

```go
func (u *User) SetName(name string) {
    u.Name = name
}
```

The pointer receiver is particularly useful when the method needs to modify the original value or when copying the value would be undesirable.

### How to avoid it

Understand the difference:

```text
User  → method receives a copy
*User → method can work with the original value
```

---

## 4. Real-World Applications

### Application 1: REST API Backend

Suppose you're building a Go HRMS application.

You might define:

```go
type Employee struct {
    ID     int
    Name   string
    Salary float64
}
```

Then associate behavior with it:

```go
func (e Employee) IsEligibleForBonus() bool {
    return e.Salary > 50000
}
```

Your application can organize its business logic around types such as:

```text
Employee
Department
Attendance
Salary
Leave
```

Interfaces can then abstract services:

```go
type EmployeeService interface {
    GetEmployee(id int) Employee
    CreateEmployee(e Employee)
}
```

This approach is very useful when building large Go backend applications.

---

### Application 2: Microservices

In a microservices system, you can represent different domain concepts using structs and interfaces.

For example:

```text
Payment Service
       |
       +-- Payment
       +-- PaymentRepository
       +-- PaymentService

Position Service
       |
       +-- Position
       +-- PositionRepository
       +-- PositionService
```

Interfaces can separate business logic from implementation.

For example:

```go
type PaymentRepository interface {
    SavePayment()
    GetPayment()
}
```

You could have different implementations:

```text
PaymentRepository
       |
       +-- MySQLPaymentRepository
       +-- MockPaymentRepository
```

This makes testing and changing implementations easier.

---

## 5. Progressive Exercises

### Exercise 1 — Beginner: Bank Account

Create a `BankAccount` struct containing:

- `AccountNumber`
- `HolderName`
- `Balance`

Create methods to:

- Deposit money
- Withdraw money
- Display account information
- Prevent withdrawal when the balance is insufficient

**Goal:** Practice structs and methods.

---

### Exercise 2 — Intermediate: Shape Interface

Create different shapes such as:

- `Circle`
- `Rectangle`
- `Triangle`

Create an interface that represents a common behavior for all shapes.

Each shape should provide its own implementation of that behavior.

Your program should be able to store different shapes together and perform the common operation without knowing their concrete type.

**Goal:** Practice interfaces and polymorphism.

---

### Exercise 3 — Advanced: Payment System

Design a small payment system supporting:

- Credit Card
- UPI
- PayPal

Create a common interface for payment processing.

Each payment type should implement the interface differently.

Your system should:

- Accept different payment types
- Process payments through the common interface
- Return success/failure information
- Keep payment-processing logic independent from the concrete payment implementation
- Make it possible to add a new payment method later without changing the existing payment-processing logic

**Goal:** Practice **interfaces, polymorphism, abstraction, composition, and clean design**.

---

## Key Takeaway

Go is best described as:

> **Not a traditional OOP language, but a language that supports many object-oriented programming concepts.**

The important concepts to remember are:

```text
Structs       → Data
Methods       → Behavior
Interfaces    → Abstraction + Polymorphism
Embedding     → Composition + Reuse
Encapsulation → Exported / Unexported identifiers
```

The biggest mindset shift for someone coming from Java/C++ is:

> **Don't try to force Go into a traditional OOP model. Use structs, methods, interfaces, and composition where they make the design simpler.**

---

## Thought-Provoking Question

If Go can achieve **encapsulation, abstraction, polymorphism, and code reuse without classes and traditional inheritance**, **what advantages might this simpler approach have when designing a large microservices-based backend—and when might traditional class-based OOP actually be a better choice?**
