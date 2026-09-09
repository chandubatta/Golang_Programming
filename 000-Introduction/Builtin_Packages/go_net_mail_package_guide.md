# Go `net/mail` Package

The Go standard library's `net/mail` package is specifically for **parsing email messages and email addresses**. It follows most of the syntax defined by **RFC 5322**, with extensions from RFC 6532. It is useful when your application needs to inspect, validate, or extract information from email headers and addresses.

> **Important:** `net/mail` is primarily a **parser**, not an email-sending package. For actually sending mail through SMTP, you would look at `net/smtp` or a third-party email library.

---

## 1. What is `net/mail`?

Import it with:

```go
import "net/mail"
```

The package provides functionality for:

- Parsing a single email address
- Parsing multiple email addresses
- Parsing email dates
- Reading an entire email message
- Accessing email headers
- Extracting `From`, `To`, `Cc`, etc.
- Handling RFC 2047 encoded display names
- Formatting parsed addresses back into valid RFC 5322 form

### Think of it this way

Suppose you receive:

```text
From: Alice <alice@example.com>
To: Bob <bob@example.com>
Date: Wed, 09 Sep 2026 10:30:00 +0000
Subject: Welcome
```

`net/mail` helps you turn that raw email text into structured Go values:

```text
Raw Email
    ↓
net/mail
    ↓
Structured information
    ├── Sender
    ├── Recipients
    ├── Date
    ├── Subject
    └── Body
```

---

# 2. Simple Example

Let's start with the most commonly used functionality: parsing an email address.

```go
package main

import (
	"fmt"
	"log"
	"net/mail"
)

func main() {
	address, err := mail.ParseAddress("Alice Johnson <alice@example.com>")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Name:", address.Name)
	fmt.Println("Email:", address.Address)
	fmt.Println("Formatted:", address.String())
}
```

Output:

```text
Name: Alice Johnson
Email: alice@example.com
Formatted: Alice Johnson <alice@example.com>
```

`ParseAddress` converts the textual RFC 5322 address into an `*mail.Address`, whose two important fields are `Name` and `Address`.

---

# 3. Every Important Function and Method in `net/mail`

The public API is relatively small, which makes `net/mail` a nice package to learn thoroughly.

---

## A. `mail.ParseAddress()`

### Signature

```go
func ParseAddress(address string) (*Address, error)
```

### Purpose

Parses **one email address**.

For example:

```text
Alice Johnson <alice@example.com>
```

becomes:

```go
&mail.Address{
	Name:    "Alice Johnson",
	Address: "alice@example.com",
}
```

### Example

```go
package main

import (
	"fmt"
	"net/mail"
)

func main() {
	addr, err := mail.ParseAddress("Alice Johnson <alice@example.com>")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(addr.Name)
	fmt.Println(addr.Address)
}
```

Output:

```text
Alice Johnson
alice@example.com
```

It also accepts:

```go
mail.ParseAddress("alice@example.com")
```

Then:

```text
Name: ""
Address: alice@example.com
```

### When to use it

Use `ParseAddress` when you expect **exactly one address**.

For example:

```text
From: Alice <alice@example.com>
```

---

# B. `mail.ParseAddressList()`

### Signature

```go
func ParseAddressList(list string) ([]*Address, error)
```

This parses **multiple comma-separated email addresses**.

### Example

```go
package main

import (
	"fmt"
	"net/mail"
)

func main() {
	list := "Alice <alice@example.com>, Bob <bob@example.com>"

	addresses, err := mail.ParseAddressList(list)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, addr := range addresses {
		fmt.Println("Name:", addr.Name)
		fmt.Println("Email:", addr.Address)
	}
}
```

Output:

```text
Name: Alice
Email: alice@example.com
Name: Bob
Email: bob@example.com
```

### When useful

Imagine an email header:

```text
To: Alice <alice@example.com>, Bob <bob@example.com>, Charlie <charlie@example.com>
```

You can parse the entire value:

```go
addresses, err := mail.ParseAddressList(headerValue)
```

Instead of manually splitting the string with:

```go
strings.Split(...)
```

That's important because email address syntax is more complicated than simply splitting strings on commas.

---

# C. `mail.ParseDate()`

### Signature

```go
func ParseDate(date string) (time.Time, error)
```

This parses an RFC 5322 email date into a Go `time.Time`.

### Example

```go
package main

import (
	"fmt"
	"net/mail"
)

func main() {
	dateString := "Wed, 09 Sep 2026 10:30:00 +0000"

	t, err := mail.ParseDate(dateString)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(t)
}
```

You can then use normal `time.Time` operations:

```go
fmt.Println(t.Year())
fmt.Println(t.Month())
fmt.Println(t.Day())
```

### Why is this useful?

Email dates aren't necessarily in the exact format your application wants.

For example:

```text
Wed, 09 Sep 2026 10:30:00 +0000
```

can become a Go `time.Time`.

Then you can:

```go
t.Format(time.RFC3339)
```

and obtain:

```text
2026-09-09T10:30:00Z
```

---

# D. `mail.ReadMessage()`

### Signature

```go
func ReadMessage(r io.Reader) (msg *Message, err error)
```

This is one of the most important functions in the package.

It reads an email message from an `io.Reader` and separates the **headers from the body**.

For example:

```text
From: alice@example.com
To: bob@example.com
Subject: Hello
Date: Wed, 09 Sep 2026 10:30:00 +0000

Hello Bob!
How are you?
```

You can use:

```go
msg, err := mail.ReadMessage(reader)
```

and then access:

```go
msg.Header
```

and:

```go
msg.Body
```

### Example

```go
package main

import (
	"fmt"
	"net/mail"
	"strings"
)

func main() {
	rawEmail := `From: Alice <alice@example.com>
To: Bob <bob@example.com>
Subject: Hello
Date: Wed, 09 Sep 2026 10:30:00 +0000

Hello Bob!
How are you?
`

	msg, err := mail.ReadMessage(strings.NewReader(rawEmail))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("From:", msg.Header.Get("From"))
	fmt.Println("To:", msg.Header.Get("To"))
	fmt.Println("Subject:", msg.Header.Get("Subject"))

	fmt.Println("Body:")

	buf := make([]byte, 1024)
	n, _ := msg.Body.Read(buf)

	fmt.Println(string(buf[:n]))
}
```

### Important concept

`ReadMessage` does **not** mean:

> "Parse the entire MIME email including attachments and HTML parts."

It primarily parses the message structure into headers and body. For MIME multipart processing, you typically use packages such as `mime`, `mime/multipart`, etc.

---

# E. `Address.String()`

### Signature

```go
func (a *Address) String() string
```

This converts a `mail.Address` back into a properly formatted RFC 5322 address.

### Example

```go
addr := &mail.Address{
	Name:    "Alice Johnson",
	Address: "alice@example.com",
}

fmt.Println(addr.String())
```

Output:

```text
Alice Johnson <alice@example.com>
```

### Why is this useful?

Suppose you construct an address:

```go
addr := &mail.Address{
	Name:    "Alice",
	Address: "alice@example.com",
}
```

You can safely turn it into:

```text
Alice <alice@example.com>
```

The method also handles non-ASCII display names using RFC 2047 encoding when necessary.

---

# F. `AddressParser`

There is also:

```go
type AddressParser struct {
	WordDecoder *mime.WordDecoder
}
```

An `AddressParser` lets you configure how encoded display names are decoded.

For normal email parsing, you can often simply use:

```go
mail.ParseAddress(...)
```

But if you need custom RFC 2047 decoding behavior, `AddressParser` becomes useful.

---

# G. `AddressParser.Parse()`

### Signature

```go
func (p *AddressParser) Parse(address string) (*Address, error)
```

Conceptually, it is the configurable equivalent of:

```go
mail.ParseAddress(...)
```

### Example

```go
package main

import (
	"fmt"
	"mime"
	"net/mail"
)

func main() {
	parser := mail.AddressParser{
		WordDecoder: &mime.WordDecoder{
			CharsetReader: nil,
		},
	}

	addr, err := parser.Parse("Alice <alice@example.com>")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(addr.Name)
	fmt.Println(addr.Address)
}
```

### When should you use it?

Use:

```go
mail.ParseAddress(...)
```

for ordinary cases.

Use:

```go
mail.AddressParser{...}
```

when you specifically need to control encoded-word decoding.

---

# H. `AddressParser.ParseList()`

### Signature

```go
func (p *AddressParser) ParseList(list string) ([]*Address, error)
```

It parses a comma-separated list of addresses using the parser's configured `WordDecoder`.

Conceptually:

```go
parser.ParseList(...)
```

is the configurable version of:

```go
mail.ParseAddressList(...)
```

Example:

```go
parser := mail.AddressParser{}

addresses, err := parser.ParseList(
	"Alice <alice@example.com>, Bob <bob@example.com>",
)
```

---

# I. `Header.Get()`

`Header` represents the email headers.

Conceptually:

```go
type Header map[string][]string
```

`Get` retrieves the **first value** associated with a header key and treats the key case-insensitively.

### Example

```go
msg, _ := mail.ReadMessage(strings.NewReader(rawEmail))

subject := msg.Header.Get("Subject")

fmt.Println(subject)
```

If the email contains:

```text
Subject: Welcome to our service
```

you get:

```text
Welcome to our service
```

### Important

This:

```go
msg.Header.Get("subject")
```

and:

```go
msg.Header.Get("Subject")
```

are treated case-insensitively.

---

# J. `Header.Date()`

### Signature

```go
func (h Header) Date() (time.Time, error)
```

This gets the `Date` header and parses it using `ParseDate`.

Instead of:

```go
dateString := msg.Header.Get("Date")

t, err := mail.ParseDate(dateString)
```

you can simply do:

```go
t, err := msg.Header.Date()
```

### Example

```go
date, err := msg.Header.Date()
if err != nil {
	fmt.Println("Could not parse date:", err)
	return
}

fmt.Println(date)
```

If the `Date` header doesn't exist, the package returns:

```go
mail.ErrHeaderNotPresent
```

---

# K. `Header.AddressList()`

### Signature

```go
func (h Header) AddressList(key string) ([]*Address, error)
```

This retrieves a header and parses it as a list of email addresses.

For example:

```text
To: Alice <alice@example.com>, Bob <bob@example.com>
```

You can write:

```go
addresses, err := msg.Header.AddressList("To")
```

Then:

```go
for _, addr := range addresses {
	fmt.Println(addr.Name, addr.Address)
}
```

This is especially useful for:

```text
From
To
Cc
Reply-To
```

and other address-bearing headers.

---

# L. `mail.ErrHeaderNotPresent`

The package defines:

```go
var ErrHeaderNotPresent = errors.New("mail: header not in message")
```

It is returned when methods such as:

```go
Header.Date()
```

or:

```go
Header.AddressList()
```

need a header that isn't present.

You can check it with:

```go
if errors.Is(err, mail.ErrHeaderNotPresent) {
	fmt.Println("Header doesn't exist")
}
```

---

# 4. Complete Practical Example

Here's a more realistic example combining several parts of the package:

```go
package main

import (
	"fmt"
	"net/mail"
	"strings"
)

func main() {
	rawEmail := `From: Alice Johnson <alice@example.com>
To: Bob <bob@example.com>, Charlie <charlie@example.com>
Cc: David <david@example.com>
Subject: Go Learning
Date: Wed, 09 Sep 2026 10:30:00 +0000

Hello,

Today we are learning Go's net/mail package.

Regards,
Alice
`

	msg, err := mail.ReadMessage(strings.NewReader(rawEmail))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Subject
	fmt.Println("Subject:", msg.Header.Get("Subject"))

	// From
	from, err := msg.Header.AddressList("From")
	if err != nil {
		fmt.Println("From error:", err)
		return
	}

	fmt.Println("From:", from[0].Name)
	fmt.Println("Email:", from[0].Address)

	// To
	to, err := msg.Header.AddressList("To")
	if err != nil {
		fmt.Println("To error:", err)
		return
	}

	fmt.Println("Recipients:")

	for _, recipient := range to {
		fmt.Printf("- %s <%s>\n",
			recipient.Name,
			recipient.Address,
		)
	}

	// Date
	date, err := msg.Header.Date()
	if err != nil {
		fmt.Println("Date error:", err)
		return
	}

	fmt.Println("Date:", date)
}
```

This demonstrates a useful mental model:

```text
Raw email
   │
   ▼
ReadMessage()
   │
   ├── Header.Get()       → Subject
   │
   ├── Header.AddressList()
   │        │
   │        └── From / To / Cc
   │
   └── Header.Date()      → time.Time
```

---

# 5. Three Common Beginner Mistakes

## Mistake 1: Thinking `net/mail` sends emails

A common misconception is:

```text
net/mail
```

means:

> "This package sends email."

It doesn't.

Its primary purpose is **parsing mail messages and addresses**.

For example:

```go
mail.ParseAddress(...)
```

parses an address; it doesn't contact an SMTP server.

### Avoid it

Separate your concepts:

```text
net/mail
    ↓
Parse email data

SMTP
    ↓
Send email
```

---

## Mistake 2: Using `strings.Split()` to parse email addresses

Beginners might do:

```go
emails := strings.Split(input, ",")
```

This can be problematic because email address syntax is more complicated than simple comma-separated strings.

Prefer:

```go
addresses, err := mail.ParseAddressList(input)
```

The package understands RFC 5322-style address syntax.

---

## Mistake 3: Assuming `ParseAddress()` verifies that an inbox exists

Consider:

```go
addr, err := mail.ParseAddress("alice@example.com")
```

A successful result means the string is syntactically acceptable to the parser.

It does **not** mean:

> Alice's mailbox exists.

It doesn't send a verification email, contact the mail server, or prove deliverability.

### Avoid it

Think of validation in layers:

```text
Syntax validation
       ↓
net/mail

Domain/DNS checks
       ↓
net.LookupMX / DNS

Actual mailbox verification
       ↓
Usually requires application-level verification
```

---

# 6. Two Real-World Applications

## Application 1: Email ingestion system

Imagine a support system receiving incoming emails.

An email arrives:

```text
From: customer@example.com
To: support@company.com
Subject: Payment problem
```

Your Go service can use:

```go
mail.ReadMessage(...)
```

and then:

```go
msg.Header.Get("Subject")
```

and:

```go
msg.Header.AddressList("From")
```

to extract structured information.

You could then store:

```text
Sender: customer@example.com
Subject: Payment problem
Body: ...
```

in your database.

This is useful for:

- Support ticket systems
- Email-to-ticket applications
- Email archiving
- Email analytics
- Internal mail processing

---

## Application 2: Email processing/security pipeline

Suppose your application receives thousands of emails.

You could parse:

```text
From
To
Cc
Reply-To
Date
Subject
```

and use that information for processing or filtering.

For example:

```text
Incoming email
      ↓
ReadMessage()
      ↓
Extract From
      ↓
Extract To
      ↓
Extract Date
      ↓
Inspect Subject
      ↓
Application processing
```

This can form part of a larger email security or automation pipeline.

---

# 7. Three Progressively Challenging Exercises

## Exercise 1 — Parse an Email Address

Write a Go program that accepts an email address such as:

```text
Alice Johnson <alice@example.com>
```

Use `net/mail` to parse it and print:

```text
Name: Alice Johnson
Email: alice@example.com
```

Your program should gracefully handle invalid addresses.

**Do not use `strings.Split()` to perform the parsing.**

---

## Exercise 2 — Build an Email Header Analyzer

Given this raw email:

```text
From: Alice <alice@example.com>
To: Bob <bob@example.com>, Charlie <charlie@example.com>
Cc: David <david@example.com>
Subject: Project Update
Date: Wed, 09 Sep 2026 10:30:00 +0000

The project is progressing well.
```

Use `mail.ReadMessage()` to analyze the email.

Your program should display:

```text
From:
  Alice <alice@example.com>

To:
  Bob <bob@example.com>
  Charlie <charlie@example.com>

Cc:
  David <david@example.com>

Subject:
  Project Update

Date:
  <parsed time>
```

Handle missing headers without crashing.

---

## Exercise 3 — Build a Mini Email Processing Pipeline

Create a program that accepts multiple raw email messages.

For every email, your program should:

1. Parse the complete message.
2. Extract the sender.
3. Extract all recipients.
4. Extract the subject.
5. Parse the date.
6. Read the body.
7. Reject malformed messages.
8. Detect messages whose `To` header contains more than one recipient.
9. Produce a structured summary for each valid message.

Then extend the program so that it can answer:

```text
How many emails came from each sender?
How many recipients did each email have?
Which email was the oldest?
Which email was the newest?
```

Try to design the program so that the email parsing logic is separated from the reporting logic.

**Do not use third-party email parsing libraries. Use the Go standard library.**

---

# 8. Important Limitations to Remember

`net/mail` follows most of RFC 5322 and RFC 6532, but it intentionally has some divergences. In particular, it does not parse obsolete address formats, doesn't perform Unicode normalization, and doesn't support the full range of CFWS spacing syntax.

Also remember that:

```text
net/mail
```

is not a complete MIME email-processing framework.

A real email can contain:

```text
Headers
   ↓
MIME multipart
   ├── text/plain
   ├── text/html
   ├── attachment
   └── attachment
```

For more complicated MIME processing, you'll typically combine `net/mail` with other standard-library packages such as `mime` and `mime/multipart`.

---

# Quick API Cheat Sheet

| API | Purpose |
|---|---|
| `mail.ParseAddress()` | Parse one email address |
| `mail.ParseAddressList()` | Parse multiple addresses |
| `mail.ParseDate()` | Parse an RFC 5322 date |
| `mail.ReadMessage()` | Read an email's headers/body |
| `Address.String()` | Format an address |
| `AddressParser.Parse()` | Configurable single-address parsing |
| `AddressParser.ParseList()` | Configurable address-list parsing |
| `Header.Get()` | Get a header value |
| `Header.Date()` | Get and parse `Date` |
| `Header.AddressList()` | Get and parse address header |
| `mail.ErrHeaderNotPresent` | Indicates a requested header is absent |

---

# 🤔 Thought-provoking question

Imagine you're building a service that receives **1 million emails per day**.

You use `net/mail` to parse every incoming message, but an attacker deliberately sends malformed or extremely unusual email headers designed to consume CPU or memory.

**Where would you draw the boundary between "the email parser should accept as much RFC-compliant input as possible" and "the application should reject or limit suspicious input before parsing"?**

Think about **security, resource exhaustion, RFC compatibility, and reliability**—not just whether an email address is syntactically valid.
