package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("time study in golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) //standard way of formatting dates day time in GO god knows why
	//crazy way to format but it is what it is brrr

	createdDate := time.Date(2022, time.January, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)
}
