package main

import "fmt"

func main() {
	fmt.Println("maps in go")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["TS"] = "TypeScript"
	languages["NS"] = "NodeJS"

	fmt.Println("List of all languages:", languages)
	fmt.Println("JS full form is:", languages["JS"])

}
