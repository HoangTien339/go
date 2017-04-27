package main

import (
	"fmt"
	"time"
)

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			fmt.Printf("--- Assigne n = %v ---\n", n)
			out <- n
			time.Sleep(time.Second)
		}
		fmt.Println("--- Close Gen channel ---")
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			fmt.Printf("--- n * n = %v ---\n", n*n)
			out <- n * n
			fmt.Println("--- Waiting for Printing ---")
			time.Sleep(time.Second)
		}
		fmt.Println("--- Close Sq channel ---")
		close(out)
	}()
	return out
}

func main() {
	// Set up the pipeline.
	c := gen(2, 3, 4)
	out := sq(c)

	for n := range out {
		fmt.Println(n)
	}
}
