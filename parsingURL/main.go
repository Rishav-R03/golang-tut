package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=jfus12"

func main() {
	fmt.Println("This tutorial includes the URL parsing technique")

	result, err := url.Parse(myurl)

	if err != nil {
		panic(err)
	}
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Port())
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)

	qparams := result.Query()

	courseName := qparams["coursename"]
	fmt.Println(courseName)

	for _, val := range qparams {
		fmt.Println("params are", val)
	}

	partsofURL := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=rishav",
	}
	anotherURL := partsofURL.String()
	fmt.Println(anotherURL)

}
