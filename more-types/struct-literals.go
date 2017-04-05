package main

import (
	"fmt"
)

type Vertex struct {
	X, Y int
}

var (
	v1 Vertex = Vertex{1, 2}  // has type Vertex
	v2        = Vertex{3, 4}  // has implicit type Vertex
	v3        = Vertex{X: 1}  // Y:0 is implicit
	v4        = Vertex{}      // X:0 and Y:0
	p         = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(v1, v2, v3, v4, p)
}
