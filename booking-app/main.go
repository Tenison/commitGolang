package main

import "fmt"

func main() {
	var firstName string
	var lastName string
	var email string
	var userTickets int8

	fmt.Print("Please enter your first name : ")
	fmt.Scan(&firstName)

	fmt.Print("Please enter your last name : ")
	fmt.Scan(&lastName)

	fmt.Print("Email Address : ")
	fmt.Scan(&email)

	fmt.Print("Please enter number of ticket : ")
	fmt.Scan(&userTickets)

	fmt.Printf("%v %v, you purhased %d tickets, confirmation will be sent to your email at %s", firstName, lastName,  userTickets, email)

}
