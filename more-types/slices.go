package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes[2])
	var s []int = primes[0:1]
	fmt.Println(s)
}
