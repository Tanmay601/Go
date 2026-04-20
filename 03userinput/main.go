package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "weclome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter ratinf for our pizza:")

	// comma ok || err ok syntax

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of this ratinf is %T", input) // it shows string type but the ratinf is numeric type

}
