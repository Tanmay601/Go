package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// define the struct that represents your JSON shape
// the backtick tags control what the JSON keys look like
type Course struct {
	CourseName string `json:"coursename"`
	Price      int    `json:"price"`
	Platform   string `json:"platform"`
}

// this is what the response from httpbin looks like
// we can decode it back into a struct too
type HttpbinResponse struct {
	Json Course `json:"json"` // httpbin puts your sent JSON inside a "json" field
	Url  string `json:"url"`
}

func main() {
	fmt.Println("--- Sending JSON data in Go ---")
	PerformPostJsonRequest()
}

func PerformPostJsonRequest() {
	const myurl = "https://httpbin.org/post"

	// step 1: create a struct instance with your data
	course := Course{
		CourseName: "Let's go with golang",
		Price:      299,
		Platform:   "learnCodeOnline.in",
	}

	// step 2: marshal the struct into JSON bytes
	// json.Marshal converts Go struct → JSON
	jsonData, err := json.Marshal(course)
	if err != nil {
		panic(err)
	}

	// see exactly what JSON is being sent
	fmt.Println("JSON being sent:")
	fmt.Println(string(jsonData))
	fmt.Println()

	// step 3: send the request
	// bytes.NewReader wraps the JSON bytes so http.Post can read them
	response, err := http.Post(myurl, "application/json", bytes.NewReader(jsonData))
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("Status code:", response.StatusCode)
	fmt.Println()

	// step 4: read the raw response body
	content, _ := ioutil.ReadAll(response.Body)

	// step 5: unmarshal the response back into a Go struct
	// json.Unmarshal converts JSON → Go struct
	var httpbinResp HttpbinResponse
	err = json.Unmarshal(content, &httpbinResp)
	if err != nil {
		panic(err)
	}

	// now you can access response fields as Go values, not raw strings
	fmt.Println("--- Decoded Response ---")
	fmt.Println("Course Name :", httpbinResp.Json.CourseName)
	fmt.Println("Price       :", httpbinResp.Json.Price)
	fmt.Println("Platform    :", httpbinResp.Json.Platform)
	fmt.Println("Request URL :", httpbinResp.Url)

	fmt.Println()

	// raw response for comparison
	fmt.Println("--- Raw Response ---")
	fmt.Println(string(content))
}
