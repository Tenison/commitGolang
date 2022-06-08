package varExp

import (
	"fmt"
)

func DeclareBool() {
	var a bool //default value false
	var b bool = true
	c := true

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

//Signed VS Unsigned
//Signed - positive and negative values
//Unsigned - positive values only

func DeclareInt() {
	var a int = 32
	var b int8 = 8
	var c int16 = 16
	var d int32 = 32
	var e int64 = 64

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

}

func DeclareFloat() {
	var a float32 = 32
	var b float64

	fmt.Println(a)
	fmt.Println(b)

}
