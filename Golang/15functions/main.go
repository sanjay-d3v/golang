package main

import "fmt"

func main() {
	greeter()
	greeterTwo()
	fmt.Println("functions necessary from get go")

	result := adder(3, 5)

	fmt.Println("result is", result)

	proResult, myMessage := proAdder(2, 5, 8, 7)
	fmt.Println("pro result is: ", proResult)
	fmt.Println("my message is: ", myMessage)

}

func adder(ValOne int, ValTwo int) int {
	return ValOne + ValTwo
}

func greeterTwo() {
	fmt.Println("another method")
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}
	return total, "Go Lang Is Good"
}

func greeter() {
	fmt.Println("hey kali")
}
