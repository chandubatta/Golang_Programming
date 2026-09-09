# Go `database/sql` Package

The Go `database/sql` package provides a **generic, database-independent interface for working with SQL databases**. It handles connection pooling, executing SQL statements, querying rows, transactions, prepared statements, scanning database values into Go variables, and more.

Importantly, `database/sql` **does not contain a database driver itself**. You use it together with a driver for PostgreSQL, MySQL, SQLite, SQL Server, etc.

## 1. What is `database/sql`?

Think of `database/sql` as a **common language between your Go application and different SQL databases**.

```text
Go application
      │
      ▼
database/sql
      │
      ▼
Database driver
      │
      ▼
PostgreSQL / MySQL / SQLite / SQL Server
```

Your Go code primarily interacts with `database/sql`, while the driver handles database-specific communication.

### What does it provide?

The package gives you APIs for:

- Opening databases
- Managing connection pools
- Executing `INSERT`, `UPDATE`, and `DELETE`
- Executing `SELECT` queries
- Reading multiple rows
- Reading a single row
- Prepared statements
- Transactions
- Transaction isolation levels
- Context cancellation and timeouts
- Handling SQL `NULL`
- Inspecting column metadata
- Database connection statistics
- Custom scanning through the `Scanner` interface

A particularly important concept is that `*sql.DB` is **not simply one database connection**. It represents a database handle and manages a pool of underlying connections. It is safe for concurrent use by multiple goroutines.

### When is it commonly used?

You will commonly use `database/sql` when building:

- REST APIs
- Web applications
- Backend services
- Authentication systems
- E-commerce applications
- Financial applications
- Inventory systems
- Reporting systems
- CLI applications that persist data
- Microservices that use relational databases

---

# 2. Simple Example

For a realistic example, suppose we have a `users` table:

```sql
CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    name TEXT,
    email TEXT
);
```

A simple Go program could look like this:

```go
package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	// Import the appropriate database driver here.
	// Example:
	// _ "github.com/example/database-driver"
)

func main() {
	ctx := context.Background()

	db, err := sql.Open("your-driver", "your-database-connection-string")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Verify that the database is reachable.
	if err := db.PingContext(ctx); err != nil {
		log.Fatal(err)
	}

	// Insert a user.
	result, err := db.ExecContext(
		ctx,
		"INSERT INTO users (name, email) VALUES (?, ?)",
		"Alice",
		"alice@example.com",
	)
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Rows inserted:", rowsAffected)

	// Retrieve one user.
	var name string
	var email string

	err = db.QueryRowContext(
		ctx,
		"SELECT name, email FROM users WHERE id = ?",
		1,
	).Scan(&name, &email)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("User not found")
			return
		}

		log.Fatal(err)
	}

	fmt.Println("Name:", name)
	fmt.Println("Email:", email)
}
```

The exact driver name, connection string, and placeholder syntax depend on the database driver.

Notice the three important operations:

```go
db.ExecContext(...)
```

for statements that don't return rows,

```go
db.QueryContext(...)
```

for multiple rows, and

```go
db.QueryRowContext(...)
```

for a query expected to return at most one row.

---

# 3. The Most Important Concepts

Before learning every function, understand these six types:

| Type | Purpose |
|---|---|
| `sql.DB` | Database handle + connection pool |
| `sql.Conn` | One dedicated database connection |
| `sql.Rows` | Multiple query results |
| `sql.Row` | One query result |
| `sql.Stmt` | Prepared SQL statement |
| `sql.Tx` | Database transaction |

The typical flow is:

```text
sql.Open()
    ↓
*sql.DB
    ↓
Exec / Query / QueryRow
    ↓
Result / Rows / Row
```

For transactions:

```text
db.BeginTx()
    ↓
*sql.Tx
    ↓
Exec / Query / QueryRow
    ↓
Commit() or Rollback()
```

---

# 4. Every Exported Function in `database/sql`

The package-level API currently contains three functions: `ConvertAssign`, `Drivers`, and `Register`.

## `sql.ConvertAssign`

```go
sql.ConvertAssign(scanCtx, dest, src)
```

Introduced in Go 1.27.

It copies a database driver value into a destination, performing the conversions used by `Rows.Scan`.

It is primarily intended for **database driver implementations**, rather than normal application code.

As a beginner, you normally won't call this directly.

---

## `sql.Drivers`

```go
sql.Drivers()
```

Returns the names of all registered database drivers.

Example:

```go
fmt.Println(sql.Drivers())
```

Possible output might look like:

```text
[mysql postgres sqlite]
```

The actual result depends on which drivers your program has registered.

This can be useful for debugging driver registration.

---

## `sql.Register`

```go
sql.Register(name, driver)
```

Registers a database driver under a name.

Normally, **you don't implement or call this yourself**. Database drivers commonly register themselves when imported.

For example:

```go
import (
	"database/sql"
	_ "some/database/driver"
)
```

The blank import executes the driver's initialization code, which can register the driver.

Registering the same name twice, or registering a nil driver, causes a panic.

---

# 5. `sql.DB`

`DB` is arguably the most important type in the package.

It represents a database handle and manages a pool of connections. It is safe for concurrent use.

## `sql.Open`

```go
db, err := sql.Open(driverName, dataSourceName)
```

Creates a `*sql.DB`.

A crucial beginner misconception is that `Open()` necessarily establishes a connection immediately.

It may only validate configuration. Use `Ping()` or `PingContext()` when you need to verify connectivity.

---

## `sql.OpenDB`

```go
db := sql.OpenDB(connector)
```

Creates a `*sql.DB` using a `driver.Connector`.

This is useful when the driver exposes a connector instead of requiring a traditional string-based DSN.

Most beginners will use `sql.Open` or a driver-specific helper instead.

---

## `db.Begin`

```go
tx, err := db.Begin()
```

Starts a transaction.

It uses `context.Background()` internally, so modern applications often prefer:

```go
db.BeginTx(ctx, opts)
```

---

## `db.BeginTx`

```go
tx, err := db.BeginTx(ctx, opts)
```

Starts a transaction with:

- a `context.Context`
- optional `TxOptions`

Example:

```go
tx, err := db.BeginTx(ctx, &sql.TxOptions{
	Isolation: sql.LevelSerializable,
	ReadOnly:  false,
})
```

If the context is canceled, `database/sql` rolls back the transaction.

---

## `db.Close`

```go
err := db.Close()
```

Closes the database handle and prevents new queries from starting.

However, a `*sql.DB` is normally a **long-lived object**, so you generally create one during application startup and share it throughout the application.

You don't normally open and close a database connection for every request.

---

## `db.Conn`

```go
conn, err := db.Conn(ctx)
```

Obtains a single dedicated connection from the pool.

Normally:

```go
db.Query(...)
```

is preferable.

Use `Conn` when you specifically need operations to occur on the **same underlying database session/connection**.

Always return it:

```go
defer conn.Close()
```

---

## `db.Driver`

```go
driver := db.Driver()
```

Returns the underlying database driver.

This is mainly useful for advanced driver-level functionality and diagnostics.

---

## `db.Exec`

```go
result, err := db.Exec(query, args...)
```

Executes a query that does not return rows.

Typical uses:

```sql
INSERT
UPDATE
DELETE
CREATE TABLE
ALTER TABLE
```

Example:

```go
result, err := db.Exec(
	"UPDATE users SET name = ? WHERE id = ?",
	"Bob",
	10,
)
```

`Exec` uses `context.Background()` internally.

For applications where cancellation or timeouts matter, prefer `ExecContext`.

---

## `db.ExecContext`

```go
result, err := db.ExecContext(ctx, query, args...)
```

Same basic purpose as `Exec`, but accepts a context.

This is generally preferable in server applications:

```go
ctx, cancel := context.WithTimeout(
	context.Background(),
	3*time.Second,
)
defer cancel()

_, err := db.ExecContext(
	ctx,
	"UPDATE users SET active = ? WHERE id = ?",
	true,
	10,
)
```

If the request is canceled or the deadline expires, the context can allow the database operation to stop, subject to driver support.

---

## `db.Ping`

```go
err := db.Ping()
```

Checks whether the database is accessible.

Useful for:

- startup checks
- health checks
- diagnostics

It uses `context.Background()`.

---

## `db.PingContext`

```go
err := db.PingContext(ctx)
```

Context-aware version of `Ping`.

For example:

```go
ctx, cancel := context.WithTimeout(
	context.Background(),
	2*time.Second,
)
defer cancel()

if err := db.PingContext(ctx); err != nil {
	log.Fatal(err)
}
```

---

## `db.Prepare`

```go
stmt, err := db.Prepare(query)
```

Creates a prepared statement.

Example:

```go
stmt, err := db.Prepare(
	"INSERT INTO users (name) VALUES (?)",
)
if err != nil {
	log.Fatal(err)
}
defer stmt.Close()
```

You can then reuse it:

```go
stmt.Exec("Alice")
stmt.Exec("Bob")
stmt.Exec("Charlie")
```

Prepared statements are useful when executing the same SQL statement repeatedly.

---

## `db.PrepareContext`

```go
stmt, err := db.PrepareContext(ctx, query)
```

Context-aware version of `Prepare`.

Important detail: the context applies to **preparing the statement**, not to subsequent executions of the resulting statement.

---

## `db.Query`

```go
rows, err := db.Query(query, args...)
```

Executes a query that returns rows.

Typical use:

```sql
SELECT id, name FROM users
```

Then:

```go
defer rows.Close()

for rows.Next() {
	// Scan each row.
}
```

For context-aware code, prefer `QueryContext`.

---

## `db.QueryContext`

```go
rows, err := db.QueryContext(ctx, query, args...)
```

Executes a query returning multiple rows while accepting a context.

Example:

```go
rows, err := db.QueryContext(
	ctx,
	"SELECT id, name FROM users WHERE active = ?",
	true,
)
if err != nil {
	log.Fatal(err)
}
defer rows.Close()

for rows.Next() {
	var id int
	var name string

	if err := rows.Scan(&id, &name); err != nil {
		log.Fatal(err)
	}

	fmt.Println(id, name)
}

if err := rows.Err(); err != nil {
	log.Fatal(err)
}
```

---

## `db.QueryRow`

```go
row := db.QueryRow(query, args...)
```

Used when you expect **at most one row**.

Example:

```go
var name string

err := db.QueryRow(
	"SELECT name FROM users WHERE id = ?",
	42,
).Scan(&name)
```

One important detail:

`QueryRow` doesn't return the query error immediately. The error is generally returned by `Scan`.

---

## `db.QueryRowContext`

```go
row := db.QueryRowContext(ctx, query, args...)
```

Context-aware version.

This is one of the most frequently used methods in backend Go programs.

Example:

```go
var name string

err := db.QueryRowContext(
	ctx,
	"SELECT name FROM users WHERE id = ?",
	42,
).Scan(&name)

if err == sql.ErrNoRows {
	fmt.Println("User doesn't exist")
}
```

`sql.ErrNoRows` is returned when no row exists.

---

# 6. Connection Pool Configuration

`database/sql` automatically maintains a connection pool.

## `SetConnMaxIdleTime`

```go
db.SetConnMaxIdleTime(duration)
```

Controls how long a connection can remain idle before it becomes eligible for closure.

Example:

```go
db.SetConnMaxIdleTime(5 * time.Minute)
```

---

## `SetConnMaxLifetime`

```go
db.SetConnMaxLifetime(duration)
```

Controls the maximum amount of time a connection can be reused.

Example:

```go
db.SetConnMaxLifetime(30 * time.Minute)
```

This can be useful with databases or infrastructure that periodically terminate old connections.

---

## `SetMaxIdleConns`

```go
db.SetMaxIdleConns(n)
```

Controls the maximum number of idle connections retained in the pool.

Example:

```go
db.SetMaxIdleConns(10)
```

---

## `SetMaxOpenConns`

```go
db.SetMaxOpenConns(n)
```

Controls the maximum number of simultaneously open database connections.

Example:

```go
db.SetMaxOpenConns(20)
```

This is particularly important in production because an unlimited or poorly chosen pool can overwhelm the database.

---

## `db.Stats`

```go
stats := db.Stats()
```

Returns connection-pool statistics.

For example:

```go
stats := db.Stats()

fmt.Println("Open:", stats.OpenConnections)
fmt.Println("In use:", stats.InUse)
fmt.Println("Idle:", stats.Idle)
```

This is useful for monitoring and diagnosing connection-pool problems.

---

# 7. `sql.Conn`

`Conn` represents **one database connection**, rather than the entire connection pool.

Its methods are:

### `BeginTx`

```go
conn.BeginTx(ctx, opts)
```

Starts a transaction using that particular connection.

### `Close`

```go
conn.Close()
```

Returns the connection to the pool.

### `ExecContext`

```go
conn.ExecContext(ctx, query, args...)
```

Executes a statement using that particular connection.

### `PingContext`

```go
conn.PingContext(ctx)
```

Checks that the connection is alive.

### `PrepareContext`

```go
conn.PrepareContext(ctx, query)
```

Prepares a statement on that connection.

### `QueryContext`

```go
conn.QueryContext(ctx, query, args...)
```

Executes a multi-row query on that connection.

### `QueryRowContext`

```go
conn.QueryRowContext(ctx, query, args...)
```

Executes a single-row query on that connection.

### `Raw`

```go
conn.Raw(func(driverConn any) error {
	// driver-specific operations
	return nil
})
```

Exposes the underlying driver connection temporarily.

This is an **advanced API**. The underlying driver connection must not be used after the callback returns.

---

# 8. `sql.Rows`

`Rows` represents multiple query results.

The normal pattern is:

```go
rows, err := db.QueryContext(ctx, query)
if err != nil {
	return err
}
defer rows.Close()

for rows.Next() {
	// Scan row
}

if err := rows.Err(); err != nil {
	return err
}
```

## `rows.Next`

```go
rows.Next()
```

Moves the cursor to the next row.

Returns:

```text
true  → another row exists
false → no more rows
```

---

## `rows.Scan`

```go
rows.Scan(&id, &name)
```

Copies the current row's columns into Go variables.

The number and compatible types of destinations must match the selected columns.

---

## `rows.Close`

```go
rows.Close()
```

Closes the result set.

Use:

```go
defer rows.Close()
```

in ordinary query code.

---

## `rows.Err`

```go
rows.Err()
```

Returns an error encountered during iteration.

This is important because an error can occur **after `Query` succeeds**, while rows are being read.

That's why this is good practice:

```go
for rows.Next() {
	// ...
}

if err := rows.Err(); err != nil {
	return err
}
```

---

## `rows.Columns`

```go
columns, err := rows.Columns()
```

Returns the column names.

For example:

```text
[id name email]
```

This can be useful when implementing generic database tools.

---

## `rows.ColumnTypes`

```go
types, err := rows.ColumnTypes()
```

Returns metadata about the result columns.

This can be useful for dynamic SQL tools, database explorers, CSV exporters, and generic query utilities.

---

## `rows.NextResultSet`

```go
rows.NextResultSet()
```

Moves to the next result set when a database/driver supports multiple result sets.

This is more advanced and is especially relevant when calling stored procedures or executing database-specific batches.

---

# 9. `sql.Row`

`Row` represents the result of a query expected to produce at most one row.

## `row.Scan`

```go
err := row.Scan(&id, &name)
```

Copies the selected columns into Go variables.

If there is no row:

```go
err == sql.ErrNoRows
```

---

## `row.Err`

```go
err := row.Err()
```

Returns any error associated with the row.

In many common situations you'll primarily encounter errors through `Scan`, but `Err` exists for explicitly checking the stored query error.

---

# 10. `sql.Stmt`

`Stmt` represents a prepared SQL statement.

It provides:

```go
stmt.Exec(...)
stmt.ExecContext(...)
stmt.Query(...)
stmt.QueryContext(...)
stmt.QueryRow(...)
stmt.QueryRowContext(...)
stmt.Close()
```

### `stmt.Exec`

Executes the prepared statement.

### `stmt.ExecContext`

Context-aware execution.

### `stmt.Query`

Executes the prepared statement and returns multiple rows.

### `stmt.QueryContext`

Context-aware multi-row query.

### `stmt.QueryRow`

Executes the prepared statement and returns a `Row`.

### `stmt.QueryRowContext`

Context-aware single-row query.

### `stmt.Close`

Releases resources associated with the statement.

For example:

```go
stmt, err := db.Prepare(
	"SELECT name FROM users WHERE id = ?",
)
if err != nil {
	return err
}
defer stmt.Close()
```

A prepared statement created from a `DB` can be used concurrently and remains usable for the lifetime of the database handle.

---

# 11. `sql.Tx`

`Tx` represents a database transaction.

A transaction lets multiple database operations behave as one logical unit.

For example:

```text
BEGIN
  ↓
UPDATE account A
  ↓
UPDATE account B
  ↓
COMMIT
```

If something fails:

```text
BEGIN
  ↓
UPDATE account A
  ↓
ERROR
  ↓
ROLLBACK
```

## `tx.Commit`

```go
err := tx.Commit()
```

Makes the transaction's changes permanent.

---

## `tx.Rollback`

```go
err := tx.Rollback()
```

Cancels the transaction's changes.

A common pattern is:

```go
tx, err := db.BeginTx(ctx, nil)
if err != nil {
	return err
}

defer tx.Rollback()

// operations...

return tx.Commit()
```

If `Commit()` succeeds, the deferred rollback is harmless.

---

## `tx.Exec`

```go
tx.Exec(query, args...)
```

Executes a non-row-returning statement inside the transaction.

---

## `tx.ExecContext`

```go
tx.ExecContext(ctx, query, args...)
```

Context-aware version.

---

## `tx.Prepare`

```go
tx.Prepare(query)
```

Creates a prepared statement associated with the transaction.

---

## `tx.PrepareContext`

```go
tx.PrepareContext(ctx, query)
```

Context-aware preparation.

---

## `tx.Query`

```go
tx.Query(query, args...)
```

Executes a multi-row query inside the transaction.

---

## `tx.QueryContext`

```go
tx.QueryContext(ctx, query, args...)
```

Context-aware version.

---

## `tx.QueryRow`

```go
tx.QueryRow(query, args...)
```

Executes a single-row query inside the transaction.

---

## `tx.QueryRowContext`

```go
tx.QueryRowContext(ctx, query, args...)
```

Context-aware version.

---

## `tx.Stmt`

```go
tx.Stmt(stmt)
```

Uses an existing prepared statement within the transaction.

---

## `tx.StmtContext`

```go
tx.StmtContext(ctx, stmt)
```

Context-aware version.

---

# 12. Transactions in Practice

Suppose you are transferring ₹500 between two bank accounts.

You don't want this:

```text
Account A -₹500
Account B +₹500  ← application crashes here
```

because money would disappear.

Instead:

```text
BEGIN TRANSACTION

Account A -₹500

Account B +₹500

COMMIT
```

If the second operation fails:

```text
BEGIN TRANSACTION

Account A -₹500

Account B +₹500 ← ERROR

ROLLBACK
```

The transaction ensures that the operations are treated as one unit.

---

# 13. `sql.TxOptions`

`TxOptions` controls transaction configuration.

Example:

```go
opts := &sql.TxOptions{
	Isolation: sql.LevelSerializable,
	ReadOnly:  false,
}

tx, err := db.BeginTx(ctx, opts)
```

Important isolation levels include:

```go
sql.LevelDefault
sql.LevelReadUncommitted
sql.LevelReadCommitted
sql.LevelWriteCommitted
sql.LevelRepeatableRead
sql.LevelSnapshot
sql.LevelSerializable
sql.LevelLinearizable
```

Not every database supports every isolation level, so the database driver determines what is actually available.

---

# 14. `sql.Result`

Methods that modify data usually return:

```go
sql.Result
```

For example:

```go
result, err := db.ExecContext(...)
```

The `Result` interface provides two important operations.

## `LastInsertId`

```go
id, err := result.LastInsertId()
```

Attempts to obtain an ID generated by the database.

However, **not every database/driver supports this operation**.

---

## `RowsAffected`

```go
count, err := result.RowsAffected()
```

Returns the number of rows affected.

Example:

```go
result, err := db.ExecContext(
	ctx,
	"DELETE FROM users WHERE id = ?",
	10,
)

count, err := result.RowsAffected()

fmt.Println("Deleted:", count)
```

---

# 15. Handling SQL `NULL`

SQL has a special value:

```sql
NULL
```

It is not the same thing as:

```go
""
```

or:

```go
0
```

or:

```go
false
```

Go's `database/sql` provides nullable types for this situation.

Examples include:

```go
sql.NullString
sql.NullInt64
sql.NullFloat64
sql.NullBool
sql.NullTime
```

And the generic:

```go
sql.Null[T]
```

is available in modern Go.

For example:

```go
var name sql.NullString

err := row.Scan(&name)

if name.Valid {
	fmt.Println(name.String)
} else {
	fmt.Println("name is NULL")
}
```

The `Scan` method reads the database value, while `Value` converts the Go nullable value back into a driver value.

---

# 16. `sql.Null[T]`

The generic nullable type is useful when you want a Go value that explicitly tracks whether a database value is valid.

Conceptually:

```go
sql.Null[T]
```

contains:

```text
Value
Valid
```

For example:

```text
Value = "Alice"
Valid = true
```

versus:

```text
Value = ""
Valid = false
```

This distinction is extremely useful when mapping nullable database columns to Go structures.

---

# 17. `sql.Named`

```go
sql.Named("id", 42)
```

Creates a named argument.

Example:

```go
db.QueryRowContext(
	ctx,
	"SELECT name FROM users WHERE id = :id",
	sql.Named("id", 42),
)
```

Named arguments can make SQL calls clearer, particularly for queries with many parameters.

Support for named parameters ultimately depends on the database driver.

---

# 18. `sql.NamedArg`

`NamedArg` represents a named SQL argument.

For example:

```go
arg := sql.NamedArg{
	Name:  "id",
	Value: 42,
}
```

Normally, however, you would simply write:

```go
sql.Named("id", 42)
```

---

# 19. `sql.Scanner`

`Scanner` is an interface that lets your own Go type control how a database value is converted into it.

Conceptually:

```go
type Scanner interface {
	Scan(src any) error
}
```

This is powerful for custom database types.

For example, you could create a custom type:

```go
type UserStatus string
```

and implement `Scan` so database values are converted into `UserStatus`.

This is useful for:

- custom enums
- JSON-like database values
- custom IDs
- domain-specific types
- database-specific representations

---

# 20. `sql.RawBytes`

```go
sql.RawBytes
```

Represents raw database bytes without copying them in the normal way.

Example:

```go
var data sql.RawBytes

rows.Scan(&data)
```

**Important:** the data is only valid until the next `Next`, `Scan`, or `Close` operation on the rows.

Therefore, beginners should generally prefer `[]byte` unless they specifically need the behavior of `RawBytes`.

---

# 21. `sql.Out`

`sql.Out` is used for output parameters, particularly with databases/drivers that support stored procedures with output parameters.

Conceptually:

```go
var result string

_, err := db.ExecContext(
	ctx,
	"some_procedure",
	sql.Named("output", sql.Out{Dest: &result}),
)
```

Not all database drivers support output parameters.

---

# 22. `sql.ColumnType`

`ColumnType` provides metadata about a result column.

Its methods include:

### `Name`

```go
column.Name()
```

Returns the column name or alias.

### `DatabaseTypeName`

```go
column.DatabaseTypeName()
```

Returns the database's type name, such as:

```text
VARCHAR
TEXT
INT
BIGINT
DECIMAL
```

### `Length`

```go
column.Length()
```

Returns the declared length where supported.

### `Nullable`

```go
column.Nullable()
```

Reports whether the column can contain `NULL`, when the driver supports that information.

### `DecimalSize`

```go
column.DecimalSize()
```

Returns decimal precision and scale when supported.

### `ScanType`

```go
column.ScanType()
```

Returns a Go `reflect.Type` suitable for scanning the column.

These methods are particularly useful when writing **generic database utilities** rather than ordinary business applications.

---

# 23. `sql.IsolationLevel`

`IsolationLevel` represents transaction isolation levels.

Example:

```go
sql.LevelSerializable
```

You use it through:

```go
sql.TxOptions{
	Isolation: sql.LevelSerializable,
}
```

### `String`

```go
level.String()
```

Returns a textual representation of the isolation level.

---

# 24. Important Error Values

Three package-level errors are particularly important.

## `sql.ErrNoRows`

```go
sql.ErrNoRows
```

Means a query expected a row but didn't find one.

Common pattern:

```go
if err == sql.ErrNoRows {
	// User doesn't exist.
}
```

Prefer `errors.Is` when wrapping may be involved:

```go
if errors.Is(err, sql.ErrNoRows) {
	// Not found.
}
```

---

## `sql.ErrTxDone`

```go
sql.ErrTxDone
```

Indicates that you attempted to use a transaction after it was already committed or rolled back.

---

## `sql.ErrConnDone`

```go
sql.ErrConnDone
```

Indicates that an operation was attempted on a connection that has already been returned to the pool/closed.

---

# 25. `database/sql` Function Cheat Sheet

| Function/method | Main purpose |
|---|---|
| `sql.Open` | Create database handle |
| `sql.OpenDB` | Create DB from connector |
| `sql.Drivers` | List registered drivers |
| `sql.Register` | Register a driver |
| `sql.ConvertAssign` | Driver-level value conversion |
| `db.Exec` | Execute non-row query |
| `db.ExecContext` | Context-aware execution |
| `db.Query` | Query multiple rows |
| `db.QueryContext` | Context-aware multi-row query |
| `db.QueryRow` | Query one row |
| `db.QueryRowContext` | Context-aware one-row query |
| `db.Prepare` | Prepare statement |
| `db.PrepareContext` | Context-aware preparation |
| `db.Begin` | Start transaction |
| `db.BeginTx` | Context-aware transaction |
| `db.Ping` | Test database connection |
| `db.PingContext` | Context-aware ping |
| `db.Conn` | Obtain dedicated connection |
| `db.Close` | Close DB handle |
| `db.Stats` | Get pool statistics |
| `db.SetMaxOpenConns` | Limit open connections |
| `db.SetMaxIdleConns` | Limit idle connections |
| `db.SetConnMaxLifetime` | Limit connection lifetime |
| `db.SetConnMaxIdleTime` | Limit idle connection age |

---

# 26. The `Context` Versions Are Important

You will frequently see pairs like:

```go
db.Exec()
db.ExecContext()
```

or:

```go
db.Query()
db.QueryContext()
```

The `Context` version lets your application control:

- deadlines
- cancellation
- request lifetime
- timeouts

For example, in an HTTP API:

```go
func getUser(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(
		r.Context(),
		2*time.Second,
	)
	defer cancel()

	var name string

	err := db.QueryRowContext(
		ctx,
		"SELECT name FROM users WHERE id = ?",
		42,
	).Scan(&name)

	// ...
}
```

This is much safer than allowing an accidental database query to run indefinitely.

---

# 27. Three Common Beginner Mistakes

## Mistake 1: Thinking `sql.DB` is one database connection

Beginners often think:

```go
db, _ := sql.Open(...)
```

means:

> "I opened one connection."

It actually creates a database handle that manages a **connection pool**.

### Avoid it

Create one long-lived `*sql.DB` and share it:

```go
var db *sql.DB
```

Don't repeatedly open and close it for every request.

---

## Mistake 2: Building SQL with string concatenation

Bad:

```go
query := "SELECT * FROM users WHERE name = '" + name + "'"
```

This can lead to SQL injection and escaping problems.

Prefer parameters:

```go
row := db.QueryRowContext(
	ctx,
	"SELECT * FROM users WHERE name = ?",
	name,
)
```

Let the driver handle parameter values.

---

## Mistake 3: Forgetting to close `Rows`

Bad:

```go
rows, err := db.QueryContext(ctx, query)

for rows.Next() {
	// ...
}
```

Better:

```go
rows, err := db.QueryContext(ctx, query)
if err != nil {
	return err
}
defer rows.Close()

for rows.Next() {
	// ...
}

if err := rows.Err(); err != nil {
	return err
}
```

Failure to properly consume/close results can contribute to connection-pool problems.

---

# 28. Two Real-World Applications

## Application 1: REST API

Imagine an e-commerce API:

```text
GET /users/42
       ↓
HTTP handler
       ↓
database/sql
       ↓
PostgreSQL/MySQL
       ↓
User record
       ↓
JSON response
```

`QueryRowContext` can retrieve a single user, while `QueryContext` can retrieve a collection of products.

Transactions can handle operations such as:

```text
Create order
    +
Decrease inventory
    +
Create payment record
```

as one logical operation.

---

## Application 2: Banking / Financial Systems

Suppose a money transfer requires:

```text
Debit Account A
Credit Account B
Create Transaction Record
```

These operations should normally happen within a transaction.

If the credit fails, you don't want the debit to remain committed.

`database/sql`'s `Tx` API provides the structure for:

```go
tx, err := db.BeginTx(ctx, opts)

...

tx.ExecContext(...)

...

tx.Commit()
```

or:

```go
tx.Rollback()
```

This makes transactions one of the most important parts of `database/sql`.

---

# 29. Three Progressively Challenging Exercises

## Exercise 1 — Basic CRUD

Create a Go program using `database/sql` that manages a `books` table.

Your program should allow the user to:

1. Insert a book.
2. Retrieve a book by ID.
3. Update a book's title.
4. Delete a book.
5. List all books.

Requirements:

- Use parameterized SQL queries.
- Use `ExecContext` for modifications.
- Use `QueryRowContext` for retrieving one book.
- Use `QueryContext` for retrieving multiple books.
- Properly close `Rows`.
- Handle `sql.ErrNoRows`.

**Do not use an ORM.**

---

## Exercise 2 — Transactions and Connection Pooling

Build a small inventory-management application.

Create tables representing:

```text
products
orders
order_items
```

When a customer places an order, your program must:

1. Create an order.
2. Insert its order items.
3. Decrease product inventory.
4. Reject the transaction if inventory is insufficient.
5. Roll back all changes when any operation fails.
6. Commit all changes when everything succeeds.

Additionally:

- Configure the database connection pool.
- Use `SetMaxOpenConns`.
- Use `SetMaxIdleConns`.
- Use `SetConnMaxLifetime`.
- Use context-aware database operations.
- Make sure transactions are properly committed or rolled back.

The important challenge is ensuring that **partial orders can never be committed**.

---

## Exercise 3 — Production-Style Database Layer

Build a small REST API for a user-management system.

Implement endpoints for:

```text
POST   /users
GET    /users/{id}
GET    /users
PUT    /users/{id}
DELETE /users/{id}
```

Requirements:

- Use `database/sql`.
- Use connection pooling.
- Use `context.Context`.
- Add query timeouts.
- Use prepared statements where appropriate.
- Correctly handle `sql.ErrNoRows`.
- Support nullable database fields.
- Use transactions where multiple related operations need atomicity.
- Expose database health through a health-check endpoint.
- Collect and expose useful `DBStats`.
- Properly handle database errors.
- Avoid SQL injection.
- Make the API safe for concurrent requests.

The goal is to design something resembling the database layer of a real Go backend rather than simply writing isolated SQL queries.

---

# 30. A Useful Mental Model

If you're learning `database/sql`, remember this:

```text
                    *sql.DB
                       │
              Connection Pool
             /       |       \
            /        |        \
        Conn        Conn      Conn
         │
         ├── Query
         ├── Exec
         └── Transaction
                 │
                 ▼
                Tx
          ┌──────┼──────┐
          │      │      │
        Exec   Query  QueryRow
          │      │      │
          ▼      ▼      ▼
       Result   Rows    Row
```

And the most important practical patterns are:

```go
// One row
db.QueryRowContext(...).Scan(...)

// Multiple rows
rows, _ := db.QueryContext(...)
defer rows.Close()

for rows.Next() {
	rows.Scan(...)
}

// Modification
db.ExecContext(...)

// Transaction
tx, _ := db.BeginTx(...)
defer tx.Rollback()

// operations...

tx.Commit()
```

Once these patterns become natural, most of everyday `database/sql` becomes much easier.

---

# 31. The Three Things I'd Learn First

If you're learning this package from scratch, don't try to memorize all of its APIs immediately.

Learn them in this order:

### Level 1 — Basic queries

```text
Open
Ping
ExecContext
QueryRowContext
QueryContext
Rows.Next
Rows.Scan
Rows.Err
Rows.Close
```

### Level 2 — Production usage

```text
Context
connection pooling
prepared statements
NULL handling
error handling
DB statistics
```

### Level 3 — Advanced database programming

```text
transactions
TxOptions
isolation levels
Conn
Stmt
ColumnType
Scanner
NamedArg
Raw
multiple result sets
driver-level APIs
```

This progression will give you practical competence without overwhelming you with the entire API at once.

---

# Thought-Provoking Question

Imagine you are building a high-traffic Go API where **10,000 requests per second** may query the database, but the database can safely handle only **200 concurrent connections**.

**How would you design your `database/sql` connection pool, request contexts, query timeouts, transactions, and error-handling strategy so that your Go application remains responsive instead of simply creating more and more database connections?**

That question gets to the heart of why understanding `database/sql` is about much more than knowing how to execute a `SELECT`.
