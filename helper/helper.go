package helper

import (
	"strings"
)

func ValidateFirstName(firstName string) bool {
	isFirstNameValid := len(firstName) >= 2
	return isFirstNameValid
}

func ValidateLastName(lastName string) bool {
	isLastNameValid := len(lastName) >= 2
	return isLastNameValid
}

func ValidateEmail(email string) bool {
	isEmailValid := strings.Contains(email, "@")
	return isEmailValid
}

func ValidateTicketAmount(userTickets uint, remainingTickets uint) bool {
	isTicketAmountValid := remainingTickets >= userTickets && userTickets > 0
	return isTicketAmountValid
}