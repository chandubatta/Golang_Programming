# Go `unicode` Package

The Go standard-library `unicode` package provides Unicode character data and functions for testing properties of Unicode **code points (runes)**. It is especially useful when your program needs to work correctly with text from many languages rather than assuming ASCII-only characters.

## 1. What is the `unicode` package?

Import it with:

```go
import "unicode"
```

The package allows you to answer questions such as:

- Is this rune a letter?
- Is it uppercase or lowercase?
- Is it a decimal digit?
- Is it whitespace?
- Is it punctuation?
- Is it a mathematical symbol?
- Is it a Unicode number?
- How should this rune be converted to uppercase/lowercase?
- Does this rune belong to a particular Unicode range?

Example:

```go
package main

import (
	"fmt"
	"unicode"
)

func main() {
	r := 'A'

	fmt.Println(unicode.IsLetter(r))
	fmt.Println(unicode.IsUpper(r))
	fmt.Println(unicode.ToLower(r))
}
```

Output:

```text
true
true
97
```

The package contains character-property functions, case-conversion functions, Unicode range types, and predefined Unicode tables.

### When is it commonly used?

Typical uses include:

- validating usernames and identifiers
- processing multilingual text
- detecting letters/digits/punctuation
- case conversion
- implementing text parsers
- searching or filtering Unicode characters
- building internationalized applications
- analyzing natural-language text

---

# 2. Important concept: Unicode, code points, runes, and UTF-8

Before learning the functions, this distinction is essential.

Consider:

```go
s := "Hello 世界"
```

A Go string contains **bytes**.

A `rune` represents a **Unicode code point**.

You can iterate through the string as runes:

```go
for _, r := range s {
	fmt.Printf("%c -> U+%04X\n", r, r)
}
```

You might see:

```text
H -> U+0048
e -> U+0065
l -> U+006C
l -> U+006C
o -> U+006F
  -> U+0020
世 -> U+4E16
界 -> U+754C
```

This is why the `unicode` package is normally used with `rune`.

---

# 3. Every function in the `unicode` package

## A. `unicode.In`

```go
func In(r rune, ranges ...*RangeTable) bool
```

`In` checks whether a rune belongs to **at least one** of the supplied Unicode range tables.

Example:

```go
r := '界'

if unicode.In(r, unicode.Han) {
	fmt.Println("This is a Han character")
}
```

You can provide multiple tables:

```go
unicode.In(r, unicode.Latin, unicode.Greek, unicode.Cyrillic)
```

This asks whether `r` belongs to any of those ranges.

### When useful

Use `In` when you want to test membership against specific Unicode categories/scripts rather than using a generic property such as `IsLetter`.

---

## B. `unicode.Is`

```go
func Is(rangeTab *RangeTable, r rune) bool
```

`Is` checks whether a rune belongs to a particular `RangeTable`.

Example:

```go
r := 'A'

if unicode.Is(unicode.Latin, r) {
	fmt.Println("Latin character")
}
```

### `Is` vs `In`

```go
unicode.Is(unicode.Latin, r)
```

checks **one** table.

```go
unicode.In(r, unicode.Latin, unicode.Greek)
```

checks **multiple** tables.

---

## C. `unicode.IsControl`

```go
func IsControl(r rune) bool
```

Checks whether `r` is a Unicode control character.

Example:

```go
fmt.Println(unicode.IsControl('\n'))
```

Output:

```text
true
```

Control characters include characters used for controlling text/data transmission rather than displaying ordinary visible text.

### Important

`IsControl` does not mean "anything unusual." It specifically tests the Unicode control-character category.

---

## D. `unicode.IsDigit`

```go
func IsDigit(r rune) bool
```

Checks whether a rune is a **decimal digit**.

Example:

```go
fmt.Println(unicode.IsDigit('5'))
```

Output:

```text
true
```

An important Unicode distinction is that this isn't simply equivalent to:

```go
r >= '0' && r <= '9'
```

Unicode contains decimal digits from scripts other than Latin/ASCII.

For example:

```go
fmt.Println(unicode.IsDigit('５'))
```

The full-width digit can be recognized as a Unicode decimal digit.

### Useful for

- internationalized numeric input
- text validation
- parsers

---

## E. `unicode.IsGraphic`

```go
func IsGraphic(r rune) bool
```

Determines whether the rune is considered a **graphic character**.

Graphic characters are characters that normally have a visible representation, including categories such as:

- letters
- marks
- numbers
- punctuation
- symbols
- spaces in the appropriate Unicode definition

Example:

```go
fmt.Println(unicode.IsGraphic('A'))
fmt.Println(unicode.IsGraphic('!'))
```

This is useful when you need to distinguish generally displayable characters from control characters.

---

## F. `unicode.IsLetter`

```go
func IsLetter(r rune) bool
```

Checks whether a rune is a Unicode letter.

Example:

```go
characters := []rune{'A', 'é', '中', 'Ж'}

for _, r := range characters {
	fmt.Printf("%c: %v\n", r, unicode.IsLetter(r))
}
```

All of those can be recognized as letters.

### Why this is better than ASCII checks

This:

```go
r >= 'A' && r <= 'Z'
```

only handles ASCII uppercase letters.

This:

```go
unicode.IsLetter(r)
```

handles Unicode letters across many scripts.

---

## G. `unicode.IsLower`

```go
func IsLower(r rune) bool
```

Checks whether a rune is lowercase.

Example:

```go
fmt.Println(unicode.IsLower('a'))
fmt.Println(unicode.IsLower('A'))
```

Output:

```text
true
false
```

It works with Unicode lowercase characters, not just ASCII.

---

## H. `unicode.IsMark`

```go
func IsMark(r rune) bool
```

Checks whether a rune is a Unicode **mark**.

Marks are characters that generally modify or combine with another character.

For example, Unicode can represent accented text using a base character plus a combining mark:

```text
e + combining acute accent
```

rather than necessarily using a single precomposed character such as:

```text
é
```

This function is useful when processing Unicode text at the code-point level.

---

## I. `unicode.IsNumber`

```go
func IsNumber(r rune) bool
```

Checks whether a rune belongs to the Unicode **Number** category.

This is broader than `IsDigit`.

Unicode contains number characters that aren't decimal digits in the usual `0–9` sense.

### Key distinction

```text
IsDigit
    ↓
decimal digits

IsNumber
    ↓
Unicode number category
```

Therefore:

```go
unicode.IsNumber(r)
```

can recognize a broader set of numeric characters than:

```go
unicode.IsDigit(r)
```

This distinction matters when processing international text.

---

## J. `unicode.IsPrint`

```go
func IsPrint(r rune) bool
```

Reports whether a rune is considered printable according to Unicode's definition used by the package.

Example:

```go
fmt.Println(unicode.IsPrint('A'))
fmt.Println(unicode.IsPrint('\n'))
```

A normal letter is printable, while a newline isn't.

### `IsPrint` vs `IsGraphic`

They are related but not identical concepts.

Use:

```go
unicode.IsPrint(r)
```

when your question is:

> Can this character be treated as printable?

Use:

```go
unicode.IsGraphic(r)
```

when your question is:

> Is this a graphic Unicode character?

---

## K. `unicode.IsPunct`

```go
func IsPunct(r rune) bool
```

Checks whether a rune is Unicode punctuation.

Example:

```go
fmt.Println(unicode.IsPunct('.'))
fmt.Println(unicode.IsPunct(','))
fmt.Println(unicode.IsPunct('!'))
```

Output:

```text
true
true
true
```

It also handles punctuation from other writing systems.

### Useful for

Text tokenization:

```go
for _, r := range text {
	if unicode.IsPunct(r) {
		fmt.Printf("Punctuation: %c\n", r)
	}
}
```

---

## L. `unicode.IsSpace`

```go
func IsSpace(r rune) bool
```

Checks whether a rune is Unicode whitespace.

Example:

```go
fmt.Println(unicode.IsSpace(' '))
fmt.Println(unicode.IsSpace('\n'))
fmt.Println(unicode.IsSpace('\t'))
```

Output:

```text
true
true
true
```

This is preferable to manually checking only:

```go
r == ' '
```

because Unicode contains multiple whitespace characters.

---

## M. `unicode.IsSymbol`

```go
func IsSymbol(r rune) bool
```

Checks whether a rune belongs to the Unicode Symbol category.

Examples include various:

- mathematical symbols
- currency symbols
- modifier symbols
- other symbols

Example:

```go
fmt.Println(unicode.IsSymbol('$'))
fmt.Println(unicode.IsSymbol('©'))
fmt.Println(unicode.IsSymbol('∞'))
```

This is useful when filtering or classifying symbols in text.

---

## N. `unicode.IsTitle`

```go
func IsTitle(r rune) bool
```

Checks whether a rune has the Unicode **titlecase** property.

Example:

```go
fmt.Println(unicode.IsTitle('ǅ'))
```

Titlecase characters are a relatively specialized part of Unicode case handling.

You won't need this function in most beginner programs, but it becomes useful when implementing sophisticated Unicode-aware text processing.

---

## O. `unicode.IsUpper`

```go
func IsUpper(r rune) bool
```

Checks whether a rune is uppercase.

Example:

```go
fmt.Println(unicode.IsUpper('A'))
fmt.Println(unicode.IsUpper('a'))
```

Output:

```text
true
false
```

It works across Unicode scripts where uppercase/lowercase distinctions exist.

---

## P. `unicode.SimpleFold`

```go
func SimpleFold(r rune) rune
```

`SimpleFold` returns the next rune in the Unicode **simple case-folding cycle** for `r`.

For example, case-folding can group uppercase/lowercase variants of a character.

Conceptually:

```text
A → a → A → ...
```

For some Unicode characters, the cycle can contain more than two characters.

Example:

```go
r := 'A'

for i := 0; i < 3; i++ {
	fmt.Printf("%c\n", r)
	r = unicode.SimpleFold(r)
}
```

### Why is this useful?

It is useful for Unicode-aware case-insensitive processing.

However, don't assume that:

```go
unicode.SimpleFold(r)
```

means simply:

```text
uppercase → lowercase
```

It represents a **case-folding relationship**, which is a broader Unicode concept.

---

## Q. `unicode.To`

```go
func To(_case int, r rune) rune
```

Converts a rune to the requested Unicode case.

The first argument specifies the desired case:

```go
unicode.UpperCase
unicode.LowerCase
unicode.TitleCase
```

Example:

```go
r := 'g'

upper := unicode.To(unicode.UpperCase, r)

fmt.Printf("%c\n", upper)
```

Output:

```text
G
```

You can also do:

```go
unicode.To(unicode.LowerCase, 'G')
```

or:

```go
unicode.To(unicode.TitleCase, r)
```

### Why use `To`?

It is useful when the desired case is determined dynamically.

For example:

```go
caseType := unicode.UpperCase

result := unicode.To(caseType, 'g')
```

---

## R. `unicode.ToLower`

```go
func ToLower(r rune) rune
```

Converts a Unicode rune to lowercase when a lowercase mapping exists.

Example:

```go
r := unicode.ToLower('G')

fmt.Printf("%c\n", r)
```

Output:

```text
g
```

It is Unicode-aware, unlike manually manipulating ASCII values.

---

## S. `unicode.ToTitle`

```go
func ToTitle(r rune) rune
```

Converts a rune to its Unicode titlecase equivalent when applicable.

Example:

```go
r := unicode.ToTitle('ǆ')

fmt.Printf("%c\n", r)
```

Titlecase is particularly relevant for characters whose Unicode casing behavior doesn't map cleanly to the simple uppercase/lowercase distinction.

---

## T. `unicode.ToUpper`

```go
func ToUpper(r rune) rune
```

Converts a Unicode rune to uppercase.

Example:

```go
r := unicode.ToUpper('g')

fmt.Printf("%c\n", r)
```

Output:

```text
G
```

For multilingual applications, this is much safer than manually doing ASCII arithmetic.

---

# 4. Unicode `RangeTable` and related types

The package also provides types used to represent Unicode ranges.

## `RangeTable`

```go
type RangeTable struct {
	R16 []Range16
	R32 []Range32
	LatinOffset int
}
```

A `RangeTable` represents a collection of Unicode code-point ranges.

You normally don't construct these yourself.

Instead, the package provides predefined tables.

For example:

```go
unicode.Latin
unicode.Greek
unicode.Cyrillic
unicode.Han
```

Then:

```go
unicode.Is(unicode.Greek, 'Ω')
```

can determine whether the character belongs to the Greek range.

---

## `Range16`

```go
type Range16 struct {
	Lo     uint16
	Hi     uint16
	Stride uint16
}
```

Represents a range of Unicode code points that fit into 16 bits.

Conceptually:

```text
Lo       → starting code point
Hi       → ending code point
Stride   → distance between values
```

You generally encounter this type when working with Unicode tables rather than ordinary application code.

---

## `Range32`

```go
type Range32 struct {
	Lo     uint32
	Hi     uint32
	Stride uint32
}
```

Similar to `Range16`, but represents ranges requiring 32-bit values.

Unicode code points can go beyond the 16-bit range, so this type accommodates them.

---

## `CaseRange`

```go
type CaseRange struct {
	Lo    uint32
	Hi    uint32
	Delta d
}
```

`CaseRange` represents a range of Unicode code points involved in simple case conversion.

It is primarily an implementation/data-structure concern for Unicode case mappings.

---

## `SpecialCase`

```go
type SpecialCase []CaseRange
```

`SpecialCase` represents special Unicode casing rules.

It provides three methods:

```go
ToLower(r rune) rune
ToTitle(r rune) rune
ToUpper(r rune) rune
```

### `SpecialCase.ToLower`

```go
func (special SpecialCase) ToLower(r rune) rune
```

Applies the special-case lowercase mapping represented by the `SpecialCase`.

### `SpecialCase.ToTitle`

```go
func (special SpecialCase) ToTitle(r rune) rune
```

Applies special titlecase conversion.

### `SpecialCase.ToUpper`

```go
func (special SpecialCase) ToUpper(r rune) rune
```

Applies special uppercase conversion.

These are useful when implementing locale/special-case-aware casing rules using Unicode's case tables.

---

# 5. Important constants

The package also exposes constants such as:

```go
unicode.MaxRune
unicode.ReplacementChar
unicode.MaxASCII
unicode.MaxLatin1
```

The standard definitions are:

```text
MaxRune         = U+10FFFF
ReplacementChar = U+FFFD
MaxASCII        = U+007F
MaxLatin1       = U+00FF
```

For example:

```go
fmt.Printf("%U\n", unicode.MaxRune)
```

produces the maximum Unicode code point.

---

# 6. Complete practical example

Here's a small program that uses several `unicode` functions together:

```go
package main

import (
	"fmt"
	"unicode"
)

func main() {
	text := "Hello, 世界! 123"

	for _, r := range text {
		fmt.Printf("%c: ", r)

		switch {
		case unicode.IsLetter(r):
			fmt.Println("letter")

		case unicode.IsDigit(r):
			fmt.Println("digit")

		case unicode.IsSpace(r):
			fmt.Println("space")

		case unicode.IsPunct(r):
			fmt.Println("punctuation")

		case unicode.IsSymbol(r):
			fmt.Println("symbol")

		default:
			fmt.Println("other")
		}
	}
}
```

This is a realistic pattern: iterate over a Go string using `range`, obtain a `rune`, and then use `unicode` to classify it.

---

# 7. Three common beginner mistakes

## Mistake 1: Treating a string as a collection of characters

Beginners often write:

```go
for i := 0; i < len(text); i++ {
	fmt.Println(text[i])
}
```

This indexes **bytes**, not Unicode code points.

For Unicode-aware character processing, prefer:

```go
for _, r := range text {
	fmt.Println(r)
}
```

### Remember

```text
string
  ↓
UTF-8 bytes

range over string
  ↓
decoded runes
```

If you need low-level UTF-8 encoding/decoding, use `unicode/utf8`, which specifically provides UTF-8 operations.

---

## Mistake 2: Assuming `IsDigit` means only `0` through `9`

This assumption is ASCII-centric.

Unicode contains digits from multiple scripts.

Therefore:

```go
unicode.IsDigit(r)
```

is more appropriate when your application is supposed to recognize Unicode decimal digits.

Also remember the distinction:

```text
IsDigit → decimal digit

IsNumber → broader Unicode number category
```

---

## Mistake 3: Assuming Unicode case conversion is the same as changing ASCII values

Don't do things such as:

```go
r + 32
```

to convert a character to lowercase.

That only makes sense for limited ASCII assumptions.

Instead:

```go
unicode.ToLower(r)
```

and:

```go
unicode.ToUpper(r)
```

are Unicode-aware approaches.

Also remember that Unicode casing can be more complicated than a simple one-character uppercase/lowercase relationship.

---

# 8. Two real-world applications

## Application 1: Internationalized username validation

Suppose you're creating a system used globally.

You may want to allow letters from different writing systems:

```go
func validName(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsSpace(r) {
			return false
		}
	}

	return true
}
```

This can recognize letters beyond ASCII.

For example, names containing characters from different scripts can be processed without assuming English-only text.

---

## Application 2: Text processing and tokenization

A text-processing application might need to identify:

```text
letters
digits
spaces
punctuation
symbols
```

For example:

```go
for _, r := range text {
	switch {
	case unicode.IsLetter(r):
		// word character

	case unicode.IsDigit(r):
		// numeric character

	case unicode.IsSpace(r):
		// whitespace

	case unicode.IsPunct(r):
		// punctuation
	}
}
```

This can form part of:

- search engines
- document processors
- syntax/token analyzers
- text statistics
- natural-language processing pipelines

---

# 9. Three progressively challenging exercises

## Exercise 1 — Character Classifier

Write a Go program that accepts a string and counts how many runes are:

- letters
- digits
- spaces
- punctuation
- symbols

Use the appropriate `unicode` functions.

**Requirement:** The program must work with multilingual input such as:

```text
Hello 世界! १२३
```

Do not assume ASCII.

---

## Exercise 2 — Unicode-Aware Identifier Validator

Create a function:

```go
func isValidIdentifier(s string) bool
```

Your function should determine whether a string is a valid identifier according to your own rules.

For example, define rules such as:

- first rune must be a letter
- subsequent runes may be letters or digits
- spaces are forbidden
- punctuation is forbidden
- symbols are forbidden

Test it with identifiers containing characters from several writing systems.

**Challenge:** Decide what should happen when the input is empty.

---

## Exercise 3 — Multilingual Text Analyzer

Build a Unicode-aware text analyzer that accepts a paragraph and produces statistics such as:

- total rune count
- number of letters
- number of uppercase letters
- number of lowercase letters
- number of digits
- number of punctuation characters
- number of spaces
- number of symbols
- number of characters belonging to selected Unicode scripts

Then add a feature that produces a transformed version of the text using Unicode-aware case conversion.

**Challenge:** Test your program with text containing multiple scripts, combining marks, punctuation, emojis, and numbers, and investigate cases where the number of bytes differs significantly from the number of runes.

---

# 10. A useful mental model

When learning `unicode`, think about the package in four groups:

```text
                    unicode
                       │
       ┌───────────────┼────────────────┐
       │               │                │
   Classification   Case handling   Range checking
       │               │                │
       ├ IsLetter      ├ ToLower         ├ Is
       ├ IsDigit       ├ ToUpper         └ In
       ├ IsNumber      ├ ToTitle
       ├ IsSpace       ├ To
       ├ IsPunct       └ SimpleFold
       ├ IsSymbol
       ├ IsUpper
       ├ IsLower
       └ ...
```

The most important functions for a beginner are:

```go
unicode.IsLetter()
unicode.IsDigit()
unicode.IsNumber()
unicode.IsSpace()
unicode.IsPunct()
unicode.IsUpper()
unicode.IsLower()

unicode.ToUpper()
unicode.ToLower()
unicode.ToTitle()

unicode.Is()
unicode.In()
unicode.SimpleFold()
```

And remember the package relationship:

```text
unicode
   ↓
"What kind of Unicode character is this?"

unicode/utf8
   ↓
"How is this character represented in UTF-8 bytes?"
```

The `unicode` package provides Unicode code-point property data and functions, while `unicode/utf8` handles UTF-8 encoding/decoding.

---

# 🤔 Thought-provoking question

Suppose you are building a **global username system** and you decide that usernames can contain any Unicode letter, but you compare usernames using `unicode.ToLower()`.

**Do you think that is enough to guarantee that two visually or linguistically equivalent usernames are treated as the same username? Why or why not?**

Think particularly about:

- Unicode normalization
- combining characters
- case folding
- different scripts
- visually confusable characters

This is where Unicode processing becomes much more interesting than simply checking `IsLetter()`.
