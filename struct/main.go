package main

import (
	"fmt"
)

func main() {

	fmt.Println("This is a struct tutorial")
	// there is no inheritance in go lang; No super or parent

	user1 := User{"Rishav", "rishav042023@gmail.com", 22, true}
	fmt.Println("The first user in the struct is ", user1)
	fmt.Printf("The Name is  %v and email is %v  ", user1.Name, user1.Email)
}

type User struct { // U is capital because it is kind of class
	Name   string
	Email  string
	Age    int
	Status bool
}
