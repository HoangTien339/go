package main

import (
	"fmt"
	"sync"
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

func sq(in <-chan int, chanNum int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			fmt.Printf("[%v] --- n * n = %v ---\n", chanNum, n*n)
			out <- n * n
			fmt.Printf("[%v] --- Waiting for Printing ---\n", chanNum)
			time.Sleep(time.Second)
		}
		fmt.Printf("[%v] --- Close Sq channel ---\n", chanNum)
		close(out)
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Start an output goroutine for each input channel in cs.
	// Output copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan int, out chan<- int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c, out)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	// Set up the pipeline.
	in := gen(2, 3, 4, 5, 6, 7, 8, 9)

	// Distribute the sq work across two goroutines that both read from in.
	c1 := sq(in, 1)
	c2 := sq(in, 2)

	// Consume the merged output from c1 and c2.
	for n := range merge(c1, c2) {
		fmt.Println(n) // 4 then 9, or 9 then 4
	}
}
