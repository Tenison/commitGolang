package main 

import "fmt"

//main : we are going to accept input from the user
func main() {
	var name string
	var age int

	fmt.Print("Please enter your name : ")
	fmt.Scan(&name)

	fmt.Print("Please enter your age : ")
	fmt.Scan(&age)

	fmt.Printf("%v, you are %d years old", name,age)
}