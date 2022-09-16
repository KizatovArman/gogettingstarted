package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := x / 2
	/*
		The z^2 - x is how away z^2 is from where it needs
		to be (x), and the division by 2z is the derivative
		of z^2, to scale how much we adjust z by how quickly
		z^2 is changing. Method is Newton's method: (https://en.wikipedia.org/wiki/Newton%27s_method)
	*/
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println("Go.")
}
