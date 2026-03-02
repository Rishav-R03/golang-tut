package mapsingo

func main() {
	m := make(map[string]int)

	m["one"] = 1
	m["two"] = 2
	m["three"] = 3

	println(m["one"])

	// creating a map using literal syntax
	ma := map[string]int{
		"four": 4,
	}
	println(ma["four"])

	//insert/ update
	m["one"] = 10
	println(m["one"])

	//delete
	delete(m, "two")
	println(m["two"]) // Output: 0 (zero value for int)

	//check if key exists
	value, exists := m["three"]
	if exists {
		println("Key 'three' exists with value:", value)
	} else {
		println("Key 'three' does not exist")
	}

	//iterate over map
	for key, value := range m {
		println("Key:", key, "Value:", value)
	}

	//length of map
	println("Length of map:", len(m))

	//maps are reference types
	m2 := m
	m2["one"] = 100
	println(m["one"]) // Output: 100 (modification through m2 affects m)

	for key := range m {
		println("Key:", key)
	}

	for _, val := range m {
		println("Value:", val)
	}

	map_slices := map[string][]int{
		"evens": {2, 4, 6},
		"odds":  {1, 3, 5},
	}
	println(map_slices["evens"][0]) // Output: 2

	//Frequency count using maps
	str := "hello world"
	frequency := make(map[string]int)
	for _, char := range str {
		frequency[string(char)]++
	}
	println("Frequency of 'l':", frequency["l"]) // Output: 3

	//maps are reference types
	m3 := m
	m3["three"] = 300
	println(m["three"]) // Output: 300 (modification through m3 affects m)
	
}
