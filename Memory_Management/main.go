package main

// Memory Management - Exercise 1
/*
*A Torn Read is like trying to read a two-word sign while someone is halfway through changing it.
 If the sign says "GO LEFT" and a worker changes it to "STOP RIGHT," a "torn read" happens if you look at the sign at the exact moment the worker has swapped "GO" for "STOP" but hasn't touched "LEFT" yet.
 You end up reading "STOP LEFT"—a message that was never intended.
*/

import "fmt"

func main() {
	//For this we are using 'Slice' in Go lang slice has 3 parts
	// A pointer to the data
	// The length of slice
	// The capacity

	//Because it takes the CPU multiple steps to update all three, a data race can cause one goroutine to read the new pointer but the old length,
	// leading to a crash or corrupted data.
	fmt.Println("Welcome to Torn Read exercise")

	//A slice is a multi-word structure (pointer + length + capacity)
	//We start with a slice of 0 int
	sharedSlice := []int{}

	//Go routine 1: The writer
	//It constantly toggles the slice between two states
	go func() {
		for {
			//State A: A slice with 10 zeroes
			sharedSlice = make([]int, 10)
			//State B: A slice with 1000 zeroes
			sharedSlice = make([]int, 1000)
		}
	}()

	//Go routine 2: The reader
	//It constantly tries to read the slice

	for {
		//TORN Read Potential:
		//The reader might grab the "pointer" from the state B (1000 items)
		//but the "length" from state A (10 items).
		//Or worse, a pointer to memory that was just replaced.
		localCopy := sharedSlice

		//if w catch it at just the right (wrong) time, the program
		//might crack or show inconsistent length
		if len(localCopy) > 0 {
			_ = localCopy[0] // just access the data to trigger a crash if torn
		}
	}
}
