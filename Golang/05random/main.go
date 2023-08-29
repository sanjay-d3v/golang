package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("welcome to my code")

	//var myNumberOne int = 2
	//var myNumberTwo float64 = 4
	//fmt.Println("The sum is:", myNumberOne+int(myNumberTwo))

	//random number
	//rand.Seed(time.Now().UnixNano())
	//fmt.Println(rand.Intn(5))

	//crypto random
	myRandNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandNum)

}
