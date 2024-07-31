package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Taking String from the User
	fmt.Println("User Input in Golang")
	fmt.Println("Enter User Name")
	reder := bufio.NewReader(os.Stdin)
	UserName, err := reder.ReadString('\n')
	if err != nil {
		panic(err)
	}
	fmt.Printf("User name is %v and Data Type is %T\n", UserName, UserName)

	//Taking Integer and Float from the user
	fmt.Println("Enter User Height ")
	UserHeight, err := reder.ReadString('\n')
	if err != nil {
		panic(err)
	}
	Userheight, err := strconv.ParseFloat(strings.TrimSpace(UserHeight), 64)
	if err != nil {
		panic(err)
	}
	fmt.Printf("The height is %v and Data Type is %T \n", Userheight, Userheight)

	fmt.Println("Enter User age")
	age, err := reder.ReadString('\n')
	if err != nil {
		panic(err)
	}
	Userage, err := strconv.Atoi(strings.TrimSpace(age))
	if err != nil {
		panic(err)
	}
	fmt.Printf("User age is %v and Data type is %T \n", Userage, Userage)

	//Taking boolean value from the user
	fmt.Println("Enter the married status: ")
	status, err := reder.ReadString('\n')
	if err != nil {
		panic(err)
	}
	MarridStatus, err := strconv.ParseBool(strings.TrimSpace(status))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Married status is %v and Data type is %T \n", MarridStatus, MarridStatus)


	// Using Scan from the fmt package f
}
