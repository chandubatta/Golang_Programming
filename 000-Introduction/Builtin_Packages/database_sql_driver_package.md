# `database/sql/driver` Package in Go

## 1. Purpose and When It Is Used

The `database/sql/driver` package is a **low-level interface layer** used to implement database drivers for Go's `database/sql` package.

A useful mental model is:

```text
Your Go application
        │
        ▼
 database/sql
  (high-level API)
        │
        ▼
database/sql/driver
  (driver interfaces)
        │
        ▼
 MySQL / PostgreSQL / SQLite / Oracle / etc.
```

The important distinction is that **most application developers should use `database/sql`, not `database/sql/driver` directly**. The `driver` package exists primarily for people implementing database drivers.

The package defines interfaces that database drivers implement so they can be used by `database/sql`.

### Common situations where you use `database/sql/driver`

You would typically encounter it when:

- Writing a **new SQL database driver**
- Implementing support for a proprietary database
- Building a wrapper around an existing database protocol
- Adding custom connection/session behavior
- Implementing custom parameter types
- Working on advanced database tooling
- Reading or debugging the implementation of an existing Go database driver

For ordinary CRUD applications, you normally **do not implement these interfaces yourself**.

---

# 2. Important Architecture

The package contains many interfaces rather than high-level database functions.

The central relationship is:

```text
Driver
  │
  └── Connector
        │
        └── Conn
              │
              ├── Stmt
              ├── Rows
              └── Tx
```

Additional optional interfaces provide things such as:

```text
Context support
     │
     ├── DriverContext
     ├── ConnPrepareContext
     ├── ConnBeginTx
     ├── ExecerContext
     └── QueryerContext

Connection management
     │
     ├── Pinger
     ├── Validator
     └── SessionResetter

Advanced values
     │
     ├── NamedValueChecker
     ├── Valuer
     └── ValueConverter
```

The current Go documentation recommends that modern drivers implement `Connector`/`DriverContext`, and that connections support interfaces such as `Pinger`, `SessionResetter`, and `Validator` where appropriate.

---

# 3. Core Types and Interfaces — In Detail

Because `database/sql/driver` is primarily an interface package, many of the important "functions" you encounter are actually **methods required by interfaces**.

## A. `Driver`

```go
type Driver interface {
    Open(name string) (Conn, error)
}
```

`Driver` is the fundamental interface that a database driver must implement.

### `Open`

```go
Open(name string) (Conn, error)
```

Creates a new database connection.

`name` is usually a driver-specific DSN or connection string.

For example, conceptually:

```text
"username:password@tcp(localhost:3306)/mydb"
```

The returned `Conn` represents a connection to the database.

One important detail: the `database/sql` package manages connection pooling, so a driver generally shouldn't create its own pool simply to reuse connections.

---

# B. `DriverContext`

Modern drivers can implement:

```go
type DriverContext interface {
    OpenConnector(name string) (Connector, error)
}
```

### `OpenConnector`

```go
OpenConnector(name string) (Connector, error)
```

Creates a `Connector` from a DSN.

Instead of repeatedly parsing:

```text
username
password
host
port
database
options
```

for every connection, the driver can parse the configuration once and store it in a connector.

Conceptually:

```text
DSN
 │
 ▼
OpenConnector()
 │
 ▼
Connector
 │
 ├── Connect()
 ├── Connect()
 └── Connect()
```

This is one reason modern drivers generally prefer `DriverContext` + `Connector`.

---

# C. `Connector`

```go
type Connector interface {
    Connect(ctx context.Context) (Conn, error)
    Driver() Driver
}
```

A `Connector` represents a driver configured with a particular database configuration.

### `Connect`

```go
Connect(ctx context.Context) (Conn, error)
```

Creates a new database connection.

The context allows the connection attempt to respect cancellation and deadlines.

For example:

```go
ctx, cancel := context.WithTimeout(
    context.Background(),
    5*time.Second,
)
defer cancel()
```

A driver can use this context while establishing the connection.

### `Driver`

```go
Driver() Driver
```

Returns the underlying `Driver`.

This allows `database/sql` to retrieve the driver associated with the connector.

---

# D. `Conn`

This represents a single database connection.

Conceptually:

```go
type Conn interface {
    Prepare(query string) (Stmt, error)
    Close() error
    Begin() (Tx, error)
}
```

### `Prepare`

```go
Prepare(query string) (Stmt, error)
```

Prepares an SQL statement.

For example:

```sql
SELECT id, name FROM users WHERE id = ?
```

The returned `Stmt` can then execute the prepared statement.

### `Close`

```go
Close() error
```

Closes the connection.

The driver should release resources associated with that connection.

### `Begin`

```go
Begin() (Tx, error)
```

Starts a database transaction.

For example:

```text
BEGIN
   │
   ├── INSERT
   ├── UPDATE
   └── COMMIT
```

---

# E. `ConnPrepareContext`

```go
type ConnPrepareContext interface {
    PrepareContext(
        ctx context.Context,
        query string,
    ) (Stmt, error)
}
```

### `PrepareContext`

This is the context-aware version of `Prepare`.

It allows preparation to respect:

- deadlines
- cancellation
- request lifetime

Modern drivers should generally support context-aware operations.

---

# F. `ConnBeginTx`

```go
type ConnBeginTx interface {
    BeginTx(
        ctx context.Context,
        opts TxOptions,
    ) (Tx, error)
}
```

### `BeginTx`

Starts a transaction with:

- a `context.Context`
- transaction options

The options can include things such as:

```go
TxOptions{
    Isolation: ...,
    ReadOnly: true,
}
```

This provides more control than the older `Begin()` method.

---

# G. `Stmt`

A `Stmt` represents a prepared SQL statement.

Conceptually:

```go
type Stmt interface {
    Close() error
    NumInput() int
    Exec(args []Value) (Result, error)
    Query(args []Value) (Rows, error)
}
```

### `Close`

```go
Close() error
```

Releases resources associated with the prepared statement.

### `NumInput`

```go
NumInput() int
```

Reports how many placeholder parameters the statement expects.

For example:

```sql
SELECT * FROM users WHERE id = ? AND age = ?
```

would conceptually have:

```text
NumInput() = 2
```

A driver can return:

```go
-1
```

when it cannot determine the number of parameters.

### `Exec`

```go
Exec(args []Value) (Result, error)
```

Executes a prepared statement without returning rows.

Typical examples:

```text
INSERT
UPDATE
DELETE
```

### `Query`

```go
Query(args []Value) (Rows, error)
```

Executes a prepared statement and returns rows.

Typical example:

```sql
SELECT ...
```

---

# H. Context-Aware Statement Interfaces

Modern drivers can additionally implement:

```text
StmtExecContext
StmtQueryContext
```

### `ExecContext`

Conceptually:

```go
ExecContext(
    ctx context.Context,
    args []NamedValue,
) (Result, error)
```

Executes a prepared statement with context support.

### `QueryContext`

Conceptually:

```go
QueryContext(
    ctx context.Context,
    args []NamedValue,
) (Rows, error)
```

Executes a prepared statement and returns rows while respecting the context.

This is important for server applications where a request might be canceled before the database operation finishes.

---

# I. `ExecerContext`

```go
type ExecerContext interface {
    ExecContext(
        ctx context.Context,
        query string,
        args []NamedValue,
    ) (Result, error)
}
```

This allows a connection to execute SQL directly without requiring `database/sql` to prepare a statement first.

For example:

```text
database/sql
      │
      ├── ExecContext()
      │
      ▼
driver.Conn
      │
      ▼
database
```

The older `Execer` interface is deprecated; modern drivers should generally implement `ExecerContext` instead.

---

# J. `QueryerContext`

```go
type QueryerContext interface {
    QueryContext(
        ctx context.Context,
        query string,
        args []NamedValue,
    ) (Rows, error)
}
```

Allows a driver to execute a query directly with context support.

For example:

```sql
SELECT id, name FROM users
```

and return a `Rows` implementation.

---

# K. `Tx`

A transaction represents an active database transaction.

It provides:

```go
type Tx interface {
    Commit() error
    Rollback() error
}
```

### `Commit`

```go
Commit() error
```

Makes the transaction's changes permanent.

### `Rollback`

```go
Rollback() error
```

Cancels the transaction's changes.

Conceptually:

```text
Begin
  │
  ├── operation
  ├── operation
  │
  └── Commit
       OR
      Rollback
```

---

# L. `Rows`

`Rows` represents query results.

Its core methods are:

```go
type Rows interface {
    Columns() []string
    Close() error
    Next(dest []Value) error
}
```

### `Columns`

```go
Columns() []string
```

Returns the names of the columns.

For:

```sql
SELECT id, name FROM users
```

you might receive:

```go
[]string{"id", "name"}
```

### `Close`

```go
Close() error
```

Releases resources associated with the result set.

### `Next`

```go
Next(dest []Value) error
```

Retrieves the next row.

The driver places the values into `dest`.

When there are no more rows, it returns:

```go
io.EOF
```

This is one of the most important methods in a database driver because it connects the database's result-streaming mechanism to `database/sql`.

---

# M. `RowsNextResultSet`

Some databases support multiple result sets from one query.

For example:

```text
Query
 │
 ├── Result set 1
 │
 └── Result set 2
```

A driver can implement:

```go
NextResultSet() error
```

to move to the next result set.

This is useful for databases or stored procedures that return multiple sets of results.

---

# N. `RowsColumnType...` Interfaces

These optional interfaces provide metadata about columns.

There are several:

## `RowsColumnTypeDatabaseTypeName`

Provides:

```go
DatabaseTypeName() string
```

For example:

```text
VARCHAR
INTEGER
TIMESTAMP
```

## `RowsColumnTypeLength`

Provides:

```go
ColumnTypeLength() (length int64, ok bool)
```

Provides the maximum length of a column when the database exposes that information.

## `RowsColumnTypeNullable`

Provides:

```go
ColumnTypeNullable() (
    nullable bool,
    ok bool,
)
```

Reports whether a column can contain `NULL`.

## `RowsColumnTypePrecisionScale`

Provides:

```go
ColumnTypePrecisionScale() (
    precision int64,
    scale int64,
    ok bool,
)
```

Useful for numeric types such as:

```sql
DECIMAL(10, 2)
```

## `RowsColumnTypeScanType`

Provides:

```go
ColumnTypeScanType() reflect.Type
```

Reports the Go type that the driver recommends for scanning the database value.

These interfaces allow `database/sql` to expose column metadata to applications.

---

# O. `Result`

After an `INSERT`, `UPDATE`, or `DELETE`, a driver returns a `Result`.

It provides:

```go
type Result interface {
    LastInsertId() (int64, error)
    RowsAffected() (int64, error)
}
```

### `LastInsertId`

```go
LastInsertId() (int64, error)
```

Returns the ID generated by an insert when the database/driver supports this concept.

For example:

```text
INSERT user
     │
     ▼
ID = 42
```

The driver could return:

```go
42
```

Not every database supports this mechanism.

### `RowsAffected`

```go
RowsAffected() (int64, error)
```

Returns how many rows were affected.

For example:

```sql
UPDATE users SET active = false
WHERE last_login < ...
```

could return:

```text
Rows affected: 127
```

---

# P. `RowsAffected`

The standard package provides a simple implementation:

```go
type RowsAffected int64
```

It implements:

```go
LastInsertId()
RowsAffected()
```

This is convenient for drivers.

For example, a driver can conceptually do:

```go
return driver.RowsAffected(10), nil
```

meaning:

```text
10 rows were affected
```

---

# Q. `Pinger`

```go
type Pinger interface {
    Ping(ctx context.Context) error
}
```

### `Ping`

Tests whether a connection is alive and usable.

This can be used by `database/sql` for health checking.

For example:

```text
Application
    │
    ▼
database/sql
    │
    ▼
Ping()
    │
    ▼
Database
```

---

# R. `Validator`

```go
type Validator interface {
    IsValid() bool
}
```

### `IsValid`

Determines whether a connection is still valid.

This is particularly useful with connection pooling.

If:

```go
IsValid() == false
```

the connection should no longer be reused.

---

# S. `SessionResetter`

```go
type SessionResetter interface {
    ResetSession(ctx context.Context) error
}
```

### `ResetSession`

Resets connection/session state before the connection is reused.

This matters because database sessions can contain state such as:

```text
SET commands
temporary state
session variables
transaction state
```

A connection returned to a pool should not accidentally leak inappropriate state into the next user's operation.

---

# T. `NamedValue`

```go
type NamedValue struct {
    Name    string
    Ordinal int
    Value   Value
}
```

Represents a query argument.

For example:

```go
sql.Named("userID", 42)
```

can eventually become a `NamedValue` containing information such as:

```text
Name:    userID
Ordinal: 1
Value:   42
```

This is especially important for drivers supporting named parameters or special argument handling.

---

# U. `NamedValueChecker`

```go
type NamedValueChecker interface {
    CheckNamedValue(*NamedValue) error
}
```

### `CheckNamedValue`

This lets a driver inspect and potentially transform query arguments before they reach the database.

This is particularly useful for:

- custom Go types
- database-specific types
- special driver options
- named parameters

A driver can return `ErrSkip` to allow normal conversion to continue.

It can also return `ErrRemoveArgument` to remove an argument from the final argument list.

This is an advanced but powerful driver feature.

---

# V. `Value`

```go
type Value any
```

`Value` represents a value that can travel between `database/sql` and the driver.

Typical values include:

```text
int64
float64
bool
[]byte
string
time.Time
nil
```

The exact conversion rules are handled by the driver/sql machinery.

---

# W. `Valuer`

```go
type Valuer interface {
    Value() (Value, error)
}
```

### `Value`

Converts a custom Go value into a database-driver value.

Imagine a custom type:

```go
type Email string
```

A type implementing `Valuer` can control how its value is sent to the database.

Conceptually:

```text
Email("user@example.com")
          │
          ▼
       Value()
          │
          ▼
     driver.Value
```

---

# X. `ValueConverter`

```go
type ValueConverter interface {
    ConvertValue(any) (Value, error)
}
```

### `ConvertValue`

Converts an arbitrary Go value into a valid driver value.

This is useful when a driver needs specialized conversion behavior.

---

# Y. `NotNull`

`NotNull` wraps another converter and rejects `nil`.

Conceptually:

```text
NotNull
   │
   └── Converter
```

Its method:

```go
ConvertValue(v any) (Value, error)
```

will reject `nil`, while delegating other conversions to the underlying converter.

---

# Z. `Null`

`Null` does the opposite.

It permits `nil` and delegates other values to its converter.

It also provides:

```go
ConvertValue(v any) (Value, error)
```

These two helpers are useful when a driver wants explicit control over nullable and non-nullable values.

---

# AA. `IsolationLevel`

Represents a transaction isolation level.

Examples include:

```text
LevelDefault
LevelReadUncommitted
LevelReadCommitted
LevelRepeatableRead
LevelSnapshot
LevelSerializable
LevelLinearizable
```

The actual levels supported depend on the database.

A driver can reject isolation levels that its database does not support.

---

# AB. `TxOptions`

Contains transaction configuration.

Conceptually:

```go
type TxOptions struct {
    Isolation IsolationLevel
    ReadOnly  bool
}
```

For example:

```text
Isolation = Serializable
ReadOnly  = true
```

A driver uses this information when implementing `BeginTx`.

---

# AC. `ColumnConverter`

This is an optional interface associated with prepared statements.

Its purpose is to allow a statement to specify how particular parameters should be converted.

It is an older/advanced mechanism; modern drivers often use `NamedValueChecker` for more sophisticated argument handling.

---

# AD. `DriverContext` + `Connector`: Why They Matter

Consider an application connecting to a database 1,000 times.

A basic driver might repeatedly do:

```text
DSN
 │
 ├── parse username
 ├── parse password
 ├── parse host
 ├── parse port
 └── parse options
```

With `Connector`, the configuration can instead be parsed once:

```text
                 DSN
                  │
                  ▼
           OpenConnector()
                  │
                  ▼
              Connector
             /    |    \
            /     |     \
         Conn   Conn   Conn
```

This is one of the major reasons the newer interfaces exist.

---

# 4. Simple Example: Implementing a Tiny Driver

Because `database/sql/driver` is a **driver implementation package**, the example should demonstrate an actual driver rather than merely calling `database/sql`.

Here is a deliberately simplified driver:

```go
package main

import (
    "context"
    "database/sql"
    "database/sql/driver"
    "fmt"
    "io"
)

type MyDriver struct{}

type MyConnector struct{}

type MyConn struct{}

type MyStmt struct{}

type MyRows struct {
    done bool
}

func (d MyDriver) Open(name string) (driver.Conn, error) {
    return &MyConn{}, nil
}

func (d MyDriver) OpenConnector(name string) (driver.Connector, error) {
    return MyConnector{}, nil
}

func (c MyConnector) Connect(ctx context.Context) (driver.Conn, error) {
    return &MyConn{}, nil
}

func (c MyConnector) Driver() driver.Driver {
    return MyDriver{}
}

func (c *MyConn) Prepare(query string) (driver.Stmt, error) {
    return &MyStmt{}, nil
}

func (c *MyConn) Close() error {
    return nil
}

func (c *MyConn) Begin() (driver.Tx, error) {
    return nil, fmt.Errorf("transactions not implemented")
}

func (s *MyStmt) Close() error {
    return nil
}

func (s *MyStmt) NumInput() int {
    return 0
}

func (s *MyStmt) Exec(args []driver.Value) (driver.Result, error) {
    return driver.RowsAffected(1), nil
}

func (s *MyStmt) Query(args []driver.Value) (driver.Rows, error) {
    return &MyRows{}, nil
}

func (r *MyRows) Columns() []string {
    return []string{"message"}
}

func (r *MyRows) Close() error {
    return nil
}

func (r *MyRows) Next(dest []driver.Value) error {
    if r.done {
        return io.EOF
    }

    dest[0] = "Hello from driver"
    r.done = true

    return nil
}

func main() {
    sql.Register("mydriver", MyDriver{})

    db, err := sql.Open("mydriver", "")
    if err != nil {
        panic(err)
    }
    defer db.Close()

    var message string

    err = db.QueryRow("SELECT message").Scan(&message)
    if err != nil {
        panic(err)
    }

    fmt.Println(message)
}
```

The important thing to notice is that the application still uses:

```text
database/sql
```

rather than directly manipulating the driver.

The architecture is:

```text
Application
     │
     ▼
database/sql
     │
     ▼
database/sql/driver interfaces
     │
     ▼
MyDriver
```

Real production drivers are substantially more complicated because they must handle network communication, authentication, transactions, prepared statements, errors, cancellation, pooling semantics, types, and database-specific behavior.

---

# 5. Three Common Beginner Mistakes

## Mistake 1: Using `database/sql/driver` for Normal Application Development

A beginner might write:

```go
import "database/sql/driver"
```

and attempt to perform application queries directly through it.

### Why this is a mistake

`driver` is the **low-level implementation layer**.

For normal applications, use:

```go
import "database/sql"
```

and a third-party database driver.

The official documentation explains that most code should use `database/sql`.

### Remember

```text
Application developer
        ↓
database/sql

Driver developer
        ↓
database/sql/driver
```

---

## Mistake 2: Thinking `Driver.Open` Means "Create a Connection Pool"

A beginner might assume:

```go
func (d Driver) Open(name string) (driver.Conn, error)
```

means the driver should maintain its own connection pool.

### Why this is wrong

`database/sql` already manages the connection pool.

The driver should generally create the individual connection requested by the SQL layer.

### Better mental model

```text
database/sql
    │
    ├── connection pool
    │
    ├── connection
    ├── connection
    └── connection
             │
             ▼
          driver
```

The Go documentation notes that `database/sql` maintains the pool of idle connections.

---

## Mistake 3: Implementing Only the Old Interfaces

A beginner studying older tutorials may encounter:

```go
Execer
Queryer
```

and assume these are the preferred interfaces for modern drivers.

They are not.

Modern drivers should favor context-aware interfaces such as:

```text
ExecerContext
QueryerContext
ConnPrepareContext
ConnBeginTx
```

and should support modern connection lifecycle interfaces where appropriate.

### Better approach

When writing a new driver, start from the **current Go documentation**, not an old driver tutorial.

---

# 6. Two Real-World Applications

## Application 1: Building a Database Driver

Imagine a company develops a proprietary SQL-compatible database:

```text
CompanyDB
```

Go applications should ideally be able to write:

```go
db, err := sql.Open("companydb", dsn)
```

and then use:

```go
db.Query(...)
db.Exec(...)
db.Begin(...)
```

To make that possible, the company can implement the interfaces in:

```text
database/sql/driver
```

This is the most direct real-world use case.

Many existing database drivers work this way. For example, Oracle's Go driver implements the `database/sql/driver` interfaces.

---

## Application 2: Supporting Specialized Database Behavior

A sophisticated driver may need to support things such as:

```text
custom data types
named parameters
context cancellation
connection validation
session reset
transaction isolation
multiple result sets
database-specific metadata
```

The optional interfaces in `database/sql/driver` allow a driver to expose those capabilities to `database/sql`.

For example:

```text
Application
     │
     ▼
database/sql
     │
     ▼
NamedValueChecker
     │
     ▼
Custom database type
```

This allows the driver to remain compatible with the standard Go SQL API while still supporting database-specific capabilities.

---

# 7. Three Progressively Challenging Exercises

## Exercise 1 — Beginner: Build a Minimal Driver

Create a custom driver called `memorydriver` that implements the minimum interfaces necessary to work with `database/sql`.

Requirements:

- Implement `Driver`.
- Implement `Conn`.
- Implement `Stmt`.
- Implement `Rows`.
- Register the driver using `sql.Register`.
- Make a simple `SELECT` query return one row containing a string.
- Use `database/sql` from the application layer to retrieve and print that value.
- Do not use an actual external database.

**Goal:** Understand how `database/sql` communicates with the low-level driver interfaces.

---

## Exercise 2 — Intermediate: Add Transactions and Context Support

Extend your driver from Exercise 1.

Add:

- Transaction support using `Tx`.
- `Commit`.
- `Rollback`.
- `ConnBeginTx`.
- `ConnPrepareContext`.
- `ExecerContext`.
- `QueryerContext`.
- Context cancellation behavior.
- Appropriate handling of invalid or closed connections.

Create a test program that:

1. Begins a transaction.
2. Performs multiple operations.
3. Commits one transaction.
4. Rolls back another transaction.
5. Uses a canceled context for a database operation.

**Goal:** Understand how modern context-aware database drivers interact with `database/sql`.

---

## Exercise 3 — Advanced: Build a Functional In-Memory SQL Driver

Build a more complete in-memory database driver.

Your driver should support:

```text
CREATE
INSERT
SELECT
UPDATE
DELETE
```

and should maintain data in memory.

Add support for:

- `Connector`
- `DriverContext`
- Connection pooling compatibility
- `Pinger`
- `Validator`
- `SessionResetter`
- Prepared statements
- Transactions
- Context cancellation
- Named parameters
- `NamedValueChecker`
- Multiple result sets
- Column metadata
- A custom Go type implementing `driver.Valuer`
- Appropriate driver errors
- Detection of invalid connections

Your goal is to make it possible for an application to interact with your database entirely through:

```go
database/sql
```

without knowing anything about your internal implementation.

**Goal:** Understand how a production-quality database driver fits into Go's `database/sql` architecture.

---

# 8. One Important Takeaway

The most important thing to understand about `database/sql/driver` is:

> **It is not primarily a database client API; it is a contract for implementing database clients that `database/sql` can use.**

The separation is intentional:

```text
                YOUR APPLICATION
                       │
                       ▼
               ┌──────────────┐
               │ database/sql │
               └──────┬───────┘
                      │
             standard driver API
                      │
                      ▼
             ┌─────────────────┐
             │ database/sql/   │
             │     driver      │
             └───────┬─────────┘
                     │
        ┌────────────┼────────────┐
        ▼            ▼            ▼
     MySQL       PostgreSQL     SQLite
     driver        driver        driver
        │            │            │
        ▼            ▼            ▼
    Database      Database     Database
```

This separation is one of the fundamental design goals of Go's SQL architecture.

---

# 9. Thought-Provoking Question

Imagine you are designing a new database that supports a feature unavailable in MySQL, PostgreSQL, or SQLite—such as a special query parameter type or a unique transaction model.

**How would you design your `database/sql/driver` implementation so that applications can use your database-specific capabilities without sacrificing compatibility with the standard `database/sql` API?**

---

## Official References

- Go Package Documentation: `database/sql/driver`
  https://pkg.go.dev/database/sql/driver
- Go Source Code: `database/sql/driver`
  https://go.dev/src/database/sql/driver/
- Go SQL Drivers Guide
  https://go.dev/wiki/SQLDrivers
- Go `database/sql` documentation
  https://pkg.go.dev/database/sql
