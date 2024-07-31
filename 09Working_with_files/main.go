package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	fmt.Println("working with files")

	creating_file()
	Read_Entire_file_content()
	Delete_File()
	appending()

}

// Creating file
func creating_file() {
	content := "chandu is good boy"
	file, err := os.Create("./chandu.txt")
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("lenght is : ", length)
	defer file.Close()
}

// Read_Entire_file_content
func Read_Entire_file_content() {
	data, err := ioutil.ReadFile("./chandu.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
	//defer file.Close()
}

//Append

func appending() {
	file, err := os.OpenFile("chandu.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if _, err := file.WriteString("Appended line\n"); err != nil {
		log.Fatal(err)
	}
	defer file.Close()
}

//

//Deleting a file

func Delete_File() {
	err := os.Remove("chandu.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("deleted successfully")
}
