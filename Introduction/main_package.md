# main package

- In Golang, the main package is a special package that acts as the entry point of a Go program.

**When you run a Go application using:**

go run main.go

**Go looks for:**

A package named *main*
A function named *main()*

Without these, the program cannot execute as a standalone application.

**Purpose of main package**

The main package is used to:

Start program execution
Organize executable applications
Connect different packages/modules together
Control application flow

**Commonly used when**

The main package is commonly used in:

CLI (Command Line Interface) tools
Web servers
APIs
Automation scripts
Microservices
Desktop/backend applications

**2. Simple Example of main Package**
```package main

  import "fmt"

  func main() {
	  fmt.Println("Hello, Golang!")
	  ``` }

**Explanation

package main
→ Defines this file as the executable package.
func main()
→ The first function Go executes.
fmt.Println()
→ Prints output to the console.

**3. Common Beginner Mistakes About main Package**

***Mistake 1: Forgetting package main Wrong***

package demo

If the package is not main, Go will not create an executable program.

How to Avoid

Always use:

package main

for runnable applications.

***Mistake 2: Writing Main() instead of main()***
Wrong
func Main() {
}

Go is case-sensitive.

The runtime only recognizes:

func main()
How to Avoid

Remember:

main() → lowercase only
No parameters
No return type

Correct format:

func main() {
}

***Mistake 3: Thinking Multiple main() Functions Are Allowed ***

Beginners sometimes create multiple main() functions in the same package.

**Wrong**
func main() {
}

func main() {
}

This causes a compilation error.

**How to Avoid**

Only one main() function is allowed per executable package.

You can create many helper functions:

func display() {
}

func calculate() {
}

func main() {
	display()
	calculate()
}

***4. Real-World Applications of main Package***
**Scenario 1: Building REST APIs**

Frameworks like:

Gin
Fiber
Echo

typically start from a main package.

Example responsibilities:

Start server
Configure routes
Connect database
Load environment variables

Example:

func main() {
	connectDB()
	startServer()
}

**Scenario 2: CLI Tools**

Popular Go CLI applications use the main package to:

Parse commands
Execute operations
Handle user input

Examples include:

Docker CLI
Kubernetes tools
Git-like utilities

5. Practice Exercises
***Exercise 1 — Beginner***

Create a Go program using the main package that:

Prints your name
Prints your favorite programming language
Prints today's goal
***Exercise 2 — Intermediate***

Create a Go application with:

One main() function
Three helper functions:
add()
subtract()
multiply()

Call all helper functions from main().

***Exercise 3 — Advanced***

Build a small console-based student management system using:

main package
Multiple helper functions
User input
Loops
Conditional statements

Features:

Add student
Display students
Search student
Exit application
Thought-Provoking Question

If Go allowed multiple main() functions in a single executable package, how do you think that would affect:

program structure,
maintainability,
debugging,
and application startup flow?