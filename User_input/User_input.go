package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to User input in the golang")
	var (
		intvalue     int
		uintvalue    uint
		floatvalue   float32
		stringvalue  string
		boolvalue    bool
		complexvalue complex64
	)
	fmt.Println("Enter interger value")
	if _, err := fmt.Scan(&intvalue); err != nil {
		fmt.Println("Error while scaning int value", err)
		return
	}
	fmt.Println("Enter unsigned interger value")
	if _, err := fmt.Scan(&uintvalue); err != nil {
		fmt.Println("Error while scaning the uint vale", err)
		return
	}
	fmt.Println("Enter float value")
	if _, err := fmt.Scan(&floatvalue); err != nil {
		fmt.Println("error while scaning the float value", err)
		return
	}
	fmt.Println("Enter string value")
	if _, err := fmt.Scan(&stringvalue); err != nil {
		fmt.Println("Error while scaning the string value")
		return
	}
	fmt.Println("Enter boolean value")
	if _, err := fmt.Scan(&boolvalue); err != nil {
		fmt.Println("Error while scaning the boolean value", err)
		return
	}
	fmt.Println("Enter complex value")
	if _, err := fmt.Scan(&complexvalue); err != nil {
		fmt.Println("Error while scanning the complex value", err)
		return
	}

	fmt.Printf(" Passed value is: %v and type is: %T \n", intvalue, intvalue)
	fmt.Printf(" Passed value is: %v and type is: %T \n", uintvalue, uintvalue)
	fmt.Printf(" Passed value is: %v and type is: %T \n", floatvalue, floatvalue)
	fmt.Printf(" Passed value is: %v and type is: %T \n", stringvalue, stringvalue)
	fmt.Printf(" Passed value is: %v and type is: %T \n", boolvalue, boolvalue)
	fmt.Printf(" Passed value is: %v and type is: %T \n", complexvalue, complexvalue)
}
