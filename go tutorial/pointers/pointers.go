package main

import "fmt"

/*
	Go has pointers. A pointer holds the memory address
	of a value.
	The type *T is a pointer to a T value. Its zero value
	is nil.
	The & operator generates a pointer to its operand.
	i:=42
	p = &i
	The * operator denotes the pointer's underlying value.
	fmt.Println(*p)
	*p = 21
	This is known as "dereferencing" or "indirecting".
*/

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i trhough the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p /= 37       // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
