package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	//String declaration
	var s1 string
	fmt.Println(s1)

	var s2 string = "hello"
	var s3 = "hi"
	s4 := "chandu"
	fmt.Println(s1, s2, s3, s4)

	//length of the string
	l1 := len(s1) //0
	fmt.Println(l1)

	var s5 = "héllo"
	l2 := len(s5)
	fmt.Println(l2)

	//Comparing strings
	s6 := s5 == s4
	fmt.Println(s6)
	v1 := strings.Compare(s4, s5)
	fmt.Println(v1)

	v2 := strings.EqualFold(s4, s5)
	fmt.Println(v2)

	v3 := s1 > s2
	v4 := s2 < s4
	fmt.Println(v3, v4)

	//Concatination
	s8 := s4 + "" + s5
	fmt.Println(s8)

	//string builder
	b := strings.Builder{}
	b.Grow(1024)
	b.WriteString("hi")
	b.WriteString(" ")
	b.WriteString("chandu")
	//b.String()
	fmt.Println(b.String())

	//Accessing index
	s9 := "chandu"
	fmt.Println(string(s9[5]))

	//substring
	s10 := s9[:4]
	fmt.Println(s10)

	//string convert
	number := 123
	str := strconv.Itoa(number)
	fmt.Printf("%T %v", str, str)

	fmt.Println()
	s11 := "123"
	number1, err := strconv.Atoi(s11)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T %v", number1, number1)

	//looing or iteration
	fmt.Println()

	s12 := "chandu"

	for i := 0; i < len(s12); i++ {
		fmt.Print(s12[i])
		fmt.Println()
		fmt.Print(string(s12[i]))
	}

	for i, v := range s12 {
		fmt.Println(i, v)
		fmt.Printf("%T", v)
	}

}
