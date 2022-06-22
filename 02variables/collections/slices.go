package collections

import "fmt"


func DeclareSliceAndAppend(){
	//make([]type, size, capacity) //capacity is optional parameter
	newSliceOne := make([]int, 3)
	newSliceTwo := make([]int, 4, 50)

	var newSliceThree []int = []int{}

	//Adding elements to slice
	newSliceOne =  append(newSliceOne, 1)
	newSliceTwo = append(newSliceTwo, 1, 2, 3, 4)

	//Adding two slices together
	newSliceThree = append(newSliceOne, newSliceTwo...)	

	//special string combine
	newSliceFour := append([]byte("Hello"), " world"...)

	fmt.Println(newSliceThree)
	fmt.Println(string(newSliceFour))//bytes are represented in a series of int.

}