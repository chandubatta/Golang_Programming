# Lexer in Golang

## 1. What is a Lexer?

A **Lexer (Lexical Analyzer)** is a component that reads raw input text and breaks it into meaningful units called **tokens**.

For example, given:

```text
age = 25 + 10
```

A lexer can produce:

```text
IDENTIFIER(age)
ASSIGN(=)
NUMBER(25)
PLUS(+)
NUMBER(10)
```

### Purpose of a Lexer

The main purpose of a lexer is to:

- Read source code or text character by character.
- Group characters into meaningful tokens.
- Ignore unnecessary characters such as spaces and comments.
- Identify keywords, identifiers, numbers, operators, strings, etc.
- Provide tokens to the next stage, usually a **parser**.

A typical compiler/interpreter pipeline looks like:

```text
Source Code
    ↓
   Lexer
    ↓
  Tokens
    ↓
  Parser
    ↓
Syntax Tree / AST
    ↓
Compiler / Interpreter
```

### When is a Lexer commonly used?

Lexers are commonly used when building:

- Compilers
- Interpreters
- Programming languages
- SQL parsers
- Configuration-file parsers
- Query languages
- Command-line expression parsers
- Template engines

> **Important:** Go doesn't have a function simply called `Lexer`. The standard library provides packages such as `text/scanner` that can perform lexical scanning.

---

# 2. Simple Lexer Example in Go

Go's `text/scanner` package can scan an input string and identify tokens.

```go
package main

import (
	"fmt"
	"strings"
	"text/scanner"
)

func main() {
	input := "age = 25 + 10"

	var s scanner.Scanner
	s.Init(strings.NewReader(input))

	for {
		token := s.Scan()

		if token == scanner.EOF {
			break
		}

		fmt.Printf("Token: %-10s Text: %q\n",
			scanner.TokenString(token),
			s.TokenText(),
		)
	}
}
```

### Example output

```text
Token: Identifier Text: "age"
Token: =          Text: "="
Token: Int        Text: "25"
Token: +          Text: "+"
Token: Int        Text: "10"
```

### What is happening?

The input:

```text
age = 25 + 10
```

is processed approximately like this:

```text
"age"   → Identifier
"="     → Operator
"25"    → Integer
"+"     → Operator
"10"    → Integer
```

The lexer doesn't try to understand the complete meaning of:

```text
age = 25 + 10
```

It only identifies the individual pieces.

The **parser** would later determine how those pieces are structured.

---

# Lexer vs Parser

This distinction is very important.

Consider:

```text
10 + 20 * 30
```

### Lexer

The lexer identifies:

```text
INT(10)
PLUS(+)
INT(20)
MULTIPLY(*)
INT(30)
```

### Parser

The parser determines the structure:

```text
10 + (20 * 30)
```

because multiplication has higher precedence than addition.

So:

> **Lexer = What are these pieces?**

> **Parser = How are these pieces organized?**

---

# 3. Three Common Mistakes and Misconceptions

## Mistake 1: Thinking Lexer understands program meaning

A beginner might think:

```text
age = 25
```

means the lexer understands that `age` is a variable containing `25`.

That's not the lexer's job.

The lexer primarily identifies:

```text
IDENTIFIER → age
ASSIGN     → =
NUMBER     → 25
```

### How to avoid it

Remember:

```text
Lexer → Tokens
Parser → Structure
Semantic Analysis → Meaning
```

---

## Mistake 2: Confusing characters with tokens

Consider:

```text
count = 100
```

A lexer doesn't necessarily treat every character as a separate meaningful token.

Instead:

```text
c o u n t
```

can become one token:

```text
IDENTIFIER("count")
```

Similarly:

```text
1 0 0
```

can become:

```text
NUMBER("100")
```

### How to avoid it

Think of a token as a **meaningful group of characters**.

```text
Characters
    ↓
Grouping
    ↓
Tokens
```

---

## Mistake 3: Thinking whitespace is always a token

Given:

```text
x = 10
```

the spaces usually aren't important to the parser.

A lexer commonly skips:

```text
space
tab
newline
```

unless the language specifically gives them meaning.

### How to avoid it

Understand whether the language treats whitespace as significant.

For example:

```text
x=10
```

and:

```text
x = 10
```

produce essentially the same tokens in many languages.

But some languages or special syntaxes may give indentation or whitespace special meaning.

---

# 4. Real-World Applications of Lexers

## Application 1: Building a Programming Language

Suppose you want to create your own programming language:

```text
x = 10
y = x + 20
print(y)
```

Your lexer could convert it into:

```text
IDENTIFIER(x)
ASSIGN
NUMBER(10)

IDENTIFIER(y)
ASSIGN
IDENTIFIER(x)
PLUS
NUMBER(20)

KEYWORD(print)
LPAREN
IDENTIFIER(y)
RPAREN
```

The parser can then use these tokens to build an **Abstract Syntax Tree (AST)**.

A simplified architecture would be:

```text
Your Programming Language
          ↓
        Lexer
          ↓
        Tokens
          ↓
        Parser
          ↓
         AST
          ↓
     Interpreter
```

---

## Application 2: SQL / Query Processing

Consider:

```sql
SELECT name FROM users WHERE age > 18;
```

A lexer can identify:

```text
KEYWORD(SELECT)
IDENTIFIER(name)
KEYWORD(FROM)
IDENTIFIER(users)
KEYWORD(WHERE)
IDENTIFIER(age)
GREATER_THAN(>)
NUMBER(18)
SEMICOLON(;)
```

The parser can then determine whether the SQL query follows the correct grammar.

This approach is used in database systems, SQL tools, IDEs, query analyzers, and many other systems that need to understand structured text.

---

# 5. Progressive Lexer Exercises

## Exercise 1 — Basic Tokenizer

Write a Go program that accepts:

```text
age = 25
```

and identifies the following token types:

```text
Identifier
Assignment Operator
Integer
```

Your program should print each token and its corresponding text.

**Goal:** Practice identifying basic identifiers, operators, and numbers.

---

## Exercise 2 — Build Your Own Expression Lexer

Create a lexer for expressions such as:

```text
total = price + tax * 2
```

Your lexer should recognize:

- Identifiers
- Integers
- `=`
- `+`
- `-`
- `*`
- `/`
- Parentheses
- Whitespace

For example, it should conceptually produce:

```text
IDENTIFIER(total)
ASSIGN(=)
IDENTIFIER(price)
PLUS(+)
IDENTIFIER(tax)
MULTIPLY(*)
NUMBER(2)
```

**Goal:** Move beyond `text/scanner` usage and start thinking about how a custom lexer would be designed.

---

## Exercise 3 — Create a Lexer for a Mini Language

Design a lexer for a small programming language with statements such as:

```text
let age = 25;
let name = "Chandu";
if age > 18 {
    print(name);
}
```

Your lexer should recognize at least:

- Keywords: `let`, `if`
- Identifiers
- Integers
- Strings
- Assignment operator
- Arithmetic operators
- Comparison operators
- `{` and `}`
- `(` and `)`
- Semicolon
- Whitespace
- Unknown/invalid characters

For invalid input such as:

```text
let age = @25;
```

your lexer should be able to identify `@` as an invalid or unknown token.

**Goal:** Think about how the lexical layer of a real programming language is designed.

---

# A Useful Mental Model

When learning Lexer, remember this:

```text
Raw Text
   ↓
Characters
   ↓
Lexer
   ↓
Tokens
   ↓
Parser
   ↓
AST
   ↓
Program Meaning
```

For example:

```text
x = 10 + 20
```

becomes:

```text
        LEXER
          ↓
IDENTIFIER(x)
ASSIGN(=)
NUMBER(10)
PLUS(+)
NUMBER(20)
          ↓
        PARSER
          ↓
      Expression
          ↓
     x = (10 + 20)
```

# Thought-Provoking Question

**If a lexer can identify tokens but does not understand their meaning, where should you place the responsibility for detecting an expression such as `10 + * 20` as invalid—and why?**
