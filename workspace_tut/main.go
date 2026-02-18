// go mod init example.com/hello
// go: creating new go.mod: module example.com/hello

package main

// adding a dependency
// go get golang.org/x/example/hello/reverse

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {
	fmt.Println(reverse.String("Hello"))
}

