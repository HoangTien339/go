package main

import (
	"fmt"
	"math/cmplx"
)

var (
	hex    int8       = 12
	ToBe   bool       = false
	MaxInt uint64     = 6 << 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	const f = "%T(%v)\n"
	fmt.Printf(f, hex, hex)
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
}
