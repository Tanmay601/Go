package main

<<<<<<< HEAD
import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello MOD")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))

}
func greeter() {

	fmt.Println("Hey mod users")

}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Tanmay Page</h1>"))
=======
import "fmt"

func main() {

	var choice int

	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Print("Enter choice: ")

	fmt.Scanln(&choice)

	switch choice {

	case 1:
		fmt.Println("Addition Selected")

	case 2:
		fmt.Println("Subtraction Selected")

	default:
		fmt.Println("Wrong Choice")
	}
>>>>>>> 3cc9b20b8d088f8ceb451483cc42036cac0b4d1c
}
