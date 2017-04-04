package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	z_next := float64(0)
	delta := float64(1)
	for i := 0; delta > 0.001 || delta != 0; i++ {
		// Implement the square root function using Newton's method.
		z_next = z - (math.Pow(z, 2)-x)/(2*z)
		delta = math.Abs(z - z_next)
		z = z_next
	}
	return z_next
}

func main() {
	fmt.Println(Sqrt(5))
}
