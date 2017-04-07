package main

import (
	"fmt"
)

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	f0, f1, fn := 0, 1, 0
	return func() int {
		fn, f0, f1 = f0, f1, f0+f1
		return fn
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
