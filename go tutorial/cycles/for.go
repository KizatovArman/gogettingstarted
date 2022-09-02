package main

import (
	"fmt"
)

/*
	Go has only one looping construct, the for loop

	The basic for loop has three components separated by semicolons:

		The init statement: executed before the first iteration
		The condition expression: evaluated before every iteration
		The post statement: executed at the end of every iteration

	The init statement will often be a short variable declaration,
	and the variables declared there are visible only in the scope
	of the for statement.

	The loop will stop iterating once the boolean condition evaluates
	to false.

	The init and post statements are optional. At that point you can
	drop the semicolons: C's while is spelled for in Go.

	If you omit the loop condition it loops forever, so an infinite
	loop is compactly expressed as:
		for {}

*/

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		fmt.Println("Now sum is", sum, "and i is", i)
		sum += i
	}
	fmt.Println(sum)

	sum1 := 1
	for sum1 < 1000 {
		sum1 += sum1
	}
	fmt.Println(sum1)
}
