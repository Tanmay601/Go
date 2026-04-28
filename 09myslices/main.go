package main

import "fmt"

func main() {
	fmt.Println("Slices") //slices are arrays that can grow and shrink dynamically in go

	var fruitList = []string{"apple", "orange"}
	fmt.Printf(" Type of Fruit list is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:]) //slices the slice
	fmt.Println(fruitList)

	highScore := make([]int, 4) //make is used for memory allocation

	highScore[0] = 234
	highScore[1] = 945
	highScore[2] = 565
	highScore[3] = 12

	fmt.Println(highScore)

}
