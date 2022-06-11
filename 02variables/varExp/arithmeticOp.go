package varExp

import "fmt"

func SimpleOperations() {
	a, b := 10, 3

	c := a + b
	d := a - b
	e := a / b
	f := a * b
	g := a % b

	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

}

func ShiftOperations() {
	// bit shifting
	a := 8  // = 2^3 or 2 pow 3
	b := 16 // = 2^4 or 2 pow 4

	fmt.Println(a << 3) // (2^3) * (2^3) = (2^6)// Ans = 64
	fmt.Println(a >> 2) // (2^3) / (2^2) = (2^1)// Ans = 2

	fmt.Println(b << 3) // (2^4) * (2^3) = (2^7)// Ans = 128
	fmt.Println(b >> 2) // (2^4) / (2^2) = (2^2)// Ans = 4
}

func BitOperations() {
	a := 10 //1010
	b := 3  //0011

	//these operations are carried out using binary value
	c := a & b //AND // 2 //0010
	d := a | b //OR // 11 //1011
	e := a ^ b //XOR // 9 //1001 //means one of the numbers has (bit set or not) but not both.

	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

}
