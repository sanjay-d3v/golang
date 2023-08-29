package main

import "fmt"

func main() {
	fmt.Println("if else in golang")

	var result string

	if loginCount := 4; loginCount < 10 {
		result = "regular user"
	} else if loginCount > 10 {
		result = "nonregular user"
	} else {
		result = "same user"
	}
	fmt.Println(result)
}
