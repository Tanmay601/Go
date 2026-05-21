package main

import "fmt"

func add(a int, b int) {
	fmt.Println("Sum =", a+b)
}

func checkEven(num int) {

	if num%2 == 0 {
		fmt.Println("Even Number")
	} else {
		fmt.Println("Odd Number")
	}
}

func main() {

	checkEven(7)
}
