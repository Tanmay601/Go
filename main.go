package main

import "fmt"

func main() {

	var choice int

	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Print("Enter choice: ")

	fmt.Scanln(&choice)

	switch choice {

	case 1:
		fmt.Println("Addition Selected")

	case 2:
		fmt.Println("Subtraction Selected")

	default:
		fmt.Println("Wrong Choice")
	}
}
