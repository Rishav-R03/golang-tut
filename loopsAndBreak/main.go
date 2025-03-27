package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, welcome to loops and break tutorial")
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturaday"}
	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	for i := range days {
		fmt.Println(days[i])
	} // same as above but using range

	for index, day := range days { // using comma ok syntax
		fmt.Printf("index is %v and value is %v\n", index, day)
	}

	rougeValue := 1
	for rougeValue < 10 {
		fmt.Println("Value is : ", rougeValue)
		rougeValue++
	}
}
