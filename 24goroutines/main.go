package main

import (
	"fmt"
	"time"
)

func main() {

	go greeter("hello") // go keyword is used to run the function in a separate goroutine. It allows the function to run concurrently with the main function.
	greeter("Tanmay")

}

func greeter(s string) {

	for i := 0; i < 6; i++ {
		time.Sleep(4 * time.Millisecond) // Sleep function is used to pause the execution of the current goroutine for a specified duration. In this case, it pauses for 4 milliseconds.
		fmt.Println(s)
	}

}
