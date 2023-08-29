package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("imp slices")

	var fruitList = []string{"apple", "tomato", "peach"}
	fmt.Printf("type of fruitList %T\n", fruitList)

	fruitList = append(fruitList, "dragonfruit", "custardapple")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScore := make([]int, 4)

	highScore[0] = 234
	highScore[1] = 945
	highScore[2] = 467
	highScore[3] = 453
	//highScore[4] = 567

	highScore = append(highScore, 444, 555, 666)

	fmt.Println("high scores", highScore)

	sort.Ints(highScore)
	fmt.Println(highScore)

	//how to remove values from slices based on index

	var courses = []string{"reactjs", "javascript", "pyhton", "ruby", "terra"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
