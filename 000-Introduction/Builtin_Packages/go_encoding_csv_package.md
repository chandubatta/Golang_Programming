# Go `encoding/csv` Package

`encoding/csv` is an important Go standard-library package for working with **CSV (Comma-Separated Values)** data. It is useful for data import/export, reports, ETL pipelines, configuration data, and batch processing.

## 1. What is `encoding/csv`?

The `encoding/csv` package provides functionality for reading and writing CSV data.

Example CSV:

```text
id,name,age
1,Chandu,25
2,Ravi,30
3,Sita,28
```

Import it with:

```go
import "encoding/csv"
```

The package handles:

- Reading CSV files
- Writing CSV files
- Parsing records
- Quoted fields
- Commas inside fields
- Different delimiters
- Malformed CSV errors
- CSV streams

### When is it commonly used?

- Importing customer data
- Exporting database records
- Generating reports
- Processing large datasets
- Migrating data
- Reading spreadsheet-exported CSV files
- ETL/data-processing pipelines

---

# 2. Important Types in `encoding/csv`

| Type | Purpose |
|---|---|
| `csv.Reader` | Reads and parses CSV data |
| `csv.Writer` | Writes CSV data |
| `csv.ParseError` | Represents a CSV parsing error |

The basic flow is:

```text
CSV file
   |
   v
csv.Reader
   |
   v
[]string
   |
   v
Go application
```

For writing:

```text
Go application
   |
   v
[]string
   |
   v
csv.Writer
   |
   v
CSV file
```

---

# 3. `csv.Reader`

`csv.Reader` reads records from an input source such as a file, string, network stream, `bytes.Buffer`, or any `io.Reader`.

Example:

```go
file, err := os.Open("users.csv")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

reader := csv.NewReader(file)

record, err := reader.Read()
if err != nil {
    log.Fatal(err)
}

fmt.Println(record)
```

If the CSV contains:

```text
1,Chandu,25
```

the record is:

```go
[]string{"1", "Chandu", "25"}
```

---

# 4. `csv.NewReader()`

### Function

```go
func NewReader(r io.Reader) *Reader
```

Creates a new `csv.Reader`.

### Parameter

```go
r io.Reader
```

The source from which CSV data will be read.

Example:

```go
file, _ := os.Open("users.csv")
reader := csv.NewReader(file)
```

It can also read from a string:

```go
data := `id,name
1,Chandu
2,Ravi`

reader := csv.NewReader(strings.NewReader(data))
```

`NewReader()` creates the reader; actual reading happens through methods such as `Read()`.

---

# 5. `Reader.Read()`

### Function

```go
func (r *Reader) Read() (record []string, err error)
```

Reads **one CSV record**.

Example:

```go
record, err := reader.Read()

if err != nil {
    log.Fatal(err)
}

fmt.Println(record)
```

For:

```text
id,name,age
1,Chandu,25
```

successive calls return:

```go
[]string{"id", "name", "age"}
```

and:

```go
[]string{"1", "Chandu", "25"}
```

When there is no more data, it returns:

```go
io.EOF
```

A common pattern:

```go
for {
    record, err := reader.Read()

    if err == io.EOF {
        break
    }

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(record)
}
```

---

# 6. `Reader.ReadAll()`

### Function

```go
func (r *Reader) ReadAll() (records [][]string, err error)
```

Reads **all remaining CSV records**.

Example:

```go
records, err := reader.ReadAll()
if err != nil {
    log.Fatal(err)
}

for _, record := range records {
    fmt.Println(record)
}
```

For:

```text
id,name
1,Chandu
2,Ravi
3,Sita
```

you get:

```go
[][]string{
    {"id", "name"},
    {"1", "Chandu"},
    {"2", "Ravi"},
    {"3", "Sita"},
}
```

### `Read()` vs `ReadAll()`

`Read()`:

```text
Read one record
```

`ReadAll()`:

```text
Read everything
```

For very large CSV files, processing records one at a time with `Read()` is generally preferable because `ReadAll()` keeps all records in memory.

---

# 7. `Reader.ReuseRecord`

`ReuseRecord` is a field, not a function:

```go
reader.ReuseRecord = true
```

It allows `Read()` to reuse the backing array of the returned record, which can reduce allocations when processing many records.

However, if you need to keep a record after another call to `Read()`, don't assume the returned slice remains independent when `ReuseRecord` is enabled.

For beginner programs, leaving the default:

```go
false
```

is usually easier.

---

# 8. `Reader.FieldsPerRecord`

Controls how many fields each record should contain.

Default:

```go
reader.FieldsPerRecord = 0
```

With `0`, the first record determines the expected number of fields.

Example:

```text
id,name,age
1,Chandu,25
2,Ravi
```

The first record has 3 fields, while the second has 2, so the reader reports an error.

You can explicitly specify:

```go
reader.FieldsPerRecord = 3
```

Now every record must contain exactly 3 fields.

---

# 9. `Reader.LazyQuotes`

```go
reader.LazyQuotes = true
```

Normally CSV quoting rules are strictly validated. `LazyQuotes` makes the reader more tolerant of incorrect quoting.

It can be useful when processing poorly formatted CSV files from external systems.

Do not enable it blindly; strict parsing is generally safer.

---

# 10. `Reader.TrimLeadingSpace`

```go
reader.TrimLeadingSpace = true
```

Removes spaces immediately following delimiters.

For:

```text
1, Chandu, 25
```

the reader can interpret fields without those leading spaces.

---

# 11. `Reader.Comment`

```go
reader.Comment = '#'
```

Specifies a character used for comments.

Example:

```text
# This is a comment
id,name
1,Chandu
2,Ravi
```

With:

```go
reader.Comment = '#'
```

the comment line is ignored.

The comment character must satisfy the CSV reader's validity rules; it cannot be a quote or delimiter.

---

# 12. `Reader.Comma`

Default:

```go
reader.Comma = ','
```

CSV does not necessarily have to use commas.

For:

```text
id;name;age
1;Chandu;25
```

use:

```go
reader.Comma = ';'
```

For tab-separated data:

```go
reader.Comma = '\t'
```

This is useful for semicolon-separated files, tab-separated data, and custom-delimited files.

---

# 13. `Reader.InputOffset()`

### Function

```go
func (r *Reader) InputOffset() int64
```

Returns the current input offset.

Example:

```go
offset := reader.InputOffset()
fmt.Println("Current offset:", offset)
```

Useful for:

- Debugging
- Tracking processing progress
- Large-file processing
- Error reporting

---

# 14. `Reader.FieldPos()`

### Function

```go
func (r *Reader) FieldPos(field int) (line, column int)
```

Returns the line and column position of a field from the most recently read record.

Example:

```go
record, err := reader.Read()
if err != nil {
    log.Fatal(err)
}

line, column := reader.FieldPos(1)

fmt.Println(line, column)
```

This is useful for detailed error messages such as:

```text
Invalid age at line 5, column 12
```

Field indexes are zero-based:

```text
field 0 -> first field
field 1 -> second field
field 2 -> third field
```

---

# 15. `csv.Writer`

`csv.Writer` converts Go data into CSV format.

Conceptually:

```text
[]string
   |
   v
csv.Writer
   |
   v
CSV
```

---

# 16. `csv.NewWriter()`

### Function

```go
func NewWriter(w io.Writer) *Writer
```

Creates a new CSV writer.

Example:

```go
file, err := os.Create("users.csv")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

writer := csv.NewWriter(file)
```

`NewWriter()` creates the writer. Output is written through methods such as `Write()` and flushed with `Flush()`.

---

# 17. `Writer.Write()`

### Function

```go
func (w *Writer) Write(record []string) error
```

Writes one CSV record.

Example:

```go
err := writer.Write([]string{
    "1",
    "Chandu",
    "25",
})

if err != nil {
    log.Fatal(err)
}
```

Output:

```text
1,Chandu,25
```

It also handles CSV escaping.

For example:

```go
writer.Write([]string{
    "1",
    "Chandu, Batta",
    "25",
})
```

produces:

```text
1,"Chandu, Batta",25
```

This is a major reason to use `encoding/csv` instead of manually concatenating strings.

---

# 18. `Writer.WriteAll()`

### Function

```go
func (w *Writer) WriteAll(records [][]string) error
```

Writes multiple CSV records.

Example:

```go
records := [][]string{
    {"id", "name", "age"},
    {"1", "Chandu", "25"},
    {"2", "Ravi", "30"},
    {"3", "Sita", "28"},
}

err := writer.WriteAll(records)
if err != nil {
    log.Fatal(err)
}
```

This is convenient when all records are already available.

---

# 19. `Writer.Flush()`

### Function

```go
func (w *Writer) Flush()
```

Flushes buffered CSV data to the underlying `io.Writer`.

Important:

```go
writer.Write([]string{"1", "Chandu", "25"})
writer.Flush()
```

A common pattern is:

```go
writer := csv.NewWriter(file)
defer writer.Flush()
```

---

# 20. `Writer.Error()`

### Function

```go
func (w *Writer) Error() error
```

Returns any error that occurred during writing or flushing.

Example:

```go
writer.Write([]string{"1", "Chandu", "25"})
writer.Flush()

if err := writer.Error(); err != nil {
    log.Fatal(err)
}
```

A robust pattern is:

```go
writer.Flush()

if err := writer.Error(); err != nil {
    log.Fatal(err)
}
```

---

# 21. `Writer.Comma`

Default:

```go
writer.Comma = ','
```

You can change it:

```go
writer.Comma = ';'
```

Then:

```go
writer.Write([]string{
    "1",
    "Chandu",
    "25",
})
```

produces:

```text
1;Chandu;25
```

---

# 22. `Writer.UseCRLF`

```go
writer.UseCRLF = true
```

Controls whether records use Windows-style CRLF line endings:

```text
\r\n
```

instead of:

```text
\n
```

This can be useful for systems that expect CRLF line endings.

---

# 23. `csv.ParseError`

`ParseError` represents an error encountered while parsing CSV.

Conceptually, it contains information such as:

```go
type ParseError struct {
    StartLine int
    Line      int
    Column    int
    Err       error
}
```

You can inspect it:

```go
record, err := reader.Read()

if err != nil {
    var parseErr *csv.ParseError

    if errors.As(err, &parseErr) {
        fmt.Println("Line:", parseErr.Line)
        fmt.Println("Column:", parseErr.Column)
    }
}
```

This is particularly useful when processing CSV files supplied by users or external systems.

---

# 24. Complete Simple CSV Reader Example

Suppose `users.csv` contains:

```text
id,name,age
1,Chandu,25
2,Ravi,30
3,Sita,28
```

Program:

```go
package main

import (
    "encoding/csv"
    "fmt"
    "io"
    "log"
    "os"
)

func main() {
    file, err := os.Open("users.csv")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    reader := csv.NewReader(file)

    for {
        record, err := reader.Read()

        if err == io.EOF {
            break
        }

        if err != nil {
            log.Fatal(err)
        }

        fmt.Println(record)
    }
}
```

Output:

```text
[id name age]
[1 Chandu 25]
[2 Ravi 30]
[3 Sita 28]
```

---

# 25. Complete CSV Writer Example

```go
package main

import (
    "encoding/csv"
    "log"
    "os"
)

func main() {
    file, err := os.Create("users.csv")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    writer.Write([]string{"id", "name", "age"})
    writer.Write([]string{"1", "Chandu", "25"})
    writer.Write([]string{"2", "Ravi", "30"})
    writer.Write([]string{"3", "Sita", "28"})

    if err := writer.Error(); err != nil {
        log.Fatal(err)
    }
}
```

Result:

```text
id,name,age
1,Chandu,25
2,Ravi,30
3,Sita,28
```

---

# 26. Important CSV Concept: Quoting

Consider:

```text
1,Chandu Batta,25
```

Easy.

But:

```text
1,"Chandu, Batta",25
```

The comma inside the name does **not** represent another column.

`encoding/csv` understands this automatically.

That's why manually doing:

```go
fields := strings.Split(line, ",")
```

is dangerous.

For example:

```go
line := `1,"Chandu, Batta",25`
fields := strings.Split(line, ",")
```

can incorrectly produce something like:

```text
1
"Chandu
 Batta"
25
```

`csv.Reader` correctly understands the quoted comma.

---

# 27. Common Mistake #1 — Using `strings.Split()`

### Wrong approach

```go
fields := strings.Split(line, ",")
```

This does not properly implement CSV parsing.

CSV supports rules involving:

- Quoted fields
- Embedded commas
- Embedded newlines
- Escaped quotes

### Correct approach

```go
reader := csv.NewReader(...)
record, err := reader.Read()
```

Use `encoding/csv` when you are actually dealing with CSV.

---

# 28. Common Mistake #2 — Forgetting `Flush()`

Beginners sometimes write:

```go
writer.Write(record)
```

and assume the data is immediately written.

`csv.Writer` buffers output.

Use:

```go
writer.Flush()
```

or:

```go
defer writer.Flush()
```

For robust error handling:

```go
writer.Flush()

if err := writer.Error(); err != nil {
    log.Fatal(err)
}
```

---

# 29. Common Mistake #3 — Assuming Every CSV Has Exactly Three Columns

You might write:

```go
record, _ := reader.Read()

fmt.Println(record[0])
fmt.Println(record[1])
fmt.Println(record[2])
```

But external CSV data may be malformed:

```text
1,Chandu
```

or:

```text
1,Chandu,25,India
```

Use `FieldsPerRecord` when you need strict validation:

```go
reader.FieldsPerRecord = 3
```

Or validate explicitly before accessing indexes.

---

# 30. Real-World Application #1 — Importing Customer Data

Imagine an application receives:

```text
customers.csv
```

```text
id,name,email,city
101,Chandu,chandu@example.com,Mumbai
102,Ravi,ravi@example.com,Hyderabad
103,Sita,sita@example.com,Delhi
```

A Go application can process it as:

```text
CSV
 |
 v
csv.Reader
 |
 v
Validate fields
 |
 v
Convert strings to appropriate types
 |
 v
Business validation
 |
 v
Database
```

This is common in:

- CRM systems
- Banking systems
- E-commerce
- Employee management
- Data migration

---

# 31. Real-World Application #2 — Generating Reports

Suppose your application calculates sales:

```text
Product,Quantity,Revenue
Laptop,20,1500000
Phone,50,2000000
Monitor,30,450000
```

Go can generate this CSV:

```text
Database
    |
    v
Go application
    |
    v
csv.Writer
    |
    v
sales-report.csv
```

Users can open the file with spreadsheet software or upload it to another system.

---

# 32. Exercise 1 — Basic CSV Reader

Create a program that reads:

```text
students.csv
```

with:

```text
id,name,age
1,Chandu,25
2,Ravi,23
3,Sita,24
```

Your program should:

1. Open the file.
2. Create a `csv.Reader`.
3. Read every record.
4. Print each student's ID, name, and age.
5. Correctly handle `io.EOF`.
6. Handle file and CSV parsing errors.
7. Do not use `strings.Split()`.

**No solution is provided.**

---

# 33. Exercise 2 — CSV Validation and Processing

Create a program that reads a CSV file containing:

```text
id,name,age,email
```

The program should:

1. Validate that every record has exactly four fields.
2. Validate that `id` is a valid integer.
3. Validate that `age` is a valid integer.
4. Validate that the email field isn't empty.
5. Separate valid and invalid records.
6. Write invalid records to:
   ```text
   invalid_users.csv
   ```
7. Print a summary showing how many records were valid and invalid.

**No solution is provided.**

---

# 34. Exercise 3 — CSV Data Processing Pipeline

Build a small **CSV-to-report pipeline**.

Input:

```text
sales.csv
```

containing:

```text
product,category,quantity,price
Laptop,Electronics,5,75000
Phone,Electronics,10,30000
Chair,Furniture,20,5000
Desk,Furniture,10,12000
```

Your program should:

1. Read the CSV.
2. Validate every row.
3. Convert `quantity` and `price` into numeric values.
4. Calculate total revenue for every product.
5. Calculate total revenue for every category.
6. Find the highest-revenue product.
7. Generate:
   ```text
   sales_report.csv
   ```
   containing appropriate summary information.
8. Handle malformed rows without crashing the entire program.
9. Report the line number of malformed records.
10. Make the program work with a large CSV without loading the entire file into memory.

This exercise combines `csv.Reader`, validation, type conversion, `csv.Writer`, error handling, and streaming.

**No solution is provided.**

---

# 35. `encoding/csv` Function Cheat Sheet

| Function / Method | Purpose |
|---|---|
| `csv.NewReader()` | Create a CSV reader |
| `Reader.Read()` | Read one CSV record |
| `Reader.ReadAll()` | Read all remaining records |
| `Reader.FieldPos()` | Get line/column position of a field |
| `Reader.InputOffset()` | Get current input offset |
| `csv.NewWriter()` | Create a CSV writer |
| `Writer.Write()` | Write one CSV record |
| `Writer.WriteAll()` | Write multiple records |
| `Writer.Flush()` | Flush buffered output |
| `Writer.Error()` | Get writing/flush errors |

Important configuration fields:

```text
Reader:
    Comma
    Comment
    FieldsPerRecord
    LazyQuotes
    TrimLeadingSpace
    ReuseRecord

Writer:
    Comma
    UseCRLF
```

---

# 36. Big Picture

Remember `encoding/csv` like this:

```text
                    encoding/csv
                         |
             +-----------+-----------+
             |                       |
          Reader                  Writer
             |                       |
       +-----+-----+           +-----+-----+
       |           |           |           |
     Read       ReadAll      Write      WriteAll
       |                       |
       +-----------+   +-------+
                   |   |
                CSV Data
```

The most important lesson is:

> **CSV is more complicated than simply splitting a string by commas.**

CSV has rules for quoting, escaping, delimiters, embedded commas, and even embedded newlines. `encoding/csv` handles those rules for you.

---

# 37. Thought-Provoking Question

Suppose your Go application receives a **10-million-row CSV file from an external company**, and some rows contain malformed data, quoted commas, missing fields, and unexpected values.

**Would you choose `ReadAll()` or process the file one record at a time with `Read()`? Why—and how would your design balance memory usage, error handling, performance, and the requirement to report exactly which rows failed?**
