package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://loc.dev"

func main() {
	fmt.Println("LCO web requests")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("type pf response: %T", response)
	defer response.Body.Close() //responsibility to close

	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)
}
