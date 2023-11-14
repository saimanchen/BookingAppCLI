package app

import (
	"booking-app/entities"
	"fmt"
	"time"
)

func getUserFirstNameInput() string {
	var firstName string
	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)
	return firstName
}

func getUserLastNameInput() string {
	var lastName string
	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)
	return lastName
}

func getUserEmailInput() string {
	var email string
	fmt.Print("Enter your e-mail address: ")
	fmt.Scan(&email)
	return email
}

func getUserTicketAmountInput() uint {
	var userTickets uint
	fmt.Print("Enter how many tickets you want to buy: ")
	fmt.Scan(&userTickets)
	return userTickets
}

func greetUsers() {
	fmt.Println()
	fmt.Println("------------------------------------------------------------")
	fmt.Println()
	fmt.Println("############################################################")
	fmt.Println()
	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend!")
	fmt.Println()
	fmt.Println("############################################################")
	fmt.Println()
	fmt.Println("------------------------------------------------------------")
	fmt.Println()
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	userData := entities.User {
			FirstName: firstName,
			LastName: lastName,
			Email: email,
			NumberOfTickets: userTickets,
	}
		
	bookings = append(bookings, userData)
	
	fmt.Println()
	fmt.Printf(
		"Thank you %v %v for booking %v tickets. \nYou will receive a confirmation mail at %v.\n", 
		firstName,
		lastName,
		userTickets,
		email,
	)

	fmt.Println()
	fmt.Printf("REMAINING TICKETS: %v\n", remainingTickets)
	fmt.Println()
	fmt.Println("------------------------------------------------------------")
	fmt.Println()
}

func sendTickets(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var tickets = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	
	fmt.Println("############################################################")
	fmt.Println()
	fmt.Printf("Sending tickets:\n%v \nto e-mail address %v\n", tickets, email)
	fmt.Println()
	fmt.Println("############################################################")
	fmt.Println()
	fmt.Println("------------------------------------------------------------")
	fmt.Println()
	wg.Done()
}