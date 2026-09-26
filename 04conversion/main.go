package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza app")
	fmt.Println("Please rate our pizza ebtween 1 and 5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating, ", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) // simple (input, 64) does not work due to a trailing /n always hidden at the end so after running a syntax error comes to trim that /n we use Trimspace

	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println("Added 1 to your rating:", numRating+1)
	}

}
