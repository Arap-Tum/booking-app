package main

import "fmt"

func ForLOOP() {

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

		// updating the remainig tickets
		remainingTickets = conferenceTicket - userTickets

		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you  for booking ticket with us  \n first name  %v  \n last name	%v  \n email %v \n with  booked  %v tickets. \n", firstName, lastName, email, userTickets)

		fmt.Println("We have ", remainingTickets, "tickets remaining ")

		fmt.Printf("These are all the bookings : %v\n", bookings)
	}

}
