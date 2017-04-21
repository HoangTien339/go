package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var c int
	for i := 0; i < 10000000; i++ {
		go func() {
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
			c++
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Println("Value of c:", c)
	time.Sleep(1 * time.Second)
}
