package main

import (
	"fmt"
	"sort"
)

// slices is not an array, it is a reference to an array. It is a data structure that provides a dynamic view into the elements of an array. Slices are more flexible than arrays because they can grow and shrink in size, and they can be easily passed around as function arguments.

func main() {
	//creating slice for integers
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums) // Output: [1 2 3 4 5]
	//ptr -> [1,2,3,4,5]
	//len -> 5
	//cap -> 5

	//creating slice

	//1. Using literal syntax
	literal_slice := []string{"Go", "is", "awesome"}

	//2. Using make function
	make_slice := make([]int, 5)

	// make(type, length, capacity)
	// make([]int, 5) -> make([]int, 5, 5)
	fmt.Println(literal_slice) // Output: [Go is awesome]
	fmt.Println(make_slice[0]) // Output: [0 0 0 0 0]

	//3. Using slicing operator
	original_array := [5]int{1, 2, 3, 4, 5}
	sliced_slice := original_array[1:4] // This creates a slice that includes elements from index 1 to index 3 (not including index 4)

	fmt.Println(sliced_slice) // Output: [2 3 4]

	// Modifying the sliced slice will also modify the original array
	sliced_slice[0] = 10
	fmt.Println(original_array) // Output: [1 10 3 4 5]

	//==============================================================================

	//1. Operation 1: Appending to a slice
	slice_one := []int{1, 3, 5}
	slice_one = append(slice_one, 7)
	slice_one = append(slice_one, 9, 11) // multi append
	fmt.Println(slice_one)               // Output: [1 3 5 7]

	//2. Operation 2: Copying a slice
	slice_two := make([]int, len(slice_one))
	copy(slice_two, slice_one)
	fmt.Println(slice_two) // Output: [1 3 5 7]

	//3. Operation 3: Slicing a slice
	slice_three := slice_one[1:4] // This creates a new slice that includes elements from index 1 to index 3 (not including index 4)
	fmt.Println(slice_three)      // Output: [3 5 7]

	//4. Operation 4: Modifying a slice
	slice_three[0] = 10
	fmt.Println(slice_one) // Output: [1 10 5 7]

	//5. Operation 5: Length and Capacity
	fmt.Println(len(slice_one)) // Output: 4
	fmt.Println(cap(slice_one)) // Output: 4 (capacity is the total size of the underlying array)

	//6. Operation 6: Removing elements from a slice
	slice_one = append(slice_one[:1], slice_one[2:]...) // This removes the element at index 1 (10 in this case)
	fmt.Println(slice_one)                              // Output: [1 5 7]

	var i int = 0
	//7. Operation 7: Removing at element at index i
	var s []int = append(slice_one[:i], slice_one[i+1:]...)
	fmt.Println(s)

	ind := 2
	value := 10
	slice_one = append(s[:ind], append([]int{value}, slice_one[ind:]...)...)
	fmt.Println(slice_one) // Output: [1 5 10 7]

	//8. Operation 8: Clearing a slice
	slice_one = slice_one[:0]
	// or s = nil
	fmt.Println(slice_one) // Output: []

	//9. Operation 9: Iterationg

	//1. Range loop
	for indx := 0; i < len(slice_two); indx++ {
		fmt.Println(slice_two[indx])
	}

	//2. For loop
	for _, value := range slice_two {
		fmt.Println(value)
	}

	sort.Ints(slice_two)
	fmt.Println(slice_two) // Output: [1 3 5 7]

	//Custom sorting using sort.Slice
	sort.Slice(slice_two, func(i, j int) bool {
		return slice_two[i] > slice_two[j] // Sort in descending order
	})
	fmt.Println(slice_two) // Output: [7 5 3 1]
}
