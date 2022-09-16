package main

import "fmt"

/*
	A struct is a collection of fields.

	Struct fields are accessed using a dot.

	Struct fields can be accessed through
	a struct pointer.
	To access the field X of a struct when
	we have the struct pointer p we could
	write (*p).X. However, that notation is
	cumbersome, so the language permits us
	instead to write just p.X, without the
	explicit deference.

	A struct literal denotes a newly allocated
	struct value by listing the values of its
	fields.
	You can list just a subset of fields by
	using the Name: syntax.(Order is irrelevant)

	The special prefix & returns a pointer to
	the struct value.
*/

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2} // has type Vertex
	v2 = Vertex{X: 1} // Y:0 is implicit
	v3 = Vertex{}     // X:0 and Y:0
	p1 = &Vertex{1, 2}
)

func main() {
	v := Vertex{1, 2}
	v.X = 4
	p := &v
	p.Y = 5
	// fmt.Println(v.X, v.Y)
	fmt.Println(v1, p1, v2, v3)
}
