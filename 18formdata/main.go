package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb video - LCO")
	PerformGetRequest()
	PerformPostJsonRequest()
	PerformPostFormRequest()
	PerformPostMultipartFormRequest()
}

func PerformGetRequest() {
	const myurl = "https://httpbin.org/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("ByteCount is: ", byteCount)
	fmt.Println(responseString.String())
}

func PerformPostJsonRequest() {
	const myurl = "https://httpbin.org/post"

	//fake json payload

	requestBody := strings.NewReader(`
		{
			"coursename":"Let's go with golang",
			"price": 0,
			"platform":"learnCodeOnline.in"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	// httpbin has no dedicated /postform route, so we reuse /post
	// it inspects the content type and handles form data correctly
	const myurl = "https://httpbin.org/post"

	//formdata

	data := url.Values{}
	data.Add("firstname", "hitesh")
	data.Add("lastname", "choudhary")
	data.Add("email", "hitesh@go.dev")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformPostMultipartFormRequest() {
	const myurl = "https://httpbin.org/post"

	// step 1: create an in-memory buffer that will hold our multipart body
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// step 2: add regular text fields, just like a normal form
	writer.WriteField("firstname", "hitesh")
	writer.WriteField("lastname", "choudhary")

	// step 3: create a temporary sample file on disk so this example
	// is self-contained and runnable without needing an existing file
	tempFile, err := os.CreateTemp("", "sample-*.txt")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tempFile.Name()) // clean up after we're done

	tempFile.WriteString("Hello! This is a sample file uploaded from Go.")
	tempFile.Close()

	// step 4: open the file and attach it to the multipart writer
	file, err := os.Open(tempFile.Name())
	if err != nil {
		panic(err)
	}
	defer file.Close()

	part, err := writer.CreateFormFile("file", "sample.txt")
	if err != nil {
		panic(err)
	}
	io.Copy(part, file) // stream file content into the form part

	// step 5: close the writer to finalize the multipart boundary
	writer.Close()

	// step 6: build the request manually since http.PostForm doesn't support files
	req, err := http.NewRequest(http.MethodPost, myurl, body)
	if err != nil {
		panic(err)
	}
	// this header includes the boundary string the server needs to parse the parts
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
