package app

import (
	"booking-app/entities"
	"booking-app/helper"
	"fmt"
	"sync"
)

var conferenceName = "Go Conference"
const conferenceTickets = 50
var remainingTickets uint = 50
var bookings = make([]entities.User, 0)

var wg = sync.WaitGroup {}

func App() {
	greetUsers()

	firstName := getUserFirstNameInput()
	isFirstNameValid := helper.ValidateFirstName(firstName)
	for !isFirstNameValid {
		fmt.Println("Your first name is to short!")
		firstName = getUserFirstNameInput()
		isFirstNameValid = helper.ValidateFirstName(firstName)
	}

	lastName := getUserLastNameInput()
	isLastNameValid := helper.ValidateLastName(lastName)
	for !isLastNameValid {
		fmt.Println("Your last name is to short!")
		lastName = getUserLastNameInput()
		isLastNameValid = helper.ValidateLastName(lastName)
	}

	email := getUserEmailInput()
	isEmailValid := helper.ValidateEmail(email)
	for !isEmailValid {
		fmt.Println("Your e-mail isn't valid!")
		email = getUserEmailInput()
		isEmailValid = helper.ValidateEmail(email)
	}

	userTickets := getUserTicketAmountInput()
	isTicketAmountValid := helper.ValidateTicketAmount(userTickets, remainingTickets)
	for !isTicketAmountValid {
		if (userTickets > remainingTickets) {
			fmt.Printf("We only have %v tickets left!\n", remainingTickets)
		}

		if (userTickets == 0) {
			fmt.Println("You can't buy 0 tickets!")
		}

		userTickets = getUserTicketAmountInput()
		isTicketAmountValid = helper.ValidateTicketAmount(userTickets, remainingTickets)
	}

	if isFirstNameValid && isLastNameValid && isEmailValid && isTicketAmountValid {
		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTickets(userTickets, firstName, lastName, email)

		if remainingTickets == 0 {
			fmt.Println("Our conference is booked out. Come back next year!")
		}
	}
	wg.Wait()
}