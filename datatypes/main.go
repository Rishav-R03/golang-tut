package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var Answer bool = false

func isAdult(age int8) bool {
	if age > 18 {
		Answer = true
	}
	return Answer
}
func main() {
	inputReader := bufio.NewReader(os.Stdin)
	var userAge, _ = inputReader.ReadString('\n')
	fmt.Println("user's age: ", userAge)
	convertAge, _ := strconv.Atoi(userAge)
	isUserAdult := isAdult(int8(convertAge))
	if isUserAdult {
		fmt.Println("You are an adult! ")
	} else {
		fmt.Println("You are a kid! ")
	}
}
