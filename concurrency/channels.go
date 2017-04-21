package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
	c <- sum
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[len(s)/2:], c)          // SUM[-9 4 0] = -5
	go sum(s[:len(s)/2], c)          // SUM[7 2 8] = 17
	go sum(s[len(s)/3:], c)          // SUM[8 -9 4 0] = 3
	go sum(s[len(s)/6:], c)          // SUM[2 8 -9 4 0] = 5
	x, y, z, t := <-c, <-c, <-c, <-c // receive from c
	fmt.Println(x, y, z, t)
}
