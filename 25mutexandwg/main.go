package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("race condition")
	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{} // mutex to protect the shared resource

	var score = []int{0}

	wg.Add(3) // tell waitgroup to expect 3 more goroutines

	go func(wg *sync.WaitGroup, m *sync.Mutex) { //
		fmt.Println("Adding 1 to score")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Adding 2 to score")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Adding 3 to score")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println(score) // the output will be different every time as the goroutines will run concurrently and the order of execution is not guaranteed. This is called a race condition. A race condition occurs when two or more goroutines access a shared resource (in this case, the score slice) at the same time and at least one of them modifies it. This can lead to unexpected behavior and bugs in the program.

}

// run cmd go run --race main.go to detect race conditions in the program. The --race flag enables the race detector, which will report any race conditions that occur during the execution of the program.
// exit status 66 indicates that a race condition was detected. The race detector will print a message indicating the location of the race condition and the goroutines involved. It is important to fix race conditions in concurrent programs to ensure correct behavior and avoid bugs.

// after using mutex the output will be consistent every time as the mutex will ensure that only one goroutine can access the shared resource at a time, preventing race conditions.
