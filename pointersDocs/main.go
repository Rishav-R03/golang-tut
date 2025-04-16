package main

import "fmt"

func main() {

	fmt.Println("Welcome to pointers tutorial in go lang!")

	var i int = 9
	address_to_i := &i
	pointer_to_i := i
	fmt.Println(i)
	fmt.Println(address_to_i)
	fmt.Println(pointer_to_i)
	*address_to_i = 21
	fmt.Println(address_to_i)
	fmt.Println(i)

}
