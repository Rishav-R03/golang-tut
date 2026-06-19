package main

import "fmt"

func main() {
	name := "Rishav" // looks strings but
	// inside golang
	// R  i  s  h  a  v
	// 82 105 115 104 97 118

	// go stores
	// type string struct {
	// 	ptr *byte
	// 	len int
	// }

	/**
	* Strings are immutable
	 */
	for i := 0; i < len(name); i++ {
		fmt.Println(name[i])
	}

	for _, r := range name {
		fmt.Printf("%c\n", r)
	}
}
