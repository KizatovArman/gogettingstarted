package main

import "fmt"

/*
	A function can take 0 or more arguments.
	Type comes after the variable name.
	When 2 or more consecutive named function paramenters
	share a type, you can omit the type from all but last.
	A fucntion can return any number of results.
	Go's return values may be named. If so, they are treated as
	variables defined at the top of the function.
	These names should be used to document the meaning of the
	return values. A return statement without arguments returns
	the named return values. This is known as a "naked" return.
	Naked return statements should be used only in short functions,
	as with the example shown in split function. They can harm
	readability in longer functions.
*/

func add(x int, y int) int {
	var sum = x + y
	return sum
}

func addTwo(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(85))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(add(42, 13))
	fmt.Println("With new parameters result", addTwo(5, 6))
}
