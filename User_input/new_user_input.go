package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welcome to our piza app")
	r := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the Rating :")

	//comma ok or // err ok syntax
	ratting, _ := r.ReadString('\n')
	fmt.Printf("Thanks for rating %v and type %T \n", ratting, ratting)
	//fmt.Printf("Type %T", ratting)

	//Type Conversion from one data type to another
	// As of now we read the String // converst into float
	Convert_float, _ := strconv.ParseFloat(strings.TrimSpace(ratting), 64)
	fmt.Printf("Thanks for rating %v and type %T \n", Convert_float, Convert_float)
	//convert into int
	//Convert_int, _ := strconv.ParseInt(strings.TrimSpace(ratting),32)
	//fmt.Printf("Thanks for rating %v and type %T \n", Convert_int, Convert_int)

}
