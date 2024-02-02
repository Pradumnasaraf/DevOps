package main

import "fmt"

type UserData struct { // struct is a collection of fields.
	fistName        string
	lastName        string
	email           string
	numberOfTickets int
}

func main() {
	user := UserData{ // struct initialization
		fistName:        "John",
		lastName:        "Doe",
		email:           "pradumnasaraf@gmail.com",
		numberOfTickets: 2,
	}

	fmt.Println(user)

	user.greetUser() // Method call
}

func (u UserData) greetUser() { // method with receiver
	fmt.Println("Hello", u.fistName, u.lastName)
}
