package main

import (
	"calculator_cli/handler"
	"fmt"
	"net/http"
)

const (
	AddOperation  = 1
	SubOperation  = 2
	MulOperation  = 3
	DivOperation  = 4
	ExitOperation = 5
)

func main() {

	// for {
	// 	util.ShowMenu()

	// 	var choice int
	// 	fmt.Print("Choose: ")
	// 	fmt.Scanln(&choice)

	// 	switch choice {

	// 	case AddOperation:
	// 		// add
	// 		a, b, err := util.ReadOperands()
	// 		if err != nil {
	// 			fmt.Printf("[Error] %s", err)
	// 			return
	// 		}
	// 		result := operations.Add(a, b)

	// 		fmt.Println("Result =", result)
	// 	case SubOperation:
	// 		// subtract
	// 		a, b, err := util.ReadOperands()
	// 		if err != nil {
	// 			fmt.Printf("[Error] %s", err)
	// 			return
	// 		}
	// 		result := operations.Substract(a, b)
	// 		fmt.Println("Result =", result)
	// 	case MulOperation:
	// 		// multiply
	// 		a, b, err := util.ReadOperands()
	// 		if err != nil {
	// 			fmt.Printf("[Error] %s", err)
	// 			return
	// 		}
	// 		result := operations.Multiply(a, b)
	// 		fmt.Println("Result =", result)
	// 	case DivOperation:
	// 		// divide
	// 		a, b, err := util.ReadOperands()
	// 		if err != nil {
	// 			fmt.Printf("[Error] %s", err)
	// 			return
	// 		}
	// 		result, err := operations.Divide(a, b)
	// 		if err != nil {
	// 			fmt.Printf("[Error] %s\n", err)
	// 			return
	// 		}
	// 		fmt.Println("Result =", result)
	// 	case ExitOperation:
	// 		fmt.Printf("Exiting gracefully.")
	// 		return
	// 	default:
	// 		fmt.Println("Invalid choice")
	// 	}
	// }

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello API")
	})

	// add endpoint
	http.HandleFunc("/add", handler.AddHandler)
	http.HandleFunc("/sub", handler.SubHandler)
	http.HandleFunc("/mul", handler.MulHandler)
	http.HandleFunc("/div", handler.DivHandler)
	http.ListenAndServe(":8000", nil)
}
