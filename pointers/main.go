package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")

	var ptr *int                                             // pointer to integer
	fmt.Println("Value of pointer before assignment: ", ptr) // nil
	var num int = 10
	ptr = &num // address of num is assigned to ptr, ampersand is for reference
	fmt.Println("Value of number: ", num)
	fmt.Println("Value of pointer after assignment: ", ptr)
	fmt.Println("Value of pointer address after assignment: ", *ptr)

	*ptr = *ptr * 2 //works for +,-,*,/
	fmt.Println("Value of number after dereferencing: ", num)
	
}
