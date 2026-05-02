package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1 // generates a number between 1 and 6
	fmt.Println("You rolled a", diceNumber)

}
