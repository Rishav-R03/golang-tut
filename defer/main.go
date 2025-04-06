package main

import "fmt"

func main() {
	defer fmt.Println("Welcome")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Defer tutorial!")
	deferFunc()
}

func deferFunc(){
	for i:= 0;i<5;i++{
		defer fmt.Println(i)
	}
}

