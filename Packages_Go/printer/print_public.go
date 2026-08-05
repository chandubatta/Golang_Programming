package printer

import "fmt"

func Public_function(s string) {
	private_function(s)
}
func private_function(s string) {
	fmt.Println(s)
}
