package main

import "fmt"

func main() {
	fmt.Println("Variables in Golang")

	// using var key word
	var name string
	name = "chandu"
	fmt.Printf("Name is %v and Type is %T\n ", name, name)
	var age int = 24
	fmt.Printf("Age is %v and Type is %T \n", age, age)

	//implicit type ( with out using of data type) LEXER done the job
	var married = false
	fmt.Printf("married %v and type is %T\n", married, married)

	// using valurs operator and with out using var key word
	height := 5.2
	weight := 61
	fmt.Printf("Height is %v and Type is %T\n", height, height)
	fmt.Printf("Weight is %v and Type is %T\n", weight, weight)
}
