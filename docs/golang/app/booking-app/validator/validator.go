package validator

import (
	"strings"
)

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) > 2 && len(lastName) > 3
	isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
	isVaildTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isVaildTicketNumber
}
