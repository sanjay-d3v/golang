package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	var input 
	fmt.Println("welcome to time study")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	presentTime := time.Now()
	if presentTime == input {
		fmt.Println("It is today")
	}
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createDate := time.Date(2040, time.April, 04, 20, 0, 0, 0, time.UTC)
	fmt.Println(createDate)

	fmt.Println(createDate.Format("01-02-2006 Monday"))

}
