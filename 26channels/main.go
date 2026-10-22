package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("channels in golang")

	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}

	wg.Add(2) // tell waitgroup to expect 2 more goroutines

	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println(<-myCh)
		wg.Done()

	}(myCh, wg) //
	go func(ch chan int, wg *sync.WaitGroup) {
		myCh <- 5
		myCh <- 10
		wg.Done()

	}(myCh, wg)

	wg.Wait()

	// myCh := make(chan int) // create a channel of type int. Channels are used to communicate between goroutines. They provide a way for one goroutine to send data to another goroutine.
	// // myCh <- 42          // send 42 to the channel.
	// fmt.Println(<-myCh) // receive 42 from the channel.
	// entire code section above created deadlock.
}
