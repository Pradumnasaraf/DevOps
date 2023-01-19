package main

import (
	"fmt"
	"strings"
)

func main() {
	var confrenceName = "Go Confrence"
	const confrenceTickets = 50
	var remainingTickets uint = 50
	var bookings []string

	fmt.Printf("Welcome to %v. \n", confrenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", confrenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Please enter your first name to get your ticket")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address:")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets")
		fmt.Scan(&userTickets)
		if userTickets > remainingTickets {
			fmt.Println("Sorry, we have only", remainingTickets, "tickets left. Please try entering a lower number of tickets")
			continue
		}

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you, %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v remaining tickets for %v\n", remainingTickets, confrenceName)

		firstNames := []string{}
		for _, booking := range bookings {
			var name = strings.Fields(booking)
			firstNames = append(firstNames, name[0])
		}

		fmt.Printf("The first names of the booking are: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Sorry, tickets are sold out")
			break
		}
	}
}
