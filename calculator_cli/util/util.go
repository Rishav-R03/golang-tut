package util

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var Reader = bufio.NewReader(os.Stdin)

func ShowMenu() {
	menu := []string{
		"Add",
		"Subtract",
		"Multiply",
		"Divide",
		"Exit",
	}

	for i, item := range menu {
		fmt.Printf("%d. %s\n", i+1, item)
	}

}

func ReadOperands() (float64, float64, error) {
	fmt.Printf("Enter first number")
	a, err := Reader.ReadString('\n')
	if err != nil {
		return 0, 0, fmt.Errorf("invalid input: %w", err)
	}
	fmt.Printf("Enter second number")
	b, err := Reader.ReadString('\n')
	if err != nil {
		return 0, 0, fmt.Errorf("invalid input %w", err)
	}
	n1, err := strconv.ParseFloat(strings.TrimSpace(b), 64)
	if err != nil {
		return 0, 0, fmt.Errorf("Invalid input: %w", err)
	}
	n2, err := strconv.ParseFloat(strings.TrimSpace(a), 64)
	if err != nil {
		return 0, 0, fmt.Errorf("Invalid input: %w", err)
	}
	return n1, n2, nil
}
