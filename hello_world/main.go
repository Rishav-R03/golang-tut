package main

import (
	"fmt"
	// "time"

	"github.com/google/uuid"
	// an external package
)

func makeToast() { // private
	fmt.Println("Toast is ready!")
}

func main() {
	id := uuid.New()
	fmt.Println("Hello World")
	fmt.Println("Generated ID: ", id)
	///Using Go routines
	go makeToast() // this runs independently
	fmt.Println("water is boiling")
	//adding a delay because if main() finishes then it kills all the process
	// time.Sleep(time.Second) //

	sum := Find_Sum(5, 6)
	fmt.Println(sum)

}
