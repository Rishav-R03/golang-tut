package main

import (
	"fmt"
	"os"
)

func main() {
	// Print all arguments, including the program path
	fmt.Println("All arguments:", os.Args)

	// Get only the arguments without the program name
	argsWithoutProg := os.Args[1:]
	fmt.Println("Arguments without program name:", argsWithoutProg)
	//if no arguments
	if len(argsWithoutProg) <= 1{ //not comparing to 0 because os always has program path as a argument
		//this is means no user arguments 
		fmt.Println("No argument provided")
	} else {
		// Access individual arguments
		fmt.Println("First argument:", argsWithoutProg[0])
	}
}
