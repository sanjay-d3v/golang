package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("welcome to web verb")
	PerfromGetRequest()
}

func PerfromGetRequest() {
	const myUrl = "http://localhost:8000/get"

	response, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("status code", response.StatusCode)
	fmt.Println("content length ", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(content)
}
