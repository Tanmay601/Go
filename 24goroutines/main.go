package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup // waitgroup to wait for all goroutines to finish
// WaitGroup is a struct that is used to wait for a collection of goroutines to finish. It provides a way to wait for multiple goroutines to complete their execution before proceeding further in the program.
// The WaitGroup has three main methods: Add, Done, and Wait. Add is used to set the number of goroutines to wait for, Done is called by each goroutine when it finishes its work, and Wait is called by the main goroutine to block until all the other goroutines have finished.
// WaitGroup is useful when you have multiple concurrent operations and you want to ensure that all of them complete before moving on to the next step in your program.

var mut sync.Mutex // mutex to protect the shared resource
func main() {

	// go greeter("hello") // go keyword is used to run the function in a separate goroutine. It allows the function to run concurrently with the main function.
	// greeter("Tanmay")

	websiteList := []string{
		"https://www.google.com",
		"https://www.youtube.com",
		"https://www.facebook.com",
	}

	for _, website := range websiteList {
		wg.Add(1)                 // tell waitgroup to expect 1 more goroutine
		go getStatusCode(website) // go keyword is used to run the function in a separate goroutine. It allows the function to run concurrently with the main function.

		// if no go keyword above the status coddes will be printed in order but with go keyword the status codes will be printed in random order as the goroutines will run concurrently and the order of execution is not guaranteed.
	}

	wg.Wait() // main will WAIT here until all goroutines call wg.Done()
	fmt.Println("All done!")

}

// func greeter(s string) {

// 	for i := 0; i < 6; i++ {
// 		time.Sleep(4 * time.Millisecond) // Sleep function is used to pause the execution of the current goroutine for a specified duration. In this case, it pauses for 4 milliseconds.
// 		fmt.Println(s)
// 	}

// }

func getStatusCode(endpoint string) {

	defer wg.Done() // tell waitgroup that this goroutine is done when the function returns
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOPS in endpiont")
	} else {
		mut.Lock()   // lock the mutex to protect the shared resource
		mut.Unlock() // unlock the mutex to allow other goroutines to access the shared resource

		fmt.Printf("%d 200 status code for %s", res.StatusCode, endpoint)

	}
}
