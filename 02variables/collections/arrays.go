package collections

import "fmt"

//UniArray := [...]int{1,2,4}
var packageArray [5]float32

//array declaration
func DeclareArray (){
	var newArray [6]int //Declare basic array of size 6
	var newArrayRedudent [4]float32 = [4]float32{34.3, 56.6, 40.3, 12}
	newArray1 := [6]string{"Osei", "Yaw", "Alex", "Nash", "Michael", "Frank"} //Array Literal
	newArray2 := [...]string{"Osei", "Yaw", "Michael", "Frank"} ///Initize array to whatever value size that is passed in

	newArray[0] = 1
	newArray[1] = 2
	newArray[2] = 3 
	
	fmt.Printf("%v", newArrayRedudent)
	fmt.Printf("%v", newArray[0])
	fmt.Printf("%v", newArray1)
	fmt.Printf("%v", newArray2)
}

func ArrayLen(){
	newArray2 := [...]string{"Osei", "Yaw", "Michael", "Frank"} ///Initize array to whatever value size that is passed in

	lenght := len(newArray2)

	fmt.Println("Array lenght is ", lenght)
}

func MultiDimensional(){
	var multiDimensionalArray [3][2]int = [3][2]int{[2]int{0,0}, [2]int{1,3}, [2]int{1,5}}
	
	//another way.. much better
	var multiDimensionalArrayTwo [3][2]int
	
	multiDimensionalArrayTwo[0] = [2]int{0,0}
	multiDimensionalArrayTwo[1] = [2]int{1,3}
	multiDimensionalArrayTwo[2] = [2]int{1,5}

	fmt.Println("Array lenght is ", multiDimensionalArray)
}

