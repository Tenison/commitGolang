package main

import (
	"fmt"
	"strings"
)

func main() {
	var firstName string
	var lastName string
	var email string
	var userTickets int8
	userFullname := []string{}

	for {
		fmt.Print("Please enter your first name : ")
		fmt.Scan(&firstName)

		fmt.Print("Please enter your last name : ")
		fmt.Scan(&lastName)

		fmt.Print("Email Address : ")
		fmt.Scan(&email)

		fmt.Print("Please enter number of ticket : ")
		fmt.Scan(&userTickets)

		userFullname = append(userFullname, firstName+" "+lastName)

		if userTickets > 5 {
			break
		}
	}

	for i := 0; i < len(userFullname); i++ {
		fristnameOnly := strings.Fields(userFullname[i]) // divide sting in array
		fmt.Printf("Thankyou %s, you purhased %d tickets, confirmation will be sent to your email at %s\n", fristnameOnly[0], userTickets, email)
	}

}
