package main

import "fmt"

func main() {
	fmt.Println("Pointers in Golang")
	var ptr *int
	fmt.Println(ptr)
	myNumber := 23
	ptr = &myNumber
	fmt.Println(ptr)

}
