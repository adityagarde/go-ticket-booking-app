package main

import (
	"fmt"
	"go-booking-app/helper"
	"strings"
)

const conferenceTickets = 50

var conferenceName = "Go Conference"
var remainingTickets uint = conferenceTickets
var bookings = []string{}

func main() {

	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidUserTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidUserTicketNumber {

			bookTickets(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("These are all the bookings - %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("All ticket booked out, please come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Invalid Input. First Name or Last name is too short.")
			}
			if !isValidEmail {
				fmt.Println("Invalid Email Id. Email Id should have @ sign.")
			}
			if !isValidUserTicketNumber {
				fmt.Printf("%v - Invalid Input. %v tickets are remaining.\n", userTickets, remainingTickets)
			}
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("Tickets remaining = %v / %v\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here...")
}

func getFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings { // _ is Blank Identifier - Ignoring unused variable
		names := strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// Take user input
	fmt.Print("Enter your first name - ")
	fmt.Scan(&firstName) // passing the reference &firstName and not the value firstName which is empty "".
	fmt.Print("Enter your last name - ")
	fmt.Scan(&lastName)
	fmt.Print("Enter your email - ")
	fmt.Scan(&email)
	fmt.Print("Enter number of tickets required - ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)
}
