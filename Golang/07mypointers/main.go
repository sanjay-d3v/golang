package main

import "fmt"

func main() {
	fmt.Println("welcome to pointers")

	//var myptr *int

	//fmt.Println("my pointer ", myptr)

	myNumber := 23

	var ptr = &myNumber

	fmt.Println("my pointer is", ptr)
	fmt.Println("my pointer is", *ptr)

}
