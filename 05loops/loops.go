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

	//Q2
	//use len() to find the length of the slice
	numberList := []int8{12, 3, 45, 2, 5, 65, 12, 43, 54, 10}

	for i := 0; i < len(numberList); i++ {
		fmt.Println("Hello", numberList[i])
	}

	//range iterates over elements in a variety of data structures
	//Example : Add the series of 10 number
	var total int8
	numberList2 := []int8{12, 3, 45, 2, 5, 65, 12, 43, 54, 10}

	for _, number := range numberList2 {
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
