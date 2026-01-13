package main

import "fmt"

func switchStatment() {
	city := "London"

	switch city {
	case "New York":
		//Execute code for booking New York conference ticket
	case "Singapore", "Hong Kong":
		// Execute code for booking
	case "London", "Berlin":
		// Execute code for booking
	case "Mexico city":
		// Execute code for booking

	default:
		fmt.Print("No valid city selected ")
	}
}
