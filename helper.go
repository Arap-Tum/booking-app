package main

import "strings"

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	// inpuut vlidation
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEamil := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets < remainingTickets
	return isValidName, isValidEamil, isValidTicketNumber
}
