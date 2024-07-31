package main

import (
	"fmt"
)

func main() {
	// Using new to allocate memory for an int
	intPointer := new(int)
	fmt.Println("Value of intPointer before assignment:", *intPointer)
	*intPointer = 10
	fmt.Println("Value of intPointer after assignment:", *intPointer)

	// Using new to allocate memory for a struct
	type Car struct {
		Brand string
		Year  int
	}

	carPointer := new(Car)
	fmt.Println("Car struct before assignment:", *carPointer)
	carPointer.Brand = "Toyota"
	carPointer.Year = 2020
	fmt.Println("Car struct after assignment:", *carPointer)

	// Using make to create and initialize a slice
	slice := make([]int, 3)
	fmt.Println("Slice before assignment:", slice)
	slice[0] = 1
	slice[1] = 2
	slice[2] = 3
	fmt.Println("Slice after assignment:", slice)

	// Using make to create and initialize a map
	m := make(map[string]int)
	fmt.Println("Map before assignment:", m)
	m["a"] = 1
	m["b"] = 2
	fmt.Println("Map after assignment:", m)

	// Using make to create and initialize a channel
	ch := make(chan string)
	go func() {
		ch <- "Hello, Go!"
	}()
	fmt.Println("Channel received value:", <-ch)
}
