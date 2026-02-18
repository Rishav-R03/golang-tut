package main

import (
	"bufio"
	"fmt"
	"os"
)

// other way of putting variables

var CollegeName = "VIT Bhopal"

// capital letter at beginning shows that it is publicly accessible
func main() {
	//bufio.NewReader is a function to read anything
	reader := bufio.NewReader(os.Stdin) // os.Stdin is the standard input stream
	// := is a short hand for declaring a variable called walrus
	fmt.Printf("Enter your username: ")
	// comma of || err ok is a kind of try catch in go
	username, _ := reader.ReadString('\n') // \n is a new line character
	fmt.Printf("Enter your age: ")
	userage, _ := reader.ReadString('\n')
	fmt.Println("Your username is: ", username)
	fmt.Println("Your age is: ", userage)
	fmt.Println("Your college is: ", CollegeName)

}
