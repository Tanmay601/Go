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
	log.Fatal(http.ListenAndServe(":8000", r))
	fmt.Println("Server started at port 8000")
}
