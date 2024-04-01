package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickers uint = 50
	bookings := []string{}

	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickers)
	fmt.Println("Get your tickets here to attend.")

	fmt.Println()

	for {
		var userName string
		var userTickets uint

		fmt.Println("Enter your first name: ")
		scan, err := fmt.Scan(&userName)
		if err != nil {
			if scan == 0 {
				fmt.Println("Error reading input.")
				return
			}
			return
		} else if len(userName) < 2 || !IsAlpha(userName) {
			fmt.Println("Invalid name. Please enter a valid name.")
			continue
		}

		fmt.Println("Enter number of tickets: ")
		scan, err = fmt.Scan(&userTickets)
		if err != nil {
			if scan == 0 {
				fmt.Println("Error reading input.")
				return
			}
			return
		} else if userTickets == 0 || userTickets > remainingTickers {
			fmt.Printf("There are only %v tickets available, so you can't book %v tickets.\n", remainingTickers, userTickets)
			continue
		}

		remainingTickers = remainingTickers - userTickets
		bookings = append(bookings, userName)

		fmt.Printf("Thank you %v for booking %v tickets.\n", userName, userTickets)
		fmt.Printf("There are %v tickets remaining.\n", remainingTickers)

		fmt.Printf("All bookings: %v\n", bookings)

		if remainingTickers == 0 {
			fmt.Println("All tickets have been booked.")
			break
		}
	}
}
