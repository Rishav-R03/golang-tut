package main // Defines the main package, required for standalone Go applications

import (
	"bufio"   // Used for buffered input
	"fmt"     // Provides formatted I/O functions
	"os"      // Provides platform-independent interface to operating system functionality
	"strconv" // Converts strings to different numeric types
	"strings" // Provides string manipulation functions
)

func main() {
	// Create a new buffered reader to read input from standard input (console)
	reader := bufio.NewReader(os.Stdin)

	// Prompt the user to enter their name
	fmt.Print("Enter your name: ")
	
	// Read the user input until a newline character is encountered
	name, _ := reader.ReadString('\n')
	
	// Remove any leading or trailing whitespace (including newline characters)
	name = strings.TrimSpace(name)

	// Prompt the user to enter a rating between 1 to 5
	fmt.Print("Enter rating between 1 to 5: ")
	
	// Read the user input for the rating
	ratingStr, _ := reader.ReadString('\n')
	
	// Remove any leading or trailing whitespace (including newline characters)
	ratingStr = strings.TrimSpace(ratingStr)
	
	// Convert the rating string to an integer
	rating, err := strconv.Atoi(ratingStr)

	// Check if conversion failed (invalid input) or rating is out of range
	if err != nil || rating < 1 || rating > 5 {
		// Print an error message if the rating is not valid
		fmt.Println(name, "please enter a valid rating between 1 and 5.")
	} else {
		// Print a thank you message with the user's name and rating
		fmt.Printf("Dear %s, thank you for rating us with %d!\n", name, rating)
	}
}