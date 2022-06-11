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

	//default foalt64
	c := 64.7

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}

func DeclareComplex() {
	var a complex64 = 1 + 3i //float32 + float32
	var b complex128         ////float64 + float64
	c := complex(2, 3)

	fmt.Printf("%v %v %T", real(a), imag(a), a)
	fmt.Println(b)
	fmt.Println(c)
}

func DeclareString() {
	//Types of string declaration
	var a string = "Hello"

	b := "Hii"

	var c string
	c = "$#"

	//String is a series of charaters
	//Each charater is represented with a byte or sometimes bytes depending on the charater.
	//Each byte or bytes has a numerical value(s)(Integer Number(s)) or code point value.
	//code point value  -> numerical representation of a charater in a string Eg:  d => 100
	//GO uses UTF8 encoding to transform/encode these charaters into binary or sequence of binary with is then stored in memory
	//Because Computers understand data in o's and 1's.
	//A UFT8 charaters' code point value can range between 1 and 4 bytes. This is because there are a lot of charaters and symbols for different languages.

	//accessing a string index or charater gives the numerical value of the charater
	fmt.Printf("%v \n", a[0])
	fmt.Printf("%v \n", c[0])

	//Access Charater in a string
	fmt.Printf("%v \n", string(b[0]))

	//len() fucntion gives the number of bytes(1 to 4 [UTF8]) that make up the code point value for a charater
	fmt.Printf("%v \n", len(string(c[1])))
}

func TypeCasting() {
	a := 4
	var b int8 = 7

	///type converting 'b'
	c := int(b)

	//d := a + b //won't work because a is INT and b is INT8
	d := a + c

	fmt.Println(d)
}
