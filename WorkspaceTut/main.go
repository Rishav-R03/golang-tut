package main

import (
	"fmt"
	"reverseString/reverseString"
)

func main() {
	fmt.Println("From main.go")
	rev := reverseString.ReverseString()

	fmt.Println(rev)
}
