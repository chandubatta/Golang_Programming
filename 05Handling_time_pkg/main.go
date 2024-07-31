package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Handling time package in Golang")

	//getting a present time

	presentime := time.Now()
	fmt.Printf("Present time is %v and Data type is %T \n", presentime, presentime)

	//formating a time

	fmt.Println(presentime.Format("01-02-2006 15:04:05 Monday "))

	//passing a time

	layout := "2006-01-02 15:04:05"
	str := "2023-07-26 14:20:00"
	t, err := time.Parse(layout, str)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return
	}
	fmt.Println("Parsed time:", t)
}
