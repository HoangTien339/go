package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a, b Abser
	f := MyFloat(-math.Sqrt2)
	f2 := MyFloat2(-math.Sqrt2)

	a = &f

	b = f2

	fmt.Println(a.Abs())
	fmt.Println(b.Abs())
}

type MyFloat float64

func (f *MyFloat) Abs() float64 {
	if *f < 0 {
		return float64(-(*f))
	}
	return float64(*f)
}

type MyFloat2 float64

func (f MyFloat2) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
