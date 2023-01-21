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

	greetUser(confrenceName, confrenceTickets, remainingTickets)

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

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isVaildTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isVaildTicketNumber {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you, %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v remaining tickets for %v\n", remainingTickets, confrenceName)

			firstNames := getFirstNames(bookings)
			fmt.Printf("The First names of booking are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Sorry, tickets are sold out")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("You name is too short. Try entring at least 3 char for first and last name")
			}
			if !isValidEmail {
				fmt.Println("You have entred a wrong email ID")
			}
			if !isVaildTicketNumber {
				fmt.Println("Try enterning correct number of tickets.")
			}
		}
	}
}

func greetUser(confName string, confTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v. \n", confName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", confTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var name = strings.Fields(booking)
		firstNames = append(firstNames, name[0])
	}
	return firstNames
}
