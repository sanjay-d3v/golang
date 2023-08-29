package main

import "fmt"

func main() {
	fmt.Println("arrays")

	var fruitList [4]string

	fruitList[0] = "apple"
	fruitList[1] = "banana"
	//fruitList[2] = "mango"
	fruitList[3] = "peach"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("lemgth is: ", len(fruitList))
}
