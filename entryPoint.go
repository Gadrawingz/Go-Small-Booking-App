package main

import "fmt"

func main() {
	const meetingName = "Go Meeting!"
	var meetingYear = 2023
	fmt.Printf("Welcome to %v in  %v  \n", meetingName, meetingYear)
	// Syntactical sugar of declaring variables in Go.
	meetingAliasName := "Go-Lang Meeting 2k23"

	fmt.Printf("The meeting is also known as %v \n", meetingAliasName)

	// Pointer
	var bookingTickets = 40
	fmt.Println(&bookingTickets) // Memory address!

	// Variables.
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)
	fmt.Println("Enter a number of tickets you want: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thanks for booking %v tickets. You\" ll soon receive message at %v. \n", userTickets, email)
	fmt.Printf("The remaing tickets are %v ", remainingTickets)

	// Bookig
	var newFirstName = "Peter"
	var newLastName = "Tosh"
	var bookings = [40]string{"Aki", "Nana", "Jack", "Kim"}
	bookings[0] = newFirstName + " " + newLastName

	fmt.Printf("The whole array is %v \n", bookings)
	fmt.Printf("The first name is %v \n", bookings[0])
	fmt.Printf("The Array length is %v \n", len(bookings))
}
