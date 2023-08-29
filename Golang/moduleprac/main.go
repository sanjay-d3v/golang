package main

import "fmt"

func main(name string) string {
	message := fmt.Sprintf("hi, %v. welcome!", name)
	return message
}
