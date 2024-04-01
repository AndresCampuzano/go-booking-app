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
		fmt.Scan(&userName)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		if userTickets > remainingTickers {
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
