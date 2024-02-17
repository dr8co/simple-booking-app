package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceTickets = 50

var (
	conferenceName        = "Go Conference"
	remainingTickets uint = 50
	bookings              = make([]UserData, 0)
)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {
	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)

			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			switch {
			case !isValidName:
				sendUnvalidatedUserInputMessage(1)
			case !isValidEmail:
				sendUnvalidatedUserInputMessage(2)
			case !isValidTicketNumber:
				sendUnvalidatedUserInputMessage(3)
			}
		}
	}
	wg.Wait()
}

func sendUnvalidatedUserInputMessage(error_code int) {
	switch error_code {
	case 1:
		fmt.Println("The first name or last name you entered is too short.")
	case 2:
		fmt.Println("The email address you entered doesn't contain @ sign.")
	case 3:
		fmt.Println("The number of tickets you entered is invalid.")
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\nWe have a total of %v tickets and %v are still available.\n", conferenceName, conferenceTickets, remainingTickets)
}

func getFirstNames() []string {
	var firstNames []string
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var (
		firstName   string
		lastName    string
		email       string
		userTickets uint
	)

	// Ask the user their info and how many tickets they'd like to order.

	fmt.Println("Enter your first name: ")
	_, err := fmt.Scan(&firstName)
	if err != nil {
		return "", "", "", 0
	}

	fmt.Println("Enter your last name: ")
	_, err = fmt.Scan(&lastName)
	if err != nil {
		return "", "", "", 0
	}

	fmt.Println("Enter your email address: ")
	_, err = fmt.Scan(&email)
	if err != nil {
		return "", "", "", 0
	}

	fmt.Println("Enter number of tickets: ")
	_, err = fmt.Scan(&userTickets)
	if err != nil {
		return "", "", "", 0
	}

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets -= userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(50 * time.Second) // Simulate delay
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)

	fmt.Println("###################")
	fmt.Printf("Sending ticket:\n %v \n to email address %v\n", ticket, email)
	fmt.Println("###################")

	wg.Done()
}
