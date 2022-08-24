package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

/*
	The var statement declares a list of variables;
	as in function arugment list, the type is last.
	A var statement can be at package or function level.
	The declaration of var can include intializers, one per variable.
	If initializer is present, the type can be omitted;

	A var declaration can include initializers, one per variable.
	If an initializer is present, the type can be omitted;
	the variable will take the type of the initializer.

	Inside a function, the := short assignment can be used
	in place of a var declaration with implicit type.
	Outside a function, every statement begins with a keyword
	(var, func, etc.) and so the := construct is not available.

	Go's basic types are:

	1) bool
	2) string
	3) int  int8  int16  int32  int64
	   uint uint8 uint16 uint32 uint64 uintptr
	4) byte alias for uint8
	5) rune alias for uint32 (represents a Unicode code point)
	6) float32 float64
	7) complex64 complex128

	Varialble declaration may be "factored" into blocks,
	as with import statements.

	The int, uint, and uintptr types are usually 32 bits wide
	on 32-bit systems and 64 bits wide on 64-bit systems.
	When you need an integer value you should use int unless
	you have a specific reason to use a sized or unsigned integer type.

	Variables declared without an explicit initial value
	are given their zero value. The zero value is:
	0 for numeric types,
	false for boolean type,
	"" (the empty string) for strings.const

	When declaring a variable without specifying an explicit type,
	the variables type is inferred from the value on the right hand side.
	When the right hand side contains an untyped numeric constnt, the new
	variable may be (for example) an int, float64 or complex128 depending
	on the precision of the constant:
	i := 42 // int
	f := 3.142 // float64
	g := 0.867 + 0.5i //complex128
*/

var c, python, java bool
var i1, j int = 1, 2

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	var i int
	fmt.Println(i, c, python, java)

	var c1, python1, java1 = true, false, "no!"
	fmt.Println(i1, j, c1, python1, java1)

	k := 3
	fmt.Println(k)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	var i2 int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i2, f, b, s)

	var x, y int = 3, 4
	var f1 float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f1)
	fmt.Println(x, y, z)

	v := (1+5i)*(1+5i) - 10i
	fmt.Printf("v %v is of type %T\n", v, v)
}
