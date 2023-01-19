package main

import "fmt"

func main(){
	var confrenceName = "Go Confrence"
	const confrenceTickets = 50
	var remainingTickets uint = 50
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Printf("Welcome to %v. \n", confrenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", confrenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	fmt.Println("Please enter your first name to get your ticket")
	fmt.Scan(&firstName)	

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)	

	fmt.Println("Enter your email address:")
	fmt.Scan(&email)	

	fmt.Println("Enter number of tickets")
	fmt.Scan(&userTickets)	

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you, %v %v for booking %v tickets. You will receive a confirmation email at %v.\n",firstName, lastName, userTickets, email)
	fmt.Printf("%v remaining tickets for %v\n", remainingTickets, confrenceName)
}