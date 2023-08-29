package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("switch case explored")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("value of dice: ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("dice value is 1 and you can open")
	case 2:
		fmt.Println("you can move 2 spots")
		fallthrough
	case 3:
		fmt.Println("you can move 3 spots")
	case 4:
		fmt.Println("you can move 4 spots")
	case 5:
		fmt.Println("you can move 5 spots")
	case 6:
		fmt.Println("you can move 6 spots and roll again")
	default:
		fmt.Println("what was that")
	}
}
