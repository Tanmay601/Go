package main

import "fmt"

func grade(marks int) {

	switch {

	case marks >= 90:
		fmt.Println("Grade A")

	case marks >= 75:
		fmt.Println("Grade B")

	case marks >= 50:
		fmt.Println("Grade C")

	default:
		fmt.Println("Fail")
	}
}

func main() {

	grade(82)
}
