package main

import (
	"fmt"
	"strings"
)

func main() {

	conferenceName := "Go  confrence"
	const conferenceTicket uint = 50
	var remainingTickets uint = 50
	bookings := []string{}

	greetUsers(conferenceName, conferenceTicket, remainingTickets)

	for {
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

		// inpuut vlidation
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEamil := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets < remainingTickets

		// check  if the remaining tickets are availabe to the users demand
		if isValidName && isValidEamil && isValidTicketNumber {
			// updating the remainig tickets
			remainingTickets = remainingTickets - userTickets

			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you  for booking ticket with us  \n first name  %v  \n last name	%v  \n email %v \n with  booked  %v tickets. \n", firstName, lastName, email, userTickets)

			fmt.Println("We have ", remainingTickets, "tickets remaining ")

			//call the function  prnt firstnames
			printFirstNames(bookings)

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

func greetUsers(confName string, confTicket uint, remTickets uint) {
	fmt.Printf("Welcome to %v booking application \n", confName)

	fmt.Printf("We have a total of %v tickets and %v are still available. \n", confTicket, remTickets)
	fmt.Println("Get your ticket  here to attend ")
}

func printFirstNames(bookings []string) {
	firstNames := []string{}

	for _, booking := range bookings {
		var names = strings.Fields(booking)

		firstNames = append(firstNames, names[0])
	}

	fmt.Printf("The first name of bookings are : %v\n", firstNames)
}
