package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps in Golang")
	languages := make(map[string]string)
	languages["Py"] = "Python"
	languages["Go"] = "Golang"
	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"

	fmt.Println(languages)
	fmt.Println(languages["Py"])
}
