package main

import "fmt"


type UserData struct  { // struct is a collection of fields. 
	fistName string
	lastName string
	email string
	numberOfTickets int
}

func main() {
	user := UserData{ // struct initialization
		fistName: "John",
		lastName: "Doe",
		email: "pradumnasaraf@gmail.com",
		numberOfTickets: 2,
	}
	fmt.Println(user.fistName) 
	fmt.Println(user)

	user2 := UserData{ "Jon", "Doe", "jpo@mail.com", 2 } // struct initialization
	fmt.Println(user2.fistName)
	fmt.Println(user2) // prints the struct
	fmt.Printf("Details are %+v", user2) // %+v prints the struct with field names
}