package main

import "fmt"

func main() {
	fmt.Println("Pointers")

	//var ptr *int
	//fmt.Println("value of the pointer is ", ptr)
	//default value of  apointer is nil obviooo

	myNumber := 23

	var ptr = &myNumber //a pointer which is referencing to some memory address using & sign

	fmt.Println("value of the pointer is ", ptr)
	fmt.Println("value of the pointer is ", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New Value is:", myNumber)

}
