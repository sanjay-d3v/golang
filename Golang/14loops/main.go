package main

import "fmt"

func main() {
	fmt.Println("welcome to loops")

	days := []string{"sunday", "monday", "tuesday", "friday", "saturday"}

	fmt.Println(days)

	//for d := 0; d < len(days); d++ {
	//fmt.Println(days[d])
	//}

	//for i := range days {
	//fmt.Println(days[i])
	//}

	for index, day := range days {
		fmt.Printf("index is %v and value is%v\n", index, day)
	}

}
