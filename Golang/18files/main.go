package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("import files, read and exec")
	content := "this needs to go in a file"

	file, err := os.Create("./mylcogogile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("length is: ", length)
	defer file.Close()
	readFile("./mylcogogile.txt")

}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	fmt.Println("Text data inside the file\n ", string(databyte))
}
