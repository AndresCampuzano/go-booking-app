package helper

// IsAlpha Function to check if a string contains only letters
func IsAlpha(s string) bool {
	for _, char := range s {
		if !('a' <= char && char <= 'z') && !('A' <= char && char <= 'Z') && char != ' ' {
			return false
		}
	}
	return true
}

func ValidateUserInput(userName string) bool {
	if len(userName) > 2 && IsAlpha(userName) {
		return true
	}
	return false
}

func ValidateTicketsInput(userTickets, remainingTickers uint) bool {
	if userTickets == 0 || userTickets <= remainingTickers {
		return true
	}
	return false
}
