package main

import (
	"fmt"
	"packages/printer"

	"github.com/fatih/color"
)

func main() {
	fmt.Println("hi")
	printer.Public_function("chandu")
	s := "chandu"
	color.Red(s)
}
