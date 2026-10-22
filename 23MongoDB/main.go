package main

import (
	"MongoDB/router"
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("MongoDB API")
	r := router.Router()
	fmt.Println("Starting the application...")
	http.ListenAndServe(":4000", r)
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Server started at port 4000")
}
