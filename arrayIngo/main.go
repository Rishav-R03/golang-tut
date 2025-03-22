package main

import "fmt"

func main() {
	fmt.Println("Arrays in Golang")

	var fruitlist [5]string

	fruitlist[0] = "Apple"
	fruitlist[1] = "Tomato"
	fruitlist[2] = "Peach"
	fruitlist[4] = "Banana"

	fmt.Println("Fruit list is: ", fruitlist)
	fmt.Println("Fruit list is: ", len(fruitlist)) // get length of the array
	// inline initialization and assignment
	var veglist = [3]string{"Potato", "Beans", "Mushroom"}
	fmt.Println("Veg list is: ", veglist)
	fmt.Println("Veg list is: ", len(veglist))
}
