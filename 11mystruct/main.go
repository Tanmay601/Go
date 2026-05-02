package main

import "fmt"

func main() {

	fmt.Println("struct")

	//no inheritence goland; no super or parent inm golang

	Tanmay := user{"Tanmay", "4VW4j@OG.com", true, 23}
	fmt.Println("User is:", Tanmay)
	fmt.Printf("Tanmay details are: %+v\n", Tanmay)
}

type user struct {
	name   string
	email  string
	Status bool
	Age    int
}
