package main

import "fmt"

func main() {

	conferenceName := "Go  confrence"
	const conferenceTicket int = 50

	var remainingTickets int = 50

	fmt.Printf("confrenceTicket is %T,  remaining tickets is %T,  confrencece Name  is %T\n", conferenceTicket, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v our confrence booking application \n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available. \n", conferenceTicket, remainingTickets)
	fmt.Println("Get your ticket  here to attend ")

	var userName string
	var userTickets int
	// ask user for their name

	userName = "Tom"
	userTickets = 2

	fmt.Printf("User %v booked  %v tickets. \n", userName, userTickets)
}
