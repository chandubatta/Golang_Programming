package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Handling Web Request")
	myurl := "https://chandubatta.netlify.app/"
	response, Err := http.Get(myurl)
	if Err != nil {
		panic(Err)
	}
	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(content)
	}
	res := string(content)
	fmt.Println(res)
}
