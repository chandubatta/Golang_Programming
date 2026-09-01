# Go `bufio` Package — Detailed Guide

## 1. What is the `bufio` package?

`bufio` stands for **buffered I/O**.

It is a Go standard-library package used to make reading and writing data more efficient and convenient by introducing a **buffer** between your program and an underlying input/output source.

Import it with:

```go
import "bufio"
```

The package commonly works with:

- keyboard input (`os.Stdin`)
- files
- terminal output (`os.Stdout`)
- network connections
- strings
- other `io.Reader` and `io.Writer` implementations

The package provides three major types:

```text
bufio
 │
 ├── Reader
 │     └── Efficiently read data
 │
 ├── Scanner
 │     └── Conveniently read tokens/lines/words
 │
 ├── Writer
 │     └── Efficiently write data
 │
 └── ReadWriter
       └── Read + Write
```

---

## 2. Why do we need buffering?

Suppose you read a file one byte at a time.

Without buffering, your program may repeatedly communicate with the underlying file/device:

```text
Program
   ↓
File
   ↓
Program
   ↓
File
   ↓
Program
   ↓
File
```

That can be inefficient.

With buffering:

```text
             ┌──────────────┐
Program ────►│   Buffer     │
             └──────┬───────┘
                    │
                    ▼
                  File
```

The `bufio.Reader` can read a larger block into memory and then serve subsequent reads from that buffer.

Similarly, `bufio.Writer` can collect multiple writes and send them to the underlying destination in larger batches.

---

## 3. Simple `bufio` example

A very common beginner example is reading keyboard input line by line.

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")

	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Hello,", name)
}
```

### What's happening?

```go
os.Stdin
```

represents standard input.

Then:

```go
bufio.NewReader(os.Stdin)
```

creates a buffered reader around it.

Then:

```go
reader.ReadString('\n')
```

reads until it encounters a newline.

---

# 4. Main components of `bufio`

The package provides:

### Functions

```text
ScanBytes
ScanLines
ScanRunes
ScanWords
```

### Types

```text
Reader
Scanner
Writer
ReadWriter
SplitFunc
```

---

# 5. `Reader`

`bufio.Reader` is used for **buffered reading**.

```go
reader := bufio.NewReader(os.Stdin)
```

It wraps an `io.Reader`.

For example:

```text
os.Stdin
   ↓
bufio.Reader
   ↓
Your program
```

Important Reader methods include:

```text
NewReader
NewReaderSize

Buffered
Discard
Peek
Read
ReadByte
ReadBytes
ReadLine
ReadRune
ReadSlice
ReadString
Reset
Size
UnreadByte
UnreadRune
WriteTo
```

---

# 6. `bufio.NewReader()`

### Syntax

```go
bufio.NewReader(rd io.Reader) *bufio.Reader
```

Creates a new buffered reader using the default buffer size.

Example:

```go
reader := bufio.NewReader(os.Stdin)
```

Another example:

```go
file, err := os.Open("data.txt")
if err != nil {
	return
}
defer file.Close()

reader := bufio.NewReader(file)
```

### When to use?

Use it when you want convenient buffered reading from:

- files
- standard input
- network connections
- other readers

---

# 7. `bufio.NewReaderSize()`

### Syntax

```go
bufio.NewReaderSize(rd io.Reader, size int) *bufio.Reader
```

This lets you specify the minimum buffer size.

Example:

```go
reader := bufio.NewReaderSize(os.Stdin, 4096)
```

### Why use it?

Useful when you have particular performance or input-size requirements.

For most beginners:

```go
bufio.NewReader(...)
```

is enough.

---

# 8. `Reader.Read()`

It reads bytes into a byte slice.

Example:

```go
package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	reader := bufio.NewReader(strings.NewReader("Hello Go"))

	buffer := make([]byte, 5)

	n, err := reader.Read(buffer)

	fmt.Println("Bytes read:", n)
	fmt.Println("Data:", string(buffer[:n]))
	fmt.Println("Error:", err)
}
```

Possible output:

```text
Bytes read: 5
Data: Hello
Error: <nil>
```

### Important point

Beginners sometimes assume:

```go
reader.Read(buffer)
```

will always fill the entire buffer.

It does **not** guarantee that.

If you need exactly `len(buffer)` bytes, `io.ReadFull` is generally the appropriate tool.

---

# 9. `Reader.ReadByte()`

Reads exactly one byte.

```go
b, err := reader.ReadByte()
```

Example:

```go
reader := bufio.NewReader(strings.NewReader("Go"))

b, err := reader.ReadByte()

fmt.Println(string(b))
fmt.Println(err)
```

Output:

```text
G
<nil>
```

Then another call would read:

```text
o
```

### Useful for

Parsing input character-by-character.

---

# 10. `Reader.ReadRune()`

A **rune** represents a Unicode code point.

```go
r, size, err := reader.ReadRune()
```

Example:

```go
reader := bufio.NewReader(strings.NewReader("你好"))

r, size, err := reader.ReadRune()

fmt.Println("Rune:", string(r))
fmt.Println("Bytes:", size)
fmt.Println("Error:", err)
```

This is important because Unicode characters can occupy multiple bytes in UTF-8.

For example:

```text
ASCII character → usually 1 byte
Unicode character → potentially multiple bytes
```

Therefore, `ReadByte()` and `ReadRune()` are not interchangeable.

---

# 11. `Reader.ReadString()`

One of the most commonly used Reader methods.

### Syntax

```go
ReadString(delim byte)
```

It reads until the specified delimiter.

Example:

```go
reader := bufio.NewReader(os.Stdin)

fmt.Print("Enter your name: ")

name, err := reader.ReadString('\n')

if err != nil {
	fmt.Println(err)
	return
}

fmt.Println("Name:", name)
```

The delimiter is included in the returned string when found.

For example, input:

```text
Chandu
```

may produce:

```text
"Chandu\n"
```

So you may often need:

```go
name = strings.TrimSpace(name)
```

---

# 12. `Reader.ReadBytes()`

Very similar to `ReadString`, but returns a `[]byte`.

```go
data, err := reader.ReadBytes('\n')
```

Example:

```go
reader := bufio.NewReader(strings.NewReader("Hello\nWorld"))

data, err := reader.ReadBytes('\n')

fmt.Println(string(data))
fmt.Println(err)
```

Result:

```text
Hello

<nil>
```

### Difference

```text
ReadString → string
ReadBytes  → []byte
```

---

# 13. `Reader.ReadSlice()`

```go
data, err := reader.ReadSlice('\n')
```

Reads until the delimiter.

However, the returned data refers to the Reader's internal buffer and can become invalid after subsequent operations.

It is therefore more low-level than `ReadBytes()` or `ReadString()`.

This is useful when you want to minimize allocations and understand buffer management.

---

# 14. `Reader.ReadLine()`

```go
line, isPrefix, err := reader.ReadLine()
```

Reads a line without the line-ending characters.

Example:

```go
line, isPrefix, err := reader.ReadLine()
```

The important value is:

```go
isPrefix
```

If:

```text
isPrefix == true
```

the returned data represents only part of a very long line.

You need to continue reading the remaining pieces.

`ReadLine` is a low-level primitive; for most callers, `ReadBytes`, `ReadString`, or `Scanner` are easier.

---

# 15. `Reader.Peek()`

This is a particularly useful method.

```go
data, err := reader.Peek(5)
```

It looks at upcoming bytes **without consuming them**.

Example:

```go
reader := bufio.NewReader(strings.NewReader("Hello"))

data, err := reader.Peek(3)

fmt.Println(string(data))

b, _ := reader.ReadByte()

fmt.Println(string(b))
```

The `Peek` does not advance the reader.

So the next read still starts with:

```text
H
```

### Mental model

```text
Input:

H e l l o
↑
Reader position

Peek(3)

H e l
↑
Still here
```

The bytes returned by `Peek` are only valid until the next read operation.

---

# 16. `Reader.Discard()`

```go
reader.Discard(n)
```

Skips the next `n` bytes.

Example:

```go
reader := bufio.NewReader(strings.NewReader("ABCDEFG"))

reader.Discard(3)

data, _ := reader.ReadString('\n')

fmt.Println(data)
```

After discarding:

```text
ABC
```

the remaining input begins with:

```text
DEFG
```

Useful when parsing protocols or structured data where certain bytes should simply be skipped.

---

# 17. `Reader.Buffered()`

Returns the number of bytes currently available in the internal buffer.

```go
n := reader.Buffered()

fmt.Println(n)
```

This does **not** mean how many bytes remain in the entire input source.

It means how many bytes are currently buffered and available.

---

# 18. `Reader.Size()`

Returns the size of the Reader's underlying buffer.

```go
size := reader.Size()

fmt.Println(size)
```

For example:

```go
reader := bufio.NewReaderSize(os.Stdin, 4096)

fmt.Println(reader.Size())
```

---

# 19. `Reader.UnreadByte()`

Allows you to put the most recently read byte back.

Example:

```go
reader := bufio.NewReader(strings.NewReader("ABC"))

b, _ := reader.ReadByte()

fmt.Println(string(b))

err := reader.UnreadByte()

fmt.Println(err)

b, _ = reader.ReadByte()

fmt.Println(string(b))
```

Output:

```text
A
<nil>
A
```

Only the most recently read byte can be unread.

---

# 20. `Reader.UnreadRune()`

Similar concept, but specifically for the most recently read rune.

```go
r, _, _ := reader.ReadRune()

reader.UnreadRune()
```

Important distinction:

```text
UnreadByte → works with last byte read
UnreadRune → requires last operation to be ReadRune
```

The latter is stricter than `UnreadByte`.

---

# 21. `Reader.Reset()`

Changes the underlying reader and discards buffered data.

```go
reader.Reset(newReader)
```

Example:

```go
reader := bufio.NewReader(strings.NewReader("First"))

reader.Reset(strings.NewReader("Second"))

data, _ := reader.ReadString('\n')

fmt.Println(data)
```

This can be useful when reusing a Reader instead of allocating another one.

---

# 22. `Reader.WriteTo()`

```go
reader.WriteTo(writer)
```

Copies data from the buffered reader to a writer.

Conceptually:

```text
Reader
  ↓
WriteTo()
  ↓
Writer
```

Example:

```go
reader := bufio.NewReader(strings.NewReader("Hello Go"))

var output strings.Builder

n, err := reader.WriteTo(&output)

fmt.Println(n)
fmt.Println(output.String())
fmt.Println(err)
```

This is particularly useful when transferring data between streams.

---

# 23. `Scanner`

`bufio.Scanner` is designed for convenient token-based reading.

The default behavior is:

```text
one Scan()
    ↓
one line
```

Example:

```go
scanner := bufio.NewScanner(os.Stdin)

for scanner.Scan() {
	fmt.Println("Input:", scanner.Text())
}

if err := scanner.Err(); err != nil {
	fmt.Println("Error:", err)
}
```

The default split function is `ScanLines`.

---

# 24. `NewScanner()`

```go
scanner := bufio.NewScanner(reader)
```

Creates a Scanner.

Example:

```go
scanner := bufio.NewScanner(os.Stdin)
```

---

# 25. `Scanner.Scan()`

```go
scanner.Scan()
```

Moves the scanner to the next token.

Typical pattern:

```go
for scanner.Scan() {
	fmt.Println(scanner.Text())
}
```

When there are no more tokens, it returns `false`.

Errors should then be checked using:

```go
scanner.Err()
```

---

# 26. `Scanner.Text()`

Returns the current token as a string.

```go
for scanner.Scan() {
	text := scanner.Text()
	fmt.Println(text)
}
```

`Text()` returns the current token as a string.

---

# 27. `Scanner.Bytes()`

Returns the current token as `[]byte`.

```go
for scanner.Scan() {
	data := scanner.Bytes()

	fmt.Println(string(data))
}
```

Important:

```text
Bytes() → []byte
Text()  → string
```

The byte slice may be overwritten by the next call to `Scan()`, so copy it if you need to retain it.

---

# 28. `Scanner.Err()`

After scanning:

```go
if err := scanner.Err(); err != nil {
	fmt.Println("Error:", err)
}
```

This tells you whether scanning stopped because of an actual error.

Normal end-of-input is not reported as an error by `Scanner.Err()`.

---

# 29. `Scanner.Split()`

This changes how Scanner determines tokens.

Example:

```go
scanner.Split(bufio.ScanWords)
```

Now instead of lines:

```text
Hello Go World
```

you get:

```text
Hello
Go
World
```

Example:

```go
scanner := bufio.NewScanner(strings.NewReader("Go is powerful"))

scanner.Split(bufio.ScanWords)

for scanner.Scan() {
	fmt.Println(scanner.Text())
}
```

---

# 30. `Scanner.Buffer()`

Scanner has a maximum token size.

You can customize its buffer:

```go
scanner.Buffer(make([]byte, 1024), 1024*1024)
```

Here:

```text
initial buffer = 1 KB
maximum buffer = 1 MB
```

This is particularly important when processing long lines.

`Buffer` must be called before scanning begins.

---

# 31. Important Scanner limitation

A common misconception is:

> "Scanner can read arbitrarily large lines."

Not by default.

Scanner has a maximum token size.

For very large tokens or when you need more precise control over reading, `bufio.Reader` is often more appropriate.

---

# 32. Built-in `SplitFunc`s

`bufio` provides four important split functions:

```text
ScanLines
ScanWords
ScanBytes
ScanRunes
```

---

## `bufio.ScanLines`

The default Scanner split function.

```go
scanner.Split(bufio.ScanLines)
```

Reads input line-by-line.

Example:

```go
scanner := bufio.NewScanner(strings.NewReader(
	"Apple\nBanana\nOrange",
))

scanner.Split(bufio.ScanLines)

for scanner.Scan() {
	fmt.Println(scanner.Text())
}
```

Output:

```text
Apple
Banana
Orange
```

---

# 33. `bufio.ScanWords`

Splits input into space-separated words.

```go
scanner.Split(bufio.ScanWords)
```

Example:

```go
scanner := bufio.NewScanner(
	strings.NewReader("Go makes programming enjoyable"),
)

scanner.Split(bufio.ScanWords)

for scanner.Scan() {
	fmt.Println(scanner.Text())
}
```

Output:

```text
Go
makes
programming
enjoyable
```

The definition of whitespace follows `unicode.IsSpace`.

---

# 34. `bufio.ScanBytes`

Returns individual bytes as tokens.

Conceptually:

```text
Hello

↓

H
e
l
l
o
```

Example:

```go
scanner := bufio.NewScanner(strings.NewReader("Go"))

scanner.Split(bufio.ScanBytes)

for scanner.Scan() {
	fmt.Println(scanner.Text())
}
```

---

# 35. `bufio.ScanRunes`

Similar to `ScanBytes`, but operates on UTF-8 encoded runes.

This distinction matters for Unicode.

```text
byte-oriented → bytes
rune-oriented → Unicode code points
```

Use `ScanRunes` when you need Unicode-aware tokenization.

---

# 36. `SplitFunc`

`SplitFunc` is the type used by Scanner to decide how input should be divided into tokens.

Its signature is:

```go
type SplitFunc func(
	data []byte,
	atEOF bool,
) (
	advance int,
	token []byte,
	err error,
)
```

Conceptually:

```text
Input data
    ↓
 SplitFunc
    ↓
┌──────────┬──────────┬───────┐
│ advance  │  token   │ error │
└──────────┴──────────┴───────┘
```

This allows you to create custom tokenization rules.

For example, you could create a Scanner that separates records using:

```text
,
|
;
:
```

or some application-specific format.

---

# 37. `Writer`

`bufio.Writer` provides buffered writing.

```go
writer := bufio.NewWriter(os.Stdout)
```

Instead of immediately sending every write to the underlying writer:

```text
Program
  ↓
Writer Buffer
  ↓
Underlying Writer
```

multiple writes can be accumulated.

Important Writer methods include:

```text
NewWriter
NewWriterSize

Available
AvailableBuffer
Buffered
Flush
ReadFrom
Reset
Size
Write
WriteByte
WriteRune
WriteString
```

---

# 38. `NewWriter()`

```go
writer := bufio.NewWriter(os.Stdout)
```

Creates a Writer using the default buffer size.

Example:

```go
writer := bufio.NewWriter(os.Stdout)

writer.WriteString("Hello ")
writer.WriteString("Go!")

writer.Flush()
```

---

# 39. `NewWriterSize()`

Allows you to specify a buffer size.

```go
writer := bufio.NewWriterSize(os.Stdout, 4096)
```

Useful when you have specific performance requirements.

---

# 40. `Writer.Write()`

Writes bytes into the Writer's buffer.

```go
data := []byte("Hello Go")

n, err := writer.Write(data)
```

Remember:

```text
Write()
  ↓
buffer
  ↓
Flush()
  ↓
actual destination
```

Calling `Write()` does not necessarily mean the data has already reached the final destination.

---

# 41. `Writer.WriteString()`

Writes a string.

```go
writer.WriteString("Hello Go")
```

Example:

```go
writer := bufio.NewWriter(os.Stdout)

writer.WriteString("Hello ")
writer.WriteString("World!")

writer.Flush()
```

This is one of the most convenient Writer methods.

---

# 42. `Writer.WriteByte()`

Writes one byte.

```go
writer.WriteByte('A')
```

Example:

```go
writer.WriteByte('H')
writer.WriteByte('i')
writer.WriteByte('\n')
```

---

# 43. `Writer.WriteRune()`

Writes a Unicode code point.

```go
writer.WriteRune('中')
```

Useful when working with Unicode text.

Unlike `WriteByte`, it can write a rune that requires multiple UTF-8 bytes.

---

# 44. `Writer.Flush()`

**This is one of the most important `bufio.Writer` methods.**

```go
writer.Flush()
```

It sends buffered data to the underlying Writer.

Example:

```go
writer := bufio.NewWriter(os.Stdout)

writer.WriteString("Hello Go")

writer.Flush()
```

### Beginner mistake

Forgetting:

```go
Flush()
```

can result in buffered data not being written when you expect.

A common pattern is:

```go
writer := bufio.NewWriter(file)
defer writer.Flush()
```

For important file-writing code, consider explicitly handling the `Flush()` error rather than discarding it.

---

# 45. `Writer.Buffered()`

Returns how many bytes are currently stored in the Writer's buffer.

```go
n := writer.Buffered()

fmt.Println(n)
```

Example conceptually:

```text
Writer buffer:

[ H ][ e ][ l ][ l ][ o ]

Buffered() → 5
```

---

# 46. `Writer.Available()`

Returns how much unused space remains in the buffer.

```go
available := writer.Available()
```

Conceptually:

```text
Buffer size = 4096
Currently used = 100

Available() = 3996
```

---

# 47. `Writer.AvailableBuffer()`

This is a more advanced optimization-oriented method.

It provides a byte slice representing the available space in the Writer's buffer.

This can help reduce allocations in performance-sensitive code.

It is generally **not necessary for beginners**.

---

# 48. `Writer.Size()`

Returns the size of the Writer's underlying buffer.

```go
size := writer.Size()

fmt.Println(size)
```

---

# 49. `Writer.Reset()`

Changes the destination Writer and discards buffered data.

```go
writer.Reset(newDestination)
```

This can be useful when reusing the same Writer object.

---

# 50. `Writer.ReadFrom()`

Copies data from an `io.Reader` into the Writer.

```go
n, err := writer.ReadFrom(reader)
```

Conceptually:

```text
Reader
  ↓
ReadFrom()
  ↓
bufio.Writer
  ↓
Flush()
  ↓
Destination
```

This can be useful for stream-copying operations.

---

# 51. `ReadWriter`

`bufio.ReadWriter` combines:

```text
Reader + Writer
```

It contains pointers to a Reader and Writer.

You create one using:

```go
rw := bufio.NewReadWriter(reader, writer)
```

This can be useful in situations where an object needs both reading and writing capabilities.

---

# 52. `NewReadWriter()`

### Syntax

```go
bufio.NewReadWriter(r, w)
```

Example:

```go
reader := bufio.NewReader(os.Stdin)
writer := bufio.NewWriter(os.Stdout)

rw := bufio.NewReadWriter(reader, writer)
```

Now both reading and writing can be performed through `rw`.

---

# 53. Complete function/method map

| Component | Function / Method | Purpose |
|---|---|---|
| Reader | `NewReader` | Create buffered Reader |
| Reader | `NewReaderSize` | Create Reader with specified buffer size |
| Reader | `Read` | Read bytes |
| Reader | `ReadByte` | Read one byte |
| Reader | `ReadRune` | Read one Unicode rune |
| Reader | `ReadString` | Read until delimiter as string |
| Reader | `ReadBytes` | Read until delimiter as bytes |
| Reader | `ReadSlice` | Read until delimiter using internal buffer |
| Reader | `ReadLine` | Low-level line reading |
| Reader | `Peek` | Inspect upcoming bytes without consuming |
| Reader | `Discard` | Skip bytes |
| Reader | `Buffered` | Get currently buffered bytes |
| Reader | `Size` | Get buffer size |
| Reader | `UnreadByte` | Undo last byte read |
| Reader | `UnreadRune` | Undo last rune read |
| Reader | `Reset` | Reuse Reader with another source |
| Reader | `WriteTo` | Transfer Reader data to Writer |
| Scanner | `NewScanner` | Create Scanner |
| Scanner | `Scan` | Move to next token |
| Scanner | `Text` | Get token as string |
| Scanner | `Bytes` | Get token as bytes |
| Scanner | `Err` | Get scanning error |
| Scanner | `Split` | Choose tokenization strategy |
| Scanner | `Buffer` | Configure Scanner buffer |
| Writer | `NewWriter` | Create buffered Writer |
| Writer | `NewWriterSize` | Create Writer with specified size |
| Writer | `Write` | Write bytes |
| Writer | `WriteString` | Write string |
| Writer | `WriteByte` | Write byte |
| Writer | `WriteRune` | Write Unicode rune |
| Writer | `Flush` | Send buffered data |
| Writer | `Buffered` | Get buffered byte count |
| Writer | `Available` | Get unused buffer capacity |
| Writer | `AvailableBuffer` | Access available buffer |
| Writer | `Size` | Get buffer size |
| Writer | `Reset` | Reuse Writer |
| Writer | `ReadFrom` | Copy from Reader |
| ReadWriter | `NewReadWriter` | Combine Reader + Writer |
| Package | `ScanLines` | Split by lines |
| Package | `ScanWords` | Split by words |
| Package | `ScanBytes` | Split by bytes |
| Package | `ScanRunes` | Split by runes |

---

# 54. `bufio.Reader` vs `bufio.Scanner`

| Feature | Reader | Scanner |
|---|---|---|
| Low-level control | Excellent | Limited |
| Read lines | Yes | Yes |
| Read words | Manual | Easy with `ScanWords` |
| Read delimiter | Yes | Custom split |
| Large tokens | Better suited | Has token-size limit |
| Convenient API | Moderate | Very convenient |
| Custom parsing | Excellent | Excellent with SplitFunc |
| Typical beginner choice | When control is needed | When reading lines |

### Simple rule

Use:

```go
bufio.Scanner
```

when you want convenient token/line processing.

Use:

```go
bufio.Reader
```

when you need more control over exactly how bytes are consumed.

---

# 55. `bufio.Reader` vs `fmt.Scan`

You can use `fmt.Scan` for simple formatted input.

But `bufio` gives you more direct control over input buffering and tokenization.

For example:

```go
scanner := bufio.NewScanner(os.Stdin)
scanner.Split(bufio.ScanWords)
```

gives you a simple word-based input mechanism.

And:

```go
reader := bufio.NewReader(os.Stdin)
reader.ReadString('\n')
```

gives you line-based input.

---

# 56. Common mistake #1 — Forgetting `Flush()`

Incorrect:

```go
writer := bufio.NewWriter(file)

writer.WriteString("Hello")
```

Better:

```go
writer := bufio.NewWriter(file)

writer.WriteString("Hello")

err := writer.Flush()
if err != nil {
	fmt.Println(err)
}
```

Remember:

```text
Write()
  ↓
Buffer
  ↓
Flush()
  ↓
Destination
```

---

# 57. Common mistake #2 — Assuming Scanner has unlimited input size

Incorrect assumption:

```text
Scanner can process any size line.
```

It cannot.

For very large tokens, configure the buffer:

```go
scanner.Buffer(make([]byte, 1024), 1024*1024)
```

Or consider using `bufio.Reader` when you need more control.

---

# 58. Common mistake #3 — Confusing bytes and runes

Consider:

```go
reader.ReadByte()
```

versus:

```go
reader.ReadRune()
```

They operate at different conceptual levels:

```text
ReadByte
   ↓
raw byte

ReadRune
   ↓
Unicode code point
```

For Unicode text, choosing the wrong one can produce unexpected results.

---

# 59. Real-world application #1 — Processing log files

Imagine a server produces:

```text
2026-09-01 INFO Server started
2026-09-01 ERROR Database unavailable
2026-09-01 INFO Server restarted
```

You could process it line by line:

```go
file, err := os.Open("server.log")
if err != nil {
	panic(err)
}
defer file.Close()

scanner := bufio.NewScanner(file)

for scanner.Scan() {
	line := scanner.Text()

	// Analyze log line
	fmt.Println(line)
}

if err := scanner.Err(); err != nil {
	fmt.Println("Error:", err)
}
```

This is a very common `bufio.Scanner` use case.

---

# 60. Real-world application #2 — Efficient file generation

Suppose your application needs to generate a large report.

Instead of performing many small writes directly to a file:

```text
Write
Write
Write
Write
Write
...
```

you can use:

```go
writer := bufio.NewWriter(file)
```

and then:

```go
writer.WriteString(...)
```

followed by:

```go
writer.Flush()
```

This is useful for:

- CSV generation
- report generation
- log writing
- exporting data
- large batch processing
- text file generation

---

# 61. A complete practical example

Let's combine several concepts.

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	fmt.Fprint(writer, "Enter your full name: ")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintln(writer, "Error:", err)
		return
	}

	input = strings.TrimSpace(input)

	fmt.Fprintln(writer, "Hello,", input)
}
```

Here we use:

```text
bufio.NewReader()
bufio.NewWriter()
ReadString()
Fprint
Flush()
```

This gives you a good first practical understanding of buffered input/output.

---

# 62. Three progressively challenging exercises

## Exercise 1 — Basic Line Reader

Create a Go program that:

1. Creates a `bufio.Scanner` for `os.Stdin`.
2. Asks the user to enter five names.
3. Reads each name line by line.
4. Prints each name with its position.

Example interaction:

```text
Enter 5 names:

1. Ravi
2. Priya
3. Arun
4. Meena
5. Kiran
```

**Do not use `fmt.Scan` for reading the names.**

---

## Exercise 2 — Word Counter

Create a program that:

1. Accepts a sentence from the user.
2. Uses `bufio.Scanner`.
3. Changes the Scanner's split function to `bufio.ScanWords`.
4. Counts the number of words.
5. Prints each word.
6. Finally prints the total word count.

Example:

```text
Input:
Go is simple and powerful

Output:
Go
is
simple
and
powerful

Total words: 5
```

---

## Exercise 3 — Large Log File Analyzer

Create a program that:

1. Opens a log file using `os.Open`.
2. Creates a `bufio.Scanner`.
3. Reads the file line by line.
4. Counts:
   - total lines
   - `INFO` lines
   - `WARNING` lines
   - `ERROR` lines
5. Prints a summary.
6. Properly handles file and Scanner errors.
7. Consider what should happen if a log line is extremely long.
8. Modify the Scanner buffer appropriately.

Example summary:

```text
Log Analysis
------------
Total lines : 12500
INFO        : 9500
WARNING     : 2100
ERROR       : 900
```

**Do not use a solution from this guide—implement it yourself.**

---

# 63. The most important mental model

Try to remember `bufio` like this:

```text
                 bufio
                   │
        ┌──────────┼──────────┐
        │          │          │
      Reader     Scanner    Writer
        │          │          │
        ▼          ▼          ▼
      Read       Tokens      Write
        │          │          │
        └──────────┼──────────┘
                   │
                Buffering
```

Remember these four ideas:

```text
Reader  → controlled buffered reading
Scanner → convenient token/line reading
Writer  → buffered writing
Flush   → push Writer's buffered data to destination
```

If you understand those four concepts, you've understood the foundation of `bufio`.

---

# 64. Thought-provoking question

Imagine you are processing a **10 GB log file** containing millions of lines.

You could read the entire file into memory, or process it incrementally using `bufio.Scanner` or `bufio.Reader`.

**What trade-offs would you consider between memory usage, performance, token-size limitations, error handling, and the level of control you need—and under what circumstances would you choose `Scanner` over `Reader`?**
