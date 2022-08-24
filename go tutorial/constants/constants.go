package main

import (
	"fmt"
)

/*
	Constants are declared like variables, but with const keyword.
	Constants can be character, string, boolean, or numeric values.
	Constants can't be declareed using the := syntax.

	Numeric constants are high-precision values
	An untyped constant takes the type needed by its context.
	(An int can store at maximum a 64-bit integer, and sometimes less.)
*/

const (
	Big   = 1 << 100
	Small = Big >> 99
	Pi    = 3.14
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	const World = "World"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go Rules?", Truth)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
