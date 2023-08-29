package main

import "fmt"

const Logintoken string = "123"

func main() {
	var username string = "rupesh"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallval uint8 = 255
	fmt.Println(smallval)
	fmt.Printf("Variable is of type: %T \n", smallval)

	var usefloat float32 = 34.5678453627
	fmt.Println(usefloat)
	fmt.Printf("Variable is of type: %T \n", usefloat)

	var nostr string
	fmt.Println(nostr)
	fmt.Printf("Variable is of type: %T \n", nostr)

	//implicit type

	var website = "gogoanime.com"
	fmt.Println(website)

	randomnumber := 8956743
	fmt.Println(randomnumber)
	fmt.Printf("Variable is of type: %T \n", randomnumber)

	fmt.Println(Logintoken)
	fmt.Printf("Variable is of type: %T \n", Logintoken)
}
