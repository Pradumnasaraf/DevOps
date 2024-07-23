package main

import (
	"booking-app/validator"
	"fmt"
	"time"
	"sync"
)

const conferenceTickets = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName string
	lastName  string
	email     string
	tickets   uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUser()

	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isVaildTicketNumber := validator.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isVaildTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)
			wg.Add(1) 
			go sendEmail(firstName, lastName, userTickets, email)

			firstNames := getFirstNames()
			fmt.Printf("The First names of booking are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Sorry, tickets are sold out")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("You name is too short. Try entered at least 3 char for first and last name")
			}
			if !isValidEmail {
				fmt.Println("You have entered a wrong email ID")
			}
			if !isVaildTicketNumber {
				fmt.Println("Try entering correct number of tickets.")
			}
		}
	}
	wg.Wait()
}

func greetUser() {
	fmt.Printf("Welcome to %v. \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {

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

	return firstName, lastName, email, userTickets
}
func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		tickets:   userTickets,
	}
	bookings = append(bookings, userData)

	fmt.Println("List of bookings: ", bookings)

	fmt.Printf("Thank you, %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v remaining tickets for %v\n", remainingTickets, conferenceName)
}

func sendEmail(firstName string, lastName string, userTickets uint, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets %v %v\n", userTickets, firstName, lastName)
	fmt.Printf("####################\n")
	fmt.Printf("Sending ticket %v to %v\n", ticket, email)
	fmt.Printf("####################\n")
}
