package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// isAdult determines if the given age qualifies as an adult (18+)
func isAdult(age int) bool {
	return age >= 18
}

func main() {
	// Create a new buffered reader to handle user input
	reader := bufio.NewReader(os.Stdin)

	// Prompt the user for their age
	fmt.Print("Enter your age: ")
	ageInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Trim any leading or trailing whitespace (including newlines)
	ageInput = strings.TrimSpace(ageInput)

	// Convert the input string to an integer
	age, err := strconv.Atoi(ageInput)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return
	}

	// Determine if the user is an adult
	if isAdult(age) {
		fmt.Println("You are an adult!")
	} else {
		fmt.Println("You are a kid!")
	}
}
