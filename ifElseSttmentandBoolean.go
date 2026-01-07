package main

import (
	"fmt"
	"strings"
)

func ifElseSttment() {

	conferenceName := "Go  confrence"
	const conferenceTicket uint = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("Welcome to %v our confrence booking application \n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available. \n", conferenceTicket, remainingTickets)
	fmt.Println("Get your ticket  here to attend ")

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

		// check  if the remaining tickets are availabe to the users demand
		if userTickets <= remainingTickets {
			// updating the remainig tickets
			remainingTickets = remainingTickets - userTickets

			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you  for booking ticket with us  \n first name  %v  \n last name	%v  \n email %v \n with  booked  %v tickets. \n", firstName, lastName, email, userTickets)

			fmt.Println("We have ", remainingTickets, "tickets remaining ")

			firstNames := []string{}

			for _, booking := range bookings {
				var names = strings.Fields(booking)

				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("The first name of bookings are : %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our confrence is booked out. Come back next year. ")

				break
			}
		} else {
			fmt.Printf("We only have %v tickets remaining , so you cnnot book  %v tickets \n", remainingTickets, userTickets)

		}

	}

}
