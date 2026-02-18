package functions

import "fmt"

func main() {
	result := adder(4, 5)
	fmt.Print(result)
}

func adder(number1 int, number2 int) int {
	return number1 + number2
}
