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

// GetStatus is a 'Value Receiver'
// It creates a COPY of the struct. Good for reading data.
func (u user) GetStatus() {
	fmt.Println("Is user active?:", u.Status)
}

// NewMail is a 'Pointer Receiver'
// It passes the actual ADDRESS of the struct.
// Without the '*', the email would only change inside this function, not in main.
func (u *user) NewMail(newMail string) {
	u.email = newMail
}
