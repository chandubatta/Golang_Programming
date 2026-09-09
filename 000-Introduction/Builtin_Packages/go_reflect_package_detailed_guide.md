# Go `reflect` Package — Detailed Guide

The Go `reflect` package is one of the most powerful—and easiest to misuse—parts of the standard library. It lets a program inspect and manipulate types and values **at runtime**, even when the exact type was not known when the program was compiled.

## 1. What is the `reflect` package?

Import it with:

```go
import "reflect"
```

The `reflect` package provides **runtime reflection**.

Normally, Go is strongly and statically typed:

```go
name := "Chandu"
age := 25
```

The compiler knows that:

- `name` is a `string`
- `age` is an `int`

But sometimes a program receives a value without knowing its concrete type beforehand:

```go
func inspect(x any) {
    // What type is x?
    // What fields does it have?
    // Can we modify it?
    // Does it have a particular method?
}
```

Reflection allows you to answer those questions at runtime.

The two most important concepts are:

```go
reflect.Type
```

and:

```go
reflect.Value
```

A useful mental model:

```text
                 Reflection
                     │
          ┌──────────┴──────────┐
          ↓                     ↓
     reflect.Type          reflect.Value
          │                     │
     "What is it?"         "What data is it?"
          │                     │
       int, struct,         actual value,
       slice, map...        fields, elements...
```

`TypeOf` obtains the runtime type, while `ValueOf` obtains a reflection `Value`.

---

# 2. Simple example

Let's inspect a struct dynamically:

```go
package main

import (
    "fmt"
    "reflect"
)

type Person struct {
    Name string
    Age  int
}

func main() {
    p := Person{
        Name: "Chandu",
        Age: 25,
    }

    t := reflect.TypeOf(p)
    v := reflect.ValueOf(p)

    fmt.Println("Type:", t)
    fmt.Println("Kind:", t.Kind())

    fmt.Println("Name:", v.FieldByName("Name").String())
    fmt.Println("Age:", v.FieldByName("Age").Int())

    fmt.Println("Number of fields:", t.NumField())
}
```

Output:

```text
Type: main.Person
Kind: struct
Name: Chandu
Age: 25
Number of fields: 2
```

Here:

```go
t := reflect.TypeOf(p)
```

asks:

> "What type is `p`?"

while:

```go
v := reflect.ValueOf(p)
```

asks:

> "Give me a reflection object representing the actual value stored in `p`."

---

# 3. The most important concepts

Before learning every API, understand these four concepts.

## `Type`

`reflect.Type` describes a Go type.

```go
t := reflect.TypeOf(100)

fmt.Println(t.Name())
fmt.Println(t.Kind())
```

Possible output:

```text
int
int
```

`Type` is an interface describing type information. Its methods inspect structs, arrays, functions, maps, pointers, interfaces, channels, and more.

---

## `Value`

`reflect.Value` represents an actual runtime value.

```go
v := reflect.ValueOf(100)

fmt.Println(v.Int())
```

Output:

```text
100
```

A `Value` can represent:

- integers
- strings
- structs
- pointers
- slices
- arrays
- maps
- functions
- channels
- interfaces
- etc.

Not every `Value` method is valid for every kind. Calling an inappropriate method can cause a runtime panic.

---

## `Kind`

`Kind` tells you the general category of a value/type.

Examples:

```go
reflect.Int
reflect.String
reflect.Struct
reflect.Slice
reflect.Map
reflect.Pointer
reflect.Interface
```

Example:

```go
v := reflect.ValueOf("hello")

fmt.Println(v.Kind())
```

Output:

```text
string
```

---

## `CanSet`

A crucial rule:

> A reflection value can only be modified if it is settable.

This will not work:

```go
name := "Chandu"

v := reflect.ValueOf(name)

v.SetString("Rahul") // panic
```

But this can work:

```go
name := "Chandu"

v := reflect.ValueOf(&name).Elem()

v.SetString("Rahul")

fmt.Println(name)
```

Output:

```text
Rahul
```

The pointer allows reflection to reach the original variable.

---

# 4. Package-level functions

## `TypeOf`

```go
reflect.TypeOf(value)
```

Returns the runtime type of a value.

```go
t := reflect.TypeOf(42)

fmt.Println(t)
fmt.Println(t.Kind())
```

Output:

```text
int
int
```

Useful when you need to inspect an unknown value.

---

## `ValueOf`

```go
reflect.ValueOf(value)
```

Returns a `reflect.Value` representing the supplied value.

```go
v := reflect.ValueOf(42)

fmt.Println(v.Int())
```

---

## `Zero`

```go
reflect.Zero(t)
```

Creates the zero value for a given type.

```go
t := reflect.TypeOf(0)

v := reflect.Zero(t)

fmt.Println(v.Int())
```

Output:

```text
0
```

For a string, the zero value is `""`.

For a pointer, it is `nil`.

For a struct, it creates a zero-valued struct.

---

## `New`

```go
reflect.New(t)
```

Creates a pointer to a new zero value of type `t`.

```go
t := reflect.TypeOf(0)

v := reflect.New(t)

fmt.Println(v.Type())
```

Conceptually:

```text
reflect.New(int)
       ↓
*int
```

---

## `MakeSlice`

```go
reflect.MakeSlice(sliceType, length, capacity)
```

Creates a slice dynamically.

```go
t := reflect.TypeOf([]int{})

v := reflect.MakeSlice(t, 3, 5)

fmt.Println(v)
```

---

## `MakeMap`

```go
reflect.MakeMap(t)
```

Creates a new map dynamically.

```go
t := reflect.TypeOf(map[string]int{})

m := reflect.MakeMap(t)

m.SetMapIndex(
    reflect.ValueOf("Go"),
    reflect.ValueOf(10),
)

fmt.Println(m.Interface())
```

---

## `MakeMapWithSize`

Similar to `MakeMap`, but accepts an initial size hint.

```go
m := reflect.MakeMapWithSize(t, 100)
```

Useful when you know approximately how many entries will be inserted.

---

## `MakeChan`

Creates a channel dynamically.

```go
t := reflect.ChanOf(reflect.BothDir, reflect.TypeOf(0))

ch := reflect.MakeChan(t, 10)
```

The second argument is the channel buffer size.

---

## `ChanOf`

Creates a channel type dynamically.

```go
t := reflect.ChanOf(
    reflect.BothDir,
    reflect.TypeOf(0),
)
```

Possible directions include:

```go
reflect.SendDir
reflect.RecvDir
reflect.BothDir
```

---

## `FuncOf`

Creates a function type dynamically.

For example, a type conceptually equivalent to:

```go
func(int, string) bool
```

can be constructed with reflection:

```go
t := reflect.FuncOf(
    []reflect.Type{
        reflect.TypeOf(0),
        reflect.TypeOf(""),
    },
    []reflect.Type{
        reflect.TypeOf(false),
    },
    false,
)
```

The final `bool` specifies whether the function is variadic.

---

## `MapOf`

Creates a map type dynamically.

```go
t := reflect.MapOf(
    reflect.TypeOf(""),
    reflect.TypeOf(0),
)
```

This represents:

```go
map[string]int
```

---

## `PtrTo`

Creates a pointer type.

```go
t := reflect.PtrTo(reflect.TypeOf(42))
```

Conceptually:

```go
*int
```

---

## `SliceOf`

Creates a slice type dynamically.

```go
t := reflect.SliceOf(reflect.TypeOf(0))
```

This represents:

```go
[]int
```

---

## `ArrayOf`

Creates an array type dynamically.

```go
t := reflect.ArrayOf(5, reflect.TypeOf(0))
```

Conceptually:

```go
[5]int
```

---

## `StructOf`

Creates a struct type dynamically.

You can construct something conceptually similar to:

```go
struct {
    Name string
    Age  int
}
```

using `reflect.StructOf`.

This is an advanced reflection technique with restrictions, so it is normally used by framework/library authors.

---

## `Append`

```go
reflect.Append(slice, values...)
```

Appends reflection values to a slice.

```go
slice := reflect.ValueOf([]int{1, 2})

slice = reflect.Append(
    slice,
    reflect.ValueOf(3),
)

fmt.Println(slice.Interface())
```

---

## `AppendSlice`

Appends one reflected slice to another.

```go
result := reflect.AppendSlice(dst, src)
```

Both values must represent compatible slices.

---

## `Copy`

Copies elements between reflected slices or arrays.

```go
n := reflect.Copy(dst, src)
```

Returns the number of elements copied.

---

## `Swapper`

```go
swap := reflect.Swapper(slice)
```

Returns a function that swaps two elements of a slice.

```go
numbers := []int{10, 20, 30}

swap := reflect.Swapper(numbers)

swap(0, 2)

fmt.Println(numbers)
```

The first and third elements are exchanged.

---

## `DeepEqual`

```go
reflect.DeepEqual(a, b)
```

Performs deep comparison.

```go
a := []int{1, 2, 3}
b := []int{1, 2, 3}

fmt.Println(reflect.DeepEqual(a, b))
```

Output:

```text
true
```

It can compare structures containing slices, maps, arrays, structs, pointers, and interfaces.

---

## `MakeFunc`

```go
reflect.MakeFunc(type, implementation)
```

Dynamically creates a function with the specified function type.

The implementation receives:

```go
[]reflect.Value
```

and returns:

```go
[]reflect.Value
```

This is an advanced API used when a function must be generated dynamically.

---

## `VisibleFields`

```go
reflect.VisibleFields(t)
```

Returns fields visible through `FieldByName`.

It is particularly useful for structs containing embedded fields.

```go
fields := reflect.VisibleFields(
    reflect.TypeOf(MyStruct{}),
)
```

---

## `TypeAssert`

Modern Go provides:

```go
reflect.TypeAssert[T](v)
```

It performs a type assertion on a `reflect.Value` and returns:

```go
(T, bool)
```

This is useful when moving from reflection back into ordinary Go values.

---

# 5. Important `reflect.Type` methods

Once you have:

```go
t := reflect.TypeOf(value)
```

you can investigate the type.

## `Kind`

```go
t.Kind()
```

Returns the broad category:

```text
Bool
Int
String
Slice
Array
Map
Struct
Pointer
Interface
Func
Chan
...
```

---

## `Name`

Returns the type's name.

```go
type Person struct{}

t := reflect.TypeOf(Person{})

fmt.Println(t.Name())
```

Output:

```text
Person
```

For unnamed types, the name can be empty.

---

## `PkgPath`

Returns the package path associated with a defined type.

Useful for identifying where a named type originates.

---

## `String`

Returns a human-readable representation of the type.

```go
fmt.Println(t.String())
```

---

## `Bits`

For numeric types, returns their size in bits.

```go
reflect.TypeOf(int64(0)).Bits()
```

Returns:

```text
64
```

---

## `Size`

Returns the number of bytes required by a value of the type.

```go
t.Size()
```

---

## `Align`

Returns the alignment requirement of the type.

Primarily useful for low-level programming.

---

## `FieldAlign`

Returns the alignment when the type is used as a struct field.

---

## `Comparable`

Reports whether values of the type can be compared using `==`.

---

## `AssignableTo`

```go
t.AssignableTo(other)
```

Checks whether a value of this type can be assigned to another type.

---

## `ConvertibleTo`

```go
t.ConvertibleTo(other)
```

Checks whether values can be converted to another type.

For example, some numeric types are convertible:

```text
int → int64
```

Conversion and assignment are not the same concept.

---

## `Implements`

```go
t.Implements(interfaceType)
```

Checks whether a type implements a particular interface.

This is very useful for framework code.

---

## `Elem`

For types such as pointers, arrays, slices, maps, and channels:

```go
t.Elem()
```

returns the element type.

Example:

```go
t := reflect.TypeOf([]int{})

fmt.Println(t.Elem())
```

Output:

```text
int
```

---

## `Len`

For arrays, returns the array length.

```go
reflect.TypeOf([5]int{}).Len()
```

returns:

```text
5
```

---

## `Key`

For maps:

```go
t.Key()
```

returns the map's key type.

---

## `ChanDir`

For channel types:

```go
t.ChanDir()
```

tells you whether the channel is:

- send-only
- receive-only
- bidirectional

---

## `NumField`

For structs:

```go
t.NumField()
```

returns the number of fields.

---

## `Field`

```go
t.Field(i)
```

returns information about the `i`th struct field as a `reflect.StructField`.

---

## `FieldByName`

```go
t.FieldByName("Name")
```

looks up a struct field by name.

---

## `FieldByIndex`

```go
t.FieldByIndex(index)
```

looks up a field using its index path.

This is useful with embedded structs.

---

## `FieldByNameFunc`

Allows custom matching logic when searching for a field.

Example:

```go
t.FieldByNameFunc(func(name string) bool {
    return name == "Name"
})
```

---

## `NumMethod`

Returns the number of methods in the method set.

---

## `Method`

```go
t.Method(i)
```

returns information about the `i`th method.

---

## `MethodByName`

```go
t.MethodByName("DoSomething")
```

looks up a method by name.

---

## `In`

Used with function types.

```go
t.In(i)
```

returns an input parameter type.

---

## `Out`

Used with function types.

```go
t.Out(i)
```

returns an output parameter type.

---

## `NumIn`

Returns the number of function input parameters.

---

## `NumOut`

Returns the number of function results.

---

## `IsVariadic`

Reports whether a function type is variadic.

For:

```go
func(x ...int)
```

it returns `true`.

---

# 6. Important `reflect.Value` methods

Suppose:

```go
v := reflect.ValueOf(data)
```

## `Kind`

```go
v.Kind()
```

Returns the value's `Kind`.

---

## `Type`

```go
v.Type()
```

Returns its `reflect.Type`.

---

## `Interface`

```go
v.Interface()
```

Converts the reflection value back into an ordinary Go interface value.

```go
v := reflect.ValueOf(42)

x := v.Interface()

fmt.Println(x)
```

---

## `IsValid`

Checks whether a `Value` represents a valid value.

```go
if !v.IsValid() {
    fmt.Println("invalid")
}
```

The zero `reflect.Value` represents no value.

---

## `IsZero`

Checks whether the value is its type's zero value.

```go
v := reflect.ValueOf(0)

fmt.Println(v.IsZero())
```

Output:

```text
true
```

---

## `IsNil`

Checks whether nil-able values are nil.

Applicable to:

- pointers
- maps
- slices
- functions
- interfaces
- channels

Do not call it indiscriminately on every kind.

---

## `CanSet`

Determines whether the value can be changed.

```go
v.CanSet()
```

---

## `CanAddr`

Determines whether the value is addressable.

---

## `CanInterface`

Determines whether calling:

```go
v.Interface()
```

is permitted.

---

# 7. Reading primitive values

Reflection provides type-specific getters.

## `Bool`

```go
v.Bool()
```

Reads a boolean.

## `Int`

```go
v.Int()
```

Reads signed integers.

## `Uint`

```go
v.Uint()
```

Reads unsigned integers.

## `Float`

```go
v.Float()
```

Reads floating-point values.

## `Complex`

```go
v.Complex()
```

Reads complex numbers.

## `String`

```go
v.String()
```

Reads strings.

These methods require the appropriate `Kind`; otherwise reflection can panic.

---

# 8. Setting values

## `Set`

```go
v.Set(other)
```

Sets the value from another `reflect.Value`.

## `SetBool`

```go
v.SetBool(true)
```

## `SetInt`

```go
v.SetInt(100)
```

## `SetUint`

```go
v.SetUint(100)
```

## `SetFloat`

```go
v.SetFloat(3.14)
```

## `SetComplex`

```go
v.SetComplex(1 + 2i)
```

## `SetString`

```go
v.SetString("Go")
```

## `SetBytes`

Sets a `[]byte`.

## `SetLen`

Changes the length of a slice.

## `SetCap`

Changes the capacity of a slice.

The value generally must be settable.

---

# 9. Working with structs

## `Field`

```go
v.Field(i)
```

Gets a struct field by index.

---

## `FieldByName`

```go
v.FieldByName("Name")
```

Gets a struct field by name.

This is one of the most frequently used reflection operations.

---

## `FieldByIndex`

```go
v.FieldByIndex(index)
```

Gets a nested field using an index sequence.

---

## `FieldByNameFunc`

Searches fields using a custom matching function.

---

# 10. Working with slices and arrays

## `Len`

```go
v.Len()
```

Gets the length.

---

## `Cap`

```go
v.Cap()
```

Gets the capacity.

---

## `Index`

```go
v.Index(i)
```

Gets an element.

Example:

```go
numbers := []int{10, 20, 30}

v := reflect.ValueOf(numbers)

fmt.Println(v.Index(1).Int())
```

Output:

```text
20
```

---

## `Slice`

Creates a sub-slice.

```go
v.Slice(1, 3)
```

Conceptually:

```text
[10 20 30 40]
     ↓
[20 30]
```

---

## `Slice3`

Allows specifying:

```text
low
high
max
```

Equivalent to Go's three-index slice expression.

---

## `SetLen`

Changes a slice's length.

---

## `SetCap`

Changes its capacity.

---

## `Grow`

Grows a slice's capacity if necessary.

---

# 11. Working with maps

## `MapIndex`

```go
v.MapIndex(key)
```

Retrieves a map value.

---

## `SetMapIndex`

```go
v.SetMapIndex(key, value)
```

Adds or updates a map entry.

A zero `Value` can be used to delete an entry.

---

## `MapKeys`

Returns all keys.

```go
keys := v.MapKeys()
```

---

## `MapRange`

Provides an iterator for map entries.

```go
iter := v.MapRange()

for iter.Next() {
    fmt.Println(iter.Key())
    fmt.Println(iter.Value())
}
```

---

## `MapIndex` vs `MapRange`

Use `MapIndex` when you want a particular key.

Use `MapRange` when iterating through the map.

---

# 12. Working with pointers

## `Elem`

Dereferences a pointer or interface.

```go
x := 100

v := reflect.ValueOf(&x)

fmt.Println(v.Elem().Int())
```

Output:

```text
100
```

This is one of the most important reflection operations.

---

## `Addr`

Gets the address of an addressable value.

---

## `Pointer`

Returns a pointer representation for applicable kinds.

This is an advanced operation and should not be confused with ordinary safe Go pointer manipulation.

---

## `UnsafePointer`

Provides an unsafe pointer representation for applicable values.

Use this only when you understand the consequences.

---

# 13. Working with interfaces

## `Elem`

If a `Value` contains an interface, `Elem()` can retrieve the concrete value stored inside it.

---

## `IsNil`

Can determine whether an interface contains `nil`.

Be careful with the distinction:

```go
var x interface{} = nil
```

versus:

```go
var p *Person = nil
var x interface{} = p
```

The second interface itself is not nil because it contains a typed nil pointer.

Reflection can make these distinctions visible.

---

# 14. Calling functions dynamically

## `Call`

```go
v.Call(args)
```

Calls a function represented by a `reflect.Value`.

Example:

```go
func add(a, b int) int {
    return a + b
}

v := reflect.ValueOf(add)

result := v.Call([]reflect.Value{
    reflect.ValueOf(10),
    reflect.ValueOf(20),
})

fmt.Println(result[0].Int())
```

Output:

```text
30
```

This is powerful, but slower and more complex than an ordinary function call.

---

## `CallSlice`

Used for calling variadic functions when the final argument is represented as a slice.

---

# 15. Calling methods dynamically

## `Method`

Gets a method by index.

```go
v.Method(i)
```

---

## `MethodByName`

Gets a method by name.

```go
method := v.MethodByName("Print")
```

Then:

```go
method.Call(nil)
```

can invoke it if the method takes no arguments.

---

## `NumMethod`

Returns the number of accessible methods.

---

## `Methods`

Modern Go also provides an iterator over the methods of a value:

```go
for method, value := range v.Methods() {
    fmt.Println(method.Name)
}
```

---

# 16. Conversion

## `Convert`

```go
v.Convert(targetType)
```

Converts a reflection value to another compatible type.

Example:

```go
v := reflect.ValueOf(int32(10))

target := reflect.TypeOf(int64(0))

converted := v.Convert(target)

fmt.Println(converted.Int())
```

Conversion must follow Go's type-conversion rules.

---

# 17. Checking convertibility

## `CanConvert`

```go
v.CanConvert(targetType)
```

Checks whether a conversion is possible before calling `Convert`.

This is a good defensive programming practice.

---

# 18. Creating values dynamically

Several APIs are useful when constructing values whose types are not known at compile time.

### `New`

Creates a pointer to a new zero value.

### `NewAt`

Creates a `Value` representing memory at a specified address.

`NewAt` is an advanced/unsafe-oriented API and should generally be avoided unless there is a specific low-level reason.

### `MakeSlice`

Creates a slice.

### `MakeMap`

Creates a map.

### `MakeChan`

Creates a channel.

### `MakeFunc`

Creates a function.

---

# 19. Channels and reflection

Reflection can dynamically interact with channels.

## `Send`

```go
v.Send(x)
```

Sends a value through a channel.

---

## `TrySend`

Attempts a send without blocking and reports whether it succeeded.

---

## `Recv`

```go
value, ok := v.Recv()
```

Receives from a channel.

---

## `TryRecv`

Attempts to receive without blocking.

---

## `Close`

Closes the channel.

---

# 20. Dynamic `select`

Reflection can dynamically construct a `select` operation.

Instead of writing:

```go
select {
case x := <-ch1:
    ...
case x := <-ch2:
    ...
}
```

you can construct cases dynamically using:

```go
reflect.SelectCase
```

and call:

```go
reflect.Select(cases)
```

This is useful when the number of channels is not known at compile time.

---

# 21. `SelectCase`

A `SelectCase` describes one case of a dynamic `select`.

It can represent:

- receive
- send
- default

The associated `SelectDir` identifies the direction.

---

# 22. Struct tags

Reflection is heavily used with struct tags.

Example:

```go
type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}
```

You can inspect the tag:

```go
t := reflect.TypeOf(User{})

field, _ := t.FieldByName("Name")

fmt.Println(field.Tag.Get("json"))
```

Output:

```text
name
```

---

## `StructTag.Get`

```go
tag.Get("json")
```

Returns the value associated with the requested tag key.

---

## `StructTag.Lookup`

```go
value, ok := tag.Lookup("json")
```

`Lookup` is useful when you need to distinguish:

```text
tag doesn't exist
```

from:

```text
tag exists but contains an empty value
```

---

# 23. `StructField`

When inspecting struct fields:

```go
field := t.Field(0)
```

you get a `reflect.StructField`.

It provides information such as:

```text
Name
PkgPath
Type
Tag
Offset
Index
Anonymous
```

---

## `StructField.IsExported`

Reports whether a struct field is exported.

For example:

```go
Name string
```

is exported because it begins with an uppercase letter.

```go
name string
```

is unexported.

---

# 24. `Kind`

`reflect.Kind` is an important enumeration.

Common kinds include:

```text
Invalid
Bool
Int
Int8
Int16
Int32
Int64
Uint
Uint8
Uint16
Uint32
Uint64
Uintptr
Float32
Float64
Complex64
Complex128
Array
Chan
Func
Interface
Map
Pointer
Slice
String
Struct
UnsafePointer
```

A common reflection pattern is:

```go
switch v.Kind() {
case reflect.String:
    // string
case reflect.Int:
    // integer
case reflect.Struct:
    // struct
case reflect.Slice:
    // slice
}
```

This is safer than blindly calling methods that may not apply to the value.

---

# 25. `ChanDir`

`ChanDir` represents channel direction.

Values include:

```text
SendDir
RecvDir
BothDir
```

Its `String()` method gives a readable representation.

---

# 26. `Method`

A `reflect.Method` describes a method.

Important fields include:

```text
Name
PkgPath
Type
Func
Index
```

Its:

```go
IsExported()
```

method tells you whether the method is exported.

---

# 27. `MapIter`

`MapIter` is used for dynamically iterating through maps.

Its main methods are:

```text
Next()
Key()
Value()
Reset()
```

Example:

```go
iter := value.MapRange()

for iter.Next() {
    key := iter.Key()
    value := iter.Value()

    fmt.Println(key.Interface(), value.Interface())
}
```

---

# 28. Why reflection is powerful

Reflection enables software where the programmer does **not necessarily know the type beforehand**.

For example:

```text
JSON data
   ↓
unknown Go value
   ↓
reflection
   ↓
inspect fields
   ↓
apply rules
   ↓
construct result
```

This is why reflection is heavily associated with framework and library development.

---

# 29. Real-world application #1: JSON/XML/database libraries

Consider:

```go
type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}
```

A serialization library needs to discover:

```text
User
 ├── Name → string
 └── Age  → int
```

It can inspect:

```go
reflect.TypeOf(User{})
```

then discover:

```go
NumField()
Field()
Field.Tag
```

and use struct tags to determine external field names.

This general technique is why reflection is important in serialization, ORM, validation, configuration, and similar libraries.

---

# 30. Real-world application #2: Dependency injection / frameworks

Imagine:

```go
func NewService(db *Database, logger *Logger) *Service
```

A dependency injection framework can inspect the function's type:

```text
NumIn()
In(0)
In(1)
Out(0)
```

and determine:

```text
argument 1 → *Database
argument 2 → *Logger
result     → *Service
```

It can then construct dependencies and dynamically invoke the function.

This is a classic example of reflection being useful when a framework needs to operate on arbitrary user-defined types.

---

# 31. Three common beginner mistakes

## Mistake 1: Calling the wrong method for a `Kind`

This can panic:

```go
v := reflect.ValueOf("hello")

v.Int()
```

because the value is a string.

### Better approach

Check the kind first:

```go
switch v.Kind() {
case reflect.String:
    fmt.Println(v.String())
}
```

---

## Mistake 2: Forgetting `CanSet`

This is a common mistake:

```go
x := 10

v := reflect.ValueOf(x)

v.SetInt(20)
```

This panics because `v` is not settable.

Instead:

```go
v := reflect.ValueOf(&x).Elem()

v.SetInt(20)
```

---

## Mistake 3: Using reflection everywhere

Reflection is not automatically better than ordinary Go code.

This:

```go
func add(a, b int) int {
    return a + b
}
```

is clearer and safer than dynamically calling the function through reflection when you already know its type.

Use reflection when **runtime type information is actually required**.

---

# 32. Three progressively challenging exercises

## Exercise 1 — Basic Type Inspector

Write:

```go
func inspect(value any)
```

using `reflect` to print:

1. The concrete type.
2. The `Kind`.
3. Whether the value is valid.
4. Whether the value is nil when the kind supports nil.
5. The value itself.

Test it with:

```go
42
"hello"
3.14
true
[]int{1, 2, 3}
map[string]int{"Go": 100}
```

**Goal:** Become comfortable with `TypeOf`, `ValueOf`, `Kind`, `IsValid`, and safe use of `IsNil`.

---

## Exercise 2 — Generic Struct Inspector

Create:

```go
type Employee struct {
    Name       string `json:"name"`
    Department string `json:"department"`
    Salary     int    `json:"salary"`
}
```

Write:

```go
func inspectStruct(value any)
```

that dynamically:

1. Determines whether the supplied value is a struct.
2. Prints every field name.
3. Prints every field type.
4. Prints every field value.
5. Prints the `json` struct tag.
6. Reports whether each field is exported.
7. Handles embedded structs.

Do not hard-code the field names.

**Goal:** Practice `Type`, `Value`, `NumField`, `Field`, `FieldByName`, `StructTag`, `Get`, `Lookup`, and `IsExported`.

---

## Exercise 3 — Dynamic Function Runner

Create:

```go
func runFunction(fn any, args ...any) []any
```

It should use reflection to:

1. Verify that `fn` is actually a function.
2. Inspect its input parameters.
3. Verify that the number of supplied arguments is correct.
4. Verify that the arguments have compatible types.
5. Dynamically invoke the function.
6. Convert returned `reflect.Value` objects back into ordinary Go values.
7. Support functions returning multiple values.
8. Produce useful errors rather than panicking when the input is invalid.

Test it with several functions having different signatures, including a function with multiple return values.

**Goal:** Combine `Type`, `Value`, `Kind`, `NumIn`, `In`, `NumOut`, `Call`, `Interface`, and type validation.

---

# 33. Practical reflection learning path

Recommended order:

```text
1. any / interface{}
       ↓
2. reflect.Type
       ↓
3. reflect.Value
       ↓
4. Kind
       ↓
5. TypeOf / ValueOf
       ↓
6. Struct fields
       ↓
7. CanSet + Elem
       ↓
8. Slices and maps
       ↓
9. Struct tags
       ↓
10. Methods
       ↓
11. Call
       ↓
12. MakeFunc
       ↓
13. MakeMap / MakeSlice / StructOf
       ↓
14. Select / channels
       ↓
15. Framework-level reflection
```

The most important rule to remember is:

> **`Type` tells you what something is; `Value` lets you inspect or manipulate the actual runtime value.**

Also remember that the zero `reflect.Value` is special: `IsValid()` returns false, `Kind()` is `Invalid`, and most other operations on it panic.

---

# Thought-provoking question

Suppose Go's generics can solve a problem without reflection.

**What reasons might still justify choosing reflection instead—and what trade-offs in type safety, performance, readability, and maintainability would you accept by doing so?**

That question gets to the heart of *why reflection exists*, rather than merely learning how its APIs work.
