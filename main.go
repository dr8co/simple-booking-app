package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)

	var userName string
	var userTickets int

	// TODO: ask user their name and how many tickets they would like to order.

	userName = "Ian"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets", userName, userTickets)

}
