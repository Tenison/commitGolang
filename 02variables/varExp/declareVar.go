package varExp

import "fmt"

//package level variables
var pkgVar int
var pkgVarTwo float32 = 3

var (
	pkgGroupVar      int     = 10
	pkgGroupVarTwo   string  = "Hello World"
	pkgGroupVarThree float64 = 34.2
)

func DeclareFunc() {
	//block level variable
	var funcVar int //defaul value 0
	var funcVarTwo float32 = 3
	funcVarThree := "Hii"

	//initialize variable funcVar
	//funcVar = 20

	fmt.Printf("%v %v %v \n", funcVar, funcVarTwo, funcVarThree)

	//Finding out the type
	fmt.Printf("%T %T %T", funcVar, funcVarTwo, funcVarThree)

}

func PrintPkgLevelVarFunc() {
	fmt.Printf("%v %v \n", pkgVar, pkgVarTwo)
	fmt.Printf("%v %v %v", pkgGroupVar, pkgGroupVarTwo, pkgGroupVarThree)

}
