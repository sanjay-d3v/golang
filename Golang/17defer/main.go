package main

import "fmt"

func main() {
	fmt.Println("defer LIFO")
	defer fmt.Println("world")
	//defer used for last execution
	myDefer()
	defer fmt.Println("one")
	defer fmt.Println("two")
	fmt.Println("hello")

}
func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
