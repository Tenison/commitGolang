package main

import (
	"fmt"
	"time"
)

func main() {
	//Simple Switch
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	//Simple Switch
	//STAGE 2
	gg := 2
	fmt.Print("Write ", i, " as ")
	switch gg {
	case 1:
		fmt.Println("one")
	case 2, 3:
		fmt.Println("two or three")
	case 4:
		fmt.Println("four")
	}

	//Having multiple cases :time.Saturday, time.Sunday
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	//Even and Odd switch
	// change number to anything you want
	number := 2
	switch {
	case number%2 == 0:
		println("Even number : ", number)
	default:
		println("Odd number : ", number)

	}

	//Cases with conditions
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
}
