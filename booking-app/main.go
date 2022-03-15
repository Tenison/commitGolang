package main

import (
	"fmt"
	"strings"
)

func main() {
	maxTickets := 50
	var firstName string
	var lastName string
	var email string
	var userTickets int8
	userFullname := []string{}
	var emails []string
	ticketPerUser := []int8{}

	for {
		fmt.Print("Please enter your first name : ")
		fmt.Scan(&firstName)

		fmt.Print("Please enter your last name : ")
		fmt.Scan(&lastName)

		fmt.Print("Email Address : ")
		fmt.Scan(&email)

		fmt.Print("Please enter number of ticket : ")
		fmt.Scan(&userTickets)

		maxTickets -= int(userTickets)

		userFullname = append(userFullname, firstName+" "+lastName)
		emails = append(emails, email)
		ticketPerUser = append(ticketPerUser, userTickets)

		if maxTickets <= 0 {
			break
		}
	}

	for index, userName := range userFullname {
		fristnameOnly := strings.Fields(userName) // divide sting in array
		fmt.Printf("Thankyou %s, you purhased %d tickets, confirmation will be sent to your email at %s\n", fristnameOnly[0], ticketPerUser[index], emails[index])
	}

}
