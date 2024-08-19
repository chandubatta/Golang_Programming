package main

import "fmt"

func main() {
	fmt.Println("Variables and constant")
	var name string = "chandu"
	var age = 24
	married_status := false
	const quality = "good"
	fmt.Printf("variable type is %T and value is %v : \n", name, name)
	fmt.Printf("variable type is %T and value is %v \n", age, age)
	fmt.Printf("variable type is %T and value is %v \n:", married_status, married_status)
	fmt.Printf("variable type is %T and value is %v \n:", quality, quality)

}
