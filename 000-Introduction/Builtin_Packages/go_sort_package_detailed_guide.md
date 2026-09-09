# Go `sort` Package — Detailed Learning Guide

## 1. What is the `sort` package?

The Go `sort` package provides functions and types for sorting slices and custom collections, checking whether data is sorted, and performing binary searches.

It is part of Go's standard library:

```go
import "sort"
```

It is commonly used when you need to:

- Sort `int`, `float64`, or `string` slices.
- Sort structs using custom comparison logic.
- Sort data in ascending or descending order.
- Perform stable sorting.
- Check whether data is already sorted.
- Search sorted data using binary search.
- Build custom sortable collections with `sort.Interface`.

Sorting normally happens **in-place**, so the original slice is modified.

---

## 2. Simple example

```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{50, 10, 40, 20, 30}

	fmt.Println("Before:", numbers)

	sort.Ints(numbers)

	fmt.Println("After:", numbers)
}
```

Output:

```text
Before: [50 10 40 20 30]
After: [10 20 30 40 50]
```

The important line is:

```go
sort.Ints(numbers)
```

It sorts the existing slice in ascending order.

---

## 3. Important concept: `sort` works in-place

Consider:

```go
numbers := []int{5, 2, 8, 1}

sort.Ints(numbers)

fmt.Println(numbers)
```

The slice itself changes to:

```text
[1 2 5 8]
```

You generally should not write:

```go
sorted := sort.Ints(numbers) // Incorrect
```

because `sort.Ints` does not return the sorted slice.

Instead:

```go
sort.Ints(numbers)
```

---

# 4. Functions in the `sort` package

The package provides functions for sorting, checking sortedness, searching, and working with custom sorting implementations.

Important exported functions include:

1. `Find`
2. `Float64s`
3. `Float64sAreSorted`
4. `Ints`
5. `IntsAreSorted`
6. `IsSorted`
7. `Search`
8. `SearchFloat64s`
9. `SearchInts`
10. `SearchStrings`
11. `Slice`
12. `SliceIsSorted`
13. `SliceStable`
14. `Sort`
15. `Stable`
16. `Strings`
17. `StringsAreSorted`
18. `Reverse`

The package also provides important types such as:

- `Interface`
- `IntSlice`
- `StringSlice`
- `Float64Slice`

---

# 5. `sort.Ints()`

### Syntax

```go
sort.Ints(x []int)
```

Sorts an integer slice in increasing/ascending order.

### Example

```go
numbers := []int{40, 10, 30, 20}

sort.Ints(numbers)

fmt.Println(numbers)
```

Output:

```text
[10 20 30 40]
```

### When to use

Use it when you simply need to sort a `[]int` in ascending order.

---

# 6. `sort.IntsAreSorted()`

### Syntax

```go
sort.IntsAreSorted(x []int) bool
```

Checks whether an integer slice is already sorted in ascending order.

Example:

```go
numbers := []int{10, 20, 30, 40}

if sort.IntsAreSorted(numbers) {
	fmt.Println("Already sorted")
}
```

Output:

```text
Already sorted
```

For:

```go
[]int{10, 30, 20, 40}
```

the function returns:

```text
false
```

---

# 7. `sort.Float64s()`

### Syntax

```go
sort.Float64s(x []float64)
```

Sorts floating-point numbers in increasing order.

Example:

```go
prices := []float64{99.5, 10.2, 50.8, 25.3}

sort.Float64s(prices)

fmt.Println(prices)
```

Output:

```text
[10.2 25.3 50.8 99.5]
```

An important detail is that `NaN` values are placed before other floating-point values according to the package's ordering rules.

---

# 8. `sort.Float64sAreSorted()`

### Syntax

```go
sort.Float64sAreSorted(x []float64) bool
```

Checks whether a `[]float64` is sorted.

Example:

```go
numbers := []float64{1.1, 2.2, 3.3}

fmt.Println(sort.Float64sAreSorted(numbers))
```

Output:

```text
true
```

It follows the package's special ordering rules for floating-point values, including `NaN`.

---

# 9. `sort.Strings()`

### Syntax

```go
sort.Strings(x []string)
```

Sorts strings in increasing lexicographical order.

Example:

```go
names := []string{
	"Charlie",
	"Alice",
	"Bob",
}

sort.Strings(names)

fmt.Println(names)
```

Output:

```text
[Alice Bob Charlie]
```

---

# 10. `sort.StringsAreSorted()`

### Syntax

```go
sort.StringsAreSorted(x []string) bool
```

Checks whether strings are already sorted.

Example:

```go
names := []string{
	"Alice",
	"Bob",
	"Charlie",
}

fmt.Println(sort.StringsAreSorted(names))
```

Output:

```text
true
```

---

# 11. `sort.Slice()`

This is one of the most important functions for beginners to learn.

### Syntax

```go
sort.Slice(x any, less func(i, j int) bool)
```

It allows you to sort a slice according to your own comparison rule.

Example:

```go
package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	people := []Person{
		{"Alice", 30},
		{"Bob", 20},
		{"Charlie", 25},
	}

	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})

	fmt.Println(people)
}
```

Output:

```text
[{Bob 20} {Charlie 25} {Alice 30}]
```

The comparison function:

```go
func(i, j int) bool {
	return people[i].Age < people[j].Age
}
```

means:

> Put element `i` before element `j` when its age is smaller.

### Descending order

Reverse the comparison:

```go
sort.Slice(people, func(i, j int) bool {
	return people[i].Age > people[j].Age
})
```

---

# 12. `sort.SliceStable()`

### Syntax

```go
sort.SliceStable(x any, less func(i, j int) bool)
```

It is similar to `sort.Slice`, but performs a **stable sort**.

A stable sort preserves the original relative order of elements that compare equal.

Example:

```go
type Student struct {
	Name  string
	Class int
}

students := []Student{
	{"Alice", 10},
	{"Bob", 10},
	{"Charlie", 9},
	{"David", 10},
}

sort.SliceStable(students, func(i, j int) bool {
	return students[i].Class < students[j].Class
})
```

The class-10 students remain in their original relative order:

```text
Alice
Bob
David
```

Stable sorting is especially useful for multi-level sorting.

---

# 13. `sort.SliceIsSorted()`

### Syntax

```go
sort.SliceIsSorted(x any, less func(i, j int) bool) bool
```

Checks whether a custom slice is sorted according to your comparison function.

Example:

```go
people := []Person{
	{"Alice", 20},
	{"Bob", 25},
	{"Charlie", 30},
}

result := sort.SliceIsSorted(people, func(i, j int) bool {
	return people[i].Age < people[j].Age
})

fmt.Println(result)
```

Output:

```text
true
```

---

# 14. `sort.Sort()`

### Syntax

```go
sort.Sort(data sort.Interface)
```

It sorts data according to the `Less` method of a type implementing `sort.Interface`.

The sort is **not stable**.

`sort.Interface` requires three methods:

```go
Len()
Less(i, j int)
Swap(i, j int)
```

Example:

```go
type People []Person

func (p People) Len() int {
	return len(p)
}

func (p People) Less(i, j int) bool {
	return p[i].Age < p[j].Age
}

func (p People) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
```

Then:

```go
people := People{
	{"Alice", 30},
	{"Bob", 20},
	{"Charlie", 25},
}

sort.Sort(people)
```

This is more verbose than `sort.Slice`, but it is useful for understanding Go's sorting abstraction.

---

# 15. `sort.Stable()`

### Syntax

```go
sort.Stable(data sort.Interface)
```

Sorts using `sort.Interface` while preserving the relative order of equal elements.

You implement:

```go
Len()
Less()
Swap()
```

just like with `sort.Sort`.

### Difference

```go
sort.Sort(...)
```

is unstable.

```go
sort.Stable(...)
```

is stable.

---

# 16. `sort.Reverse()`

### Syntax

```go
sort.Reverse(data sort.Interface) sort.Interface
```

It creates a sorting interface that reverses the ordering.

Example:

```go
numbers := []int{5, 2, 8, 1, 3}

sort.Sort(sort.Reverse(sort.IntSlice(numbers)))

fmt.Println(numbers)
```

Output:

```text
[8 5 3 2 1]
```

For modern custom sorting, `sort.Slice` can often express descending order more directly:

```go
sort.Slice(numbers, func(i, j int) bool {
	return numbers[i] > numbers[j]
})
```

---

# 17. `sort.Search()`

This function performs a **binary search**.

### Syntax

```go
sort.Search(n int, f func(int) bool) int
```

Suppose:

```go
numbers := []int{10, 20, 30, 40, 50}
```

You want to find where `30` is:

```go
index := sort.Search(len(numbers), func(i int) bool {
	return numbers[i] >= 30
})

fmt.Println(index)
```

Output:

```text
2
```

The function returns the **smallest index for which the condition becomes true**.

For:

```text
10 20 30 40 50
```

the condition:

```go
numbers[i] >= 30
```

first becomes true at index `2`.

### Important

The condition supplied to `sort.Search` must have the monotonic property required for binary search. The input is normally sorted in a way that makes the predicate transition from false to true.

---

# 18. `sort.SearchInts()`

### Syntax

```go
sort.SearchInts(a []int, x int) int
```

Searches for an integer in a sorted integer slice using binary search.

Example:

```go
numbers := []int{10, 20, 30, 40, 50}

index := sort.SearchInts(numbers, 30)

fmt.Println(index)
```

Output:

```text
2
```

If the value does not exist:

```go
index := sort.SearchInts(numbers, 35)
```

the result is:

```text
3
```

Why?

Because `35` could be inserted at index `3`:

```text
10 20 30 35 40 50
         ↑
```

Therefore, don't automatically interpret the result as "found."

Correct existence checking:

```go
index := sort.SearchInts(numbers, 35)

if index < len(numbers) && numbers[index] == 35 {
	fmt.Println("Found")
} else {
	fmt.Println("Not found")
}
```

---

# 19. `sort.SearchFloat64s()`

### Syntax

```go
sort.SearchFloat64s(a []float64, x float64) int
```

Performs binary search on a sorted `[]float64`.

Example:

```go
numbers := []float64{
	1.1,
	2.2,
	3.3,
	4.4,
}

index := sort.SearchFloat64s(numbers, 3.3)

fmt.Println(index)
```

Output:

```text
2
```

---

# 20. `sort.SearchStrings()`

### Syntax

```go
sort.SearchStrings(a []string, x string) int
```

Performs binary search on a sorted string slice.

Example:

```go
names := []string{
	"Alice",
	"Bob",
	"Charlie",
	"David",
}

index := sort.SearchStrings(names, "Charlie")

fmt.Println(index)
```

Output:

```text
2
```

The input must be sorted in ascending order.

---

# 21. `sort.Find()`

`sort.Find` is a binary-search API that uses a three-way comparison result.

### Syntax

```go
sort.Find(n int, cmp func(int) int) (int, bool)
```

The comparison function returns:

```text
negative → target is after this position
0        → match
positive → target is before this position
```

Example:

```go
numbers := []int{10, 20, 30, 40, 50}

target := 30

index, found := sort.Find(len(numbers), func(i int) int {
	if numbers[i] < target {
		return -1
	}
	if numbers[i] > target {
		return 1
	}
	return 0
})

fmt.Println(index, found)
```

Output:

```text
2 true
```

`Find` returns the appropriate boundary index and a boolean indicating whether the target was actually found.

It is useful when a three-way comparison is more natural than a simple true/false predicate.

---

# 22. `sort.IsSorted()`

### Syntax

```go
sort.IsSorted(data sort.Interface) bool
```

Checks whether an object implementing `sort.Interface` is sorted.

Example:

```go
numbers := []int{1, 2, 3, 4}

fmt.Println(sort.IsSorted(sort.IntSlice(numbers)))
```

Output:

```text
true
```

It is the generic counterpart to functions such as:

```go
sort.IntsAreSorted()
```

---

# 23. `sort.Interface`

Although it is not a function, `sort.Interface` is one of the most important parts of the package.

The interface is:

```go
type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}
```

It lets you define how a custom collection can be sorted.

Think of it as telling Go:

> "Here is my collection. I will tell you how many elements it contains, how two elements should be compared, and how two elements can be exchanged."

### `Len()`

Returns the number of elements:

```go
func (p People) Len() int {
	return len(p)
}
```

### `Less(i, j)`

Returns `true` when element `i` should appear before element `j`:

```go
func (p People) Less(i, j int) bool {
	return p[i].Age < p[j].Age
}
```

### `Swap(i, j)`

Exchanges two elements:

```go
func (p People) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
```

These three methods form the foundation of the traditional `sort.Sort` API.

Your `Less` function must define a consistent ordering.

---

# 24. `sort.IntSlice`

`IntSlice` is a type representing a slice of integers:

```go
type IntSlice []int
```

It implements `sort.Interface`.

Example:

```go
numbers := []int{5, 2, 8, 1}

sort.Sort(sort.IntSlice(numbers))
```

`IntSlice` provides methods such as:

```text
Len()
Less()
Swap()
Search()
Sort()
```

### `IntSlice.Sort()`

A convenience method:

```go
sort.IntSlice(numbers).Sort()
```

### `IntSlice.Search()`

Searches the sorted integer slice:

```go
index := sort.IntSlice(numbers).Search(5)
```

---

# 25. `sort.StringSlice`

`StringSlice` is the string equivalent:

```go
type StringSlice []string
```

It provides methods such as:

```text
Len()
Less()
Swap()
Search()
Sort()
```

Example:

```go
names := sort.StringSlice{
	"Charlie",
	"Alice",
	"Bob",
}

names.Sort()

fmt.Println(names)
```

Output:

```text
[Alice Bob Charlie]
```

You can also use:

```go
index := names.Search("Bob")
```

to perform binary search on the sorted slice.

---

# 26. `sort.Float64Slice`

This is the floating-point equivalent:

```go
type Float64Slice []float64
```

It provides:

```text
Len()
Less()
Swap()
Search()
Sort()
```

Example:

```go
numbers := sort.Float64Slice{
	3.3,
	1.1,
	2.2,
}

numbers.Sort()

fmt.Println(numbers)
```

Output:

```text
[1.1 2.2 3.3]
```

`Float64Slice.Less` also handles `NaN` specially so that `NaN` values have a defined position in the ordering.

---

# 27. Sorting structs

One of the most practical uses of `sort` is sorting structs.

Suppose:

```go
type Product struct {
	Name  string
	Price float64
}
```

Sort by price:

```go
products := []Product{
	{"Laptop", 800},
	{"Mouse", 20},
	{"Keyboard", 50},
}

sort.Slice(products, func(i, j int) bool {
	return products[i].Price < products[j].Price
})
```

Result:

```text
Mouse     20
Keyboard  50
Laptop    800
```

For descending order:

```go
sort.Slice(products, func(i, j int) bool {
	return products[i].Price > products[j].Price
})
```

---

# 28. Multi-field sorting

Suppose:

```go
type Employee struct {
	Name       string
	Department string
	Salary     int
}
```

You want:

1. Department ascending.
2. Salary descending.
3. Name ascending.

You can write:

```go
sort.Slice(employees, func(i, j int) bool {
	if employees[i].Department != employees[j].Department {
		return employees[i].Department < employees[j].Department
	}

	if employees[i].Salary != employees[j].Salary {
		return employees[i].Salary > employees[j].Salary
	}

	return employees[i].Name < employees[j].Name
})
```

This pattern is extremely useful in real applications.

---

# 29. Stable vs. unstable sorting

Suppose you have:

```text
Alice   Engineering
Bob     Sales
Charlie Engineering
David   Sales
```

You sort by department.

An **unstable** sort is allowed to rearrange elements that compare equal.

A **stable** sort preserves their original relative order.

For Engineering:

```text
Alice
Charlie
```

their relative order remains the same with a stable sort.

Use:

```go
sort.SliceStable(...)
```

when preserving equal-element order matters.

Use:

```go
sort.Slice(...)
```

when you don't care about the order of equal elements.

---

# 30. Three common beginner mistakes

## Mistake 1: Forgetting that sorting modifies the slice

Beginners sometimes expect:

```go
sorted := sort.Ints(numbers)
```

This is incorrect because `sort.Ints` does not return a slice.

Use:

```go
sort.Ints(numbers)
```

Remember: sorting is generally in-place.

---

## Mistake 2: Assuming `SearchInts` tells you whether an item exists

This is wrong:

```go
index := sort.SearchInts(numbers, 50)

if index != -1 {
	fmt.Println("Found")
}
```

`SearchInts` does not return `-1` for "not found."

It returns the index where the value exists or where it could be inserted.

Correct:

```go
index := sort.SearchInts(numbers, 50)

if index < len(numbers) && numbers[index] == 50 {
	fmt.Println("Found")
}
```

---

## Mistake 3: Writing an inconsistent `Less` function

For example:

```go
sort.Slice(numbers, func(i, j int) bool {
	return numbers[i] <= numbers[j]
})
```

Prefer:

```go
return numbers[i] < numbers[j]
```

The `Less` function should define a consistent strict ordering.

---

# 31. Two real-world applications

## Application 1: E-commerce product sorting

An online store may allow users to sort products by:

- Price: low → high
- Price: high → low
- Rating
- Name
- Popularity
- Discount percentage

For example:

```go
sort.Slice(products, func(i, j int) bool {
	return products[i].Price < products[j].Price
})
```

---

## Application 2: Leaderboards and rankings

Imagine a gaming application:

```go
type Player struct {
	Name  string
	Score int
}
```

Sort players from highest score to lowest:

```go
sort.Slice(players, func(i, j int) bool {
	return players[i].Score > players[j].Score
})
```

You could display:

```text
1. Alice   9500
2. Bob     8200
3. Charlie 7100
```

The same concept applies to:

- Sports rankings
- Employee performance reports
- Student rankings
- Financial reports
- Search results

---

# 32. Three progressively challenging exercises

## Exercise 1 — Basic sorting

Create a Go program that:

1. Creates a slice containing 10 integers in random order.
2. Prints the original slice.
3. Sorts the slice in ascending order using the `sort` package.
4. Prints the sorted slice.
5. Checks whether the slice is sorted.
6. Prints the result.

Do not use a manual sorting algorithm such as bubble sort.

---

## Exercise 2 — Sort a slice of structs

Create:

```go
type Product struct {
	Name   string
	Price  float64
	Rating float64
}
```

Create at least 8 products.

Your program should allow the products to be sorted:

1. By price from low to high.
2. By price from high to low.
3. By rating from high to low.

Use `sort.Slice`.

Challenge: If two products have the same rating, order those products alphabetically by name.

---

## Exercise 3 — Leaderboard with binary search

Create:

```go
type Player struct {
	Name  string
	Score int
}
```

Build a leaderboard containing at least 15 players.

Your program must:

1. Sort players by score from highest to lowest.
2. Assign each player a ranking.
3. Display the leaderboard.
4. Allow a player name to be searched.
5. Determine the player's score.
6. Maintain a separate sorted collection of scores.
7. Use an appropriate `sort` binary-search function to determine where a new score belongs.
8. Correctly handle duplicate scores.

Advanced challenge: Make your ranking logic correctly handle ties—for example, two players with the same score should receive the same rank.

---

# 33. Modern Go note

If you're learning Go today, don't assume that `sort` is always the newest or preferred solution.

Modern Go also has the generic `slices` package. In many situations, APIs such as:

```go
slices.Sort
slices.SortFunc
slices.SortStableFunc
slices.IsSorted
slices.IsSortedFunc
```

are more ergonomic and can be faster.

For example, instead of:

```go
sort.Ints(numbers)
```

modern Go can use:

```go
slices.Sort(numbers)
```

However, learning `sort` is still worthwhile because `sort.Interface`, custom comparisons, stable sorting, and binary search teach important concepts that are useful throughout Go programming.

---

# 34. Recommended learning order

A good sequence for learning the package is:

```text
sort.Ints
     ↓
sort.Strings
     ↓
sort.Float64s
     ↓
sort.Slice
     ↓
sort.SliceStable
     ↓
sort.SliceIsSorted
     ↓
sort.Sort
     ↓
sort.Interface
     ↓
sort.Reverse
     ↓
sort.SearchInts / SearchStrings
     ↓
sort.Search
     ↓
sort.Find
```

---

# 35. Thought-provoking question

Suppose you are building a **large e-commerce search system with 10 million products**.

You need to let users sort by:

- Price
- Rating
- Popularity
- Relevance

You also need to quickly determine where a newly arriving product belongs in a sorted list.

**Would you simply sort the entire 10-million-item slice every time a user changes the sorting option? Why or why not? What data structures or algorithms might you combine with Go's `sort`/`slices` functionality to make this system efficient?**

This question takes you beyond simply using `sort` and into:

- Algorithmic complexity
- Data structures
- Binary search
- Caching
- Indexing
- Precomputed ordering
- System design
- Performance optimization
