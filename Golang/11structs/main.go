package main

import "fmt"

func main() {
	fmt.Println("structs in golang")
	//no inheritence in golang : no super or no parent

	sanjay := User{"sanjay", "sanjayhardageri97@gmail.com", true, 26}
	fmt.Println(sanjay)
	fmt.Printf("assign struct: %+v\n", sanjay)
	fmt.Printf("name:%v\n,age:%v\n", sanjay.Name, sanjay.Age)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
