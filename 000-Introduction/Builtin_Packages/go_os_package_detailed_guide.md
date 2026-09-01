# Go `os` Package — Detailed Guide

## 1. What is the `os` package?

The `os` package is one of Go's standard-library packages.

```go
import "os"
```

It provides an interface for interacting with the **operating system**.

You can use it to:

- Create, open, read, write, and delete files
- Create and remove directories
- Rename and move files
- Read file information
- Change file permissions
- Work with environment variables
- Get the current working directory
- Get the hostname
- Get process IDs
- Start, signal, and terminate processes
- Read command-line arguments
- Work with temporary directories/files
- Work with symbolic and hard links
- Access platform-independent OS functionality

The important idea is:

> **`os` is the bridge between your Go program and the operating system.**

The package is designed to provide a relatively uniform interface across operating systems such as Windows, Linux, and macOS.

---

## 2. When is the `os` package commonly used?

You will frequently use `os` when building:

### File-based applications

```text
Application
    ↓
os package
    ↓
Operating System
    ↓
File System
```

For example:

```go
data, err := os.ReadFile("config.txt")
```

### Command-line applications

```go
args := os.Args
```

### Configuration-driven applications

```go
port := os.Getenv("PORT")
```

### File-management tools

```go
os.Rename("old.txt", "new.txt")
```

### System utilities

```go
pid := os.Getpid()
```

### Server applications

For example, reading environment variables:

```go
databaseURL := os.Getenv("DATABASE_URL")
```

---

## 3. Simple example

Let's start with a small example that creates, writes, reads, and deletes a file.

```go
package main

import (
	"fmt"
	"os"
)

func main() {

	// Create a file
	err := os.WriteFile(
		"example.txt",
		[]byte("Hello from Go!"),
		0644,
	)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("File created successfully")

	// Read the file
	data, err := os.ReadFile("example.txt")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("File content:", string(data))

	// Delete the file
	err = os.Remove("example.txt")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("File deleted successfully")
}
```

Here we used three important functions:

```go
os.WriteFile()
os.ReadFile()
os.Remove()
```

---

# 4. Important `os` package categories

It helps to learn the package in groups rather than memorizing functions randomly.

| Category | Examples |
|---|---|
| Files | `Open`, `Create`, `ReadFile`, `WriteFile` |
| Directories | `Mkdir`, `MkdirAll`, `ReadDir` |
| File management | `Remove`, `Rename`, `Truncate` |
| File metadata | `Stat`, `Lstat` |
| Permissions | `Chmod`, `Chown` |
| Environment | `Getenv`, `Setenv`, `LookupEnv` |
| OS information | `Hostname`, `Getpid`, `Getuid` |
| Working directory | `Getwd`, `Chdir` |
| Temporary files | `CreateTemp`, `MkdirTemp`, `TempDir` |
| Links | `Link`, `Symlink`, `Readlink` |
| Processes | `StartProcess`, `FindProcess` |
| Process control | `Kill`, `Signal`, `Wait` |
| Errors | `IsExist`, `IsNotExist`, `IsPermission` |
| Rooted filesystem access | `OpenRoot`, `Root` |

---

# 5. File Functions

## `os.Create()`

Creates a file.

```go
file, err := os.Create("hello.txt")
```

If the file doesn't exist, it is created.

If it already exists, its contents are **truncated**.

Example:

```go
file, err := os.Create("hello.txt")

if err != nil {
	fmt.Println(err)
	return
}

defer file.Close()

file.WriteString("Hello Go")
```

### Important

Don't use `Create()` when you want to preserve existing contents.

---

## `os.Open()`

Opens an existing file for reading.

```go
file, err := os.Open("hello.txt")
```

Example:

```go
file, err := os.Open("hello.txt")

if err != nil {
	fmt.Println(err)
	return
}

defer file.Close()

data := make([]byte, 100)

n, err := file.Read(data)

if err != nil {
	fmt.Println(err)
	return
}

fmt.Println(string(data[:n]))
```

`Open()` is primarily used when you need a `*os.File` and want to perform operations such as `Read`, `Seek`, `Stat`, etc.

---

## `os.OpenFile()`

This is one of the most important functions in the package.

```go
file, err := os.OpenFile(
	"hello.txt",
	os.O_CREATE|os.O_WRONLY|os.O_APPEND,
	0644,
)
```

The three arguments are:

```text
filename
flags
permissions
```

### Common flags

```go
os.O_RDONLY
os.O_WRONLY
os.O_RDWR
os.O_CREATE
os.O_APPEND
os.O_TRUNC
os.O_EXCL
os.O_SYNC
```

For example:

```go
os.O_CREATE | os.O_WRONLY | os.O_APPEND
```

means:

```text
Create if necessary
+
Write only
+
Append to existing content
```

This is extremely useful for log files.

```go
file, err := os.OpenFile(
	"app.log",
	os.O_CREATE|os.O_WRONLY|os.O_APPEND,
	0644,
)

if err != nil {
	fmt.Println(err)
	return
}

defer file.Close()

file.WriteString("Application started\n")
```

---

## `os.ReadFile()`

Reads an entire file into memory.

```go
data, err := os.ReadFile("hello.txt")
```

Example:

```go
data, err := os.ReadFile("hello.txt")

if err != nil {
	fmt.Println(err)
	return
}

fmt.Println(string(data))
```

### Important

`ReadFile()` is convenient for small/medium files.

For a very large file, reading the entire file into memory may be inappropriate. In that situation, use `os.Open()` with streaming reads.

---

## `os.WriteFile()`

Writes data to a file.

```go
err := os.WriteFile(
	"hello.txt",
	[]byte("Hello World"),
	0644,
)
```

If the file exists, its previous contents are replaced.

Example:

```go
data := []byte("Hello from Go")

err := os.WriteFile("hello.txt", data, 0644)

if err != nil {
	fmt.Println(err)
	return
}
```

---

## `os.Remove()`

Deletes a file or an empty directory.

```go
err := os.Remove("hello.txt")
```

Example:

```go
if err := os.Remove("hello.txt"); err != nil {
	fmt.Println(err)
}
```

---

## `os.RemoveAll()`

Removes a path and everything underneath it.

```go
err := os.RemoveAll("temp")
```

Suppose:

```text
temp/
 ├── a.txt
 ├── b.txt
 └── logs/
     └── app.log
```

Then:

```go
os.RemoveAll("temp")
```

removes the entire tree.

### Be careful

This is powerful and potentially destructive.

Never blindly call:

```go
os.RemoveAll(userInput)
```

without validating what the user supplied.

---

## `os.Rename()`

Renames or moves a file.

```go
err := os.Rename("old.txt", "new.txt")
```

Example:

```go
err := os.Rename(
	"old-name.txt",
	"new-name.txt",
)

if err != nil {
	fmt.Println(err)
}
```

---

## `os.Truncate()`

Changes the size of a file.

```go
err := os.Truncate("data.txt", 100)
```

If the new size is smaller, data is removed.

If the new size is larger, the file is extended.

Example:

```go
err := os.Truncate("data.txt", 0)
```

This effectively empties the file while keeping the file itself.

---

# 6. Directory Functions

## `os.Mkdir()`

Creates one directory.

```go
err := os.Mkdir("documents", 0755)
```

It does **not** automatically create missing parent directories.

For example:

```text
project/
```

works:

```go
os.Mkdir("project", 0755)
```

But:

```go
os.Mkdir("a/b/c", 0755)
```

fails if `a` and `b` don't already exist.

---

## `os.MkdirAll()`

Creates a directory and all missing parent directories.

```go
err := os.MkdirAll(
	"a/b/c",
	0755,
)
```

It creates:

```text
a/
└── b/
    └── c/
```

This is generally more convenient when creating application directory structures.

---

## `os.ReadDir()`

Reads directory entries.

```go
entries, err := os.ReadDir(".")
```

Example:

```go
entries, err := os.ReadDir(".")

if err != nil {
	fmt.Println(err)
	return
}

for _, entry := range entries {
	fmt.Println(entry.Name())
}
```

You can determine whether an entry is a directory:

```go
if entry.IsDir() {
	fmt.Println("Directory:", entry.Name())
} else {
	fmt.Println("File:", entry.Name())
}
```

`ReadDir` returns directory entries sorted by filename.

---

# 7. File Information

## `os.Stat()`

Returns information about a file.

```go
info, err := os.Stat("hello.txt")
```

You can obtain:

```go
fmt.Println(info.Name())
fmt.Println(info.Size())
fmt.Println(info.Mode())
fmt.Println(info.ModTime())
fmt.Println(info.IsDir())
```

Example:

```go
info, err := os.Stat("hello.txt")

if err != nil {
	fmt.Println(err)
	return
}

fmt.Println("Name:", info.Name())
fmt.Println("Size:", info.Size())
fmt.Println("Directory:", info.IsDir())
fmt.Println("Mode:", info.Mode())
fmt.Println("Modified:", info.ModTime())
```

---

## `os.Lstat()`

`Lstat()` is similar to `Stat()`, but there is an important difference with symbolic links.

```go
info, err := os.Lstat("link.txt")
```

If the path is a symbolic link:

```text
link.txt → actual.txt
```

`Lstat()` describes the **link itself**.

`Stat()` normally follows the link and describes the target.

---

## `os.SameFile()`

Determines whether two `FileInfo` values describe the same file.

```go
same := os.SameFile(info1, info2)
```

Useful when dealing with:

- hard links
- aliases
- filesystem comparisons

---

# 8. Working Directory Functions

## `os.Getwd()`

Gets the current working directory.

```go
dir, err := os.Getwd()

if err != nil {
	fmt.Println(err)
	return
}

fmt.Println(dir)
```

---

## `os.Chdir()`

Changes the current working directory.

```go
err := os.Chdir("documents")
```

After this:

```go
os.Getwd()
```

will report the new working directory.

### Important

Changing the working directory affects the **current process**, so be careful in larger applications.

---

# 9. Environment Variables

Environment variables are extremely important in real-world Go applications.

For example:

```text
PORT=8080
DATABASE_URL=...
APP_ENV=production
```

---

## `os.Getenv()`

Gets an environment variable.

```go
port := os.Getenv("PORT")
```

Example:

```go
port := os.Getenv("PORT")

fmt.Println("Port:", port)
```

If the variable doesn't exist:

```go
os.Getenv("UNKNOWN")
```

returns:

```text
""
```

---

## `os.LookupEnv()`

Use `LookupEnv()` when you need to distinguish:

```text
variable doesn't exist
```

from:

```text
variable exists but is empty
```

Example:

```go
value, exists := os.LookupEnv("PORT")

if exists {
	fmt.Println("PORT exists:", value)
} else {
	fmt.Println("PORT does not exist")
}
```

---

## `os.Setenv()`

Sets an environment variable for the current process.

```go
err := os.Setenv("APP_ENV", "development")
```

Then:

```go
fmt.Println(os.Getenv("APP_ENV"))
```

prints:

```text
development
```

It modifies the environment of the current process, not a permanent system-wide configuration.

---

## `os.Unsetenv()`

Removes an environment variable from the current process.

```go
err := os.Unsetenv("APP_ENV")
```

---

## `os.Clearenv()`

Removes all environment variables from the current process environment.

```go
os.Clearenv()
```

This is powerful and should be used carefully.

---

## `os.Environ()`

Returns all environment variables.

```go
env := os.Environ()

for _, value := range env {
	fmt.Println(value)
}
```

The returned strings have the form:

```text
KEY=value
```

---

## `os.Expand()`

Replaces variables using a custom mapping function.

Example:

```go
result := os.Expand(
	"Hello $NAME",
	func(key string) string {
		return "Chandu"
	},
)

fmt.Println(result)
```

Output:

```text
Hello Chandu
```

---

## `os.ExpandEnv()`

Expands variables using the current environment.

```go
os.Setenv("NAME", "Chandu")

result := os.ExpandEnv("Hello $NAME")

fmt.Println(result)
```

Output:

```text
Hello Chandu
```

---

# 10. Temporary Files

## `os.CreateTemp()`

Creates a temporary file.

```go
file, err := os.CreateTemp("", "example-*.txt")
```

Example:

```go
file, err := os.CreateTemp("", "myapp-*.txt")

if err != nil {
	fmt.Println(err)
	return
}

defer os.Remove(file.Name())
defer file.Close()

file.WriteString("Temporary data")

fmt.Println("Temporary file:", file.Name())
```

The `*` is replaced with random characters.

Useful for:

- temporary uploads
- intermediate processing
- test files
- temporary generated data

---

## `os.MkdirTemp()`

Creates a temporary directory.

```go
dir, err := os.MkdirTemp("", "myapp-*")
```

Example:

```go
dir, err := os.MkdirTemp("", "processing-*")

if err != nil {
	fmt.Println(err)
	return
}

defer os.RemoveAll(dir)

fmt.Println("Temporary directory:", dir)
```

---

## `os.TempDir()`

Returns the default temporary directory.

```go
dir := os.TempDir()

fmt.Println(dir)
```

The actual result depends on the operating system and environment.

---

# 11. User Directories

## `os.UserHomeDir()`

Returns the current user's home directory.

```go
home, err := os.UserHomeDir()
```

Useful for application data such as:

```text
~/.myapp
```

---

## `os.UserConfigDir()`

Returns an appropriate location for user-specific configuration files.

```go
dir, err := os.UserConfigDir()
```

---

## `os.UserCacheDir()`

Returns an appropriate location for user-specific cached data.

```go
dir, err := os.UserCacheDir()
```

These functions help avoid hard-coding OS-specific paths.

---

# 12. Executable Information

## `os.Executable()`

Returns the path of the executable that started the current process.

```go
path, err := os.Executable()

if err != nil {
	fmt.Println(err)
	return
}

fmt.Println(path)
```

This can be useful when an application needs to locate resources relative to its executable.

---

# 13. Hostname

## `os.Hostname()`

Returns the machine's hostname.

```go
name, err := os.Hostname()

if err != nil {
	fmt.Println(err)
	return
}

fmt.Println("Hostname:", name)
```

Useful for:

- server identification
- diagnostics
- logging
- distributed systems

---

# 14. Process Information

## `os.Getpid()`

Gets the current process ID.

```go
pid := os.Getpid()

fmt.Println("PID:", pid)
```

---

## `os.Getppid()`

Gets the parent process ID.

```go
ppid := os.Getppid()

fmt.Println("Parent PID:", ppid)
```

---

# 15. User and Group IDs

These are particularly relevant on Unix-like systems.

### `os.Getuid()`

Gets user ID.

```go
uid := os.Getuid()
```

### `os.Geteuid()`

Gets effective user ID.

```go
euid := os.Geteuid()
```

### `os.Getgid()`

Gets group ID.

```go
gid := os.Getgid()
```

### `os.Getegid()`

Gets effective group ID.

```go
egid := os.Getegid()
```

### `os.Getgroups()`

Returns the groups to which the process belongs.

```go
groups, err := os.Getgroups()
```

These functions are OS-dependent. Some have special behavior or are unavailable on Windows.

---

# 16. `os.Getpagesize()`

Returns the system's memory page size.

```go
size := os.Getpagesize()

fmt.Println("Page size:", size)
```

This is generally more relevant to lower-level/system programming than ordinary application development.

---

# 17. File Permissions

## `os.Chmod()`

Changes file permissions.

```go
err := os.Chmod("script.sh", 0755)
```

On Unix-like systems:

```text
0755
```

generally means:

```text
Owner:  read + write + execute
Group:  read + execute
Others: read + execute
```

The exact behavior of permission bits varies by operating system.

---

## `os.Chown()`

Changes the owner and group of a file.

```go
err := os.Chown(
	"file.txt",
	1000,
	1000,
)
```

This is primarily relevant to Unix-like systems.

---

## `os.Lchown()`

Similar to `Chown()`, but handles symbolic links differently.

```go
err := os.Lchown(
	"link.txt",
	1000,
	1000,
)
```

---

# 18. File Timestamps

## `os.Chtimes()`

Changes access and modification times.

```go
err := os.Chtimes(
	"file.txt",
	atime,
	mtime,
)
```

Useful in:

- synchronization tools
- backup programs
- file restoration
- testing

---

# 19. Symbolic and Hard Links

## `os.Link()`

Creates a hard link.

```go
err := os.Link(
	"original.txt",
	"copy.txt",
)
```

A hard link is not simply a normal file copy.

---

## `os.Symlink()`

Creates a symbolic link.

```go
err := os.Symlink(
	"original.txt",
	"shortcut.txt",
)
```

Conceptually:

```text
shortcut.txt
      ↓
original.txt
```

---

## `os.Readlink()`

Reads where a symbolic link points.

```go
target, err := os.Readlink("shortcut.txt")
```

---

# 20. Error Checking Functions

## `os.IsExist()`

Historically used to determine whether an error indicates that something already exists.

```go
if os.IsExist(err) {
	fmt.Println("Already exists")
}
```

For modern Go code, prefer `errors.Is` with the appropriate sentinel where possible.

---

## `os.IsNotExist()`

Checks whether an error indicates that something doesn't exist.

```go
if os.IsNotExist(err) {
	fmt.Println("File does not exist")
}
```

Modern code often prefers:

```go
errors.Is(err, os.ErrNotExist)
```

---

## `os.IsPermission()`

Checks whether an error indicates a permission problem.

```go
if os.IsPermission(err) {
	fmt.Println("Permission denied")
}
```

---

## `os.IsTimeout()`

Checks whether an error indicates a timeout.

```go
if os.IsTimeout(err) {
	fmt.Println("Operation timed out")
}
```

---

## `os.IsPathSeparator()`

Checks whether a byte is an OS path separator.

```go
if os.IsPathSeparator('/') {
	fmt.Println("Path separator")
}
```

For normal path manipulation, however, prefer the `path/filepath` package.

---

# 21. `os.Exit()`

Terminates the current process immediately.

```go
os.Exit(1)
```

Conventionally:

```text
0     → success
non-0 → error
```

### Very important difference

If you do:

```go
defer fmt.Println("This will not execute")

os.Exit(1)
```

the deferred statement will not run.

---

# 22. `os.Pipe()`

Creates a pipe for communication between a reader and writer.

```go
reader, writer, err := os.Pipe()
```

Conceptually:

```text
Writer
   ↓
 PIPE
   ↓
Reader
```

It is useful for connecting streams of data between parts of a program or process-related operations.

---

# 23. `os.StartProcess()`

Starts a new operating-system process.

```go
process, err := os.StartProcess(
	name,
	argv,
	attr,
)
```

This is relatively low-level.

For most applications, you should normally use:

```go
os/exec
```

instead.

---

# 24. `os.FindProcess()`

Finds a process by process ID.

```go
process, err := os.FindProcess(pid)
```

This gives you a:

```go
*os.Process
```

which can then be used to:

```go
process.Kill()
process.Signal(...)
process.Wait()
```

---

# 25. `os.Process` Methods

Once you have:

```go
process *os.Process
```

you have several important methods.

## `Process.Kill()`

Terminates the process.

```go
err := process.Kill()
```

---

## `Process.Signal()`

Sends an OS signal.

```go
err := process.Signal(os.Interrupt)
```

Signals are particularly important in Unix/Linux applications.

---

## `Process.Release()`

Releases resources associated with the process.

```go
err := process.Release()
```

---

## `Process.Wait()`

Waits for the process to finish.

```go
state, err := process.Wait()
```

It returns:

```go
*os.ProcessState
```

---

## `Process.WithHandle()`

Provides access to the process handle through a callback.

```go
err := process.WithHandle(func(handle uintptr) {
	// use handle
})
```

This is a lower-level API and is generally unnecessary for beginners.

---

# 26. `os.ProcessState`

`ProcessState` describes the result/status of a completed process.

## `ExitCode()`

Returns the process exit code.

```go
code := state.ExitCode()
```

---

## `Exited()`

Determines whether the process has exited.

```go
if state.Exited() {
	fmt.Println("Process exited")
}
```

---

## `Pid()`

Gets the process ID.

```go
pid := state.Pid()
```

---

## `Success()`

Determines whether the process completed successfully.

```go
if state.Success() {
	fmt.Println("Success")
}
```

---

## `String()`

Returns a human-readable representation.

```go
fmt.Println(state.String())
```

---

## `SystemTime()`

Returns the amount of system CPU time used by the process.

```go
duration := state.SystemTime()
```

---

## `UserTime()`

Returns the amount of user CPU time used by the process.

```go
duration := state.UserTime()
```

---

## `Sys()`

Returns OS-specific process status information.

```go
value := state.Sys()
```

This is lower-level and OS-dependent.

---

## `SysUsage()`

Returns OS-specific resource-usage information.

```go
usage := state.SysUsage()
```

---

# 27. `os.File`

When you use:

```go
file, err := os.Open("hello.txt")
```

the `file` variable is:

```go
*os.File
```

`File` represents an open file.

It provides many methods.

---

## `File.Close()`

Closes the file.

```go
file.Close()
```

Usually:

```go
file, err := os.Open("hello.txt")

if err != nil {
	return
}

defer file.Close()
```

This is one of the most important patterns to learn.

---

## `File.Read()`

Reads bytes from a file.

```go
buffer := make([]byte, 100)

n, err := file.Read(buffer)
```

`n` tells you how many bytes were actually read.

Use:

```go
buffer[:n]
```

rather than assuming the entire buffer was filled.

---

## `File.ReadAt()`

Reads from a specific position without relying on the file's current offset.

```go
n, err := file.ReadAt(buffer, 100)
```

Meaning:

```text
Read into buffer
starting at byte offset 100
```

Useful for random-access file processing.

---

## `File.Write()`

Writes bytes.

```go
n, err := file.Write([]byte("Hello"))
```

---

## `File.WriteAt()`

Writes at a specific file offset.

```go
n, err := file.WriteAt(
	[]byte("Hello"),
	100,
)
```

---

## `File.WriteString()`

Writes a string directly.

```go
n, err := file.WriteString("Hello Go")
```

---

## `File.Seek()`

Moves the file's current position.

```go
position, err := file.Seek(
	0,
	os.SEEK_SET,
)
```

Modern Go code commonly uses:

```go
io.SeekStart
io.SeekCurrent
io.SeekEnd
```

instead of the deprecated `os.SEEK_*` constants.

---

## `File.Stat()`

Gets information about the open file.

```go
info, err := file.Stat()
```

This returns `FileInfo`.

---

## `File.Name()`

Returns the name associated with the file.

```go
fmt.Println(file.Name())
```

---

## `File.Fd()`

Returns the underlying file descriptor/handle.

```go
fd := file.Fd()
```

This is a low-level operation and should generally not be needed for ordinary Go programs.

---

## `File.Sync()`

Synchronizes file contents with the underlying storage as appropriate.

```go
err := file.Sync()
```

This can matter when durability is important, such as certain database or logging scenarios.

---

## `File.Truncate()`

Changes the size of the currently opened file.

```go
err := file.Truncate(100)
```

---

## `File.Chmod()`

Changes permissions on the opened file.

```go
err := file.Chmod(0644)
```

---

## `File.Chown()`

Changes owner/group.

```go
err := file.Chown(uid, gid)
```

Again, this is primarily relevant to Unix-like systems.

---

## `File.Chdir()`

Changes the process's current working directory to the directory represented by the file.

```go
err := file.Chdir()
```

The file must represent a directory.

---

## `File.ReadDir()`

Reads directory entries from an opened directory.

```go
entries, err := file.ReadDir(-1)
```

This is related to:

```go
os.ReadDir()
```

but works on an already-opened directory.

---

## `File.Readdir()`

Reads directory entries and returns `FileInfo` values.

```go
entries, err := file.Readdir(-1)
```

It is an older-style directory API; `ReadDir()` is generally preferable for newer code.

---

## `File.Readdirnames()`

Returns directory entry names.

```go
names, err := file.Readdirnames(-1)
```

---

## `File.ReadFrom()`

Reads data from an `io.Reader` and writes it into the file.

```go
n, err := file.ReadFrom(reader)
```

Useful for copying streams.

---

## `File.WriteTo()`

Writes file contents to an `io.Writer`.

```go
n, err := file.WriteTo(writer)
```

This can be useful for efficiently copying file data to another destination.

---

## `File.SetDeadline()`

Sets both read and write deadlines.

```go
err := file.SetDeadline(time.Now().Add(time.Second))
```

---

## `File.SetReadDeadline()`

Sets a read deadline.

```go
err := file.SetReadDeadline(deadline)
```

---

## `File.SetWriteDeadline()`

Sets a write deadline.

```go
err := file.SetWriteDeadline(deadline)
```

These are more relevant to certain OS-backed file descriptors and devices than ordinary disk-file programming.

---

## `File.SyscallConn()`

Provides access to the underlying system-level connection.

```go
conn, err := file.SyscallConn()
```

This is an advanced API intended for interaction with lower-level OS facilities.

For normal Go application development, avoid it unless you specifically need system-level access.

---

# 28. `FileInfo`

`FileInfo` describes a file.

Common methods include:

```go
info.Name()
info.Size()
info.Mode()
info.ModTime()
info.IsDir()
info.Sys()
```

Example:

```go
info, err := os.Stat("hello.txt")

if err != nil {
	return
}

fmt.Println(info.Name())
fmt.Println(info.Size())
fmt.Println(info.Mode())
fmt.Println(info.ModTime())
fmt.Println(info.IsDir())
```

---

# 29. `FileMode`

`FileMode` represents file permissions and file type information.

Example:

```go
mode := os.FileMode(0644)
```

Common permission values:

```text
0644
0755
0700
```

You can inspect modes:

```go
info.Mode().IsRegular()
info.Mode().IsDir()
```

---

# 30. `DirEntry`

A `DirEntry` represents a directory entry.

For example:

```go
entries, _ := os.ReadDir(".")

for _, entry := range entries {

	fmt.Println(entry.Name())

	if entry.IsDir() {
		fmt.Println("Directory")
	}
}
```

Useful methods include:

```go
entry.Name()
entry.IsDir()
entry.Type()
entry.Info()
```

---

# 31. `Root` — Modern Rooted Filesystem Access

Modern versions of Go also provide `os.Root`.

It allows operations to be constrained to a particular directory tree.

You can create one with:

```go
root, err := os.OpenRoot("sandbox")
```

Then:

```go
file, err := root.Open("hello.txt")
```

The important idea is:

```text
sandbox/
   ├── hello.txt
   ├── data/
   └── config/
```

Operations through `root` are intended to stay within that root.

This is particularly interesting for applications that need to safely work with paths supplied by another component.

---

# 32. Important `Root` methods

The major ones are:

```text
root.Open()
root.OpenFile()
root.Create()
root.ReadFile()
root.WriteFile()
root.Mkdir()
root.MkdirAll()
root.Remove()
root.RemoveAll()
root.Rename()
root.Stat()
root.Lstat()
root.Readlink()
root.Symlink()
root.Link()
root.Chmod()
root.Chown()
root.Lchown()
root.Chtimes()
root.FS()
root.Name()
root.Close()
root.OpenRoot()
```

Their purpose largely mirrors the corresponding package-level filesystem operations, but the path is interpreted relative to the root.

For example:

```go
root, err := os.OpenRoot("sandbox")
if err != nil {
	return
}

defer root.Close()

err = root.WriteFile(
	"data.txt",
	[]byte("Hello"),
	0644,
)
```

This writes:

```text
sandbox/data.txt
```

---

# 33. `os.LinkError`

Some filesystem operations can return a `*os.LinkError`.

It contains information about:

```text
operation
path
underlying error
```

Its methods include:

```go
Error()
Unwrap()
```

This allows normal Go error inspection.

For example:

```go
if err != nil {
	var linkErr *os.LinkError

	if errors.As(err, &linkErr) {
		fmt.Println("Operation:", linkErr.Op)
	}
}
```

---

# 34. `os.SyscallError`

`SyscallError` records an error associated with a system call.

It provides:

```text
Error()
Unwrap()
Timeout()
```

This is mainly useful when dealing with lower-level OS errors.

---

# 35. `os.Signal`

`Signal` represents an operating-system signal.

Two especially important portable signals are:

```go
os.Interrupt
os.Kill
```

For example:

```go
process.Signal(os.Interrupt)
```

The exact available signals depend on the operating system.

---

# 36. Practical `os` Example

Here is a more realistic example combining several functions.

```go
package main

import (
	"fmt"
	"os"
)

func main() {

	// Get current working directory
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Getwd error:", err)
		return
	}

	fmt.Println("Working directory:", dir)

	// Create directory
	err = os.MkdirAll("app/data", 0755)
	if err != nil {
		fmt.Println("Mkdir error:", err)
		return
	}

	// Write file
	err = os.WriteFile(
		"app/data/config.txt",
		[]byte("application=golang\n"),
		0644,
	)

	if err != nil {
		fmt.Println("Write error:", err)
		return
	}

	// Read file
	data, err := os.ReadFile("app/data/config.txt")

	if err != nil {
		fmt.Println("Read error:", err)
		return
	}

	fmt.Println("Content:")
	fmt.Println(string(data))

	// File information
	info, err := os.Stat("app/data/config.txt")

	if err != nil {
		fmt.Println("Stat error:", err)
		return
	}

	fmt.Println("File:", info.Name())
	fmt.Println("Size:", info.Size())
	fmt.Println("Mode:", info.Mode())
	fmt.Println("Modified:", info.ModTime())

	// Environment variable
	err = os.Setenv("APP_ENV", "development")

	if err != nil {
		fmt.Println("Setenv error:", err)
		return
	}

	fmt.Println(
		"Environment:",
		os.Getenv("APP_ENV"),
	)
}
```

This demonstrates:

```text
Getwd
MkdirAll
WriteFile
ReadFile
Stat
Setenv
Getenv
```

---

# 37. Three Common Beginner Mistakes

## Mistake 1: Forgetting to close files

Bad:

```go
file, _ := os.Open("data.txt")

// use file
```

Better:

```go
file, err := os.Open("data.txt")

if err != nil {
	return
}

defer file.Close()
```

### Why?

An open file consumes an OS resource.

---

## Mistake 2: Ignoring errors

Bad:

```go
data, _ := os.ReadFile("config.txt")
```

This hides potentially important errors.

Better:

```go
data, err := os.ReadFile("config.txt")

if err != nil {
	fmt.Println("Error:", err)
	return
}
```

Go's `os` functions commonly return errors rather than OS-specific error numbers.

---

## Mistake 3: Confusing `Create()` and `Open()`

Beginners sometimes assume:

```go
os.Create("data.txt")
```

simply opens an existing file.

It doesn't.

`Create()` creates the file if needed and truncates an existing file.

If you want to read an existing file:

```go
os.Open("data.txt")
```

If you want controlled read/write behavior:

```go
os.OpenFile(...)
```

---

# 38. Two Real-World Applications

## Application 1: Configuration management

A production application might read configuration from environment variables:

```go
databaseURL := os.Getenv("DATABASE_URL")
port := os.Getenv("PORT")
environment := os.Getenv("APP_ENV")
```

For example:

```text
DATABASE_URL=postgres://...
PORT=8080
APP_ENV=production
```

This allows the same binary to run in:

```text
Development
Testing
Staging
Production
```

without hard-coding configuration.

---

## Application 2: File-processing service

Imagine an image-processing service:

```text
uploads/
   ↓
Go application
   ↓
os.ReadFile / os.Open
   ↓
Process image
   ↓
os.WriteFile
   ↓
processed/
```

The `os` package can handle the underlying file operations while other packages perform the actual image processing.

---

# 39. Three Progressive Exercises

As requested, **no solutions** are provided.

## Exercise 1 — Beginner: File Manager

Create a Go program that:

1. Creates a directory named `practice`.
2. Creates a file named `notes.txt` inside it.
3. Writes three lines of text into the file.
4. Reads the file.
5. Prints the contents.
6. Prints the file size.
7. Deletes the file.

### Functions to practice

```text
Mkdir
WriteFile
ReadFile
Stat
Remove
```

---

## Exercise 2 — Intermediate: Log Manager

Create a command-line log-management program.

The program should:

1. Create a `logs` directory if it doesn't exist.
2. Create `application.log`.
3. Append new log messages instead of replacing existing content.
4. Read the log file.
5. Display the number of bytes in the log.
6. Display the last modification time.
7. Allow the user to clear the log.

### Functions to practice

```text
MkdirAll
OpenFile
Write
ReadFile
Stat
Truncate
```

Do not use `os.WriteFile()` for the append operation.

---

## Exercise 3 — Advanced: Secure File Workspace

Build a small file workspace application.

The application should:

1. Create a dedicated workspace directory.
2. Open that directory as an `os.Root`.
3. Allow files to be created inside the workspace.
4. Allow files to be read.
5. Allow directories to be created.
6. Display file information.
7. Rename files.
8. Delete files.
9. Reject attempts to access files outside the workspace.
10. Clean up the workspace when the program finishes.

### Functions/types to investigate

```text
os.OpenRoot
os.Root
Root.Open
Root.OpenFile
Root.ReadFile
Root.WriteFile
Root.MkdirAll
Root.Stat
Root.Rename
Root.Remove
Root.RemoveAll
Root.Close
```

This exercise will give you a much deeper understanding of why filesystem boundaries matter in applications.

---

# 40. The Most Important `os` Functions to Learn First

Although the package contains a large API, don't try to memorize everything at once.

## Level 1 — Essential

```text
os.Open
os.Create
os.ReadFile
os.WriteFile
os.Remove
os.Stat
os.Mkdir
os.MkdirAll
```

## Level 2 — Practical

```text
os.OpenFile
os.Rename
os.ReadDir
os.Getwd
os.Chdir
os.TempDir
os.CreateTemp
os.MkdirTemp
```

## Level 3 — Configuration

```text
os.Getenv
os.LookupEnv
os.Setenv
os.Unsetenv
os.Environ
os.ExpandEnv
```

## Level 4 — System information

```text
os.Hostname
os.Getpid
os.Getppid
os.Executable
os.Getuid
os.Getgid
```

## Level 5 — Advanced

```text
os.StartProcess
os.FindProcess
os.Pipe
os.Root
os.OpenRoot
os.SyscallConn
```

---

# 41. `os` vs Related Go Packages

A very important distinction for beginners:

| Requirement | Package |
|---|---|
| OS/files/processes | `os` |
| Path manipulation | `path/filepath` |
| External commands | `os/exec` |
| Buffered I/O | `bufio` |
| General I/O | `io` |
| File-system abstraction | `io/fs` |
| File formats | Depends on format |

For example, don't try to manually concatenate paths:

```go
path := "data/" + filename
```

Prefer:

```go
filepath.Join("data", filename)
```

---

# 42. Big Picture

Think about the `os` package like this:

```text
                     Go Application
                           |
                           |
                    +------v------+
                    |     os     |
                    +------+------+
                           |
          +----------------+----------------+
          |                |                |
          v                v                v
       Files          Environment       Processes
          |                |                |
     +----+----+       +---+---+       +----+----+
     |         |       |       |       |         |
   Read      Write   Getenv  Setenv  Start     Kill
     |         |       |       |       |         |
     +----+----+       +---+---+       +----+----+
          |
          v
      File System
```

Once you understand this model, the large `os` API becomes much easier to organize mentally.

---

# 43. Thought-Provoking Question

Suppose you are building a **multi-user file-upload server**, and users are allowed to specify filenames such as:

```text
report.pdf
images/photo.jpg
../../secret.txt
```

**How would you design your Go program so that users can freely manage files inside their own directory while preventing them from accessing files belonging to other users or the operating system?**

Think specifically about how `os`, `filepath`, symbolic links, and the newer `os.Root` API could work together—and where simply "checking the path first" might still create a security problem.
