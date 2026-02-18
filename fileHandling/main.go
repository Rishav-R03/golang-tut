package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in go lang")
	content := "This needs to go in a file - LearnCodeOnline.in"

	file, err := os.Create("./mylcofile.txt")

	checkNilError(err)
	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}
	fmt.Println("length is :", length)
	defer file.Close()
	readFile("./mylcofile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)

	fmt.Println("The text data inside file is \n", string(databyte))
	checkNilError(err)
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
