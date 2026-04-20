package main

import "fmt"

func main() {
	var username string = "Tanmay Rajpoot"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255 //(for uint more than 255 number cant be assigned)
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.3243232 //(for float 32 only 5 values after decimal will appear in the log)
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	//default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	//implicit type
	var website = "youtube.com"
	fmt.Println(website)

	//no var type
	numberOfUser := 3000000
	fmt.Println(numberOfUser)
}
