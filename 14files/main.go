package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	fmt.Println("files in golang")
	content := "This is a sample text to write into the file."

	file, err := os.Create("./mylcogofile.txt")

	checkNilErr(err)

	//if err != nil {
	//	panic(err)
	//}

	length, err := io.WriteString(file, content)

	checkNilErr(err)
	//if err != nil {
	//	panic(err)
	//}
	fmt.Println("length is:", length)
	defer file.Close()

	readFile("./mylcogofile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)

	checkNilErr(err)
	//if err != nil {
	//	panic(err)
	//}

	fmt.Println("Data read from file is \n", string(databyte))
}

func checkNilErr(err error) {

	if err != nil {
		panic(err)
	}
}
