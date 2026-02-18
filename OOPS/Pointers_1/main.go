package main

import "fmt"

// var x int = 10 stores value 10
// var p *int = &x stores the address where 10 is sitting

// &: (Address-of) => Where is this variable located?
// *: (Dereferencing) => Go to this address and tell me what value is there

func main() {
	var x int = 10
	fmt.Println("X has value: ", x)

	var y int = 5
	var p *int = &x
	fmt.Println("P is the pointer to address of value of X", p)

	var p2 *int = &y

	var z = *p + *p2
	fmt.Println("The sum of X and Y ", z)

}
