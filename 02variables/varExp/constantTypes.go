package varExp

import (
	"fmt"
)

//const don't use := notation in assigning integers

const MyFirstConst int = 5       //exported constant
const mySecondConst float32 = 10 // availiable within the package//Package level scope

//using iota counters
const (
	a = iota
	b = iota
	c = iota
)

//Same as Before(Above)
const (
	d = iota
	f
	g
)

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

	///Next Set
	fmt.Printf("%v \n", d)
	fmt.Printf("%v \n", f)
	fmt.Printf("%v \n", g)
}

const (
	isAdmin = 1 << iota //1
	isHeadquarters      //2
	canSeeFinancials    //4
	canSeePayment       //8
)

func IotaOpe() {
	fmt.Printf("%v \n", isAdmin)
	fmt.Printf("%v \n", isHeadquarters)
	fmt.Printf("%v \n", canSeeFinancials)
	fmt.Printf("%v \n", canSeePayment)
}

const (
	_ = iota + 3 //write only(blank identifier )
	o
	s
	e
)

func IotaOpe1() {
	fmt.Printf("%v \n", o)
	fmt.Printf("%v \n", s)
	fmt.Printf("%v \n", e)
}

const (
	_ = iota      //write only(blank identifier )
	x = 10 - iota //start here
	y
	z
)

func IotaOpe2() {
	fmt.Printf("%v \n", x)
	fmt.Printf("%v \n", y)
	fmt.Printf("%v \n", z)
}
