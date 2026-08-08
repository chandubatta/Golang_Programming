# User Input (Different Types) in Go

## 1. What is User Input?

**User input** means accepting data from a user while a Go program is running.

Go commonly uses the `fmt` package to read input from the keyboard:

- `fmt.Scan()` — reads space-separated input.
- `fmt.Scanln()` — reads input until Enter.
- `fmt.Scanf()` — reads input according to a specified format.

Different Go data types can be read from the user, such as:

- `string`
- `int`
- `float64`
- `bool`
- `byte` / `rune`

### Basic syntax

```go
fmt.Scan(&variable)
```

The `&` is important because `Scan` needs the **address of the variable** so it can store the user's input there.

### When is it commonly used?

User input is useful in:

- Command-line applications
- Learning and practice programs
- Calculator programs
- Menu-driven applications
- Configuration/setup programs
- Small interactive utilities

---

## 2. Simple Example — Different Input Types

```go
package main

import "fmt"

func main() {
    var name string
    var age int
    var salary float64
    var isDeveloper bool

    fmt.Print("Enter your name: ")
    fmt.Scan(&name)

    fmt.Print("Enter your age: ")
    fmt.Scan(&age)

    fmt.Print("Enter your salary: ")
    fmt.Scan(&salary)

    fmt.Print("Are you a developer? (true/false): ")
    fmt.Scan(&isDeveloper)

    fmt.Println("\n--- User Information ---")
    fmt.Println("Name:", name)
    fmt.Println("Age:", age)
    fmt.Println("Salary:", salary)
    fmt.Println("Developer:", isDeveloper)
}
```

### Example input

```text
Enter your name: Chandu
Enter your age: 25
Enter your salary: 45000.50
Are you a developer? (true/false): true
```

### Output

```text
--- User Information ---
Name: Chandu
Age: 25
Salary: 45000.5
Developer: true
```

### Important point

Notice the `&`:

```go
fmt.Scan(&age)
```

It means:

> "Give `Scan` the memory address of `age` so it can put the user's input into it."

---

## 3. Three Common Beginner Mistakes

### Mistake 1: Forgetting `&`

Incorrect:

```go
var age int
fmt.Scan(age)
```

Correct:

```go
var age int
fmt.Scan(&age)
```

**Why?**

`Scan` needs the address of the variable to modify it.

---

### Mistake 2: Expecting `Scan()` to read a complete sentence

Suppose the user enters:

```text
Chandu Batta
```

This:

```go
var name string
fmt.Scan(&name)
```

usually reads only:

```text
Chandu
```

because `Scan` separates input using whitespace.

For complete lines/sentences, use `bufio.Reader` or `fmt.Scanln` depending on the requirement.

For example:

```go
reader := bufio.NewReader(os.Stdin)

name, _ := reader.ReadString('\n')
```

This is more appropriate when the input can contain spaces.

---

### Mistake 3: Assuming input is always valid

For example:

```go
var age int
fmt.Scan(&age)
```

If the user enters:

```text
hello
```

Go cannot convert `"hello"` into an `int`.

A better approach is to check the return value:

```go
_, err := fmt.Scan(&age)

if err != nil {
    fmt.Println("Invalid input")
}
```

For production applications, input validation should be more thorough.

---

## 4. Real-World Applications

### Application 1: Command-Line Registration

A CLI application could ask:

```text
Enter username:
Enter age:
Enter email:
Enter phone number:
```

The program can then validate the values and create a user account.

For example:

```go
var username string
var age int

fmt.Print("Username: ")
fmt.Scan(&username)

fmt.Print("Age: ")
fmt.Scan(&age)
```

This approach is useful for command-line tools and setup utilities.

---

### Application 2: Menu-Driven Applications

A Go CLI application might provide:

```text
1. Add User
2. View Users
3. Delete User
4. Exit

Enter your choice:
```

The program can read the user's choice:

```go
var choice int

fmt.Print("Enter your choice: ")
fmt.Scan(&choice)

switch choice {
case 1:
    fmt.Println("Add User")
case 2:
    fmt.Println("View User")
case 3:
    fmt.Println("Delete User")
case 4:
    fmt.Println("Exit")
default:
    fmt.Println("Invalid choice")
}
```

This pattern is commonly used when learning Go and when building simple CLI tools.

---

## 5. Progressive Exercises

### Exercise 1 — Basic

Create a Go program that asks the user for:

- Name
- Age
- City

Then display the information in a formatted output.

Example:

```text
Enter your name: Chandu
Enter your age: 25
Enter your city: Tirupati

--- Profile ---
Name: Chandu
Age: 25
City: Tirupati
```

**Requirement:** Use appropriate Go data types for each input.

---

### Exercise 2 — Intermediate

Create a **Student Marks Calculator**.

Ask the user to enter:

- Student name
- Roll number
- Marks in English
- Marks in Mathematics
- Marks in Science

Then calculate and display:

- Total marks
- Average marks
- Whether the student passed or failed

Define your own passing criteria.

**Requirement:** Handle numeric inputs using appropriate data types.

---

### Exercise 3 — Advanced

Create a **CLI Employee Registration and Salary Calculator**.

Ask the user for:

- Employee name
- Employee ID
- Age
- Basic salary
- HRA percentage
- Bonus amount
- Whether the employee is a permanent employee (`true/false`)

Then calculate the employee's final salary according to rules you define.

Your program should:

1. Read different data types.
2. Validate numeric input.
3. Handle invalid input gracefully.
4. Display a formatted employee report.
5. Use a menu that allows the user to register another employee or exit.

**Challenge:** Try to handle names containing spaces rather than accepting only a single word.

---

## 💭 Thought-Provoking Question

If `fmt.Scan()` can read different data types automatically, **how do you think Go knows whether the input `"25"` should become an `int`, `float64`, or `string`?**

And more importantly:

**What problems could occur in a real application if you blindly trust the type conversion performed on user input?**

Try answering this in your own words before looking for the answer.
