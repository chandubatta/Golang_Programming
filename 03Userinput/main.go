package main

import (
	"fmt"
)

func main() {
	var (
		intValue     int
		uintValue    uint
		floatValue   float64
		boolValue    bool
		complexValue complex128
		stringValue  string
	)

	fmt.Println("Enter an integer value:")
	if _, err := fmt.Scan(&intValue); err != nil {
		fmt.Println("Error scanning int:", err)
		return
	}

	fmt.Println("Enter an unsigned integer value:")
	if _, err := fmt.Scan(&uintValue); err != nil {
		fmt.Println("Error scanning uint:", err)
		return
	}

	fmt.Println("Enter a float value:")
	if _, err := fmt.Scan(&floatValue); err != nil {
		fmt.Println("Error scanning float:", err)
		return
	}

	fmt.Println("Enter a boolean value (true/false):")
	if _, err := fmt.Scan(&boolValue); err != nil {
		fmt.Println("Error scanning bool:", err)
		return
	}

	fmt.Println("Enter a complex value (e.g., 1.23+4.56i):")
	if _, err := fmt.Scan(&complexValue); err != nil {
		fmt.Println("Error scanning complex:", err)
		return
	}

	fmt.Println("Enter a string value:")
	if _, err := fmt.Scan(&stringValue); err != nil {
		fmt.Println("Error scanning string:", err)
		return
	}

	fmt.Printf("Parsed integer value: %d\n", intValue)
	fmt.Printf("Parsed unsigned integer value: %d\n", uintValue)
	fmt.Printf("Parsed float value: %f\n", floatValue)
	fmt.Printf("Parsed boolean value: %t\n", boolValue)
	fmt.Printf("Parsed complex value: %v\n", complexValue)
	fmt.Printf("Parsed string value: %s\n", stringValue)
}
