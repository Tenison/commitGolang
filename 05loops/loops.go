package main

import (
	"fmt"
)

//loops in go
func main() {
	//loop basic form
	max := 5
	for i := 0; i < max; i++ {
		fmt.Println("Hello", i)
	}

	//range iterates over elements in a variety of data structures
	//Example : Add the series of 10 number
	var total int8
	numberList := []int8{12, 3, 45, 2, 5, 65, 12, 43, 54, 10}

	for _, number := range numberList {
		total += number
	}
	fmt.Println("Total of 10 numbers is : ", total)

	//forever loop
	for {
		max++
		if max == 10 {
			break
		}
	}
}
