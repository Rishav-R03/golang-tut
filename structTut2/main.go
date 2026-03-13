package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println("This tutorial is for struct application !")
	fmt.Println("Vertex are ", Vertex{1, 2})

	robot := Robot{Name: "rd2"}
	human := Human{FirstName: "alice"}

	Introduce(robot)
	Introduce(human)

}

type Robot struct {
	Name string
}

func (r Robot) Speak() string {
	return "Hi, I am " + r.Name
}

/*
	Interfaces

	1. Define the Interface
*/

type Speaker interface {
	Speak() string
}

//2. Define another struct

type Human struct {
	FirstName string
}

func (h Human) Speak() string {
	return "Hello, my name is " + h.FirstName
}

//3. A function that accepts the Interface

func Introduce(s Speaker) {
	fmt.Println(s.Speak())
}
