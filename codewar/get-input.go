package main

import (
	"fmt"
)

func main() {
	var a []int
	for {
		var n int
		var c byte
		fmt.Scanf("%d%c", &n, &c)
		a = append(a, n)
		if c == '\n' {
			break
		}
	}
	fmt.Println(a)
	/*
		var n int
		fmt.Scanf("%d", &n)
		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scanf("%d", &a[i])
		}
		fmt.Println(a)
	*/
}
