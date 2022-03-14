package main

import "fmt"

//loops in go
func main() {
	max := 5

	for i := 0; i < max; i++ {
		fmt.Println("Hello", i)
	}

	for {
		max++
		if max == 10 {
			break
		}
	}
}
