package main

import (
	"fmt"
	"net/mail"
	"strings"
)

func main() {
	const conference_tickets = 50

	var user_email_addr string
	var user_name_first string
	var user_name_last string
	var user_tickets int

	bookings := []string{}
	conference_name := "Intercogni Monthly"
	remaining_tickets := 50

	// user_name = "Taib"
	// user_tickets = 1

	fmt.Printf("Sign up for %v for IoT and Robotics\n", conference_name)
	fmt.Printf(
		"%v are given this month and %v are still available\n",
		conference_tickets, remaining_tickets)
	fmt.Printf("Get your tickets here to attend\n\n")

	for remaining_tickets > 0 && len(bookings) < 50 {
	__input_name:
		fmt.Printf("Please enter your first name: ")
		fmt.Scan(&user_name_first)
		fmt.Printf("Please enter your last name: ")
		fmt.Scan(&user_name_last)

		stat_valid_user_name := len(user_name_first) >= 2 && len(user_name_last) >= 2
		if !stat_valid_user_name {
			fmt.Printf("!!! You seem to have entered an invalid name, please try again !!!\n")
			goto __input_name
		}
		bookings = append(bookings, user_name_first+" "+user_name_last)

	__input_email:
		fmt.Printf("Please enter your email address: ")
		fmt.Scan(&user_email_addr)

		_, err := mail.ParseAddress(user_email_addr)
		if err != nil {
			fmt.Printf("!!! You seem to have entered an invalid email address, please try again !!!\n")
			goto __input_email
		}

	__book_tickets:
		fmt.Printf("Please enter the number of tickets to book: ")
		fmt.Scan(&user_tickets)

		if user_tickets > remaining_tickets {
			fmt.Printf(
				"!!! Only %v tickets remain, please try booking with less tickets !!!\n",
				remaining_tickets)
			goto __book_tickets
		}
		remaining_tickets -= user_tickets

		fmt.Printf(
			"\nSuccessfully booked %v ticket/s for %v. Confirmation sent to %v \n",
			user_tickets, bookings[len(bookings)-1], user_email_addr)

		first_names := []string{}
		for _, booking := range bookings {
			this_first_name := strings.Fields(booking)[0]
			first_names = append(first_names, this_first_name)
		}

		fmt.Printf("----------\n")
		fmt.Printf("booking reservations: %v\n", first_names)
		fmt.Printf(
			"%v tickets now remains for %v\n",
			remaining_tickets, conference_name)
		fmt.Printf("----------\n\n")

		stat_tickets_unavailable := remaining_tickets == 0
		if stat_tickets_unavailable {
			fmt.Printf(
				"This month's %v is booked out. Please come again next month!\n",
				conference_name)
			break
		}
	}
}
