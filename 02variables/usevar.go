package main

import "fmt"

//Variables in go
//string
//int [int, int8, int16, int32, int64 AND uint8, uint16, uint32, uint64]
//float [float32, float64, AND complex64, complex128]
//bool
//arrays
//map

func main() {
	//variable assignment
	var myname = "osei owusu"
	var boolean = true

	//More Options
	var number int = 5 //declared but not used!!!
	sentence := "hii"

	//constant
	const ad = "hold this value forever"

	fmt.Printf("type: %T \n", boolean)

	fmt.Println("My is ", myname)

	//define a varible.
	var holdValue int8

	//assigning the variable
	holdValue = 4
	fmt.Printf("the variable contains : %d", holdValue)

}
