# Value Type & Reference Type in Go

## 1. Concise Explanation

In Go, the terms **value type** and **reference type** are useful for understanding what happens when you assign variables, pass them to functions, or modify data.

> **Important:** Go is technically a **pass-by-value language**. Even pointers, slices, maps, and channels are passed by value—the value being copied may contain a reference to shared underlying data.

### Value Types

A **value type** stores the actual data directly. When you assign it to another variable or pass it to a function, the value is **copied**.

Common value types include:

- `int`
- `float64`
- `bool`
- `string`
- `array`
- `struct`

Example:

```go
a := 10
b := a

b = 20

fmt.Println(a) // 10
fmt.Println(b) // 20
```

`b := a` creates a copy of `a`.

Changing `b` does **not** change `a`.

### Reference-Like Types

Go doesn't have a formal language category called "reference types" in the same sense as some other languages. However, several types contain references to underlying/shared data.

Common examples are:

- `slice`
- `map`
- `channel`
- `pointer`
- `function`
- `interface`

For example, a slice is a small descriptor containing information about an underlying array. Copying the slice copies that descriptor, while the underlying array can remain shared.

```go
numbers1 := []int{10, 20, 30}
numbers2 := numbers1

numbers2[0] = 100

fmt.Println(numbers1) // [100 20 30]
fmt.Println(numbers2) // [100 20 30]
```

Both slices refer to the same underlying array in this example.

### Quick Comparison

| Type | Assignment behavior | Can underlying data be shared? |
|---|---|---|
| `int` | Value copied | No |
| `bool` | Value copied | No |
| `string` | Value copied | String data is immutable |
| `array` | Entire array copied | No |
| `struct` | Entire struct copied | Depends on its fields |
| `slice` | Slice descriptor copied | **Yes** |
| `map` | Map value copied | **Yes** |
| `channel` | Channel value copied | **Yes** |
| `pointer` | Address copied | **Yes** |

---

## 2. Simple Code Example

```go
package main

import "fmt"

func main() {
    // Value type
    a := 10
    b := a

    b = 20

    fmt.Println("Value type:")
    fmt.Println("a =", a)
    fmt.Println("b =", b)

    // Reference-like behavior with slice
    numbers1 := []int{10, 20, 30}
    numbers2 := numbers1

    numbers2[0] = 100

    fmt.Println("\nSlice:")
    fmt.Println("numbers1 =", numbers1)
    fmt.Println("numbers2 =", numbers2)
}
```

### Output

```text
Value type:
a = 10
b = 20

Slice:
numbers1 = [100 20 30]
numbers2 = [100 20 30]
```

### Why?

For `int`:

```text
a = 10
  ↓ copy
b = 10
```

After:

```go
b = 20
```

you have:

```text
a = 10
b = 20
```

For the slice:

```text
numbers1 ──┐
           ├──> underlying array [10 20 30]
numbers2 ──┘
```

Therefore:

```go
numbers2[0] = 100
```

can affect what `numbers1` observes.

---

# 3. Three Common Mistakes

## Mistake 1: Thinking Go has traditional "reference types"

A common misconception is:

> "Slices and maps are reference types, and Go passes them by reference."

That's not technically correct.

Go **always passes arguments by value**.

For example:

```go
func change(numbers []int) {
    numbers[0] = 100
}
```

The slice descriptor itself is copied, but the copied descriptor can still point to the **same underlying array**.

**Avoid it:** Think in terms of **what is copied** and **what underlying data is shared**.

---

## Mistake 2: Assuming every slice assignment creates a copy

Beginners may write:

```go
a := []int{1, 2, 3}
b := a

b[0] = 100
```

and expect:

```text
a = [1 2 3]
b = [100 2 3]
```

But they may share the same underlying array.

**Avoid it:** If you need an independent slice, explicitly copy the elements.

```go
b := make([]int, len(a))
copy(b, a)
```

---

## Mistake 3: Assuming structs are always completely independent

Structs are value types:

```go
type User struct {
    Name string
    Age  int
}
```

So:

```go
u2 := u1
```

copies the struct.

But a struct can contain a slice or map:

```go
type User struct {
    Name string
    Tags []string
}
```

Now copying the struct copies the slice descriptor, while the underlying slice data may still be shared.

**Avoid it:** Check the fields inside a struct. A value-type struct can contain reference-like fields.

---

# 4. Real-World Applications

## Application 1: API Request/Response Structures

Go backend applications commonly use structs:

```go
type User struct {
    ID    int
    Name  string
    Email string
}
```

You can safely create copies when you want to work with independent user values.

This is useful for:

- REST APIs
- JSON requests/responses
- database models
- business logic
- validation

For example:

```go
user2 := user1
user2.Name = "Ravi"
```

Changing the copied struct's `Name` doesn't change `user1.Name`.

---

## Application 2: Shared Collections

Slices and maps are extremely common in backend applications:

```go
users := []string{
    "Chandu",
    "Ravi",
    "John",
}
```

You may pass the slice to different functions:

```go
processUsers(users)
```

The slice descriptor is passed by value, but the underlying collection can be shared.

This is useful when working with:

- lists of users
- database query results
- JSON data
- configuration
- caching
- collections processed by multiple functions

When you **don't want sharing**, create an explicit copy.

---

# 5. Three Progressively Challenging Exercises

## Exercise 1 — Basic Value vs Slice

Create a Go program containing:

1. An `int` variable.
2. Copy it into another variable.
3. Modify the second variable.
4. Print both variables.
5. Create a slice of integers.
6. Assign it to another slice.
7. Modify one element through the second slice.
8. Print both slices.

**Goal:** Observe the difference between copying an ordinary value and copying a slice.

---

## Exercise 2 — Function Parameters

Create:

```go
type Employee struct {
    Name   string
    Salary int
}
```

Implement two functions:

```text
updateEmployee(...)
updateSalaries(...)
```

Requirements:

1. Pass an `Employee` to the first function.
2. Try modifying its `Salary`.
3. Observe whether the original employee changes.
4. Create a slice of employees.
5. Pass the slice to the second function.
6. Modify the salary of one employee inside the function.
7. Observe what happens to the original slice.

**Goal:** Understand how Go passes structs and slices to functions.

---

## Exercise 3 — Struct Containing a Slice

Create:

```go
type Student struct {
    Name  string
    Marks []int
}
```

Then:

1. Create one `Student`.
2. Assign it to another `Student` variable.
3. Modify the second student's `Name`.
4. Modify one element in the second student's `Marks`.
5. Print both students.
6. Determine which modifications affect the original.
7. Create a second version where `Marks` is independently copied.
8. Compare the results.

**Goal:** Understand why a value-type struct can still contain shared underlying data.

---

# Thought-Provoking Question

Suppose you have a large struct containing several fields, including slices and maps.

**Would you prefer to pass that struct by value or pass a pointer to it—and what trade-offs involving performance, data sharing, and accidental modification would influence your decision?**
