# Go `encoding/xml` Package — Complete Learning Guide

## 1. What is `encoding/xml`?

The Go `encoding/xml` package is part of the standard library and provides tools for reading, writing, and processing XML 1.0 documents, including XML namespaces.

Import it with:

```go
import "encoding/xml"
```

The package lets you convert between Go data and XML:

```text
Marshal
Go struct ──────────→ XML

Unmarshal
XML ────────────────→ Go struct
```

For example:

```go
type Person struct {
    Name string `xml:"name"`
    Age  int    `xml:"age"`
}
```

can represent:

```xml
<Person>
    <name>Chandu</name>
    <age>25</age>
</Person>
```

The package uses struct tags to tell Go how fields correspond to XML elements and attributes.

---

## 2. When is `encoding/xml` commonly used?

Typical situations include:

### 1. SOAP APIs

Many enterprise and banking systems still expose SOAP/XML APIs.

```text
Go Application
     ↓
SOAP Request
     ↓
XML
     ↓
Bank / Enterprise System
```

### 2. XML configuration files

For example:

```xml
<config>
    <server>
        <host>localhost</host>
        <port>8080</port>
    </server>
</config>
```

### 3. RSS/Atom feeds

News and content systems commonly exchange structured XML.

### 4. Legacy APIs

Older systems may expose XML instead of JSON.

### 5. Enterprise integrations

ERP, banking, logistics, government, and other enterprise systems frequently use XML.

---

# 3. Simple example

Let's create an XML document from a Go struct.

```go
package main

import (
    "encoding/xml"
    "fmt"
)

type Person struct {
    XMLName xml.Name `xml:"person"`
    Name    string   `xml:"name"`
    Age     int      `xml:"age"`
    Email   string   `xml:"email"`
}

func main() {

    person := Person{
        Name:  "Chandu",
        Age:   25,
        Email: "chandu@example.com",
    }

    data, err := xml.MarshalIndent(person, "", "    ")
    if err != nil {
        panic(err)
    }

    fmt.Println(string(data))
}
```

Output:

```xml
<person>
    <name>Chandu</name>
    <age>25</age>
    <email>chandu@example.com</email>
</person>
```

`xml.MarshalIndent(person, "", "    ")` converts the Go struct into nicely formatted XML.

---

# 4. Reading XML with `Unmarshal`

Suppose we receive:

```xml
<person>
    <name>Chandu</name>
    <age>25</age>
</person>
```

We can convert it back to a Go struct:

```go
package main

import (
    "encoding/xml"
    "fmt"
)

type Person struct {
    Name string `xml:"name"`
    Age  int    `xml:"age"`
}

func main() {

    data := []byte(`
        <person>
            <name>Chandu</name>
            <age>25</age>
        </person>
    `)

    var person Person

    err := xml.Unmarshal(data, &person)
    if err != nil {
        panic(err)
    }

    fmt.Println(person.Name)
    fmt.Println(person.Age)
}
```

Output:

```text
Chandu
25
```

Notice the `&`:

```go
xml.Unmarshal(data, &person)
```

We pass a pointer because `Unmarshal` needs to modify `person`.

---

# 5. Every package-level function

The package provides five package-level functions:

```text
Escape
EscapeText
Marshal
MarshalIndent
Unmarshal
```

Let's understand each one.

---

## Function 1: `xml.Marshal()`

### Purpose

Converts a Go value into XML.

```go
xml.Marshal(v)
```

Signature:

```go
func Marshal(v any) ([]byte, error)
```

It returns:

```text
[]byte → XML data
error  → possible error
```

Example:

```go
type Person struct {
    Name string `xml:"name"`
    Age  int    `xml:"age"`
}

person := Person{
    Name: "Chandu",
    Age:  25,
}

data, err := xml.Marshal(person)
if err != nil {
    panic(err)
}

fmt.Println(string(data))
```

Result:

```xml
<Person><name>Chandu</name><age>25</age></Person>
```

### When should you use it?

Use `Marshal` when you need XML in memory:

```text
Go struct
   ↓
Marshal
   ↓
[]byte
   ↓
send/store/write
```

For example:

```go
http.Post(..., "application/xml", bytes.NewBuffer(data))
```

### Important behavior

`Marshal` can handle:

- structs
- arrays
- slices
- pointers
- interfaces
- strings
- numbers
- booleans

It returns an error for unsupported values such as functions, channels, and other unsupported data.

---

# 6. `xml.MarshalIndent()`

### Purpose

Same basic job as `Marshal`, but produces human-readable formatted XML.

Signature:

```go
func MarshalIndent(v any, prefix, indent string) ([]byte, error)
```

Example:

```go
data, err := xml.MarshalIndent(person, "", "    ")
```

Output:

```xml
<Person>
    <name>Chandu</name>
    <age>25</age>
</Person>
```

Compare that with normal `Marshal`:

```xml
<Person><name>Chandu</name><age>25</age></Person>
```

### Parameters

```go
xml.MarshalIndent(
    person,
    "",      // prefix
    "    ",  // indentation
)
```

The second argument is the prefix.

The third argument determines indentation.

You can use:

```go
"    "
```

or:

```go
"\t"
```

### When to use it?

Use `MarshalIndent` when humans need to read the XML:

- debugging
- logs
- configuration files
- generated XML files
- development

For high-volume machine-to-machine communication, compact XML from `Marshal` can be preferable.

---

# 7. `xml.Unmarshal()`

This is the reverse of `Marshal`.

```text
Marshal

Go
 ↓
XML


Unmarshal

XML
 ↓
Go
```

Signature:

```go
func Unmarshal(data []byte, v any) error
```

Example:

```go
xmlData := []byte(`
<Person>
    <name>Chandu</name>
    <age>25</age>
</Person>
`)

var person Person

err := xml.Unmarshal(xmlData, &person)
if err != nil {
    panic(err)
}
```

Now:

```go
person.Name
```

contains:

```text
Chandu
```

and:

```go
person.Age
```

contains:

```text
25
```

### Critical beginner concept

The XML field must correspond to your struct.

For example:

```xml
<name>Chandu</name>
```

requires something like:

```go
Name string `xml:"name"`
```

Struct tags are extremely important when working with XML.

---

# 8. `xml.EscapeText()`

This function escapes special XML characters.

Signature:

```go
func EscapeText(w io.Writer, s []byte) error
```

Suppose you have:

```text
5 < 10 && 10 > 5
```

Those characters have special meaning in XML.

You can escape them:

```go
var buffer bytes.Buffer

err := xml.EscapeText(
    &buffer,
    []byte(`5 < 10 && 10 > 5`),
)

if err != nil {
    panic(err)
}

fmt.Println(buffer.String())
```

The output contains XML-safe escaped representations:

```text
5 &lt; 10 &amp;&amp; 10 &gt; 5
```

### Why is escaping important?

This:

```xml
<message>
    5 < 10
</message>
```

contains a `<` that can be interpreted as the beginning of XML markup.

Escaping makes it safe:

```xml
<message>
    5 &lt; 10
</message>
```

---

# 9. `xml.Escape()`

Signature:

```go
func Escape(w io.Writer, s []byte)
```

It also escapes XML text.

However, there is an important historical detail:

```text
Escape
```

is essentially the older compatibility version of:

```text
EscapeText
```

For modern Go code, `EscapeText` is generally the preferred function.

---

# 10. Important `encoding/xml` types

The package contains important types including:

```text
Attr
CharData
Comment
Decoder
Directive
Encoder
EndElement
Name
ProcInst
StartElement
Token
UnmarshalError
Marshaler
Unmarshaler
UnmarshalerAttr
```

The most important for beginners are:

```text
Encoder
Decoder
StartElement
EndElement
Attr
Name
Token
```

---

# 11. `xml.Encoder`

An `Encoder` lets you write XML directly to an `io.Writer`.

You create one with:

```go
encoder := xml.NewEncoder(os.Stdout)
```

Instead of:

```text
Go struct
 ↓
[]byte
 ↓
write []byte
```

you can have:

```text
Go data
 ↓
Encoder
 ↓
io.Writer
```

The writer could be:

- `os.File`
- `bytes.Buffer`
- HTTP response
- network connection

---

# 12. `xml.NewEncoder()`

Signature:

```go
func NewEncoder(w io.Writer) *Encoder
```

It creates an XML encoder around an `io.Writer`.

Example:

```go
file, err := os.Create("person.xml")
if err != nil {
    panic(err)
}
defer file.Close()

encoder := xml.NewEncoder(file)

err = encoder.Encode(person)
if err != nil {
    panic(err)
}
```

This writes XML directly to the file.

---

# 13. `Encoder.Encode()`

Signature:

```go
func (enc *Encoder) Encode(v any) error
```

It writes a Go value as XML.

Example:

```go
encoder := xml.NewEncoder(os.Stdout)

err := encoder.Encode(person)
if err != nil {
    panic(err)
}
```

Unlike:

```go
xml.Marshal()
```

which returns:

```go
[]byte
```

`Encoder.Encode()` writes directly to the underlying `io.Writer`.

---

# 14. `Encoder.Indent()`

Signature:

```go
func (enc *Encoder) Indent(prefix, indent string)
```

Example:

```go
encoder.Indent("", "    ")
```

This produces nicely formatted XML.

Very useful for:

- configuration files
- debugging
- generated XML
- human-readable output

---

# 15. `Encoder.Flush()`

Signature:

```go
func (enc *Encoder) Flush() error
```

It forces buffered XML data to be written to the underlying writer.

Normally:

```go
Encode()
```

handles flushing for you.

But when manually writing XML tokens using:

```go
EncodeToken()
```

you may need:

```go
encoder.Flush()
```

---

# 16. `Encoder.Close()`

Signature:

```go
func (enc *Encoder) Close() error
```

This was added in Go 1.20.

It:

1. finishes buffered XML output
2. checks whether the XML is structurally valid
3. reports errors such as unclosed elements

For example, if you manually create:

```xml
<person>
```

but forget:

```xml
</person>
```

`Close()` can detect that the XML is incomplete.

---

# 17. `Encoder.EncodeToken()`

This is a more advanced API:

```go
encoder.EncodeToken(token)
```

Instead of:

```go
encoder.Encode(person)
```

you can manually construct XML tokens.

Conceptually:

```text
StartElement
     ↓
Character Data
     ↓
EndElement
```

For example:

```go
encoder.EncodeToken(
    xml.StartElement{
        Name: xml.Name{Local: "person"},
    },
)

encoder.EncodeToken(
    xml.CharData("Chandu"),
)

encoder.EncodeToken(
    xml.EndElement{
        Name: xml.Name{Local: "person"},
    },
)

encoder.Flush()
```

This gives you much more control over XML generation.

---

# 18. `Encoder.EncodeElement()`

Signature:

```go
func (enc *Encoder) EncodeElement(
    v any,
    start StartElement,
) error
```

This is particularly useful when implementing custom:

```go
MarshalXML()
```

For example:

```go
return e.EncodeElement(value, start)
```

It tells the encoder:

> Encode this Go value using this XML start element.

---

# 19. `xml.Decoder`

`Decoder` is the streaming counterpart to `Encoder`.

You create one using:

```go
decoder := xml.NewDecoder(reader)
```

It reads XML from an `io.Reader`.

Example:

```go
file, err := os.Open("person.xml")
if err != nil {
    panic(err)
}
defer file.Close()

decoder := xml.NewDecoder(file)

var person Person

err = decoder.Decode(&person)
if err != nil {
    panic(err)
}
```

---

# 20. `xml.NewDecoder()`

Signature:

```go
func NewDecoder(r io.Reader) *Decoder
```

It creates an XML decoder around an `io.Reader`.

Useful for:

- files
- HTTP responses
- network streams
- large XML documents

---

# 21. `Decoder.Decode()`

Signature:

```go
func (d *Decoder) Decode(v any) error
```

It reads XML and decodes it into a Go value.

Conceptually:

```text
XML stream
    ↓
Decoder
    ↓
Go struct
```

Example:

```go
decoder.Decode(&person)
```

This is similar to:

```go
xml.Unmarshal(data, &person)
```

but `Decoder` works directly with an `io.Reader`, which is particularly useful for streaming data.

---

# 22. `Decoder.DecodeElement()`

Signature:

```go
func (d *Decoder) DecodeElement(
    v any,
    start *StartElement,
) error
```

This is commonly used when implementing:

```go
UnmarshalXML()
```

Imagine you already encountered:

```xml
<person>
```

and want to decode everything inside that element into a Go value.

That's where:

```go
DecodeElement()
```

is useful.

---

# 23. `Decoder.Token()`

One of the most important advanced methods is:

```go
Token()
```

It reads XML one token at a time.

The XML:

```xml
<person>
    <name>Chandu</name>
</person>
```

can conceptually be viewed as:

```text
StartElement(person)

StartElement(name)

CharData("Chandu")

EndElement(name)

EndElement(person)
```

This is useful when you don't know the exact XML structure beforehand or need very fine-grained control.

---

# 24. `Decoder.RawToken()`

`RawToken()` is similar to `Token()` but intentionally performs less XML processing.

It does not verify that start/end elements match and does not translate namespace prefixes into their corresponding URLs.

This is generally an advanced API.

For beginners:

```text
Token()
```

is usually preferable.

---

# 25. `Decoder.InputOffset()`

```go
decoder.InputOffset()
```

returns the current byte position in the XML input stream.

This can be useful for debugging malformed XML.

Conceptually:

```text
XML error
     ↓
Find approximate byte offset
     ↓
Investigate XML
```

---

# 26. `Decoder.InputPos()`

`InputPos()` provides the current line and column position.

Example:

```go
line, column := decoder.InputPos()
```

This is particularly useful when diagnosing XML parsing errors.

---

# 27. Struct tags — extremely important

A huge part of `encoding/xml` is understanding tags.

### Basic element

```go
Name string `xml:"name"`
```

maps to:

```xml
<name>Chandu</name>
```

### Attribute

```go
ID int `xml:"id,attr"`
```

maps to:

```xml
<person id="10">
```

### Ignore field

```go
Password string `xml:"-"`
```

The field is excluded.

### Omit empty values

```go
Email string `xml:"email,omitempty"`
```

If `Email` is empty, it is omitted.

### Character data

```go
Text string `xml:",chardata"`
```

### CDATA

```go
Text string `xml:",cdata"`
```

### Nested XML

```go
FirstName string `xml:"name>first"`
```

can produce:

```xml
<name>
    <first>Chandu</first>
</name>
```

These struct-tag rules are central to both `Marshal` and `Unmarshal`.

---

# 28. XML attributes

Consider:

```xml
<person id="100" country="India">
    <name>Chandu</name>
</person>
```

You can represent it as:

```go
type Person struct {
    ID      int    `xml:"id,attr"`
    Country string `xml:"country,attr"`
    Name    string `xml:"name"`
}
```

The `,attr` tells Go:

> This is an XML attribute, not a child element.

---

# 29. `xml.Name`

`xml.Name` represents an XML name.

It contains:

```go
type Name struct {
    Space string
    Local string
}
```

For:

```xml
<person>
```

you might get:

```text
Space: ""
Local: "person"
```

It becomes particularly important when dealing with XML namespaces.

---

# 30. XML namespaces

For example:

```xml
<soap:Envelope
    xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
```

XML namespaces are supported by `encoding/xml`.

You may encounter:

```go
xml.Name{
    Space: "http://schemas.xmlsoap.org/soap/envelope/",
    Local: "Envelope",
}
```

This is especially important when integrating with SOAP and enterprise systems.

---

# 31. `xml.Attr`

An XML attribute is represented by:

```go
type Attr struct {
    Name  Name
    Value string
}
```

For:

```xml
<person id="100">
```

the attribute can conceptually be:

```text
Name:
    Local = "id"

Value:
    "100"
```

---

# 32. `xml.CharData`

Character data represents text inside XML.

For:

```xml
<name>Chandu</name>
```

the text:

```text
Chandu
```

is character data.

`CharData` is:

```go
type CharData []byte
```

It also has:

```go
Copy()
```

which creates an independent copy of the data.

This matters when working directly with XML tokens.

---

# 33. `xml.Comment`

Represents:

```xml
<!-- This is a comment -->
```

The comment's underlying bytes don't include the `<!--` and `-->` delimiters.

It also provides:

```go
Copy()
```

for creating an independent copy.

---

# 34. `xml.Directive`

Represents XML directives such as:

```xml
<!DOCTYPE ...>
```

The underlying bytes exclude the `<!` and `>` delimiters.

It also provides:

```go
Copy()
```

---

# 35. `xml.StartElement`

Represents an opening XML element:

```xml
<person id="10">
```

Conceptually:

```go
xml.StartElement{
    Name: xml.Name{
        Local: "person",
    },
    Attr: []xml.Attr{
        // attributes
    },
}
```

---

# 36. `xml.EndElement`

Represents:

```xml
</person>
```

It contains the corresponding `Name`.

Example:

```go
xml.EndElement{
    Name: xml.Name{
        Local: "person",
    },
}
```

---

# 37. `xml.ProcInst`

Represents an XML processing instruction.

For example:

```xml
<?xml version="1.0"?>
```

It has a target and instruction data.

You'll encounter this more often when working with XML at the token level rather than ordinary struct marshaling.

---

# 38. `xml.Token`

`Token` is the common concept used when processing XML token-by-token.

The token types include:

```text
StartElement
EndElement
CharData
Comment
Directive
ProcInst
```

This makes it possible to process an XML document like a stream:

```text
XML
 ↓
Token
 ↓
Token
 ↓
Token
 ↓
Token
```

rather than loading the entire document into a struct.

---

# 39. Custom marshaling

Sometimes the default struct-to-XML mapping isn't enough.

You can implement:

```go
MarshalXML()
```

through:

```go
xml.Marshaler
```

The interface is:

```go
type Marshaler interface {
    MarshalXML(e *Encoder, start StartElement) error
}
```

This lets your type control exactly how it is represented in XML.

Similarly, custom XML decoding can be implemented using:

```go
UnmarshalXML()
```

through:

```go
xml.Unmarshaler
```

This is useful when the external XML format doesn't naturally map to your Go struct.

---

# 40. Three common beginner mistakes

## Mistake 1 — Forgetting exported fields

Consider:

```go
type Person struct {
    name string
}
```

This is unexported.

Reflection-based XML decoding cannot assign to it.

Instead:

```go
type Person struct {
    Name string
}
```

### Avoid it

Use:

```go
Name string
```

instead of:

```go
name string
```

---

## Mistake 2 — Forgetting `,attr`

Suppose the XML is:

```xml
<person id="10">
```

A beginner might write:

```go
ID int `xml:"id"`
```

But that's looking for an XML element:

```xml
<id>10</id>
```

For an attribute, use:

```go
ID int `xml:"id,attr"`
```

Remember:

```text
<id>10</id>
       ↑
    element

id="10"
   ↑
attribute
```

---

## Mistake 3 — Assuming XML behaves exactly like JSON

XML is not simply "JSON with angle brackets."

XML has:

```text
elements
attributes
namespaces
character data
CDATA
comments
processing instructions
ordered token streams
```

The mapping between XML and ordinary Go data structures has inherent limitations because XML is an order-dependent collection of anonymous values while typical data structures are named fields.

### Avoid it

Before designing your Go structs, inspect the actual XML structure.

Ask:

```text
What are the elements?
What are the attributes?
Are there namespaces?
Are elements repeated?
Is ordering meaningful?
Is CDATA involved?
```

---

# 41. Two real-world applications

## Application 1 — SOAP / enterprise APIs

Imagine your Go application communicates with an enterprise banking service.

The request may look like:

```xml
<soap:Envelope>
    <soap:Body>
        <GetAccount>
            <AccountNumber>12345</AccountNumber>
        </GetAccount>
    </soap:Body>
</soap:Envelope>
```

Your Go application can use:

```text
encoding/xml
      ↓
construct request XML
      ↓
HTTP
      ↓
SOAP server
      ↓
XML response
      ↓
encoding/xml
      ↓
Go structs
```

This is one of the situations where XML remains highly relevant.

---

## Application 2 — RSS/feed processing

Suppose your application reads an RSS feed:

```xml
<rss>
    <channel>
        <title>Technology News</title>

        <item>
            <title>Go Programming</title>
            <link>https://example.com</link>
        </item>
    </channel>
</rss>
```

You can define structs:

```go
type RSS struct {
    Channel Channel `xml:"channel"`
}

type Channel struct {
    Title string `xml:"title"`
    Items []Item `xml:"item"`
}

type Item struct {
    Title string `xml:"title"`
    Link  string `xml:"link"`
}
```

Then:

```go
xml.Unmarshal(data, &rss)
```

converts the XML feed into ordinary Go structures.

---

# 42. Three progressively challenging exercises

## Exercise 1 — Beginner: Student XML

Create a Go program that represents a student:

```text
ID
Name
Age
Course
```

Convert the struct into XML using `encoding/xml`.

Requirements:

- Use `xml.MarshalIndent`.
- Make `ID` an XML attribute.
- Make the remaining values XML elements.
- Print the generated XML.
- Then unmarshal the XML back into another `Student` struct.
- Print the resulting Go values.
- Do not use JSON.

---

## Exercise 2 — Intermediate: Product catalog

Create a product catalog represented by XML.

The XML should conceptually contain:

```text
catalog
 ├── product
 │    ├── name
 │    ├── price
 │    └── category
 │
 ├── product
 │    ├── name
 │    ├── price
 │    └── category
 │
 └── ...
```

Requirements:

- A catalog should contain multiple products.
- Each product must have an ID attribute.
- Use a slice to represent multiple products.
- Use XML struct tags correctly.
- Marshal the catalog into formatted XML.
- Unmarshal the XML back into Go.
- Calculate and display the total price of all products.
- Do not use solutions from this explanation.

---

## Exercise 3 — Advanced: Streaming XML processor

Create a program that reads a potentially very large XML file containing thousands of:

```xml
<transaction>
    ...
</transaction>
```

elements.

Requirements:

- Do not load the entire XML document into memory with `xml.Unmarshal`.
- Use `xml.Decoder`.
- Process the XML using tokens and/or `DecodeElement`.
- Extract transaction information.
- Count the number of transactions.
- Calculate the total transaction amount.
- Detect malformed XML and report useful parsing information.
- Use `InputOffset()` or `InputPos()` to help identify where an error occurs.

The goal is to understand the difference between:

```text
Unmarshal
```

and:

```text
Decoder
```

and why streaming becomes important for large XML documents.

---

# 43. Useful mental model

If you're learning `encoding/xml`, remember this hierarchy:

```text
                    encoding/xml
                         │
          ┌──────────────┴──────────────┐
          │                             │
       Writing                        Reading
          │                             │
      Marshal                        Unmarshal
          │                             │
   MarshalIndent                    Decoder
          │                             │
       Encoder                       Token()
          │                       DecodeElement()
    EncodeToken()
```

For most beginner programs:

```text
Need XML from Go?
        ↓
   xml.Marshal
        or
 xml.MarshalIndent


Have XML and need Go?
        ↓
   xml.Unmarshal


Large XML / streaming?
        ↓
   xml.Decoder


Need precise XML control?
        ↓
   Encoder + Decoder + Tokens
```

---

# 44. Thought-provoking question

Imagine you're building a modern Go microservice where you control both the client and server.

You can choose either:

```text
JSON
```

or:

```text
XML
```

for communication.

If JSON is simpler and usually more compact, **why might an experienced engineer deliberately choose XML anyway?**

Think beyond "legacy systems."

Consider:

- schemas
- namespaces
- validation
- document structure
- interoperability
- backward compatibility
- the kinds of organizations and systems your service needs to communicate with

---

## Quick Summary

| Feature | Purpose |
|---|---|
| `xml.Marshal` | Convert Go value to XML |
| `xml.MarshalIndent` | Convert Go value to formatted XML |
| `xml.Unmarshal` | Convert XML bytes to Go value |
| `xml.Escape` | Escape XML text; legacy compatibility API |
| `xml.EscapeText` | Escape XML text |
| `xml.NewEncoder` | Create streaming XML encoder |
| `Encoder.Encode` | Encode Go value directly to writer |
| `Encoder.EncodeToken` | Write XML tokens manually |
| `Encoder.EncodeElement` | Encode a value with a specific start element |
| `Encoder.Indent` | Format encoder output |
| `Encoder.Flush` | Flush buffered XML |
| `Encoder.Close` | Finish/check XML output |
| `xml.NewDecoder` | Create streaming XML decoder |
| `Decoder.Decode` | Decode XML from a reader |
| `Decoder.DecodeElement` | Decode the current XML element |
| `Decoder.Token` | Read XML one token at a time |
| `Decoder.RawToken` | Read tokens with less processing |
| `Decoder.InputOffset` | Get current byte offset |
| `Decoder.InputPos` | Get current line/column |
| `xml.Name` | Represent XML names/namespaces |
| `xml.Attr` | Represent XML attributes |
| `xml.StartElement` | Represent opening elements |
| `xml.EndElement` | Represent closing elements |
| `xml.CharData` | Represent XML text |
| `xml.Comment` | Represent XML comments |
| `xml.Directive` | Represent XML directives |
| `xml.ProcInst` | Represent processing instructions |
| `xml.Token` | Common token concept |
| `xml.Marshaler` | Custom XML marshaling |
| `xml.Unmarshaler` | Custom XML unmarshaling |
