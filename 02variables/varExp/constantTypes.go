package varExp

import (
	"fmt"
)

//const don't use := notation in assigning integers

const (
	a = iota
	b = iota
	c = iota
)

const MyFirstConst int = 5       //exported constant
const mySecondConst float32 = 10 // availiable within the package//Package level scope

func DeclareConst() {
	const myThirdConst int = 5

	fmt.Printf("%v %T \n", MyFirstConst, MyFirstConst)
	fmt.Printf("%v %T \n", mySecondConst, mySecondConst)
	fmt.Printf("sum is %v \n", myThirdConst+MyFirstConst)
}

func PrintIota() {
	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
}
