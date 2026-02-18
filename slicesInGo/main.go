package main

import "fmt"

/*
*Unlike arrays, slices are typed only by the
elements they contain (not the number of elements).
An uninitialized slice equals to nil and has length 0.
*/
func main() {
	fmt.Println("Welcome to Slices class")

	var fruitlist = []string{}
	var veglist = []string{"Potato", "Beans", "Mushroom"}
	fmt.Printf("Type of fruitlist is: %T\n", fruitlist)

	// fruitlist = append(fruitlist, "Apple", "Tomato", "Peach")
	fruitlist = append(fruitlist, "Apple", "Tomato", "Peach")
	fmt.Println("Fruit list is: ", fruitlist)
	fmt.Println("Fruit list is: ", fruitlist[1:])
	fmt.Println("Fruit list is: ", fruitlist[1:3]) // end range in not included
	fmt.Println("Fruit list is: ", veglist)

	// allocation of memory
	var newVegList = make([]string, 3)
	fmt.Println("New veg list is: ", newVegList)

	// copy one slice to another
	copy(newVegList, veglist)
	fmt.Println("New veg list is: ", newVegList)

	// highscores array and assigning values to it
	highscores := make([]int, 4)

	highscores[0] = 234
	highscores[1] = 945
	highscores[2] = 565
	highscores[3] = 444

	fmt.Println("Highscores are: ", highscores)

	// highscores = append(highscores, 555, 666, 777)
	/**
	below the append() is able to add extra elements to slice
	without increasing size of the initialized slice instead it
	re assign the slice.
	*/
	highscores = append(highscores, 555, 666, 777)
	fmt.Println("Highscores are: ", highscores)

	//remove slice's elemets using the index

	var courses = []string{"Python", "Go", "JavaScript"}
	fmt.Println(courses)
	index := 2
	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println(courses)

	names := make([]string, 3) // pre defined size
	names[0] = "Abhinav"
	names[1] = "Raghav"
	names[2] = "Akhand"
	print("set", names)
	print("set", names[2])
}
