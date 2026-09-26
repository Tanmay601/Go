package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=avhsa234"

func main() {
	fmt.Println("URLs in golang")
	fmt.Println(myurl)

	//parsing
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())

	qparams := result.Query()

	fmt.Printf("The type of query params is: %T\n", qparams)

	partsOfUrl := &url.URL{

		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/tutcss",
		RawQuery: "user=TR",
	}
	anotherURL := partsOfUrl.String()
	fmt.Println("Another URL is:", anotherURL)

}
