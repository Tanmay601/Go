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

// same struct but with omitempty on Price and Discount
// omitempty skips the field in JSON if it has a zero value
// zero value: "" for string, 0 for int, false for bool, nil for pointer
type CourseOmit struct {
	CourseName string `json:"coursename"`
	Price      int    `json:"price,omitempty"`    // skipped if Price == 0
	Discount   string `json:"discount,omitempty"` // skipped if Discount == ""
	Platform   string `json:"platform,omitempty"` // skipped if Platform == ""
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
	fmt.Println()
	fmt.Println("====================================")
	fmt.Println()
	DemoOmitempty()
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

func DemoOmitempty() {
	const myurl = "https://httpbin.org/post"

	fmt.Println("--- omitempty Demo ---")
	fmt.Println()

	// CASE 1: without omitempty — zero values still appear in JSON
	fmt.Println("WITHOUT omitempty (Price=0, Platform=''):")
	course1 := Course{
		CourseName: "Let's go with golang",
		Price:      0,  // zero value — still appears in JSON
		Platform:   "", // empty string — still appears in JSON
	}
	jsonData1, _ := json.MarshalIndent(course1, "", "  ")
	fmt.Println(string(jsonData1))
	fmt.Println()

	// CASE 2: with omitempty — zero values are stripped out
	fmt.Println("WITH omitempty (Price=0, Discount='', Platform=''):")
	course2 := CourseOmit{
		CourseName: "Let's go with golang",
		Price:      0,  // zero value → SKIPPED
		Discount:   "", // empty string → SKIPPED
		Platform:   "", // empty string → SKIPPED
	}
	jsonData2, _ := json.MarshalIndent(course2, "", "  ")
	fmt.Println(string(jsonData2))
	fmt.Println()

	// CASE 3: with omitempty but all fields filled — all appear normally
	fmt.Println("WITH omitempty (all fields filled):")
	course3 := CourseOmit{
		CourseName: "Let's go with golang",
		Price:      299,
		Discount:   "10%",
		Platform:   "learnCodeOnline.in",
	}
	jsonData3, _ := json.MarshalIndent(course3, "", "  ")
	fmt.Println(string(jsonData3))
	fmt.Println()

	// CASE 4: actually send the omitempty JSON to httpbin to prove
	// the fields are stripped before they even leave your machine
	fmt.Println("Sending omitempty JSON (case 2) to httpbin...")
	response, err := http.Post(myurl, "application/json", bytes.NewReader(jsonData2))
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	var httpbinResp struct {
		Json map[string]interface{} `json:"json"`
	}
	json.Unmarshal(content, &httpbinResp)

	fmt.Println("What httpbin actually received:")
	for key, val := range httpbinResp.Json {
		fmt.Printf("  %s: %v\n", key, val)
	}
	fmt.Println()
	fmt.Println("Notice: price, discount, platform are gone —")
	fmt.Println("omitempty stripped them before sending, not on the server side")
}
