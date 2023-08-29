package main

import "fmt"

func main() {
	fmt.Println("maps assign")

	languages := make(map[string]string)

	languages["35"] = "javascript"
	languages["45"] = "typescript"
	languages["py"] = "python"
	languages["go"] = "golang"

	fmt.Println("languages", languages)

	delete(languages, "py")
	fmt.Println(languages)

	//loops are intersting in golang

	for _, value := range languages {
		fmt.Printf("%v\n", key, value)
	}

}
