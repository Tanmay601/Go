package main

import "fmt"

func main() {
	fmt.Println("Arrays")

	var fruitList [4]string //arrya declaration requires numeric specification in Go

	fruitList[0] = "apple"
	fruitList[1] = "banana"
	fruitList[2] = "pineapple"
	fruitList[3] = "mango"

	fmt.Println("Fruit list is:", fruitList)

}
