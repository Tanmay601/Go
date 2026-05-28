package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("web req in GO")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Println("Response is of type:", response)

	defer response.Body.Close() //callers responsibility to close the connection

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databytes)
	fmt.Println("Content:", content)

	fmt.Println("Status code is:", response.StatusCode)

	fmt.Println("Header is:", response.Header)

	fmt.Println("Header is:", response.Header.Get("Content-Type"))

	fmt.Println("Header is:", response.Header.Get("Content-Length"))

}
