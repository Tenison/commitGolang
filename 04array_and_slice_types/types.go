package main

import (
	"fmt"
)

func main() {
	//arrays -  static list //reserves memory for entire array which is not efficient.
	//declaring an array with SIZE and DATATYPE stored in the array
	var new_array [5]string //empty array

	second_array := [10]int{5, 10, 8} // array with values

	// print type of array
	fmt.Printf("Array type is : %T", second_array)

	//add elements to array
	new_array[0] = "Cars"

	//using loops
	//len gives the lenght of the array
	for i := 1; i < len(new_array); i++ {
		new_array[i] = "hi"
	}

	//output
	fmt.Printf("Array type : %T\n", new_array)
	fmt.Printf("Array lenght : %v\n", len(new_array))
	fmt.Printf("list array : %v\n", new_array)
	fmt.Printf("list array : %v\n", second_array)
	fmt.Printf("Array index 2: %v\n", second_array[1])
	fmt.Printf("Array index 5 : %v\n", new_array[4])

	//Slice  - dynamic list //dynamically assign values into memory making slices more efficient.
	var new_slice []string //empty slice

	myslice2 := []string{"Go", "Slices", "Are", "Powerful"} //slice with values

	//adding values to slice
	new_slice = append(new_slice, "value1")

	new_slice = append(new_slice, "value2")

	fmt.Printf("Slice Output!!!\n")
	fmt.Printf("Slice lenght : %v\n", len(new_slice))
	fmt.Printf("list slice : %v\n", new_slice)

	//cap - returns the capacity of the slice (the number of elements the slice can grow or shrink to)
	//cap show you how slice is working under the hood
	fmt.Printf("Slice capacity : %v\n", cap(myslice2))

	myslice2 = append(myslice2, "last")
	fmt.Printf("Slice capacity after adding new data: %v\n", cap(myslice2))
	fmt.Printf("list slice : %v\n", myslice2)
}
