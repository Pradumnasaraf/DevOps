package main

import "fmt"

type UserData struct {
	fistName        string
	lastName        string
	email           string
	numberOfTickets int
}

/*
func (receiver Type) MethodName(parameters) returnType {
    // method body
}
*/

func (u UserData) greetUser() {
	fmt.Println("Hello", u.fistName, u.lastName)
}

func main() {
	user := UserData{ 
		fistName:        "John",
		lastName:        "Doe",
		email:           "pradumnasaraf@gmail.com",
		numberOfTickets: 2,
	}
	
	user.greetUser() // Call the method
}
