package main

import "fmt"

func main() {
	fmt.Println("Welcome to Slices practice")
	names := make([]string, 3)
	names[0] = "Rishav"
	names[1] = "Harsh"
	names[2] = "Abhinav"

	println("Set", names)

	fmt.Println("Size of the names", len(names))

	names = append(names, "Balaji")
	names = append(names, "Harshal", "Namokar")

	fmt.Println("Appended", names)
	fmt.Println("Size of the appended names", len(names))

	copyNames := make([]string, len(names))
	copy(copyNames, names)

	fmt.Println("copied ", copyNames)

	slice1 := copyNames[2:5]
	fmt.Println(slice1)

	for i := 0; i < 3; i++ {
		if names[i] == "Abhinav" {
			fmt.Println("index", i)
		}
	}

	
}
