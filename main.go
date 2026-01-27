package main

import (
	"booking-app/helper"
	"fmt"
	"time"
)

const conferenceTicket uint = 50

var conferenceName = "Go  confrence"

var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()

		// VALIDATION OF INPUT
		isValidName, isValidEamil, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		// check  if the remaining tickets are availabe to the users demand
		if isValidName && isValidEamil && isValidTicketNumber {
			//  booking ticket function goes here

			bookTicket(userTickets, firstName, lastName, email)
			go sendTicket(userTickets, firstName, lastName, email)

			//call the function  prnt firstnames
			firstNames := printFirstNames()

			fmt.Printf("The first name of bookings are : %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our confrence is booked out. Come back next year. ")

				break
			}
		} else {

			if !isValidName {
				fmt.Println("first name or last name you entered is too short ")
			}
			if !isValidEamil {
				fmt.Println("email address you enter does not contaim @ sign  ")
			}

			if !isValidTicketNumber {
				fmt.Println("numbeer of tickets you r=entered is invalid ")
			}

		}

	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application \n", conferenceName)

	fmt.Printf("We have a total of %v tickets and %v are still available. \n", conferenceTicket, remainingTickets)
	fmt.Println("Get your ticket  here to attend ")
}

func printFirstNames() []string {
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

	// ask user for their name
	fmt.Println("Enter your first name ?  ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name ? ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your  email ? ")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets ? ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	// updating the remainig tickets
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("List of bookings is %v \n", bookings)

	fmt.Printf("Thank you  for booking ticket with us  \n first name  %v  \n last name	%v  \n email %v \n with  booked  %v tickets. \n", firstName, lastName, email, userTickets)

	fmt.Println("We have ", remainingTickets, "tickets remaining ")
}

func sendTicket(userTickets uint, firstnames string, lastName string, email string) {
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets %v %v ", userTickets, firstnames, lastName)

	fmt.Println("#########################")

	fmt.Printf("Sending ticket \n %v  to email address %v \n", ticket, email)

	fmt.Println("#########################")

}
