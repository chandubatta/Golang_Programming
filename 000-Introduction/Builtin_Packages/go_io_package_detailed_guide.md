# Go `io` Package — Detailed Guide

The Go `io` package is one of the most important packages to understand because many other packages—`os`, `net/http`, `bytes`, `strings`, `bufio`, etc.—use its interfaces.

> **Core idea:** `io` is mainly about moving data from a source to a destination.

```text
File ──────> Reader ──────> Your Program ──────> Writer ──────> File
Network ───> Reader ──────> Your Program ──────> Writer ──────> HTTP response
Memory ────> Reader ──────> Your Program ──────> Writer ──────> Memory
```

---

# 1. What is the `io` package?

The `io` package is part of Go's standard library:

```go
import "io"
```

It provides:

- Interfaces for reading data
- Interfaces for writing data
- Functions for copying data
- Functions for reading exact amounts of data
- Functions for combining readers/writers
- Functions for limiting data
- Functions for connecting readers and writers
- Utilities for working with streams

The two most important interfaces are:

```go
io.Reader
io.Writer
```

## `io.Reader`

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

A `Reader` represents something that **provides data**.

Examples:

```go
os.File
strings.Reader
bytes.Reader
http.Request.Body
```

Conceptually:

```text
Source
  |
  | Read()
  v
[]byte
```

## `io.Writer`

```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

A `Writer` represents something that **accepts data**.

Examples:

```go
os.File
os.Stdout
bytes.Buffer
http.ResponseWriter
```

Conceptually:

```text
[]byte
  |
  | Write()
  v
Destination
```

This abstraction is extremely powerful.

A function can accept an `io.Reader` without caring whether the data comes from:

- a file
- network
- memory
- a string
- another program

That is the major reason the `io` package is so important.

---

# 2. Simple example

Let's start with the most important concept: copying from a `Reader` to a `Writer`.

```go
package main

import (
    "io"
    "os"
    "strings"
)

func main() {
    reader := strings.NewReader("Hello from Go io package!")

    _, err := io.Copy(os.Stdout, reader)
    if err != nil {
        panic(err)
    }
}
```

Output:

```text
Hello from Go io package!
```

Here:

```go
strings.NewReader(...)
```

creates a `Reader`.

And:

```go
os.Stdout
```

is a `Writer`.

Then:

```go
io.Copy(os.Stdout, reader)
```

moves the data:

```text
strings.Reader
      |
      | io.Copy()
      v
  os.Stdout
```

`io.Copy` keeps copying until the reader reaches `EOF` or another error occurs.

---

# 3. Important `io` interfaces

Before learning the functions, you should understand these interfaces.

## 3.1 `io.Reader`

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

The most important interface in the package.

Example:

```go
func readSomething(r io.Reader) {
    // r can be a file, network connection,
    // memory buffer, etc.
}
```

The function doesn't care about the actual source.

---

## 3.2 `io.Writer`

```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

Represents something that accepts data.

For example:

```go
func saveSomething(w io.Writer) {
    // w could be a file,
    // network connection,
    // bytes.Buffer, etc.
}
```

---

## 3.3 `io.ReadWriter`

Combines `Reader` and `Writer`:

```go
type ReadWriter interface {
    Reader
    Writer
}
```

Therefore a `ReadWriter` can:

```text
READ  <---->  WRITE
```

---

## 3.4 `io.ReadCloser`

```go
type ReadCloser interface {
    Reader
    Closer
}
```

It can read data and be closed.

A common example is:

```go
http.Request.Body
```

---

## 3.5 `io.WriteCloser`

```go
type WriteCloser interface {
    Writer
    Closer
}
```

It can write and then be closed.

Files commonly behave this way.

---

## 3.6 `io.ReadSeeker`

```go
type ReadSeeker interface {
    Reader
    Seeker
}
```

It can read and change its position.

For example:

```text
0 ---- 1 ---- 2 ---- 3 ---- 4
                ^
              current
```

You can move the position backward or forward using `Seek`.

---

## 3.7 `io.WriteSeeker`

```go
type WriteSeeker interface {
    Writer
    Seeker
}
```

It can write and move its position.

---

## 3.8 `io.ReadWriteSeeker`

Combines:

```go
Reader
Writer
Seeker
```

Useful for files where you need both reading/writing and positioning.

---

# 4. Every major exported function in `io`

The Go `io` package provides functions including:

- `Copy`
- `CopyBuffer`
- `CopyN`
- `Pipe`
- `ReadAll`
- `ReadAtLeast`
- `ReadFull`
- `WriteString`
- `LimitReader`
- `MultiReader`
- `MultiWriter`
- `TeeReader`
- `NewSectionReader`
- `NopCloser`

Let's go through them.

---

## 4.1 `io.Copy`

Signature:

```go
func Copy(dst Writer, src Reader) (written int64, err error)
```

Purpose:

> Copy everything from a `Reader` to a `Writer`.

Example:

```go
reader := strings.NewReader("Hello Go")

n, err := io.Copy(os.Stdout, reader)

fmt.Println("\nBytes copied:", n)
```

Conceptually:

```text
Reader
   |
   | Copy
   v
Writer
```

It reads until:

```text
EOF
```

or an error occurs.

A successful `io.Copy` returns `nil`, not `io.EOF`.

### Real-world example

Copy a file:

```go
src, _ := os.Open("input.txt")
defer src.Close()

dst, _ := os.Create("output.txt")
defer dst.Close()

_, err := io.Copy(dst, src)
```

---

## 4.2 `io.CopyBuffer`

Signature:

```go
func CopyBuffer(dst Writer, src Reader, buf []byte) (written int64, err error)
```

Similar to `Copy`, but you provide the buffer.

```go
buffer := make([]byte, 32*1024)

_, err := io.CopyBuffer(dst, src, buffer)
```

Why use it?

You may want to control the memory buffer being used during copying.

Conceptually:

```text
Reader
   |
   v
[ custom buffer ]
   |
   v
Writer
```

Useful when:

- controlling allocations
- reusing buffers
- processing many copy operations

---

## 4.3 `io.CopyN`

Signature:

```go
func CopyN(dst Writer, src Reader, n int64) (written int64, err error)
```

Copies up to the requested number of bytes, stopping after `n`.

Example:

```go
reader := strings.NewReader("ABCDEFGHIJK")

n, err := io.CopyN(os.Stdout, reader, 5)
```

Output:

```text
ABCDE
```

This is useful when you don't want to copy the entire stream.

---

## 4.4 `io.ReadAll`

Signature:

```go
func ReadAll(r Reader) ([]byte, error)
```

Reads the entire reader into memory.

Example:

```go
reader := strings.NewReader("Hello Go")

data, err := io.ReadAll(reader)
if err != nil {
    panic(err)
}

fmt.Println(string(data))
```

Output:

```text
Hello Go
```

### Important warning

Don't blindly use `ReadAll` for potentially huge input.

For example:

```go
data, _ := io.ReadAll(hugeFile)
```

could consume a large amount of memory.

For large streams, process data incrementally instead.

`ReadAll` reads until EOF or another error and returns the accumulated data.

---

## 4.5 `io.ReadAtLeast`

Signature:

```go
func ReadAtLeast(r Reader, buf []byte, min int) (n int, err error)
```

It attempts to read **at least `min` bytes**.

Example:

```go
buf := make([]byte, 10)

n, err := io.ReadAtLeast(reader, buf, 5)
```

You are saying:

> "I need at least 5 bytes."

If fewer than the required amount are available, you can get:

```go
io.ErrUnexpectedEOF
```

depending on the situation.

Useful when processing data structures where a minimum amount of data is required.

---

## 4.6 `io.ReadFull`

Signature:

```go
func ReadFull(r Reader, buf []byte) (n int, err error)
```

This is stricter than `ReadAtLeast`.

If:

```go
buf := make([]byte, 10)
```

then `ReadFull` tries to fill **all 10 bytes**.

Example:

```go
buf := make([]byte, 10)

n, err := io.ReadFull(reader, buf)
```

Think:

```text
ReadAtLeast:
"I need at least X bytes."

ReadFull:
"I need this entire buffer filled."
```

This is particularly useful when reading fixed-size binary structures.

---

## 4.7 `io.WriteString`

Signature:

```go
func WriteString(w Writer, s string) (n int, err error)
```

Writes a string to an `io.Writer`.

Example:

```go
io.WriteString(os.Stdout, "Hello from Go!")
```

Instead of:

```go
w.Write([]byte("Hello from Go!"))
```

you can use:

```go
io.WriteString(w, "Hello from Go!")
```

This is convenient when you already have a string.

---

## 4.8 `io.LimitReader`

Signature:

```go
func LimitReader(r Reader, n int64) Reader
```

Creates a new reader that stops after `n` bytes.

Example:

```go
reader := strings.NewReader("ABCDEFGHIJK")

limited := io.LimitReader(reader, 5)

data, _ := io.ReadAll(limited)

fmt.Println(string(data))
```

Output:

```text
ABCDE
```

Think:

```text
Original Reader
ABCDEFGHIJK

        |
        | LimitReader(..., 5)
        v

Limited Reader
ABCDE
```

### Real-world importance

Suppose a user uploads a file.

You expect:

```text
maximum 10 MB
```

You don't necessarily want to blindly read unlimited data.

`io.LimitReader` can put a boundary around how much your code reads.

---

## 4.9 `io.MultiReader`

Signature:

```go
func MultiReader(readers ...Reader) Reader
```

Combines multiple readers into one logical reader.

Example:

```go
r1 := strings.NewReader("Hello ")
r2 := strings.NewReader("Go ")
r3 := strings.NewReader("World")

reader := io.MultiReader(r1, r2, r3)

data, _ := io.ReadAll(reader)

fmt.Println(string(data))
```

Output:

```text
Hello Go World
```

Conceptually:

```text
Reader 1 ──┐
           │
Reader 2 ──┼──> MultiReader ──> One Reader
           │
Reader 3 ──┘
```

It reads them sequentially.

---

## 4.10 `io.MultiWriter`

Signature:

```go
func MultiWriter(writers ...Writer) Writer
```

This does the opposite idea.

One write gets sent to multiple writers.

Example:

```go
var buffer1 bytes.Buffer
var buffer2 bytes.Buffer

writer := io.MultiWriter(&buffer1, &buffer2)

io.WriteString(writer, "Hello")
```

Now both contain:

```text
buffer1 → Hello
buffer2 → Hello
```

Conceptually:

```text
             ┌──> Writer 1
Input ───────┼──> Writer 2
             └──> Writer 3
```

Useful for simultaneously writing to:

- a file
- a log
- another destination

The writers are written sequentially; if one returns an error, the overall operation stops.

---

## 4.11 `io.TeeReader`

Signature:

```go
func TeeReader(r Reader, w Writer) Reader
```

This returns a reader that:

1. reads from the original reader
2. simultaneously writes the data it reads to another writer

Example:

```go
source := strings.NewReader("Hello Go")

var copy bytes.Buffer

reader := io.TeeReader(source, &copy)

data, _ := io.ReadAll(reader)

fmt.Println("Read:", string(data))
fmt.Println("Copied:", copy.String())
```

Conceptually:

```text
             ┌──────────> Your program
             |
Source ──────┤
             |
             └──────────> Writer
```

Unlike a buffered duplicate, the write must complete before the read completes.

Useful for:

```text
HTTP request
     |
     +----> process request
     |
     +----> audit/log/capture data
```

---

## 4.12 `io.Pipe`

Signature:

```go
func Pipe() (*PipeReader, *PipeWriter)
```

Creates a synchronous in-memory pipe.

Example:

```go
reader, writer := io.Pipe()
```

You get:

```text
PipeWriter ─────────> PipeReader
```

The writer writes data and the reader receives it.

It is particularly useful for connecting two pieces of code through a stream without first storing the entire data set in memory.

For example:

```text
Producer
   |
   | Write
   v
PipeWriter
   |
   v
PipeReader
   |
   | Read
   v
Consumer
```

One important property is that operations synchronize between the reader and writer rather than acting like an ordinary buffered queue.

---

## 4.13 `io.NopCloser`

Signature:

```go
func NopCloser(r Reader) ReadCloser
```

Sometimes you have:

```go
io.Reader
```

but an API requires:

```go
io.ReadCloser
```

`NopCloser` wraps the reader and provides a `Close()` method that does nothing.

Example:

```go
reader := strings.NewReader("Hello")

readCloser := io.NopCloser(reader)
```

Now:

```go
readCloser.Read(...)
readCloser.Close()
```

are available.

Very useful when an API requires a `ReadCloser`, but your underlying source doesn't need closing.

---

## 4.14 `io.NewSectionReader`

Signature:

```go
func NewSectionReader(
    r ReaderAt,
    off int64,
    n int64,
) *SectionReader
```

Creates a reader that exposes only a particular section of another `ReaderAt`.

Imagine a large file:

```text
0
|
v
+-------------------------------+
| Header | Data | Metadata | ...|
+-------------------------------+
         ^
         |
       section
```

You can create a reader for only a particular portion.

Example:

```go
section := io.NewSectionReader(file, 1000, 500)
```

This means:

```text
Start at byte 1000
Read at most 500 bytes
```

Very useful for:

- large files
- binary file formats
- archives
- random-access data

---

# 5. `SectionReader` methods

`SectionReader` itself has several methods.

## `Read`

```go
func (s *SectionReader) Read(p []byte) (n int, err error)
```

Reads sequentially from the selected section.

---

## `ReadAt`

```go
func (s *SectionReader) ReadAt(p []byte, off int64) (n int, err error)
```

Reads from a specific offset within the section.

This is useful when you need random access.

---

## `Seek`

```go
func (s *SectionReader) Seek(offset int64, whence int) (int64, error)
```

Moves the current reading position.

The `whence` values are:

```go
io.SeekStart
io.SeekCurrent
io.SeekEnd
```

For example:

```go
s.Seek(100, io.SeekStart)
```

means:

> Move to position 100 relative to the beginning.

---

## `Size`

```go
func (s *SectionReader) Size() int64
```

Returns the size of the section.

---

## `Outer`

```go
func (s *SectionReader) Outer() (r ReaderAt, off int64, n int64)
```

Returns information about the underlying reader and the section boundaries.

This can be useful when you need to recover the original reader and the section's:

```text
reader
offset
length
```

---

# 6. Important `io` interfaces beyond `Reader`/`Writer`

There are a few more that are worth knowing.

## `io.ReaderAt`

```go
type ReaderAt interface {
    ReadAt(p []byte, off int64) (n int, err error)
}
```

Unlike ordinary `Reader`, it lets you request:

> "Read these bytes starting at this exact offset."

Example concept:

```text
File:

0   1   2   3   4   5   6
A   B   C   D   E   F   G
            ^
            |
         ReadAt(..., 3)
```

This is useful for random-access data.

---

## `io.WriterAt`

```go
type WriterAt interface {
    WriteAt(p []byte, off int64) (n int, err error)
}
```

Writes at a particular offset.

---

## `io.Seeker`

```go
type Seeker interface {
    Seek(offset int64, whence int) (int64, error)
}
```

Allows movement within a data source.

---

## `io.ReaderFrom`

```go
type ReaderFrom interface {
    ReadFrom(r Reader) (n int64, err error)
}
```

Allows a destination to efficiently read directly from a `Reader`.

`io.Copy` can take advantage of this interface when available.

---

## `io.WriterTo`

```go
type WriterTo interface {
    WriteTo(w Writer) (n int64, err error)
}
```

Allows a source to efficiently write itself to a destination.

Again, `io.Copy` can use this when available.

---

## `io.RuneReader`

```go
type RuneReader interface {
    ReadRune() (r rune, size int, err error)
}
```

Reads one Unicode character at a time.

---

## `io.RuneScanner`

```go
type RuneScanner interface {
    RuneReader
    UnreadRune() error
}
```

Adds the ability to "undo" the last successful rune read.

---

## `io.StringWriter`

```go
type StringWriter interface {
    WriteString(s string) (n int, err error)
}
```

Represents a writer that can directly accept a string.

---

# 7. Important constants and errors

The package also provides seek constants:

```go
io.SeekStart
io.SeekCurrent
io.SeekEnd
```

And important errors:

```go
io.EOF
io.ErrUnexpectedEOF
io.ErrShortWrite
```

## `io.EOF`

Means:

> There is no more input.

A beginner mistake is thinking:

```go
if err == io.EOF {
    // something went wrong
}
```

Not necessarily.

`EOF` normally means:

```text
The reader has reached the end.
```

---

## `io.ErrUnexpectedEOF`

Means:

> The input ended before the expected amount of data was available.

For example, if you expected a 100-byte structure but only 70 bytes arrived.

---

## `io.ErrShortWrite`

Means:

> Fewer bytes were written than expected and no explicit error was returned.

---

# 8. Three common beginner mistakes

## Mistake 1: Assuming `Read` fills the buffer

Beginners might write:

```go
buf := make([]byte, 100)

n, _ := reader.Read(buf)
```

and assume:

```text
n == 100
```

That's not guaranteed.

A `Reader` may return fewer bytes than the buffer size.

### Better approach

If you require a specific number:

```go
io.ReadFull(...)
```

or:

```go
io.ReadAtLeast(...)
```

---

## Mistake 2: Treating `io.EOF` like a normal failure

`EOF` usually means:

```text
"The input is finished."
```

For example, a reader may eventually produce:

```go
err == io.EOF
```

That doesn't necessarily mean your program failed.

Also, functions such as `io.Copy` and `io.ReadAll` intentionally consume the reader until EOF and return `nil` on successful completion.

---

## Mistake 3: Using `io.ReadAll` on unlimited/huge data

This:

```go
data, err := io.ReadAll(r)
```

loads the entire stream into memory.

For a small JSON response:

```text
10 KB → fine
```

For a huge upload:

```text
10 GB → potentially disastrous
```

Instead, consider:

```go
io.Copy(...)
```

or:

```go
io.LimitReader(...)
```

or incremental processing.

---

# 9. Two real-world applications

## Application 1: File uploads/downloads

Suppose a web server receives a large file.

Instead of:

```text
Upload
   ↓
Read entire file into RAM
   ↓
Save file
```

you can stream:

```text
HTTP Request
     |
     | io.Reader
     v
   io.Copy
     |
     | io.Writer
     v
   Disk
```

This avoids loading the entire file into memory.

---

## Application 2: Streaming and data pipelines

Imagine:

```text
Database
   ↓
Reader
   ↓
Transformation
   ↓
Writer
   ↓
Compressed file
```

The `io.Reader` / `io.Writer` abstractions allow you to build components that don't care about the exact underlying source or destination.

For example:

```text
File Reader
     ↓
   io.Copy
     ↓
Network Writer
```

The same function can potentially work with different readers and writers because they implement the common interfaces.

---

# 10. The most important mental model

If you're learning Go, remember this:

```text
                    io.Reader
                       |
                       v
                 +-----------+
                 |   DATA    |
                 +-----------+
                       |
                       v
                    io.Writer
```

`Reader` answers:

> **"Where can I get bytes from?"**

`Writer` answers:

> **"Where can I send bytes to?"**

And `io` gives you tools to connect them.

For example:

```go
io.Copy(destination, source)
```

means:

> Take data from this `Reader` and continuously send it to this `Writer`.

That's the heart of Go's I/O design.

---

# 11. Three progressively challenging exercises

## Exercise 1 — Beginner: Copy a string

Create a Go program that:

1. Creates a `strings.Reader` containing a sentence.
2. Creates a `bytes.Buffer`.
3. Uses `io.Copy` to copy the data from the reader into the buffer.
4. Prints the contents of the buffer.
5. Prints how many bytes were copied.

**Goal:** Understand `io.Reader`, `io.Writer`, and `io.Copy`.

---

## Exercise 2 — Intermediate: Build a limited streaming reader

Create a program that contains a long string representing a large piece of data.

Your program must:

1. Create a `strings.Reader`.
2. Use `io.LimitReader` to expose only the first 20 bytes.
3. Use `io.Copy` to copy those bytes into another destination.
4. Verify that data after the first 20 bytes isn't copied.
5. Experiment by changing the limit to different values.

**Goal:** Understand how `LimitReader` controls how much data can be consumed.

---

## Exercise 3 — Advanced: Build a streaming pipeline

Build a small data-processing pipeline using:

```text
io.Reader
    ↓
io.TeeReader
    ↓
io.MultiWriter
    ↓
multiple destinations
```

Your program should:

1. Create a source containing a large text stream.
2. Create two separate output destinations.
3. Use `io.TeeReader` so the incoming data is simultaneously captured somewhere else.
4. Use `io.MultiWriter` to send data to multiple destinations.
5. Process the stream without using `io.ReadAll` on the original source.
6. Report how many bytes were processed.
7. Verify that the destinations received the expected data.

**Goal:** Understand how Go's I/O abstractions can be combined to create streaming pipelines.

---

# 12. One deeper thing to notice

The really powerful part of `io` isn't actually `io.Copy`.

It's **interface-based design**.

Consider:

```go
func process(r io.Reader) error {
    // process data
    return nil
}
```

You could call this with:

```text
file
 ↓
io.Reader
```

or:

```text
network connection
 ↓
io.Reader
```

or:

```text
strings.Reader
 ↓
io.Reader
```

or:

```text
bytes.Reader
 ↓
io.Reader
```

The function doesn't need to know the concrete type.

That's one of the major ideas behind Go's design:

> **Program against small interfaces rather than concrete implementations.**

---

# 13. Thought-provoking question

Imagine you are building a file-upload API that could receive **10,000 simultaneous uploads**, and each upload could be several gigabytes.

**Why would designing your application around `io.Reader` and streaming operations such as `io.Copy` potentially be much safer and more scalable than calling `io.ReadAll` for every request? What could happen to memory usage if all 10,000 uploads were read completely into memory at the same time?**

That question gets to the heart of **why Go's `io` package is designed around streams rather than forcing you to load entire datasets into memory**.
