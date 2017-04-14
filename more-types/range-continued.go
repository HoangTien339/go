package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
        fmt.Printf("-%d-", i)
		pow[i] = 1 << uint(i) // == 2**i
	}
    fmt.Println()

	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
