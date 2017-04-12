package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	t.S = "World"
	fmt.Println(t.S)
}

func M(t T) {
	fmt.Println(t.S)
}

type F float64

func (f *F) M() {
	fmt.Println(*f)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i I

	v := T{"Hihi"}
	M(v)

	i = &T{"Hello"}
	describe(i)
	i.M()

	f := F(math.Pi)
	i = &f
	describe(i)
	i.M()
}
