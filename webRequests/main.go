package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://google.com"

func main() {
	fmt.Println("LCO web request !")

	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Printf("%T", res)
	// fmt.Println(res)
	defer res.Body.Close()

	databytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)
}
