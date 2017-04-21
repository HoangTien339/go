package main

import (
	"fmt"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	var n, a, b int
	var times []int

	fmt.Scanf("%d", &n)
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)

	times = make([]int, n+1)
	times[0] = 0
	times[1] = a
	for i := 2; i <= n; i++ {
		times[i] = times[i-1] + a
		if i%2 == 0 {
			times[i] = min(times[i], times[i/2]+b)
		} else {
			times[i] = min(times[i], times[i/2+1]+a+b)
		}
	}
	fmt.Printf("%d", times[n])
}
