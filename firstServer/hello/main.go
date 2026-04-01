package main

import (
	"fmt"
	"greetings"
)

func main() {
	fmt.Println("Enter your age:")
	var age int
	_, age_err := fmt.Scanf("%d", &age)
	if age_err != nil || age < 18 {
		fmt.Println("Error reading age or age is less than 18:", age_err)
		return
	}
	fmt.Println("Enter your name:")
	var name string
	_, name_err := fmt.Scanf("%s", &name)
	if name_err != nil || len(name) == 0 || name == " " || len(name) < 3 {
		fmt.Println("Error reading name or enter a valid name", name_err)
		return
	} else {
		fmt.Println(greetings.Greet(name))
	}
}
