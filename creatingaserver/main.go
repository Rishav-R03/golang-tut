package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("status code: ", response.StatusCode)
	fmt.Println("content length is: ", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}

func main() {
	fmt.Println("Creating a server in Go-lang")
	PerformGetRequest()
}
