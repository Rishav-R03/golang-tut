package reverseString

// go mod init example.com/hello
// go: creating new go.mod: module example.com/hello
// adding a dependency
// go get golang.org/x/example/hello/reverse

import (
	"golang.org/x/example/hello/reverse"
)

func ReverseString() {
	// fmt.Println(reverse.String("Hello"))
	return reverse.String("Hello")
}
