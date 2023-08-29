package main

import "fmt"

func main() {
	fmt.Println("structs in golang")
	//no inheritence in golang : no super or no parent

	sanjay := User{"sanjay", "sanjayhardageri97@gmail.com", true, 26}
	fmt.Println(sanjay)
	fmt.Printf("assign struct: %+v\n", sanjay)
	fmt.Printf("name:%v\n,age:%v\n", sanjay.Name, sanjay.Age)
	sanjay.GetStatus()
	sanjay.NewMail()

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() { //use pointers for no copies
	u.Email = "test@go.dev"
	fmt.Println("email is: ", u.Email)
}
