package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	pow = append(pow, 256, 512)
	fmt.Printf("len=%d cap=%d %v\n", len(pow), cap(pow), pow)
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
